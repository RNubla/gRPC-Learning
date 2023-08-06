# Learning gRPC with golang and python

This repo is for my own learning. It is to learn about gRPC between python and golang. This will help with iterop of the two language to comunicate with each other.

## Command listed are only tested on windows.

```
https://stackoverflow.com/questions/64048132/proto-path-passed-empty-directory-name
protoc --proto_path=. --go_out=. ./addressBook.proto
```

### To generate RPC for GoLang

```
protoc --proto_path=. --go_out=. --go-grpc_out=. ./helloWorld.proto
```

### To generate RPC for Python

```
 python -m grpc_tools.protoc --proto_path=. --python_out=./services/hello-world --grpc_python_out=./services/hello-world ./helloWorld.proto
```

Thanks to Miki Tebeka for writing this [article](https://www.ardanlabs.com/blog/2020/06/python-go-grpc.html)

## Key takeaways

1. If you want go to call python functions you must generate rpc for both python and go.
2. To send data from go to python, we use the `HelloWorldRequest()` from Go RPC which will be caught by the python server on the `request` parameter.
3. In order for python to send back the data, `HelloWorldResponse` is used to return the response.
4. We also need to override the service `Print` so that we can manipulate the data that was sent from the client/go.
5. `add_HelloWorldServicer_to_server(HelloWorldServer(), server)` This code will add the service to the server

# Load Balancer References

https://www.useanvil.com/blog/engineering/grpc-load-balancing/
https://medium.com/bitaksi-tech/round-robin-load-balancing-with-nginx-via-docker-eng-b183a11bcc8b
https://linuxhint.com/scale-docker-containers-using-nginx-reverse-proxy-and-load-balancer/

### Reference

https://www.ardanlabs.com/blog/2020/06/python-go-grpc.html

https://grpc.io/docs/languages/go/basics/
