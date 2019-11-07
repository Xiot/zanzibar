// Code generated by thriftrw v1.20.1. DO NOT EDIT.
// @generated

package base

import (
	errors "errors"
	fmt "fmt"
	strings "strings"

	multierr "go.uber.org/multierr"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
)

type BazResponse struct {
	Message string `json:"message,required"`
}

// ToWire translates a BazResponse struct into a Thrift-level intermediate
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
func (v *BazResponse) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a BazResponse struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a BazResponse struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v BazResponse
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *BazResponse) FromWire(w wire.Value) error {
	var err error

	messageIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		}
	}

	if !messageIsSet {
		return errors.New("field Message of BazResponse is required")
	}

	return nil
}

// String returns a readable string representation of a BazResponse
// struct.
func (v *BazResponse) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++

	return fmt.Sprintf("BazResponse{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this BazResponse match the
// provided BazResponse.
//
// This function performs a deep comparison.
func (v *BazResponse) Equals(rhs *BazResponse) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Message == rhs.Message) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of BazResponse.
func (v *BazResponse) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("message", v.Message)
	return err
}

// GetMessage returns the value of Message if it is set or its
// zero value if it is unset.
func (v *BazResponse) GetMessage() (o string) {
	if v != nil {
		o = v.Message
	}
	return
}

type NestHeaders struct {
	UUID  string  `json:"UUID,required"`
	Token *string `json:"token,omitempty"`
}

// ToWire translates a NestHeaders struct into a Thrift-level intermediate
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
func (v *NestHeaders) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.UUID), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Token != nil {
		w, err = wire.NewValueString(*(v.Token)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a NestHeaders struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a NestHeaders struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v NestHeaders
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *NestHeaders) FromWire(w wire.Value) error {
	var err error

	UUIDIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.UUID, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				UUIDIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Token = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if !UUIDIsSet {
		return errors.New("field UUID of NestHeaders is required")
	}

	return nil
}

// String returns a readable string representation of a NestHeaders
// struct.
func (v *NestHeaders) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("UUID: %v", v.UUID)
	i++
	if v.Token != nil {
		fields[i] = fmt.Sprintf("Token: %v", *(v.Token))
		i++
	}

	return fmt.Sprintf("NestHeaders{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this NestHeaders match the
// provided NestHeaders.
//
// This function performs a deep comparison.
func (v *NestHeaders) Equals(rhs *NestHeaders) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.UUID == rhs.UUID) {
		return false
	}
	if !_String_EqualsPtr(v.Token, rhs.Token) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of NestHeaders.
func (v *NestHeaders) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("UUID", v.UUID)
	if v.Token != nil {
		enc.AddString("token", *v.Token)
	}
	return err
}

// GetUUID returns the value of UUID if it is set or its
// zero value if it is unset.
func (v *NestHeaders) GetUUID() (o string) {
	if v != nil {
		o = v.UUID
	}
	return
}

// GetToken returns the value of Token if it is set or its
// zero value if it is unset.
func (v *NestHeaders) GetToken() (o string) {
	if v != nil && v.Token != nil {
		return *v.Token
	}

	return
}

// IsSetToken returns true if Token is not nil.
func (v *NestHeaders) IsSetToken() bool {
	return v != nil && v.Token != nil
}

type NestedStruct struct {
	Msg   string `json:"msg,required"`
	Check *int32 `json:"check,omitempty"`
}

