# building a binary Image.
image-create:
	buf build ./example/grpc/exampleapis -o ./images/example.bin --as-file-descriptor-set
