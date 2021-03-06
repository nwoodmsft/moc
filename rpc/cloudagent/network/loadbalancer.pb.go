// Code generated by protoc-gen-go. DO NOT EDIT.
// source: loadbalancer.proto

package network

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

type LoadBalancerRequest struct {
	LoadBalancers        []*LoadBalancer  `protobuf:"bytes,1,rep,name=LoadBalancers,proto3" json:"LoadBalancers,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LoadBalancerRequest) Reset()         { *m = LoadBalancerRequest{} }
func (m *LoadBalancerRequest) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerRequest) ProtoMessage()    {}
func (*LoadBalancerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d60ee00bafda1595, []int{0}
}

func (m *LoadBalancerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerRequest.Unmarshal(m, b)
}
func (m *LoadBalancerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerRequest.Marshal(b, m, deterministic)
}
func (m *LoadBalancerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerRequest.Merge(m, src)
}
func (m *LoadBalancerRequest) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerRequest.Size(m)
}
func (m *LoadBalancerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerRequest proto.InternalMessageInfo

func (m *LoadBalancerRequest) GetLoadBalancers() []*LoadBalancer {
	if m != nil {
		return m.LoadBalancers
	}
	return nil
}

func (m *LoadBalancerRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type LoadBalancerResponse struct {
	LoadBalancers        []*LoadBalancer     `protobuf:"bytes,1,rep,name=LoadBalancers,proto3" json:"LoadBalancers,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LoadBalancerResponse) Reset()         { *m = LoadBalancerResponse{} }
func (m *LoadBalancerResponse) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerResponse) ProtoMessage()    {}
func (*LoadBalancerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d60ee00bafda1595, []int{1}
}

