install:
	go install ./app/

tools:
	go install github.com/goreleaser/goreleaser@latest
	go install github.com/spf13/cobra-cli@latest


snapshot:
	goreleaser release --snapshot --clean