// ToWire translates a NestedStruct struct into a Thrift-level intermediate
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
func (v *NestedStruct) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Msg), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Check != nil {
		w, err = wire.NewValueI32(*(v.Check)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a NestedStruct struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a NestedStruct struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v NestedStruct
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *NestedStruct) FromWire(w wire.Value) error {
	var err error

	msgIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Msg, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				msgIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TI32 {
				var x int32
				x, err = field.Value.GetI32(), error(nil)
				v.Check = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if !msgIsSet {
		return errors.New("field Msg of NestedStruct is required")
	}

	return nil
}

// String returns a readable string representation of a NestedStruct
// struct.
func (v *NestedStruct) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Msg: %v", v.Msg)
	i++
	if v.Check != nil {
		fields[i] = fmt.Sprintf("Check: %v", *(v.Check))
		i++
	}

	return fmt.Sprintf("NestedStruct{%v}", strings.Join(fields[:i], ", "))
}

func _I32_EqualsPtr(lhs, rhs *int32) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this NestedStruct match the
// provided NestedStruct.
//
// This function performs a deep comparison.
func (v *NestedStruct) Equals(rhs *NestedStruct) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Msg == rhs.Msg) {
		return false
	}
	if !_I32_EqualsPtr(v.Check, rhs.Check) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of NestedStruct.
func (v *NestedStruct) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("msg", v.Msg)
	if v.Check != nil {
		enc.AddInt32("check", *v.Check)
	}
	return err
}

// GetMsg returns the value of Msg if it is set or its
// zero value if it is unset.
func (v *NestedStruct) GetMsg() (o string) {
	if v != nil {
		o = v.Msg
	}
	return
}

// GetCheck returns the value of Check if it is set or its
// zero value if it is unset.
func (v *NestedStruct) GetCheck() (o int32) {
	if v != nil && v.Check != nil {
		return *v.Check
	}

	return
}

// IsSetCheck returns true if Check is not nil.
func (v *NestedStruct) IsSetCheck() bool {
	return v != nil && v.Check != nil
}

type ServerErr struct {
	Message string `json:"message,required"`
}

// ToWire translates a ServerErr struct into a Thrift-level intermediate
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
func (v *ServerErr) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ServerErr struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ServerErr struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ServerErr
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ServerErr) FromWire(w wire.Value) error {
	var err error

	messageIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		}
	}

	if !messageIsSet {
		return errors.New("field Message of ServerErr is required")
	}

	return nil
}

// String returns a readable string representation of a ServerErr
// struct.
func (v *ServerErr) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++

	return fmt.Sprintf("ServerErr{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ServerErr match the
// provided ServerErr.
//
// This function performs a deep comparison.
func (v *ServerErr) Equals(rhs *ServerErr) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Message == rhs.Message) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ServerErr.
func (v *ServerErr) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("message", v.Message)
	return err
}

// GetMessage returns the value of Message if it is set or its
// zero value if it is unset.
func (v *ServerErr) GetMessage() (o string) {
	if v != nil {
		o = v.Message
	}
	return
}

func (v *ServerErr) Error() string {
	return v.String()
}

type TransHeaders struct {
	W1 *Wrapped `json:"w1,required"`
	W2 *Wrapped `json:"w2,omitempty"`
}

