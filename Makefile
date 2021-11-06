.PHONY: buf-build
buf-build: # compile check
	buf build ./example/grpc/exampleapis --error-format=json 

.PHONY: buf-lint
buf-lint: # lint Protobuf files
	buf lint ./example/grpc/exampleapis --error-format=json

.PHONY: buf-generate
buf-generate: # generate Go code
	bash ./example/grpc/buf.gen.sh

.PHONY: image-create
image-create: # build a binary Image.
	buf build ./example/grpc/exampleapis -o ./envoy/descriptors/example.pb --as-file-descriptor-set

.PHONY: docker-start
docker-start:
	docker-compose -f docker-compose.yml up -d

.PHONY: docker-stop
docker-stop:
	docker-compose -f docker-compose.yml down	