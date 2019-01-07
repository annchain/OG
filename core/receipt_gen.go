package core

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import "github.com/tinylib/msgp/msgp"

// DecodeMsg implements msgp.Decodable
func (z *Receipt) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	err = z.TxHash.DecodeMsg(dc)
	if err != nil {
		return
	}
	{
		var zb0002 int
		zb0002, err = dc.ReadInt()
		if err != nil {
			return
		}
		z.Status = ReceiptStatus(zb0002)
	}
	z.ProcessResult, err = dc.ReadString()
	if err != nil {
		return
	}
	err = z.ContractAddress.DecodeMsg(dc)
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Receipt) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	err = z.TxHash.EncodeMsg(en)
	if err != nil {
		return
	}
	err = en.WriteInt(int(z.Status))
	if err != nil {
		return
	}
	err = en.WriteString(z.ProcessResult)
	if err != nil {
		return
	}
	err = z.ContractAddress.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Receipt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o, err = z.TxHash.MarshalMsg(o)
	if err != nil {
		return
	}
	o = msgp.AppendInt(o, int(z.Status))
	o = msgp.AppendString(o, z.ProcessResult)
	o, err = z.ContractAddress.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Receipt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	bts, err = z.TxHash.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	{
		var zb0002 int
		zb0002, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			return
		}
		z.Status = ReceiptStatus(zb0002)
	}
	z.ProcessResult, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		return
	}
	bts, err = z.ContractAddress.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Receipt) Msgsize() (s int) {
	s = 1 + z.TxHash.Msgsize() + msgp.IntSize + msgp.StringPrefixSize + len(z.ProcessResult) + z.ContractAddress.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ReceiptSet) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0003 uint32
	zb0003, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	if (*z) == nil {
		(*z) = make(ReceiptSet, zb0003)
	} else if len((*z)) > 0 {
		for key := range *z {
			delete((*z), key)
		}
	}
	for zb0003 > 0 {
		zb0003--
		var zb0001 string
		var zb0002 *Receipt
		zb0001, err = dc.ReadString()
		if err != nil {
			return
		}
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				return
			}
			zb0002 = nil
		} else {
			if zb0002 == nil {
				zb0002 = new(Receipt)
			}
			err = zb0002.DecodeMsg(dc)
			if err != nil {
				return
			}
		}
		(*z)[zb0001] = zb0002
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z ReceiptSet) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0004, zb0005 := range z {
		err = en.WriteString(zb0004)
		if err != nil {
			return
		}
		if zb0005 == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = zb0005.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ReceiptSet) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, uint32(len(z)))
	for zb0004, zb0005 := range z {
		o = msgp.AppendString(o, zb0004)
		if zb0005 == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zb0005.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ReceiptSet) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	if (*z) == nil {
		(*z) = make(ReceiptSet, zb0003)
	} else if len((*z)) > 0 {
		for key := range *z {
			delete((*z), key)
		}
	}
	for zb0003 > 0 {
		var zb0001 string
		var zb0002 *Receipt
		zb0003--
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			return
		}
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			zb0002 = nil
		} else {
			if zb0002 == nil {
				zb0002 = new(Receipt)
			}
			bts, err = zb0002.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
		(*z)[zb0001] = zb0002
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ReceiptSet) Msgsize() (s int) {
	s = msgp.MapHeaderSize
	if z != nil {
		for zb0004, zb0005 := range z {
			_ = zb0005
			s += msgp.StringPrefixSize + len(zb0004)
			if zb0005 == nil {
				s += msgp.NilSize
			} else {
				s += zb0005.Msgsize()
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ReceiptStatus) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int
		zb0001, err = dc.ReadInt()
		if err != nil {
			return
		}
		(*z) = ReceiptStatus(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z ReceiptStatus) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt(int(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ReceiptStatus) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt(o, int(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ReceiptStatus) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int
		zb0001, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			return
		}
		(*z) = ReceiptStatus(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ReceiptStatus) Msgsize() (s int) {
	s = msgp.IntSize
	return
}