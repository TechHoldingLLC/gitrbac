build:
	go build -o bin/gitrbac .

compile:
	GOOS=darwin GOARCH=amd64 go build -o bin/gitrbac-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o bin/gitrbac-darwin-arm64 .
	GOOS=linux GOARCH=amd64 go build -o bin/gitrbac-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -o bin/gitrbac-linux-arm64 .

	tar czf bin/gitrbac-darwin-amd64.tar.gz bin/gitrbac-darwin-amd64
	tar czf bin/gitrbac-darwin-arm64.tar.gz bin/gitrbac-darwin-arm64
	tar czf bin/gitrbac-linux-amd64.tar.gz bin/gitrbac-linux-amd64
	tar czf bin/gitrbac-linux-arm64.tar.gz bin/gitrbac-linux-arm64

	sha256sum bin/gitrbac-darwin-amd64.tar.gz > bin/gitrbac-darwin-amd64.tar.gz.sha256
	sha256sum bin/gitrbac-darwin-arm64.tar.gz > bin/gitrbac-darwin-arm64.tar.gz.sha256
	sha256sum bin/gitrbac-linux-amd64.tar.gz > bin/gitrbac-linux-amd64.tar.gz.sha256
	sha256sum bin/gitrbac-linux-arm64.tar.gz > bin/gitrbac-linux-arm64.tar.gz.sha256
