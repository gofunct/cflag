// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/assistant/embedded/v1alpha1/embedded_assistant.proto

package google_assistant_embedded_v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import google_rpc "go.pedge.io/pb/gogo/google/rpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Audio encoding of the data sent in the audio message.
// Audio must be one-channel (mono). The only language supported is "en-US".
type AudioInConfig_Encoding int32

const (
	// Not specified. Will return result [google.rpc.Code.INVALID_ARGUMENT][].
	AudioInConfig_ENCODING_UNSPECIFIED AudioInConfig_Encoding = 0
	// Uncompressed 16-bit signed little-endian samples (Linear PCM).
	// This encoding includes no header, only the raw audio bytes.
	AudioInConfig_LINEAR16 AudioInConfig_Encoding = 1
	// [`FLAC`](https://xiph.org/flac/documentation.html) (Free Lossless Audio
	// Codec) is the recommended encoding because it is
	// lossless--therefore recognition is not compromised--and
	// requires only about half the bandwidth of `LINEAR16`. This encoding
	// includes the `FLAC` stream header followed by audio data. It supports
	// 16-bit and 24-bit samples, however, not all fields in `STREAMINFO` are
	// supported.
	AudioInConfig_FLAC AudioInConfig_Encoding = 2
)

var AudioInConfig_Encoding_name = map[int32]string{
	0: "ENCODING_UNSPECIFIED",
	1: "LINEAR16",
	2: "FLAC",
}
var AudioInConfig_Encoding_value = map[string]int32{
	"ENCODING_UNSPECIFIED": 0,
	"LINEAR16":             1,
	"FLAC":                 2,
}

func (x AudioInConfig_Encoding) String() string {
	return proto.EnumName(AudioInConfig_Encoding_name, int32(x))
}
func (AudioInConfig_Encoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorEmbeddedAssistant, []int{1, 0}
}

// Audio encoding of the data returned in the audio message. All encodings are
// raw audio bytes with no header, except as indicated below.
type AudioOutConfig_Encoding int32

const (
	// Not specified. Will return result [google.rpc.Code.INVALID_ARGUMENT][].
	AudioOutConfig_ENCODING_UNSPECIFIED AudioOutConfig_Encoding = 0
	// Uncompressed 16-bit signed little-endian samples (Linear PCM).
	AudioOutConfig_LINEAR16 AudioOutConfig_Encoding = 1
	// MP3 audio encoding. The sample rate is encoded in the payload.
	AudioOutConfig_MP3 AudioOutConfig_Encoding = 2
	// Opus-encoded audio wrapped in an ogg container. The result will be a
	// file which can be played natively on Android and in some browsers (such
	// as Chrome). The quality of the encoding is considerably higher than MP3
	// while using the same bitrate. The sample rate is encoded in the payload.
	AudioOutConfig_OPUS_IN_OGG AudioOutConfig_Encoding = 3
)

var AudioOutConfig_Encoding_name = map[int32]string{
	0: "ENCODING_UNSPECIFIED",
	1: "LINEAR16",
	2: "MP3",
	3: "OPUS_IN_OGG",
}
var AudioOutConfig_Encoding_value = map[string]int32{
	"ENCODING_UNSPECIFIED": 0,
	"LINEAR16":             1,
	"MP3":                  2,
	"OPUS_IN_OGG":          3,
}

func (x AudioOutConfig_Encoding) String() string {
	return proto.EnumName(AudioOutConfig_Encoding_name, int32(x))
}
func (AudioOutConfig_Encoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorEmbeddedAssistant, []int{2, 0}
}

// Possible states of the microphone after a `Converse` RPC completes.
type ConverseResult_MicrophoneMode int32

const (
	// No mode specified.
	ConverseResult_MICROPHONE_MODE_UNSPECIFIED ConverseResult_MicrophoneMode = 0
	// The service is not expecting a follow-on question from the user.
	// The microphone should remain off until the user re-activates it.
	ConverseResult_CLOSE_MICROPHONE ConverseResult_MicrophoneMode = 1
	// The service is expecting a follow-on question from the user. The
	// microphone should be re-opened when the `AudioOut` playback completes
	// (by starting a new `Converse` RPC call to send the new audio).
	ConverseResult_DIALOG_FOLLOW_ON ConverseResult_MicrophoneMode = 2
)

var ConverseResult_MicrophoneMode_name = map[int32]string{
	0: "MICROPHONE_MODE_UNSPECIFIED",
	1: "CLOSE_MICROPHONE",
	2: "DIALOG_FOLLOW_ON",
}
var ConverseResult_MicrophoneMode_value = map[string]int32{
	"MICROPHONE_MODE_UNSPECIFIED": 0,
	"CLOSE_MICROPHONE":            1,
	"DIALOG_FOLLOW_ON":            2,
}

