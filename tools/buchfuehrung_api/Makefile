build:
	rm -f buchfuerhung_api
	GO15VENDOREXPERIMENT=1 go build -ldflags "-X main.BuildTime=`date +%s` -X main.BuildVersion=`git describe --always`"

build_freebsd:
	rm -f buchfuerhung_api_freebsd
	GO15VENDOREXPERIMENT=1 env GOOS=freebsd go build -o buchfuerhung_api_freebsd -ldflags "-X main.BuildTime=`date +%s` -X main.BuildVersion=`git describe --always`"

dependencies_init:
	GO15VENDOREXPERIMENT=1 glide init

dependencies_update:
	GO15VENDOREXPERIMENT=1 glide up
