// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/genomics/v1/references.proto

package google_genomics_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A reference is a canonical assembled DNA sequence, intended to act as a
// reference coordinate space for other genomic annotations. A single reference
// might represent the human chromosome 1 or mitochandrial DNA, for instance. A
// reference belongs to one or more reference sets.
//
// For more genomics resource definitions, see [Fundamentals of Google
// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
type Reference struct {
	// The server-generated reference ID, unique across all references.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The length of this reference's sequence.
	Length int64 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	// MD5 of the upper-case sequence excluding all whitespace characters (this
	// is equivalent to SQ:M5 in SAM). This value is represented in lower case
	// hexadecimal format.
	Md5Checksum string `protobuf:"bytes,3,opt,name=md5checksum" json:"md5checksum,omitempty"`
	// The name of this reference, for example `22`.
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	// The URI from which the sequence was obtained. Typically specifies a FASTA
	// format file.
	SourceUri string `protobuf:"bytes,5,opt,name=source_uri,json=sourceUri" json:"source_uri,omitempty"`
	// All known corresponding accession IDs in INSDC (GenBank/ENA/DDBJ) ideally
	// with a version number, for example `GCF_000001405.26`.
	SourceAccessions []string `protobuf:"bytes,6,rep,name=source_accessions,json=sourceAccessions" json:"source_accessions,omitempty"`
	// ID from http://www.ncbi.nlm.nih.gov/taxonomy. For example, 9606 for human.
	NcbiTaxonId int32 `protobuf:"varint,7,opt,name=ncbi_taxon_id,json=ncbiTaxonId" json:"ncbi_taxon_id,omitempty"`
}

func (m *Reference) Reset()                    { *m = Reference{} }
func (m *Reference) String() string            { return proto.CompactTextString(m) }
func (*Reference) ProtoMessage()               {}
func (*Reference) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *Reference) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Reference) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Reference) GetMd5Checksum() string {
	if m != nil {
		return m.Md5Checksum
	}
	return ""
}

func (m *Reference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Reference) GetSourceUri() string {
	if m != nil {
		return m.SourceUri
	}
	return ""
}

func (m *Reference) GetSourceAccessions() []string {
	if m != nil {
		return m.SourceAccessions
	}
	return nil
}

func (m *Reference) GetNcbiTaxonId() int32 {
	if m != nil {
		return m.NcbiTaxonId
	}
	return 0
}

// A reference set is a set of references which typically comprise a reference
// assembly for a species, such as `GRCh38` which is representative
// of the human genome. A reference set defines a common coordinate space for
// comparing reference-aligned experimental data. A reference set contains 1 or
// more references.
//
// For more genomics resource definitions, see [Fundamentals of Google
// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
type ReferenceSet struct {
	// The server-generated reference set ID, unique across all reference sets.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The IDs of the reference objects that are part of this set.
	// `Reference.md5checksum` must be unique within this set.
	ReferenceIds []string `protobuf:"bytes,2,rep,name=reference_ids,json=referenceIds" json:"reference_ids,omitempty"`
	// Order-independent MD5 checksum which identifies this reference set. The
	// checksum is computed by sorting all lower case hexidecimal string
	// `reference.md5checksum` (for all reference in this set) in
	// ascending lexicographic order, concatenating, and taking the MD5 of that
	// value. The resulting value is represented in lower case hexadecimal format.
	Md5Checksum string `protobuf:"bytes,3,opt,name=md5checksum" json:"md5checksum,omitempty"`
	// ID from http://www.ncbi.nlm.nih.gov/taxonomy (for example, 9606 for human)
	// indicating the species which this reference set is intended to model. Note
	// that contained references may specify a different `ncbiTaxonId`, as
	// assemblies may contain reference sequences which do not belong to the
	// modeled species, for example EBV in a human reference genome.
	NcbiTaxonId int32 `protobuf:"varint,4,opt,name=ncbi_taxon_id,json=ncbiTaxonId" json:"ncbi_taxon_id,omitempty"`
	// Free text description of this reference set.
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	// Public id of this reference set, such as `GRCh37`.
	AssemblyId string `protobuf:"bytes,6,opt,name=assembly_id,json=assemblyId" json:"assembly_id,omitempty"`
	// The URI from which the references were obtained.
	SourceUri string `protobuf:"bytes,7,opt,name=source_uri,json=sourceUri" json:"source_uri,omitempty"`
	// All known corresponding accession IDs in INSDC (GenBank/ENA/DDBJ) ideally
	// with a version number, for example `NC_000001.11`.
	SourceAccessions []string `protobuf:"bytes,8,rep,name=source_accessions,json=sourceAccessions" json:"source_accessions,omitempty"`
}

