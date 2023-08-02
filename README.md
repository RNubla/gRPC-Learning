# Learning gRPC with golang and python

This repo is for my own learning. It is to learn about gRPC between python and golang. This will help with iterop of the two language to comunicate with each other.

Command for windows

```
https://stackoverflow.com/questions/64048132/proto-path-passed-empty-directory-name
protoc --proto_path=. --go_out=. ./addressBook.proto
```

To generate RPC

```
protoc --proto_path=. --go_out=. --go-grpc_out=. ./helloWorld.proto
```

### Reference

https://www.ardanlabs.com/blog/2020/06/python-go-grpc.html
https://grpc.io/docs/languages/go/basics/
