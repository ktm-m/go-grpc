## Protocol Buffer
Protocol buffer similar to an interface of each end point,
services communicate by using protocol that define in Protocol buffer (.proto)

## Tags
In Protocol buffer field names are not important, important element is tag,
tag start from 1 except 19000 - 19999

Tag number 1 - 15 use 1 byte in space, 16 - 2047 use 2 byte 

Don't change tag when update protocol rules, it makes client and server don't need to use same
protocol version everytime

## Naming Convension
- Message: pascal case
- Field name: snake case (go will build to pascal case)
- Enum: snake case (capital)

## Type of API in gRPC
1. Unary
2. Server Streaming
3. Client Streaming
4. Bi Directional Streaming

## Protoc CLI
- Build proto buffer
protoc file.proto --go_out=path --go_grpc_out=path
Ex. protoc calculator.proto --go_out=../server --go-grpc_out=../server

## Evans
1. Open evans: evans --proto=path_to_file.proto
Ex. evans --proto=./calculator.proto
2. Show package: show package
3. Show service: show service
4. Show description: desc message
Ex. desc HelloRequest
Ex. desc HelloResponse
5. Call service: call service_name
Ex. call Hello
6. Exit evans: exit
7. Open evans with use reflexion: evens -r 
it don't need to set proto

In evans use Ctrl + D to stop sending stream request

## Generate SSL
1. Generate Certificate Authority (CA key)
2. Generate CA Root Public Certificate (CA cert) -> client use
3. Generate server private key
4. Convert server private key to .pem format (gRPC use .pem) -> server use
5. Generate server.csr (certificate service request)
6. Sign server.csr with CA key, result is server.crt -> server use



