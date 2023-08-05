## Command for python

This will generate the pb2.py and pb2.pyi. The pyi is good for having types shown when developing

```
protoc -I..\..\proto\ --python_out=. --pyi_out=. ..\..\proto\youtube-dl.proto
```

This will generate the grpc.py

```
python -m grpc_tools.protoc -I..\..\proto\ --python_out=. --pyi_out=. --grpc_python_out=. ..\..\proto\youtube-dl.proto
```
