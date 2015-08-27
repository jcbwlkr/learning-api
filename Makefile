package = github.com/jcbwlkr/learning-api

.PHONY: release

release:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/myproject-linux-amd64 $(package)
	GOOS=linux GOARCH=386 go build -o release/myproject-linux-386 $(package)
	GOOS=linux GOARCH=arm go build -o release/myproject-linux-arm $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/myproject-darwin-amd64 $(package)
	GOOS=darwin GOARCH=386 go build -o release/myproject-darwin-386 $(package)
	GOOS=windows GOARCH=amd64 go build -o release/myproject-windows-amd64 $(package)
	GOOS=windows GOARCH=386 go build -o release/myproject-windows-386 $(package)
