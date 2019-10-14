package dkg

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *DkgBasicInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	z.TermId, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "TermId")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z DkgBasicInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.TermId)
	if err != nil {
		err = msgp.WrapError(err, "TermId")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z DkgBasicInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 1
	o = append(o, 0x91)
	o = msgp.AppendUint32(o, z.TermId)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DkgBasicInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	z.TermId, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "TermId")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z DkgBasicInfo) Msgsize() (s int) {
	s = 1 + msgp.Uint32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *DkgMessageType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint16
		zb0001, err = dc.ReadUint16()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = DkgMessageType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z DkgMessageType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint16(uint16(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z DkgMessageType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint16(o, uint16(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DkgMessageType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint16
		zb0001, bts, err = msgp.ReadUint16Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = DkgMessageType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z DkgMessageType) Msgsize() (s int) {
	s = msgp.Uint16Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MessageDkgDeal) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo")
		return
	}
	if zb0002 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0002}
		return
	}
	z.DkgBasicInfo.TermId, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo", "TermId")
		return
	}
	z.Data, err = dc.ReadBytes(z.Data)
	if err != nil {
		err = msgp.WrapError(err, "Data")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MessageDkgDeal) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	// array header, size 1
	err = en.Append(0x92, 0x91)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.DkgBasicInfo.TermId)
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo", "TermId")
		return
	}
	err = en.WriteBytes(z.Data)
	if err != nil {
		err = msgp.WrapError(err, "Data")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MessageDkgDeal) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	// array header, size 1
	o = append(o, 0x92, 0x91)
	o = msgp.AppendUint32(o, z.DkgBasicInfo.TermId)
	o = msgp.AppendBytes(o, z.Data)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MessageDkgDeal) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo")
		return
	}
	if zb0002 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0002}
		return
	}
	z.DkgBasicInfo.TermId, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo", "TermId")
		return
	}
	z.Data, bts, err = msgp.ReadBytesBytes(bts, z.Data)
	if err != nil {
		err = msgp.WrapError(err, "Data")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MessageDkgDeal) Msgsize() (s int) {
	s = 1 + 1 + msgp.Uint32Size + msgp.BytesPrefixSize + len(z.Data)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MessageDkgDealResponse) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo")
		return
	}
	if zb0002 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0002}
		return
	}
	z.DkgBasicInfo.TermId, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo", "TermId")
		return
	}
	z.Data, err = dc.ReadBytes(z.Data)
	if err != nil {
		err = msgp.WrapError(err, "Data")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MessageDkgDealResponse) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	// array header, size 1
	err = en.Append(0x92, 0x91)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.DkgBasicInfo.TermId)
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo", "TermId")
		return
	}
	err = en.WriteBytes(z.Data)
	if err != nil {
		err = msgp.WrapError(err, "Data")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MessageDkgDealResponse) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	// array header, size 1
	o = append(o, 0x92, 0x91)
	o = msgp.AppendUint32(o, z.DkgBasicInfo.TermId)
	o = msgp.AppendBytes(o, z.Data)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MessageDkgDealResponse) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo")
		return
	}
	if zb0002 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0002}
		return
	}
	z.DkgBasicInfo.TermId, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo", "TermId")
		return
	}
	z.Data, bts, err = msgp.ReadBytesBytes(bts, z.Data)
	if err != nil {
		err = msgp.WrapError(err, "Data")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MessageDkgDealResponse) Msgsize() (s int) {
	s = 1 + 1 + msgp.Uint32Size + msgp.BytesPrefixSize + len(z.Data)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MessageDkgGenesisPublicKey) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo")
		return
	}
	if zb0002 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0002}
		return
	}
	z.DkgBasicInfo.TermId, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo", "TermId")
		return
	}
	z.PublicKeyBytes, err = dc.ReadBytes(z.PublicKeyBytes)
	if err != nil {
		err = msgp.WrapError(err, "PublicKeyBytes")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MessageDkgGenesisPublicKey) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	// array header, size 1
	err = en.Append(0x92, 0x91)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.DkgBasicInfo.TermId)
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo", "TermId")
		return
	}
	err = en.WriteBytes(z.PublicKeyBytes)
	if err != nil {
		err = msgp.WrapError(err, "PublicKeyBytes")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MessageDkgGenesisPublicKey) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	// array header, size 1
	o = append(o, 0x92, 0x91)
	o = msgp.AppendUint32(o, z.DkgBasicInfo.TermId)
	o = msgp.AppendBytes(o, z.PublicKeyBytes)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MessageDkgGenesisPublicKey) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo")
		return
	}
	if zb0002 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0002}
		return
	}
	z.DkgBasicInfo.TermId, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo", "TermId")
		return
	}
	z.PublicKeyBytes, bts, err = msgp.ReadBytesBytes(bts, z.PublicKeyBytes)
	if err != nil {
		err = msgp.WrapError(err, "PublicKeyBytes")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MessageDkgGenesisPublicKey) Msgsize() (s int) {
	s = 1 + 1 + msgp.Uint32Size + msgp.BytesPrefixSize + len(z.PublicKeyBytes)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MessageDkgSigSets) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo")
		return
	}
	if zb0002 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0002}
		return
	}
	z.DkgBasicInfo.TermId, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo", "TermId")
		return
	}
	z.PkBls, err = dc.ReadBytes(z.PkBls)
	if err != nil {
		err = msgp.WrapError(err, "PkBls")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MessageDkgSigSets) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	// array header, size 1
	err = en.Append(0x92, 0x91)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.DkgBasicInfo.TermId)
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo", "TermId")
		return
	}
	err = en.WriteBytes(z.PkBls)
	if err != nil {
		err = msgp.WrapError(err, "PkBls")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MessageDkgSigSets) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	// array header, size 1
	o = append(o, 0x92, 0x91)
	o = msgp.AppendUint32(o, z.DkgBasicInfo.TermId)
	o = msgp.AppendBytes(o, z.PkBls)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MessageDkgSigSets) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo")
		return
	}
	if zb0002 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0002}
		return
	}
	z.DkgBasicInfo.TermId, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "DkgBasicInfo", "TermId")
		return
	}
	z.PkBls, bts, err = msgp.ReadBytesBytes(bts, z.PkBls)
	if err != nil {
		err = msgp.WrapError(err, "PkBls")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MessageDkgSigSets) Msgsize() (s int) {
	s = 1 + 1 + msgp.Uint32Size + msgp.BytesPrefixSize + len(z.PkBls)
	return
}
