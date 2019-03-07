// Code generated by thriftrw v1.16.1. DO NOT EDIT.
// @generated

package baz

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// SimpleService_SillyNoop_Args represents the arguments for the SimpleService.sillyNoop function.
//
// The arguments for sillyNoop are sent and received over the wire as this struct.
type SimpleService_SillyNoop_Args struct {
}

// ToWire translates a SimpleService_SillyNoop_Args struct into a Thrift-level intermediate
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
func (v *SimpleService_SillyNoop_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a SimpleService_SillyNoop_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SimpleService_SillyNoop_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SimpleService_SillyNoop_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SimpleService_SillyNoop_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a SimpleService_SillyNoop_Args
// struct.
func (v *SimpleService_SillyNoop_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("SimpleService_SillyNoop_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SimpleService_SillyNoop_Args match the
// provided SimpleService_SillyNoop_Args.
//
// This function performs a deep comparison.
func (v *SimpleService_SillyNoop_Args) Equals(rhs *SimpleService_SillyNoop_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of SimpleService_SillyNoop_Args.
func (v *SimpleService_SillyNoop_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "sillyNoop" for this struct.
func (v *SimpleService_SillyNoop_Args) MethodName() string {
	return "sillyNoop"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *SimpleService_SillyNoop_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// SimpleService_SillyNoop_Helper provides functions that aid in handling the
// parameters and return values of the SimpleService.sillyNoop
// function.
var SimpleService_SillyNoop_Helper = struct {
	// Args accepts the parameters of sillyNoop in-order and returns
	// the arguments struct for the function.
	Args func() *SimpleService_SillyNoop_Args

	// IsException returns true if the given error can be thrown
	// by sillyNoop.
	//
	// An error can be thrown by sillyNoop only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for sillyNoop
	// given the error returned by it. The provided error may
	// be nil if sillyNoop did not fail.
	//
	// This allows mapping errors returned by sillyNoop into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// sillyNoop
	//
	//   err := sillyNoop(args)
	//   result, err := SimpleService_SillyNoop_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from sillyNoop: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*SimpleService_SillyNoop_Result, error)

	// UnwrapResponse takes the result struct for sillyNoop
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if sillyNoop threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := SimpleService_SillyNoop_Helper.UnwrapResponse(result)
	UnwrapResponse func(*SimpleService_SillyNoop_Result) error
}{}

func init() {
	SimpleService_SillyNoop_Helper.Args = func() *SimpleService_SillyNoop_Args {
		return &SimpleService_SillyNoop_Args{}
	}

	SimpleService_SillyNoop_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *AuthErr:
			return true
		case *ServerErr:
			return true
		default:
			return false
		}
	}

	SimpleService_SillyNoop_Helper.WrapResponse = func(err error) (*SimpleService_SillyNoop_Result, error) {
		if err == nil {
			return &SimpleService_SillyNoop_Result{}, nil
		}

		switch e := err.(type) {
		case *AuthErr:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for SimpleService_SillyNoop_Result.AuthErr")
			}
			return &SimpleService_SillyNoop_Result{AuthErr: e}, nil
		case *ServerErr:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for SimpleService_SillyNoop_Result.ServerErr")
			}
			return &SimpleService_SillyNoop_Result{ServerErr: e}, nil
		}

		return nil, err
	}
	SimpleService_SillyNoop_Helper.UnwrapResponse = func(result *SimpleService_SillyNoop_Result) (err error) {
		if result.AuthErr != nil {
			err = result.AuthErr
			return
		}
		if result.ServerErr != nil {
			err = result.ServerErr
			return
		}
		return
	}

}

// SimpleService_SillyNoop_Result represents the result of a SimpleService.sillyNoop function call.
//
// The result of a sillyNoop execution is sent and received over the wire as this struct.
type SimpleService_SillyNoop_Result struct {
	AuthErr   *AuthErr   `json:"authErr,omitempty"`
	ServerErr *ServerErr `json:"serverErr,omitempty"`
}

// ToWire translates a SimpleService_SillyNoop_Result struct into a Thrift-level intermediate
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
func (v *SimpleService_SillyNoop_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.AuthErr != nil {
		w, err = v.AuthErr.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.ServerErr != nil {
		w, err = v.ServerErr.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	if i > 1 {
		return wire.Value{}, fmt.Errorf("SimpleService_SillyNoop_Result should have at most one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ServerErr_Read(w wire.Value) (*ServerErr, error) {
	var v ServerErr
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a SimpleService_SillyNoop_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SimpleService_SillyNoop_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SimpleService_SillyNoop_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SimpleService_SillyNoop_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.AuthErr, err = _AuthErr_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.ServerErr, err = _ServerErr_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.AuthErr != nil {
		count++
	}
	if v.ServerErr != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("SimpleService_SillyNoop_Result should have at most one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a SimpleService_SillyNoop_Result
// struct.
func (v *SimpleService_SillyNoop_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.AuthErr != nil {
		fields[i] = fmt.Sprintf("AuthErr: %v", v.AuthErr)
		i++
	}
	if v.ServerErr != nil {
		fields[i] = fmt.Sprintf("ServerErr: %v", v.ServerErr)
		i++
	}

	return fmt.Sprintf("SimpleService_SillyNoop_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SimpleService_SillyNoop_Result match the
// provided SimpleService_SillyNoop_Result.
//
// This function performs a deep comparison.
func (v *SimpleService_SillyNoop_Result) Equals(rhs *SimpleService_SillyNoop_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.AuthErr == nil && rhs.AuthErr == nil) || (v.AuthErr != nil && rhs.AuthErr != nil && v.AuthErr.Equals(rhs.AuthErr))) {
		return false
	}
	if !((v.ServerErr == nil && rhs.ServerErr == nil) || (v.ServerErr != nil && rhs.ServerErr != nil && v.ServerErr.Equals(rhs.ServerErr))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of SimpleService_SillyNoop_Result.
func (v *SimpleService_SillyNoop_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.AuthErr != nil {
		err = multierr.Append(err, enc.AddObject("authErr", v.AuthErr))
	}
	if v.ServerErr != nil {
		err = multierr.Append(err, enc.AddObject("serverErr", v.ServerErr))
	}
	return err
}

// GetAuthErr returns the value of AuthErr if it is set or its
// zero value if it is unset.
func (v *SimpleService_SillyNoop_Result) GetAuthErr() (o *AuthErr) {
	if v != nil && v.AuthErr != nil {
		return v.AuthErr
	}

	return
}

// IsSetAuthErr returns true if AuthErr is not nil.
func (v *SimpleService_SillyNoop_Result) IsSetAuthErr() bool {
	return v != nil && v.AuthErr != nil
}

// GetServerErr returns the value of ServerErr if it is set or its
// zero value if it is unset.
func (v *SimpleService_SillyNoop_Result) GetServerErr() (o *ServerErr) {
	if v != nil && v.ServerErr != nil {
		return v.ServerErr
	}

	return
}

// IsSetServerErr returns true if ServerErr is not nil.
func (v *SimpleService_SillyNoop_Result) IsSetServerErr() bool {
	return v != nil && v.ServerErr != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "sillyNoop" for this struct.
func (v *SimpleService_SillyNoop_Result) MethodName() string {
	return "sillyNoop"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *SimpleService_SillyNoop_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
