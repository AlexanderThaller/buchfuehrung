generate:
	go generate

dependencies_init:
	GO15VENDOREXPERIMENT=1 glide init

dependencies_update:
	GO15VENDOREXPERIMENT=1 glide up
