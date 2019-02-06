// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/api/servicecontrol/v1/log_entry.proto

package google_api_servicecontrol_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import google_logging_type "go.pedge.io/pb/gogo/google/logging/type"
import google_protobuf1 "github.com/gogo/protobuf/types"
import google_protobuf2 "github.com/gogo/protobuf/types"
import google_protobuf3 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An individual log entry.
type LogEntry struct {
	// Required. The log to which this log entry belongs. Examples: `"syslog"`,
	// `"book_log"`.
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// The time the event described by the log entry occurred. If
	// omitted, defaults to operation start time.
	Timestamp *google_protobuf3.Timestamp `protobuf:"bytes,11,opt,name=timestamp" json:"timestamp,omitempty"`
	// The severity of the log entry. The default value is
	// `LogSeverity.DEFAULT`.
	Severity google_logging_type.LogSeverity `protobuf:"varint,12,opt,name=severity,proto3,enum=google.logging.type.LogSeverity" json:"severity,omitempty"`
	// A unique ID for the log entry used for deduplication. If omitted,
	// the implementation will generate one based on operation_id.
	InsertId string `protobuf:"bytes,4,opt,name=insert_id,json=insertId,proto3" json:"insert_id,omitempty"`
	// A set of user-defined (key, value) data that provides additional
	// information about the log entry.
	Labels map[string]string `protobuf:"bytes,13,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The log entry payload, which can be one of multiple types.
	//
	// Types that are valid to be assigned to Payload:
	//	*LogEntry_ProtoPayload
	//	*LogEntry_TextPayload
	//	*LogEntry_StructPayload
	Payload isLogEntry_Payload `protobuf_oneof:"payload"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptorLogEntry, []int{0} }

type isLogEntry_Payload interface {
	isLogEntry_Payload()
}

type LogEntry_ProtoPayload struct {
	ProtoPayload *google_protobuf1.Any `protobuf:"bytes,2,opt,name=proto_payload,json=protoPayload,oneof"`
}
type LogEntry_TextPayload struct {
	TextPayload string `protobuf:"bytes,3,opt,name=text_payload,json=textPayload,proto3,oneof"`
}
type LogEntry_StructPayload struct {
	StructPayload *google_protobuf2.Struct `protobuf:"bytes,6,opt,name=struct_payload,json=structPayload,oneof"`
}

func (*LogEntry_ProtoPayload) isLogEntry_Payload()  {}
func (*LogEntry_TextPayload) isLogEntry_Payload()   {}
func (*LogEntry_StructPayload) isLogEntry_Payload() {}

func (m *LogEntry) GetPayload() isLogEntry_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *LogEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogEntry) GetTimestamp() *google_protobuf3.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogEntry) GetSeverity() google_logging_type.LogSeverity {
	if m != nil {
		return m.Severity
	}
	return google_logging_type.LogSeverity_DEFAULT
}

func (m *LogEntry) GetInsertId() string {
	if m != nil {
		return m.InsertId
	}
	return ""
}

func (m *LogEntry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *LogEntry) GetProtoPayload() *google_protobuf1.Any {
	if x, ok := m.GetPayload().(*LogEntry_ProtoPayload); ok {
		return x.ProtoPayload
	}
	return nil
}

func (m *LogEntry) GetTextPayload() string {
	if x, ok := m.GetPayload().(*LogEntry_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

func (m *LogEntry) GetStructPayload() *google_protobuf2.Struct {
	if x, ok := m.GetPayload().(*LogEntry_StructPayload); ok {
		return x.StructPayload
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LogEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LogEntry_OneofMarshaler, _LogEntry_OneofUnmarshaler, _LogEntry_OneofSizer, []interface{}{
		(*LogEntry_ProtoPayload)(nil),
		(*LogEntry_TextPayload)(nil),
		(*LogEntry_StructPayload)(nil),
	}
}

func _LogEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LogEntry)
	// payload
	switch x := m.Payload.(type) {
	case *LogEntry_ProtoPayload:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProtoPayload); err != nil {
			return err
		}
	case *LogEntry_TextPayload:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.TextPayload)
	case *LogEntry_StructPayload:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StructPayload); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LogEntry.Payload has unexpected type %T", x)
	}
	return nil
}

func _LogEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LogEntry)
	switch tag {
	case 2: // payload.proto_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Any)
		err := b.DecodeMessage(msg)
		m.Payload = &LogEntry_ProtoPayload{msg}
		return true, err
	case 3: // payload.text_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Payload = &LogEntry_TextPayload{x}
		return true, err
	case 6: // payload.struct_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf2.Struct)
		err := b.DecodeMessage(msg)
		m.Payload = &LogEntry_StructPayload{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LogEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LogEntry)
	// payload
	switch x := m.Payload.(type) {
	case *LogEntry_ProtoPayload:
		s := proto.Size(x.ProtoPayload)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LogEntry_TextPayload:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TextPayload)))
		n += len(x.TextPayload)
	case *LogEntry_StructPayload:
		s := proto.Size(x.StructPayload)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "google.api.servicecontrol.v1.LogEntry")
}

func init() {
	proto.RegisterFile("google/api/servicecontrol/v1/log_entry.proto", fileDescriptorLogEntry)
}

