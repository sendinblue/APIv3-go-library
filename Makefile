SWAGGER_DEFINITION_URL=https://api.sendinblue.com/v3/swagger_definition.yml
GO_SWAGGER_PACKAGE=github.com/go-swagger/go-swagger/cmd/swagger

generate: \
	build/swagger_definition.yml \
	install-go-swagger
	rm -rf client models
	swagger generate client -f build/swagger_definition.yml --additional-initialism=SMS

build/swagger_definition.yml: Makefile
	mkdir -p build
	curl -o $@ $(SWAGGER_DEFINITION_URL)

install-go-swagger:
	go get -v -u $(GO_SWAGGER_PACKAGE)

clean:
	go clean -i ./... ./vendor/...
	rm -rf build vendor

.PHONY: generate install-go-swagger install-dependencies test clean