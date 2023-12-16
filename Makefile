install:
	go install ./app/

tools:
	#go install github.com/goreleaser/goreleaser@latest
	curl -sfL https://goreleaser.com/static/run | bash


snapshot:
	goreleaser release --snapshot --clean