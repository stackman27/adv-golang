 build: 
	go build -buildvcs=false 

run: 
	./go-microservice3

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    types/machine.proto    

