
FILES = $(shell find . -type f -name '*.go' -not -path './vendor/*')

gofmt:
	@gofmt -w $(FILES)
	@gofmt -r '&α{} -> new(α)' -w $(FILES)

deps:
	go get -u github.com/mgechev/revive

	go get -u github.com/go-redis/redis
	go get -u github.com/golang/protobuf

test:
	revive -formatter friendly
	go install .
