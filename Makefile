gen:
	@go-present -in main.slide -out ./docs/index.html


install:
	@https://pkg.go.dev/golang.org/x/tools/present
	@go get github.com/alextanhongpin/go-present


present:
	open http://localhost:3999
	go run golang.org/x/tools/cmd/present


help:
	open https://pkg.go.dev/golang.org/x/tools/present


# To highlight code snippets.
# .code asset/snippets/example_http.go /^func/,/\}/
