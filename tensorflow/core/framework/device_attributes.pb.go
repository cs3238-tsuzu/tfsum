// Code generated by protoc-gen-go.
// source: tensorflow/core/framework/device_attributes.proto
// DO NOT EDIT!

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeviceLocality struct {
	// Optional bus locality of device.  Default value of 0 means
	// no specific locality.  Specific localities are indexed from 1.
	BusId int32 `protobuf:"varint,1,opt,name=bus_id,json=busId" json:"bus_id,omitempty"`
}

func (m *DeviceLocality) Reset()                    { *m = DeviceLocality{} }
func (m *DeviceLocality) String() string            { return proto.CompactTextString(m) }
func (*DeviceLocality) ProtoMessage()               {}
func (*DeviceLocality) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *DeviceLocality) GetBusId() int32 {
	if m != nil {
		return m.BusId
	}
	return 0
}

type DeviceAttributes struct {
	// Fully specified name of the device within a cluster.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// String representation of device_type.
	DeviceType string `protobuf:"bytes,2,opt,name=device_type,json=deviceType" json:"device_type,omitempty"`
	// Memory capacity of device in bytes.
	MemoryLimit int64 `protobuf:"varint,4,opt,name=memory_limit,json=memoryLimit" json:"memory_limit,omitempty"`
	// Platform-specific data about device that may be useful
	// for supporting efficient data transfers.
	Locality *DeviceLocality `protobuf:"bytes,5,opt,name=locality" json:"locality,omitempty"`
	// A device is assigned a global unique number each time it is
	// initialized. "incarnation" should never be 0.
	Incarnation uint64 `protobuf:"fixed64,6,opt,name=incarnation" json:"incarnation,omitempty"`
	// String representation of the physical device that this device maps to.
	PhysicalDeviceDesc string `protobuf:"bytes,7,opt,name=physical_device_desc,json=physicalDeviceDesc" json:"physical_device_desc,omitempty"`
}

func (m *DeviceAttributes) Reset()                    { *m = DeviceAttributes{} }
func (m *DeviceAttributes) String() string            { return proto.CompactTextString(m) }
func (*DeviceAttributes) ProtoMessage()               {}
func (*DeviceAttributes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *DeviceAttributes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceAttributes) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *DeviceAttributes) GetMemoryLimit() int64 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *DeviceAttributes) GetLocality() *DeviceLocality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *DeviceAttributes) GetIncarnation() uint64 {
	if m != nil {
		return m.Incarnation
	}
	return 0
}

func (m *DeviceAttributes) GetPhysicalDeviceDesc() string {
	if m != nil {
		return m.PhysicalDeviceDesc
	}
	return ""
}

func init() {
	proto.RegisterType((*DeviceLocality)(nil), "tensorflow.DeviceLocality")
	proto.RegisterType((*DeviceAttributes)(nil), "tensorflow.DeviceAttributes")
}

func init() { proto.RegisterFile("tensorflow/core/framework/device_attributes.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x14, 0xc4, 0x89, 0x6e, 0xab, 0xbe, 0x8a, 0x48, 0x50, 0x09, 0x5e, 0xac, 0x7b, 0xb1, 0xa7, 0xd6,
	0x3f, 0xa0, 0x67, 0x97, 0xbd, 0x08, 0x7b, 0x58, 0x8a, 0xf7, 0x92, 0xa6, 0x6f, 0x35, 0xd8, 0x36,
	0x25, 0x49, 0x5d, 0xfa, 0xc5, 0xc5, 0xa3, 0x6c, 0x5a, 0xbb, 0xab, 0xb7, 0x30, 0x33, 0x79, 0xf3,
	0x63, 0xe0, 0xce, 0x62, 0x6d, 0x94, 0x5e, 0x95, 0x6a, 0x9d, 0x08, 0xa5, 0x31, 0x59, 0x69, 0x5e,
	0xe1, 0x5a, 0xe9, 0x8f, 0xa4, 0xc0, 0x4f, 0x29, 0x30, 0xe3, 0xd6, 0x6a, 0x99, 0xb7, 0x16, 0x4d,
	0xdc, 0x68, 0x65, 0x15, 0x85, 0xed, 0x97, 0xe9, 0x0d, 0x9c, 0xcc, 0x5d, 0x6c, 0xa1, 0x04, 0x2f,
	0xa5, 0xed, 0xe8, 0x39, 0xf8, 0x79, 0x6b, 0x32, 0x59, 0x30, 0x12, 0x92, 0xc8, 0x4b, 0xbd, 0xbc,
	0x35, 0x2f, 0xc5, 0xf4, 0x8b, 0xc0, 0x69, 0x9f, 0x7c, 0x1e, 0xef, 0x51, 0x0a, 0x93, 0x9a, 0x57,
	0xe8, 0x92, 0x47, 0xa9, 0x7b, 0xd3, 0x2b, 0x08, 0x86, 0x62, 0xdb, 0x35, 0xc8, 0xf6, 0x9c, 0x05,
	0xbd, 0xf4, 0xda, 0x35, 0x48, 0xaf, 0xe1, 0xb8, 0xc2, 0x4a, 0xe9, 0x2e, 0x2b, 0x65, 0x25, 0x2d,
	0x9b, 0x84, 0x24, 0xda, 0x4f, 0x83, 0x5e, 0x5b, 0x6c, 0x24, 0xfa, 0x08, 0x87, 0xe5, 0xc0, 0xc3,
	0xbc, 0x90, 0x44, 0xc1, 0xfd, 0x65, 0xbc, 0x85, 0x8e, 0xff, 0x12, 0xa7, 0x63, 0x96, 0x86, 0x10,
	0xc8, 0x5a, 0x70, 0x5d, 0x73, 0x2b, 0x55, 0xcd, 0xfc, 0x90, 0x44, 0x7e, 0xba, 0x2b, 0xd1, 0x5b,
	0x38, 0x6b, 0xde, 0x3b, 0x23, 0x05, 0x2f, 0xb3, 0x01, 0xb3, 0x40, 0x23, 0xd8, 0x81, 0xc3, 0xa4,
	0xbf, 0x5e, 0xdf, 0x30, 0x47, 0x23, 0x66, 0x4f, 0xc0, 0x94, 0x7e, 0xdb, 0xad, 0x1f, 0x17, 0x9e,
	0x5d, 0xfc, 0x5f, 0x64, 0xb9, 0x19, 0xd8, 0x2c, 0xc9, 0x37, 0x21, 0xb9, 0xef, 0xd6, 0x7e, 0xf8,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xb1, 0x4d, 0x88, 0xa2, 0x01, 0x00, 0x00,
}
