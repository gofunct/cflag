package cflag

import (
	_ "github.com/lyft/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/plugin/face"
	_ "github.com/gogo/protobuf/plugin/unmarshal"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/gogo/protobuf/vanity"
	_ "github.com/gogo/protobuf/vanity/command"
	_ "github.com/gogo/protobuf/plugin/testgen"
	_ "github.com/gogo/protobuf/plugin/stringer"
	_ "github.com/gogo/protobuf/plugin/compare"
	_ "github.com/gogo/protobuf/plugin/populate"
	_ "github.com/gogo/protobuf/plugin/marshalto"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/protoc-gen-gogoslick"
	_ "github.com/gogo/protobuf/protoc-gen-gogofaster"
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "github.com/gogo/protobuf/proto"
)