func (x ConverseResult_MicrophoneMode) String() string {
	return proto.EnumName(ConverseResult_MicrophoneMode_name, int32(x))
}
func (ConverseResult_MicrophoneMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorEmbeddedAssistant, []int{5, 0}
}

// Indicates the type of event.
type ConverseResponse_EventType int32

const (
	// No event specified.
	ConverseResponse_EVENT_TYPE_UNSPECIFIED ConverseResponse_EventType = 0
	// This event indicates that the server has detected the end of the user's
	// speech utterance and expects no additional speech. Therefore, the server
	// will not process additional audio (although it may subsequently return
	// additional results). The client should stop sending additional audio
	// data, half-close the gRPC connection, and wait for any additional results
	// until the server closes the gRPC connection.
	ConverseResponse_END_OF_UTTERANCE ConverseResponse_EventType = 1
)

var ConverseResponse_EventType_name = map[int32]string{
	0: "EVENT_TYPE_UNSPECIFIED",
	1: "END_OF_UTTERANCE",
}
var ConverseResponse_EventType_value = map[string]int32{
	"EVENT_TYPE_UNSPECIFIED": 0,
	"END_OF_UTTERANCE":       1,
}

func (x ConverseResponse_EventType) String() string {
	return proto.EnumName(ConverseResponse_EventType_name, int32(x))
}
func (ConverseResponse_EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorEmbeddedAssistant, []int{7, 0}
}

// Specifies how to process the `ConverseRequest` messages.
type ConverseConfig struct {
	// *Required* Specifies how to process the subsequent incoming audio.
	AudioInConfig *AudioInConfig `protobuf:"bytes,1,opt,name=audio_in_config,json=audioInConfig" json:"audio_in_config,omitempty"`
	// *Required* Specifies how to format the audio that will be returned.
	AudioOutConfig *AudioOutConfig `protobuf:"bytes,2,opt,name=audio_out_config,json=audioOutConfig" json:"audio_out_config,omitempty"`
	// *Required* Represents the current dialog state.
	ConverseState *ConverseState `protobuf:"bytes,3,opt,name=converse_state,json=converseState" json:"converse_state,omitempty"`
}

func (m *ConverseConfig) Reset()                    { *m = ConverseConfig{} }
func (m *ConverseConfig) String() string            { return proto.CompactTextString(m) }
func (*ConverseConfig) ProtoMessage()               {}
func (*ConverseConfig) Descriptor() ([]byte, []int) { return fileDescriptorEmbeddedAssistant, []int{0} }

func (m *ConverseConfig) GetAudioInConfig() *AudioInConfig {
	if m != nil {
		return m.AudioInConfig
	}
	return nil
}

func (m *ConverseConfig) GetAudioOutConfig() *AudioOutConfig {
	if m != nil {
		return m.AudioOutConfig
	}
	return nil
}

func (m *ConverseConfig) GetConverseState() *ConverseState {
	if m != nil {
		return m.ConverseState
	}
	return nil
}

// Specifies how to process the `audio_in` data that will be provided in
// subsequent requests. For recommended settings, see the Google Assistant SDK
// [best practices](https://developers.google.com/assistant/sdk/develop/grpc/best-practices/audio).
type AudioInConfig struct {
	// *Required* Encoding of audio data sent in all `audio_in` messages.
	Encoding AudioInConfig_Encoding `protobuf:"varint,1,opt,name=encoding,proto3,enum=google.assistant.embedded.v1alpha1.AudioInConfig_Encoding" json:"encoding,omitempty"`
	// *Required* Sample rate (in Hertz) of the audio data sent in all `audio_in`
	// messages. Valid values are from 16000-24000, but 16000 is optimal.
	// For best results, set the sampling rate of the audio source to 16000 Hz.
	// If that's not possible, use the native sample rate of the audio source
	// (instead of re-sampling).
	SampleRateHertz int32 `protobuf:"varint,2,opt,name=sample_rate_hertz,json=sampleRateHertz,proto3" json:"sample_rate_hertz,omitempty"`
}

func (m *AudioInConfig) Reset()                    { *m = AudioInConfig{} }
func (m *AudioInConfig) String() string            { return proto.CompactTextString(m) }
func (*AudioInConfig) ProtoMessage()               {}
func (*AudioInConfig) Descriptor() ([]byte, []int) { return fileDescriptorEmbeddedAssistant, []int{1} }

func (m *AudioInConfig) GetEncoding() AudioInConfig_Encoding {
	if m != nil {
		return m.Encoding
	}
	return AudioInConfig_ENCODING_UNSPECIFIED
}

func (m *AudioInConfig) GetSampleRateHertz() int32 {
	if m != nil {
		return m.SampleRateHertz
	}
	return 0
}

