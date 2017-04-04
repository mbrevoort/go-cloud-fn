// Code generated by go-bindata.
// sources:
// index.js
// package.json
// DO NOT EDIT!

package template

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xc1\x6e\xdb\x38\x10\xbd\xfb\x2b\xa6\x27\x49\xa8\x23\xf5\xb2\xc0\x22\x80\x77\x91\x36\xee\x26\x8b\xd6\x29\x62\x17\x5d\xa0\x28\x0a\x46\x1c\x5b\x4c\x28\x8e\x4a\x8e\xec\x18\x81\xff\x7d\x41\x52\x92\xe5\xb4\x39\x2c\xb0\x37\x9b\x33\x7a\xf3\x38\xf3\xf8\x48\x8b\x3f\x5a\x65\x31\x4d\x24\x31\x9a\x6d\x92\xe5\x25\x99\xb5\xda\xa4\xd9\x64\x52\x14\x70\x25\x8c\xd4\x08\x6f\x45\xf9\xb0\xb1\xd4\x1a\x09\xb8\x45\xc3\x0e\x44\x59\x92\x95\xca\x6c\x80\x09\x5c\x83\xe5\x64\xdd\x9a\x92\x15\x19\x70\x95\xaa\xe3\x77\x16\x52\x29\x58\x64\xf0\x34\x01\xb0\xc8\xad\x35\x60\x70\x07\x9f\x2c\xd5\xca\x61\x9a\x5a\x74\xa4\xb7\x38\x05\x8b\xf7\x58\x72\x06\xb3\x3f\x42\x2e\x40\x49\xc6\x91\xc6\x5c\xd3\x26\x4d\xe6\x8b\xd5\xfc\x16\x96\x57\xd7\x1f\xe1\xea\x62\x71\xf9\x61\x7e\x9b\x64\x21\xab\x28\x60\xd9\x88\x9d\x01\xae\x10\x06\x02\xc2\x48\x50\xc6\x03\x86\x75\x34\x5b\x58\x5b\xaa\xc3\x9f\x46\x58\x34\x0c\x8d\xa5\x12\x9d\xcb\x87\x5a\x0c\x0d\xcc\x60\xe8\x46\x59\x29\x2d\xbf\x77\x59\x49\x96\xe3\x23\x96\xef\x95\xc6\x34\xc9\x8b\xa7\xa7\xfc\x7d\x57\x6a\x21\x6a\x3c\x1c\x92\x29\x7c\xfd\x36\xed\x88\x83\xaf\x77\x3e\x14\x40\xb3\x0d\xcb\x87\x48\x78\x2b\x2c\x68\xe1\xf8\x23\x3a\x27\x36\x18\xd6\x9a\xdc\xb1\x54\x26\x77\xc8\x73\x53\x92\x6f\x6a\x9a\xb4\xbc\x3e\xfb\xfd\xb8\xcb\x0f\xb4\x01\xc7\xc2\x48\x61\x25\xa0\xb5\x50\x47\x00\x17\xda\x3f\x0a\x1c\x11\xd1\xda\x9c\x4c\x9a\xf8\x09\x24\x53\x48\xd1\xda\x51\x7f\x8f\x1d\x46\x6b\xc9\xfa\x68\xce\xb4\x64\xeb\xab\x67\xd9\x98\x73\x40\xa3\x96\x4f\xd0\xa8\xe5\x5f\xa2\xf9\x79\xf9\xdc\xe7\x58\x30\xde\x36\xcc\x80\x5a\x3e\xad\xe1\xc1\x4b\x4d\x0e\x3d\x7a\x49\x12\x5f\x84\x4f\xe6\xff\xcc\xdf\xc1\xbb\x0f\x37\xcb\xf9\x65\xd2\xa3\xab\x75\xfc\x0a\x5e\xcd\x66\xf0\x26\x1b\x3e\x0c\xdd\x5b\x55\xca\x41\x8d\xc2\xb8\xa0\x01\x2f\x50\x58\x0b\xa5\x51\x42\x01\x8d\x30\xaa\x7c\x40\x99\xc3\x92\x60\x87\x9d\x14\xa1\x12\x56\xe6\x03\x48\x5c\x4c\xfb\x6a\x07\x40\xed\xf0\xb4\xc8\x6d\xd4\x72\x54\x59\x14\x38\xec\x14\x57\x61\x41\x0b\x46\xc7\x7e\xd7\x4d\xcb\x51\x8d\xb1\xa9\x63\x84\x6b\x03\xa5\x70\x08\xb4\x0e\x14\x6b\x7f\xba\x2a\xe6\x66\x0a\xec\x37\xa0\x22\x7b\x8b\xae\x21\xe3\xd3\xee\x3c\xa7\x31\xc7\x40\x20\x1d\x35\x7a\xe0\x3b\xee\x75\x51\xc0\x17\xab\x38\x32\x8d\x20\x45\x27\xa7\xc2\x1f\x00\x4f\x94\x69\xe8\x54\xe2\x20\xe8\x33\x1c\x2b\xa7\x36\x46\xe8\x1e\x66\x6e\xa4\x27\xab\x4c\xd3\x76\x3c\x7a\x2d\xef\x7c\x81\xf4\xef\xe5\xcd\x22\x77\x41\x07\x6a\xbd\x8f\x5e\x90\x9d\xe4\xa1\x91\x5d\x53\x9f\x1d\xf8\x4b\x58\xdc\x5c\xce\xe1\x7a\xf1\xe9\xf3\x2a\x0c\xf9\x90\x4d\x0e\x63\x4b\xf2\x8d\x81\x8e\xee\xd1\x7b\xaa\x10\xbc\xf2\xb1\xd4\xe2\x0f\x6f\x2c\x2e\x8a\xc1\x9f\xbc\x2e\xfd\x2d\xc9\xfd\x64\x02\xe0\x76\x8a\xcb\x2a\x24\xe6\x1b\xe4\x34\x29\xc9\x30\x1a\x3e\xe3\x7d\x83\x49\xd6\x8b\x28\xcc\x24\x11\x4d\xa3\x55\x29\x7c\x95\xe2\xde\x91\x49\xce\x27\x7d\xdb\x07\x50\x98\xc1\xb3\x2d\x7b\xe8\x3b\x92\xfb\x7e\x10\x77\x16\xc5\xc3\x0b\xa8\x8f\x67\xbb\xdd\xee\x6c\x4d\xb6\x3e\x6b\xad\x46\x6f\x05\x28\x87\x32\x41\xc6\x08\x1e\xcc\x9b\x98\x43\x0b\x6b\xb2\x50\x6a\x6a\xe5\x60\x7d\x0e\x24\x79\x4f\xa8\x94\x9b\x82\x23\xb8\x6f\xbd\xb1\x69\xb1\x07\xa1\xc9\x6c\x8e\x50\x41\x98\x8a\x7d\x92\xb5\x7b\xa8\x85\x79\x05\x1f\xc5\xfe\x0e\xfd\x09\x70\x15\xb5\x5a\x46\x5b\xb4\x6d\xc9\xe0\xa8\x46\x78\x50\x71\xda\x8d\xa5\x06\xed\x11\xca\x13\xee\x9b\x10\xe8\xfd\x09\x64\xc1\x10\xe7\xa0\x91\x13\x07\x0f\x88\x0d\x28\x8e\x22\xde\x89\x7d\xe0\x6d\x68\x37\x05\xe1\x26\xa7\x67\x54\x39\xa8\x68\xf7\xd3\xa6\xee\xb0\x12\x5b\x74\xf9\xb1\xe3\xde\x2c\x9f\xcd\x6b\xfa\x8b\x19\x65\xff\xf7\x8c\xa8\x64\xe4\x33\xc7\x16\x45\xfd\x82\x02\x7a\xb8\x17\xd0\x18\x1f\xb9\x68\xb4\x50\x2f\x29\xe8\x85\xef\x0f\x93\x4e\xc4\xeb\x56\xeb\xcf\x56\x77\x99\x8d\x25\xa6\x92\x34\xbc\x86\xe4\xbc\x28\x12\x78\x0d\x83\x9c\x2b\x72\x9c\x64\xdd\x0a\x59\xb5\x51\x46\xf8\x4f\x7b\x24\x7f\x84\x6e\xbb\xc1\xcd\x3a\xb1\x27\xbe\x74\x72\x3e\xe6\x34\x8d\x81\x0a\x85\x44\xeb\x62\x2c\xef\xfe\x75\xb1\x1a\xb9\x22\xd9\x85\xe2\x9f\x2e\x62\xb1\x26\xc6\xef\x42\x4a\xdb\x85\x55\xd3\x85\x5a\xab\x93\xf3\x7e\x3b\xfd\x0e\x47\x8f\x87\x74\xc4\xcf\xcf\x26\xe7\x0a\x4d\x78\x33\xb4\x7a\x7c\xf9\xf8\xad\x78\x6b\xe9\x27\x1b\x4e\x47\x9f\x36\xe9\x9c\x31\x77\x2c\xb8\x75\xc1\x83\xba\xdf\xdf\xc3\x2d\x73\x4c\x40\x8e\xd1\x6e\x6b\xe3\x88\x91\x31\x74\x94\xc9\x89\x5d\x5d\xde\x2c\xe6\xd1\xaf\xbe\xdc\x5e\xaf\xae\x17\x7f\xc1\xe5\xc5\xea\xa2\xb7\x2d\x80\xbc\x14\x5c\x56\x69\x3a\x22\x3d\xa2\xf4\xdb\x9b\x37\xd9\x60\x84\x83\xcd\x3d\x3d\xf9\x3b\x2d\x5f\x59\xb5\xd9\xa0\xbd\x5a\xad\x3e\xc1\xe1\x30\xc1\xc7\x86\x2c\xbb\xaf\xc9\xcf\x8f\x90\x6f\x30\x3b\x3e\x81\x9e\x99\xdf\x09\xdb\x40\xf4\xfd\xe7\xc5\x3b\x58\xae\x2e\x6e\xa3\xbb\x76\x8f\xb3\xa3\x7d\x1e\x01\x26\x87\xc8\x26\xdc\x78\xff\x85\x42\x78\x29\x9e\xbc\xfd\xc6\xc3\x0d\xd1\x3c\x5c\x09\x43\x05\x23\x7d\x81\x7f\x03\x00\x00\xff\xff\xa9\xee\xb6\xab\x93\x0a\x00\x00")

