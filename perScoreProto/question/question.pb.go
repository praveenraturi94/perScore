// Code generated by protoc-gen-go. DO NOT EDIT.
// source: question.proto

/*
Package question is a generated protocol buffer package.

It is generated from these files:
	question.proto

It has these top-level messages:
	CreateQuestionRequest
	CreateQuestionResponse
	GetQuestionRequest
	GetQuestionResponse
*/
package question

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateQuestionRequest struct {
	AuthToken  string                            `protobuf:"bytes,1,opt,name=auth_token,json=authToken" json:"auth_token,omitempty"`
	Title      string                            `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Body       string                            `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	Answer     *CreateQuestionRequest_Answer     `protobuf:"bytes,4,opt,name=answer" json:"answer,omitempty"`
	Weight     *CreateQuestionRequest_Weight     `protobuf:"bytes,5,opt,name=weight" json:"weight,omitempty"`
	Categories []*CreateQuestionRequest_Category `protobuf:"bytes,6,rep,name=categories" json:"categories,omitempty"`
}

func (m *CreateQuestionRequest) Reset()                    { *m = CreateQuestionRequest{} }
func (m *CreateQuestionRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateQuestionRequest) ProtoMessage()               {}
func (*CreateQuestionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateQuestionRequest) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func (m *CreateQuestionRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateQuestionRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *CreateQuestionRequest) GetAnswer() *CreateQuestionRequest_Answer {
	if m != nil {
		return m.Answer
	}
	return nil
}

func (m *CreateQuestionRequest) GetWeight() *CreateQuestionRequest_Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *CreateQuestionRequest) GetCategories() []*CreateQuestionRequest_Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type CreateQuestionRequest_Answer struct {
	Option1    string                                   `protobuf:"bytes,1,opt,name=option1" json:"option1,omitempty"`
	Option2    string                                   `protobuf:"bytes,2,opt,name=option2" json:"option2,omitempty"`
	Option3    string                                   `protobuf:"bytes,3,opt,name=option3" json:"option3,omitempty"`
	Option4    string                                   `protobuf:"bytes,4,opt,name=option4" json:"option4,omitempty"`
	Option5    string                                   `protobuf:"bytes,5,opt,name=option5" json:"option5,omitempty"`
	Weights    []*CreateQuestionRequest_Answer_Weight   `protobuf:"bytes,6,rep,name=weights" json:"weights,omitempty"`
	Categories []*CreateQuestionRequest_Answer_Category `protobuf:"bytes,7,rep,name=categories" json:"categories,omitempty"`
}

func (m *CreateQuestionRequest_Answer) Reset()                    { *m = CreateQuestionRequest_Answer{} }
func (m *CreateQuestionRequest_Answer) String() string            { return proto.CompactTextString(m) }
func (*CreateQuestionRequest_Answer) ProtoMessage()               {}
func (*CreateQuestionRequest_Answer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *CreateQuestionRequest_Answer) GetOption1() string {
	if m != nil {
		return m.Option1
	}
	return ""
}

func (m *CreateQuestionRequest_Answer) GetOption2() string {
	if m != nil {
		return m.Option2
	}
	return ""
}

func (m *CreateQuestionRequest_Answer) GetOption3() string {
	if m != nil {
		return m.Option3
	}
	return ""
}

func (m *CreateQuestionRequest_Answer) GetOption4() string {
	if m != nil {
		return m.Option4
	}
	return ""
}

func (m *CreateQuestionRequest_Answer) GetOption5() string {
	if m != nil {
		return m.Option5
	}
	return ""
}

func (m *CreateQuestionRequest_Answer) GetWeights() []*CreateQuestionRequest_Answer_Weight {
	if m != nil {
		return m.Weights
	}
	return nil
}

func (m *CreateQuestionRequest_Answer) GetCategories() []*CreateQuestionRequest_Answer_Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type CreateQuestionRequest_Answer_Category struct {
	Id         int32                                    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name       string                                   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Option     int32                                    `protobuf:"varint,3,opt,name=option" json:"option,omitempty"`
	Parent     int32                                    `protobuf:"varint,4,opt,name=parent" json:"parent,omitempty"`
	Categories []*CreateQuestionRequest_Answer_Category `protobuf:"bytes,5,rep,name=categories" json:"categories,omitempty"`
}

