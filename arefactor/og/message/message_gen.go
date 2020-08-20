package message

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *OgMessageHeaderRequest) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Amount":
			z.Amount, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Amount")
				return
			}
		case "Skip":
			z.Skip, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Skip")
				return
			}
		case "Reverse":
			z.Reverse, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Reverse")
				return
			}
		case "RequestId":
			z.RequestId, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "RequestId")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *OgMessageHeaderRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Amount"
	err = en.Append(0x84, 0xa6, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Amount)
	if err != nil {
		err = msgp.WrapError(err, "Amount")
		return
	}
	// write "Skip"
	err = en.Append(0xa4, 0x53, 0x6b, 0x69, 0x70)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Skip)
	if err != nil {
		err = msgp.WrapError(err, "Skip")
		return
	}
	// write "Reverse"
	err = en.Append(0xa7, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Reverse)
	if err != nil {
		err = msgp.WrapError(err, "Reverse")
		return
	}
	// write "RequestId"
	err = en.Append(0xa9, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.RequestId)
	if err != nil {
		err = msgp.WrapError(err, "RequestId")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *OgMessageHeaderRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Amount"
	o = append(o, 0x84, 0xa6, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendUint64(o, z.Amount)
	// string "Skip"
	o = append(o, 0xa4, 0x53, 0x6b, 0x69, 0x70)
	o = msgp.AppendUint64(o, z.Skip)
	// string "Reverse"
	o = append(o, 0xa7, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65)
	o = msgp.AppendBool(o, z.Reverse)
	// string "RequestId"
	o = append(o, 0xa9, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64)
	o = msgp.AppendUint32(o, z.RequestId)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OgMessageHeaderRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Amount":
			z.Amount, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Amount")
				return
			}
		case "Skip":
			z.Skip, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Skip")
				return
			}
		case "Reverse":
			z.Reverse, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Reverse")
				return
			}
		case "RequestId":
			z.RequestId, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RequestId")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *OgMessageHeaderRequest) Msgsize() (s int) {
	s = 1 + 7 + msgp.Uint64Size + 5 + msgp.Uint64Size + 8 + msgp.BoolSize + 10 + msgp.Uint32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *OgMessagePing) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Protocol":
			z.Protocol, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Protocol")
				return
			}
		case "NetworkId":
			z.NetworkId, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "NetworkId")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z OgMessagePing) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Protocol"
	err = en.Append(0x82, 0xa8, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteString(z.Protocol)
	if err != nil {
		err = msgp.WrapError(err, "Protocol")
		return
	}
	// write "NetworkId"
	err = en.Append(0xa9, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.NetworkId)
	if err != nil {
		err = msgp.WrapError(err, "NetworkId")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z OgMessagePing) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Protocol"
	o = append(o, 0x82, 0xa8, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c)
	o = msgp.AppendString(o, z.Protocol)
	// string "NetworkId"
	o = append(o, 0xa9, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64)
	o = msgp.AppendString(o, z.NetworkId)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OgMessagePing) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Protocol":
			z.Protocol, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Protocol")
				return
			}
		case "NetworkId":
			z.NetworkId, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "NetworkId")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z OgMessagePing) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.Protocol) + 10 + msgp.StringPrefixSize + len(z.NetworkId)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *OgMessagePong) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Protocol":
			z.Protocol, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Protocol")
				return
			}
		case "NetworkId":
			z.NetworkId, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "NetworkId")
				return
			}
		case "Close":
			z.Close, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Close")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z OgMessagePong) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Protocol"
	err = en.Append(0x83, 0xa8, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteString(z.Protocol)
	if err != nil {
		err = msgp.WrapError(err, "Protocol")
		return
	}
	// write "NetworkId"
	err = en.Append(0xa9, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.NetworkId)
	if err != nil {
		err = msgp.WrapError(err, "NetworkId")
		return
	}
	// write "Close"
	err = en.Append(0xa5, 0x43, 0x6c, 0x6f, 0x73, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Close)
	if err != nil {
		err = msgp.WrapError(err, "Close")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z OgMessagePong) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Protocol"
	o = append(o, 0x83, 0xa8, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c)
	o = msgp.AppendString(o, z.Protocol)
	// string "NetworkId"
	o = append(o, 0xa9, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64)
	o = msgp.AppendString(o, z.NetworkId)
	// string "Close"
	o = append(o, 0xa5, 0x43, 0x6c, 0x6f, 0x73, 0x65)
	o = msgp.AppendBool(o, z.Close)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OgMessagePong) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Protocol":
			z.Protocol, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Protocol")
				return
			}
		case "NetworkId":
			z.NetworkId, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "NetworkId")
				return
			}
		case "Close":
			z.Close, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Close")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z OgMessagePong) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.Protocol) + 10 + msgp.StringPrefixSize + len(z.NetworkId) + 6 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *OgMessageType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int
		zb0001, err = dc.ReadInt()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = OgMessageType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z OgMessageType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt(int(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z OgMessageType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt(o, int(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OgMessageType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int
		zb0001, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = OgMessageType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z OgMessageType) Msgsize() (s int) {
	s = msgp.IntSize
	return
}