// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.5
// source: protobuf/sneaker.proto

package models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sneaker_SiteSoldOn int32

const (
	Sneaker_STOCKX       Sneaker_SiteSoldOn = 0
	Sneaker_NIKE         Sneaker_SiteSoldOn = 1
	Sneaker_ADIDAS       Sneaker_SiteSoldOn = 2
	Sneaker_PUMA         Sneaker_SiteSoldOn = 3
	Sneaker_STADIUMGOODS Sneaker_SiteSoldOn = 4
	Sneaker_FLIGHTCLUB   Sneaker_SiteSoldOn = 5
)

// Enum value maps for Sneaker_SiteSoldOn.
var (
	Sneaker_SiteSoldOn_name = map[int32]string{
		0: "STOCKX",
		1: "NIKE",
		2: "ADIDAS",
		3: "PUMA",
		4: "STADIUMGOODS",
		5: "FLIGHTCLUB",
	}
	Sneaker_SiteSoldOn_value = map[string]int32{
		"STOCKX":       0,
		"NIKE":         1,
		"ADIDAS":       2,
		"PUMA":         3,
		"STADIUMGOODS": 4,
		"FLIGHTCLUB":   5,
	}
)

func (x Sneaker_SiteSoldOn) Enum() *Sneaker_SiteSoldOn {
	p := new(Sneaker_SiteSoldOn)
	*p = x
	return p
}

func (x Sneaker_SiteSoldOn) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sneaker_SiteSoldOn) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_sneaker_proto_enumTypes[0].Descriptor()
}

func (Sneaker_SiteSoldOn) Type() protoreflect.EnumType {
	return &file_protobuf_sneaker_proto_enumTypes[0]
}

func (x Sneaker_SiteSoldOn) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sneaker_SiteSoldOn.Descriptor instead.
func (Sneaker_SiteSoldOn) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_sneaker_proto_rawDescGZIP(), []int{0, 0}
}

type Sneaker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price       int64                  `protobuf:"varint,1,opt,name=price,proto3" json:"price,omitempty"`
	Brand       string                 `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Model       string                 `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Sku         string                 `protobuf:"bytes,4,opt,name=sku,proto3" json:"sku,omitempty"`
	ReleaseDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	Sites       []Sneaker_SiteSoldOn   `protobuf:"varint,6,rep,packed,name=sites,proto3,enum=protobuf.Sneaker_SiteSoldOn" json:"sites,omitempty"`
}

func (x *Sneaker) Reset() {
	*x = Sneaker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_sneaker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sneaker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sneaker) ProtoMessage() {}

func (x *Sneaker) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_sneaker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sneaker.ProtoReflect.Descriptor instead.
func (*Sneaker) Descriptor() ([]byte, []int) {
	return file_protobuf_sneaker_proto_rawDescGZIP(), []int{0}
}

func (x *Sneaker) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Sneaker) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Sneaker) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Sneaker) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *Sneaker) GetReleaseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ReleaseDate
	}
	return nil
}

func (x *Sneaker) GetSites() []Sneaker_SiteSoldOn {
	if x != nil {
		return x.Sites
	}
	return nil
}

var File_protobuf_sneaker_proto protoreflect.FileDescriptor

var file_protobuf_sneaker_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x6e, 0x65, 0x61, 0x6b,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xac, 0x02, 0x0a, 0x07, 0x53, 0x6e, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x6b, 0x75, 0x12, 0x3d, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x69, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x6e, 0x65,
	0x61, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x64, 0x4f, 0x6e, 0x52,
	0x05, 0x73, 0x69, 0x74, 0x65, 0x73, 0x22, 0x5a, 0x0a, 0x0a, 0x53, 0x69, 0x74, 0x65, 0x53, 0x6f,
	0x6c, 0x64, 0x4f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x4f, 0x43, 0x4b, 0x58, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x49, 0x4b, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x44,
	0x49, 0x44, 0x41, 0x53, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x55, 0x4d, 0x41, 0x10, 0x03,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x44, 0x49, 0x55, 0x4d, 0x47, 0x4f, 0x4f, 0x44, 0x53,
	0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x43, 0x4c, 0x55, 0x42,
	0x10, 0x05, 0x42, 0x0c, 0x5a, 0x0a, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_sneaker_proto_rawDescOnce sync.Once
	file_protobuf_sneaker_proto_rawDescData = file_protobuf_sneaker_proto_rawDesc
)

func file_protobuf_sneaker_proto_rawDescGZIP() []byte {
	file_protobuf_sneaker_proto_rawDescOnce.Do(func() {
		file_protobuf_sneaker_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_sneaker_proto_rawDescData)
	})
	return file_protobuf_sneaker_proto_rawDescData
}

var file_protobuf_sneaker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_sneaker_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_sneaker_proto_goTypes = []interface{}{
	(Sneaker_SiteSoldOn)(0),       // 0: protobuf.Sneaker.SiteSoldOn
	(*Sneaker)(nil),               // 1: protobuf.Sneaker
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_protobuf_sneaker_proto_depIdxs = []int32{
	2, // 0: protobuf.Sneaker.release_date:type_name -> google.protobuf.Timestamp
	0, // 1: protobuf.Sneaker.sites:type_name -> protobuf.Sneaker.SiteSoldOn
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protobuf_sneaker_proto_init() }
func file_protobuf_sneaker_proto_init() {
	if File_protobuf_sneaker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_sneaker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sneaker); i {
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
			RawDescriptor: file_protobuf_sneaker_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_sneaker_proto_goTypes,
		DependencyIndexes: file_protobuf_sneaker_proto_depIdxs,
		EnumInfos:         file_protobuf_sneaker_proto_enumTypes,
		MessageInfos:      file_protobuf_sneaker_proto_msgTypes,
	}.Build()
	File_protobuf_sneaker_proto = out.File
	file_protobuf_sneaker_proto_rawDesc = nil
	file_protobuf_sneaker_proto_goTypes = nil
	file_protobuf_sneaker_proto_depIdxs = nil
}
