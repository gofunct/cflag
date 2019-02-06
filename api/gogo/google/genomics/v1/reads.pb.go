// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/genomics/v1/reads.proto

package google_genomics_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import _ "go.pedge.io/pb/gogo/google/longrunning"
import _ "github.com/gogo/protobuf/types"
import google_protobuf2 "go.pedge.io/pb/gogo/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ImportReadGroupSetsRequest_PartitionStrategy int32

const (
	ImportReadGroupSetsRequest_PARTITION_STRATEGY_UNSPECIFIED ImportReadGroupSetsRequest_PartitionStrategy = 0
	// In most cases, this strategy yields one read group set per file. This is
	// the default behavior.
	//
	// Allocate one read group set per file per sample. For BAM files, read
	// groups are considered to share a sample if they have identical sample
	// names. Furthermore, all reads for each file which do not belong to a read
	// group, if any, will be grouped into a single read group set per-file.
	ImportReadGroupSetsRequest_PER_FILE_PER_SAMPLE ImportReadGroupSetsRequest_PartitionStrategy = 1
	// Includes all read groups in all imported files into a single read group
	// set. Requires that the headers for all imported files are equivalent. All
	// reads which do not belong to a read group, if any, will be grouped into a
	// separate read group set.
	ImportReadGroupSetsRequest_MERGE_ALL ImportReadGroupSetsRequest_PartitionStrategy = 2
)

var ImportReadGroupSetsRequest_PartitionStrategy_name = map[int32]string{
	0: "PARTITION_STRATEGY_UNSPECIFIED",
	1: "PER_FILE_PER_SAMPLE",
	2: "MERGE_ALL",
}
var ImportReadGroupSetsRequest_PartitionStrategy_value = map[string]int32{
	"PARTITION_STRATEGY_UNSPECIFIED": 0,
	"PER_FILE_PER_SAMPLE":            1,
	"MERGE_ALL":                      2,
}

func (x ImportReadGroupSetsRequest_PartitionStrategy) String() string {
	return proto.EnumName(ImportReadGroupSetsRequest_PartitionStrategy_name, int32(x))
}
func (ImportReadGroupSetsRequest_PartitionStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorReads, []int{2, 0}
}

