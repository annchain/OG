package ogsyncer_interface

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *OgSyncBlockByHashRequest) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Hash":
			z.Hash, err = dc.ReadBytes(z.Hash)
			if err != nil {
				err = msgp.WrapError(err, "Hash")
				return
			}
		case "Offset":
			z.Offset, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Offset")
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
func (z *OgSyncBlockByHashRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Hash"
	err = en.Append(0x82, 0xa4, 0x48, 0x61, 0x73, 0x68)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Hash)
	if err != nil {
		err = msgp.WrapError(err, "Hash")
		return
	}
	// write "Offset"
	err = en.Append(0xa6, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Offset)
	if err != nil {
		err = msgp.WrapError(err, "Offset")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *OgSyncBlockByHashRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hash"
	o = append(o, 0x82, 0xa4, 0x48, 0x61, 0x73, 0x68)
	o = msgp.AppendBytes(o, z.Hash)
	// string "Offset"
	o = append(o, 0xa6, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74)
	o = msgp.AppendInt(o, z.Offset)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OgSyncBlockByHashRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Hash":
			z.Hash, bts, err = msgp.ReadBytesBytes(bts, z.Hash)
			if err != nil {
				err = msgp.WrapError(err, "Hash")
				return
			}
		case "Offset":
			z.Offset, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Offset")
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
func (z *OgSyncBlockByHashRequest) Msgsize() (s int) {
	s = 1 + 5 + msgp.BytesPrefixSize + len(z.Hash) + 7 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *OgSyncBlockByHashResponse) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "HasMore":
			z.HasMore, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "HasMore")
				return
			}
		case "Sequencers":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Sequencers")
				return
			}
			if cap(z.Sequencers) >= int(zb0002) {
				z.Sequencers = (z.Sequencers)[:zb0002]
			} else {
				z.Sequencers = make([]MessageContentSequencer, zb0002)
			}
			for za0001 := range z.Sequencers {
				err = z.Sequencers[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Sequencers", za0001)
					return
				}
			}
		case "Ints":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Ints")
				return
			}
			if cap(z.Ints) >= int(zb0003) {
				z.Ints = (z.Ints)[:zb0003]
			} else {
				z.Ints = make([]MessageContentInt, zb0003)
			}
			for za0002 := range z.Ints {
				err = z.Ints[za0002].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Ints", za0002)
					return
				}
			}
		case "Txs":
			var zb0004 uint32
			zb0004, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Txs")
				return
			}
			if cap(z.Txs) >= int(zb0004) {
				z.Txs = (z.Txs)[:zb0004]
			} else {
				z.Txs = make([]MessageContentTx, zb0004)
			}
			for za0003 := range z.Txs {
				err = z.Txs[za0003].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Txs", za0003)
					return
				}
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
func (z *OgSyncBlockByHashResponse) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "HasMore"
	err = en.Append(0x84, 0xa7, 0x48, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBool(z.HasMore)
	if err != nil {
		err = msgp.WrapError(err, "HasMore")
		return
	}
	// write "Sequencers"
	err = en.Append(0xaa, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Sequencers)))
	if err != nil {
		err = msgp.WrapError(err, "Sequencers")
		return
	}
	for za0001 := range z.Sequencers {
		err = z.Sequencers[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Sequencers", za0001)
			return
		}
	}
	// write "Ints"
	err = en.Append(0xa4, 0x49, 0x6e, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Ints)))
	if err != nil {
		err = msgp.WrapError(err, "Ints")
		return
	}
	for za0002 := range z.Ints {
		err = z.Ints[za0002].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Ints", za0002)
			return
		}
	}
	// write "Txs"
	err = en.Append(0xa3, 0x54, 0x78, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Txs)))
	if err != nil {
		err = msgp.WrapError(err, "Txs")
		return
	}
	for za0003 := range z.Txs {
		err = z.Txs[za0003].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Txs", za0003)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *OgSyncBlockByHashResponse) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "HasMore"
	o = append(o, 0x84, 0xa7, 0x48, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65)
	o = msgp.AppendBool(o, z.HasMore)
	// string "Sequencers"
	o = append(o, 0xaa, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Sequencers)))
	for za0001 := range z.Sequencers {
		o, err = z.Sequencers[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Sequencers", za0001)
			return
		}
	}
	// string "Ints"
	o = append(o, 0xa4, 0x49, 0x6e, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Ints)))
	for za0002 := range z.Ints {
		o, err = z.Ints[za0002].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Ints", za0002)
			return
		}
	}
	// string "Txs"
	o = append(o, 0xa3, 0x54, 0x78, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Txs)))
	for za0003 := range z.Txs {
		o, err = z.Txs[za0003].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Txs", za0003)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OgSyncBlockByHashResponse) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "HasMore":
			z.HasMore, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "HasMore")
				return
			}
		case "Sequencers":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Sequencers")
				return
			}
			if cap(z.Sequencers) >= int(zb0002) {
				z.Sequencers = (z.Sequencers)[:zb0002]
			} else {
				z.Sequencers = make([]MessageContentSequencer, zb0002)
			}
			for za0001 := range z.Sequencers {
				bts, err = z.Sequencers[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Sequencers", za0001)
					return
				}
			}
		case "Ints":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Ints")
				return
			}
			if cap(z.Ints) >= int(zb0003) {
				z.Ints = (z.Ints)[:zb0003]
			} else {
				z.Ints = make([]MessageContentInt, zb0003)
			}
			for za0002 := range z.Ints {
				bts, err = z.Ints[za0002].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Ints", za0002)
					return
				}
			}
		case "Txs":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Txs")
				return
			}
			if cap(z.Txs) >= int(zb0004) {
				z.Txs = (z.Txs)[:zb0004]
			} else {
				z.Txs = make([]MessageContentTx, zb0004)
			}
			for za0003 := range z.Txs {
				bts, err = z.Txs[za0003].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Txs", za0003)
					return
				}
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
func (z *OgSyncBlockByHashResponse) Msgsize() (s int) {
	s = 1 + 8 + msgp.BoolSize + 11 + msgp.ArrayHeaderSize
	for za0001 := range z.Sequencers {
		s += z.Sequencers[za0001].Msgsize()
	}
	s += 5 + msgp.ArrayHeaderSize
	for za0002 := range z.Ints {
		s += z.Ints[za0002].Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for za0003 := range z.Txs {
		s += z.Txs[za0003].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *OgSyncBlockByHeightRequest) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Height":
			z.Height, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Height")
				return
			}
		case "Offset":
			z.Offset, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Offset")
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
func (z OgSyncBlockByHeightRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Height"
	err = en.Append(0x82, 0xa6, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Height)
	if err != nil {
		err = msgp.WrapError(err, "Height")
		return
	}
	// write "Offset"
	err = en.Append(0xa6, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Offset)
	if err != nil {
		err = msgp.WrapError(err, "Offset")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z OgSyncBlockByHeightRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Height"
	o = append(o, 0x82, 0xa6, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74)
	o = msgp.AppendInt64(o, z.Height)
	// string "Offset"
	o = append(o, 0xa6, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74)
	o = msgp.AppendInt(o, z.Offset)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OgSyncBlockByHeightRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Height":
			z.Height, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Height")
				return
			}
		case "Offset":
			z.Offset, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Offset")
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
func (z OgSyncBlockByHeightRequest) Msgsize() (s int) {
	s = 1 + 7 + msgp.Int64Size + 7 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *OgSyncBlockByHeightResponse) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "HasMore":
			z.HasMore, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "HasMore")
				return
			}
		case "Sequencers":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Sequencers")
				return
			}
			if cap(z.Sequencers) >= int(zb0002) {
				z.Sequencers = (z.Sequencers)[:zb0002]
			} else {
				z.Sequencers = make([]MessageContentSequencer, zb0002)
			}
			for za0001 := range z.Sequencers {
				err = z.Sequencers[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Sequencers", za0001)
					return
				}
			}
		case "Ints":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Ints")
				return
			}
			if cap(z.Ints) >= int(zb0003) {
				z.Ints = (z.Ints)[:zb0003]
			} else {
				z.Ints = make([]MessageContentInt, zb0003)
			}
			for za0002 := range z.Ints {
				err = z.Ints[za0002].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Ints", za0002)
					return
				}
			}
		case "Txs":
			var zb0004 uint32
			zb0004, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Txs")
				return
			}
			if cap(z.Txs) >= int(zb0004) {
				z.Txs = (z.Txs)[:zb0004]
			} else {
				z.Txs = make([]MessageContentTx, zb0004)
			}
			for za0003 := range z.Txs {
				err = z.Txs[za0003].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Txs", za0003)
					return
				}
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
func (z *OgSyncBlockByHeightResponse) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "HasMore"
	err = en.Append(0x84, 0xa7, 0x48, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBool(z.HasMore)
	if err != nil {
		err = msgp.WrapError(err, "HasMore")
		return
	}
	// write "Sequencers"
	err = en.Append(0xaa, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Sequencers)))
	if err != nil {
		err = msgp.WrapError(err, "Sequencers")
		return
	}
	for za0001 := range z.Sequencers {
		err = z.Sequencers[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Sequencers", za0001)
			return
		}
	}
	// write "Ints"
	err = en.Append(0xa4, 0x49, 0x6e, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Ints)))
	if err != nil {
		err = msgp.WrapError(err, "Ints")
		return
	}
	for za0002 := range z.Ints {
		err = z.Ints[za0002].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Ints", za0002)
			return
		}
	}
	// write "Txs"
	err = en.Append(0xa3, 0x54, 0x78, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Txs)))
	if err != nil {
		err = msgp.WrapError(err, "Txs")
		return
	}
	for za0003 := range z.Txs {
		err = z.Txs[za0003].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Txs", za0003)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *OgSyncBlockByHeightResponse) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "HasMore"
	o = append(o, 0x84, 0xa7, 0x48, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65)
	o = msgp.AppendBool(o, z.HasMore)
	// string "Sequencers"
	o = append(o, 0xaa, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Sequencers)))
	for za0001 := range z.Sequencers {
		o, err = z.Sequencers[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Sequencers", za0001)
			return
		}
	}
	// string "Ints"
	o = append(o, 0xa4, 0x49, 0x6e, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Ints)))
	for za0002 := range z.Ints {
		o, err = z.Ints[za0002].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Ints", za0002)
			return
		}
	}
	// string "Txs"
	o = append(o, 0xa3, 0x54, 0x78, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Txs)))
	for za0003 := range z.Txs {
		o, err = z.Txs[za0003].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Txs", za0003)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OgSyncBlockByHeightResponse) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "HasMore":
			z.HasMore, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "HasMore")
				return
			}
		case "Sequencers":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Sequencers")
				return
			}
			if cap(z.Sequencers) >= int(zb0002) {
				z.Sequencers = (z.Sequencers)[:zb0002]
			} else {
				z.Sequencers = make([]MessageContentSequencer, zb0002)
			}
			for za0001 := range z.Sequencers {
				bts, err = z.Sequencers[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Sequencers", za0001)
					return
				}
			}
		case "Ints":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Ints")
				return
			}
			if cap(z.Ints) >= int(zb0003) {
				z.Ints = (z.Ints)[:zb0003]
			} else {
				z.Ints = make([]MessageContentInt, zb0003)
			}
			for za0002 := range z.Ints {
				bts, err = z.Ints[za0002].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Ints", za0002)
					return
				}
			}
		case "Txs":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Txs")
				return
			}
			if cap(z.Txs) >= int(zb0004) {
				z.Txs = (z.Txs)[:zb0004]
			} else {
				z.Txs = make([]MessageContentTx, zb0004)
			}
			for za0003 := range z.Txs {
				bts, err = z.Txs[za0003].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Txs", za0003)
					return
				}
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
func (z *OgSyncBlockByHeightResponse) Msgsize() (s int) {
	s = 1 + 8 + msgp.BoolSize + 11 + msgp.ArrayHeaderSize
	for za0001 := range z.Sequencers {
		s += z.Sequencers[za0001].Msgsize()
	}
	s += 5 + msgp.ArrayHeaderSize
	for za0002 := range z.Ints {
		s += z.Ints[za0002].Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for za0003 := range z.Txs {
		s += z.Txs[za0003].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *OgSyncByHashesRequest) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Hashes":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Hashes")
				return
			}
			if cap(z.Hashes) >= int(zb0002) {
				z.Hashes = (z.Hashes)[:zb0002]
			} else {
				z.Hashes = make([][]byte, zb0002)
			}
			for za0001 := range z.Hashes {
				z.Hashes[za0001], err = dc.ReadBytes(z.Hashes[za0001])
				if err != nil {
					err = msgp.WrapError(err, "Hashes", za0001)
					return
				}
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
func (z *OgSyncByHashesRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Hashes"
	err = en.Append(0x81, 0xa6, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Hashes)))
	if err != nil {
		err = msgp.WrapError(err, "Hashes")
		return
	}
	for za0001 := range z.Hashes {
		err = en.WriteBytes(z.Hashes[za0001])
		if err != nil {
			err = msgp.WrapError(err, "Hashes", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *OgSyncByHashesRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Hashes"
	o = append(o, 0x81, 0xa6, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Hashes)))
	for za0001 := range z.Hashes {
		o = msgp.AppendBytes(o, z.Hashes[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OgSyncByHashesRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Hashes":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Hashes")
				return
			}
			if cap(z.Hashes) >= int(zb0002) {
				z.Hashes = (z.Hashes)[:zb0002]
			} else {
				z.Hashes = make([][]byte, zb0002)
			}
			for za0001 := range z.Hashes {
				z.Hashes[za0001], bts, err = msgp.ReadBytesBytes(bts, z.Hashes[za0001])
				if err != nil {
					err = msgp.WrapError(err, "Hashes", za0001)
					return
				}
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
func (z *OgSyncByHashesRequest) Msgsize() (s int) {
	s = 1 + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.Hashes {
		s += msgp.BytesPrefixSize + len(z.Hashes[za0001])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *OgSyncByHashesResponse) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Sequencers":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Sequencers")
				return
			}
			if cap(z.Sequencers) >= int(zb0002) {
				z.Sequencers = (z.Sequencers)[:zb0002]
			} else {
				z.Sequencers = make([]MessageContentSequencer, zb0002)
			}
			for za0001 := range z.Sequencers {
				err = z.Sequencers[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Sequencers", za0001)
					return
				}
			}
		case "Txs":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Txs")
				return
			}
			if cap(z.Txs) >= int(zb0003) {
				z.Txs = (z.Txs)[:zb0003]
			} else {
				z.Txs = make([]MessageContentTx, zb0003)
			}
			for za0002 := range z.Txs {
				err = z.Txs[za0002].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Txs", za0002)
					return
				}
			}
		case "Ints":
			var zb0004 uint32
			zb0004, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Ints")
				return
			}
			if cap(z.Ints) >= int(zb0004) {
				z.Ints = (z.Ints)[:zb0004]
			} else {
				z.Ints = make([]MessageContentInt, zb0004)
			}
			for za0003 := range z.Ints {
				err = z.Ints[za0003].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Ints", za0003)
					return
				}
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
func (z *OgSyncByHashesResponse) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Sequencers"
	err = en.Append(0x83, 0xaa, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Sequencers)))
	if err != nil {
		err = msgp.WrapError(err, "Sequencers")
		return
	}
	for za0001 := range z.Sequencers {
		err = z.Sequencers[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Sequencers", za0001)
			return
		}
	}
	// write "Txs"
	err = en.Append(0xa3, 0x54, 0x78, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Txs)))
	if err != nil {
		err = msgp.WrapError(err, "Txs")
		return
	}
	for za0002 := range z.Txs {
		err = z.Txs[za0002].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Txs", za0002)
			return
		}
	}
	// write "Ints"
	err = en.Append(0xa4, 0x49, 0x6e, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Ints)))
	if err != nil {
		err = msgp.WrapError(err, "Ints")
		return
	}
	for za0003 := range z.Ints {
		err = z.Ints[za0003].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Ints", za0003)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *OgSyncByHashesResponse) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Sequencers"
	o = append(o, 0x83, 0xaa, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Sequencers)))
	for za0001 := range z.Sequencers {
		o, err = z.Sequencers[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Sequencers", za0001)
			return
		}
	}
	// string "Txs"
	o = append(o, 0xa3, 0x54, 0x78, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Txs)))
	for za0002 := range z.Txs {
		o, err = z.Txs[za0002].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Txs", za0002)
			return
		}
	}
	// string "Ints"
	o = append(o, 0xa4, 0x49, 0x6e, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Ints)))
	for za0003 := range z.Ints {
		o, err = z.Ints[za0003].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Ints", za0003)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OgSyncByHashesResponse) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Sequencers":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Sequencers")
				return
			}
			if cap(z.Sequencers) >= int(zb0002) {
				z.Sequencers = (z.Sequencers)[:zb0002]
			} else {
				z.Sequencers = make([]MessageContentSequencer, zb0002)
			}
			for za0001 := range z.Sequencers {
				bts, err = z.Sequencers[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Sequencers", za0001)
					return
				}
			}
		case "Txs":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Txs")
				return
			}
			if cap(z.Txs) >= int(zb0003) {
				z.Txs = (z.Txs)[:zb0003]
			} else {
				z.Txs = make([]MessageContentTx, zb0003)
			}
			for za0002 := range z.Txs {
				bts, err = z.Txs[za0002].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Txs", za0002)
					return
				}
			}
		case "Ints":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Ints")
				return
			}
			if cap(z.Ints) >= int(zb0004) {
				z.Ints = (z.Ints)[:zb0004]
			} else {
				z.Ints = make([]MessageContentInt, zb0004)
			}
			for za0003 := range z.Ints {
				bts, err = z.Ints[za0003].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Ints", za0003)
					return
				}
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
func (z *OgSyncByHashesResponse) Msgsize() (s int) {
	s = 1 + 11 + msgp.ArrayHeaderSize
	for za0001 := range z.Sequencers {
		s += z.Sequencers[za0001].Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for za0002 := range z.Txs {
		s += z.Txs[za0002].Msgsize()
	}
	s += 5 + msgp.ArrayHeaderSize
	for za0003 := range z.Ints {
		s += z.Ints[za0003].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *OgSyncLatestHeightRequest) DecodeMsg(dc *msgp.Reader) (err error) {
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
func (z OgSyncLatestHeightRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 0
	err = en.Append(0x80)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z OgSyncLatestHeightRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 0
	o = append(o, 0x80)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OgSyncLatestHeightRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
func (z OgSyncLatestHeightRequest) Msgsize() (s int) {
	s = 1
	return
}

// DecodeMsg implements msgp.Decodable
func (z *OgSyncLatestHeightResponse) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "MyHeight":
			z.MyHeight, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "MyHeight")
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
func (z OgSyncLatestHeightResponse) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "MyHeight"
	err = en.Append(0x81, 0xa8, 0x4d, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.MyHeight)
	if err != nil {
		err = msgp.WrapError(err, "MyHeight")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z OgSyncLatestHeightResponse) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "MyHeight"
	o = append(o, 0x81, 0xa8, 0x4d, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74)
	o = msgp.AppendInt64(o, z.MyHeight)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OgSyncLatestHeightResponse) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "MyHeight":
			z.MyHeight, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MyHeight")
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
func (z OgSyncLatestHeightResponse) Msgsize() (s int) {
	s = 1 + 9 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *OgSyncMessageType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int
		zb0001, err = dc.ReadInt()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = OgSyncMessageType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z OgSyncMessageType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt(int(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z OgSyncMessageType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt(o, int(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OgSyncMessageType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int
		zb0001, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = OgSyncMessageType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z OgSyncMessageType) Msgsize() (s int) {
	s = msgp.IntSize
	return
}
