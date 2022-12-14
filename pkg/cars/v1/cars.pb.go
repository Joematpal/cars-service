// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: cars.proto

package v1_cars

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

type ListCarsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query  *CarReq `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Cursor *Cursor `protobuf:"bytes,2,opt,name=Cursor,proto3" json:"Cursor,omitempty"`
}

func (x *ListCarsReq) Reset() {
	*x = ListCarsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cars_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCarsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCarsReq) ProtoMessage() {}

func (x *ListCarsReq) ProtoReflect() protoreflect.Message {
	mi := &file_cars_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCarsReq.ProtoReflect.Descriptor instead.
func (*ListCarsReq) Descriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{0}
}

func (x *ListCarsReq) GetQuery() *CarReq {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ListCarsReq) GetCursor() *Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Make     string `protobuf:"bytes,2,opt,name=make,proto3" json:"make,omitempty"`
	Package  string `protobuf:"bytes,3,opt,name=package,proto3" json:"package,omitempty"`
	Color    string `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	Year     int32  `protobuf:"varint,5,opt,name=year,proto3" json:"year,omitempty"`
	Mileage  int64  `protobuf:"varint,6,opt,name=mileage,proto3" json:"mileage,omitempty"`
	Price    int64  `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	Category string `protobuf:"bytes,8,opt,name=category,proto3" json:"category,omitempty"`
	Model    string `protobuf:"bytes,9,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cars_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_cars_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{1}
}

func (x *Car) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Car) GetMake() string {
	if x != nil {
		return x.Make
	}
	return ""
}

func (x *Car) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *Car) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Car) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Car) GetMileage() int64 {
	if x != nil {
		return x.Mileage
	}
	return 0
}

func (x *Car) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Car) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Car) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

type CarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CarReq) Reset() {
	*x = CarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cars_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarReq) ProtoMessage() {}

func (x *CarReq) ProtoReflect() protoreflect.Message {
	mi := &file_cars_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarReq.ProtoReflect.Descriptor instead.
func (*CarReq) Descriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{2}
}

func (x *CarReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cars_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_cars_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{3}
}

func (x *Success) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Success) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Cars struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cars []*Car `protobuf:"bytes,1,rep,name=cars,proto3" json:"cars,omitempty"`
	Page *Page  `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *Cars) Reset() {
	*x = Cars{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cars_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cars) ProtoMessage() {}

func (x *Cars) ProtoReflect() protoreflect.Message {
	mi := &file_cars_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cars.ProtoReflect.Descriptor instead.
func (*Cars) Descriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{4}
}

func (x *Cars) GetCars() []*Car {
	if x != nil {
		return x.Cars
	}
	return nil
}

func (x *Cars) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type Cursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size   int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Cursor) Reset() {
	*x = Cursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cars_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cursor) ProtoMessage() {}

func (x *Cursor) ProtoReflect() protoreflect.Message {
	mi := &file_cars_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cursor.ProtoReflect.Descriptor instead.
func (*Cursor) Descriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{5}
}

func (x *Cursor) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Cursor) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Response pointing to the next set of pages
type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// update for how many returned
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Returns next offset
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Total number of rows
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cars_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_cars_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{6}
}

func (x *Page) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Page) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Page) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_cars_proto protoreflect.FileDescriptor

var file_cars_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61,
	0x72, 0x73, 0x22, 0x57, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x22, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x52, 0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0xcf, 0x01, 0x0a, 0x03,
	0x43, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x61, 0x6b, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x69,
	0x6c, 0x65, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x18, 0x0a,
	0x06, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x45, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x04,
	0x63, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x63, 0x61, 0x72,
	0x73, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x04, 0x63, 0x61, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x61, 0x72, 0x73,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x34, 0x0a, 0x06, 0x43,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x22, 0x48, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xa8, 0x01, 0x0a, 0x0b,
	0x43, 0x61, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x12, 0x09, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e,
	0x43, 0x61, 0x72, 0x1a, 0x09, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x12, 0x29,
	0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x72,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e,
	0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x12, 0x0c, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x09, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x12, 0x28, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x12, 0x0c, 0x2e, 0x63, 0x61, 0x72, 0x73,
	0x2e, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x65, 0x6d, 0x61, 0x74, 0x70, 0x61, 0x6c, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x63, 0x61, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3a, 0x63, 0x61, 0x72, 0x73, 0x3b,
	0x76, 0x31, 0x5f, 0x63, 0x61, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cars_proto_rawDescOnce sync.Once
	file_cars_proto_rawDescData = file_cars_proto_rawDesc
)

func file_cars_proto_rawDescGZIP() []byte {
	file_cars_proto_rawDescOnce.Do(func() {
		file_cars_proto_rawDescData = protoimpl.X.CompressGZIP(file_cars_proto_rawDescData)
	})
	return file_cars_proto_rawDescData
}

var file_cars_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cars_proto_goTypes = []interface{}{
	(*ListCarsReq)(nil), // 0: cars.ListCarsReq
	(*Car)(nil),         // 1: cars.Car
	(*CarReq)(nil),      // 2: cars.CarReq
	(*Success)(nil),     // 3: cars.Success
	(*Cars)(nil),        // 4: cars.Cars
	(*Cursor)(nil),      // 5: cars.Cursor
	(*Page)(nil),        // 6: cars.Page
}
var file_cars_proto_depIdxs = []int32{
	2, // 0: cars.ListCarsReq.query:type_name -> cars.CarReq
	5, // 1: cars.ListCarsReq.Cursor:type_name -> cars.Cursor
	1, // 2: cars.Cars.cars:type_name -> cars.Car
	6, // 3: cars.Cars.page:type_name -> cars.Page
	1, // 4: cars.CarsService.CreateCar:input_type -> cars.Car
	0, // 5: cars.CarsService.ListCars:input_type -> cars.ListCarsReq
	2, // 6: cars.CarsService.GetCar:input_type -> cars.CarReq
	2, // 7: cars.CarsService.DeleteCar:input_type -> cars.CarReq
	1, // 8: cars.CarsService.CreateCar:output_type -> cars.Car
	4, // 9: cars.CarsService.ListCars:output_type -> cars.Cars
	1, // 10: cars.CarsService.GetCar:output_type -> cars.Car
	3, // 11: cars.CarsService.DeleteCar:output_type -> cars.Success
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cars_proto_init() }
func file_cars_proto_init() {
	if File_cars_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cars_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCarsReq); i {
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
		file_cars_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_cars_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarReq); i {
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
		file_cars_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
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
		file_cars_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cars); i {
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
		file_cars_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cursor); i {
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
		file_cars_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
			RawDescriptor: file_cars_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cars_proto_goTypes,
		DependencyIndexes: file_cars_proto_depIdxs,
		MessageInfos:      file_cars_proto_msgTypes,
	}.Build()
	File_cars_proto = out.File
	file_cars_proto_rawDesc = nil
	file_cars_proto_goTypes = nil
	file_cars_proto_depIdxs = nil
}
