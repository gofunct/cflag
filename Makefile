protoc:
	protoc \
		-I flag \
		-I vendor/ \
		-I vendor/github.com/gogo/googleapis/ \
		-I vendor/github.com/lyft/protoc-gen-validate \
		--gogoslick_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:flag flag/cflag.proto

install:
	go install \
		./vendor/github.com/gogo/protobuf/protoc-gen-gogo \
