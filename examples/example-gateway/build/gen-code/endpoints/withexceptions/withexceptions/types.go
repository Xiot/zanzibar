// Code generated by thriftrw v1.16.1. DO NOT EDIT.
// @generated

package withexceptions

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

type EndpointExceptionType1 struct {
	Message1 string `json:"message1,required"`
}

// ToWire translates a EndpointExceptionType1 struct into a Thrift-level intermediate
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
func (v *EndpointExceptionType1) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Message1), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a EndpointExceptionType1 struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a EndpointExceptionType1 struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v EndpointExceptionType1
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *EndpointExceptionType1) FromWire(w wire.Value) error {
	var err error

	message1IsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Message1, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				message1IsSet = true
			}
		}
	}

	if !message1IsSet {
		return errors.New("field Message1 of EndpointExceptionType1 is required")
	}

	return nil
}

// String returns a readable string representation of a EndpointExceptionType1
// struct.
func (v *EndpointExceptionType1) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Message1: %v", v.Message1)
	i++

	return fmt.Sprintf("EndpointExceptionType1{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this EndpointExceptionType1 match the
// provided EndpointExceptionType1.
//
// This function performs a deep comparison.
func (v *EndpointExceptionType1) Equals(rhs *EndpointExceptionType1) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Message1 == rhs.Message1) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of EndpointExceptionType1.
func (v *EndpointExceptionType1) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("message1", v.Message1)
	return err
}

// GetMessage1 returns the value of Message1 if it is set or its
// zero value if it is unset.
func (v *EndpointExceptionType1) GetMessage1() (o string) {
	if v != nil {
		o = v.Message1
	}
	return
}

func (v *EndpointExceptionType1) Error() string {
	return v.String()
}

type EndpointExceptionType2 struct {
	Message2 string `json:"message2,required"`
}

// ToWire translates a EndpointExceptionType2 struct into a Thrift-level intermediate
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
func (v *EndpointExceptionType2) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Message2), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a EndpointExceptionType2 struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a EndpointExceptionType2 struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v EndpointExceptionType2
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *EndpointExceptionType2) FromWire(w wire.Value) error {
	var err error

	message2IsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Message2, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				message2IsSet = true
			}
		}
	}

	if !message2IsSet {
		return errors.New("field Message2 of EndpointExceptionType2 is required")
	}

	return nil
}

// String returns a readable string representation of a EndpointExceptionType2
// struct.
func (v *EndpointExceptionType2) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Message2: %v", v.Message2)
	i++

	return fmt.Sprintf("EndpointExceptionType2{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this EndpointExceptionType2 match the
// provided EndpointExceptionType2.
//
// This function performs a deep comparison.
func (v *EndpointExceptionType2) Equals(rhs *EndpointExceptionType2) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Message2 == rhs.Message2) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of EndpointExceptionType2.
func (v *EndpointExceptionType2) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("message2", v.Message2)
	return err
}

// GetMessage2 returns the value of Message2 if it is set or its
// zero value if it is unset.
func (v *EndpointExceptionType2) GetMessage2() (o string) {
	if v != nil {
		o = v.Message2
	}
	return
}

func (v *EndpointExceptionType2) Error() string {
	return v.String()
}

type Response struct {
}

// ToWire translates a Response struct into a Thrift-level intermediate
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
func (v *Response) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Response struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Response struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Response
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Response) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a Response
// struct.
func (v *Response) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("Response{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Response match the
// provided Response.
//
// This function performs a deep comparison.
func (v *Response) Equals(rhs *Response) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Response.
func (v *Response) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}
