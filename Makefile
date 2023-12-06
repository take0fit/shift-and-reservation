NAME := lispo-customer-api
VERSION := v0.0.1

generate:
	java -cp ./tools/openapi-generator-cli.jar:./tools/go-custom-server-openapi-generator-1.0.0.jar org.openapitools.codegen.OpenAPIGenerator generate -g go-custom-server -i ./api/openapi.yaml -o ./ --git-user-id=LifeSports --git-repo-id=lispo-customer-api --additional-properties=sourceFolder=. --additional-properties=enumClassPrefix=true

combine-openapi:
	java -cp ./tools/openapi-generator-cli.jar:./tools/go-custom-server-openapi-generator-1.0.0.jar org.openapitools.codegen.OpenAPIGenerator generate -g openapi -i ./api/openapi.yaml -o ./api/output

wire:
	go install github.com/google/wire/cmd/wire@latest
	cd app/adapter/http && wire

moq:
	go install github.com/matryer/moq@latest
	go generate ./app/application/...
	go generate ./app/domain/...

test:
	. ./.env.dev && TEST_DB_USER=$${TEST_DB_USER} TEST_DB_PASSWORD=$${TEST_DB_PASSWORD} TEST_DB_HOST=$${TEST_DB_HOST} TEST_DB_PORT=$${TEST_DB_PORT} go test -v ./...

generate-all:
	make generate
	make wire
	make moq
	make combine-openapi