// Specifies the desired format for the server to use when it returns
// `audio_out` messages.
type AudioOutConfig struct {
	// *Required* The encoding of audio data to be returned in all `audio_out`
	// messages.
	Encoding AudioOutConfig_Encoding `protobuf:"varint,1,opt,name=encoding,proto3,enum=google.assistant.embedded.v1alpha1.AudioOutConfig_Encoding" json:"encoding,omitempty"`
	// *Required* The sample rate in Hertz of the audio data returned in
	// `audio_out` messages. Valid values are: 16000-24000.
	SampleRateHertz int32 `protobuf:"varint,2,opt,name=sample_rate_hertz,json=sampleRateHertz,proto3" json:"sample_rate_hertz,omitempty"`
	// *Required* Current volume setting of the device's audio output.
	// Valid values are 1 to 100 (corresponding to 1% to 100%).
	VolumePercentage int32 `protobuf:"varint,3,opt,name=volume_percentage,json=volumePercentage,proto3" json:"volume_percentage,omitempty"`
}

func (m *AudioOutConfig) Reset()                    { *m = AudioOutConfig{} }
func (m *AudioOutConfig) String() string            { return proto.CompactTextString(m) }
func (*AudioOutConfig) ProtoMessage()               {}
func (*AudioOutConfig) Descriptor() ([]byte, []int) { return fileDescriptorEmbeddedAssistant, []int{2} }

func (m *AudioOutConfig) GetEncoding() AudioOutConfig_Encoding {
	if m != nil {
		return m.Encoding
	}
	return AudioOutConfig_ENCODING_UNSPECIFIED
}

func (m *AudioOutConfig) GetSampleRateHertz() int32 {
	if m != nil {
		return m.SampleRateHertz
	}
	return 0
}

func (m *AudioOutConfig) GetVolumePercentage() int32 {
	if m != nil {
		return m.VolumePercentage
	}
	return 0
}

// Provides information about the current dialog state.
type ConverseState struct {
	// *Required* The `conversation_state` value returned in the prior
	// `ConverseResponse`. Omit (do not set the field) if there was no prior
	// `ConverseResponse`. If there was a prior `ConverseResponse`, do not omit
	// this field; doing so will end that conversation (and this new request will
	// start a new conversation).
	ConversationState []byte `protobuf:"bytes,1,opt,name=conversation_state,json=conversationState,proto3" json:"conversation_state,omitempty"`
}

func (m *ConverseState) Reset()                    { *m = ConverseState{} }
func (m *ConverseState) String() string            { return proto.CompactTextString(m) }
func (*ConverseState) ProtoMessage()               {}
func (*ConverseState) Descriptor() ([]byte, []int) { return fileDescriptorEmbeddedAssistant, []int{3} }

func (m *ConverseState) GetConversationState() []byte {
	if m != nil {
		return m.ConversationState
	}
	return nil
}

// The audio containing the assistant's response to the query. Sequential chunks
// of audio data are received in sequential `ConverseResponse` messages.
type AudioOut struct {
	// *Output-only* The audio data containing the assistant's response to the
	// query. Sequential chunks of audio data are received in sequential
	// `ConverseResponse` messages.
	AudioData []byte `protobuf:"bytes,1,opt,name=audio_data,json=audioData,proto3" json:"audio_data,omitempty"`
}

func (m *AudioOut) Reset()                    { *m = AudioOut{} }
func (m *AudioOut) String() string            { return proto.CompactTextString(m) }
func (*AudioOut) ProtoMessage()               {}
func (*AudioOut) Descriptor() ([]byte, []int) { return fileDescriptorEmbeddedAssistant, []int{4} }

func (m *AudioOut) GetAudioData() []byte {
	if m != nil {
		return m.AudioData
	}
	return nil
}

// The semantic result for the user's spoken query.
type ConverseResult struct {
	// *Output-only* The recognized transcript of what the user said.
	SpokenRequestText string `protobuf:"bytes,1,opt,name=spoken_request_text,json=spokenRequestText,proto3" json:"spoken_request_text,omitempty"`
	// *Output-only* The text of the assistant's spoken response. This is only
	// returned for an IFTTT action.
	SpokenResponseText string `protobuf:"bytes,2,opt,name=spoken_response_text,json=spokenResponseText,proto3" json:"spoken_response_text,omitempty"`
	// *Output-only* State information for subsequent `ConverseRequest`. This
	// value should be saved in the client and returned in the
	// `conversation_state` with the next `ConverseRequest`. (The client does not
	// need to interpret or otherwise use this value.) There is no need to save
	// this information across device restarts.
	ConversationState []byte `protobuf:"bytes,3,opt,name=conversation_state,json=conversationState,proto3" json:"conversation_state,omitempty"`
	// *Output-only* Specifies the mode of the microphone after this `Converse`
	// RPC is processed.
	MicrophoneMode ConverseResult_MicrophoneMode `protobuf:"varint,4,opt,name=microphone_mode,json=microphoneMode,proto3,enum=google.assistant.embedded.v1alpha1.ConverseResult_MicrophoneMode" json:"microphone_mode,omitempty"`
	// *Output-only* Updated volume level. The value will be 0 or omitted
	// (indicating no change) unless a voice command such as "Increase the volume"
	// or "Set volume level 4" was recognized, in which case the value will be
	// between 1 and 100 (corresponding to the new volume level of 1% to 100%).
	// Typically, a client should use this volume level when playing the
	// `audio_out` data, and retain this value as the current volume level and
	// supply it in the `AudioOutConfig` of the next `ConverseRequest`. (Some
	// clients may also implement other ways to allow the current volume level to
	// be changed, for example, by providing a knob that the user can turn.)
	VolumePercentage int32 `protobuf:"varint,5,opt,name=volume_percentage,json=volumePercentage,proto3" json:"volume_percentage,omitempty"`
}

