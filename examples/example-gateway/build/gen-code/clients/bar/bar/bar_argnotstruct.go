// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// Bar_ArgNotStruct_Args represents the arguments for the Bar.argNotStruct function.
//
// The arguments for argNotStruct are sent and received over the wire as this struct.
type Bar_ArgNotStruct_Args struct {
	Request string `json:"request,required"`
}

// ToWire translates a Bar_ArgNotStruct_Args struct into a Thrift-level intermediate
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
func (v *Bar_ArgNotStruct_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Request), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Bar_ArgNotStruct_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_ArgNotStruct_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_ArgNotStruct_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_ArgNotStruct_Args) FromWire(w wire.Value) error {
	var err error

	requestIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Request, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				requestIsSet = true
			}
		}
	}

	if !requestIsSet {
		return errors.New("field Request of Bar_ArgNotStruct_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Bar_ArgNotStruct_Args
// struct.
func (v *Bar_ArgNotStruct_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Request: %v", v.Request)
	i++

	return fmt.Sprintf("Bar_ArgNotStruct_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_ArgNotStruct_Args match the
// provided Bar_ArgNotStruct_Args.
//
// This function performs a deep comparison.
func (v *Bar_ArgNotStruct_Args) Equals(rhs *Bar_ArgNotStruct_Args) bool {
	if !(v.Request == rhs.Request) {
		return false
	}

	return true
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *Bar_ArgNotStruct_Args) GetRequest() (o string) { return v.Request }

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "argNotStruct" for this struct.
func (v *Bar_ArgNotStruct_Args) MethodName() string {
	return "argNotStruct"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Bar_ArgNotStruct_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Bar_ArgNotStruct_Helper provides functions that aid in handling the
// parameters and return values of the Bar.argNotStruct
// function.
var Bar_ArgNotStruct_Helper = struct {
	// Args accepts the parameters of argNotStruct in-order and returns
	// the arguments struct for the function.
	Args func(
		request string,
	) *Bar_ArgNotStruct_Args

	// IsException returns true if the given error can be thrown
	// by argNotStruct.
	//
	// An error can be thrown by argNotStruct only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for argNotStruct
	// given the error returned by it. The provided error may
	// be nil if argNotStruct did not fail.
	//
	// This allows mapping errors returned by argNotStruct into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// argNotStruct
	//
	//   err := argNotStruct(args)
	//   result, err := Bar_ArgNotStruct_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from argNotStruct: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*Bar_ArgNotStruct_Result, error)

	// UnwrapResponse takes the result struct for argNotStruct
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if argNotStruct threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := Bar_ArgNotStruct_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Bar_ArgNotStruct_Result) error
}{}

func init() {
	Bar_ArgNotStruct_Helper.Args = func(
		request string,
	) *Bar_ArgNotStruct_Args {
		return &Bar_ArgNotStruct_Args{
			Request: request,
		}
	}

	Bar_ArgNotStruct_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BarException:
			return true
		default:
			return false
		}
	}

	Bar_ArgNotStruct_Helper.WrapResponse = func(err error) (*Bar_ArgNotStruct_Result, error) {
		if err == nil {
			return &Bar_ArgNotStruct_Result{}, nil
		}

		switch e := err.(type) {
		case *BarException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Bar_ArgNotStruct_Result.BarException")
			}
			return &Bar_ArgNotStruct_Result{BarException: e}, nil
		}

		return nil, err
	}
	Bar_ArgNotStruct_Helper.UnwrapResponse = func(result *Bar_ArgNotStruct_Result) (err error) {
		if result.BarException != nil {
			err = result.BarException
			return
		}
		return
	}

}

// Bar_ArgNotStruct_Result represents the result of a Bar.argNotStruct function call.
//
// The result of a argNotStruct execution is sent and received over the wire as this struct.
type Bar_ArgNotStruct_Result struct {
	BarException *BarException `json:"barException,omitempty"`
}

// ToWire translates a Bar_ArgNotStruct_Result struct into a Thrift-level intermediate
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
func (v *Bar_ArgNotStruct_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.BarException != nil {
		w, err = v.BarException.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	if i > 1 {
		return wire.Value{}, fmt.Errorf("Bar_ArgNotStruct_Result should have at most one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _BarException_Read(w wire.Value) (*BarException, error) {
	var v BarException
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Bar_ArgNotStruct_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_ArgNotStruct_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_ArgNotStruct_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_ArgNotStruct_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BarException, err = _BarException_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.BarException != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("Bar_ArgNotStruct_Result should have at most one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Bar_ArgNotStruct_Result
// struct.
func (v *Bar_ArgNotStruct_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.BarException != nil {
		fields[i] = fmt.Sprintf("BarException: %v", v.BarException)
		i++
	}

	return fmt.Sprintf("Bar_ArgNotStruct_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_ArgNotStruct_Result match the
// provided Bar_ArgNotStruct_Result.
//
// This function performs a deep comparison.
func (v *Bar_ArgNotStruct_Result) Equals(rhs *Bar_ArgNotStruct_Result) bool {
	if !((v.BarException == nil && rhs.BarException == nil) || (v.BarException != nil && rhs.BarException != nil && v.BarException.Equals(rhs.BarException))) {
		return false
	}

	return true
}

// GetBarException returns the value of BarException if it is set or its
// zero value if it is unset.
func (v *Bar_ArgNotStruct_Result) GetBarException() (o *BarException) {
	if v.BarException != nil {
		return v.BarException
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "argNotStruct" for this struct.
func (v *Bar_ArgNotStruct_Result) MethodName() string {
	return "argNotStruct"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Bar_ArgNotStruct_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