// ToWire translates a TransHeaders struct into a Thrift-level intermediate
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
func (v *TransHeaders) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.W1 == nil {
		return w, errors.New("field W1 of TransHeaders is required")
	}
	w, err = v.W1.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.W2 != nil {
		w, err = v.W2.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Wrapped_Read(w wire.Value) (*Wrapped, error) {
	var v Wrapped
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a TransHeaders struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a TransHeaders struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v TransHeaders
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *TransHeaders) FromWire(w wire.Value) error {
	var err error

	w1IsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.W1, err = _Wrapped_Read(field.Value)
				if err != nil {
					return err
				}
				w1IsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.W2, err = _Wrapped_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	if !w1IsSet {
		return errors.New("field W1 of TransHeaders is required")
	}

	return nil
}

// String returns a readable string representation of a TransHeaders
// struct.
func (v *TransHeaders) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("W1: %v", v.W1)
	i++
	if v.W2 != nil {
		fields[i] = fmt.Sprintf("W2: %v", v.W2)
		i++
	}

	return fmt.Sprintf("TransHeaders{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this TransHeaders match the
// provided TransHeaders.
//
// This function performs a deep comparison.
func (v *TransHeaders) Equals(rhs *TransHeaders) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !v.W1.Equals(rhs.W1) {
		return false
	}
	if !((v.W2 == nil && rhs.W2 == nil) || (v.W2 != nil && rhs.W2 != nil && v.W2.Equals(rhs.W2))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of TransHeaders.
func (v *TransHeaders) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	err = multierr.Append(err, enc.AddObject("w1", v.W1))
	if v.W2 != nil {
		err = multierr.Append(err, enc.AddObject("w2", v.W2))
	}
	return err
}

// GetW1 returns the value of W1 if it is set or its
// zero value if it is unset.
func (v *TransHeaders) GetW1() (o *Wrapped) {
	if v != nil {
		o = v.W1
	}
	return
}

// IsSetW1 returns true if W1 is not nil.
func (v *TransHeaders) IsSetW1() bool {
	return v != nil && v.W1 != nil
}

// GetW2 returns the value of W2 if it is set or its
// zero value if it is unset.
func (v *TransHeaders) GetW2() (o *Wrapped) {
	if v != nil && v.W2 != nil {
		return v.W2
	}

	return
}

// IsSetW2 returns true if W2 is not nil.
func (v *TransHeaders) IsSetW2() bool {
	return v != nil && v.W2 != nil
}

type TransStruct struct {
	Message string        `json:"message,required"`
	Driver  *NestedStruct `json:"driver,omitempty"`
	Rider   *NestedStruct `json:"rider,required"`
}

// ToWire translates a TransStruct struct into a Thrift-level intermediate
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
func (v *TransStruct) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Driver != nil {
		w, err = v.Driver.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.Rider == nil {
		return w, errors.New("field Rider of TransStruct is required")
	}
	w, err = v.Rider.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 3, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _NestedStruct_Read(w wire.Value) (*NestedStruct, error) {
	var v NestedStruct
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a TransStruct struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a TransStruct struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v TransStruct
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *TransStruct) FromWire(w wire.Value) error {
	var err error

	messageIsSet := false

	riderIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.Driver, err = _NestedStruct_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TStruct {
				v.Rider, err = _NestedStruct_Read(field.Value)
				if err != nil {
					return err
				}
				riderIsSet = true
			}
		}
	}

	if !messageIsSet {
		return errors.New("field Message of TransStruct is required")
	}

	if !riderIsSet {
		return errors.New("field Rider of TransStruct is required")
	}

	return nil
}

// String returns a readable string representation of a TransStruct
// struct.
func (v *TransStruct) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	if v.Driver != nil {
		fields[i] = fmt.Sprintf("Driver: %v", v.Driver)
		i++
	}
	fields[i] = fmt.Sprintf("Rider: %v", v.Rider)
	i++

	return fmt.Sprintf("TransStruct{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this TransStruct match the
// provided TransStruct.
//
// This function performs a deep comparison.
func (v *TransStruct) Equals(rhs *TransStruct) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Message == rhs.Message) {
		return false
	}
	if !((v.Driver == nil && rhs.Driver == nil) || (v.Driver != nil && rhs.Driver != nil && v.Driver.Equals(rhs.Driver))) {
		return false
	}
	if !v.Rider.Equals(rhs.Rider) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of TransStruct.
func (v *TransStruct) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("message", v.Message)
	if v.Driver != nil {
		err = multierr.Append(err, enc.AddObject("driver", v.Driver))
	}
	err = multierr.Append(err, enc.AddObject("rider", v.Rider))
	return err
}

// GetMessage returns the value of Message if it is set or its
// zero value if it is unset.
func (v *TransStruct) GetMessage() (o string) {
	if v != nil {
		o = v.Message
	}
	return
}

// GetDriver returns the value of Driver if it is set or its
// zero value if it is unset.
func (v *TransStruct) GetDriver() (o *NestedStruct) {
	if v != nil && v.Driver != nil {
		return v.Driver
	}

	return
}

