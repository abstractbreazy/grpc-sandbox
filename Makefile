# env
DOCKER_COMPOSE=docker-compose -f docker/docker-compose.yaml

.PHONY: run buf-build buf-lint buf-generate image-gen docker-run docker-stop
run: go run ./example/cmd/example/main.go

buf-build: # compile check
	buf build ./example/grpc/exampleapis --error-format=json 

buf-lint: # lint Protobuf files
	buf lint ./example/grpc/exampleapis --error-format=json

buf-gen: # generate Go code
	bash ./example/grpc/buf.gen.sh
	
image-gen: # generate buf image
	buf build ./example/grpc/exampleapis -o ./docker/protos/example.pb --as-file-descriptor-set

docker-run:
	${DOCKER_COMPOSE} up -d

docker-stop:
	${DOCKER_COMPOSE} down
 	