func (m *ConverseResult) Reset()                    { *m = ConverseResult{} }
func (m *ConverseResult) String() string            { return proto.CompactTextString(m) }
func (*ConverseResult) ProtoMessage()               {}
func (*ConverseResult) Descriptor() ([]byte, []int) { return fileDescriptorEmbeddedAssistant, []int{5} }

func (m *ConverseResult) GetSpokenRequestText() string {
	if m != nil {
		return m.SpokenRequestText
	}
	return ""
}

func (m *ConverseResult) GetSpokenResponseText() string {
	if m != nil {
		return m.SpokenResponseText
	}
	return ""
}

func (m *ConverseResult) GetConversationState() []byte {
	if m != nil {
		return m.ConversationState
	}
	return nil
}

func (m *ConverseResult) GetMicrophoneMode() ConverseResult_MicrophoneMode {
	if m != nil {
		return m.MicrophoneMode
	}
	return ConverseResult_MICROPHONE_MODE_UNSPECIFIED
}

func (m *ConverseResult) GetVolumePercentage() int32 {
	if m != nil {
		return m.VolumePercentage
	}
	return 0
}

// The top-level message sent by the client. Clients must send at least two, and
// typically numerous `ConverseRequest` messages. The first message must
// contain a `config` message and must not contain `audio_in` data. All
// subsequent messages must contain `audio_in` data and must not contain a
// `config` message.
type ConverseRequest struct {
	// Exactly one of these fields must be specified in each `ConverseRequest`.
	//
	// Types that are valid to be assigned to ConverseRequest:
	//	*ConverseRequest_Config
	//	*ConverseRequest_AudioIn
	ConverseRequest isConverseRequest_ConverseRequest `protobuf_oneof:"converse_request"`
}

func (m *ConverseRequest) Reset()                    { *m = ConverseRequest{} }
func (m *ConverseRequest) String() string            { return proto.CompactTextString(m) }
func (*ConverseRequest) ProtoMessage()               {}
func (*ConverseRequest) Descriptor() ([]byte, []int) { return fileDescriptorEmbeddedAssistant, []int{6} }

type isConverseRequest_ConverseRequest interface {
	isConverseRequest_ConverseRequest()
}

type ConverseRequest_Config struct {
	Config *ConverseConfig `protobuf:"bytes,1,opt,name=config,oneof"`
}
type ConverseRequest_AudioIn struct {
	AudioIn []byte `protobuf:"bytes,2,opt,name=audio_in,json=audioIn,proto3,oneof"`
}

func (*ConverseRequest_Config) isConverseRequest_ConverseRequest()  {}
func (*ConverseRequest_AudioIn) isConverseRequest_ConverseRequest() {}

func (m *ConverseRequest) GetConverseRequest() isConverseRequest_ConverseRequest {
	if m != nil {
		return m.ConverseRequest
	}
	return nil
}

func (m *ConverseRequest) GetConfig() *ConverseConfig {
	if x, ok := m.GetConverseRequest().(*ConverseRequest_Config); ok {
		return x.Config
	}
	return nil
}

func (m *ConverseRequest) GetAudioIn() []byte {
	if x, ok := m.GetConverseRequest().(*ConverseRequest_AudioIn); ok {
		return x.AudioIn
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ConverseRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ConverseRequest_OneofMarshaler, _ConverseRequest_OneofUnmarshaler, _ConverseRequest_OneofSizer, []interface{}{
		(*ConverseRequest_Config)(nil),
		(*ConverseRequest_AudioIn)(nil),
	}
}

func _ConverseRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ConverseRequest)
	// converse_request
	switch x := m.ConverseRequest.(type) {
	case *ConverseRequest_Config:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Config); err != nil {
			return err
		}
	case *ConverseRequest_AudioIn:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.AudioIn)
	case nil:
	default:
		return fmt.Errorf("ConverseRequest.ConverseRequest has unexpected type %T", x)
	}
	return nil
}

