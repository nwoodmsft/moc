// Code generated by protoc-gen-go. DO NOT EDIT.
// source: location.proto

package cloud

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type LocationRequest struct {
	Locations            []*Location      `protobuf:"bytes,1,rep,name=Locations,proto3" json:"Locations,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LocationRequest) Reset()         { *m = LocationRequest{} }
func (m *LocationRequest) String() string { return proto.CompactTextString(m) }
func (*LocationRequest) ProtoMessage()    {}
func (*LocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0f35158dcf9f2c, []int{0}
}

func (m *LocationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationRequest.Unmarshal(m, b)
}
func (m *LocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationRequest.Marshal(b, m, deterministic)
}
func (m *LocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationRequest.Merge(m, src)
}
func (m *LocationRequest) XXX_Size() int {
	return xxx_messageInfo_LocationRequest.Size(m)
}
func (m *LocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LocationRequest proto.InternalMessageInfo

func (m *LocationRequest) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *LocationRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type LocationResponse struct {
	Locations            []*Location         `protobuf:"bytes,1,rep,name=Locations,proto3" json:"Locations,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LocationResponse) Reset()         { *m = LocationResponse{} }
func (m *LocationResponse) String() string { return proto.CompactTextString(m) }
func (*LocationResponse) ProtoMessage()    {}
func (*LocationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0f35158dcf9f2c, []int{1}
}

func (m *LocationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationResponse.Unmarshal(m, b)
}
func (m *LocationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationResponse.Marshal(b, m, deterministic)
}
func (m *LocationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationResponse.Merge(m, src)
}
func (m *LocationResponse) XXX_Size() int {
	return xxx_messageInfo_LocationResponse.Size(m)
}
func (m *LocationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LocationResponse proto.InternalMessageInfo

func (m *LocationResponse) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *LocationResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *LocationResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Location struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Status               *common.Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0f35158dcf9f2c, []int{2}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Location) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Location) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*LocationRequest)(nil), "moc.cloudagent.location.LocationRequest")
	proto.RegisterType((*LocationResponse)(nil), "moc.cloudagent.location.LocationResponse")
	proto.RegisterType((*Location)(nil), "moc.cloudagent.location.Location")
}

func init() { proto.RegisterFile("location.proto", fileDescriptor_4f0f35158dcf9f2c) }

