GOPATH := ${PWD}
export GOPATH

build:
	cd src/isubata && go build -o isubata ... && mv isubata ../../

vet:
	go vet ./src/isubata/...
