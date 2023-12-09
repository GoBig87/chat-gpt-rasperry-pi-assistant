.PHONY: build

build:
	CGO_LDFLAGS="-latomic -lpthread -lm -ldl" go build -o ${GOBIN}/gptctl *.go