func (m *LoadBalancerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerResponse.Unmarshal(m, b)
}
func (m *LoadBalancerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerResponse.Marshal(b, m, deterministic)
}
func (m *LoadBalancerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerResponse.Merge(m, src)
}
func (m *LoadBalancerResponse) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerResponse.Size(m)
}
func (m *LoadBalancerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerResponse proto.InternalMessageInfo

func (m *LoadBalancerResponse) GetLoadBalancers() []*LoadBalancer {
	if m != nil {
		return m.LoadBalancers
	}
	return nil
}

func (m *LoadBalancerResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *LoadBalancerResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type LoadBalancingRule struct {
	FrontendPort         uint32          `protobuf:"varint,1,opt,name=frontendPort,proto3" json:"frontendPort,omitempty"`
	BackendPort          uint32          `protobuf:"varint,2,opt,name=backendPort,proto3" json:"backendPort,omitempty"`
	Protocol             common.Protocol `protobuf:"varint,3,opt,name=protocol,proto3,enum=moc.Protocol" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LoadBalancingRule) Reset()         { *m = LoadBalancingRule{} }
func (m *LoadBalancingRule) String() string { return proto.CompactTextString(m) }
func (*LoadBalancingRule) ProtoMessage()    {}
func (*LoadBalancingRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_d60ee00bafda1595, []int{2}
}

func (m *LoadBalancingRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancingRule.Unmarshal(m, b)
}
func (m *LoadBalancingRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancingRule.Marshal(b, m, deterministic)
}
func (m *LoadBalancingRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancingRule.Merge(m, src)
}
func (m *LoadBalancingRule) XXX_Size() int {
	return xxx_messageInfo_LoadBalancingRule.Size(m)
}
func (m *LoadBalancingRule) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancingRule.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancingRule proto.InternalMessageInfo

func (m *LoadBalancingRule) GetFrontendPort() uint32 {
	if m != nil {
		return m.FrontendPort
	}
	return 0
}

func (m *LoadBalancingRule) GetBackendPort() uint32 {
	if m != nil {
		return m.BackendPort
	}
	return 0
}

func (m *LoadBalancingRule) GetProtocol() common.Protocol {
	if m != nil {
		return m.Protocol
	}
	return common.Protocol_All
}

type LoadBalancer struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	FrontendIP           string               `protobuf:"bytes,3,opt,name=frontendIP,proto3" json:"frontendIP,omitempty"`
	Backendpoolnames     []string             `protobuf:"bytes,4,rep,name=backendpoolnames,proto3" json:"backendpoolnames,omitempty"`
	Networkid            string               `protobuf:"bytes,5,opt,name=networkid,proto3" json:"networkid,omitempty"`
	Loadbalancingrules   []*LoadBalancingRule `protobuf:"bytes,6,rep,name=loadbalancingrules,proto3" json:"loadbalancingrules,omitempty"`
	Nodefqdn             string               `protobuf:"bytes,7,opt,name=nodefqdn,proto3" json:"nodefqdn,omitempty"`
	GroupName            string               `protobuf:"bytes,8,opt,name=groupName,proto3" json:"groupName,omitempty"`
	LocationName         string               `protobuf:"bytes,9,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Status               *common.Status       `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LoadBalancer) Reset()         { *m = LoadBalancer{} }
func (m *LoadBalancer) String() string { return proto.CompactTextString(m) }
func (*LoadBalancer) ProtoMessage()    {}
func (*LoadBalancer) Descriptor() ([]byte, []int) {
	return fileDescriptor_d60ee00bafda1595, []int{3}
}

func (m *LoadBalancer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancer.Unmarshal(m, b)
}
func (m *LoadBalancer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancer.Marshal(b, m, deterministic)
}
func (m *LoadBalancer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancer.Merge(m, src)
}
func (m *LoadBalancer) XXX_Size() int {
	return xxx_messageInfo_LoadBalancer.Size(m)
}
func (m *LoadBalancer) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancer.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancer proto.InternalMessageInfo

func (m *LoadBalancer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoadBalancer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LoadBalancer) GetFrontendIP() string {
	if m != nil {
		return m.FrontendIP
	}
	return ""
}

func (m *LoadBalancer) GetBackendpoolnames() []string {
	if m != nil {
		return m.Backendpoolnames
	}
	return nil
}

func (m *LoadBalancer) GetNetworkid() string {
	if m != nil {
		return m.Networkid
	}
	return ""
}

func (m *LoadBalancer) GetLoadbalancingrules() []*LoadBalancingRule {
	if m != nil {
		return m.Loadbalancingrules
	}
	return nil
}

func (m *LoadBalancer) GetNodefqdn() string {
	if m != nil {
		return m.Nodefqdn
	}
	return ""
}

func (m *LoadBalancer) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *LoadBalancer) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *LoadBalancer) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*LoadBalancerRequest)(nil), "moc.cloudagent.network.LoadBalancerRequest")
	proto.RegisterType((*LoadBalancerResponse)(nil), "moc.cloudagent.network.LoadBalancerResponse")
	proto.RegisterType((*LoadBalancingRule)(nil), "moc.cloudagent.network.LoadBalancingRule")
	proto.RegisterType((*LoadBalancer)(nil), "moc.cloudagent.network.LoadBalancer")
}

func init() { proto.RegisterFile("loadbalancer.proto", fileDescriptor_d60ee00bafda1595) }