var fileDescriptorLogEntry = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4d, 0x8b, 0xdb, 0x30,
	0x10, 0x86, 0xe3, 0xcd, 0x36, 0x8d, 0xc7, 0xc9, 0x52, 0xc4, 0x42, 0x5d, 0x37, 0x50, 0xd3, 0x42,
	0xc9, 0xa1, 0xc8, 0x6c, 0xf6, 0xb2, 0xfd, 0x38, 0xb4, 0x81, 0x42, 0x5a, 0x72, 0x58, 0xb4, 0xbd,
	0x07, 0x25, 0x51, 0x85, 0xa8, 0x22, 0x19, 0x5b, 0x31, 0xf5, 0x4f, 0xee, 0xbf, 0x28, 0xd6, 0x87,
	0xbb, 0xdd, 0x85, 0xdc, 0x66, 0x34, 0xcf, 0x3b, 0x33, 0xef, 0x08, 0xde, 0x71, 0xad, 0xb9, 0x64,
	0x05, 0x2d, 0x45, 0x51, 0xb3, 0xaa, 0x11, 0x3b, 0xb6, 0xd3, 0xca, 0x54, 0x5a, 0x16, 0xcd, 0x55,
	0x21, 0x35, 0xdf, 0x30, 0x65, 0xaa, 0x16, 0x97, 0x95, 0x36, 0x1a, 0xcd, 0x1c, 0x8d, 0x69, 0x29,
	0xf0, 0xff, 0x34, 0x6e, 0xae, 0xb2, 0xd9, 0xbd, 0x5e, 0x54, 0x29, 0x6d, 0xa8, 0x11, 0x5a, 0xd5,
	0x4e, 0x9b, 0xbd, 0xf5, 0x55, 0xa9, 0x39, 0x17, 0x8a, 0x17, 0xa6, 0x2d, 0x6d, 0xb2, 0xa9, 0x59,
	0xc3, 0x2a, 0x61, 0xfc, 0x8c, 0xec, 0x85, 0xe7, 0x6c, 0xb6, 0x3d, 0xfe, 0x2c, 0xa8, 0x0a, 0xa5,
	0xd9, 0xc3, 0x52, 0x6d, 0xaa, 0xe3, 0xce, 0xf8, 0xea, 0xab, 0x87, 0x55, 0x23, 0x0e, 0xac, 0x36,
	0xf4, 0x50, 0x3a, 0xe0, 0xf5, 0x9f, 0x21, 0x8c, 0xd7, 0x9a, 0x7f, 0xed, 0x0c, 0x21, 0x04, 0xe7,
	0x8a, 0x1e, 0x58, 0x0a, 0x79, 0x34, 0x8f, 0x89, 0x8d, 0xd1, 0x0d, 0xc4, 0xbd, 0x26, 0x4d, 0xf2,
	0x68, 0x9e, 0x2c, 0x32, 0xec, 0x2d, 0x87, 0xae, 0xf8, 0x47, 0x20, 0xc8, 0x3f, 0x18, 0x7d, 0x82,
	0x71, 0xb0, 0x91, 0x4e, 0xf2, 0x68, 0x7e, 0xb1, 0xc8, 0x83, 0xd0, 0xfb, 0xc5, 0x9d, 0x5f, 0xbc,
	0xd6, 0xfc, 0xce, 0x73, 0xa4, 0x57, 0xa0, 0x97, 0x10, 0x0b, 0x55, 0xb3, 0xca, 0x6c, 0xc4, 0x3e,
	0x3d, 0xb7, 0x0b, 0x8d, 0xdd, 0xc3, 0xb7, 0x3d, 0xfa, 0x0e, 0x23, 0x49, 0xb7, 0x4c, 0xd6, 0xe9,
	0x34, 0x1f, 0xce, 0x93, 0xc5, 0x02, 0x9f, 0xfa, 0x04, 0x1c, 0x0c, 0xe2, 0xb5, 0x15, 0xd9, 0x98,
	0xf8, 0x0e, 0xe8, 0x23, 0x4c, 0xad, 0x8f, 0x4d, 0x49, 0x5b, 0xa9, 0xe9, 0x3e, 0x3d, 0xb3, 0x26,
	0x2f, 0x1f, 0x99, 0xfc, 0xa2, 0xda, 0xd5, 0x80, 0x4c, 0x6c, 0x7e, 0xeb, 0x58, 0xf4, 0x06, 0x26,
	0x86, 0xfd, 0x36, 0xbd, 0x76, 0xd8, 0x2d, 0xba, 0x1a, 0x90, 0xa4, 0x7b, 0x0d, 0xd0, 0x67, 0xb8,
	0x70, 0x9f, 0xd2, 0x63, 0x23, 0x3b, 0xe2, 0xf9, 0xa3, 0x11, 0x77, 0x16, 0x5b, 0x0d, 0xc8, 0xd4,
	0x09, 0x7c, 0x87, 0xec, 0x3d, 0x24, 0xf7, 0x56, 0x47, 0xcf, 0x60, 0xf8, 0x8b, 0xb5, 0x69, 0x64,
	0xaf, 0xd2, 0x85, 0xe8, 0x12, 0x9e, 0x34, 0x54, 0x1e, 0x99, 0x5d, 0x3e, 0x26, 0x2e, 0xf9, 0x70,
	0x76, 0x13, 0x2d, 0x63, 0x78, 0xea, 0xa7, 0x2e, 0xaf, 0x21, 0xdf, 0xe9, 0xc3, 0xc9, 0x53, 0x2d,
	0xa7, 0xe1, 0x56, 0xb7, 0xd6, 0x66, 0xb4, 0x1d, 0xd9, 0xe5, 0xae, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x5b, 0x7a, 0x37, 0x09, 0x15, 0x03, 0x00, 0x00,
}