var fileDescriptor_4f0f35158dcf9f2c = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x7d, 0x69, 0xdf, 0x0b, 0xaf, 0x93, 0xd7, 0x3c, 0x19, 0x04, 0x43, 0x16, 0x52, 0xe3, 0x26,
	0x2e, 0x9c, 0x81, 0xe8, 0x5e, 0x2c, 0xb8, 0x10, 0x04, 0x61, 0x2a, 0x2e, 0xdc, 0x48, 0x3a, 0x9d,
	0xc6, 0x60, 0x26, 0x77, 0x9c, 0x0f, 0xc5, 0x7f, 0xe0, 0x9f, 0xf0, 0xbf, 0x8a, 0x93, 0xa6, 0xc5,
	0x85, 0x28, 0xb8, 0x9b, 0x7b, 0xef, 0x39, 0xe7, 0x9e, 0x7b, 0x06, 0xc5, 0x0d, 0xf0, 0xd2, 0xd6,
	0xd0, 0x12, 0xa5, 0xc1, 0x02, 0xde, 0x91, 0xc0, 0x09, 0x6f, 0xc0, 0x2d, 0xca, 0x4a, 0xb4, 0x96,
	0xf4, 0xe3, 0x74, 0xb7, 0x02, 0xa8, 0x1a, 0x41, 0x3d, 0x6c, 0xee, 0x96, 0xf4, 0x49, 0x97, 0x4a,
	0x09, 0x6d, 0x3a, 0x62, 0xfa, 0x8f, 0x83, 0x94, 0xbd, 0x4c, 0xf6, 0x12, 0xa0, 0xff, 0x17, 0x2b,
	0x2a, 0x13, 0x0f, 0x4e, 0x18, 0x8b, 0x4f, 0xd0, 0xa8, 0x6f, 0x99, 0x24, 0x98, 0x0c, 0xf3, 0xa8,
	0xd8, 0x23, 0x9f, 0xac, 0x23, 0x6b, 0xf2, 0x86, 0x83, 0x8f, 0xd1, 0xf8, 0x52, 0x09, 0xed, 0xab,
	0xab, 0x67, 0x25, 0x92, 0xc1, 0x24, 0xc8, 0xe3, 0x22, 0xf6, 0x22, 0xeb, 0x09, 0xfb, 0x08, 0xca,
	0x5e, 0x03, 0xb4, 0xb5, 0xb1, 0x62, 0x14, 0xb4, 0x46, 0xfc, 0xdc, 0x4b, 0x81, 0x42, 0x26, 0x8c,
	0x6b, 0xac, 0x37, 0x11, 0x15, 0x29, 0xe9, 0xf2, 0x21, 0x7d, 0x3e, 0x64, 0x0a, 0xd0, 0x5c, 0x97,
	0x8d, 0x13, 0x6c, 0x85, 0xc4, 0xdb, 0xe8, 0xcf, 0x99, 0xd6, 0xa0, 0x93, 0xe1, 0x24, 0xc8, 0x47,
	0xac, 0x2b, 0xb2, 0x19, 0xfa, 0xdb, 0xcb, 0x62, 0x8c, 0x7e, 0xb7, 0xa5, 0x14, 0x49, 0xe0, 0x01,
	0xfe, 0x8d, 0x63, 0x34, 0xa8, 0x17, 0x7e, 0xcb, 0x88, 0x0d, 0xea, 0x05, 0xde, 0x47, 0xa1, 0xb1,
	0xa5, 0x75, 0xc6, 0xcb, 0x44, 0x45, 0xe4, 0x7d, 0xcf, 0x7c, 0x8b, 0xad, 0x46, 0x85, 0x42, 0xe3,
	0x5e, 0xf4, 0xf4, 0xfd, 0x18, 0x7c, 0x8b, 0xc2, 0xf3, 0xf6, 0x11, 0xee, 0x05, 0xce, 0xbf, 0xbe,
	0xb3, 0xfb, 0xb0, 0xf4, 0xe0, 0x1b, 0xc8, 0x2e, 0xcf, 0xec, 0xd7, 0x94, 0xde, 0x1c, 0x56, 0xb5,
	0xbd, 0x73, 0x73, 0xc2, 0x41, 0x52, 0x59, 0x73, 0x0d, 0x06, 0x96, 0x96, 0x4a, 0xe0, 0x54, 0x2b,
	0x4e, 0x37, 0x32, 0xdd, 0x73, 0x1e, 0xfa, 0xa4, 0x8e, 0xde, 0x02, 0x00, 0x00, 0xff, 0xff, 0x69,
	0x7b, 0x98, 0xa2, 0x82, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LocationAgentClient is the client API for LocationAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LocationAgentClient interface {
	Invoke(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*LocationResponse, error)
}

type locationAgentClient struct {
	cc *grpc.ClientConn
}

func NewLocationAgentClient(cc *grpc.ClientConn) LocationAgentClient {
	return &locationAgentClient{cc}
}

func (c *locationAgentClient) Invoke(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*LocationResponse, error) {
	out := new(LocationResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.location.LocationAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationAgentServer is the server API for LocationAgent service.
type LocationAgentServer interface {
	Invoke(context.Context, *LocationRequest) (*LocationResponse, error)
}

// UnimplementedLocationAgentServer can be embedded to have forward compatible implementations.
type UnimplementedLocationAgentServer struct {
}

func (*UnimplementedLocationAgentServer) Invoke(ctx context.Context, req *LocationRequest) (*LocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterLocationAgentServer(s *grpc.Server, srv LocationAgentServer) {
	s.RegisterService(&_LocationAgent_serviceDesc, srv)
}

func _LocationAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.location.LocationAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationAgentServer).Invoke(ctx, req.(*LocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LocationAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.location.LocationAgent",
	HandlerType: (*LocationAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _LocationAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "location.proto",
}