func _ConverseRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ConverseRequest)
	switch tag {
	case 1: // converse_request.config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ConverseConfig)
		err := b.DecodeMessage(msg)
		m.ConverseRequest = &ConverseRequest_Config{msg}
		return true, err
	case 2: // converse_request.audio_in
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.ConverseRequest = &ConverseRequest_AudioIn{x}
		return true, err
	default:
		return false, nil
	}
}

func _ConverseRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ConverseRequest)
	// converse_request
	switch x := m.ConverseRequest.(type) {
	case *ConverseRequest_Config:
		s := proto.Size(x.Config)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConverseRequest_AudioIn:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AudioIn)))
		n += len(x.AudioIn)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The top-level message received by the client. A series of one or more
// `ConverseResponse` messages are streamed back to the client.
type ConverseResponse struct {
	// Exactly one of these fields will be populated in each `ConverseResponse`.
	//
	// Types that are valid to be assigned to ConverseResponse:
	//	*ConverseResponse_Error
	//	*ConverseResponse_EventType_
	//	*ConverseResponse_AudioOut
	//	*ConverseResponse_Result
	ConverseResponse isConverseResponse_ConverseResponse `protobuf_oneof:"converse_response"`
}

func (m *ConverseResponse) Reset()         { *m = ConverseResponse{} }
func (m *ConverseResponse) String() string { return proto.CompactTextString(m) }
func (*ConverseResponse) ProtoMessage()    {}
func (*ConverseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorEmbeddedAssistant, []int{7}
}

type isConverseResponse_ConverseResponse interface {
	isConverseResponse_ConverseResponse()
}

type ConverseResponse_Error struct {
	Error *google_rpc.Status `protobuf:"bytes,1,opt,name=error,oneof"`
}
type ConverseResponse_EventType_ struct {
	EventType ConverseResponse_EventType `protobuf:"varint,2,opt,name=event_type,json=eventType,proto3,enum=google.assistant.embedded.v1alpha1.ConverseResponse_EventType,oneof"`
}
type ConverseResponse_AudioOut struct {
	AudioOut *AudioOut `protobuf:"bytes,3,opt,name=audio_out,json=audioOut,oneof"`
}
type ConverseResponse_Result struct {
	Result *ConverseResult `protobuf:"bytes,5,opt,name=result,oneof"`
}

func (*ConverseResponse_Error) isConverseResponse_ConverseResponse()      {}
func (*ConverseResponse_EventType_) isConverseResponse_ConverseResponse() {}
func (*ConverseResponse_AudioOut) isConverseResponse_ConverseResponse()   {}
func (*ConverseResponse_Result) isConverseResponse_ConverseResponse()     {}

func (m *ConverseResponse) GetConverseResponse() isConverseResponse_ConverseResponse {
	if m != nil {
		return m.ConverseResponse
	}
	return nil
}

func (m *ConverseResponse) GetError() *google_rpc.Status {
	if x, ok := m.GetConverseResponse().(*ConverseResponse_Error); ok {
		return x.Error
	}
	return nil
}

func (m *ConverseResponse) GetEventType() ConverseResponse_EventType {
	if x, ok := m.GetConverseResponse().(*ConverseResponse_EventType_); ok {
		return x.EventType
	}
	return ConverseResponse_EVENT_TYPE_UNSPECIFIED
}

func (m *ConverseResponse) GetAudioOut() *AudioOut {
	if x, ok := m.GetConverseResponse().(*ConverseResponse_AudioOut); ok {
		return x.AudioOut
	}
	return nil
}

func (m *ConverseResponse) GetResult() *ConverseResult {
	if x, ok := m.GetConverseResponse().(*ConverseResponse_Result); ok {
		return x.Result
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ConverseResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ConverseResponse_OneofMarshaler, _ConverseResponse_OneofUnmarshaler, _ConverseResponse_OneofSizer, []interface{}{
		(*ConverseResponse_Error)(nil),
		(*ConverseResponse_EventType_)(nil),
		(*ConverseResponse_AudioOut)(nil),
		(*ConverseResponse_Result)(nil),
	}
}

func _ConverseResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ConverseResponse)
	// converse_response
	switch x := m.ConverseResponse.(type) {
	case *ConverseResponse_Error:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *ConverseResponse_EventType_:
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.EventType))
	case *ConverseResponse_AudioOut:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AudioOut); err != nil {
			return err
		}
	case *ConverseResponse_Result:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Result); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ConverseResponse.ConverseResponse has unexpected type %T", x)
	}
	return nil
}