var fileDescriptor_d60ee00bafda1595 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x25, 0xed, 0x6e, 0x69, 0xa6, 0x1f, 0x02, 0xef, 0x0a, 0x45, 0x15, 0x5a, 0x45, 0x81, 0x43,
	0x17, 0x50, 0x22, 0x02, 0x7f, 0x80, 0x4a, 0x1c, 0x16, 0x21, 0xa8, 0x0c, 0x42, 0x82, 0x5b, 0x9a,
	0xb8, 0x21, 0xaa, 0xe3, 0xc9, 0xda, 0x0e, 0x2b, 0x38, 0xf3, 0x1f, 0xf8, 0x13, 0xfc, 0x42, 0x4e,
	0xa8, 0x8e, 0xdb, 0x26, 0x62, 0x25, 0xf6, 0xb0, 0xb7, 0xf8, 0xf9, 0xf9, 0xcd, 0x9b, 0x99, 0x17,
	0x20, 0x1c, 0x93, 0x6c, 0x95, 0xf0, 0x44, 0xa4, 0x4c, 0x86, 0x95, 0x44, 0x8d, 0xe4, 0x41, 0x89,
	0x69, 0x98, 0x72, 0xac, 0xb3, 0x24, 0x67, 0x42, 0x87, 0x82, 0xe9, 0x2b, 0x94, 0x9b, 0xd9, 0x59,
	0x8e, 0x98, 0x73, 0x16, 0x19, 0xd6, 0xaa, 0x5e, 0x47, 0x57, 0x32, 0xa9, 0x2a, 0x26, 0x55, 0xf3,
	0x6e, 0x36, 0x4e, 0xb1, 0x2c, 0x51, 0xd8, 0xd3, 0x89, 0x7d, 0xd6, 0x06, 0x83, 0x5f, 0x0e, 0x9c,
	0xbc, 0xc5, 0x24, 0x5b, 0xd8, 0x8a, 0x94, 0x5d, 0xd6, 0x4c, 0x69, 0xf2, 0x06, 0x26, 0x6d, 0x58,
	0x79, 0x8e, 0xdf, 0x9f, 0x8f, 0xe2, 0xc7, 0xe1, 0xf5, 0x56, 0xc2, 0x8e, 0x46, 0xf7, 0x29, 0x79,
	0x09, 0x93, 0xf7, 0x15, 0x93, 0x89, 0x2e, 0x50, 0x7c, 0xfc, 0x5e, 0x31, 0xaf, 0xe7, 0x3b, 0xf3,
	0x69, 0x3c, 0x35, 0x5a, 0xfb, 0x1b, 0xda, 0x25, 0x05, 0xbf, 0x1d, 0x38, 0xed, 0x3a, 0x53, 0x15,
	0x0a, 0xc5, 0x6e, 0xd5, 0x5a, 0x0c, 0x03, 0xca, 0x54, 0xcd, 0xb5, 0xf1, 0x34, 0x8a, 0x67, 0x61,
	0x33, 0xd2, 0x70, 0x37, 0xd2, 0x70, 0x81, 0xc8, 0x3f, 0x25, 0xbc, 0x66, 0xd4, 0x32, 0xc9, 0x29,
	0x1c, 0xbf, 0x96, 0x12, 0xa5, 0xd7, 0xf7, 0x9d, 0xb9, 0x4b, 0x9b, 0x43, 0xf0, 0xd3, 0x81, 0xfb,
	0x07, 0xed, 0x42, 0xe4, 0xb4, 0xe6, 0x8c, 0x04, 0x30, 0x5e, 0x4b, 0x14, 0x9a, 0x89, 0x6c, 0x89,
	0x52, 0x7b, 0x8e, 0xef, 0xcc, 0x27, 0xb4, 0x83, 0x11, 0x1f, 0x46, 0xab, 0x24, 0xdd, 0xec, 0x28,
	0x3d, 0x43, 0x69, 0x43, 0xe4, 0x1c, 0x86, 0xc6, 0x4f, 0x8a, 0xdc, 0x14, 0x9d, 0xc6, 0x13, 0xd3,
	0xec, 0xd2, 0x82, 0x74, 0x7f, 0x1d, 0xfc, 0xe9, 0xc1, 0xb8, 0xdd, 0x22, 0x21, 0x70, 0x24, 0x92,
	0x92, 0x99, 0xca, 0x2e, 0x35, 0xdf, 0x64, 0x0a, 0xbd, 0x22, 0x33, 0x85, 0x5c, 0xda, 0x2b, 0x32,
	0x72, 0x06, 0xb0, 0x73, 0x74, 0xb1, 0xb4, 0x6d, 0xb5, 0x10, 0xf2, 0x04, 0xee, 0x59, 0x3b, 0x15,
	0x22, 0xdf, 0x4a, 0x28, 0xef, 0xc8, 0xef, 0xcf, 0x5d, 0xfa, 0x0f, 0x4e, 0x1e, 0x82, 0x6b, 0x07,
	0x5f, 0x64, 0xde, 0xb1, 0x91, 0x3a, 0x00, 0xe4, 0x73, 0x3b, 0xdf, 0x85, 0xc8, 0x65, 0xcd, 0x99,
	0xf2, 0x06, 0x66, 0x81, 0xe7, 0xff, 0x5f, 0xa0, 0x1d, 0x2b, 0xbd, 0x46, 0x84, 0xcc, 0x60, 0x28,
	0x30, 0x63, 0xeb, 0xcb, 0x4c, 0x78, 0x77, 0x4d, 0xdd, 0xfd, 0x79, 0x6b, 0x2a, 0x97, 0x58, 0x57,
	0xef, 0xb6, 0x93, 0x18, 0x36, 0xa6, 0xf6, 0xc0, 0x76, 0x49, 0x1c, 0x53, 0x93, 0x3c, 0x43, 0x70,
	0x0d, 0xa1, 0x83, 0x91, 0x47, 0x30, 0x50, 0x3a, 0xd1, 0xb5, 0xf2, 0xc0, 0x04, 0x65, 0x64, 0xcc,
	0x7e, 0x30, 0x10, 0xb5, 0x57, 0xf1, 0x8f, 0x76, 0x04, 0x98, 0x7c, 0xb5, 0xed, 0x82, 0x30, 0x18,
	0x5c, 0x88, 0x6f, 0xb8, 0x61, 0xe4, 0xe9, 0x8d, 0x12, 0xda, 0xfc, 0x80, 0xb3, 0x67, 0x37, 0x23,
	0x37, 0xff, 0x44, 0x70, 0x67, 0xf1, 0xfc, 0x4b, 0x94, 0x17, 0xfa, 0x6b, 0xbd, 0x0a, 0x53, 0x2c,
	0xa3, 0xb2, 0x48, 0x25, 0x2a, 0x5c, 0xeb, 0xa8, 0xc4, 0x34, 0x92, 0x55, 0x1a, 0x1d, 0x94, 0x22,
	0xab, 0xb4, 0x1a, 0x98, 0xd4, 0xbc, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xb3, 0x66, 0x95,
	0x73, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadBalancerAgentClient is the client API for LoadBalancerAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadBalancerAgentClient interface {
	Invoke(ctx context.Context, in *LoadBalancerRequest, opts ...grpc.CallOption) (*LoadBalancerResponse, error)
}

