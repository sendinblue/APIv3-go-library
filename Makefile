SWAGGER_CODEGEN_VERSION=2.2.3
SWAGGER_CODEGEN_JAR_URL=http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/$(SWAGGER_CODEGEN_VERSION)/swagger-codegen-cli-$(SWAGGER_CODEGEN_VERSION).jar
SWAGGER_DEFINITION_URL=https://api.sendinblue.com/v3/swagger_definition.yml

generate: sibapiv3

sibapiv3: build/generated
	rm -rf $@
	mkdir -p $@
	cp build/generated/*.go $@
	gofmt -w $@

build/generated:\
 build/swagger-codegen-cli.jar\
 build/swagger_definition.yml\
 swagger-codegen-config.json
	rm -rf $@
	java -jar build/swagger-codegen-cli.jar\
	 generate\
	 -i build/swagger_definition.yml\
	 -l go -c swagger-codegen-config.json\
	 -o $@

build/swagger-codegen-cli.jar: Makefile
	mkdir -p build
	curl -o $@ $(SWAGGER_CODEGEN_JAR_URL)

build/swagger_definition.yml: Makefile
	mkdir -p build
	curl -o $@ $(SWAGGER_DEFINITION_URL)

install-dependencies:
	go get -v -u github.com/golang/dep/cmd/dep
	dep ensure -v
	dep prune -v

test:
	go build -v ./sibapiv3 # TODO We should write real tests (use the code)

clean:
	rm -rf build vendor

.PHONY: generate install-dependencies test clean
