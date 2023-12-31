install:

build:
	mkdir -p ./bin
	go build -o ./bin/sumdiff ./cli/main.go

install: build
	sudo rm -rf ~/go/bin/sumdiff
	sudo rm -rf /usr/local/bin/sumdiff
	sudo cp ./bin/sumdiff /usr/local/bin/

tools:
	go install github.com/goreleaser/goreleaser@latest
	go install github.com/spf13/cobra-cli@latest

release:
	goreleaser release --clean

snapshot:
	goreleaser release --snapshot --clean