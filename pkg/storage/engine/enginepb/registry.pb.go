// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/engine/enginepb/registry.proto

package enginepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RegistryVersion int32

const (
	// The only version so far.
	RegistryVersion_Base RegistryVersion = 0
)

var RegistryVersion_name = map[int32]string{
	0: "Base",
}
var RegistryVersion_value = map[string]int32{
	"Base": 0,
}

func (x RegistryVersion) String() string {
	return proto.EnumName(RegistryVersion_name, int32(x))
}
func (RegistryVersion) EnumDescriptor() ([]byte, []int) { return fileDescriptorRegistry, []int{0} }

type EncryptionType int32

const (
	// No encryption applied, not used for the registry.
	EncryptionType_Plaintext EncryptionType = 0
	// AES in counter mode.
	EncryptionType_AES_CTR EncryptionType = 1
)

var EncryptionType_name = map[int32]string{
	0: "Plaintext",
	1: "AES_CTR",
}
var EncryptionType_value = map[string]int32{
	"Plaintext": 0,
	"AES_CTR":   1,
}

func (x EncryptionType) String() string {
	return proto.EnumName(EncryptionType_name, int32(x))
}
func (EncryptionType) EnumDescriptor() ([]byte, []int) { return fileDescriptorRegistry, []int{1} }

type Registry struct {
	// version is currently always Base.
	Version RegistryVersion `protobuf:"varint,1,opt,name=version,proto3,enum=cockroach.storage.engine.enginepb.RegistryVersion" json:"version,omitempty"`
	Files   []*FileEntry    `protobuf:"bytes,2,rep,name=files" json:"files,omitempty"`
}

func (m *Registry) Reset()                    { *m = Registry{} }
func (m *Registry) String() string            { return proto.CompactTextString(m) }
func (*Registry) ProtoMessage()               {}
func (*Registry) Descriptor() ([]byte, []int) { return fileDescriptorRegistry, []int{0} }

