.PHONY: clean build

clean:
	rm -f get-subordinates dean-balac-deputy-code-challenge.tar.gz cover.out

build:
	export CGO_ENABLED=0
	go build -o get-subordinates

test:
	go test -v ./...
	go test -coverprofile cover.out
	go tool cover -html=cover.out

package: clean test build
	tar -czvf dean-balac-deputy-code-challenge.tar.gz *