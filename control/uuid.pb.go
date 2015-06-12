// Code generated by protoc-gen-gogo.
// source: uuid.proto
// DO NOT EDIT!

package control

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io "io"
import fmt "fmt"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// / Type representing a 128-bit UUID.
//
// The bytes of the UUID should be packed in little-endian **byte** (not bit) order. For example, the UUID `f47ac10b-58cc-4372-a567-0e02b2c3d479` should be encoded as `UUID{ low: 0x7243cc580bc17af4, high: 0x79d4c3b2020e67a5 }`
type UUID struct {
	Low              *uint64 `protobuf:"varint,1,req,name=low" json:"low,omitempty"`
	High             *uint64 `protobuf:"varint,2,req,name=high" json:"high,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UUID) Reset()         { *m = UUID{} }
func (m *UUID) String() string { return proto.CompactTextString(m) }
func (*UUID) ProtoMessage()    {}

func (m *UUID) GetLow() uint64 {
	if m != nil && m.Low != nil {
		return *m.Low
	}
	return 0
}

func (m *UUID) GetHigh() uint64 {
	if m != nil && m.High != nil {
		return *m.High
	}
	return 0
}

func init() {
}
func (m *UUID) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Low", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Low = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field High", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.High = &v
			hasFields[0] |= uint64(0x00000002)
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := skipUuid(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("low")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("high")
	}

	return nil
}
func skipUuid(data []byte) (n int, err error) {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if index >= l {
					return 0, io.ErrUnexpectedEOF
				}
				index++
				if data[index-1] < 0x80 {
					break
				}
			}
			return index, nil
		case 1:
			index += 8
			return index, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			index += length
			return index, nil
		case 3:
			for {
				var wire uint64
				var start int = index
				for shift := uint(0); ; shift += 7 {
					if index >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[index]
					index++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				wireType := int(wire & 0x7)
				if wireType == 4 {
					break
				}
				next, err := skipUuid(data[start:])
				if err != nil {
					return 0, err
				}
				index = start + next
			}
			return index, nil
		case 4:
			return index, nil
		case 5:
			index += 4
			return index, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}
func (m *UUID) Size() (n int) {
	var l int
	_ = l
	if m.Low != nil {
		n += 1 + sovUuid(uint64(*m.Low))
	}
	if m.High != nil {
		n += 1 + sovUuid(uint64(*m.High))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovUuid(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUuid(x uint64) (n int) {
	return sovUuid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UUID) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *UUID) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Low == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("low")
	} else {
		data[i] = 0x8
		i++
		i = encodeVarintUuid(data, i, uint64(*m.Low))
	}
	if m.High == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("high")
	} else {
		data[i] = 0x10
		i++
		i = encodeVarintUuid(data, i, uint64(*m.High))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Uuid(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Uuid(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintUuid(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