type FileEntry struct {
	// File path relative to the DB directory.
	// TODO(mberhault): figure out if we need an alternate representation for Windows.
	Filename string `protobuf:"bytes,1,opt,name=Filename,proto3" json:"Filename,omitempty"`
	// The type of encryption applied.
	Type EncryptionType `protobuf:"varint,2,opt,name=type,proto3,enum=cockroach.storage.engine.enginepb.EncryptionType" json:"type,omitempty"`
	// Encryption fields.
	// ID (hash) of the key in use, if any.
	KeyId []byte `protobuf:"bytes,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// Initialization vector, of size 96 bits (12 bytes) for AES.
	Nonce []byte `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Counter, allowing 2^32 blocks per file, so 64GiB.
	Counter uint32 `protobuf:"varint,5,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (m *FileEntry) Reset()                    { *m = FileEntry{} }
func (m *FileEntry) String() string            { return proto.CompactTextString(m) }
func (*FileEntry) ProtoMessage()               {}
func (*FileEntry) Descriptor() ([]byte, []int) { return fileDescriptorRegistry, []int{1} }

func init() {
	proto.RegisterType((*Registry)(nil), "cockroach.storage.engine.enginepb.Registry")
	proto.RegisterType((*FileEntry)(nil), "cockroach.storage.engine.enginepb.FileEntry")
	proto.RegisterEnum("cockroach.storage.engine.enginepb.RegistryVersion", RegistryVersion_name, RegistryVersion_value)
	proto.RegisterEnum("cockroach.storage.engine.enginepb.EncryptionType", EncryptionType_name, EncryptionType_value)
}
func (m *Registry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Registry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRegistry(dAtA, i, uint64(m.Version))
	}
	if len(m.Files) > 0 {
		for _, msg := range m.Files {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRegistry(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *FileEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileEntry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Filename) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.Filename)))
		i += copy(dAtA[i:], m.Filename)
	}
	if m.Type != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRegistry(dAtA, i, uint64(m.Type))
	}
	if len(m.KeyId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.KeyId)))
		i += copy(dAtA[i:], m.KeyId)
	}
	if len(m.Nonce) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.Nonce)))
		i += copy(dAtA[i:], m.Nonce)
	}
	if m.Counter != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintRegistry(dAtA, i, uint64(m.Counter))
	}
	return i, nil
}

func encodeVarintRegistry(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Registry) Size() (n int) {
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovRegistry(uint64(m.Version))
	}
	if len(m.Files) > 0 {
		for _, e := range m.Files {
			l = e.Size()
			n += 1 + l + sovRegistry(uint64(l))
		}
	}
	return n
}

func (m *FileEntry) Size() (n int) {
	var l int
	_ = l
	l = len(m.Filename)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovRegistry(uint64(m.Type))
	}
	l = len(m.KeyId)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	if m.Counter != 0 {
		n += 1 + sovRegistry(uint64(m.Counter))
	}
	return n
}

func sovRegistry(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRegistry(x uint64) (n int) {
	return sovRegistry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Registry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegistry
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
			return fmt.Errorf("proto: Registry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Registry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (RegistryVersion(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Files", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Files = append(m.Files, &FileEntry{})
			if err := m.Files[len(m.Files)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRegistry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegistry
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
func (m *FileEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegistry
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
			return fmt.Errorf("proto: FileEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filename", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
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
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filename = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (EncryptionType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyId = append(m.KeyId[:0], dAtA[iNdEx:postIndex]...)
			if m.KeyId == nil {
				m.KeyId = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = append(m.Nonce[:0], dAtA[iNdEx:postIndex]...)
			if m.Nonce == nil {
				m.Nonce = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRegistry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegistry
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
func skipRegistry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRegistry
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
					return 0, ErrIntOverflowRegistry
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
					return 0, ErrIntOverflowRegistry
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
				return 0, ErrInvalidLengthRegistry
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRegistry
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
				next, err := skipRegistry(dAtA[start:])
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
	ErrInvalidLengthRegistry = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRegistry   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("storage/engine/enginepb/registry.proto", fileDescriptorRegistry) }

var fileDescriptorRegistry = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4e, 0xfa, 0x40,
	0x14, 0xc6, 0x3b, 0x40, 0xa1, 0x0c, 0x7f, 0xf8, 0x93, 0x09, 0x26, 0x0d, 0x26, 0x4d, 0x65, 0x61,
	0x1a, 0x42, 0x4a, 0xc4, 0x13, 0x88, 0xa9, 0x89, 0x89, 0x0b, 0x33, 0x12, 0x17, 0x6e, 0x48, 0x29,
	0xcf, 0x3a, 0x01, 0x67, 0x9a, 0xe9, 0x68, 0xec, 0x2d, 0x3c, 0x80, 0xd7, 0xf0, 0x0e, 0x2c, 0x5d,
	0xba, 0xd4, 0x7a, 0x11, 0x43, 0x4b, 0x49, 0x74, 0x23, 0xab, 0x99, 0xef, 0xe5, 0xfd, 0xbe, 0xf7,
	0xcd, 0x3c, 0x7c, 0x18, 0x2b, 0x21, 0xfd, 0x10, 0x86, 0xc0, 0x43, 0xc6, 0x8b, 0x23, 0x9a, 0x0d,
	0x25, 0x84, 0x2c, 0x56, 0x32, 0x71, 0x23, 0x29, 0x94, 0x20, 0x07, 0x81, 0x08, 0x16, 0x52, 0xf8,
	0xc1, 0x9d, 0xbb, 0x21, 0xdc, 0xbc, 0xd5, 0x2d, 0x88, 0x6e, 0x27, 0x14, 0xa1, 0xc8, 0xba, 0x87,
	0xeb, 0x5b, 0x0e, 0xf6, 0x5e, 0x10, 0x36, 0xe8, 0xc6, 0x8b, 0x5c, 0xe0, 0xda, 0x23, 0xc8, 0x98,
	0x09, 0x6e, 0x22, 0x1b, 0x39, 0xad, 0xd1, 0xc8, 0xfd, 0xd3, 0xd7, 0x2d, 0xe8, 0xeb, 0x9c, 0xa4,
	0x85, 0x05, 0x19, 0x63, 0xfd, 0x96, 0x2d, 0x21, 0x36, 0x4b, 0x76, 0xd9, 0x69, 0x8c, 0x06, 0x3b,
	0x78, 0x9d, 0xb1, 0x25, 0x78, 0x5c, 0xc9, 0x84, 0xe6, 0x68, 0xef, 0x15, 0xe1, 0xfa, 0xb6, 0x48,
	0xba, 0xd8, 0x58, 0x0b, 0xee, 0xdf, 0x43, 0x16, 0xb0, 0x4e, 0xb7, 0x9a, 0x78, 0xb8, 0xa2, 0x92,
	0x08, 0xcc, 0x52, 0x16, 0xfc, 0x68, 0x87, 0x61, 0x1e, 0x0f, 0x64, 0x12, 0x29, 0x26, 0xf8, 0x24,
	0x89, 0x80, 0x66, 0x38, 0xd9, 0xc3, 0xd5, 0x05, 0x24, 0x53, 0x36, 0x37, 0xcb, 0x36, 0x72, 0xfe,
	0x51, 0x7d, 0x01, 0xc9, 0xf9, 0x9c, 0x74, 0xb0, 0xce, 0x05, 0x0f, 0xc0, 0xac, 0xe4, 0xd5, 0x4c,
	0x10, 0x13, 0xd7, 0x02, 0xf1, 0xc0, 0x15, 0x48, 0x53, 0xb7, 0x91, 0xd3, 0xa4, 0x85, 0xec, 0xef,
	0xe3, 0xff, 0xbf, 0xfe, 0x85, 0x18, 0xb8, 0x32, 0xf6, 0x63, 0x68, 0x6b, 0xfd, 0x01, 0x6e, 0xfd,
	0x9c, 0x4d, 0x9a, 0xb8, 0x7e, 0xb9, 0xf4, 0x19, 0x57, 0xf0, 0xa4, 0xda, 0x1a, 0x69, 0xe0, 0xda,
	0x89, 0x77, 0x35, 0x3d, 0x9d, 0xd0, 0x36, 0x1a, 0xf7, 0x56, 0x9f, 0x96, 0xb6, 0x4a, 0x2d, 0xf4,
	0x96, 0x5a, 0xe8, 0x3d, 0xb5, 0xd0, 0x47, 0x6a, 0xa1, 0xe7, 0x2f, 0x4b, 0xbb, 0x31, 0x8a, 0xa7,
	0xcc, 0xaa, 0xd9, 0x32, 0x8f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x90, 0x91, 0x7d, 0x2f,
	0x02, 0x00, 0x00,
}