type loadBalancerAgentClient struct {
	cc *grpc.ClientConn
}

func NewLoadBalancerAgentClient(cc *grpc.ClientConn) LoadBalancerAgentClient {
	return &loadBalancerAgentClient{cc}
}

func (c *loadBalancerAgentClient) Invoke(ctx context.Context, in *LoadBalancerRequest, opts ...grpc.CallOption) (*LoadBalancerResponse, error) {
	out := new(LoadBalancerResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.network.LoadBalancerAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoadBalancerAgentServer is the server API for LoadBalancerAgent service.
type LoadBalancerAgentServer interface {
	Invoke(context.Context, *LoadBalancerRequest) (*LoadBalancerResponse, error)
}

// UnimplementedLoadBalancerAgentServer can be embedded to have forward compatible implementations.
type UnimplementedLoadBalancerAgentServer struct {
}

func (*UnimplementedLoadBalancerAgentServer) Invoke(ctx context.Context, req *LoadBalancerRequest) (*LoadBalancerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterLoadBalancerAgentServer(s *grpc.Server, srv LoadBalancerAgentServer) {
	s.RegisterService(&_LoadBalancerAgent_serviceDesc, srv)
}

func _LoadBalancerAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.network.LoadBalancerAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerAgentServer).Invoke(ctx, req.(*LoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadBalancerAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.network.LoadBalancerAgent",
	HandlerType: (*LoadBalancerAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _LoadBalancerAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loadbalancer.proto",
}