// The read group set search request.
type SearchReadGroupSetsRequest struct {
	// Restricts this query to read group sets within the given datasets. At least
	// one ID must be provided.
	DatasetIds []string `protobuf:"bytes,1,rep,name=dataset_ids,json=datasetIds" json:"dataset_ids,omitempty"`
	// Only return read group sets for which a substring of the name matches this
	// string.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `nextPageToken` from the previous response.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The maximum number of results to return in a single page. If unspecified,
	// defaults to 256. The maximum value is 1024.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (m *SearchReadGroupSetsRequest) Reset()                    { *m = SearchReadGroupSetsRequest{} }
func (m *SearchReadGroupSetsRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchReadGroupSetsRequest) ProtoMessage()               {}
func (*SearchReadGroupSetsRequest) Descriptor() ([]byte, []int) { return fileDescriptorReads, []int{0} }

func (m *SearchReadGroupSetsRequest) GetDatasetIds() []string {
	if m != nil {
		return m.DatasetIds
	}
	return nil
}

func (m *SearchReadGroupSetsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SearchReadGroupSetsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *SearchReadGroupSetsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// The read group set search response.
type SearchReadGroupSetsResponse struct {
	// The list of matching read group sets.
	ReadGroupSets []*ReadGroupSet `protobuf:"bytes,1,rep,name=read_group_sets,json=readGroupSets" json:"read_group_sets,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *SearchReadGroupSetsResponse) Reset()                    { *m = SearchReadGroupSetsResponse{} }
func (m *SearchReadGroupSetsResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchReadGroupSetsResponse) ProtoMessage()               {}
func (*SearchReadGroupSetsResponse) Descriptor() ([]byte, []int) { return fileDescriptorReads, []int{1} }

func (m *SearchReadGroupSetsResponse) GetReadGroupSets() []*ReadGroupSet {
	if m != nil {
		return m.ReadGroupSets
	}
	return nil
}

func (m *SearchReadGroupSetsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The read group set import request.
type ImportReadGroupSetsRequest struct {
	// Required. The ID of the dataset these read group sets will belong to. The
	// caller must have WRITE permissions to this dataset.
	DatasetId string `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// The reference set to which the imported read group sets are aligned to, if
	// any. The reference names of this reference set must be a superset of those
	// found in the imported file headers. If no reference set id is provided, a
	// best effort is made to associate with a matching reference set.
	ReferenceSetId string `protobuf:"bytes,4,opt,name=reference_set_id,json=referenceSetId,proto3" json:"reference_set_id,omitempty"`
	// A list of URIs pointing at [BAM
	// files](https://samtools.github.io/hts-specs/SAMv1.pdf)
	// in Google Cloud Storage.
	// Those URIs can include wildcards (*), but do not add or remove
	// matching files before import has completed.
	//
	// Note that Google Cloud Storage object listing is only eventually
	// consistent: files added may be not be immediately visible to
	// everyone. Thus, if using a wildcard it is preferable not to start
	// the import immediately after the files are created.
	SourceUris []string `protobuf:"bytes,2,rep,name=source_uris,json=sourceUris" json:"source_uris,omitempty"`
	// The partition strategy describes how read groups are partitioned into read
	// group sets.
	PartitionStrategy ImportReadGroupSetsRequest_PartitionStrategy `protobuf:"varint,5,opt,name=partition_strategy,json=partitionStrategy,proto3,enum=google.genomics.v1.ImportReadGroupSetsRequest_PartitionStrategy" json:"partition_strategy,omitempty"`
}

func (m *ImportReadGroupSetsRequest) Reset()                    { *m = ImportReadGroupSetsRequest{} }
func (m *ImportReadGroupSetsRequest) String() string            { return proto.CompactTextString(m) }
func (*ImportReadGroupSetsRequest) ProtoMessage()               {}
func (*ImportReadGroupSetsRequest) Descriptor() ([]byte, []int) { return fileDescriptorReads, []int{2} }

func (m *ImportReadGroupSetsRequest) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *ImportReadGroupSetsRequest) GetReferenceSetId() string {
	if m != nil {
		return m.ReferenceSetId
	}
	return ""
}

func (m *ImportReadGroupSetsRequest) GetSourceUris() []string {
	if m != nil {
		return m.SourceUris
	}
	return nil
}

func (m *ImportReadGroupSetsRequest) GetPartitionStrategy() ImportReadGroupSetsRequest_PartitionStrategy {
	if m != nil {
		return m.PartitionStrategy
	}
	return ImportReadGroupSetsRequest_PARTITION_STRATEGY_UNSPECIFIED
}

// The read group set import response.
type ImportReadGroupSetsResponse struct {
	// IDs of the read group sets that were created.
	ReadGroupSetIds []string `protobuf:"bytes,1,rep,name=read_group_set_ids,json=readGroupSetIds" json:"read_group_set_ids,omitempty"`
}

func (m *ImportReadGroupSetsResponse) Reset()                    { *m = ImportReadGroupSetsResponse{} }
func (m *ImportReadGroupSetsResponse) String() string            { return proto.CompactTextString(m) }
func (*ImportReadGroupSetsResponse) ProtoMessage()               {}
func (*ImportReadGroupSetsResponse) Descriptor() ([]byte, []int) { return fileDescriptorReads, []int{3} }

func (m *ImportReadGroupSetsResponse) GetReadGroupSetIds() []string {
	if m != nil {
		return m.ReadGroupSetIds
	}
	return nil
}

// The read group set export request.
type ExportReadGroupSetRequest struct {
	// Required. The Google Cloud project ID that owns this
	// export. The caller must have WRITE access to this project.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Required. A Google Cloud Storage URI for the exported BAM file.
	// The currently authenticated user must have write access to the new file.
	// An error will be returned if the URI already contains data.
	ExportUri string `protobuf:"bytes,2,opt,name=export_uri,json=exportUri,proto3" json:"export_uri,omitempty"`
	// Required. The ID of the read group set to export. The caller must have
	// READ access to this read group set.
	ReadGroupSetId string `protobuf:"bytes,3,opt,name=read_group_set_id,json=readGroupSetId,proto3" json:"read_group_set_id,omitempty"`
	// The reference names to export. If this is not specified, all reference
	// sequences, including unmapped reads, are exported.
	// Use `*` to export only unmapped reads.
	ReferenceNames []string `protobuf:"bytes,4,rep,name=reference_names,json=referenceNames" json:"reference_names,omitempty"`
}

func (m *ExportReadGroupSetRequest) Reset()                    { *m = ExportReadGroupSetRequest{} }
func (m *ExportReadGroupSetRequest) String() string            { return proto.CompactTextString(m) }
func (*ExportReadGroupSetRequest) ProtoMessage()               {}
func (*ExportReadGroupSetRequest) Descriptor() ([]byte, []int) { return fileDescriptorReads, []int{4} }

func (m *ExportReadGroupSetRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ExportReadGroupSetRequest) GetExportUri() string {
	if m != nil {
		return m.ExportUri
	}
	return ""
}

func (m *ExportReadGroupSetRequest) GetReadGroupSetId() string {
	if m != nil {
		return m.ReadGroupSetId
	}
	return ""
}

func (m *ExportReadGroupSetRequest) GetReferenceNames() []string {
	if m != nil {
		return m.ReferenceNames
	}
	return nil
}

type UpdateReadGroupSetRequest struct {
	// The ID of the read group set to be updated. The caller must have WRITE
	// permissions to the dataset associated with this read group set.
	ReadGroupSetId string `protobuf:"bytes,1,opt,name=read_group_set_id,json=readGroupSetId,proto3" json:"read_group_set_id,omitempty"`
	// The new read group set data. See `updateMask` for details on mutability of
	// fields.
	ReadGroupSet *ReadGroupSet `protobuf:"bytes,2,opt,name=read_group_set,json=readGroupSet" json:"read_group_set,omitempty"`
	// An optional mask specifying which fields to update. Supported fields:
	//
	// * [name][google.genomics.v1.ReadGroupSet.name].
	// * [referenceSetId][google.genomics.v1.ReadGroupSet.reference_set_id].
	//
	// Leaving `updateMask` unset is equivalent to specifying all mutable
	// fields.
	UpdateMask *google_protobuf2.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
}

func (m *UpdateReadGroupSetRequest) Reset()                    { *m = UpdateReadGroupSetRequest{} }
func (m *UpdateReadGroupSetRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateReadGroupSetRequest) ProtoMessage()               {}
func (*UpdateReadGroupSetRequest) Descriptor() ([]byte, []int) { return fileDescriptorReads, []int{5} }

func (m *UpdateReadGroupSetRequest) GetReadGroupSetId() string {
	if m != nil {
		return m.ReadGroupSetId
	}
	return ""
}

func (m *UpdateReadGroupSetRequest) GetReadGroupSet() *ReadGroupSet {
	if m != nil {
		return m.ReadGroupSet
	}
	return nil
}

func (m *UpdateReadGroupSetRequest) GetUpdateMask() *google_protobuf2.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type DeleteReadGroupSetRequest struct {
	// The ID of the read group set to be deleted. The caller must have WRITE
	// permissions to the dataset associated with this read group set.
	ReadGroupSetId string `protobuf:"bytes,1,opt,name=read_group_set_id,json=readGroupSetId,proto3" json:"read_group_set_id,omitempty"`
}

func (m *DeleteReadGroupSetRequest) Reset()                    { *m = DeleteReadGroupSetRequest{} }
func (m *DeleteReadGroupSetRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteReadGroupSetRequest) ProtoMessage()               {}
func (*DeleteReadGroupSetRequest) Descriptor() ([]byte, []int) { return fileDescriptorReads, []int{6} }

func (m *DeleteReadGroupSetRequest) GetReadGroupSetId() string {
	if m != nil {
		return m.ReadGroupSetId
	}
	return ""
}

type GetReadGroupSetRequest struct {
	// The ID of the read group set.
	ReadGroupSetId string `protobuf:"bytes,1,opt,name=read_group_set_id,json=readGroupSetId,proto3" json:"read_group_set_id,omitempty"`
}

func (m *GetReadGroupSetRequest) Reset()                    { *m = GetReadGroupSetRequest{} }
func (m *GetReadGroupSetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetReadGroupSetRequest) ProtoMessage()               {}
func (*GetReadGroupSetRequest) Descriptor() ([]byte, []int) { return fileDescriptorReads, []int{7} }

func (m *GetReadGroupSetRequest) GetReadGroupSetId() string {
	if m != nil {
		return m.ReadGroupSetId
	}
	return ""
}

type ListCoverageBucketsRequest struct {
	// Required. The ID of the read group set over which coverage is requested.
	ReadGroupSetId string `protobuf:"bytes,1,opt,name=read_group_set_id,json=readGroupSetId,proto3" json:"read_group_set_id,omitempty"`
	// The name of the reference to query, within the reference set associated
	// with this query. Optional.
	ReferenceName string `protobuf:"bytes,3,opt,name=reference_name,json=referenceName,proto3" json:"reference_name,omitempty"`
	// The start position of the range on the reference, 0-based inclusive. If
	// specified, `referenceName` must also be specified. Defaults to 0.
	Start int64 `protobuf:"varint,4,opt,name=start,proto3" json:"start,omitempty"`
	// The end position of the range on the reference, 0-based exclusive. If
	// specified, `referenceName` must also be specified. If unset or 0, defaults
	// to the length of the reference.
	End int64 `protobuf:"varint,5,opt,name=end,proto3" json:"end,omitempty"`
	// The desired width of each reported coverage bucket in base pairs. This
	// will be rounded down to the nearest precomputed bucket width; the value
	// of which is returned as `bucketWidth` in the response. Defaults
	// to infinity (each bucket spans an entire reference sequence) or the length
	// of the target range, if specified. The smallest precomputed
	// `bucketWidth` is currently 2048 base pairs; this is subject to
	// change.
	TargetBucketWidth int64 `protobuf:"varint,6,opt,name=target_bucket_width,json=targetBucketWidth,proto3" json:"target_bucket_width,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `nextPageToken` from the previous response.
	PageToken string `protobuf:"bytes,7,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The maximum number of results to return in a single page. If unspecified,
	// defaults to 1024. The maximum value is 2048.
	PageSize int32 `protobuf:"varint,8,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (m *ListCoverageBucketsRequest) Reset()                    { *m = ListCoverageBucketsRequest{} }
func (m *ListCoverageBucketsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListCoverageBucketsRequest) ProtoMessage()               {}
func (*ListCoverageBucketsRequest) Descriptor() ([]byte, []int) { return fileDescriptorReads, []int{8} }

func (m *ListCoverageBucketsRequest) GetReadGroupSetId() string {
	if m != nil {
		return m.ReadGroupSetId
	}
	return ""
}

func (m *ListCoverageBucketsRequest) GetReferenceName() string {
	if m != nil {
		return m.ReferenceName
	}
	return ""
}

func (m *ListCoverageBucketsRequest) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ListCoverageBucketsRequest) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *ListCoverageBucketsRequest) GetTargetBucketWidth() int64 {
	if m != nil {
		return m.TargetBucketWidth
	}
	return 0
}

func (m *ListCoverageBucketsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListCoverageBucketsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// A bucket over which read coverage has been precomputed. A bucket corresponds
// to a specific range of the reference sequence.
type CoverageBucket struct {
	// The genomic coordinate range spanned by this bucket.
	Range *Range `protobuf:"bytes,1,opt,name=range" json:"range,omitempty"`
	// The average number of reads which are aligned to each individual
	// reference base in this bucket.
	MeanCoverage float32 `protobuf:"fixed32,2,opt,name=mean_coverage,json=meanCoverage,proto3" json:"mean_coverage,omitempty"`
}

func (m *CoverageBucket) Reset()                    { *m = CoverageBucket{} }
func (m *CoverageBucket) String() string            { return proto.CompactTextString(m) }
func (*CoverageBucket) ProtoMessage()               {}
func (*CoverageBucket) Descriptor() ([]byte, []int) { return fileDescriptorReads, []int{9} }

func (m *CoverageBucket) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *CoverageBucket) GetMeanCoverage() float32 {
	if m != nil {
		return m.MeanCoverage
	}
	return 0
}

type ListCoverageBucketsResponse struct {
	// The length of each coverage bucket in base pairs. Note that buckets at the
	// end of a reference sequence may be shorter. This value is omitted if the
	// bucket width is infinity (the default behaviour, with no range or
	// `targetBucketWidth`).
	BucketWidth int64 `protobuf:"varint,1,opt,name=bucket_width,json=bucketWidth,proto3" json:"bucket_width,omitempty"`
	// The coverage buckets. The list of buckets is sparse; a bucket with 0
	// overlapping reads is not returned. A bucket never crosses more than one
	// reference sequence. Each bucket has width `bucketWidth`, unless
	// its end is the end of the reference sequence.
	CoverageBuckets []*CoverageBucket `protobuf:"bytes,2,rep,name=coverage_buckets,json=coverageBuckets" json:"coverage_buckets,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListCoverageBucketsResponse) Reset()         { *m = ListCoverageBucketsResponse{} }
func (m *ListCoverageBucketsResponse) String() string { return proto.CompactTextString(m) }
func (*ListCoverageBucketsResponse) ProtoMessage()    {}
func (*ListCoverageBucketsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorReads, []int{10}
}

func (m *ListCoverageBucketsResponse) GetBucketWidth() int64 {
	if m != nil {
		return m.BucketWidth
	}
	return 0
}

func (m *ListCoverageBucketsResponse) GetCoverageBuckets() []*CoverageBucket {
	if m != nil {
		return m.CoverageBuckets
	}
	return nil
}

func (m *ListCoverageBucketsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The read search request.
type SearchReadsRequest struct {
	// The IDs of the read groups sets within which to search for reads. All
	// specified read group sets must be aligned against a common set of reference
	// sequences; this defines the genomic coordinates for the query. Must specify
	// one of `readGroupSetIds` or `readGroupIds`.
	ReadGroupSetIds []string `protobuf:"bytes,1,rep,name=read_group_set_ids,json=readGroupSetIds" json:"read_group_set_ids,omitempty"`
	// The IDs of the read groups within which to search for reads. All specified
	// read groups must belong to the same read group sets. Must specify one of
	// `readGroupSetIds` or `readGroupIds`.
	ReadGroupIds []string `protobuf:"bytes,5,rep,name=read_group_ids,json=readGroupIds" json:"read_group_ids,omitempty"`
	// The reference sequence name, for example `chr1`, `1`, or `chrX`. If set to
	// `*`, only unmapped reads are returned. If unspecified, all reads (mapped
	// and unmapped) are returned.
	ReferenceName string `protobuf:"bytes,7,opt,name=reference_name,json=referenceName,proto3" json:"reference_name,omitempty"`
	// The start position of the range on the reference, 0-based inclusive. If
	// specified, `referenceName` must also be specified.
	Start int64 `protobuf:"varint,8,opt,name=start,proto3" json:"start,omitempty"`
	// The end position of the range on the reference, 0-based exclusive. If
	// specified, `referenceName` must also be specified.
	End int64 `protobuf:"varint,9,opt,name=end,proto3" json:"end,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `nextPageToken` from the previous response.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The maximum number of results to return in a single page. If unspecified,
	// defaults to 256. The maximum value is 2048.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (m *SearchReadsRequest) Reset()                    { *m = SearchReadsRequest{} }
func (m *SearchReadsRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchReadsRequest) ProtoMessage()               {}
func (*SearchReadsRequest) Descriptor() ([]byte, []int) { return fileDescriptorReads, []int{11} }

func (m *SearchReadsRequest) GetReadGroupSetIds() []string {
	if m != nil {
		return m.ReadGroupSetIds
	}
	return nil
}

func (m *SearchReadsRequest) GetReadGroupIds() []string {
	if m != nil {
		return m.ReadGroupIds
	}
	return nil
}

func (m *SearchReadsRequest) GetReferenceName() string {
	if m != nil {
		return m.ReferenceName
	}
	return ""
}

func (m *SearchReadsRequest) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *SearchReadsRequest) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *SearchReadsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *SearchReadsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// The read search response.
type SearchReadsResponse struct {
	// The list of matching alignments sorted by mapped genomic coordinate,
	// if any, ascending in position within the same reference. Unmapped reads,
	// which have no position, are returned contiguously and are sorted in
	// ascending lexicographic order by fragment name.
	Alignments []*Read `protobuf:"bytes,1,rep,name=alignments" json:"alignments,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *SearchReadsResponse) Reset()                    { *m = SearchReadsResponse{} }
func (m *SearchReadsResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchReadsResponse) ProtoMessage()               {}
func (*SearchReadsResponse) Descriptor() ([]byte, []int) { return fileDescriptorReads, []int{12} }

func (m *SearchReadsResponse) GetAlignments() []*Read {
	if m != nil {
		return m.Alignments
	}
	return nil
}

func (m *SearchReadsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The stream reads request.
type StreamReadsRequest struct {
	// The Google Cloud project ID which will be billed
	// for this access. The caller must have WRITE access to this project.
	// Required.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The ID of the read group set from which to stream reads.
	ReadGroupSetId string `protobuf:"bytes,2,opt,name=read_group_set_id,json=readGroupSetId,proto3" json:"read_group_set_id,omitempty"`
	// The reference sequence name, for example `chr1`,
	// `1`, or `chrX`. If set to *, only unmapped reads are
	// returned.
	ReferenceName string `protobuf:"bytes,3,opt,name=reference_name,json=referenceName,proto3" json:"reference_name,omitempty"`
	// The start position of the range on the reference, 0-based inclusive. If
	// specified, `referenceName` must also be specified.
	Start int64 `protobuf:"varint,4,opt,name=start,proto3" json:"start,omitempty"`
	// The end position of the range on the reference, 0-based exclusive. If
	// specified, `referenceName` must also be specified.
	End int64 `protobuf:"varint,5,opt,name=end,proto3" json:"end,omitempty"`
	// Restricts results to a shard containing approximately `1/totalShards`
	// of the normal response payload for this query. Results from a sharded
	// request are disjoint from those returned by all queries which differ only
	// in their shard parameter. A shard may yield 0 results; this is especially
	// likely for large values of `totalShards`.
	//
	// Valid values are `[0, totalShards)`.
	Shard int32 `protobuf:"varint,6,opt,name=shard,proto3" json:"shard,omitempty"`
	// Specifying `totalShards` causes a disjoint subset of the normal response
	// payload to be returned for each query with a unique `shard` parameter
	// specified. A best effort is made to yield equally sized shards. Sharding
	// can be used to distribute processing amongst workers, where each worker is
	// assigned a unique `shard` number and all workers specify the same
	// `totalShards` number. The union of reads returned for all sharded queries
	// `[0, totalShards)` is equal to those returned by a single unsharded query.
	//
	// Queries for different values of `totalShards` with common divisors will
	// share shard boundaries. For example, streaming `shard` 2 of 5
	// `totalShards` yields the same results as streaming `shard`s 4 and 5 of 10
	// `totalShards`. This property can be leveraged for adaptive retries.
	TotalShards int32 `protobuf:"varint,7,opt,name=total_shards,json=totalShards,proto3" json:"total_shards,omitempty"`
}

func (m *StreamReadsRequest) Reset()                    { *m = StreamReadsRequest{} }
func (m *StreamReadsRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamReadsRequest) ProtoMessage()               {}
func (*StreamReadsRequest) Descriptor() ([]byte, []int) { return fileDescriptorReads, []int{13} }

func (m *StreamReadsRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *StreamReadsRequest) GetReadGroupSetId() string {
	if m != nil {
		return m.ReadGroupSetId
	}
	return ""
}

func (m *StreamReadsRequest) GetReferenceName() string {
	if m != nil {
		return m.ReferenceName
	}
	return ""
}

func (m *StreamReadsRequest) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *StreamReadsRequest) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *StreamReadsRequest) GetShard() int32 {
	if m != nil {
		return m.Shard
	}
	return 0
}

func (m *StreamReadsRequest) GetTotalShards() int32 {
	if m != nil {
		return m.TotalShards
	}
	return 0
}

type StreamReadsResponse struct {
	Alignments []*Read `protobuf:"bytes,1,rep,name=alignments" json:"alignments,omitempty"`
}

func (m *StreamReadsResponse) Reset()                    { *m = StreamReadsResponse{} }
func (m *StreamReadsResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamReadsResponse) ProtoMessage()               {}
func (*StreamReadsResponse) Descriptor() ([]byte, []int) { return fileDescriptorReads, []int{14} }

func (m *StreamReadsResponse) GetAlignments() []*Read {
	if m != nil {
		return m.Alignments
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchReadGroupSetsRequest)(nil), "google.genomics.v1.SearchReadGroupSetsRequest")
	proto.RegisterType((*SearchReadGroupSetsResponse)(nil), "google.genomics.v1.SearchReadGroupSetsResponse")
	proto.RegisterType((*ImportReadGroupSetsRequest)(nil), "google.genomics.v1.ImportReadGroupSetsRequest")
	proto.RegisterType((*ImportReadGroupSetsResponse)(nil), "google.genomics.v1.ImportReadGroupSetsResponse")
	proto.RegisterType((*ExportReadGroupSetRequest)(nil), "google.genomics.v1.ExportReadGroupSetRequest")
	proto.RegisterType((*UpdateReadGroupSetRequest)(nil), "google.genomics.v1.UpdateReadGroupSetRequest")
	proto.RegisterType((*DeleteReadGroupSetRequest)(nil), "google.genomics.v1.DeleteReadGroupSetRequest")
	proto.RegisterType((*GetReadGroupSetRequest)(nil), "google.genomics.v1.GetReadGroupSetRequest")
	proto.RegisterType((*ListCoverageBucketsRequest)(nil), "google.genomics.v1.ListCoverageBucketsRequest")
	proto.RegisterType((*CoverageBucket)(nil), "google.genomics.v1.CoverageBucket")
	proto.RegisterType((*ListCoverageBucketsResponse)(nil), "google.genomics.v1.ListCoverageBucketsResponse")
	proto.RegisterType((*SearchReadsRequest)(nil), "google.genomics.v1.SearchReadsRequest")
	proto.RegisterType((*SearchReadsResponse)(nil), "google.genomics.v1.SearchReadsResponse")
	proto.RegisterType((*StreamReadsRequest)(nil), "google.genomics.v1.StreamReadsRequest")
	proto.RegisterType((*StreamReadsResponse)(nil), "google.genomics.v1.StreamReadsResponse")
	proto.RegisterEnum("google.genomics.v1.ImportReadGroupSetsRequest_PartitionStrategy", ImportReadGroupSetsRequest_PartitionStrategy_name, ImportReadGroupSetsRequest_PartitionStrategy_value)
}

func init() { proto.RegisterFile("google/genomics/v1/reads.proto", fileDescriptorReads) }

var fileDescriptorReads = []byte{
	// 1305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0x67, 0xec, 0xb8, 0x6d, 0x9e, 0x9b, 0xc4, 0x19, 0x97, 0xe2, 0xd8, 0xa4, 0x35, 0x5b, 0xda,
	0xa6, 0x81, 0xda, 0xc4, 0x08, 0x15, 0x05, 0x21, 0x91, 0xb6, 0x4e, 0x30, 0x4a, 0x5a, 0x6b, 0x9d,
	0x80, 0x38, 0xad, 0x26, 0xde, 0xc9, 0x76, 0x89, 0xbd, 0xbb, 0xec, 0x8c, 0xd3, 0x7f, 0xea, 0xa5,
	0x37, 0x90, 0x80, 0x03, 0xe2, 0xc4, 0x95, 0x2b, 0x47, 0xc4, 0x87, 0xe0, 0x84, 0xb8, 0xf0, 0x01,
	0x10, 0x1f, 0x80, 0x13, 0x47, 0x34, 0xb3, 0xbb, 0xf1, 0xae, 0x77, 0xb6, 0x71, 0x54, 0x89, 0x9b,
	0xe7, 0xcd, 0x6f, 0xdf, 0xfc, 0xde, 0xff, 0x67, 0xb8, 0x64, 0xb9, 0xae, 0x35, 0xa0, 0x4d, 0x8b,
	0x3a, 0xee, 0xd0, 0xee, 0xb3, 0xe6, 0xd1, 0x5a, 0xd3, 0xa7, 0xc4, 0x64, 0x0d, 0xcf, 0x77, 0xb9,
	0x8b, 0x71, 0x70, 0xdf, 0x88, 0xee, 0x1b, 0x47, 0x6b, 0xd5, 0xd7, 0xc3, 0x6f, 0x88, 0x67, 0x37,
	0x89, 0xe3, 0xb8, 0x9c, 0x70, 0xdb, 0x75, 0xc2, 0x2f, 0xaa, 0x4a, 0x8d, 0xc4, 0xb1, 0x68, 0x78,
	0x7f, 0x2d, 0xe3, 0x45, 0x32, 0xb0, 0x2d, 0x67, 0x48, 0x1d, 0x1e, 0xe2, 0xae, 0x66, 0xe0, 0x2c,
	0xdf, 0x1d, 0x79, 0x8c, 0x46, 0xb0, 0x2b, 0x21, 0x6c, 0xe0, 0x3a, 0x96, 0x3f, 0x72, 0x1c, 0xdb,
	0xb1, 0x9a, 0xae, 0x47, 0xfd, 0x04, 0xa7, 0x5a, 0x08, 0x92, 0xa7, 0xfd, 0xd1, 0x41, 0x93, 0x0e,
	0x3d, 0xfe, 0x38, 0xbc, 0xac, 0x4f, 0x5e, 0x1e, 0xd8, 0x74, 0x60, 0x1a, 0x43, 0xc2, 0x0e, 0x03,
	0x84, 0xf6, 0x0d, 0x82, 0x6a, 0x8f, 0x12, 0xbf, 0xff, 0x40, 0xa7, 0xc4, 0xdc, 0x12, 0x04, 0x7a,
	0x94, 0x33, 0x9d, 0x7e, 0x39, 0xa2, 0x8c, 0xe3, 0xcb, 0x50, 0x34, 0x09, 0x27, 0x8c, 0x72, 0xc3,
	0x36, 0x59, 0x05, 0xd5, 0xf3, 0x2b, 0xb3, 0x3a, 0x84, 0xa2, 0x8e, 0xc9, 0x30, 0x86, 0x19, 0x87,
	0x0c, 0x69, 0x25, 0x5f, 0x47, 0x2b, 0xb3, 0xba, 0xfc, 0x8d, 0x97, 0x01, 0x3c, 0x62, 0x51, 0x83,
	0xbb, 0x87, 0xd4, 0xa9, 0xe4, 0xe4, 0xcd, 0xac, 0x90, 0xec, 0x0a, 0x01, 0xae, 0x81, 0x3c, 0x18,
	0xcc, 0x7e, 0x42, 0x2b, 0x33, 0x75, 0xb4, 0x52, 0xd0, 0xcf, 0x09, 0x41, 0xcf, 0x7e, 0x42, 0xb5,
	0xef, 0x10, 0xd4, 0x94, 0x7c, 0x98, 0xe7, 0x3a, 0x8c, 0xe2, 0x8f, 0x61, 0x41, 0x78, 0xca, 0x90,
	0xae, 0x32, 0x18, 0xe5, 0x01, 0xa9, 0x62, 0xab, 0xde, 0x48, 0x87, 0xb3, 0x11, 0xd7, 0xa1, 0xcf,
	0xf9, 0x71, 0x8d, 0xf8, 0x1a, 0x2c, 0x38, 0xf4, 0x11, 0x37, 0x52, 0x54, 0xe7, 0x84, 0xb8, 0x1b,
	0xd1, 0xd5, 0xfe, 0xcc, 0x41, 0xb5, 0x33, 0xf4, 0x5c, 0x9f, 0x2b, 0x3d, 0xb4, 0x0c, 0x30, 0xf6,
	0x50, 0x05, 0x05, 0xc6, 0x1e, 0x3b, 0x08, 0xaf, 0x40, 0xc9, 0xa7, 0x07, 0xd4, 0xa7, 0x4e, 0x9f,
	0x1a, 0x21, 0x68, 0x46, 0x82, 0xe6, 0x8f, 0xe5, 0x3d, 0x89, 0xbc, 0x0c, 0x45, 0xe6, 0x8e, 0xfc,
	0x3e, 0x35, 0x46, 0xbe, 0xcd, 0x2a, 0xb9, 0xc0, 0xd5, 0x81, 0x68, 0xcf, 0xb7, 0x19, 0x76, 0x01,
	0x7b, 0xc4, 0xe7, 0xb6, 0x88, 0xbe, 0xc1, 0xb8, 0x4f, 0x38, 0xb5, 0x1e, 0x57, 0x0a, 0x75, 0xb4,
	0x32, 0xdf, 0xfa, 0x48, 0x65, 0x7d, 0x36, 0xeb, 0x46, 0x37, 0x52, 0xd4, 0x0b, 0xf5, 0xe8, 0x8b,
	0xde, 0xa4, 0x48, 0x33, 0x60, 0x31, 0x85, 0xc3, 0x1a, 0x5c, 0xea, 0x6e, 0xe8, 0xbb, 0x9d, 0xdd,
	0xce, 0xfd, 0x7b, 0x46, 0x6f, 0x57, 0xdf, 0xd8, 0x6d, 0x6f, 0x7d, 0x6e, 0xec, 0xdd, 0xeb, 0x75,
	0xdb, 0x77, 0x3a, 0x9b, 0x9d, 0xf6, 0xdd, 0xd2, 0x2b, 0xf8, 0x35, 0x28, 0x77, 0xdb, 0xba, 0xb1,
	0xd9, 0xd9, 0x6e, 0x1b, 0xe2, 0x47, 0x6f, 0x63, 0xa7, 0xbb, 0xdd, 0x2e, 0x21, 0x3c, 0x07, 0xb3,
	0x3b, 0x6d, 0x7d, 0xab, 0x6d, 0x6c, 0x6c, 0x6f, 0x97, 0x72, 0xda, 0x27, 0x50, 0x53, 0x72, 0x0c,
	0x63, 0xfd, 0x16, 0xe0, 0x64, 0xac, 0x63, 0x39, 0xb8, 0x10, 0x0f, 0x66, 0xc7, 0x64, 0xda, 0xcf,
	0x08, 0x96, 0xda, 0x8f, 0x26, 0x95, 0xc5, 0xa2, 0xe4, 0xf9, 0xee, 0x17, 0xb4, 0x1f, 0x8f, 0x52,
	0x28, 0xe9, 0x98, 0xe2, 0x9a, 0xca, 0x6f, 0x85, 0xef, 0xa3, 0x8c, 0x0d, 0x24, 0x7b, 0xbe, 0x8d,
	0x6f, 0xc0, 0x62, 0x8a, 0x48, 0x98, 0xf1, 0xf3, 0x49, 0x1e, 0xf8, 0xba, 0xc8, 0xcf, 0x28, 0xde,
	0xa2, 0x1a, 0x58, 0x65, 0x46, 0x12, 0x1e, 0x87, 0xfb, 0x9e, 0x90, 0x6a, 0xbf, 0x21, 0x58, 0xda,
	0xf3, 0x4c, 0xc2, 0xa9, 0x8a, 0xaf, 0xf2, 0x45, 0xa4, 0x7c, 0x71, 0x13, 0xe6, 0x93, 0x50, 0xc9,
	0x7f, 0x9a, 0x82, 0x38, 0x1f, 0xd7, 0x84, 0x3f, 0x80, 0xe2, 0x48, 0xf2, 0x91, 0xed, 0x41, 0x9a,
	0x57, 0x6c, 0x55, 0x23, 0x25, 0x51, 0x07, 0x69, 0x6c, 0x8a, 0x0e, 0xb2, 0x43, 0xd8, 0xa1, 0x0e,
	0x01, 0x5c, 0xfc, 0xd6, 0x36, 0x61, 0xe9, 0x2e, 0x1d, 0xd0, 0x97, 0x35, 0x46, 0xbb, 0x03, 0x17,
	0xb7, 0x28, 0x7f, 0x49, 0x25, 0xcf, 0x73, 0x50, 0xdd, 0xb6, 0x19, 0xbf, 0xe3, 0x1e, 0x51, 0x9f,
	0x58, 0xf4, 0xf6, 0xa8, 0x7f, 0x18, 0xab, 0xd8, 0x53, 0xf8, 0xf6, 0x2a, 0xcc, 0x27, 0xa3, 0x19,
	0x46, 0x7d, 0x2e, 0x11, 0x4c, 0x7c, 0x01, 0x0a, 0x8c, 0x13, 0x9f, 0xcb, 0xca, 0xce, 0xeb, 0xc1,
	0x01, 0x97, 0x20, 0x4f, 0x1d, 0x53, 0x16, 0x68, 0x5e, 0x17, 0x3f, 0x71, 0x03, 0xca, 0x9c, 0xf8,
	0x16, 0xe5, 0xc6, 0xbe, 0xa4, 0x64, 0x3c, 0xb4, 0x4d, 0xfe, 0xa0, 0x72, 0x46, 0x22, 0x16, 0x83,
	0xab, 0x80, 0xec, 0x67, 0xe2, 0x62, 0xa2, 0x91, 0x9e, 0x7d, 0x61, 0x23, 0x3d, 0x37, 0xd1, 0x48,
	0x0f, 0x60, 0x3e, 0x69, 0x3f, 0x6e, 0x42, 0x41, 0x0e, 0x2b, 0x69, 0x6b, 0xb1, 0xb5, 0xa4, 0xcc,
	0x0f, 0x01, 0xd0, 0x03, 0x1c, 0xbe, 0x02, 0x73, 0x43, 0x4a, 0x1c, 0xa3, 0x1f, 0xea, 0x91, 0x89,
	0x95, 0xd3, 0xcf, 0x0b, 0x61, 0xa4, 0x5b, 0xfb, 0x15, 0x41, 0x4d, 0xe9, 0xec, 0xb0, 0x88, 0xdf,
	0x80, 0xf3, 0x09, 0x63, 0x91, 0x34, 0xb6, 0xb8, 0x1f, 0x33, 0x73, 0x07, 0x4a, 0xd1, 0x13, 0xa1,
	0x63, 0x82, 0xf6, 0x57, 0x6c, 0x69, 0x2a, 0x8e, 0xc9, 0x97, 0xf4, 0x85, 0x7e, 0xf2, 0x65, 0x55,
	0x63, 0xcf, 0xab, 0x1a, 0xfb, 0x3f, 0x08, 0xf0, 0x78, 0xd4, 0x1c, 0xa7, 0xc7, 0x69, 0xba, 0x0e,
	0x7e, 0x33, 0x51, 0x7c, 0x02, 0x58, 0x90, 0xc0, 0x71, 0x69, 0x09, 0x54, 0x3a, 0x8d, 0xce, 0xbe,
	0x30, 0x8d, 0xce, 0x29, 0xd2, 0x68, 0x76, 0x9c, 0x46, 0xc9, 0xb4, 0xc8, 0x9f, 0x6a, 0xbe, 0x3e,
	0x84, 0x72, 0xc2, 0xe6, 0x30, 0x4a, 0xef, 0x03, 0x1c, 0x2f, 0x29, 0xd1, 0x44, 0xad, 0x64, 0x35,
	0x10, 0x3d, 0x86, 0x9d, 0x7a, 0x8c, 0xfe, 0x2d, 0xbc, 0xcd, 0x7d, 0x4a, 0x86, 0x09, 0x6f, 0x9f,
	0xd0, 0x98, 0x95, 0xb5, 0x9a, 0xfb, 0x3f, 0x6a, 0x55, 0xe0, 0x1e, 0x10, 0xdf, 0x94, 0xd5, 0x59,
	0xd0, 0x83, 0x83, 0xc8, 0x66, 0xee, 0x72, 0x32, 0x30, 0xe4, 0x91, 0xc9, 0x38, 0x16, 0xf4, 0xa2,
	0x94, 0xf5, 0xa4, 0x48, 0xbb, 0x0f, 0xe5, 0x84, 0x9d, 0x2f, 0xeb, 0xe1, 0xd6, 0x0f, 0x08, 0x2e,
	0x04, 0x1a, 0x6d, 0xc7, 0x12, 0xb7, 0x3d, 0xea, 0x1f, 0xd9, 0x7d, 0x8a, 0x9f, 0x41, 0x31, 0xf6,
	0x12, 0xbe, 0xa6, 0xd2, 0x96, 0x76, 0x79, 0xf5, 0xfa, 0x89, 0xb8, 0x80, 0xb2, 0x56, 0x7b, 0xfe,
	0xc7, 0x5f, 0xdf, 0xe7, 0x5e, 0xd5, 0x4a, 0xc7, 0x9b, 0xf3, 0x3a, 0x93, 0xb0, 0x75, 0xb4, 0xfa,
	0x0e, 0x6a, 0xfd, 0x3e, 0x0b, 0x73, 0x31, 0x3a, 0x9f, 0xae, 0xe1, 0xaf, 0x10, 0x94, 0x15, 0x03,
	0x1d, 0x37, 0x4e, 0xb7, 0x9d, 0x54, 0x97, 0x23, 0x7c, 0x6c, 0xf3, 0x6d, 0xdc, 0x8f, 0x36, 0x5f,
	0xed, 0x8a, 0xe4, 0xb5, 0xac, 0x55, 0x26, 0xf7, 0x66, 0xb6, 0x6e, 0x4b, 0xa5, 0xeb, 0x68, 0x15,
	0xff, 0x88, 0x00, 0xa7, 0xf7, 0x01, 0x7c, 0x53, 0x45, 0x25, 0x73, 0x6f, 0x38, 0x89, 0xc9, 0x2d,
	0xc9, 0x64, 0x4d, 0x7b, 0x3b, 0xc5, 0xa4, 0xf9, 0x34, 0x95, 0xb7, 0xcf, 0xd6, 0x83, 0x8d, 0x22,
	0x64, 0x57, 0x56, 0xac, 0xb9, 0x6a, 0x4f, 0x65, 0xef, 0xe7, 0xd5, 0xe6, 0xd4, 0xf8, 0x30, 0xa6,
	0xd9, 0xbe, 0x6b, 0x32, 0xf9, 0x99, 0x60, 0xf7, 0x13, 0x02, 0x9c, 0xde, 0x4d, 0xd4, 0xbe, 0xcb,
	0xdc, 0x61, 0xaa, 0x27, 0x2e, 0x20, 0xda, 0x87, 0x92, 0xcc, 0xad, 0xd6, 0xd5, 0xe9, 0xdc, 0x37,
	0xb1, 0xe7, 0xe0, 0xaf, 0x11, 0xe0, 0xf4, 0xd6, 0xa1, 0xa6, 0x99, 0xb9, 0x9d, 0x54, 0x2f, 0xa6,
	0x56, 0x9c, 0xb6, 0xf8, 0x07, 0xa5, 0xdd, 0x94, 0xe4, 0xae, 0xaf, 0x4e, 0x47, 0x0e, 0x7f, 0x8b,
	0x60, 0x61, 0x62, 0x75, 0xc1, 0xab, 0x2a, 0x26, 0xea, 0xfd, 0x66, 0x0a, 0x6f, 0x85, 0x84, 0xf0,
	0x94, 0x84, 0x7e, 0x41, 0x50, 0x56, 0x0c, 0x66, 0x75, 0x8a, 0x65, 0xaf, 0x4b, 0xea, 0x14, 0x7b,
	0xc1, 0xc4, 0x8f, 0xa2, 0x8a, 0xdf, 0x9b, 0x8a, 0x67, 0x33, 0x1a, 0xdf, 0xe1, 0xe4, 0xc7, 0x4f,
	0xa1, 0x18, 0x9b, 0x50, 0x19, 0x5d, 0x2d, 0x35, 0xb6, 0x33, 0xba, 0x5a, 0x7a, 0xd4, 0x29, 0xba,
	0xda, 0x38, 0xf3, 0x6f, 0xdf, 0x80, 0x8b, 0x7d, 0x77, 0xa8, 0x50, 0x75, 0x1b, 0xa4, 0x96, 0xae,
	0xc8, 0x91, 0x2e, 0xfa, 0x17, 0xa1, 0xfd, 0x33, 0x32, 0x5f, 0xde, 0xfd, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x92, 0xf8, 0xeb, 0x27, 0x67, 0x10, 0x00, 0x00,
}