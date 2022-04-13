// Code generated by protoc-gen-go. DO NOT EDIT.
// source: countries.proto

package countries

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CountryRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountryRequest) Reset()         { *m = CountryRequest{} }
func (m *CountryRequest) String() string { return proto.CompactTextString(m) }
func (*CountryRequest) ProtoMessage()    {}
func (*CountryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3604625e838084c3, []int{0}
}

func (m *CountryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountryRequest.Unmarshal(m, b)
}
func (m *CountryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountryRequest.Marshal(b, m, deterministic)
}
func (m *CountryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountryRequest.Merge(m, src)
}
func (m *CountryRequest) XXX_Size() int {
	return xxx_messageInfo_CountryRequest.Size(m)
}
func (m *CountryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CountryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CountryRequest proto.InternalMessageInfo

func (m *CountryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Currencies struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol               string   `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Currencies) Reset()         { *m = Currencies{} }
func (m *Currencies) String() string { return proto.CompactTextString(m) }
func (*Currencies) ProtoMessage()    {}
func (*Currencies) Descriptor() ([]byte, []int) {
	return fileDescriptor_3604625e838084c3, []int{1}
}

func (m *Currencies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Currencies.Unmarshal(m, b)
}
func (m *Currencies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Currencies.Marshal(b, m, deterministic)
}
func (m *Currencies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Currencies.Merge(m, src)
}
func (m *Currencies) XXX_Size() int {
	return xxx_messageInfo_Currencies.Size(m)
}
func (m *Currencies) XXX_DiscardUnknown() {
	xxx_messageInfo_Currencies.DiscardUnknown(m)
}

var xxx_messageInfo_Currencies proto.InternalMessageInfo

func (m *Currencies) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Currencies) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Currencies) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type CountryResponse struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Alpha2Code           string        `protobuf:"bytes,2,opt,name=alpha2Code,proto3" json:"alpha2Code,omitempty"`
	Capital              string        `protobuf:"bytes,3,opt,name=capital,proto3" json:"capital,omitempty"`
	Subregion            string        `protobuf:"bytes,4,opt,name=subregion,proto3" json:"subregion,omitempty"`
	Population           int32         `protobuf:"varint,5,opt,name=population,proto3" json:"population,omitempty"`
	NativeName           string        `protobuf:"bytes,6,opt,name=nativeName,proto3" json:"nativeName,omitempty"`
	Currencies           []*Currencies `protobuf:"bytes,7,rep,name=currencies,proto3" json:"currencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CountryResponse) Reset()         { *m = CountryResponse{} }
func (m *CountryResponse) String() string { return proto.CompactTextString(m) }
func (*CountryResponse) ProtoMessage()    {}
func (*CountryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3604625e838084c3, []int{2}
}

func (m *CountryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountryResponse.Unmarshal(m, b)
}
func (m *CountryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountryResponse.Marshal(b, m, deterministic)
}
func (m *CountryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountryResponse.Merge(m, src)
}
func (m *CountryResponse) XXX_Size() int {
	return xxx_messageInfo_CountryResponse.Size(m)
}
func (m *CountryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountryResponse proto.InternalMessageInfo

func (m *CountryResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CountryResponse) GetAlpha2Code() string {
	if m != nil {
		return m.Alpha2Code
	}
	return ""
}

func (m *CountryResponse) GetCapital() string {
	if m != nil {
		return m.Capital
	}
	return ""
}

func (m *CountryResponse) GetSubregion() string {
	if m != nil {
		return m.Subregion
	}
	return ""
}

func (m *CountryResponse) GetPopulation() int32 {
	if m != nil {
		return m.Population
	}
	return 0
}

func (m *CountryResponse) GetNativeName() string {
	if m != nil {
		return m.NativeName
	}
	return ""
}

func (m *CountryResponse) GetCurrencies() []*Currencies {
	if m != nil {
		return m.Currencies
	}
	return nil
}

func init() {
	proto.RegisterType((*CountryRequest)(nil), "countries.CountryRequest")
	proto.RegisterType((*Currencies)(nil), "countries.Currencies")
	proto.RegisterType((*CountryResponse)(nil), "countries.CountryResponse")
}

func init() {
	proto.RegisterFile("countries.proto", fileDescriptor_3604625e838084c3)
}

var fileDescriptor_3604625e838084c3 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x35, 0xb6, 0x4d, 0xc8, 0x08, 0x16, 0x06, 0x94, 0x58, 0x44, 0x42, 0xf0, 0x90, 0x53, 0x0f,
	0x11, 0x3f, 0x40, 0x72, 0x2d, 0x1e, 0xe2, 0x17, 0x6c, 0xd6, 0xc1, 0x06, 0xd2, 0xdd, 0x75, 0x77,
	0x23, 0xf4, 0xbf, 0xfd, 0x00, 0xd9, 0xa4, 0xdd, 0x8d, 0xd0, 0xdb, 0xcc, 0x7b, 0xf3, 0xe6, 0xed,
	0xbe, 0x81, 0x35, 0x97, 0x83, 0xb0, 0xba, 0x23, 0xb3, 0x55, 0x5a, 0x5a, 0x89, 0xa9, 0x07, 0x8a,
	0x67, 0xb8, 0xad, 0xc7, 0xe6, 0xd8, 0xd0, 0xf7, 0x40, 0xc6, 0x22, 0xc2, 0x52, 0xb0, 0x03, 0x65,
	0x51, 0x1e, 0x95, 0x69, 0x33, 0xd6, 0xc5, 0x0e, 0xa0, 0x1e, 0xb4, 0x26, 0xc1, 0x3b, 0x32, 0x6e,
	0x82, 0xcb, 0x4f, 0x3f, 0xe1, 0x6a, 0xaf, 0xba, 0x0e, 0x2a, 0xbc, 0x87, 0xd8, 0x1c, 0x0f, 0xad,
	0xec, 0xb3, 0xc5, 0x88, 0x9e, 0xba, 0xe2, 0x37, 0x82, 0xb5, 0x37, 0x35, 0x4a, 0x0a, 0x43, 0x97,
	0x5c, 0xf1, 0x09, 0x80, 0xf5, 0x6a, 0xcf, 0xaa, 0xda, 0xb9, 0x4d, 0x9b, 0x67, 0x08, 0x66, 0x90,
	0x70, 0xa6, 0x3a, 0xcb, 0xce, 0x06, 0xe7, 0x16, 0x1f, 0x21, 0x35, 0x43, 0xab, 0xe9, 0xab, 0x93,
	0x22, 0x5b, 0x8e, 0x5c, 0x00, 0xdc, 0x5e, 0x25, 0xd5, 0xd0, 0x33, 0xeb, 0xe8, 0x55, 0x1e, 0x95,
	0xab, 0x66, 0x86, 0x38, 0x5e, 0x30, 0xdb, 0xfd, 0xd0, 0xbb, 0x7b, 0x51, 0x3c, 0xf9, 0x06, 0x04,
	0x5f, 0x01, 0xb8, 0x4f, 0x23, 0x4b, 0xf2, 0x45, 0x79, 0x53, 0xdd, 0x6d, 0x43, 0xc8, 0x21, 0xaa,
	0x66, 0x36, 0x58, 0xed, 0x20, 0x39, 0xfd, 0x1a, 0xdf, 0x20, 0xfe, 0x20, 0xa6, 0xf9, 0x1e, 0x1f,
	0xe6, 0xba, 0x7f, 0x87, 0xd8, 0x6c, 0x2e, 0x51, 0x53, 0x5c, 0xc5, 0x55, 0x1b, 0x8f, 0xa7, 0x7c,
	0xf9, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x80, 0x42, 0xf2, 0xbc, 0xdd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CountryClient is the client API for Country service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CountryClient interface {
	Search(ctx context.Context, in *CountryRequest, opts ...grpc.CallOption) (*CountryResponse, error)
}

type countryClient struct {
	cc grpc.ClientConnInterface
}

func NewCountryClient(cc grpc.ClientConnInterface) CountryClient {
	return &countryClient{cc}
}

func (c *countryClient) Search(ctx context.Context, in *CountryRequest, opts ...grpc.CallOption) (*CountryResponse, error) {
	out := new(CountryResponse)
	err := c.cc.Invoke(ctx, "/countries.Country/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountryServer is the server API for Country service.
type CountryServer interface {
	Search(context.Context, *CountryRequest) (*CountryResponse, error)
}

// UnimplementedCountryServer can be embedded to have forward compatible implementations.
type UnimplementedCountryServer struct {
}

func (*UnimplementedCountryServer) Search(ctx context.Context, req *CountryRequest) (*CountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterCountryServer(s *grpc.Server, srv CountryServer) {
	s.RegisterService(&_Country_serviceDesc, srv)
}

func _Country_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/countries.Country/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServer).Search(ctx, req.(*CountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Country_serviceDesc = grpc.ServiceDesc{
	ServiceName: "countries.Country",
	HandlerType: (*CountryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Country_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "countries.proto",
}
