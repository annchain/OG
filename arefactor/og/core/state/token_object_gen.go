package state

//
//// Code generated by github.com/tinylib/msgp DO NOT EDIT.
//
//import (
//	"github.com/annchain/OG/common/math"
//	"github.com/tinylib/msgp/msgp"
//)
//
//// DecodeMsg implements msgp.Decodable
//func (z *TokenObject) DecodeMsg(dc *msgp.Reader) (err error) {
//	var zb0001 uint32
//	zb0001, err = dc.ReadArrayHeader()
//	if err != nil {
//		err = msgp.WrapError(err)
//		return
//	}
//	if zb0001 != 7 {
//		err = msgp.ArrayError{Wanted: 7, Got: zb0001}
//		return
//	}
//	z.TokenID, err = dc.ReadInt32()
//	if err != nil {
//		err = msgp.WrapError(err, "TokenID")
//		return
//	}
//	z.Name, err = dc.ReadString()
//	if err != nil {
//		err = msgp.WrapError(err, "Name")
//		return
//	}
//	z.Symbol, err = dc.ReadString()
//	if err != nil {
//		err = msgp.WrapError(err, "Symbol")
//		return
//	}
//	err = z.Issuer.DecodeMsg(dc)
//	if err != nil {
//		err = msgp.WrapError(err, "Issuer")
//		return
//	}
//	z.ReIssuable, err = dc.ReadBool()
//	if err != nil {
//		err = msgp.WrapError(err, "ReIssuable")
//		return
//	}
//	var zb0002 uint32
//	zb0002, err = dc.ReadArrayHeader()
//	if err != nil {
//		err = msgp.WrapError(err, "Issues")
//		return
//	}
//	if cap(z.Issues) >= int(zb0002) {
//		z.Issues = (z.Issues)[:zb0002]
//	} else {
//		z.Issues = make([]*math.BigInt, zb0002)
//	}
//	for za0001 := range z.Issues {
//		if dc.IsNil() {
//			err = dc.ReadNil()
//			if err != nil {
//				err = msgp.WrapError(err, "Issues", za0001)
//				return
//			}
//			z.Issues[za0001] = nil
//		} else {
//			if z.Issues[za0001] == nil {
//				z.Issues[za0001] = new(math.BigInt)
//			}
//			err = z.Issues[za0001].DecodeMsg(dc)
//			if err != nil {
//				err = msgp.WrapError(err, "Issues", za0001)
//				return
//			}
//		}
//	}
//	z.Destroyed, err = dc.ReadBool()
//	if err != nil {
//		err = msgp.WrapError(err, "Destroyed")
//		return
//	}
//	return
//}
//
//// EncodeMsg implements msgp.Encodable
//func (z *TokenObject) EncodeMsg(en *msgp.Writer) (err error) {
//	// array header, size 7
//	err = en.Append(0x97)
//	if err != nil {
//		return
//	}
//	err = en.WriteInt32(z.TokenID)
//	if err != nil {
//		err = msgp.WrapError(err, "TokenID")
//		return
//	}
//	err = en.WriteString(z.Name)
//	if err != nil {
//		err = msgp.WrapError(err, "Name")
//		return
//	}
//	err = en.WriteString(z.Symbol)
//	if err != nil {
//		err = msgp.WrapError(err, "Symbol")
//		return
//	}
//	err = z.Issuer.EncodeMsg(en)
//	if err != nil {
//		err = msgp.WrapError(err, "Issuer")
//		return
//	}
//	err = en.WriteBool(z.ReIssuable)
//	if err != nil {
//		err = msgp.WrapError(err, "ReIssuable")
//		return
//	}
//	err = en.WriteArrayHeader(uint32(len(z.Issues)))
//	if err != nil {
//		err = msgp.WrapError(err, "Issues")
//		return
//	}
//	for za0001 := range z.Issues {
//		if z.Issues[za0001] == nil {
//			err = en.WriteNil()
//			if err != nil {
//				return
//			}
//		} else {
//			err = z.Issues[za0001].EncodeMsg(en)
//			if err != nil {
//				err = msgp.WrapError(err, "Issues", za0001)
//				return
//			}
//		}
//	}
//	err = en.WriteBool(z.Destroyed)
//	if err != nil {
//		err = msgp.WrapError(err, "Destroyed")
//		return
//	}
//	return
//}
//
//// MarshalMsg implements msgp.Marshaler
//func (z *TokenObject) MarshalMsg(b []byte) (o []byte, err error) {
//	o = msgp.Require(b, z.Msgsize())
//	// array header, size 7
//	o = append(o, 0x97)
//	o = msgp.AppendInt32(o, z.TokenID)
//	o = msgp.AppendString(o, z.Name)
//	o = msgp.AppendString(o, z.Symbol)
//	o, err = z.Issuer.MarshalMsg(o)
//	if err != nil {
//		err = msgp.WrapError(err, "Issuer")
//		return
//	}
//	o = msgp.AppendBool(o, z.ReIssuable)
//	o = msgp.AppendArrayHeader(o, uint32(len(z.Issues)))
//	for za0001 := range z.Issues {
//		if z.Issues[za0001] == nil {
//			o = msgp.AppendNil(o)
//		} else {
//			o, err = z.Issues[za0001].MarshalMsg(o)
//			if err != nil {
//				err = msgp.WrapError(err, "Issues", za0001)
//				return
//			}
//		}
//	}
//	o = msgp.AppendBool(o, z.Destroyed)
//	return
//}
//
//// UnmarshalMsg implements msgp.Unmarshaler
//func (z *TokenObject) UnmarshalMsg(bts []byte) (o []byte, err error) {
//	var zb0001 uint32
//	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
//	if err != nil {
//		err = msgp.WrapError(err)
//		return
//	}
//	if zb0001 != 7 {
//		err = msgp.ArrayError{Wanted: 7, Got: zb0001}
//		return
//	}
//	z.TokenID, bts, err = msgp.ReadInt32Bytes(bts)
//	if err != nil {
//		err = msgp.WrapError(err, "TokenID")
//		return
//	}
//	z.Name, bts, err = msgp.ReadStringBytes(bts)
//	if err != nil {
//		err = msgp.WrapError(err, "Name")
//		return
//	}
//	z.Symbol, bts, err = msgp.ReadStringBytes(bts)
//	if err != nil {
//		err = msgp.WrapError(err, "Symbol")
//		return
//	}
//	bts, err = z.Issuer.UnmarshalMsg(bts)
//	if err != nil {
//		err = msgp.WrapError(err, "Issuer")
//		return
//	}
//	z.ReIssuable, bts, err = msgp.ReadBoolBytes(bts)
//	if err != nil {
//		err = msgp.WrapError(err, "ReIssuable")
//		return
//	}
//	var zb0002 uint32
//	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
//	if err != nil {
//		err = msgp.WrapError(err, "Issues")
//		return
//	}
//	if cap(z.Issues) >= int(zb0002) {
//		z.Issues = (z.Issues)[:zb0002]
//	} else {
//		z.Issues = make([]*math.BigInt, zb0002)
//	}
//	for za0001 := range z.Issues {
//		if msgp.IsNil(bts) {
//			bts, err = msgp.ReadNilBytes(bts)
//			if err != nil {
//				return
//			}
//			z.Issues[za0001] = nil
//		} else {
//			if z.Issues[za0001] == nil {
//				z.Issues[za0001] = new(math.BigInt)
//			}
//			bts, err = z.Issues[za0001].UnmarshalMsg(bts)
//			if err != nil {
//				err = msgp.WrapError(err, "Issues", za0001)
//				return
//			}
//		}
//	}
//	z.Destroyed, bts, err = msgp.ReadBoolBytes(bts)
//	if err != nil {
//		err = msgp.WrapError(err, "Destroyed")
//		return
//	}
//	o = bts
//	return
//}
//
//// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
//func (z *TokenObject) Msgsize() (s int) {
//	s = 1 + msgp.Int32Size + msgp.StringPrefixSize + len(z.Name) + msgp.StringPrefixSize + len(z.Symbol) + z.Issuer.Msgsize() + msgp.BoolSize + msgp.ArrayHeaderSize
//	for za0001 := range z.Issues {
//		if z.Issues[za0001] == nil {
//			s += msgp.NilSize
//		} else {
//			s += z.Issues[za0001].Msgsize()
//		}
//	}
//	s += msgp.BoolSize
//	return
//}
