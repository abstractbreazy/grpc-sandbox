# env
DOCKER_COMPOSE=sudo docker-compose -f docker/docker-compose.yaml

.PHONY: run buf-build buf-lint buf-generate image-gen docker-run docker-stop

# starts envoy-proxy docker container.
docker-run:
	${DOCKER_COMPOSE} up -d

# stops envoy-proxy docker container.
docker-stop:
	${DOCKER_COMPOSE} down

# starts the service
run: go run ./example/cmd/example/main.go

# compile check
buf-build: 
	buf build ./example/grpc/exampleapis --error-format=json 

# lint .proto files
buf-lint: 
	buf lint ./example/grpc/exampleapis --error-format=json

# generate Go code
buf-gen: 
	bash ./example/grpc/buf.gen.sh

# generate buf image	
image-gen: 
	buf build ./example/grpc/exampleapis -o ./docker/envoy/protos/example.pb --as-file-descriptor-set

 	


