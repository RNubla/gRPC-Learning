Command for windows

```
https://stackoverflow.com/questions/64048132/proto-path-passed-empty-directory-name
protoc --proto_path=. --go_out=. ./addressBook.proto
```

To generate RPC

```
protoc --proto_path=. --go_out=. --go-grpc_out=. ./helloWorld.proto
```
