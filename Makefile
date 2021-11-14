.PHONY: buf-build
buf-build: # compile check
	buf build ./payments/grpc/paymentsapis --error-format=json 

.PHONY: buf-lint
buf-lint: # lint Protobuf files
	buf lint ./payments/grpc/paymentsapis --error-format=json

.PHONY: buf-generate
buf-generate: # generate Go code
	bash ./payments/grpc/buf.gen.sh


		