# GRPC gateway using gin

simple grpc gateway using gin

## How to run it
generate proto
```bash
make proto_hello

```
run server
```bash
 go run server/main.go
```
run client 
```bash
 go run client/main.go
```
Try API
```bash
curl  -d '{"name": "Khalid"}' http://localhost:8052/greeter
```

## How to check grpc server
check list grpc service and method
```bash
grpcurl --plaintext  localhost:50051 list
```