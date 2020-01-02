# Installation
## Mac
### gRPC
```
brew tap grpc/grpc
brew install grpc
```
### protobuf
```
brew install protobuf
```
### protobuf plugin (Go)
```
go get github.com/golang/protobuf/protoc-gen-go
```
### grpcurl (debug tool)
```
brew install grpcurl
```
# Build
```
cd ../
protoc --go_out=plugins=grpc:./server/ ./protocol/sample.proto
```
# Debug
```
go run main.go
grpcurl -plaintext -d '{ "id": 3, "requestMessage":"REQUEST_MESSAGE" }' localhost:9999 sample.SampleService/RequestSampleMethod
```