// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helper/storagepacker/types.proto

package storagepacker

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Item represents an entry that gets inserted into the storage packer
type Item struct {
	// ID must be provided by the caller; the same value, if used with GetItem,
	// can be used to fetch the item. However, when iterating through a bucket,
	// this ID will be an internal ID. In other words, outside of the use-case
	// described above, the caller *must not* rely on this value to be
	// consistent with what they passed in.
	ID string `sentinel:"" protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// message is the contents of the item
	Message              *any.Any `sentinel:"" protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0e98c66c4f51b7f, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Item) GetMessage() *any.Any {
	if m != nil {
		return m.Message
	}
	return nil
}

// Bucket is a construct to hold multiple items within itself. This
// abstraction contains multiple buckets of the same kind within itself and
// shares amont them the items that get inserted. When the bucket as a whole
// gets too big to hold more items, the contained buckets gets pushed out only
// to become independent buckets. Hence, this can grow infinitely in terms of
// storage space for items that get inserted.
type Bucket struct {
	// Key is the storage path where the bucket gets stored
	Key string `sentinel:"" protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Items holds the items contained within this bucket. Used by v1.
	Items []*Item `sentinel:"" protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	// ItemMap stores a mapping of item ID to message. Used by v2.
	ItemMap              map[string]*any.Any `sentinel:"" protobuf:"bytes,3,rep,name=item_map,json=itemMap,proto3" json:"item_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0e98c66c4f51b7f, []int{1}
}

func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Bucket) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Bucket) GetItemMap() map[string]*any.Any {
	if m != nil {
		return m.ItemMap
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "storagepacker.Item")
	proto.RegisterType((*Bucket)(nil), "storagepacker.Bucket")
	proto.RegisterMapType((map[string]*any.Any)(nil), "storagepacker.Bucket.ItemMapEntry")
}

func init() { proto.RegisterFile("helper/storagepacker/types.proto", fileDescriptor_c0e98c66c4f51b7f) }

var fileDescriptor_c0e98c66c4f51b7f = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0xeb, 0x36, 0x7d, 0x53, 0x91, 0xe8, 0xa1, 0xee, 0x54, 0x7a, 0xaa, 0x1e, 0x12,
	0x9c, 0x17, 0x11, 0x3c, 0x38, 0x50, 0xf0, 0x20, 0x48, 0x8f, 0x5e, 0x24, 0xed, 0x9e, 0x6d, 0xe8,
	0x8f, 0x84, 0x24, 0x1d, 0xf4, 0x1f, 0xf5, 0xef, 0x91, 0x36, 0x0e, 0x9c, 0x0c, 0x6f, 0x2f, 0x7c,
	0x3f, 0xf9, 0xe4, 0x1b, 0x1e, 0x44, 0x25, 0xd6, 0x0a, 0x35, 0x33, 0x56, 0x6a, 0x5e, 0xa0, 0xe2,
	0x79, 0x85, 0x9a, 0xd9, 0x5e, 0xa1, 0xa1, 0x4a, 0x4b, 0x2b, 0xc9, 0xc9, 0x4e, 0xb4, 0xb8, 0x2c,
	0xa4, 0x2c, 0x6a, 0x64, 0x63, 0x98, 0x75, 0x9f, 0x8c, 0xb7, 0xbd, 0x23, 0xe3, 0x67, 0x38, 0x78,
	0xb1, 0xd8, 0x90, 0x53, 0xf0, 0xc5, 0x3a, 0xf4, 0x22, 0x2f, 0x39, 0x4a, 0x7d, 0xb1, 0x26, 0x14,
	0x66, 0x0d, 0x1a, 0xc3, 0x0b, 0x0c, 0xfd, 0xc8, 0x4b, 0xe6, 0xcb, 0x0b, 0xea, 0x24, 0x74, 0x2b,
	0xa1, 0x8f, 0x6d, 0x9f, 0x6e, 0xa1, 0xf8, 0xcb, 0x83, 0xe9, 0xaa, 0xcb, 0x2b, 0xb4, 0xe4, 0x0c,
	0x82, 0x0a, 0xfb, 0x1f, 0xd7, 0x30, 0x92, 0x2b, 0x98, 0x08, 0x8b, 0x8d, 0x09, 0xfd, 0x28, 0x48,
	0xe6, 0xcb, 0x73, 0xba, 0x53, 0x8f, 0x0e, 0x05, 0x52, 0x47, 0x90, 0x07, 0x38, 0x1c, 0x86, 0x8f,
	0x86, 0xab, 0x30, 0x18, 0xe9, 0xf8, 0x0f, 0xed, 0x5e, 0x19, 0x2f, 0xbd, 0x72, 0xf5, 0xd4, 0x5a,
	0xdd, 0xa7, 0x33, 0xe1, 0x4e, 0x8b, 0x37, 0x38, 0xfe, 0x1d, 0xec, 0xe9, 0x72, 0x0d, 0x93, 0x0d,
	0xaf, 0xbb, 0xff, 0xbf, 0xe5, 0x90, 0x7b, 0xff, 0xce, 0x5b, 0xdd, 0xbc, 0xb3, 0x42, 0xd8, 0xb2,
	0xcb, 0x68, 0x2e, 0x1b, 0x56, 0x72, 0x53, 0x8a, 0x5c, 0x6a, 0xc5, 0x36, 0xbc, 0xab, 0x2d, 0xdb,
	0xb7, 0x89, 0x6c, 0x3a, 0xba, 0x6e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x46, 0x9d, 0x8a, 0xcb,
	0xa8, 0x01, 0x00, 0x00,
}
