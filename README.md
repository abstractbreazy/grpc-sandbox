### gRPC-Server Example

Protobuf, gRPC, Envoy    example using `buf.build` tool.

If you have `Go`, `make` and `buf.build` tool installed, you should be able to just execute.

```shell
 make
```

I recommend completing [Tour of Buf](https://docs.buf.build/tour/introduction)  to understand the functionality  of both the CLI and the BSR.


### *Envoy*

Envoy is a very flexible proxy. It's implements *gRPC-JSON transcoding* as a [filter](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_stats_filter).
In order to to start transcoding you need to:

- Generate the proto descriptor of gRPC service   
    ```bash
    make image-gen 
    ```
- Run Envoy using [Docker](https://docs.docker.com/get-docker/) 
    ```bash
    make docker-run 
    ```
- Build a gRPC-service:
    ```bash 
    make run 
    ```

By default, when transcoding occurs, *gRPC-JSON* encodes the message output of a gRPC service method into JSON and sets the HTTP response *Content-Type* header to *`application/json`*.  

You can verify that *gRPC-JSON transcoding* works fine by sending the following request:

``` bash
curl -X 'GET' \
  'https://localhost:8080/v1/example/ping' \
  -H 'accept: application/json'
```

### *gRPCurl*

You can use [gRPCurl](https://github.com/fullstorydev/grpcurl) to interact with gRPC server without using an Envoy proxy.

Example of gRPCurl usage: 

```bash
grpcurl -plaintext localhost:9092 describe                     

grpcurl -plaintext -d '{}' localhost:9092 example.v1.Example.Ping

```

### *Links*:
- [Envoy gRPC-JSON transcoder](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter)
- [gRPCurl Library](https://github.com/fullstorydev/grpcurl)
- [Buf.build documantation](https://docs.buf.build/introduction)
- [GNU Make](https://www.gnu.org/software/make/)
- [Docker](https://docs.docker.com/get-docker/)

