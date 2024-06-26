// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: inventory.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName string `protobuf:"bytes,1,opt,name=ProductName,proto3" json:"ProductName,omitempty"`
	Quantity    int32  `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
}

func (x *OrderDetails) Reset() {
	*x = OrderDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDetails) ProtoMessage() {}

func (x *OrderDetails) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDetails.ProtoReflect.Descriptor instead.
func (*OrderDetails) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *OrderDetails) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *OrderDetails) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ArrivalDate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArrivalDate string `protobuf:"bytes,1,opt,name=ArrivalDate,proto3" json:"ArrivalDate,omitempty"`
}

func (x *ArrivalDate) Reset() {
	*x = ArrivalDate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrivalDate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrivalDate) ProtoMessage() {}

func (x *ArrivalDate) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrivalDate.ProtoReflect.Descriptor instead.
func (*ArrivalDate) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *ArrivalDate) GetArrivalDate() string {
	if x != nil {
		return x.ArrivalDate
	}
	return ""
}

var File_inventory_proto protoreflect.FileDescriptor

var file_inventory_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x2f, 0x0a, 0x0b, 0x41, 0x72, 0x72, 0x69, 0x76, 0x61,
	0x6c, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x72, 0x72, 0x69,
	0x76, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x32, 0x4a, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x50,
	0x61, 0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x44, 0x61, 0x74,
	0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_proto_rawDescOnce sync.Once
	file_inventory_proto_rawDescData = file_inventory_proto_rawDesc
)

func file_inventory_proto_rawDescGZIP() []byte {
	file_inventory_proto_rawDescOnce.Do(func() {
		file_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_proto_rawDescData)
	})
	return file_inventory_proto_rawDescData
}

var file_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_inventory_proto_goTypes = []interface{}{
	(*OrderDetails)(nil), // 0: proto.OrderDetails
	(*ArrivalDate)(nil),  // 1: proto.ArrivalDate
}
var file_inventory_proto_depIdxs = []int32{
	0, // 0: proto.InventoryService.PackOrder:input_type -> proto.OrderDetails
	1, // 1: proto.InventoryService.PackOrder:output_type -> proto.ArrivalDate
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_inventory_proto_init() }
func file_inventory_proto_init() {
	if File_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrivalDate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_proto_goTypes,
		DependencyIndexes: file_inventory_proto_depIdxs,
		MessageInfos:      file_inventory_proto_msgTypes,
	}.Build()
	File_inventory_proto = out.File
	file_inventory_proto_rawDesc = nil
	file_inventory_proto_goTypes = nil
	file_inventory_proto_depIdxs = nil
}
