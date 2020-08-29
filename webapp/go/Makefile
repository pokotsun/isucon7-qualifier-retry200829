GOPATH := ${PWD}
export GOPATH

build:
	cd src/isubata && go build -v -o isubata *.go && mv isubata ../../

vet:
	go vet ./src/isubata/...