func (m *CreateQuestionRequest_Answer_Category) Reset()         { *m = CreateQuestionRequest_Answer_Category{} }
func (m *CreateQuestionRequest_Answer_Category) String() string { return proto.CompactTextString(m) }
func (*CreateQuestionRequest_Answer_Category) ProtoMessage()    {}
func (*CreateQuestionRequest_Answer_Category) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

func (m *CreateQuestionRequest_Answer_Category) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateQuestionRequest_Answer_Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateQuestionRequest_Answer_Category) GetOption() int32 {
	if m != nil {
		return m.Option
	}
	return 0
}

func (m *CreateQuestionRequest_Answer_Category) GetParent() int32 {
	if m != nil {
		return m.Parent
	}
	return 0
}

func (m *CreateQuestionRequest_Answer_Category) GetCategories() []*CreateQuestionRequest_Answer_Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type CreateQuestionRequest_Answer_Weight struct {
	Value  int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Option int32 `protobuf:"varint,2,opt,name=option" json:"option,omitempty"`
}

func (m *CreateQuestionRequest_Answer_Weight) Reset()         { *m = CreateQuestionRequest_Answer_Weight{} }
func (m *CreateQuestionRequest_Answer_Weight) String() string { return proto.CompactTextString(m) }
func (*CreateQuestionRequest_Answer_Weight) ProtoMessage()    {}
func (*CreateQuestionRequest_Answer_Weight) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 1}
}

func (m *CreateQuestionRequest_Answer_Weight) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *CreateQuestionRequest_Answer_Weight) GetOption() int32 {
	if m != nil {
		return m.Option
	}
	return 0
}

