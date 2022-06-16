// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: message/dialog.proto

package animal

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

type AnimalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Food string `protobuf:"bytes,1,opt,name=food,proto3" json:"food,omitempty"`
}

func (x *AnimalRequest) Reset() {
	*x = AnimalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_dialog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimalRequest) ProtoMessage() {}

func (x *AnimalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_dialog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimalRequest.ProtoReflect.Descriptor instead.
func (*AnimalRequest) Descriptor() ([]byte, []int) {
	return file_message_dialog_proto_rawDescGZIP(), []int{0}
}

func (x *AnimalRequest) GetFood() string {
	if x != nil {
		return x.Food
	}
	return ""
}

type AnimalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eat string `protobuf:"bytes,1,opt,name=eat,proto3" json:"eat,omitempty"`
}

func (x *AnimalResponse) Reset() {
	*x = AnimalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_dialog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimalResponse) ProtoMessage() {}

func (x *AnimalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_dialog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimalResponse.ProtoReflect.Descriptor instead.
func (*AnimalResponse) Descriptor() ([]byte, []int) {
	return file_message_dialog_proto_rawDescGZIP(), []int{1}
}

func (x *AnimalResponse) GetEat() string {
	if x != nil {
		return x.Eat
	}
	return ""
}

var File_message_dialog_proto protoreflect.FileDescriptor

var file_message_dialog_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x22, 0x23,
	0x0a, 0x0d, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x6f, 0x6f, 0x64, 0x22, 0x22, 0x0a, 0x0e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x65, 0x61, 0x74, 0x32, 0x9a, 0x02, 0x0a, 0x0c, 0x41, 0x6e, 0x69, 0x6d,
	0x61, 0x6c, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x3c, 0x0a, 0x09, 0x55, 0x6e, 0x69, 0x41,
	0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x15, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x41,
	0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61,
	0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x15, 0x2e, 0x61, 0x6e, 0x69,
	0x6d, 0x61, 0x6c, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x44, 0x0a,
	0x0f, 0x43, 0x6c, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c,
	0x12, 0x15, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c,
	0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x12, 0x3f, 0x0a, 0x08, 0x42, 0x69, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x12,
	0x15, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2e,
	0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x2f, 0x61, 0x6e,
	0x69, 0x6d, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_dialog_proto_rawDescOnce sync.Once
	file_message_dialog_proto_rawDescData = file_message_dialog_proto_rawDesc
)

func file_message_dialog_proto_rawDescGZIP() []byte {
	file_message_dialog_proto_rawDescOnce.Do(func() {
		file_message_dialog_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_dialog_proto_rawDescData)
	})
	return file_message_dialog_proto_rawDescData
}

var file_message_dialog_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_dialog_proto_goTypes = []interface{}{
	(*AnimalRequest)(nil),  // 0: animal.AnimalRequest
	(*AnimalResponse)(nil), // 1: animal.AnimalResponse
}
var file_message_dialog_proto_depIdxs = []int32{
	0, // 0: animal.AnimalDialog.UniAnimal:input_type -> animal.AnimalRequest
	0, // 1: animal.AnimalDialog.ServStreamAnimal:input_type -> animal.AnimalRequest
	0, // 2: animal.AnimalDialog.CliStreamAnimal:input_type -> animal.AnimalRequest
	0, // 3: animal.AnimalDialog.BiAnimal:input_type -> animal.AnimalRequest
	1, // 4: animal.AnimalDialog.UniAnimal:output_type -> animal.AnimalResponse
	1, // 5: animal.AnimalDialog.ServStreamAnimal:output_type -> animal.AnimalResponse
	1, // 6: animal.AnimalDialog.CliStreamAnimal:output_type -> animal.AnimalResponse
	1, // 7: animal.AnimalDialog.BiAnimal:output_type -> animal.AnimalResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_dialog_proto_init() }
func file_message_dialog_proto_init() {
	if File_message_dialog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_dialog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimalRequest); i {
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
		file_message_dialog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimalResponse); i {
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
			RawDescriptor: file_message_dialog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_dialog_proto_goTypes,
		DependencyIndexes: file_message_dialog_proto_depIdxs,
		MessageInfos:      file_message_dialog_proto_msgTypes,
	}.Build()
	File_message_dialog_proto = out.File
	file_message_dialog_proto_rawDesc = nil
	file_message_dialog_proto_goTypes = nil
	file_message_dialog_proto_depIdxs = nil
}
