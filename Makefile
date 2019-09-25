.PHONY: build clean deploy

build:
	GOOS=linux GOARCH=amd64 go build -x -ldflags="-s -w" -o bin/rest rest/main.go
	chmod +x bin/rest

clean:
	rm -rf ./bin

deploy: clean build
	npx sls deploy --verbose
