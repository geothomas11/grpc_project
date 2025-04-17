// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: pkg/pb/product.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Stock         int64                  `protobuf:"varint,2,opt,name=stock,proto3" json:"stock,omitempty"`
	Price         int64                  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	mi := &file_pkg_pb_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *CreateProductRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int64                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id            int64                  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductResponse) Reset() {
	*x = CreateProductResponse{}
	mi := &file_pkg_pb_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductResponse) ProtoMessage() {}

func (x *CreateProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductResponse.ProtoReflect.Descriptor instead.
func (*CreateProductResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateProductResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateProductResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindOneData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Stock         int64                  `protobuf:"varint,3,opt,name=stock,proto3" json:"stock,omitempty"`
	Price         int64                  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindOneData) Reset() {
	*x = FindOneData{}
	mi := &file_pkg_pb_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindOneData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneData) ProtoMessage() {}

func (x *FindOneData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneData.ProtoReflect.Descriptor instead.
func (*FindOneData) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{2}
}

func (x *FindOneData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FindOneData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindOneData) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *FindOneData) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type FindOneRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindOneRequest) Reset() {
	*x = FindOneRequest{}
	mi := &file_pkg_pb_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneRequest) ProtoMessage() {}

func (x *FindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneRequest.ProtoReflect.Descriptor instead.
func (*FindOneRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{3}
}

func (x *FindOneRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindOneResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int64                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data          *FindOneData           `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindOneResponse) Reset() {
	*x = FindOneResponse{}
	mi := &file_pkg_pb_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneResponse) ProtoMessage() {}

func (x *FindOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneResponse.ProtoReflect.Descriptor instead.
func (*FindOneResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{4}
}

func (x *FindOneResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindOneResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindOneResponse) GetData() *FindOneData {
	if x != nil {
		return x.Data
	}
	return nil
}

type DecreaseStockRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId       int64                  `protobuf:"varint,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DecreaseStockRequest) Reset() {
	*x = DecreaseStockRequest{}
	mi := &file_pkg_pb_product_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecreaseStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseStockRequest) ProtoMessage() {}

func (x *DecreaseStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseStockRequest.ProtoReflect.Descriptor instead.
func (*DecreaseStockRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{5}
}

func (x *DecreaseStockRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DecreaseStockRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type DecreaseStockResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int64                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DecreaseStockResponse) Reset() {
	*x = DecreaseStockResponse{}
	mi := &file_pkg_pb_product_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecreaseStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseStockResponse) ProtoMessage() {}

func (x *DecreaseStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseStockResponse.ProtoReflect.Descriptor instead.
func (*DecreaseStockResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{6}
}

func (x *DecreaseStockResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DecreaseStockResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_pb_product_proto protoreflect.FileDescriptor

const file_pkg_pb_product_proto_rawDesc = "" +
	"\n" +
	"\x14pkg/pb/product.proto\x12\aproduct\"V\n" +
	"\x14CreateProductRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05stock\x18\x02 \x01(\x03R\x05stock\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x03R\x05price\"U\n" +
	"\x15CreateProductResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\x03R\x06status\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error\x12\x0e\n" +
	"\x02id\x18\x03 \x01(\x03R\x02id\"]\n" +
	"\vFindOneData\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05stock\x18\x03 \x01(\x03R\x05stock\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x03R\x05price\" \n" +
	"\x0eFindOneRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"i\n" +
	"\x0fFindOneResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\x03R\x06status\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error\x12(\n" +
	"\x04data\x18\x03 \x01(\v2\x14.product.FindOneDataR\x04data\"@\n" +
	"\x14DecreaseStockRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x18\n" +
	"\aorderId\x18\x02 \x01(\x03R\aorderId\"E\n" +
	"\x15DecreaseStockResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\x03R\x06status\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error2\xf4\x01\n" +
	"\x0eProductService\x12P\n" +
	"\rCreateProduct\x12\x1d.product.CreateProductRequest\x1a\x1e.product.CreateProductResponse\"\x00\x12>\n" +
	"\aFindOne\x12\x17.product.FindOneRequest\x1a\x18.product.FindOneResponse\"\x00\x12P\n" +
	"\rDecreaseStock\x12\x1d.product.DecreaseStockRequest\x1a\x1e.product.DecreaseStockResponse\"\x00B\n" +
	"Z\b./pkg/pbb\x06proto3"

var (
	file_pkg_pb_product_proto_rawDescOnce sync.Once
	file_pkg_pb_product_proto_rawDescData []byte
)

func file_pkg_pb_product_proto_rawDescGZIP() []byte {
	file_pkg_pb_product_proto_rawDescOnce.Do(func() {
		file_pkg_pb_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_pb_product_proto_rawDesc), len(file_pkg_pb_product_proto_rawDesc)))
	})
	return file_pkg_pb_product_proto_rawDescData
}

var file_pkg_pb_product_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_pb_product_proto_goTypes = []any{
	(*CreateProductRequest)(nil),  // 0: product.CreateProductRequest
	(*CreateProductResponse)(nil), // 1: product.CreateProductResponse
	(*FindOneData)(nil),           // 2: product.FindOneData
	(*FindOneRequest)(nil),        // 3: product.FindOneRequest
	(*FindOneResponse)(nil),       // 4: product.FindOneResponse
	(*DecreaseStockRequest)(nil),  // 5: product.DecreaseStockRequest
	(*DecreaseStockResponse)(nil), // 6: product.DecreaseStockResponse
}
var file_pkg_pb_product_proto_depIdxs = []int32{
	2, // 0: product.FindOneResponse.data:type_name -> product.FindOneData
	0, // 1: product.ProductService.CreateProduct:input_type -> product.CreateProductRequest
	3, // 2: product.ProductService.FindOne:input_type -> product.FindOneRequest
	5, // 3: product.ProductService.DecreaseStock:input_type -> product.DecreaseStockRequest
	1, // 4: product.ProductService.CreateProduct:output_type -> product.CreateProductResponse
	4, // 5: product.ProductService.FindOne:output_type -> product.FindOneResponse
	6, // 6: product.ProductService.DecreaseStock:output_type -> product.DecreaseStockResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_pb_product_proto_init() }
func file_pkg_pb_product_proto_init() {
	if File_pkg_pb_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_pb_product_proto_rawDesc), len(file_pkg_pb_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_product_proto_goTypes,
		DependencyIndexes: file_pkg_pb_product_proto_depIdxs,
		MessageInfos:      file_pkg_pb_product_proto_msgTypes,
	}.Build()
	File_pkg_pb_product_proto = out.File
	file_pkg_pb_product_proto_goTypes = nil
	file_pkg_pb_product_proto_depIdxs = nil
}
