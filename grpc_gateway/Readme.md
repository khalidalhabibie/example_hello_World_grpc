# GRPC gateway using grpc-ecosystem

simple gateway using grpc-ecosystem

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
curl  -d '{"name": "Khalid"}' http://localhost:8081/api/v1/greeter
```
documentation : swagger/hello.swagger