type CreateQuestionRequest_Category struct {
	Id         int32                             `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name       string                            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Parent     int32                             `protobuf:"varint,3,opt,name=parent" json:"parent,omitempty"`
	Categories []*CreateQuestionRequest_Category `protobuf:"bytes,4,rep,name=categories" json:"categories,omitempty"`
}

func (m *CreateQuestionRequest_Category) Reset()         { *m = CreateQuestionRequest_Category{} }
func (m *CreateQuestionRequest_Category) String() string { return proto.CompactTextString(m) }
func (*CreateQuestionRequest_Category) ProtoMessage()    {}
func (*CreateQuestionRequest_Category) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1}
}

func (m *CreateQuestionRequest_Category) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateQuestionRequest_Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateQuestionRequest_Category) GetParent() int32 {
	if m != nil {
		return m.Parent
	}
	return 0
}

func (m *CreateQuestionRequest_Category) GetCategories() []*CreateQuestionRequest_Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type CreateQuestionRequest_Weight struct {
	Value int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *CreateQuestionRequest_Weight) Reset()                    { *m = CreateQuestionRequest_Weight{} }
func (m *CreateQuestionRequest_Weight) String() string            { return proto.CompactTextString(m) }
func (*CreateQuestionRequest_Weight) ProtoMessage()               {}
func (*CreateQuestionRequest_Weight) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

func (m *CreateQuestionRequest_Weight) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type CreateQuestionResponse struct {
	Success    bool                               `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message    string                             `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Categories []*CreateQuestionResponse_Category `protobuf:"bytes,3,rep,name=categories" json:"categories,omitempty"`
}

func (m *CreateQuestionResponse) Reset()                    { *m = CreateQuestionResponse{} }
func (m *CreateQuestionResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateQuestionResponse) ProtoMessage()               {}
func (*CreateQuestionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateQuestionResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CreateQuestionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateQuestionResponse) GetCategories() []*CreateQuestionResponse_Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type CreateQuestionResponse_Category struct {
	Id          int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	WeightRange string `protobuf:"bytes,3,opt,name=weight_range,json=weightRange" json:"weight_range,omitempty"`
	Parent      int32  `protobuf:"varint,4,opt,name=parent" json:"parent,omitempty"`
	Level       int32  `protobuf:"varint,5,opt,name=level" json:"level,omitempty"`
}

func (m *CreateQuestionResponse_Category) Reset()         { *m = CreateQuestionResponse_Category{} }
func (m *CreateQuestionResponse_Category) String() string { return proto.CompactTextString(m) }
func (*CreateQuestionResponse_Category) ProtoMessage()    {}
func (*CreateQuestionResponse_Category) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

func (m *CreateQuestionResponse_Category) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateQuestionResponse_Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateQuestionResponse_Category) GetWeightRange() string {
	if m != nil {
		return m.WeightRange
	}
	return ""
}

func (m *CreateQuestionResponse_Category) GetParent() int32 {
	if m != nil {
		return m.Parent
	}
	return 0
}

func (m *CreateQuestionResponse_Category) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type GetQuestionRequest struct {
	AuthToken  string                     `protobuf:"bytes,1,opt,name=auth_token,json=authToken" json:"auth_token,omitempty"`
	QuestionId int32                      `protobuf:"varint,2,opt,name=question_id,json=questionId" json:"question_id,omitempty"`
	Answer     *GetQuestionRequest_Answer `protobuf:"bytes,3,opt,name=answer" json:"answer,omitempty"`
}

func (m *GetQuestionRequest) Reset()                    { *m = GetQuestionRequest{} }
func (m *GetQuestionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetQuestionRequest) ProtoMessage()               {}
func (*GetQuestionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetQuestionRequest) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func (m *GetQuestionRequest) GetQuestionId() int32 {
	if m != nil {
		return m.QuestionId
	}
	return 0
}

func (m *GetQuestionRequest) GetAnswer() *GetQuestionRequest_Answer {
	if m != nil {
		return m.Answer
	}
	return nil
}

type GetQuestionRequest_Answer struct {
	Option1 bool `protobuf:"varint,1,opt,name=option1" json:"option1,omitempty"`
	Option2 bool `protobuf:"varint,2,opt,name=option2" json:"option2,omitempty"`
	Option3 bool `protobuf:"varint,3,opt,name=option3" json:"option3,omitempty"`
	Option4 bool `protobuf:"varint,4,opt,name=option4" json:"option4,omitempty"`
	Option5 bool `protobuf:"varint,5,opt,name=option5" json:"option5,omitempty"`
}

func (m *GetQuestionRequest_Answer) Reset()                    { *m = GetQuestionRequest_Answer{} }
func (m *GetQuestionRequest_Answer) String() string            { return proto.CompactTextString(m) }
func (*GetQuestionRequest_Answer) ProtoMessage()               {}
func (*GetQuestionRequest_Answer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *GetQuestionRequest_Answer) GetOption1() bool {
	if m != nil {
		return m.Option1
	}
	return false
}

func (m *GetQuestionRequest_Answer) GetOption2() bool {
	if m != nil {
		return m.Option2
	}
	return false
}

func (m *GetQuestionRequest_Answer) GetOption3() bool {
	if m != nil {
		return m.Option3
	}
	return false
}

func (m *GetQuestionRequest_Answer) GetOption4() bool {
	if m != nil {
		return m.Option4
	}
	return false
}

func (m *GetQuestionRequest_Answer) GetOption5() bool {
	if m != nil {
		return m.Option5
	}
	return false
}

type GetQuestionResponse struct {
	Success bool                        `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string                      `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Title   string                      `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Body    string                      `protobuf:"bytes,4,opt,name=body" json:"body,omitempty"`
	Score   float32                     `protobuf:"fixed32,5,opt,name=score" json:"score,omitempty"`
	Answer  *GetQuestionResponse_Answer `protobuf:"bytes,6,opt,name=answer" json:"answer,omitempty"`
}

func (m *GetQuestionResponse) Reset()                    { *m = GetQuestionResponse{} }
func (m *GetQuestionResponse) String() string            { return proto.CompactTextString(m) }
func (*GetQuestionResponse) ProtoMessage()               {}
func (*GetQuestionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetQuestionResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetQuestionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetQuestionResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GetQuestionResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *GetQuestionResponse) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *GetQuestionResponse) GetAnswer() *GetQuestionResponse_Answer {
	if m != nil {
		return m.Answer
	}
	return nil
}

type GetQuestionResponse_Answer struct {
	Option1 string `protobuf:"bytes,1,opt,name=option1" json:"option1,omitempty"`
	Option2 string `protobuf:"bytes,2,opt,name=option2" json:"option2,omitempty"`
	Option3 string `protobuf:"bytes,3,opt,name=option3" json:"option3,omitempty"`
	Option4 string `protobuf:"bytes,4,opt,name=option4" json:"option4,omitempty"`
	Option5 string `protobuf:"bytes,5,opt,name=option5" json:"option5,omitempty"`
}

