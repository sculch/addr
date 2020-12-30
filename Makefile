LDFLAGS := -ldflags '-s -w -linkmode external'

all: addr

addr: addr.go vendor
	go build -trimpath -buildmode=pie -mod=readonly ${LDFLAGS} -o addr \
		addr.go

vendor: go.mod go.sum
	go mod vendor

clean:
	rm -rf addr vendor

test: vendor
	go test -cover ./...

.PHONY: all clean test
