// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/cloud/vision/v1/text_annotation.proto

package google_cloud_vision_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Enum to denote the type of break found. New line, space etc.
type TextAnnotation_DetectedBreak_BreakType int32

const (
	// Unknown break label type.
	TextAnnotation_DetectedBreak_UNKNOWN TextAnnotation_DetectedBreak_BreakType = 0
	// Regular space.
	TextAnnotation_DetectedBreak_SPACE TextAnnotation_DetectedBreak_BreakType = 1
	// Sure space (very wide).
	TextAnnotation_DetectedBreak_SURE_SPACE TextAnnotation_DetectedBreak_BreakType = 2
	// Line-wrapping break.
	TextAnnotation_DetectedBreak_EOL_SURE_SPACE TextAnnotation_DetectedBreak_BreakType = 3
	// End-line hyphen that is not present in text; does
	TextAnnotation_DetectedBreak_HYPHEN TextAnnotation_DetectedBreak_BreakType = 4
	// not co-occur with SPACE, LEADER_SPACE, or
	// LINE_BREAK.
	// Line break that ends a paragraph.
	TextAnnotation_DetectedBreak_LINE_BREAK TextAnnotation_DetectedBreak_BreakType = 5
)

var TextAnnotation_DetectedBreak_BreakType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SPACE",
	2: "SURE_SPACE",
	3: "EOL_SURE_SPACE",
	4: "HYPHEN",
	5: "LINE_BREAK",
}
var TextAnnotation_DetectedBreak_BreakType_value = map[string]int32{
	"UNKNOWN":        0,
	"SPACE":          1,
	"SURE_SPACE":     2,
	"EOL_SURE_SPACE": 3,
	"HYPHEN":         4,
	"LINE_BREAK":     5,
}

func (x TextAnnotation_DetectedBreak_BreakType) String() string {
	return proto.EnumName(TextAnnotation_DetectedBreak_BreakType_name, int32(x))
}
func (TextAnnotation_DetectedBreak_BreakType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorTextAnnotation, []int{0, 1, 0}
}

// Type of a block (text, image etc) as identified by OCR.
type Block_BlockType int32

const (
	// Unknown block type.
	Block_UNKNOWN Block_BlockType = 0
	// Regular text block.
	Block_TEXT Block_BlockType = 1
	// Table block.
	Block_TABLE Block_BlockType = 2
	// Image block.
	Block_PICTURE Block_BlockType = 3
	// Horizontal/vertical line box.
	Block_RULER Block_BlockType = 4
	// Barcode block.
	Block_BARCODE Block_BlockType = 5
)

var Block_BlockType_name = map[int32]string{
	0: "UNKNOWN",
	1: "TEXT",
	2: "TABLE",
	3: "PICTURE",
	4: "RULER",
	5: "BARCODE",
}
var Block_BlockType_value = map[string]int32{
	"UNKNOWN": 0,
	"TEXT":    1,
	"TABLE":   2,
	"PICTURE": 3,
	"RULER":   4,
	"BARCODE": 5,
}

func (x Block_BlockType) String() string {
	return proto.EnumName(Block_BlockType_name, int32(x))
}
func (Block_BlockType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorTextAnnotation, []int{2, 0}
}

// TextAnnotation contains a structured representation of OCR extracted text.
// The hierarchy of an OCR extracted text structure is like this:
//     TextAnnotation -> Page -> Block -> Paragraph -> Word -> Symbol
// Each structural component, starting from Page, may further have their own
// properties. Properties describe detected languages, breaks etc.. Please
// refer to the [google.cloud.vision.v1.TextAnnotation.TextProperty][google.cloud.vision.v1.TextAnnotation.TextProperty] message
// definition below for more detail.
type TextAnnotation struct {
	// List of pages detected by OCR.
	Pages []*Page `protobuf:"bytes,1,rep,name=pages" json:"pages,omitempty"`
	// UTF-8 text detected on the pages.
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *TextAnnotation) Reset()                    { *m = TextAnnotation{} }
func (m *TextAnnotation) String() string            { return proto.CompactTextString(m) }
func (*TextAnnotation) ProtoMessage()               {}
func (*TextAnnotation) Descriptor() ([]byte, []int) { return fileDescriptorTextAnnotation, []int{0} }

func (m *TextAnnotation) GetPages() []*Page {
	if m != nil {
		return m.Pages
	}
	return nil
}

