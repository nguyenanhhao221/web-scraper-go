build: 
	@go build -o bin/web-scraper-go

run: build
	@./bin/web-scraper-go

test: 
	@go test -v -cover ./...
