// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

package baz

import (
	"errors"
	"fmt"
	"github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// SimpleService_TransHeaders_Args represents the arguments for the SimpleService.transHeaders function.
//
// The arguments for transHeaders are sent and received over the wire as this struct.
type SimpleService_TransHeaders_Args struct {
	Req *base.TransHeaders `json:"req,required"`
}

// ToWire translates a SimpleService_TransHeaders_Args struct into a Thrift-level intermediate
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
func (v *SimpleService_TransHeaders_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Req == nil {
		return w, errors.New("field Req of SimpleService_TransHeaders_Args is required")
	}
	w, err = v.Req.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _TransHeaders_Read(w wire.Value) (*base.TransHeaders, error) {
	var v base.TransHeaders
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a SimpleService_TransHeaders_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SimpleService_TransHeaders_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SimpleService_TransHeaders_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SimpleService_TransHeaders_Args) FromWire(w wire.Value) error {
	var err error

	reqIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Req, err = _TransHeaders_Read(field.Value)
				if err != nil {
					return err
				}
				reqIsSet = true
			}
		}
	}

	if !reqIsSet {
		return errors.New("field Req of SimpleService_TransHeaders_Args is required")
	}

	return nil
}

// String returns a readable string representation of a SimpleService_TransHeaders_Args
// struct.
func (v *SimpleService_TransHeaders_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Req: %v", v.Req)
	i++

	return fmt.Sprintf("SimpleService_TransHeaders_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SimpleService_TransHeaders_Args match the
// provided SimpleService_TransHeaders_Args.
//
// This function performs a deep comparison.
func (v *SimpleService_TransHeaders_Args) Equals(rhs *SimpleService_TransHeaders_Args) bool {
	if !v.Req.Equals(rhs.Req) {
		return false
	}

	return true
}

// GetReq returns the value of Req if it is set or its
// zero value if it is unset.
func (v *SimpleService_TransHeaders_Args) GetReq() (o *base.TransHeaders) { return v.Req }

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "transHeaders" for this struct.
func (v *SimpleService_TransHeaders_Args) MethodName() string {
	return "transHeaders"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *SimpleService_TransHeaders_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// SimpleService_TransHeaders_Helper provides functions that aid in handling the
// parameters and return values of the SimpleService.transHeaders
// function.
var SimpleService_TransHeaders_Helper = struct {
	// Args accepts the parameters of transHeaders in-order and returns
	// the arguments struct for the function.
	Args func(
		req *base.TransHeaders,
	) *SimpleService_TransHeaders_Args

	// IsException returns true if the given error can be thrown
	// by transHeaders.
	//
	// An error can be thrown by transHeaders only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for transHeaders
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// transHeaders into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by transHeaders
	//
	//   value, err := transHeaders(args)
	//   result, err := SimpleService_TransHeaders_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from transHeaders: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*base.TransHeaders, error) (*SimpleService_TransHeaders_Result, error)

	// UnwrapResponse takes the result struct for transHeaders
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if transHeaders threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := SimpleService_TransHeaders_Helper.UnwrapResponse(result)
	UnwrapResponse func(*SimpleService_TransHeaders_Result) (*base.TransHeaders, error)
}{}

func init() {
	SimpleService_TransHeaders_Helper.Args = func(
		req *base.TransHeaders,
	) *SimpleService_TransHeaders_Args {
		return &SimpleService_TransHeaders_Args{
			Req: req,
		}
	}

	SimpleService_TransHeaders_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *AuthErr:
			return true
		case *OtherAuthErr:
			return true
		default:
			return false
		}
	}

	SimpleService_TransHeaders_Helper.WrapResponse = func(success *base.TransHeaders, err error) (*SimpleService_TransHeaders_Result, error) {
		if err == nil {
			return &SimpleService_TransHeaders_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *AuthErr:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for SimpleService_TransHeaders_Result.AuthErr")
			}
			return &SimpleService_TransHeaders_Result{AuthErr: e}, nil
		case *OtherAuthErr:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for SimpleService_TransHeaders_Result.OtherAuthErr")
			}
			return &SimpleService_TransHeaders_Result{OtherAuthErr: e}, nil
		}

		return nil, err
	}
	SimpleService_TransHeaders_Helper.UnwrapResponse = func(result *SimpleService_TransHeaders_Result) (success *base.TransHeaders, err error) {
		if result.AuthErr != nil {
			err = result.AuthErr
			return
		}
		if result.OtherAuthErr != nil {
			err = result.OtherAuthErr
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// SimpleService_TransHeaders_Result represents the result of a SimpleService.transHeaders function call.
//
// The result of a transHeaders execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type SimpleService_TransHeaders_Result struct {
	// Value returned by transHeaders after a successful execution.
	Success      *base.TransHeaders `json:"success,omitempty"`
	AuthErr      *AuthErr           `json:"authErr,omitempty"`
	OtherAuthErr *OtherAuthErr      `json:"otherAuthErr,omitempty"`
}

// ToWire translates a SimpleService_TransHeaders_Result struct into a Thrift-level intermediate
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
func (v *SimpleService_TransHeaders_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.AuthErr != nil {
		w, err = v.AuthErr.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.OtherAuthErr != nil {
		w, err = v.OtherAuthErr.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("SimpleService_TransHeaders_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a SimpleService_TransHeaders_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SimpleService_TransHeaders_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SimpleService_TransHeaders_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SimpleService_TransHeaders_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _TransHeaders_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.AuthErr, err = _AuthErr_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.OtherAuthErr, err = _OtherAuthErr_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.AuthErr != nil {
		count++
	}
	if v.OtherAuthErr != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("SimpleService_TransHeaders_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a SimpleService_TransHeaders_Result
// struct.
func (v *SimpleService_TransHeaders_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.AuthErr != nil {
		fields[i] = fmt.Sprintf("AuthErr: %v", v.AuthErr)
		i++
	}
	if v.OtherAuthErr != nil {
		fields[i] = fmt.Sprintf("OtherAuthErr: %v", v.OtherAuthErr)
		i++
	}

	return fmt.Sprintf("SimpleService_TransHeaders_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SimpleService_TransHeaders_Result match the
// provided SimpleService_TransHeaders_Result.
//
// This function performs a deep comparison.
func (v *SimpleService_TransHeaders_Result) Equals(rhs *SimpleService_TransHeaders_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.AuthErr == nil && rhs.AuthErr == nil) || (v.AuthErr != nil && rhs.AuthErr != nil && v.AuthErr.Equals(rhs.AuthErr))) {
		return false
	}
	if !((v.OtherAuthErr == nil && rhs.OtherAuthErr == nil) || (v.OtherAuthErr != nil && rhs.OtherAuthErr != nil && v.OtherAuthErr.Equals(rhs.OtherAuthErr))) {
		return false
	}

	return true
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *SimpleService_TransHeaders_Result) GetSuccess() (o *base.TransHeaders) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// GetAuthErr returns the value of AuthErr if it is set or its
// zero value if it is unset.
func (v *SimpleService_TransHeaders_Result) GetAuthErr() (o *AuthErr) {
	if v.AuthErr != nil {
		return v.AuthErr
	}

	return
}

// GetOtherAuthErr returns the value of OtherAuthErr if it is set or its
// zero value if it is unset.
func (v *SimpleService_TransHeaders_Result) GetOtherAuthErr() (o *OtherAuthErr) {
	if v.OtherAuthErr != nil {
		return v.OtherAuthErr
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "transHeaders" for this struct.
func (v *SimpleService_TransHeaders_Result) MethodName() string {
	return "transHeaders"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *SimpleService_TransHeaders_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