func (m *GetQuestionResponse_Answer) Reset()                    { *m = GetQuestionResponse_Answer{} }
func (m *GetQuestionResponse_Answer) String() string            { return proto.CompactTextString(m) }
func (*GetQuestionResponse_Answer) ProtoMessage()               {}
func (*GetQuestionResponse_Answer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *GetQuestionResponse_Answer) GetOption1() string {
	if m != nil {
		return m.Option1
	}
	return ""
}

func (m *GetQuestionResponse_Answer) GetOption2() string {
	if m != nil {
		return m.Option2
	}
	return ""
}

func (m *GetQuestionResponse_Answer) GetOption3() string {
	if m != nil {
		return m.Option3
	}
	return ""
}

func (m *GetQuestionResponse_Answer) GetOption4() string {
	if m != nil {
		return m.Option4
	}
	return ""
}

func (m *GetQuestionResponse_Answer) GetOption5() string {
	if m != nil {
		return m.Option5
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateQuestionRequest)(nil), "question.CreateQuestionRequest")
	proto.RegisterType((*CreateQuestionRequest_Answer)(nil), "question.CreateQuestionRequest.Answer")
	proto.RegisterType((*CreateQuestionRequest_Answer_Category)(nil), "question.CreateQuestionRequest.Answer.Category")
	proto.RegisterType((*CreateQuestionRequest_Answer_Weight)(nil), "question.CreateQuestionRequest.Answer.Weight")
	proto.RegisterType((*CreateQuestionRequest_Category)(nil), "question.CreateQuestionRequest.Category")
	proto.RegisterType((*CreateQuestionRequest_Weight)(nil), "question.CreateQuestionRequest.Weight")
	proto.RegisterType((*CreateQuestionResponse)(nil), "question.CreateQuestionResponse")
	proto.RegisterType((*CreateQuestionResponse_Category)(nil), "question.CreateQuestionResponse.Category")
	proto.RegisterType((*GetQuestionRequest)(nil), "question.GetQuestionRequest")
	proto.RegisterType((*GetQuestionRequest_Answer)(nil), "question.GetQuestionRequest.Answer")
	proto.RegisterType((*GetQuestionResponse)(nil), "question.GetQuestionResponse")
	proto.RegisterType((*GetQuestionResponse_Answer)(nil), "question.GetQuestionResponse.Answer")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Question service

type QuestionClient interface {
	// Create a new question
	CreateQuestion(ctx context.Context, in *CreateQuestionRequest, opts ...grpc.CallOption) (*CreateQuestionResponse, error)
	// Get next question
	GetQuestion(ctx context.Context, in *GetQuestionRequest, opts ...grpc.CallOption) (*GetQuestionResponse, error)
}

type questionClient struct {
	cc *grpc.ClientConn
}

func NewQuestionClient(cc *grpc.ClientConn) QuestionClient {
	return &questionClient{cc}
}

