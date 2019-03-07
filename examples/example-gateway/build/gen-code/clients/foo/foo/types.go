// Code generated by thriftrw v1.16.1. DO NOT EDIT.
// @generated

package foo

import (
	"errors"
	"fmt"
	"github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/foo/base/base"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

type FooException struct {
	Teapot string `json:"teapot,required"`
}

// ToWire translates a FooException struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *FooException) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Teapot), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a FooException struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a FooException struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v FooException
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *FooException) FromWire(w wire.Value) error {
	var err error

	teapotIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Teapot, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				teapotIsSet = true
			}
		}
	}

	if !teapotIsSet {
		return errors.New("field Teapot of FooException is required")
	}

	return nil
}

// String returns a readable string representation of a FooException
// struct.
func (v *FooException) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Teapot: %v", v.Teapot)
	i++

	return fmt.Sprintf("FooException{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this FooException match the
// provided FooException.
//
// This function performs a deep comparison.
func (v *FooException) Equals(rhs *FooException) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Teapot == rhs.Teapot) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of FooException.
func (v *FooException) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("teapot", v.Teapot)
	return err
}

// GetTeapot returns the value of Teapot if it is set or its
// zero value if it is unset.
func (v *FooException) GetTeapot() (o string) {
	if v != nil {
		o = v.Teapot
	}
	return
}

func (v *FooException) Error() string {
	return v.String()
}

type FooName struct {
	Name *string `json:"name,omitempty"`
}

// ToWire translates a FooName struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *FooName) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Name != nil {
		w, err = wire.NewValueString(*(v.Name)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a FooName struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a FooName struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v FooName
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *FooName) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Name = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a FooName
// struct.
func (v *FooName) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Name != nil {
		fields[i] = fmt.Sprintf("Name: %v", *(v.Name))
		i++
	}

	return fmt.Sprintf("FooName{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this FooName match the
// provided FooName.
//
// This function performs a deep comparison.
func (v *FooName) Equals(rhs *FooName) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Name, rhs.Name) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of FooName.
func (v *FooName) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Name != nil {
		enc.AddString("name", *v.Name)
	}
	return err
}

// GetName returns the value of Name if it is set or its
// zero value if it is unset.
func (v *FooName) GetName() (o string) {
	if v != nil && v.Name != nil {
		return *v.Name
	}

	return
}

// IsSetName returns true if Name is not nil.
func (v *FooName) IsSetName() bool {
	return v != nil && v.Name != nil
}

type FooStruct struct {
	FooString string            `json:"fooString,required"`
	FooI32    *int32            `json:"fooI32,omitempty"`
	FooI16    *int16            `json:"fooI16,omitempty"`
	FooDouble *float64          `json:"fooDouble,omitempty"`
	FooBool   *bool             `json:"fooBool,omitempty"`
	FooMap    map[string]string `json:"fooMap,omitempty"`
	Message   *base.Message     `json:"message,omitempty"`
}

type _Map_String_String_MapItemList map[string]string

func (m _Map_String_String_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		kw, err := wire.NewValueString(k), error(nil)
		if err != nil {
			return err
		}

		vw, err := wire.NewValueString(v), error(nil)
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_String_String_MapItemList) Size() int {
	return len(m)
}

func (_Map_String_String_MapItemList) KeyType() wire.Type {
	return wire.TBinary
}

func (_Map_String_String_MapItemList) ValueType() wire.Type {
	return wire.TBinary
}

func (_Map_String_String_MapItemList) Close() {}

