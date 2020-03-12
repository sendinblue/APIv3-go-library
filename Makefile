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

test:
	go build -v ./client
	go test -v ./test

clean:
	go clean -i ./...
	rm -rf build

.PHONY: generate install-go-swagger mod-update test clean

mod-update:
	go get -v -u -d all
	$(MOD_TIDY)

MOD_TIDY=rm -f go.sum && go mod tidy -v
.PHONY: mod-tidy
mod-tidy:
	$(MOD_TIDY)