func (c *questionClient) CreateQuestion(ctx context.Context, in *CreateQuestionRequest, opts ...grpc.CallOption) (*CreateQuestionResponse, error) {
	out := new(CreateQuestionResponse)
	err := grpc.Invoke(ctx, "/question.Question/CreateQuestion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionClient) GetQuestion(ctx context.Context, in *GetQuestionRequest, opts ...grpc.CallOption) (*GetQuestionResponse, error) {
	out := new(GetQuestionResponse)
	err := grpc.Invoke(ctx, "/question.Question/GetQuestion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Question service

type QuestionServer interface {
	// Create a new question
	CreateQuestion(context.Context, *CreateQuestionRequest) (*CreateQuestionResponse, error)
	// Get next question
	GetQuestion(context.Context, *GetQuestionRequest) (*GetQuestionResponse, error)
}

func RegisterQuestionServer(s *grpc.Server, srv QuestionServer) {
	s.RegisterService(&_Question_serviceDesc, srv)
}

func _Question_CreateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServer).CreateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/question.Question/CreateQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServer).CreateQuestion(ctx, req.(*CreateQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Question_GetQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServer).GetQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/question.Question/GetQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServer).GetQuestion(ctx, req.(*GetQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Question_serviceDesc = grpc.ServiceDesc{
	ServiceName: "question.Question",
	HandlerType: (*QuestionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQuestion",
			Handler:    _Question_CreateQuestion_Handler,
		},
		{
			MethodName: "GetQuestion",
			Handler:    _Question_GetQuestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "question.proto",
}

func init() { proto.RegisterFile("question.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x6d, 0xec, 0xd8, 0x71, 0x27, 0xa8, 0x87, 0x25, 0x54, 0x56, 0x44, 0x69, 0x30, 0x08, 0x85,
	0x03, 0x41, 0xa4, 0x0d, 0x17, 0x10, 0x12, 0xea, 0xa1, 0x54, 0x42, 0x42, 0xac, 0x40, 0x1c, 0x23,
	0x37, 0x59, 0xa5, 0x16, 0xa9, 0x1d, 0xbc, 0x9b, 0x56, 0x3d, 0x71, 0xe7, 0x80, 0x38, 0x72, 0xe2,
	0x17, 0xf8, 0x19, 0xbe, 0x83, 0x6f, 0x40, 0x3b, 0xbb, 0xeb, 0xd8, 0x21, 0x4e, 0x52, 0x22, 0x71,
	0xcb, 0x9b, 0xf1, 0xac, 0xdf, 0x3c, 0xbd, 0x7d, 0x0e, 0xec, 0x7c, 0x9a, 0x32, 0x2e, 0xa2, 0x24,
	0xee, 0x4c, 0xd2, 0x44, 0x24, 0xc4, 0x33, 0x38, 0xf8, 0x5d, 0x83, 0x5b, 0x47, 0x29, 0x0b, 0x05,
	0x7b, 0xab, 0x4b, 0x94, 0x61, 0x93, 0xec, 0x01, 0x84, 0x53, 0x71, 0xd6, 0x17, 0xc9, 0x47, 0x16,
	0xfb, 0x95, 0x56, 0xa5, 0xbd, 0x4d, 0xb7, 0x65, 0xe5, 0x9d, 0x2c, 0x90, 0x06, 0x38, 0x22, 0x12,
	0x63, 0xe6, 0x5b, 0xd8, 0x51, 0x80, 0x10, 0xa8, 0x9e, 0x26, 0xc3, 0x2b, 0xdf, 0xc6, 0x22, 0xfe,
	0x26, 0x2f, 0xc0, 0x0d, 0x63, 0x7e, 0xc9, 0x52, 0xbf, 0xda, 0xaa, 0xb4, 0xeb, 0xdd, 0x07, 0x9d,
	0x8c, 0xcd, 0xc2, 0x37, 0x77, 0x5e, 0xe2, 0xd3, 0x54, 0x4f, 0xc9, 0xf9, 0x4b, 0x16, 0x8d, 0xce,
	0x84, 0xef, 0xac, 0x37, 0xff, 0x01, 0x9f, 0xa6, 0x7a, 0x8a, 0xbc, 0x02, 0x18, 0x84, 0x82, 0x8d,
	0x92, 0x34, 0x62, 0xdc, 0x77, 0x5b, 0x76, 0xbb, 0xde, 0x6d, 0xaf, 0x3a, 0xe3, 0x48, 0x4d, 0x5c,
	0xd1, 0xdc, 0x6c, 0xf3, 0x6b, 0x15, 0x5c, 0x45, 0x8e, 0xf8, 0x50, 0x4b, 0x26, 0x72, 0xe0, 0x89,
	0x96, 0xc6, 0xc0, 0x59, 0xa7, 0xab, 0xa5, 0x31, 0x70, 0xd6, 0x39, 0xd0, 0xfa, 0x18, 0x38, 0xeb,
	0x1c, 0xa2, 0x46, 0x59, 0xe7, 0x70, 0xd6, 0xe9, 0xe1, 0xf6, 0x59, 0xa7, 0x47, 0x8e, 0xa1, 0xa6,
	0x16, 0x34, 0x3b, 0x3d, 0x5a, 0x4f, 0x57, 0x23, 0x8f, 0x99, 0x26, 0x6f, 0x0a, 0xfa, 0xd4, 0xf0,
	0xac, 0xc7, 0x6b, 0x9e, 0xb5, 0x50, 0xa6, 0x9f, 0x15, 0xf0, 0x4c, 0x83, 0xec, 0x80, 0x15, 0x0d,
	0x51, 0x23, 0x87, 0x5a, 0xd1, 0x50, 0x3a, 0x24, 0x0e, 0xcf, 0x8d, 0x6d, 0xf0, 0x37, 0xd9, 0x05,
	0x57, 0x6d, 0x85, 0xba, 0x38, 0x54, 0x23, 0x59, 0x9f, 0x84, 0x29, 0x8b, 0x05, 0xaa, 0xe2, 0x50,
	0x8d, 0xe6, 0x18, 0x3b, 0x9b, 0x33, 0x7e, 0x0a, 0xae, 0x52, 0x45, 0xda, 0xfa, 0x22, 0x1c, 0x4f,
	0x99, 0x66, 0xac, 0x40, 0x8e, 0xa0, 0x95, 0x27, 0xd8, 0xfc, 0xf6, 0x0f, 0x9b, 0xea, 0x8d, 0xec,
	0xc2, 0x46, 0x45, 0x8f, 0x56, 0x37, 0xf0, 0xe8, 0x9d, 0xe5, 0xab, 0x04, 0xdf, 0x2d, 0xd8, 0x9d,
	0x3f, 0x8e, 0x4f, 0x92, 0x98, 0x33, 0xe9, 0x35, 0x3e, 0x1d, 0x0c, 0x18, 0xe7, 0x38, 0xe2, 0x51,
	0x03, 0x65, 0xe7, 0x9c, 0x71, 0x1e, 0x8e, 0xcc, 0x36, 0x06, 0x92, 0x93, 0x02, 0x71, 0x1b, 0x89,
	0x3f, 0x2c, 0x27, 0xae, 0xde, 0xb4, 0x98, 0xf9, 0xe7, 0x6b, 0x6a, 0x79, 0x17, 0x6e, 0x28, 0x0b,
	0xf7, 0xd3, 0x30, 0x1e, 0x31, 0x7d, 0xa7, 0xea, 0xaa, 0x46, 0x65, 0xa9, 0xd4, 0x40, 0x0d, 0x70,
	0xc6, 0xec, 0x82, 0x8d, 0xf1, 0x4e, 0x39, 0x54, 0x81, 0xe0, 0x87, 0x05, 0xe4, 0x98, 0x89, 0x6b,
	0x06, 0xe1, 0x3e, 0xd4, 0xcd, 0xba, 0xfd, 0x68, 0xa8, 0x0d, 0x02, 0xa6, 0x74, 0x32, 0x24, 0xcf,
	0xb2, 0xfc, 0xb3, 0x31, 0xbf, 0xee, 0xcd, 0xe4, 0xf9, 0xfb, 0x6d, 0x73, 0xe1, 0xd7, 0xfc, 0x52,
	0x29, 0x8b, 0x1c, 0xaf, 0x34, 0x72, 0xbc, 0xd2, 0xc8, 0xf1, 0x4a, 0x23, 0xc7, 0x2b, 0x8d, 0x9c,
	0xac, 0xd3, 0x0b, 0x7e, 0x59, 0x70, 0xb3, 0x40, 0x79, 0x03, 0xe3, 0x64, 0xdf, 0x0f, 0x7b, 0xd1,
	0xf7, 0xa3, 0x9a, 0xfb, 0x7e, 0x34, 0xc0, 0xe1, 0x83, 0x24, 0x65, 0xc8, 0xc6, 0xa2, 0x0a, 0x90,
	0xe7, 0x99, 0xaa, 0x2e, 0xaa, 0x7a, 0xbf, 0x44, 0x55, 0xed, 0xb8, 0xb5, 0x65, 0xfd, 0xff, 0x49,
	0xde, 0x95, 0x79, 0x69, 0x08, 0x93, 0xf7, 0xb0, 0x53, 0xbc, 0x34, 0x64, 0x7f, 0x45, 0x0e, 0x34,
	0x5b, 0xab, 0xee, 0x5b, 0xb0, 0x45, 0x5e, 0x43, 0x3d, 0x27, 0x0b, 0xb9, 0xbd, 0xcc, 0x83, 0xcd,
	0xbd, 0xa5, 0x5a, 0x06, 0x5b, 0xa7, 0x2e, 0xfe, 0x8d, 0x38, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x70, 0xce, 0x74, 0x18, 0x58, 0x08, 0x00, 0x00,
}
