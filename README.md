## gRPC-Server Example

Protobuf, gRPC, gRPC-gateway example using `buf.build` tool.

If you have `Go`, `make` and `buf.build` tool installed, you should be able to just execute.

```shell
make
```

I recommend completing [Tour of Buf](https://docs.buf.build/tour/introduction)  to understand the functionality  of both the CLI and the BSR.

You can use [gRPCurl](https://github.com/fullstorydev/grpcurl) to interact with gRPC server.

Example of gRPCurl usage: 

```bash
$ grpcurl -plaintext localhost:9092 describe                     

$ grpcurl -plaintext -d '{}' localhost:9092 example.v1.Example.GetStatus

```
