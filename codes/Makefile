regenerate:
	go install github.com/gogo/protobuf/protoc-gen-gogo
	protoc \
	--proto_path=../../../../ \
	--gogo_out=../../../../ \
	github.com/easyops-cn/giraffe-micro/codes/code.proto \
