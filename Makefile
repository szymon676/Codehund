generate: 
	@templ generate
	
run:
	@templ generate
	@go run .
build:
	@templ generate
	@go build -o .