func (m *ReferenceSet) Reset()                    { *m = ReferenceSet{} }
func (m *ReferenceSet) String() string            { return proto.CompactTextString(m) }
func (*ReferenceSet) ProtoMessage()               {}
func (*ReferenceSet) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *ReferenceSet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReferenceSet) GetReferenceIds() []string {
	if m != nil {
		return m.ReferenceIds
	}
	return nil
}

func (m *ReferenceSet) GetMd5Checksum() string {
	if m != nil {
		return m.Md5Checksum
	}
	return ""
}

func (m *ReferenceSet) GetNcbiTaxonId() int32 {
	if m != nil {
		return m.NcbiTaxonId
	}
	return 0
}

func (m *ReferenceSet) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ReferenceSet) GetAssemblyId() string {
	if m != nil {
		return m.AssemblyId
	}
	return ""
}

func (m *ReferenceSet) GetSourceUri() string {
	if m != nil {
		return m.SourceUri
	}
	return ""
}

func (m *ReferenceSet) GetSourceAccessions() []string {
	if m != nil {
		return m.SourceAccessions
	}
	return nil
}

type SearchReferenceSetsRequest struct {
	// If present, return reference sets for which the
	// [md5checksum][google.genomics.v1.ReferenceSet.md5checksum] matches exactly.
	Md5Checksums []string `protobuf:"bytes,1,rep,name=md5checksums" json:"md5checksums,omitempty"`
	// If present, return reference sets for which a prefix of any of
	// [sourceAccessions][google.genomics.v1.ReferenceSet.source_accessions]
	// match any of these strings. Accession numbers typically have a main number
	// and a version, for example `NC_000001.11`.
	Accessions []string `protobuf:"bytes,2,rep,name=accessions" json:"accessions,omitempty"`
	// If present, return reference sets for which a substring of their
	// `assemblyId` matches this string (case insensitive).
	AssemblyId string `protobuf:"bytes,3,opt,name=assembly_id,json=assemblyId" json:"assembly_id,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `nextPageToken` from the previous response.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// The maximum number of results to return in a single page. If unspecified,
	// defaults to 1024. The maximum value is 4096.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *SearchReferenceSetsRequest) Reset()                    { *m = SearchReferenceSetsRequest{} }
func (m *SearchReferenceSetsRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchReferenceSetsRequest) ProtoMessage()               {}
func (*SearchReferenceSetsRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *SearchReferenceSetsRequest) GetMd5Checksums() []string {
	if m != nil {
		return m.Md5Checksums
	}
	return nil
}

func (m *SearchReferenceSetsRequest) GetAccessions() []string {
	if m != nil {
		return m.Accessions
	}
	return nil
}

func (m *SearchReferenceSetsRequest) GetAssemblyId() string {
	if m != nil {
		return m.AssemblyId
	}
	return ""
}

func (m *SearchReferenceSetsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *SearchReferenceSetsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type SearchReferenceSetsResponse struct {
	// The matching references sets.
	ReferenceSets []*ReferenceSet `protobuf:"bytes,1,rep,name=reference_sets,json=referenceSets" json:"reference_sets,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *SearchReferenceSetsResponse) Reset()                    { *m = SearchReferenceSetsResponse{} }
func (m *SearchReferenceSetsResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchReferenceSetsResponse) ProtoMessage()               {}
func (*SearchReferenceSetsResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *SearchReferenceSetsResponse) GetReferenceSets() []*ReferenceSet {
	if m != nil {
		return m.ReferenceSets
	}
	return nil
}

func (m *SearchReferenceSetsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetReferenceSetRequest struct {
	// The ID of the reference set.
	ReferenceSetId string `protobuf:"bytes,1,opt,name=reference_set_id,json=referenceSetId" json:"reference_set_id,omitempty"`
}

func (m *GetReferenceSetRequest) Reset()                    { *m = GetReferenceSetRequest{} }
func (m *GetReferenceSetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetReferenceSetRequest) ProtoMessage()               {}
func (*GetReferenceSetRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *GetReferenceSetRequest) GetReferenceSetId() string {
	if m != nil {
		return m.ReferenceSetId
	}
	return ""
}

type SearchReferencesRequest struct {
	// If present, return references for which the
	// [md5checksum][google.genomics.v1.Reference.md5checksum] matches exactly.
	Md5Checksums []string `protobuf:"bytes,1,rep,name=md5checksums" json:"md5checksums,omitempty"`
	// If present, return references for which a prefix of any of
	// [sourceAccessions][google.genomics.v1.Reference.source_accessions] match
	// any of these strings. Accession numbers typically have a main number and a
	// version, for example `GCF_000001405.26`.
	Accessions []string `protobuf:"bytes,2,rep,name=accessions" json:"accessions,omitempty"`
	// If present, return only references which belong to this reference set.
	ReferenceSetId string `protobuf:"bytes,3,opt,name=reference_set_id,json=referenceSetId" json:"reference_set_id,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `nextPageToken` from the previous response.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// The maximum number of results to return in a single page. If unspecified,
	// defaults to 1024. The maximum value is 4096.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *SearchReferencesRequest) Reset()                    { *m = SearchReferencesRequest{} }
func (m *SearchReferencesRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchReferencesRequest) ProtoMessage()               {}
func (*SearchReferencesRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *SearchReferencesRequest) GetMd5Checksums() []string {
	if m != nil {
		return m.Md5Checksums
	}
	return nil
}

func (m *SearchReferencesRequest) GetAccessions() []string {
	if m != nil {
		return m.Accessions
	}
	return nil
}

func (m *SearchReferencesRequest) GetReferenceSetId() string {
	if m != nil {
		return m.ReferenceSetId
	}
	return ""
}

func (m *SearchReferencesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *SearchReferencesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type SearchReferencesResponse struct {
	// The matching references.
	References []*Reference `protobuf:"bytes,1,rep,name=references" json:"references,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *SearchReferencesResponse) Reset()                    { *m = SearchReferencesResponse{} }
func (m *SearchReferencesResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchReferencesResponse) ProtoMessage()               {}
func (*SearchReferencesResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *SearchReferencesResponse) GetReferences() []*Reference {
	if m != nil {
		return m.References
	}
	return nil
}

func (m *SearchReferencesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetReferenceRequest struct {
	// The ID of the reference.
	ReferenceId string `protobuf:"bytes,1,opt,name=reference_id,json=referenceId" json:"reference_id,omitempty"`
}

func (m *GetReferenceRequest) Reset()                    { *m = GetReferenceRequest{} }
func (m *GetReferenceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetReferenceRequest) ProtoMessage()               {}
func (*GetReferenceRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

func (m *GetReferenceRequest) GetReferenceId() string {
	if m != nil {
		return m.ReferenceId
	}
	return ""
}

type ListBasesRequest struct {
	// The ID of the reference.
	ReferenceId string `protobuf:"bytes,1,opt,name=reference_id,json=referenceId" json:"reference_id,omitempty"`
	// The start position (0-based) of this query. Defaults to 0.
	Start int64 `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	// The end position (0-based, exclusive) of this query. Defaults to the length
	// of this reference.
	End int64 `protobuf:"varint,3,opt,name=end" json:"end,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `nextPageToken` from the previous response.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// The maximum number of bases to return in a single page. If unspecified,
	// defaults to 200Kbp (kilo base pairs). The maximum value is 10Mbp (mega base
	// pairs).
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListBasesRequest) Reset()                    { *m = ListBasesRequest{} }
func (m *ListBasesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListBasesRequest) ProtoMessage()               {}
func (*ListBasesRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{8} }

func (m *ListBasesRequest) GetReferenceId() string {
	if m != nil {
		return m.ReferenceId
	}
	return ""
}

func (m *ListBasesRequest) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ListBasesRequest) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *ListBasesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListBasesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type ListBasesResponse struct {
	// The offset position (0-based) of the given `sequence` from the
	// start of this `Reference`. This value will differ for each page
	// in a paginated request.
	Offset int64 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	// A substring of the bases that make up this reference.
	Sequence string `protobuf:"bytes,2,opt,name=sequence" json:"sequence,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListBasesResponse) Reset()                    { *m = ListBasesResponse{} }
func (m *ListBasesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListBasesResponse) ProtoMessage()               {}
func (*ListBasesResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{9} }

func (m *ListBasesResponse) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListBasesResponse) GetSequence() string {
	if m != nil {
		return m.Sequence
	}
	return ""
}

func (m *ListBasesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*Reference)(nil), "google.genomics.v1.Reference")
	proto.RegisterType((*ReferenceSet)(nil), "google.genomics.v1.ReferenceSet")
	proto.RegisterType((*SearchReferenceSetsRequest)(nil), "google.genomics.v1.SearchReferenceSetsRequest")
	proto.RegisterType((*SearchReferenceSetsResponse)(nil), "google.genomics.v1.SearchReferenceSetsResponse")
	proto.RegisterType((*GetReferenceSetRequest)(nil), "google.genomics.v1.GetReferenceSetRequest")
	proto.RegisterType((*SearchReferencesRequest)(nil), "google.genomics.v1.SearchReferencesRequest")
	proto.RegisterType((*SearchReferencesResponse)(nil), "google.genomics.v1.SearchReferencesResponse")
	proto.RegisterType((*GetReferenceRequest)(nil), "google.genomics.v1.GetReferenceRequest")
	proto.RegisterType((*ListBasesRequest)(nil), "google.genomics.v1.ListBasesRequest")
	proto.RegisterType((*ListBasesResponse)(nil), "google.genomics.v1.ListBasesResponse")
}

func init() { proto.RegisterFile("google/genomics/v1/references.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 820 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x41, 0x4f, 0xeb, 0x46,
	0x10, 0xd6, 0xc6, 0x24, 0x90, 0x49, 0x80, 0x30, 0xb4, 0xa9, 0x15, 0x9a, 0x92, 0x3a, 0x40, 0x23,
	0x40, 0x89, 0xa0, 0xaa, 0x54, 0x55, 0xea, 0xa1, 0xb9, 0xa0, 0x48, 0x3d, 0x20, 0x43, 0x7b, 0xb5,
	0x1c, 0x7b, 0x09, 0x16, 0xc4, 0x4e, 0xbd, 0x1b, 0x44, 0x41, 0x1c, 0x5a, 0xa9, 0xc7, 0xb6, 0x87,
	0x5e, 0x5a, 0xf5, 0xb7, 0xf4, 0xd4, 0x9f, 0xd0, 0x53, 0xf5, 0xae, 0xef, 0x47, 0xbc, 0xe3, 0x93,
	0xd7, 0x1b, 0x67, 0xe3, 0xf8, 0x91, 0x48, 0xbc, 0x5b, 0xf6, 0xdb, 0x9d, 0x99, 0xef, 0xfb, 0x66,
	0x67, 0x63, 0x68, 0x0e, 0x82, 0x60, 0x70, 0x4b, 0x3b, 0x03, 0xea, 0x07, 0x43, 0xcf, 0x61, 0x9d,
	0xbb, 0x93, 0x4e, 0x48, 0xaf, 0x68, 0x48, 0x7d, 0x87, 0xb2, 0xf6, 0x28, 0x0c, 0x78, 0x80, 0x18,
	0x1f, 0x6a, 0x4f, 0x0e, 0xb5, 0xef, 0x4e, 0x6a, 0x1f, 0xcb, 0x40, 0x7b, 0xe4, 0x75, 0x6c, 0xdf,
	0x0f, 0xb8, 0xcd, 0xbd, 0xc0, 0x97, 0x11, 0xc6, 0xff, 0x04, 0x8a, 0xe6, 0x24, 0x0d, 0x6e, 0x40,
	0xce, 0x73, 0x75, 0xd2, 0x20, 0xad, 0xa2, 0x99, 0xf3, 0x5c, 0xac, 0x42, 0xe1, 0x96, 0xfa, 0x03,
	0x7e, 0xad, 0xe7, 0x1a, 0xa4, 0xa5, 0x99, 0x72, 0x85, 0x0d, 0x28, 0x0d, 0xdd, 0x2f, 0x9c, 0x6b,
	0xea, 0xdc, 0xb0, 0xf1, 0x50, 0xd7, 0x44, 0x80, 0x0a, 0x21, 0xc2, 0x8a, 0x6f, 0x0f, 0xa9, 0xbe,
	0x22, 0xb6, 0xc4, 0x6f, 0xac, 0x03, 0xb0, 0x60, 0x1c, 0x3a, 0xd4, 0x1a, 0x87, 0x9e, 0x9e, 0x17,
	0x3b, 0xc5, 0x18, 0xf9, 0x2e, 0xf4, 0xf0, 0x08, 0xb6, 0xe4, 0xb6, 0xed, 0x38, 0x94, 0xb1, 0x88,
	0xa5, 0x5e, 0x68, 0x68, 0xad, 0xa2, 0x59, 0x89, 0x37, 0xbe, 0x49, 0x70, 0x34, 0x60, 0xdd, 0x77,
	0xfa, 0x9e, 0xc5, 0xed, 0xfb, 0xc0, 0xb7, 0x3c, 0x57, 0x5f, 0x6d, 0x90, 0x56, 0xde, 0x2c, 0x45,
	0xe0, 0x65, 0x84, 0xf5, 0x5c, 0xe3, 0xcf, 0x1c, 0x94, 0x13, 0x6d, 0x17, 0x94, 0xcf, 0xc9, 0x6b,
	0xc2, 0x7a, 0x62, 0xa1, 0xe5, 0xb9, 0x4c, 0xcf, 0x89, 0x6a, 0xe5, 0x04, 0xec, 0xb9, 0x6c, 0x09,
	0xad, 0x73, 0x5c, 0x56, 0xe6, 0xb8, 0x44, 0x59, 0x5c, 0xca, 0x9c, 0xd0, 0x1b, 0x45, 0xee, 0x4b,
	0xf1, 0x2a, 0x84, 0xbb, 0x50, 0xb2, 0x19, 0xa3, 0xc3, 0xfe, 0xed, 0x8f, 0x51, 0x8e, 0x82, 0x38,
	0x01, 0x13, 0xa8, 0xe7, 0xa6, 0xec, 0x5b, 0x5d, 0xca, 0xbe, 0xb5, 0x6c, 0xfb, 0x8c, 0x7f, 0x08,
	0xd4, 0x2e, 0xa8, 0x1d, 0x3a, 0xd7, 0xaa, 0x41, 0xcc, 0xa4, 0x3f, 0x8c, 0x29, 0xe3, 0x68, 0x40,
	0x59, 0x11, 0xc8, 0x74, 0x12, 0xfb, 0xa2, 0x62, 0xf8, 0x09, 0x80, 0x52, 0x28, 0x76, 0x4e, 0x41,
	0xd2, 0x7a, 0xb4, 0x2c, 0x3d, 0x23, 0x7b, 0x40, 0x2d, 0x1e, 0xdc, 0x50, 0x5f, 0x5e, 0x94, 0x62,
	0x84, 0x5c, 0x46, 0x00, 0xee, 0x80, 0x58, 0x58, 0xcc, 0x7b, 0xa0, 0xc2, 0xaf, 0xbc, 0xb9, 0x16,
	0x01, 0x17, 0xde, 0x03, 0x35, 0x7e, 0x23, 0xb0, 0x93, 0xc9, 0x9f, 0x8d, 0x02, 0x9f, 0x51, 0x3c,
	0x83, 0x8d, 0x69, 0x67, 0x19, 0xe5, 0xb1, 0x84, 0xd2, 0x69, 0xa3, 0x3d, 0x3f, 0x21, 0x6d, 0x35,
	0x85, 0x39, 0xbd, 0x11, 0x51, 0x42, 0x3c, 0x80, 0x4d, 0x9f, 0xde, 0x73, 0x4b, 0x61, 0x9a, 0x13,
	0x4c, 0xd7, 0x23, 0xf8, 0x7c, 0xc2, 0xd6, 0xe8, 0x42, 0xf5, 0x8c, 0xf2, 0x99, 0x4c, 0xd2, 0xcb,
	0x16, 0x54, 0x66, 0xa8, 0x58, 0xc9, 0x15, 0xdc, 0x50, 0x4b, 0xf5, 0x5c, 0xe3, 0x5f, 0x02, 0x1f,
	0xa5, 0x44, 0xbd, 0xd7, 0x8e, 0x64, 0x31, 0xd1, 0xb2, 0x98, 0xbc, 0xa8, 0x35, 0x3f, 0x11, 0xd0,
	0xe7, 0x55, 0xc8, 0xbe, 0x7c, 0x0d, 0x30, 0x7d, 0xb4, 0x64, 0x4f, 0xea, 0xcf, 0xf6, 0xc4, 0x54,
	0x02, 0x96, 0xee, 0xc6, 0x97, 0xb0, 0xad, 0x76, 0x63, 0x62, 0xe2, 0xa7, 0x50, 0x56, 0xe7, 0x5d,
	0xb6, 0xa1, 0xa4, 0x8c, 0xbb, 0xf1, 0x17, 0x81, 0xca, 0xb7, 0x1e, 0xe3, 0x5d, 0x9b, 0x4d, 0xcd,
	0x5f, 0x1c, 0x87, 0x1f, 0x40, 0x9e, 0x71, 0x3b, 0xe4, 0xf2, 0xa1, 0x8c, 0x17, 0x58, 0x01, 0x8d,
	0xfa, 0xb1, 0xc9, 0x9a, 0x19, 0xfd, 0x7c, 0x91, 0xb3, 0x01, 0x6c, 0x29, 0xd4, 0xa4, 0xa3, 0x55,
	0x28, 0x04, 0x57, 0x57, 0x8c, 0x72, 0xc1, 0x4a, 0x33, 0xe5, 0x0a, 0x6b, 0xb0, 0xc6, 0x22, 0xfa,
	0xbe, 0x43, 0xa5, 0x47, 0xc9, 0x3a, 0xcb, 0x46, 0x2d, 0xc3, 0xc6, 0xd3, 0x57, 0x79, 0x40, 0xe5,
	0x4a, 0x87, 0x77, 0x9e, 0x43, 0xbf, 0x3f, 0xc1, 0xbf, 0x09, 0x6c, 0x67, 0x0c, 0x1f, 0xb6, 0xb3,
	0x1a, 0xf9, 0xee, 0x57, 0xa6, 0xd6, 0x59, 0xfa, 0x7c, 0xac, 0xd5, 0x68, 0xfe, 0xfc, 0xdf, 0xeb,
	0x3f, 0x72, 0x75, 0x43, 0x9f, 0xfd, 0xf3, 0xa3, 0x9c, 0x75, 0x98, 0x08, 0xfb, 0x8a, 0x1c, 0xe2,
	0xaf, 0x04, 0x36, 0x53, 0xa3, 0x88, 0x87, 0x59, 0x95, 0xb2, 0xe7, 0xb5, 0xb6, 0xf0, 0x89, 0x30,
	0x8e, 0x05, 0x8d, 0x03, 0xdc, 0x9b, 0xa7, 0xf1, 0x98, 0x1e, 0xb0, 0x27, 0xfc, 0x9d, 0x40, 0x25,
	0x3d, 0x0f, 0x78, 0xb4, 0x84, 0xf4, 0xc4, 0xa7, 0xe3, 0xe5, 0x0e, 0x4b, 0x93, 0x1a, 0x82, 0x5d,
	0xcd, 0xf8, 0x70, 0x96, 0x9d, 0xe2, 0xd0, 0x13, 0x94, 0x55, 0xed, 0xf8, 0xd9, 0x22, 0x77, 0x26,
	0x44, 0x9e, 0x9f, 0x54, 0x63, 0x5f, 0x54, 0xde, 0xc5, 0x7a, 0xaa, 0xf2, 0xa3, 0x3a, 0x3c, 0x4f,
	0xf8, 0x0b, 0x81, 0x62, 0x72, 0x8f, 0x71, 0x2f, 0x2b, 0x67, 0x7a, 0x02, 0x6b, 0xfb, 0x0b, 0x4e,
	0x49, 0xed, 0x47, 0x82, 0xc1, 0x3e, 0x36, 0x9f, 0x65, 0xd0, 0xe9, 0x47, 0x41, 0xdd, 0x36, 0x54,
	0x9d, 0x60, 0x98, 0x91, 0xb8, 0xbb, 0x39, 0xb5, 0xf5, 0x3c, 0xfa, 0x4a, 0x3a, 0x27, 0x6f, 0x08,
	0xe9, 0x17, 0xc4, 0x17, 0xd3, 0xe7, 0x6f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x36, 0x1f, 0xf6, 0xb9,
	0x8a, 0x09, 0x00, 0x00,
}