regenerate:
	GOPATH=(go env GOPATH)
	(protoc  \
    --proto_path=. \
    --proto_path=${GOPATH}/src \
    --proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
    --giraffe_out=paths=source_relative,\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:.\
    code.proto status.proto)