// IsSetDriver returns true if Driver is not nil.
func (v *TransStruct) IsSetDriver() bool {
	return v != nil && v.Driver != nil
}

// GetRider returns the value of Rider if it is set or its
// zero value if it is unset.
func (v *TransStruct) GetRider() (o *NestedStruct) {
	if v != nil {
		o = v.Rider
	}
	return
}

// IsSetRider returns true if Rider is not nil.
func (v *TransStruct) IsSetRider() bool {
	return v != nil && v.Rider != nil
}

type UUID string

// UUIDPtr returns a pointer to a UUID
func (v UUID) Ptr() *UUID {
	return &v
}

// ToWire translates UUID into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
func (v UUID) ToWire() (wire.Value, error) {
	x := (string)(v)
	return wire.NewValueString(x), error(nil)
}

// String returns a readable string representation of UUID.
func (v UUID) String() string {
	x := (string)(v)
	return fmt.Sprint(x)
}

// FromWire deserializes UUID from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
func (v *UUID) FromWire(w wire.Value) error {
	x, err := w.GetString(), error(nil)
	*v = (UUID)(x)
	return err
}

// Equals returns true if this UUID is equal to the provided
// UUID.
func (lhs UUID) Equals(rhs UUID) bool {
	return ((string)(lhs) == (string)(rhs))
}

type Wrapped struct {
	N1 *NestHeaders `json:"n1,required"`
	N2 *NestHeaders `json:"n2,omitempty"`
}

// ToWire translates a Wrapped struct into a Thrift-level intermediate
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
func (v *Wrapped) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.N1 == nil {
		return w, errors.New("field N1 of Wrapped is required")
	}
	w, err = v.N1.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.N2 != nil {
		w, err = v.N2.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _NestHeaders_Read(w wire.Value) (*NestHeaders, error) {
	var v NestHeaders
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Wrapped struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Wrapped struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Wrapped
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Wrapped) FromWire(w wire.Value) error {
	var err error

	n1IsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.N1, err = _NestHeaders_Read(field.Value)
				if err != nil {
					return err
				}
				n1IsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.N2, err = _NestHeaders_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	if !n1IsSet {
		return errors.New("field N1 of Wrapped is required")
	}

	return nil
}

// String returns a readable string representation of a Wrapped
// struct.
func (v *Wrapped) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("N1: %v", v.N1)
	i++
	if v.N2 != nil {
		fields[i] = fmt.Sprintf("N2: %v", v.N2)
		i++
	}

	return fmt.Sprintf("Wrapped{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Wrapped match the
// provided Wrapped.
//
// This function performs a deep comparison.
func (v *Wrapped) Equals(rhs *Wrapped) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !v.N1.Equals(rhs.N1) {
		return false
	}
	if !((v.N2 == nil && rhs.N2 == nil) || (v.N2 != nil && rhs.N2 != nil && v.N2.Equals(rhs.N2))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Wrapped.
func (v *Wrapped) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	err = multierr.Append(err, enc.AddObject("n1", v.N1))
	if v.N2 != nil {
		err = multierr.Append(err, enc.AddObject("n2", v.N2))
	}
	return err
}

// GetN1 returns the value of N1 if it is set or its
// zero value if it is unset.
func (v *Wrapped) GetN1() (o *NestHeaders) {
	if v != nil {
		o = v.N1
	}
	return
}

// IsSetN1 returns true if N1 is not nil.
func (v *Wrapped) IsSetN1() bool {
	return v != nil && v.N1 != nil
}

// GetN2 returns the value of N2 if it is set or its
// zero value if it is unset.
func (v *Wrapped) GetN2() (o *NestHeaders) {
	if v != nil && v.N2 != nil {
		return v.N2
	}

	return
}

// IsSetN2 returns true if N2 is not nil.
func (v *Wrapped) IsSetN2() bool {
	return v != nil && v.N2 != nil
}
