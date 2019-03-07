// Code generated by thriftrw v1.16.1. DO NOT EDIT.
// @generated

package googlenow

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// GoogleNowService_CheckCredentials_Args represents the arguments for the GoogleNowService.checkCredentials function.
//
// The arguments for checkCredentials are sent and received over the wire as this struct.
type GoogleNowService_CheckCredentials_Args struct {
}

// ToWire translates a GoogleNowService_CheckCredentials_Args struct into a Thrift-level intermediate
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
func (v *GoogleNowService_CheckCredentials_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a GoogleNowService_CheckCredentials_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a GoogleNowService_CheckCredentials_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v GoogleNowService_CheckCredentials_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *GoogleNowService_CheckCredentials_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a GoogleNowService_CheckCredentials_Args
// struct.
func (v *GoogleNowService_CheckCredentials_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("GoogleNowService_CheckCredentials_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this GoogleNowService_CheckCredentials_Args match the
// provided GoogleNowService_CheckCredentials_Args.
//
// This function performs a deep comparison.
func (v *GoogleNowService_CheckCredentials_Args) Equals(rhs *GoogleNowService_CheckCredentials_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of GoogleNowService_CheckCredentials_Args.
func (v *GoogleNowService_CheckCredentials_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "checkCredentials" for this struct.
func (v *GoogleNowService_CheckCredentials_Args) MethodName() string {
	return "checkCredentials"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *GoogleNowService_CheckCredentials_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// GoogleNowService_CheckCredentials_Helper provides functions that aid in handling the
// parameters and return values of the GoogleNowService.checkCredentials
// function.
var GoogleNowService_CheckCredentials_Helper = struct {
	// Args accepts the parameters of checkCredentials in-order and returns
	// the arguments struct for the function.
	Args func() *GoogleNowService_CheckCredentials_Args

	// IsException returns true if the given error can be thrown
	// by checkCredentials.
	//
	// An error can be thrown by checkCredentials only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for checkCredentials
	// given the error returned by it. The provided error may
	// be nil if checkCredentials did not fail.
	//
	// This allows mapping errors returned by checkCredentials into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// checkCredentials
	//
	//   err := checkCredentials(args)
	//   result, err := GoogleNowService_CheckCredentials_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from checkCredentials: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*GoogleNowService_CheckCredentials_Result, error)

	// UnwrapResponse takes the result struct for checkCredentials
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if checkCredentials threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := GoogleNowService_CheckCredentials_Helper.UnwrapResponse(result)
	UnwrapResponse func(*GoogleNowService_CheckCredentials_Result) error
}{}

func init() {
	GoogleNowService_CheckCredentials_Helper.Args = func() *GoogleNowService_CheckCredentials_Args {
		return &GoogleNowService_CheckCredentials_Args{}
	}

	GoogleNowService_CheckCredentials_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	GoogleNowService_CheckCredentials_Helper.WrapResponse = func(err error) (*GoogleNowService_CheckCredentials_Result, error) {
		if err == nil {
			return &GoogleNowService_CheckCredentials_Result{}, nil
		}

		return nil, err
	}
	GoogleNowService_CheckCredentials_Helper.UnwrapResponse = func(result *GoogleNowService_CheckCredentials_Result) (err error) {
		return
	}

}

// GoogleNowService_CheckCredentials_Result represents the result of a GoogleNowService.checkCredentials function call.
//
// The result of a checkCredentials execution is sent and received over the wire as this struct.
type GoogleNowService_CheckCredentials_Result struct {
}

// ToWire translates a GoogleNowService_CheckCredentials_Result struct into a Thrift-level intermediate
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
func (v *GoogleNowService_CheckCredentials_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a GoogleNowService_CheckCredentials_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a GoogleNowService_CheckCredentials_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v GoogleNowService_CheckCredentials_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *GoogleNowService_CheckCredentials_Result) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a GoogleNowService_CheckCredentials_Result
// struct.
func (v *GoogleNowService_CheckCredentials_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("GoogleNowService_CheckCredentials_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this GoogleNowService_CheckCredentials_Result match the
// provided GoogleNowService_CheckCredentials_Result.
//
// This function performs a deep comparison.
func (v *GoogleNowService_CheckCredentials_Result) Equals(rhs *GoogleNowService_CheckCredentials_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of GoogleNowService_CheckCredentials_Result.
func (v *GoogleNowService_CheckCredentials_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "checkCredentials" for this struct.
func (v *GoogleNowService_CheckCredentials_Result) MethodName() string {
	return "checkCredentials"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *GoogleNowService_CheckCredentials_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
