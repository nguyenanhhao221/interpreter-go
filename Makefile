build-docker: 
	CGO_ENABLED=0 GOOS=linux go build -o bin/monkey
run:
	./bin/monkey

build: 
	@go build -o bin/monkey

dev: build
	@./bin/monkey

test: 
	@go test -v ./...