func indexJsBytes() ([]byte, error) {
	return bindataRead(
		_indexJs,
		"index.js",
	)
}

func indexJs() (*asset, error) {
	bytes, err := indexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.js", size: 2707, mode: os.FileMode(420), modTime: time.Unix(1491259207, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _packageJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xe6\x52\x50\x50\xca\x4b\xcc\x4d\x55\xb2\x52\x50\xaa\xae\xd6\x73\x2b\xcd\x4b\x2e\xc9\xcc\xcf\xf3\x4b\xcc\x4d\xad\xad\x55\xd2\x01\x49\x97\xa5\x16\x15\x67\xe6\xe7\x81\x54\x18\xe8\x19\xe8\x19\x40\x44\x53\x52\x0b\x52\xf3\x52\x52\xf3\x92\x33\x53\x8b\x95\xac\x14\x40\x06\x81\x44\xf3\x4b\x52\xf3\xca\x40\x4a\xe3\x4c\xc0\x6a\xb9\x14\x14\x6a\xb9\x6a\x01\x01\x00\x00\xff\xff\xe6\x48\x09\x7c\x67\x00\x00\x00")

func packageJsonBytes() ([]byte, error) {
	return bindataRead(
		_packageJson,
		"package.json",
	)
}

func packageJson() (*asset, error) {
	bytes, err := packageJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "package.json", size: 103, mode: os.FileMode(420), modTime: time.Unix(1491255478, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"index.js": indexJs,
	"package.json": packageJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"index.js": &bintree{indexJs, map[string]*bintree{}},
	"package.json": &bintree{packageJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

