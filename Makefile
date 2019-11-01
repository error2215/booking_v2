# Prepare variables
SOURCE_PATH :=TEST_SOURCE_PATH=$(PWD)
VERSION=`git rev-parse --short HEAD`

# Download all dependencies
install: mod
	@echo "All packages successfully installed!"

mod:
	@echo "======================================================================"
	@echo "Run MOD"
	@ GO111MODULE=on go mod verify
	@ GO111MODULE=on go mod tidy
	@ GO111MODULE=on go mod vendor
	@ GO111MODULE=on go mod download
	@ GO111MODULE=on go mod verify
	@echo "======================================================================"
# Run tests
tests:
	$(SOURCE_PATH) go test  -coverprofile=coverage.out -v ./backend/api/
	go tool cover -html=coverage.out -o coverage.html
	rm coverage.out

test: mod tests

# Build the server
build:
	@echo "Building the binary..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./application ./main.go

# Run the server
run: build
	./application