func (m *TextAnnotation) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// Detected language for a structural component.
type TextAnnotation_DetectedLanguage struct {
	// The BCP-47 language code, such as "en-US" or "sr-Latn". For more
	// information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode string `protobuf:"bytes,1,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Confidence of detected language. Range [0, 1].
	Confidence float32 `protobuf:"fixed32,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (m *TextAnnotation_DetectedLanguage) Reset()         { *m = TextAnnotation_DetectedLanguage{} }
func (m *TextAnnotation_DetectedLanguage) String() string { return proto.CompactTextString(m) }
func (*TextAnnotation_DetectedLanguage) ProtoMessage()    {}
func (*TextAnnotation_DetectedLanguage) Descriptor() ([]byte, []int) {
	return fileDescriptorTextAnnotation, []int{0, 0}
}

func (m *TextAnnotation_DetectedLanguage) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *TextAnnotation_DetectedLanguage) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// Detected start or end of a structural component.
type TextAnnotation_DetectedBreak struct {
	Type TextAnnotation_DetectedBreak_BreakType `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.vision.v1.TextAnnotation_DetectedBreak_BreakType" json:"type,omitempty"`
	// True if break prepends the element.
	IsPrefix bool `protobuf:"varint,2,opt,name=is_prefix,json=isPrefix,proto3" json:"is_prefix,omitempty"`
}

func (m *TextAnnotation_DetectedBreak) Reset()         { *m = TextAnnotation_DetectedBreak{} }
func (m *TextAnnotation_DetectedBreak) String() string { return proto.CompactTextString(m) }
func (*TextAnnotation_DetectedBreak) ProtoMessage()    {}
func (*TextAnnotation_DetectedBreak) Descriptor() ([]byte, []int) {
	return fileDescriptorTextAnnotation, []int{0, 1}
}

func (m *TextAnnotation_DetectedBreak) GetType() TextAnnotation_DetectedBreak_BreakType {
	if m != nil {
		return m.Type
	}
	return TextAnnotation_DetectedBreak_UNKNOWN
}

func (m *TextAnnotation_DetectedBreak) GetIsPrefix() bool {
	if m != nil {
		return m.IsPrefix
	}
	return false
}

// Additional information detected on the structural component.
type TextAnnotation_TextProperty struct {
	// A list of detected languages together with confidence.
	DetectedLanguages []*TextAnnotation_DetectedLanguage `protobuf:"bytes,1,rep,name=detected_languages,json=detectedLanguages" json:"detected_languages,omitempty"`
	// Detected start or end of a text segment.
	DetectedBreak *TextAnnotation_DetectedBreak `protobuf:"bytes,2,opt,name=detected_break,json=detectedBreak" json:"detected_break,omitempty"`
}

func (m *TextAnnotation_TextProperty) Reset()         { *m = TextAnnotation_TextProperty{} }
func (m *TextAnnotation_TextProperty) String() string { return proto.CompactTextString(m) }
func (*TextAnnotation_TextProperty) ProtoMessage()    {}
func (*TextAnnotation_TextProperty) Descriptor() ([]byte, []int) {
	return fileDescriptorTextAnnotation, []int{0, 2}
}

func (m *TextAnnotation_TextProperty) GetDetectedLanguages() []*TextAnnotation_DetectedLanguage {
	if m != nil {
		return m.DetectedLanguages
	}
	return nil
}

func (m *TextAnnotation_TextProperty) GetDetectedBreak() *TextAnnotation_DetectedBreak {
	if m != nil {
		return m.DetectedBreak
	}
	return nil
}

// Detected page from OCR.
type Page struct {
	// Additional information detected on the page.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// Page width in pixels.
	Width int32 `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	// Page height in pixels.
	Height int32 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	// List of blocks of text, images etc on this page.
	Blocks []*Block `protobuf:"bytes,4,rep,name=blocks" json:"blocks,omitempty"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptorTextAnnotation, []int{1} }

func (m *Page) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Page) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Page) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Page) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

// Logical element on the page.
type Block struct {
	// Additional information detected for the block.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// The bounding box for the block.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox" json:"bounding_box,omitempty"`
	// List of paragraphs in this block (if this blocks is of type text).
	Paragraphs []*Paragraph `protobuf:"bytes,3,rep,name=paragraphs" json:"paragraphs,omitempty"`
	// Detected block type (text, image etc) for this block.
	BlockType Block_BlockType `protobuf:"varint,4,opt,name=block_type,json=blockType,proto3,enum=google.cloud.vision.v1.Block_BlockType" json:"block_type,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptorTextAnnotation, []int{2} }

func (m *Block) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Block) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Block) GetParagraphs() []*Paragraph {
	if m != nil {
		return m.Paragraphs
	}
	return nil
}

func (m *Block) GetBlockType() Block_BlockType {
	if m != nil {
		return m.BlockType
	}
	return Block_UNKNOWN
}

// Structural unit of text representing a number of words in certain order.
type Paragraph struct {
	// Additional information detected for the paragraph.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// The bounding box for the paragraph.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox" json:"bounding_box,omitempty"`
	// List of words in this paragraph.
	Words []*Word `protobuf:"bytes,3,rep,name=words" json:"words,omitempty"`
}

