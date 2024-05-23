# protoc calculator.proto --go_out=../server --go-grpc_out=../server
# protoc calculator.proto --go_out=../client --go-grpc_out=../client

# use *.proto if you have more than one .proto
protoc *.proto --go_out=../server --go-grpc_out=../server
protoc *.proto --go_out=../client --go-grpc_out=../client