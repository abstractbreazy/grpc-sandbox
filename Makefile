.PHONY: buf-build
buf-build: # compile check
	buf build ./example/grpc/exampleapis --error-format=json 

.PHONY: buf-lint
buf-lint: # lint Protobuf files
	buf lint ./example/grpc/exampleapis --error-format=json

.PHONY: buf-generate
buf-generate: # generate Go code
	bash ./example/grpc/buf.gen.sh


		