func (m *Paragraph) Reset()                    { *m = Paragraph{} }
func (m *Paragraph) String() string            { return proto.CompactTextString(m) }
func (*Paragraph) ProtoMessage()               {}
func (*Paragraph) Descriptor() ([]byte, []int) { return fileDescriptorTextAnnotation, []int{3} }

func (m *Paragraph) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Paragraph) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Paragraph) GetWords() []*Word {
	if m != nil {
		return m.Words
	}
	return nil
}

// A word representation.
type Word struct {
	// Additional information detected for the word.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// The bounding box for the word.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox" json:"bounding_box,omitempty"`
	// List of symbols in the word.
	// The order of the symbols follows the natural reading order.
	Symbols []*Symbol `protobuf:"bytes,3,rep,name=symbols" json:"symbols,omitempty"`
}

func (m *Word) Reset()                    { *m = Word{} }
func (m *Word) String() string            { return proto.CompactTextString(m) }
func (*Word) ProtoMessage()               {}
func (*Word) Descriptor() ([]byte, []int) { return fileDescriptorTextAnnotation, []int{4} }

func (m *Word) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Word) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Word) GetSymbols() []*Symbol {
	if m != nil {
		return m.Symbols
	}
	return nil
}

// A single symbol representation.
type Symbol struct {
	// Additional information detected for the symbol.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// The bounding box for the symbol.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox" json:"bounding_box,omitempty"`
	// The actual UTF-8 representation of the symbol.
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *Symbol) Reset()                    { *m = Symbol{} }
func (m *Symbol) String() string            { return proto.CompactTextString(m) }
func (*Symbol) ProtoMessage()               {}
func (*Symbol) Descriptor() ([]byte, []int) { return fileDescriptorTextAnnotation, []int{5} }

func (m *Symbol) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Symbol) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Symbol) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*TextAnnotation)(nil), "google.cloud.vision.v1.TextAnnotation")
	proto.RegisterType((*TextAnnotation_DetectedLanguage)(nil), "google.cloud.vision.v1.TextAnnotation.DetectedLanguage")
	proto.RegisterType((*TextAnnotation_DetectedBreak)(nil), "google.cloud.vision.v1.TextAnnotation.DetectedBreak")
	proto.RegisterType((*TextAnnotation_TextProperty)(nil), "google.cloud.vision.v1.TextAnnotation.TextProperty")
	proto.RegisterType((*Page)(nil), "google.cloud.vision.v1.Page")
	proto.RegisterType((*Block)(nil), "google.cloud.vision.v1.Block")
	proto.RegisterType((*Paragraph)(nil), "google.cloud.vision.v1.Paragraph")
	proto.RegisterType((*Word)(nil), "google.cloud.vision.v1.Word")
	proto.RegisterType((*Symbol)(nil), "google.cloud.vision.v1.Symbol")
	proto.RegisterEnum("google.cloud.vision.v1.TextAnnotation_DetectedBreak_BreakType", TextAnnotation_DetectedBreak_BreakType_name, TextAnnotation_DetectedBreak_BreakType_value)
	proto.RegisterEnum("google.cloud.vision.v1.Block_BlockType", Block_BlockType_name, Block_BlockType_value)
}

func init() {
	proto.RegisterFile("google/cloud/vision/v1/text_annotation.proto", fileDescriptorTextAnnotation)
}

