// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api/v1/speech-to-text.proto

package api

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

type ProcessSpeechRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WakeWord WakeWord `protobuf:"varint,1,opt,name=wakeWord,proto3,enum=api.v1.common.WakeWord" json:"wakeWord,omitempty"`
}

func (x *ProcessSpeechRequest) Reset() {
	*x = ProcessSpeechRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_speech_to_text_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessSpeechRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessSpeechRequest) ProtoMessage() {}

func (x *ProcessSpeechRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_speech_to_text_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessSpeechRequest.ProtoReflect.Descriptor instead.
func (*ProcessSpeechRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_speech_to_text_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessSpeechRequest) GetWakeWord() WakeWord {
	if x != nil {
		return x.WakeWord
	}
	return WakeWord_WAKE_WORD_UNSPECIFIED
}

type ProcessSpeechResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processing      bool   `protobuf:"varint,1,opt,name=processing,proto3" json:"processing,omitempty"`
	TranscribedText string `protobuf:"bytes,2,opt,name=transcribedText,proto3" json:"transcribedText,omitempty"`
}

func (x *ProcessSpeechResponse) Reset() {
	*x = ProcessSpeechResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_speech_to_text_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessSpeechResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessSpeechResponse) ProtoMessage() {}

func (x *ProcessSpeechResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_speech_to_text_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessSpeechResponse.ProtoReflect.Descriptor instead.
func (*ProcessSpeechResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_speech_to_text_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessSpeechResponse) GetProcessing() bool {
	if x != nil {
		return x.Processing
	}
	return false
}

func (x *ProcessSpeechResponse) GetTranscribedText() string {
	if x != nil {
		return x.TranscribedText
	}
	return ""
}

var File_api_v1_speech_to_text_proto protoreflect.FileDescriptor

var file_api_v1_speech_to_text_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2d,
	0x74, 0x6f, 0x2d, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x5f, 0x74, 0x6f, 0x5f,
	0x74, 0x65, 0x78, 0x74, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x14, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x33, 0x0a, 0x08, 0x77, 0x61, 0x6b, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x57, 0x61, 0x6b, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x08, 0x77, 0x61,
	0x6b, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x22, 0x61, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x12,
	0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x54, 0x65,
	0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x32, 0x85, 0x01, 0x0a, 0x13, 0x53, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x54, 0x6f, 0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6e, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x5f, 0x74, 0x6f, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x5f,
	0x74, 0x6f, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x42, 0xdf, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x5f, 0x74, 0x6f, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x42,
	0x11, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x54, 0x6f, 0x54, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x47, 0x6f, 0x42, 0x69, 0x67, 0x38, 0x37, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x67, 0x70,
	0x74, 0x2d, 0x72, 0x61, 0x73, 0x70, 0x65, 0x72, 0x72, 0x79, 0x2d, 0x70, 0x69, 0x2d, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x3b, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x53, 0xaa, 0x02, 0x13,
	0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x54, 0x6f, 0x54,
	0x65, 0x78, 0x74, 0xca, 0x02, 0x13, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x70, 0x65,
	0x65, 0x63, 0x68, 0x54, 0x6f, 0x54, 0x65, 0x78, 0x74, 0xe2, 0x02, 0x1f, 0x41, 0x70, 0x69, 0x5c,
	0x56, 0x31, 0x5c, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x54, 0x6f, 0x54, 0x65, 0x78, 0x74, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x41, 0x70,
	0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x54, 0x6f, 0x54,
	0x65, 0x78, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_speech_to_text_proto_rawDescOnce sync.Once
	file_api_v1_speech_to_text_proto_rawDescData = file_api_v1_speech_to_text_proto_rawDesc
)

func file_api_v1_speech_to_text_proto_rawDescGZIP() []byte {
	file_api_v1_speech_to_text_proto_rawDescOnce.Do(func() {
		file_api_v1_speech_to_text_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_speech_to_text_proto_rawDescData)
	})
	return file_api_v1_speech_to_text_proto_rawDescData
}

var file_api_v1_speech_to_text_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_speech_to_text_proto_goTypes = []interface{}{
	(*ProcessSpeechRequest)(nil),  // 0: api.v1.speech_to_text.ProcessSpeechRequest
	(*ProcessSpeechResponse)(nil), // 1: api.v1.speech_to_text.ProcessSpeechResponse
	(WakeWord)(0),                 // 2: api.v1.common.WakeWord
}
var file_api_v1_speech_to_text_proto_depIdxs = []int32{
	2, // 0: api.v1.speech_to_text.ProcessSpeechRequest.wakeWord:type_name -> api.v1.common.WakeWord
	0, // 1: api.v1.speech_to_text.SpeechToTextService.ProcessSpeech:input_type -> api.v1.speech_to_text.ProcessSpeechRequest
	1, // 2: api.v1.speech_to_text.SpeechToTextService.ProcessSpeech:output_type -> api.v1.speech_to_text.ProcessSpeechResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_speech_to_text_proto_init() }
func file_api_v1_speech_to_text_proto_init() {
	if File_api_v1_speech_to_text_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_speech_to_text_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessSpeechRequest); i {
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
		file_api_v1_speech_to_text_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessSpeechResponse); i {
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
			RawDescriptor: file_api_v1_speech_to_text_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_speech_to_text_proto_goTypes,
		DependencyIndexes: file_api_v1_speech_to_text_proto_depIdxs,
		MessageInfos:      file_api_v1_speech_to_text_proto_msgTypes,
	}.Build()
	File_api_v1_speech_to_text_proto = out.File
	file_api_v1_speech_to_text_proto_rawDesc = nil
	file_api_v1_speech_to_text_proto_goTypes = nil
	file_api_v1_speech_to_text_proto_depIdxs = nil
}