func _ConverseResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ConverseResponse)
	switch tag {
	case 1: // converse_response.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_rpc.Status)
		err := b.DecodeMessage(msg)
		m.ConverseResponse = &ConverseResponse_Error{msg}
		return true, err
	case 2: // converse_response.event_type
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ConverseResponse = &ConverseResponse_EventType_{ConverseResponse_EventType(x)}
		return true, err
	case 3: // converse_response.audio_out
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AudioOut)
		err := b.DecodeMessage(msg)
		m.ConverseResponse = &ConverseResponse_AudioOut{msg}
		return true, err
	case 5: // converse_response.result
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ConverseResult)
		err := b.DecodeMessage(msg)
		m.ConverseResponse = &ConverseResponse_Result{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ConverseResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ConverseResponse)
	// converse_response
	switch x := m.ConverseResponse.(type) {
	case *ConverseResponse_Error:
		s := proto.Size(x.Error)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConverseResponse_EventType_:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.EventType))
	case *ConverseResponse_AudioOut:
		s := proto.Size(x.AudioOut)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConverseResponse_Result:
		s := proto.Size(x.Result)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ConverseConfig)(nil), "google.assistant.embedded.v1alpha1.ConverseConfig")
	proto.RegisterType((*AudioInConfig)(nil), "google.assistant.embedded.v1alpha1.AudioInConfig")
	proto.RegisterType((*AudioOutConfig)(nil), "google.assistant.embedded.v1alpha1.AudioOutConfig")
	proto.RegisterType((*ConverseState)(nil), "google.assistant.embedded.v1alpha1.ConverseState")
	proto.RegisterType((*AudioOut)(nil), "google.assistant.embedded.v1alpha1.AudioOut")
	proto.RegisterType((*ConverseResult)(nil), "google.assistant.embedded.v1alpha1.ConverseResult")
	proto.RegisterType((*ConverseRequest)(nil), "google.assistant.embedded.v1alpha1.ConverseRequest")
	proto.RegisterType((*ConverseResponse)(nil), "google.assistant.embedded.v1alpha1.ConverseResponse")
	proto.RegisterEnum("google.assistant.embedded.v1alpha1.AudioInConfig_Encoding", AudioInConfig_Encoding_name, AudioInConfig_Encoding_value)
	proto.RegisterEnum("google.assistant.embedded.v1alpha1.AudioOutConfig_Encoding", AudioOutConfig_Encoding_name, AudioOutConfig_Encoding_value)
	proto.RegisterEnum("google.assistant.embedded.v1alpha1.ConverseResult_MicrophoneMode", ConverseResult_MicrophoneMode_name, ConverseResult_MicrophoneMode_value)
	proto.RegisterEnum("google.assistant.embedded.v1alpha1.ConverseResponse_EventType", ConverseResponse_EventType_name, ConverseResponse_EventType_value)
}

func init() {
	proto.RegisterFile("google/assistant/embedded/v1alpha1/embedded_assistant.proto", fileDescriptorEmbeddedAssistant)
}

