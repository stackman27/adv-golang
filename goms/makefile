build: 
	@go build -o -buildvcs=fals ./bin/fact

run: build
	./bin/fact

test:
	go test -v ./...

# check this out if help need installing proto https://grpc.io/docs/languages/go/quickstart/
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    types/types.proto     