var fileDescriptorTextAnnotation = []byte{
	// 717 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcf, 0x6f, 0xd3, 0x4a,
	0x10, 0x7e, 0xdb, 0xd8, 0x69, 0x3c, 0x69, 0x23, 0xbf, 0x7d, 0x4f, 0x55, 0x14, 0x4a, 0x55, 0x0c,
	0x88, 0x1e, 0x90, 0xa3, 0xa6, 0xfc, 0x3a, 0x21, 0xc5, 0xa9, 0xa1, 0x55, 0xa3, 0xc4, 0xda, 0x26,
	0x2a, 0x88, 0x83, 0xe5, 0xd8, 0x5b, 0xc7, 0x6a, 0xea, 0xb5, 0x6c, 0xb7, 0x4d, 0x6e, 0xfc, 0x55,
	0x9c, 0xf8, 0x2f, 0x38, 0xc1, 0x9d, 0x33, 0x57, 0x8e, 0xc8, 0x6b, 0x3b, 0x4d, 0x10, 0x46, 0x80,
	0x38, 0xf4, 0x62, 0xed, 0x8c, 0xbe, 0xf9, 0xf6, 0xfb, 0x66, 0x32, 0x59, 0x78, 0xe8, 0x32, 0xe6,
	0x4e, 0x68, 0xd3, 0x9e, 0xb0, 0x0b, 0xa7, 0x79, 0xe9, 0x45, 0x1e, 0xf3, 0x9b, 0x97, 0xbb, 0xcd,
	0x98, 0x4e, 0x63, 0xd3, 0xf2, 0x7d, 0x16, 0x5b, 0xb1, 0xc7, 0x7c, 0x35, 0x08, 0x59, 0xcc, 0xf0,
	0x46, 0x8a, 0x56, 0x39, 0x5a, 0x4d, 0xd1, 0xea, 0xe5, 0x6e, 0x63, 0x33, 0x63, 0xb1, 0x02, 0xaf,
	0x79, 0x5d, 0x14, 0xa5, 0x55, 0x8d, 0xfb, 0x05, 0x77, 0xb8, 0x94, 0x9d, 0xd3, 0x38, 0x9c, 0xa5,
	0x30, 0xe5, 0x8b, 0x00, 0xb5, 0x01, 0x9d, 0xc6, 0xed, 0x39, 0x01, 0x6e, 0x81, 0x18, 0x58, 0x2e,
	0x8d, 0xea, 0x68, 0xbb, 0xb4, 0x53, 0x6d, 0x6d, 0xaa, 0x3f, 0xbe, 0x5f, 0x35, 0x2c, 0x97, 0x92,
	0x14, 0x8a, 0x31, 0x08, 0x89, 0xf8, 0xfa, 0xca, 0x36, 0xda, 0x91, 0x08, 0x3f, 0x37, 0x4e, 0x40,
	0xde, 0xa7, 0x31, 0xb5, 0x63, 0xea, 0x74, 0x2d, 0xdf, 0xbd, 0xb0, 0x5c, 0x8a, 0xef, 0xc2, 0xfa,
	0x24, 0x3b, 0x9b, 0x36, 0x73, 0x68, 0x1d, 0xf1, 0x82, 0xb5, 0x3c, 0xd9, 0x61, 0x0e, 0xc5, 0x5b,
	0x00, 0x36, 0xf3, 0x4f, 0x3d, 0x87, 0xfa, 0x36, 0xe5, 0x94, 0x2b, 0x64, 0x21, 0xd3, 0xf8, 0x8c,
	0x60, 0x3d, 0x67, 0xd6, 0x42, 0x6a, 0x9d, 0x61, 0x02, 0x42, 0x3c, 0x0b, 0x52, 0xb6, 0x5a, 0xeb,
	0x79, 0x91, 0xe2, 0x65, 0xa3, 0xea, 0x12, 0x87, 0xca, 0xbf, 0x83, 0x59, 0x40, 0x09, 0xe7, 0xc2,
	0xb7, 0x40, 0xf2, 0x22, 0x33, 0x08, 0xe9, 0xa9, 0x37, 0xe5, 0x22, 0x2a, 0xa4, 0xe2, 0x45, 0x06,
	0x8f, 0x15, 0x1b, 0xa4, 0x39, 0x1e, 0x57, 0x61, 0x75, 0xd8, 0x3b, 0xea, 0xf5, 0x4f, 0x7a, 0xf2,
	0x3f, 0x58, 0x02, 0xf1, 0xd8, 0x68, 0x77, 0x74, 0x19, 0xe1, 0x1a, 0xc0, 0xf1, 0x90, 0xe8, 0x66,
	0x1a, 0xaf, 0x60, 0x0c, 0x35, 0xbd, 0xdf, 0x35, 0x17, 0x72, 0x25, 0x0c, 0x50, 0x3e, 0x78, 0x6d,
	0x1c, 0xe8, 0x3d, 0x59, 0x48, 0xf0, 0xdd, 0xc3, 0x9e, 0x6e, 0x6a, 0x44, 0x6f, 0x1f, 0xc9, 0x62,
	0xe3, 0x03, 0x82, 0xb5, 0x44, 0xb2, 0x11, 0xb2, 0x80, 0x86, 0xf1, 0x0c, 0x9f, 0x02, 0x76, 0x32,
	0xcd, 0x66, 0xde, 0xb1, 0x7c, 0x4c, 0x4f, 0x7f, 0xd3, 0x74, 0x3e, 0x12, 0xf2, 0xaf, 0xf3, 0x5d,
	0x26, 0xc2, 0x6f, 0xa0, 0x36, 0xbf, 0x67, 0x94, 0xd8, 0xe4, 0xfe, 0xab, 0xad, 0x47, 0x7f, 0xd2,
	0x58, 0xb2, 0xee, 0x2c, 0x86, 0xca, 0x7b, 0x04, 0x42, 0xf2, 0xd3, 0xc1, 0x7d, 0xa8, 0x04, 0x99,
	0x33, 0x3e, 0xb8, 0x6a, 0x6b, 0xef, 0x17, 0xf9, 0x17, 0x9b, 0x42, 0xe6, 0x24, 0xf8, 0x7f, 0x10,
	0xaf, 0x3c, 0x27, 0x1e, 0x73, 0xb5, 0x22, 0x49, 0x03, 0xbc, 0x01, 0xe5, 0x31, 0xf5, 0xdc, 0x71,
	0x5c, 0x2f, 0xf1, 0x74, 0x16, 0xe1, 0xc7, 0x50, 0x1e, 0x4d, 0x98, 0x7d, 0x16, 0xd5, 0x05, 0xde,
	0xc0, 0xdb, 0x45, 0x97, 0x6b, 0x09, 0x8a, 0x64, 0x60, 0xe5, 0x6d, 0x09, 0x44, 0x9e, 0xf9, 0xfb,
	0xfa, 0x5f, 0xc2, 0xda, 0x88, 0x5d, 0xf8, 0x8e, 0xe7, 0xbb, 0xe6, 0x88, 0x4d, 0xb3, 0xa6, 0xdf,
	0x2b, 0xd4, 0x95, 0x61, 0x0d, 0x36, 0x99, 0x91, 0x6a, 0x5e, 0xa9, 0xb1, 0x29, 0x6e, 0x03, 0x04,
	0x56, 0x68, 0xb9, 0xa1, 0x15, 0x8c, 0xa3, 0x7a, 0x89, 0xdb, 0xbb, 0x53, 0xbc, 0xc6, 0x19, 0x92,
	0x2c, 0x14, 0xe1, 0x17, 0x00, 0xdc, 0xb0, 0xc9, 0xf7, 0x4a, 0xe0, 0x7b, 0xf5, 0xe0, 0xa7, 0x1d,
	0x4a, 0xbf, 0x7c, 0x81, 0xa4, 0x51, 0x7e, 0x54, 0x08, 0x48, 0xf3, 0xfc, 0xf2, 0xa2, 0x54, 0x40,
	0x18, 0xe8, 0xaf, 0x06, 0x32, 0x4a, 0x56, 0x66, 0xd0, 0xd6, 0xba, 0xc9, 0x8a, 0x54, 0x61, 0xd5,
	0x38, 0xec, 0x0c, 0x86, 0x24, 0xd9, 0x0d, 0x09, 0x44, 0x32, 0xec, 0xea, 0x44, 0x16, 0x92, 0xbc,
	0xd6, 0x26, 0x9d, 0xfe, 0xbe, 0x2e, 0x8b, 0xca, 0x47, 0x04, 0xd2, 0x5c, 0xf5, 0x0d, 0x1e, 0x43,
	0x0b, 0xc4, 0x2b, 0x16, 0x3a, 0xf9, 0x04, 0x0a, 0xff, 0x48, 0x4f, 0x58, 0xe8, 0x90, 0x14, 0xaa,
	0x7c, 0x42, 0x20, 0x24, 0xf1, 0x0d, 0xb6, 0xf5, 0x0c, 0x56, 0xa3, 0xd9, 0xf9, 0x88, 0x4d, 0x72,
	0x63, 0x5b, 0x45, 0x1c, 0xc7, 0x1c, 0x46, 0x72, 0xb8, 0xf2, 0x0e, 0x41, 0x39, 0xcd, 0xdd, 0x60,
	0x7b, 0xf9, 0x53, 0x56, 0xba, 0x7e, 0xca, 0xb4, 0x27, 0xd0, 0xb0, 0xd9, 0x79, 0x01, 0x97, 0xf6,
	0xdf, 0xb2, 0x42, 0x23, 0x79, 0x58, 0x0d, 0xf4, 0x15, 0xa1, 0x51, 0x99, 0x3f, 0xb2, 0x7b, 0xdf,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x70, 0xbe, 0xbe, 0xf1, 0x07, 0x00, 0x00,
}