var fileDescriptorEmbeddedAssistant = []byte{
	// 863 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xe1, 0x6f, 0xdb, 0x44,
	0x14, 0x8f, 0x93, 0x75, 0x4b, 0xde, 0x5a, 0xc7, 0xb9, 0x55, 0x50, 0x75, 0x20, 0x90, 0x3f, 0xa0,
	0x31, 0xc0, 0x59, 0x53, 0xc4, 0x87, 0x0d, 0x26, 0xa5, 0x89, 0xdb, 0x44, 0x24, 0xb6, 0x75, 0x49,
	0x37, 0x26, 0x81, 0x4e, 0x37, 0xfb, 0x68, 0x0d, 0x89, 0xcf, 0xd8, 0x97, 0x6a, 0xe5, 0x8f, 0x98,
	0xf8, 0xca, 0x67, 0xfe, 0x22, 0xfe, 0x23, 0xe4, 0x3b, 0xdb, 0x4d, 0xa0, 0x13, 0x75, 0xf7, 0xf1,
	0xde, 0xbb, 0xf7, 0xf3, 0xef, 0xfd, 0xde, 0xef, 0x9e, 0x0c, 0xcf, 0xce, 0x38, 0x3f, 0x5b, 0xb0,
	0x2e, 0x4d, 0xd3, 0x30, 0x15, 0x34, 0x12, 0x5d, 0xb6, 0x7c, 0xcd, 0x82, 0x80, 0x05, 0xdd, 0x8b,
	0x03, 0xba, 0x88, 0xcf, 0xe9, 0x41, 0x19, 0x21, 0xe5, 0x25, 0x2b, 0x4e, 0xb8, 0xe0, 0xc8, 0x54,
	0xc5, 0xd6, 0x55, 0xbc, 0xb8, 0x6a, 0x15, 0xc5, 0xfb, 0x1f, 0x15, 0x1f, 0x88, 0xc3, 0x2e, 0x8d,
	0x22, 0x2e, 0xa8, 0x08, 0x79, 0x94, 0x2a, 0x84, 0xfd, 0x0f, 0xf3, 0x6c, 0x12, 0xfb, 0xdd, 0x54,
	0x50, 0xb1, 0xca, 0x13, 0xe6, 0x5f, 0x75, 0xd0, 0x07, 0x3c, 0xba, 0x60, 0x49, 0xca, 0x06, 0x3c,
	0xfa, 0x39, 0x3c, 0x43, 0xaf, 0xa0, 0x4d, 0x57, 0x41, 0xc8, 0x49, 0x18, 0x11, 0x5f, 0x86, 0xf6,
	0xb4, 0x4f, 0xb5, 0x47, 0xf7, 0x7b, 0x07, 0xd6, 0xff, 0xf3, 0xb0, 0xfa, 0x59, 0xe9, 0x38, 0x52,
	0x58, 0x78, 0x87, 0xae, 0x1f, 0xd1, 0x8f, 0x60, 0x28, 0x68, 0xbe, 0x12, 0x05, 0x76, 0x5d, 0x62,
	0xf7, 0x6e, 0x8c, 0xed, 0xae, 0x44, 0x0e, 0xae, 0xd3, 0x8d, 0x33, 0xfa, 0x01, 0x74, 0x3f, 0x6f,
	0x85, 0x64, 0x4d, 0xb2, 0xbd, 0xc6, 0xcd, 0x79, 0x17, 0x22, 0xcc, 0xb2, 0x42, 0xbc, 0xe3, 0xaf,
	0x1f, 0xcd, 0xbf, 0x35, 0xd8, 0xd9, 0x68, 0x0c, 0xbd, 0x80, 0x26, 0x8b, 0x7c, 0x1e, 0x84, 0x91,
	0x52, 0x47, 0xef, 0x3d, 0xad, 0xac, 0x8e, 0x65, 0xe7, 0x08, 0xb8, 0xc4, 0x42, 0x8f, 0xa1, 0x93,
	0xd2, 0x65, 0xbc, 0x60, 0x24, 0xa1, 0x82, 0x91, 0x73, 0x96, 0x88, 0xdf, 0xa5, 0x44, 0x5b, 0xb8,
	0xad, 0x12, 0x98, 0x0a, 0x36, 0xca, 0xc2, 0xe6, 0xb7, 0xd0, 0x2c, 0x10, 0xd0, 0x1e, 0xec, 0xda,
	0xce, 0xc0, 0x1d, 0x8e, 0x9d, 0x13, 0x72, 0xea, 0xcc, 0x3c, 0x7b, 0x30, 0x3e, 0x1e, 0xdb, 0x43,
	0xa3, 0x86, 0xb6, 0xa1, 0x39, 0x19, 0x3b, 0x76, 0x1f, 0x1f, 0x7c, 0x63, 0x68, 0xa8, 0x09, 0x77,
	0x8e, 0x27, 0xfd, 0x81, 0x51, 0x37, 0xff, 0xa8, 0x83, 0xbe, 0x29, 0x28, 0x7a, 0xf9, 0x9f, 0xa6,
	0x9e, 0x55, 0x1f, 0xcb, 0x7b, 0x76, 0x85, 0xbe, 0x80, 0xce, 0x05, 0x5f, 0xac, 0x96, 0x8c, 0xc4,
	0x2c, 0xf1, 0x59, 0x24, 0xe8, 0x99, 0x1a, 0xe4, 0x16, 0x36, 0x54, 0xc2, 0x2b, 0xe3, 0xe6, 0xe4,
	0x16, 0x12, 0xdc, 0x83, 0xc6, 0xd4, 0x3b, 0x34, 0xea, 0xa8, 0x0d, 0xf7, 0x5d, 0xef, 0x74, 0x46,
	0xc6, 0x0e, 0x71, 0x4f, 0x4e, 0x8c, 0x86, 0xf9, 0x1c, 0x76, 0x36, 0x6c, 0x80, 0xbe, 0x02, 0x94,
	0x1b, 0x41, 0xbe, 0xa6, 0xdc, 0x55, 0x99, 0x34, 0xdb, 0xb8, 0xb3, 0x9e, 0x51, 0x36, 0xf9, 0x1c,
	0x9a, 0x85, 0x16, 0xe8, 0x63, 0x00, 0x65, 0xf5, 0x80, 0x0a, 0x9a, 0x97, 0xb4, 0x64, 0x64, 0x48,
	0x05, 0x35, 0xff, 0x6c, 0x5c, 0xbd, 0x3b, 0xcc, 0xd2, 0xd5, 0x42, 0x20, 0x0b, 0x1e, 0xa4, 0x31,
	0xff, 0x95, 0x45, 0x24, 0x61, 0xbf, 0xad, 0x58, 0x2a, 0x88, 0x60, 0x6f, 0x84, 0x2c, 0x6d, 0xe1,
	0x8e, 0x4a, 0x61, 0x95, 0x99, 0xb3, 0x37, 0x02, 0x3d, 0x81, 0xdd, 0xf2, 0x7e, 0x1a, 0xf3, 0x28,
	0x65, 0xaa, 0xa0, 0x2e, 0x0b, 0x50, 0x51, 0xa0, 0x52, 0xb2, 0xe2, 0xfa, 0x76, 0x1a, 0xef, 0x68,
	0x07, 0xfd, 0x02, 0xed, 0x65, 0xe8, 0x27, 0x3c, 0x3e, 0xe7, 0x11, 0x23, 0x4b, 0x1e, 0xb0, 0xbd,
	0x3b, 0xd2, 0x15, 0xfd, 0x2a, 0x0f, 0x4a, 0x75, 0x67, 0x4d, 0x4b, 0xa4, 0x29, 0x0f, 0x18, 0xd6,
	0x97, 0x1b, 0xe7, 0xeb, 0xa7, 0xbe, 0xf5, 0x8e, 0xa9, 0xff, 0x04, 0xfa, 0x26, 0x1c, 0xfa, 0x04,
	0x1e, 0x4e, 0xc7, 0x03, 0xec, 0x7a, 0x23, 0xd7, 0xb1, 0xc9, 0xd4, 0x1d, 0xda, 0xff, 0xb2, 0xc0,
	0x2e, 0x18, 0x83, 0x89, 0x3b, 0xb3, 0xc9, 0xd5, 0x35, 0x43, 0xcb, 0xa2, 0xc3, 0x71, 0x7f, 0xe2,
	0x9e, 0x90, 0x63, 0x77, 0x32, 0x71, 0x5f, 0x12, 0xd7, 0xc9, 0x5e, 0x86, 0x06, 0xed, 0x2b, 0xf6,
	0x52, 0x70, 0x34, 0x81, 0xbb, 0x1b, 0xbb, 0xb0, 0x57, 0x45, 0x02, 0xf5, 0x30, 0x46, 0x35, 0x9c,
	0x63, 0xa0, 0x87, 0xd0, 0x2c, 0x56, 0xac, 0x1c, 0xd7, 0xf6, 0xa8, 0x86, 0xef, 0xe5, 0xab, 0xf2,
	0x08, 0x81, 0x51, 0xae, 0xb1, 0xdc, 0x09, 0xe6, 0xdb, 0x06, 0x18, 0x6b, 0x82, 0xca, 0x91, 0xa2,
	0xc7, 0xb0, 0xc5, 0x92, 0x84, 0x27, 0x39, 0x25, 0x54, 0x50, 0x4a, 0x62, 0xdf, 0x9a, 0xc9, 0x25,
	0x3f, 0xaa, 0x61, 0x75, 0x05, 0x11, 0x00, 0x76, 0xc1, 0x22, 0x41, 0xc4, 0x65, 0xcc, 0xe4, 0x37,
	0xf5, 0xde, 0xf3, 0x8a, 0x63, 0x94, 0x5f, 0xb5, 0xec, 0x0c, 0x66, 0x7e, 0x19, 0xb3, 0x51, 0x0d,
	0xb7, 0x58, 0x71, 0x40, 0xdf, 0x43, 0xab, 0x5c, 0xed, 0xf9, 0xde, 0xfd, 0xb2, 0xca, 0xf2, 0x18,
	0xd5, 0x70, 0xb3, 0xd8, 0xe7, 0x99, 0xda, 0x89, 0xb4, 0x8d, 0xb4, 0x40, 0x45, 0xb5, 0x95, 0xe1,
	0x32, 0xb5, 0x15, 0x86, 0xf9, 0x1d, 0xb4, 0x4a, 0xd2, 0x68, 0x1f, 0x3e, 0xb0, 0x5f, 0xd8, 0xce,
	0x9c, 0xcc, 0x5f, 0x79, 0xd7, 0x98, 0xc4, 0x76, 0x86, 0xc4, 0x3d, 0x26, 0xa7, 0xf3, 0xb9, 0x8d,
	0xfb, 0xce, 0xc0, 0x36, 0xb4, 0xa3, 0x07, 0xd0, 0x59, 0x9b, 0x87, 0x52, 0xa1, 0xf7, 0x56, 0x83,
	0x8e, 0x9d, 0x53, 0xe8, 0x17, 0xa4, 0xd0, 0x25, 0x34, 0x0b, 0x16, 0xe8, 0xb0, 0x1a, 0x67, 0x39,
	0xe7, 0xfd, 0xaf, 0x6f, 0x33, 0x92, 0x47, 0xda, 0x13, 0xed, 0xe8, 0x29, 0x7c, 0xe6, 0xf3, 0xe5,
	0x0d, 0xca, 0x8f, 0xf4, 0x92, 0xaf, 0x97, 0xfd, 0x02, 0x78, 0xda, 0xeb, 0xbb, 0xf2, 0x5f, 0xe0,
	0xf0, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xf8, 0xe0, 0xf1, 0xa5, 0x08, 0x00, 0x00,
}