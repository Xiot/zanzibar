// Code generated by thriftrw v1.19.0. DO NOT EDIT.
// @generated

package multi

import (
	errors "errors"
	fmt "fmt"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// ServiceCBack_Hello_Args represents the arguments for the ServiceCBack.hello function.
//
// The arguments for hello are sent and received over the wire as this struct.
type ServiceCBack_Hello_Args struct {
}

// ToWire translates a ServiceCBack_Hello_Args struct into a Thrift-level intermediate
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
func (v *ServiceCBack_Hello_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ServiceCBack_Hello_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ServiceCBack_Hello_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ServiceCBack_Hello_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ServiceCBack_Hello_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a ServiceCBack_Hello_Args
// struct.
func (v *ServiceCBack_Hello_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("ServiceCBack_Hello_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ServiceCBack_Hello_Args match the
// provided ServiceCBack_Hello_Args.
//
// This function performs a deep comparison.
func (v *ServiceCBack_Hello_Args) Equals(rhs *ServiceCBack_Hello_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ServiceCBack_Hello_Args.
func (v *ServiceCBack_Hello_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "hello" for this struct.
func (v *ServiceCBack_Hello_Args) MethodName() string {
	return "hello"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *ServiceCBack_Hello_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// ServiceCBack_Hello_Helper provides functions that aid in handling the
// parameters and return values of the ServiceCBack.hello
// function.
var ServiceCBack_Hello_Helper = struct {
	// Args accepts the parameters of hello in-order and returns
	// the arguments struct for the function.
	Args func() *ServiceCBack_Hello_Args

	// IsException returns true if the given error can be thrown
	// by hello.
	//
	// An error can be thrown by hello only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for hello
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// hello into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by hello
	//
	//   value, err := hello(args)
	//   result, err := ServiceCBack_Hello_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from hello: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(string, error) (*ServiceCBack_Hello_Result, error)

	// UnwrapResponse takes the result struct for hello
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if hello threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := ServiceCBack_Hello_Helper.UnwrapResponse(result)
	UnwrapResponse func(*ServiceCBack_Hello_Result) (string, error)
}{}

func init() {
	ServiceCBack_Hello_Helper.Args = func() *ServiceCBack_Hello_Args {
		return &ServiceCBack_Hello_Args{}
	}

	ServiceCBack_Hello_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	ServiceCBack_Hello_Helper.WrapResponse = func(success string, err error) (*ServiceCBack_Hello_Result, error) {
		if err == nil {
			return &ServiceCBack_Hello_Result{Success: &success}, nil
		}

		return nil, err
	}
	ServiceCBack_Hello_Helper.UnwrapResponse = func(result *ServiceCBack_Hello_Result) (success string, err error) {

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// ServiceCBack_Hello_Result represents the result of a ServiceCBack.hello function call.
//
// The result of a hello execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type ServiceCBack_Hello_Result struct {
	// Value returned by hello after a successful execution.
	Success *string `json:"success,omitempty"`
}

// ToWire translates a ServiceCBack_Hello_Result struct into a Thrift-level intermediate
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
func (v *ServiceCBack_Hello_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueString(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("ServiceCBack_Hello_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ServiceCBack_Hello_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ServiceCBack_Hello_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ServiceCBack_Hello_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ServiceCBack_Hello_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Success = &x
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
	if count != 1 {
		return fmt.Errorf("ServiceCBack_Hello_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a ServiceCBack_Hello_Result
// struct.
func (v *ServiceCBack_Hello_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}

	return fmt.Sprintf("ServiceCBack_Hello_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ServiceCBack_Hello_Result match the
// provided ServiceCBack_Hello_Result.
//
// This function performs a deep comparison.
func (v *ServiceCBack_Hello_Result) Equals(rhs *ServiceCBack_Hello_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Success, rhs.Success) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ServiceCBack_Hello_Result.
func (v *ServiceCBack_Hello_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		enc.AddString("success", *v.Success)
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *ServiceCBack_Hello_Result) GetSuccess() (o string) {
	if v != nil && v.Success != nil {
		return *v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *ServiceCBack_Hello_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "hello" for this struct.
func (v *ServiceCBack_Hello_Result) MethodName() string {
	return "hello"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *ServiceCBack_Hello_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}