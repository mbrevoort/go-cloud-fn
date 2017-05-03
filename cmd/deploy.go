package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	t "github.com/mbrevoort/go-cloud-fn/template"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy <function-name>",
	Short: "Deploy your cloud function",
	Long: `This command lets you deploy your function with a given
set of parameters.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Println("You need to provide the function name as first argument.")
			return
		}
		if !emulator && stageBucket == "" {
			log.Println("Production environment needs stage-bucket")
			return
		}
		if !triggerHTTP && triggerBucket == "" && triggerTopic == "" {
			log.Println("You need to set either trigger-http, trigger-bucket or trigger-topic")
			return
		}
		functionName := args[0]
		targetDir := "./target/" + uuid.NewV4().String() + "/"

		templateData := t.IndexTemplateData{
			FunctionName: functionName,
			TargetDir:    targetDir,
			TriggerHTTP:  triggerHTTP,
		}

		index, err := t.GenerateIndex(templateData)
		if err != nil {
			logCLIErr(err.Error())
			return
		}
		pjson, err := t.GeneratePackageJson(templateData)
		if err != nil {
			logCLIErr(err.Error())
			return
		}

		err = os.MkdirAll(targetDir, os.ModePerm)
		if err != nil {
			logCLIErr(err.Error())
			return
		}

		//Set the correct trigger (only one)
		trigger := []string{}
		if triggerHTTP {
			trigger = append(trigger, "--trigger-http")
		} else if triggerBucket != "" {
			trigger = append(trigger, "--trigger-bucket")
			trigger = append(trigger, triggerBucket)
		} else if triggerTopic != "" {
			trigger = append(trigger, "--trigger-topic")
			trigger = append(trigger, triggerTopic)
		}

		//Standard deploy arguments
		deployArguments := []string{
			"gcloud beta functions deploy",
			functionName,
			"--local-path", targetDir,
			"--timeout", timeout,
		}

		//Use functions if it's emulator we're deploying to.
		var buildCmd string
		indexPath := targetDir + "index.js"
		pjsonPath := targetDir + "package.json"
		if emulator {
			deployArguments[0] = "functions deploy"
			buildCmd = "go build -o " + targetDir + functionName
		} else {
			buildCmd = "GOOS=linux go build -o " + targetDir + functionName
			deployArguments = append(
				deployArguments,
				"--memory", memory,
				"--stage-bucket", stageBucket)
			if region != "" {
				deployArguments = append(
					deployArguments,
					"--region", region)
			}
		}

		err = ioutil.WriteFile(indexPath, []byte(index), os.ModePerm)
		if err != nil {
			logCLIErr(err.Error())
			return
		}
		err = ioutil.WriteFile(pjsonPath, []byte(pjson), os.ModePerm)
		if err != nil {
			logCLIErr(err.Error())
			return
		}

		if dotenv != "" {
			env, err := ioutil.ReadFile(dotenv)
			if err != nil {
				logCLIErr(err.Error())
				return
			}
			err = ioutil.WriteFile(targetDir+".env", env, os.ModePerm)
			if err != nil {
				logCLIErr(err.Error())
				return
			}
		}

		logCLI("Building " + functionName)
		compile := exec.Command("sh", "-c", buildCmd)
		compile.Stdout = os.Stdout
		compile.Stderr = os.Stderr
		err = compile.Run()
		if err != nil {
			logCLIErr(err.Error())
			return
		}

		logCLI("Deploying " + functionName)
		deploy := exec.Command("sh", "-c", strings.Join(append(deployArguments, trigger...), " "))
		deploy.Stdout = os.Stdout
		deploy.Stderr = os.Stderr
		err = deploy.Run()
		if err != nil {
			logCLIErr(err.Error())
			return
		}
		logCLI("Finished")
	},
}

var emulator bool
var stageBucket string
var triggerHTTP bool
var triggerBucket string
var triggerTopic string
var memory string
var timeout string
var region string
var dotenv string

func init() {
	RootCmd.AddCommand(deployCmd)
	deployCmd.Flags().BoolVarP(&emulator, "emulator", "e", false, "Deploy to emulator")
	deployCmd.Flags().StringVarP(&stageBucket, "stage-bucket", "s", "", "Set GCS bucket to upload zip bundle")
	deployCmd.Flags().BoolVarP(&triggerHTTP, "trigger-http", "j", false, "Set function to trigger on HTTP event")
	deployCmd.Flags().StringVarP(&triggerBucket, "trigger-bucket", "b", "", "Set function to trigger on this GCS bucket event")
	deployCmd.Flags().StringVarP(&triggerTopic, "trigger-topic", "t", "", "Set function to trigger on this Pubsub topic")
	deployCmd.Flags().StringVarP(&memory, "memory", "m", "1024MB", "Set function memory [MAX 2048MB]")
	deployCmd.Flags().StringVarP(&timeout, "timeout", "o", "540s", "Set function timeout [MAX 540s]")
	deployCmd.Flags().StringVarP(&region, "region", "r", "", "Set gcloud region")
	deployCmd.Flags().StringVarP(&dotenv, "dotenv", "v", "", "Pass to .env file")
}

func logCLI(line string) {
	os.Stdout.WriteString("λ " + line + "\n")
}

func logCLIErr(line string) {
	os.Stderr.WriteString("ø " + line + "\n")
}
