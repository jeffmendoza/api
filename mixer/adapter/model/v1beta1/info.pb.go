// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/model/v1beta1/info.proto

package v1beta1 // import "istio.io/api/mixer/adapter/model/v1beta1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Info describes an adapter or a backend that wants to provide telemetry and policy functionality to Mixer as an
// out of process adapter.
type Info struct {
	// Name of the adapter. It must be an RFC 1035 compatible DNS label
	// matching the `^[a-z]([-a-z0-9]*[a-z0-9])?$` regular expression.
	// Name is used in Istio configuration, therefore it should be descriptive but short.
	// example: denier
	// Vendor adapters should use a vendor prefix.
	// example: mycompany-denier
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// User-friendly description of the adapter.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Names of the templates the adapter supports.
	Templates []string `protobuf:"bytes,3,rep,name=templates" json:"templates,omitempty"`
	// Base64 encoded proto descriptor of the adapter configuration.
	Config string `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	// True if backend has implemented the
	// [InfrastructureBackend](https://github.com/istio/api/blob/master/mixer/adapter/model/v1beta1/infrastructure_backend.proto)
	// service; false otherwise.
	//
	// If true, during configuration time, Mixer calls the `InfrastructureBackend`' rpcs
	// to validate and pass the handler configuration. And during request-time, Mixer does not pass the handler
	// configuration, and only passes the template-specific Instance payload using the template-specific handle service (Example
	// [HandlerMetricService](https://github.com/istio/istio/blob/master/mixer/template/metric/template_handler_service.proto),
	// [HandlerListEntryService](https://github.com/istio/istio/blob/master/mixer/template/logentry/template_handler_service.proto),
	// [HandleQuotaService](https://github.com/istio/istio/blob/master/mixer/template/quota/template_handler_service.proto) and more).
	// If `session_based` is false, Mixer does not expect backend to implement `InfrastructureBackend` service, and
	// only communicates with the backends during request-time through the template-specific handle service. Without
	// `InfrastructureBackend` service, Mixer passes the handler configuration on each call during request-time.
	SessionBased bool `protobuf:"varint,5,opt,name=session_based,json=sessionBased,proto3" json:"session_based,omitempty"`
}

func (m *Info) Reset()      { *m = Info{} }
func (*Info) ProtoMessage() {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_info_d3fb8d8896e5148e, []int{0}
}
func (m *Info) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Info.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(dst, src)
}
func (m *Info) XXX_Size() int {
	return m.Size()
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Info) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Info) GetTemplates() []string {
	if m != nil {
		return m.Templates
	}
	return nil
}

func (m *Info) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *Info) GetSessionBased() bool {
	if m != nil {
		return m.SessionBased
	}
	return false
}

func init() {
	proto.RegisterType((*Info)(nil), "istio.mixer.adapter.model.v1beta1.Info")
}
func (this *Info) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Info)
	if !ok {
		that2, ok := that.(Info)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Templates) != len(that1.Templates) {
		return false
	}
	for i := range this.Templates {
		if this.Templates[i] != that1.Templates[i] {
			return false
		}
	}
	if this.Config != that1.Config {
		return false
	}
	if this.SessionBased != that1.SessionBased {
		return false
	}
	return true
}
func (this *Info) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&v1beta1.Info{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Description: "+fmt.Sprintf("%#v", this.Description)+",\n")
	s = append(s, "Templates: "+fmt.Sprintf("%#v", this.Templates)+",\n")
	s = append(s, "Config: "+fmt.Sprintf("%#v", this.Config)+",\n")
	s = append(s, "SessionBased: "+fmt.Sprintf("%#v", this.SessionBased)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringInfo(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Info) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Info) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.Templates) > 0 {
		for _, s := range m.Templates {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Config) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Config)))
		i += copy(dAtA[i:], m.Config)
	}
	if m.SessionBased {
		dAtA[i] = 0x28
		i++
		if m.SessionBased {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintInfo(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Info) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if len(m.Templates) > 0 {
		for _, s := range m.Templates {
			l = len(s)
			n += 1 + l + sovInfo(uint64(l))
		}
	}
	l = len(m.Config)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if m.SessionBased {
		n += 2
	}
	return n
}

func sovInfo(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInfo(x uint64) (n int) {
	return sovInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Info) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Info{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`Templates:` + fmt.Sprintf("%v", this.Templates) + `,`,
		`Config:` + fmt.Sprintf("%v", this.Config) + `,`,
		`SessionBased:` + fmt.Sprintf("%v", this.SessionBased) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringInfo(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Info) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Templates", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Templates = append(m.Templates, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Config = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionBased", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SessionBased = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthInfo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInfo
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipInfo(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthInfo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInfo   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/adapter/model/v1beta1/info.proto", fileDescriptor_info_d3fb8d8896e5148e)
}

var fileDescriptor_info_d3fb8d8896e5148e = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4e, 0xec, 0x30,
	0x10, 0x86, 0x3d, 0x6f, 0xf3, 0x56, 0xc4, 0x40, 0xe3, 0x02, 0xa5, 0x40, 0xa3, 0x00, 0x12, 0x4a,
	0x15, 0x6b, 0xc5, 0x0d, 0xb6, 0xa3, 0x4d, 0x49, 0x01, 0x72, 0x36, 0x0e, 0xb2, 0xb4, 0xb1, 0xad,
	0xd8, 0x42, 0x94, 0x1c, 0x81, 0x8e, 0x2b, 0x70, 0x14, 0xca, 0x94, 0x5b, 0x12, 0xa7, 0xa1, 0xdc,
	0x23, 0x20, 0x4c, 0x24, 0xa8, 0xe8, 0x66, 0xbe, 0xf9, 0xa6, 0xf8, 0x7f, 0x7a, 0xd9, 0xa9, 0x47,
	0xd9, 0x73, 0xd1, 0x08, 0xeb, 0x65, 0xcf, 0x3b, 0xd3, 0xc8, 0x2d, 0x7f, 0x58, 0xd5, 0xd2, 0x8b,
	0x15, 0x57, 0xba, 0x35, 0xa5, 0xed, 0x8d, 0x37, 0xec, 0x4c, 0x39, 0xaf, 0x4c, 0x19, 0xed, 0x72,
	0xb6, 0xcb, 0x68, 0x97, 0xb3, 0x7d, 0xfe, 0x02, 0x34, 0xb9, 0xd6, 0xad, 0x61, 0x8c, 0x26, 0x5a,
	0x74, 0x32, 0x83, 0x1c, 0x8a, 0xb4, 0x8a, 0x33, 0xcb, 0xe9, 0x61, 0x23, 0xdd, 0xa6, 0x57, 0xd6,
	0x2b, 0xa3, 0xb3, 0x7f, 0xf1, 0xf4, 0x1b, 0xb1, 0x53, 0x9a, 0x7a, 0xd9, 0xd9, 0xad, 0xf0, 0xd2,
	0x65, 0x8b, 0x7c, 0x51, 0xa4, 0xd5, 0x0f, 0x60, 0x27, 0x74, 0xb9, 0x31, 0xba, 0x55, 0xf7, 0x59,
	0x12, 0x5f, 0xe7, 0x8d, 0x5d, 0xd0, 0x63, 0x27, 0x9d, 0x53, 0x46, 0xdf, 0xd5, 0xc2, 0xc9, 0x26,
	0xfb, 0x9f, 0x43, 0x71, 0x50, 0x1d, 0xcd, 0x70, 0xfd, 0xc5, 0xd6, 0xb7, 0xc3, 0x88, 0x64, 0x37,
	0x22, 0xd9, 0x8f, 0x08, 0x4f, 0x01, 0xe1, 0x35, 0x20, 0xbc, 0x05, 0x84, 0x21, 0x20, 0xbc, 0x07,
	0x84, 0x8f, 0x80, 0x64, 0x1f, 0x10, 0x9e, 0x27, 0x24, 0xc3, 0x84, 0x64, 0x37, 0x21, 0xb9, 0x29,
	0xbe, 0x23, 0x2b, 0xc3, 0x85, 0x55, 0xfc, 0x8f, 0x9e, 0xea, 0x65, 0xec, 0xe8, 0xea, 0x33, 0x00,
	0x00, 0xff, 0xff, 0xf6, 0x62, 0xa4, 0x9d, 0x4d, 0x01, 0x00, 0x00,
}
