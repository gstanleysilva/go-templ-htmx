exec: 
	@go run main.go

run: gen exec

gen:
	@templ generate