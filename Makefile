run:
	@go run main.go

gen:
	@templ generate

exec: gen run