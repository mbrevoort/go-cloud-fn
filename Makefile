template:
	cd template && go-bindata -pkg template ./index.js ./package.json
.PHONY: template