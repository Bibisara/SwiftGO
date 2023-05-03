## test: runs all tests
test:
	@go test -v ./...

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## coverage: displays test coverage
coverage:
	@go test -cover ./...

## build_cli: builds the command line tool swiftgo and copies it to myapp
build_cli:
	@go build -o ../myapp/swiftgo ./cmd/cli

## build: builds the command line tool dist directory
build:
	@go build -o ./dist/swiftgo ./cmd/cli
# windows users should delete the line above this one, and use the line below instead (uncommented)
#@go build -o dist/swiftgo.exe ./cmd/cli
