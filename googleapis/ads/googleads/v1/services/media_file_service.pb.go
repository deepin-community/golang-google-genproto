// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/media_file_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [MediaFileService.GetMediaFile][google.ads.googleads.v1.services.MediaFileService.GetMediaFile]
type GetMediaFileRequest struct {
	// Required. The resource name of the media file to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMediaFileRequest) Reset()         { *m = GetMediaFileRequest{} }
func (m *GetMediaFileRequest) String() string { return proto.CompactTextString(m) }
func (*GetMediaFileRequest) ProtoMessage()    {}
func (*GetMediaFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_793e9e1492d555b0, []int{0}
}

func (m *GetMediaFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMediaFileRequest.Unmarshal(m, b)
}
func (m *GetMediaFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMediaFileRequest.Marshal(b, m, deterministic)
}
func (m *GetMediaFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMediaFileRequest.Merge(m, src)
}
func (m *GetMediaFileRequest) XXX_Size() int {
	return xxx_messageInfo_GetMediaFileRequest.Size(m)
}
func (m *GetMediaFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMediaFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMediaFileRequest proto.InternalMessageInfo

func (m *GetMediaFileRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [MediaFileService.MutateMediaFiles][google.ads.googleads.v1.services.MediaFileService.MutateMediaFiles]
type MutateMediaFilesRequest struct {
	// Required. The ID of the customer whose media files are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual media file.
	Operations []*MediaFileOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateMediaFilesRequest) Reset()         { *m = MutateMediaFilesRequest{} }
func (m *MutateMediaFilesRequest) String() string { return proto.CompactTextString(m) }
func (*MutateMediaFilesRequest) ProtoMessage()    {}
func (*MutateMediaFilesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_793e9e1492d555b0, []int{1}
}

func (m *MutateMediaFilesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMediaFilesRequest.Unmarshal(m, b)
}
func (m *MutateMediaFilesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMediaFilesRequest.Marshal(b, m, deterministic)
}
func (m *MutateMediaFilesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMediaFilesRequest.Merge(m, src)
}
func (m *MutateMediaFilesRequest) XXX_Size() int {
	return xxx_messageInfo_MutateMediaFilesRequest.Size(m)
}
func (m *MutateMediaFilesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMediaFilesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMediaFilesRequest proto.InternalMessageInfo

func (m *MutateMediaFilesRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateMediaFilesRequest) GetOperations() []*MediaFileOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateMediaFilesRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateMediaFilesRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation to create media file.
type MediaFileOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*MediaFileOperation_Create
	Operation            isMediaFileOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MediaFileOperation) Reset()         { *m = MediaFileOperation{} }
func (m *MediaFileOperation) String() string { return proto.CompactTextString(m) }
func (*MediaFileOperation) ProtoMessage()    {}
func (*MediaFileOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_793e9e1492d555b0, []int{2}
}

func (m *MediaFileOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaFileOperation.Unmarshal(m, b)
}
func (m *MediaFileOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaFileOperation.Marshal(b, m, deterministic)
}
func (m *MediaFileOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaFileOperation.Merge(m, src)
}
func (m *MediaFileOperation) XXX_Size() int {
	return xxx_messageInfo_MediaFileOperation.Size(m)
}
func (m *MediaFileOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaFileOperation.DiscardUnknown(m)
}

var xxx_messageInfo_MediaFileOperation proto.InternalMessageInfo

type isMediaFileOperation_Operation interface {
	isMediaFileOperation_Operation()
}

type MediaFileOperation_Create struct {
	Create *resources.MediaFile `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

func (*MediaFileOperation_Create) isMediaFileOperation_Operation() {}

func (m *MediaFileOperation) GetOperation() isMediaFileOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *MediaFileOperation) GetCreate() *resources.MediaFile {
	if x, ok := m.GetOperation().(*MediaFileOperation_Create); ok {
		return x.Create
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MediaFileOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MediaFileOperation_Create)(nil),
	}
}

// Response message for a media file mutate.
type MutateMediaFilesResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateMediaFileResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MutateMediaFilesResponse) Reset()         { *m = MutateMediaFilesResponse{} }
func (m *MutateMediaFilesResponse) String() string { return proto.CompactTextString(m) }
func (*MutateMediaFilesResponse) ProtoMessage()    {}
func (*MutateMediaFilesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_793e9e1492d555b0, []int{3}
}

func (m *MutateMediaFilesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMediaFilesResponse.Unmarshal(m, b)
}
func (m *MutateMediaFilesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMediaFilesResponse.Marshal(b, m, deterministic)
}
func (m *MutateMediaFilesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMediaFilesResponse.Merge(m, src)
}
func (m *MutateMediaFilesResponse) XXX_Size() int {
	return xxx_messageInfo_MutateMediaFilesResponse.Size(m)
}
func (m *MutateMediaFilesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMediaFilesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMediaFilesResponse proto.InternalMessageInfo

func (m *MutateMediaFilesResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateMediaFilesResponse) GetResults() []*MutateMediaFileResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the media file mutate.
type MutateMediaFileResult struct {
	// The resource name returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateMediaFileResult) Reset()         { *m = MutateMediaFileResult{} }
func (m *MutateMediaFileResult) String() string { return proto.CompactTextString(m) }
func (*MutateMediaFileResult) ProtoMessage()    {}
func (*MutateMediaFileResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_793e9e1492d555b0, []int{4}
}

func (m *MutateMediaFileResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMediaFileResult.Unmarshal(m, b)
}
func (m *MutateMediaFileResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMediaFileResult.Marshal(b, m, deterministic)
}
func (m *MutateMediaFileResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMediaFileResult.Merge(m, src)
}
func (m *MutateMediaFileResult) XXX_Size() int {
	return xxx_messageInfo_MutateMediaFileResult.Size(m)
}
func (m *MutateMediaFileResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMediaFileResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMediaFileResult proto.InternalMessageInfo

func (m *MutateMediaFileResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMediaFileRequest)(nil), "google.ads.googleads.v1.services.GetMediaFileRequest")
	proto.RegisterType((*MutateMediaFilesRequest)(nil), "google.ads.googleads.v1.services.MutateMediaFilesRequest")
	proto.RegisterType((*MediaFileOperation)(nil), "google.ads.googleads.v1.services.MediaFileOperation")
	proto.RegisterType((*MutateMediaFilesResponse)(nil), "google.ads.googleads.v1.services.MutateMediaFilesResponse")
	proto.RegisterType((*MutateMediaFileResult)(nil), "google.ads.googleads.v1.services.MutateMediaFileResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/media_file_service.proto", fileDescriptor_793e9e1492d555b0)
}

var fileDescriptor_793e9e1492d555b0 = []byte{
	// 715 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xbe, 0x4e, 0xae, 0x7a, 0x6f, 0x27, 0x2d, 0x54, 0x53, 0x95, 0x86, 0x80, 0x44, 0x64, 0x22,
	0x51, 0x45, 0x95, 0xad, 0x84, 0x22, 0x54, 0x43, 0x17, 0x8e, 0x20, 0x2d, 0x8b, 0xd2, 0x92, 0x4a,
	0x95, 0x40, 0x91, 0xac, 0xa9, 0x3d, 0x09, 0x23, 0xd9, 0x1e, 0x33, 0x33, 0x8e, 0x54, 0x55, 0xdd,
	0xf0, 0x0a, 0xbc, 0x01, 0x4b, 0xf6, 0x2c, 0x78, 0x85, 0x8a, 0x1d, 0xbb, 0x2e, 0x50, 0x17, 0x2c,
	0x10, 0x8f, 0x80, 0x58, 0x20, 0xff, 0x8c, 0xed, 0xa4, 0x8d, 0x4a, 0xd9, 0x1d, 0x9d, 0xf9, 0xce,
	0xf7, 0x7d, 0xe7, 0xcc, 0xf1, 0x18, 0xac, 0x0f, 0x29, 0x1d, 0xba, 0x58, 0x47, 0x0e, 0xd7, 0x93,
	0x30, 0x8a, 0x46, 0x2d, 0x9d, 0x63, 0x36, 0x22, 0x36, 0xe6, 0xba, 0x87, 0x1d, 0x82, 0xac, 0x01,
	0x71, 0xb1, 0x95, 0xe6, 0xb4, 0x80, 0x51, 0x41, 0x61, 0x3d, 0xc1, 0x6b, 0xc8, 0xe1, 0x5a, 0x56,
	0xaa, 0x8d, 0x5a, 0x9a, 0x2c, 0xad, 0xb5, 0xa7, 0x91, 0x33, 0xcc, 0x69, 0xc8, 0xc6, 0xd9, 0x13,
	0xd6, 0xda, 0x6d, 0x59, 0x13, 0x10, 0x1d, 0xf9, 0x3e, 0x15, 0x48, 0x10, 0xea, 0xf3, 0xf4, 0x74,
	0xb9, 0x70, 0x6a, 0xbb, 0x04, 0xfb, 0x22, 0x3d, 0xb8, 0x53, 0x38, 0x18, 0x10, 0xec, 0x3a, 0xd6,
	0x01, 0x7e, 0x8d, 0x46, 0x84, 0xb2, 0x14, 0x70, 0xb3, 0x00, 0x90, 0xf2, 0x13, 0xa4, 0x2c, 0xb0,
	0x75, 0x2e, 0x90, 0x08, 0x53, 0x35, 0x75, 0x00, 0x16, 0x37, 0xb1, 0xd8, 0x8e, 0x2c, 0x76, 0x89,
	0x8b, 0x7b, 0xf8, 0x4d, 0x88, 0xb9, 0x80, 0x3b, 0x60, 0x5e, 0x32, 0x58, 0x3e, 0xf2, 0x70, 0x55,
	0xa9, 0x2b, 0x2b, 0xb3, 0x9d, 0xe6, 0x99, 0x59, 0xfa, 0x69, 0x36, 0x80, 0x9a, 0x0f, 0x23, 0x8d,
	0x02, 0xc2, 0x35, 0x9b, 0x7a, 0x7a, 0xce, 0x34, 0x27, 0x09, 0x9e, 0x23, 0x0f, 0xab, 0xdf, 0x15,
	0xb0, 0xbc, 0x1d, 0x0a, 0x24, 0x70, 0x86, 0xe0, 0x52, 0xac, 0x01, 0x2a, 0x76, 0xc8, 0x05, 0xf5,
	0x30, 0xb3, 0x88, 0x93, 0x4a, 0x95, 0xcf, 0xcc, 0x52, 0x0f, 0xc8, 0xfc, 0x33, 0x07, 0xbe, 0x04,
	0x80, 0x06, 0x98, 0x25, 0xb3, 0xaa, 0x96, 0xea, 0xe5, 0x95, 0x4a, 0x7b, 0x4d, 0xbb, 0xec, 0x82,
	0xb4, 0x4c, 0x6e, 0x47, 0x16, 0xa7, 0xd4, 0x39, 0x19, 0xbc, 0x07, 0xae, 0x07, 0x88, 0x09, 0x82,
	0x5c, 0x6b, 0x80, 0x88, 0x1b, 0x32, 0x5c, 0x2d, 0xd7, 0x95, 0x95, 0xff, 0x7b, 0xd7, 0xd2, 0x74,
	0x37, 0xc9, 0xc2, 0xbb, 0x60, 0x7e, 0x84, 0x5c, 0xe2, 0x20, 0x81, 0x2d, 0xea, 0xbb, 0x87, 0xd5,
	0x7f, 0x63, 0xd8, 0x9c, 0x4c, 0xee, 0xf8, 0xee, 0xa1, 0x4a, 0x00, 0x3c, 0x2f, 0x0a, 0xbb, 0x60,
	0xc6, 0x66, 0x18, 0x89, 0x64, 0x94, 0x95, 0xf6, 0xea, 0x54, 0xeb, 0xd9, 0xe6, 0xe4, 0xde, 0xb7,
	0xfe, 0xe9, 0xa5, 0xd5, 0x9d, 0x0a, 0x98, 0xcd, 0x9c, 0xab, 0x1f, 0x15, 0x50, 0x3d, 0x3f, 0x55,
	0x1e, 0x50, 0x9f, 0x63, 0xd8, 0x05, 0x4b, 0x13, 0x5d, 0x59, 0x98, 0x31, 0xca, 0xe2, 0xde, 0x2a,
	0x6d, 0x28, 0x0d, 0xb0, 0xc0, 0xd6, 0xf6, 0xe2, 0x9d, 0xe8, 0x2d, 0x8e, 0xf7, 0xfb, 0x34, 0x82,
	0xc3, 0x17, 0xe0, 0x3f, 0x86, 0x79, 0xe8, 0x0a, 0x39, 0xf5, 0x87, 0x7f, 0x30, 0xf5, 0x71, 0x53,
	0xbd, 0xb8, 0xbe, 0x27, 0x79, 0xd4, 0xc7, 0x60, 0xe9, 0x42, 0x44, 0x34, 0xe0, 0x0b, 0xf6, 0x6e,
	0x7c, 0x97, 0xda, 0x9f, 0xcb, 0x60, 0x21, 0x2b, 0xdc, 0x4b, 0x24, 0xe1, 0x27, 0x05, 0xcc, 0x15,
	0x37, 0x19, 0x3e, 0xb8, 0xdc, 0xe5, 0x05, 0x9b, 0x5f, 0xbb, 0xd2, 0xbd, 0xa8, 0x4f, 0x4e, 0xcd,
	0x71, 0xc3, 0x6f, 0xbf, 0x7c, 0x7b, 0x57, 0xd2, 0xe0, 0x6a, 0xf4, 0x04, 0x1c, 0x8d, 0x9d, 0x6c,
	0xc8, 0x5d, 0xe6, 0x7a, 0x33, 0x79, 0x13, 0xe2, 0xeb, 0xd2, 0x9b, 0xc7, 0xf0, 0xab, 0x02, 0x16,
	0x26, 0xaf, 0x11, 0xae, 0x5f, 0x79, 0xca, 0xf2, 0x83, 0xaa, 0x19, 0x7f, 0x53, 0x9a, 0x6c, 0x8d,
	0xba, 0x77, 0x6a, 0xde, 0x28, 0x7c, 0x8d, 0xab, 0xf9, 0x67, 0x12, 0xb7, 0xb6, 0xa6, 0xea, 0x51,
	0x6b, 0x79, 0x2f, 0x47, 0x05, 0xf0, 0x46, 0xf3, 0xb8, 0xd0, 0x99, 0xe1, 0xc5, 0x1a, 0x86, 0xd2,
	0xac, 0xdd, 0x3a, 0x31, 0xab, 0xd3, 0x9e, 0x8c, 0xce, 0x2f, 0x05, 0x34, 0x6c, 0xea, 0x5d, 0xea,
	0xb9, 0xb3, 0x34, 0x79, 0xe9, 0xbb, 0xd1, 0x13, 0xb6, 0xab, 0xbc, 0xda, 0x4a, 0x4b, 0x87, 0xd4,
	0x45, 0xfe, 0x50, 0xa3, 0x6c, 0xa8, 0x0f, 0xb1, 0x1f, 0x3f, 0x70, 0x7a, 0x2e, 0x36, 0xfd, 0x07,
	0xf0, 0x48, 0x06, 0xef, 0x4b, 0xe5, 0x4d, 0xd3, 0xfc, 0x50, 0xaa, 0x6f, 0x26, 0x84, 0xa6, 0xc3,
	0xb5, 0x24, 0x8c, 0xa2, 0xfd, 0x96, 0x96, 0x0a, 0xf3, 0x13, 0x09, 0xe9, 0x9b, 0x0e, 0xef, 0x67,
	0x90, 0xfe, 0x7e, 0xab, 0x2f, 0x21, 0x3f, 0x4a, 0x8d, 0x24, 0x6f, 0x18, 0xa6, 0xc3, 0x0d, 0x23,
	0x03, 0x19, 0xc6, 0x7e, 0xcb, 0x30, 0x24, 0xec, 0x60, 0x26, 0xf6, 0x79, 0xff, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x1a, 0xc5, 0x35, 0x2a, 0xa7, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MediaFileServiceClient is the client API for MediaFileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MediaFileServiceClient interface {
	// Returns the requested media file in full detail.
	GetMediaFile(ctx context.Context, in *GetMediaFileRequest, opts ...grpc.CallOption) (*resources.MediaFile, error)
	// Creates media files. Operation statuses are returned.
	MutateMediaFiles(ctx context.Context, in *MutateMediaFilesRequest, opts ...grpc.CallOption) (*MutateMediaFilesResponse, error)
}

type mediaFileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaFileServiceClient(cc grpc.ClientConnInterface) MediaFileServiceClient {
	return &mediaFileServiceClient{cc}
}

func (c *mediaFileServiceClient) GetMediaFile(ctx context.Context, in *GetMediaFileRequest, opts ...grpc.CallOption) (*resources.MediaFile, error) {
	out := new(resources.MediaFile)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.MediaFileService/GetMediaFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaFileServiceClient) MutateMediaFiles(ctx context.Context, in *MutateMediaFilesRequest, opts ...grpc.CallOption) (*MutateMediaFilesResponse, error) {
	out := new(MutateMediaFilesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.MediaFileService/MutateMediaFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaFileServiceServer is the server API for MediaFileService service.
type MediaFileServiceServer interface {
	// Returns the requested media file in full detail.
	GetMediaFile(context.Context, *GetMediaFileRequest) (*resources.MediaFile, error)
	// Creates media files. Operation statuses are returned.
	MutateMediaFiles(context.Context, *MutateMediaFilesRequest) (*MutateMediaFilesResponse, error)
}

// UnimplementedMediaFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMediaFileServiceServer struct {
}

func (*UnimplementedMediaFileServiceServer) GetMediaFile(ctx context.Context, req *GetMediaFileRequest) (*resources.MediaFile, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetMediaFile not implemented")
}
func (*UnimplementedMediaFileServiceServer) MutateMediaFiles(ctx context.Context, req *MutateMediaFilesRequest) (*MutateMediaFilesResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateMediaFiles not implemented")
}

func RegisterMediaFileServiceServer(s *grpc.Server, srv MediaFileServiceServer) {
	s.RegisterService(&_MediaFileService_serviceDesc, srv)
}

func _MediaFileService_GetMediaFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMediaFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaFileServiceServer).GetMediaFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.MediaFileService/GetMediaFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaFileServiceServer).GetMediaFile(ctx, req.(*GetMediaFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaFileService_MutateMediaFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateMediaFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaFileServiceServer).MutateMediaFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.MediaFileService/MutateMediaFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaFileServiceServer).MutateMediaFiles(ctx, req.(*MutateMediaFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MediaFileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.MediaFileService",
	HandlerType: (*MediaFileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMediaFile",
			Handler:    _MediaFileService_GetMediaFile_Handler,
		},
		{
			MethodName: "MutateMediaFiles",
			Handler:    _MediaFileService_MutateMediaFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/media_file_service.proto",
}