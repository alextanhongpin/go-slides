gen:
	@mkdir -p docs
	@# Required for github-pages deployment.
	@# Otherwise the assets won't appear.
	@cp -r asset/ docs/asset/
	@go-present -in main.slide -out ./docs/index.html


install:
	@go install github.com/alextanhongpin/go-present@latest


present:
	@open http://localhost:3999
	@go run golang.org/x/tools/cmd/present


help:
	@open https://pkg.go.dev/golang.org/x/tools/present


# To highlight code snippets.
# .code asset/snippets/example_http.go /^func/,/\}/
