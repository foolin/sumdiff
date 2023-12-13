install:
	go install ./

tools:
	go install github.com/goreleaser/goreleaser@latest

snapshot:
	goreleaser release --snapshot --clean