// ToWire translates a FooStruct struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *FooStruct) ToWire() (wire.Value, error) {
	var (
		fields [7]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.FooString), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.FooI32 != nil {
		w, err = wire.NewValueI32(*(v.FooI32)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.FooI16 != nil {
		w, err = wire.NewValueI16(*(v.FooI16)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.FooDouble != nil {
		w, err = wire.NewValueDouble(*(v.FooDouble)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	if v.FooBool != nil {
		w, err = wire.NewValueBool(*(v.FooBool)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 5, Value: w}
		i++
	}
	if v.FooMap != nil {
		w, err = wire.NewValueMap(_Map_String_String_MapItemList(v.FooMap)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 6, Value: w}
		i++
	}
	if v.Message != nil {
		w, err = v.Message.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 7, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Map_String_String_Read(m wire.MapItemList) (map[string]string, error) {
	if m.KeyType() != wire.TBinary {
		return nil, nil
	}

	if m.ValueType() != wire.TBinary {
		return nil, nil
	}

	o := make(map[string]string, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetString(), error(nil)
		if err != nil {
			return err
		}

		v, err := x.Value.GetString(), error(nil)
		if err != nil {
			return err
		}

		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

func _Message_Read(w wire.Value) (*base.Message, error) {
	var v base.Message
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a FooStruct struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a FooStruct struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v FooStruct
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *FooStruct) FromWire(w wire.Value) error {
	var err error

	fooStringIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.FooString, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				fooStringIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TI32 {
				var x int32
				x, err = field.Value.GetI32(), error(nil)
				v.FooI32 = &x
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TI16 {
				var x int16
				x, err = field.Value.GetI16(), error(nil)
				v.FooI16 = &x
				if err != nil {
					return err
				}

			}
		case 4:
			if field.Value.Type() == wire.TDouble {
				var x float64
				x, err = field.Value.GetDouble(), error(nil)
				v.FooDouble = &x
				if err != nil {
					return err
				}

			}
		case 5:
			if field.Value.Type() == wire.TBool {
				var x bool
				x, err = field.Value.GetBool(), error(nil)
				v.FooBool = &x
				if err != nil {
					return err
				}

			}
		case 6:
			if field.Value.Type() == wire.TMap {
				v.FooMap, err = _Map_String_String_Read(field.Value.GetMap())
				if err != nil {
					return err
				}

			}
		case 7:
			if field.Value.Type() == wire.TStruct {
				v.Message, err = _Message_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	if !fooStringIsSet {
		return errors.New("field FooString of FooStruct is required")
	}

	return nil
}

// String returns a readable string representation of a FooStruct
// struct.
func (v *FooStruct) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [7]string
	i := 0
	fields[i] = fmt.Sprintf("FooString: %v", v.FooString)
	i++
	if v.FooI32 != nil {
		fields[i] = fmt.Sprintf("FooI32: %v", *(v.FooI32))
		i++
	}
	if v.FooI16 != nil {
		fields[i] = fmt.Sprintf("FooI16: %v", *(v.FooI16))
		i++
	}
	if v.FooDouble != nil {
		fields[i] = fmt.Sprintf("FooDouble: %v", *(v.FooDouble))
		i++
	}
	if v.FooBool != nil {
		fields[i] = fmt.Sprintf("FooBool: %v", *(v.FooBool))
		i++
	}
	if v.FooMap != nil {
		fields[i] = fmt.Sprintf("FooMap: %v", v.FooMap)
		i++
	}
	if v.Message != nil {
		fields[i] = fmt.Sprintf("Message: %v", v.Message)
		i++
	}

	return fmt.Sprintf("FooStruct{%v}", strings.Join(fields[:i], ", "))
}

func _I32_EqualsPtr(lhs, rhs *int32) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _I16_EqualsPtr(lhs, rhs *int16) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _Double_EqualsPtr(lhs, rhs *float64) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _Bool_EqualsPtr(lhs, rhs *bool) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _Map_String_String_Equals(lhs, rhs map[string]string) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for lk, lv := range lhs {
		rv, ok := rhs[lk]
		if !ok {
			return false
		}
		if !(lv == rv) {
			return false
		}
	}
	return true
}

// Equals returns true if all the fields of this FooStruct match the
// provided FooStruct.
//
// This function performs a deep comparison.
func (v *FooStruct) Equals(rhs *FooStruct) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.FooString == rhs.FooString) {
		return false
	}
	if !_I32_EqualsPtr(v.FooI32, rhs.FooI32) {
		return false
	}
	if !_I16_EqualsPtr(v.FooI16, rhs.FooI16) {
		return false
	}
	if !_Double_EqualsPtr(v.FooDouble, rhs.FooDouble) {
		return false
	}
	if !_Bool_EqualsPtr(v.FooBool, rhs.FooBool) {
		return false
	}
	if !((v.FooMap == nil && rhs.FooMap == nil) || (v.FooMap != nil && rhs.FooMap != nil && _Map_String_String_Equals(v.FooMap, rhs.FooMap))) {
		return false
	}
	if !((v.Message == nil && rhs.Message == nil) || (v.Message != nil && rhs.Message != nil && v.Message.Equals(rhs.Message))) {
		return false
	}

	return true
}

type _Map_String_String_Zapper map[string]string

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of _Map_String_String_Zapper.
func (m _Map_String_String_Zapper) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	for k, v := range m {
		enc.AddString((string)(k), v)
	}
	return err
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of FooStruct.
func (v *FooStruct) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("fooString", v.FooString)
	if v.FooI32 != nil {
		enc.AddInt32("fooI32", *v.FooI32)
	}
	if v.FooI16 != nil {
		enc.AddInt16("fooI16", *v.FooI16)
	}
	if v.FooDouble != nil {
		enc.AddFloat64("fooDouble", *v.FooDouble)
	}
	if v.FooBool != nil {
		enc.AddBool("fooBool", *v.FooBool)
	}
	if v.FooMap != nil {
		err = multierr.Append(err, enc.AddObject("fooMap", (_Map_String_String_Zapper)(v.FooMap)))
	}
	if v.Message != nil {
		err = multierr.Append(err, enc.AddObject("message", v.Message))
	}
	return err
}

// GetFooString returns the value of FooString if it is set or its
// zero value if it is unset.
func (v *FooStruct) GetFooString() (o string) {
	if v != nil {
		o = v.FooString
	}
	return
}

// GetFooI32 returns the value of FooI32 if it is set or its
// zero value if it is unset.
func (v *FooStruct) GetFooI32() (o int32) {
	if v != nil && v.FooI32 != nil {
		return *v.FooI32
	}

	return
}

// IsSetFooI32 returns true if FooI32 is not nil.
func (v *FooStruct) IsSetFooI32() bool {
	return v != nil && v.FooI32 != nil
}

// GetFooI16 returns the value of FooI16 if it is set or its
// zero value if it is unset.
func (v *FooStruct) GetFooI16() (o int16) {
	if v != nil && v.FooI16 != nil {
		return *v.FooI16
	}

	return
}

// IsSetFooI16 returns true if FooI16 is not nil.
func (v *FooStruct) IsSetFooI16() bool {
	return v != nil && v.FooI16 != nil
}

// GetFooDouble returns the value of FooDouble if it is set or its
// zero value if it is unset.
func (v *FooStruct) GetFooDouble() (o float64) {
	if v != nil && v.FooDouble != nil {
		return *v.FooDouble
	}

	return
}

// IsSetFooDouble returns true if FooDouble is not nil.
func (v *FooStruct) IsSetFooDouble() bool {
	return v != nil && v.FooDouble != nil
}

// GetFooBool returns the value of FooBool if it is set or its
// zero value if it is unset.
func (v *FooStruct) GetFooBool() (o bool) {
	if v != nil && v.FooBool != nil {
		return *v.FooBool
	}

	return
}

// IsSetFooBool returns true if FooBool is not nil.
func (v *FooStruct) IsSetFooBool() bool {
	return v != nil && v.FooBool != nil
}

// GetFooMap returns the value of FooMap if it is set or its
// zero value if it is unset.
func (v *FooStruct) GetFooMap() (o map[string]string) {
	if v != nil && v.FooMap != nil {
		return v.FooMap
	}

	return
}

// IsSetFooMap returns true if FooMap is not nil.
func (v *FooStruct) IsSetFooMap() bool {
	return v != nil && v.FooMap != nil
}

// GetMessage returns the value of Message if it is set or its
// zero value if it is unset.
func (v *FooStruct) GetMessage() (o *base.Message) {
	if v != nil && v.Message != nil {
		return v.Message
	}

	return
}

// IsSetMessage returns true if Message is not nil.
func (v *FooStruct) IsSetMessage() bool {
	return v != nil && v.Message != nil
}
