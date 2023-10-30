# debug
flags := -gcflags="all=-N -l -c 2"
BUILD_TARGET := altools
cli := ./bin/${BUILD_TARGET}
env := GO111MODULE=on
GOBIN := /Users/mac/go/bin

build:
	$(env) go build -o $(cli) cmd/main.go

install:build
	cp $(cli) $(GOBIN)

clean:
	@rm -f $(cli)
	@rm -rf ./bin
