# grpc-sample

This is protocol buffer test repository.  
ref: http://www.minaandrawos.com/2014/05/27/practical-guide-protocol-buffers-protobuf-go-golang

## Preparation

```
% brew install protobuf
% go get -u -v github.com/golang/protobuf/proto
% go get -u -v github.com/golang/protobuf/protoc-gen-go
% protoc --go_out=./Player Player.proto
```
