run-user:
	./users/main

run-order:
	./ordens/main

compose:
	docker-compose -f docker-compose.yml up -d 

test-user:
	go test ./users/...

test-orders:
	go test ./ordens/...

proto-user:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=./users --go-grpc_out=./users users/proto/*.proto

proto-order:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=./users --go-grpc_out=./users users/proto/*.proto