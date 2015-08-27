package = github.com/jcbwlkr/learning-api

.PHONY: release

release:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/learning-api-linux-amd64 $(package)
	GOOS=linux GOARCH=386 go build -o release/learning-api-linux-386 $(package)
	GOOS=linux GOARCH=arm go build -o release/learning-api-linux-arm $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/learning-api-darwin-amd64 $(package)
	GOOS=darwin GOARCH=386 go build -o release/learning-api-darwin-386 $(package)
	GOOS=windows GOARCH=amd64 go build -o release/learning-api-windows-amd64.exe $(package)
	GOOS=windows GOARCH=386 go build -o release/learning-api-windows-386.exe $(package)
