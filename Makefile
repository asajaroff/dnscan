GOBIN  = $(shell which go)
SOURCE = $(shell ls cmd/*.go)

all: 
	@echo Using Go binary at ${GOBIN}
	@echo go build ${SOURCE}

run:
	go run ${SOURCE}
