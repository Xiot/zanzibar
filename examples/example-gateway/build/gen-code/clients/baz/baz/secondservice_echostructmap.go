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

// SecondService_EchoStructMap_Args represents the arguments for the SecondService.echoStructMap function.
//
// The arguments for echoStructMap are sent and received over the wire as this struct.
type SecondService_EchoStructMap_Args struct {
	Arg []struct {
		Key   *base.BazResponse
		Value string
	} `json:"arg,required"`
}

type _Map_BazResponse_String_MapItemList []struct {
	Key   *base.BazResponse
	Value string
}

func (m _Map_BazResponse_String_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for _, i := range m {
		k := i.Key
		v := i.Value
		if k == nil {
			return fmt.Errorf("invalid map key: value is nil")
		}
		kw, err := k.ToWire()
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

func (m _Map_BazResponse_String_MapItemList) Size() int {
	return len(m)
}

func (_Map_BazResponse_String_MapItemList) KeyType() wire.Type {
	return wire.TStruct
}

func (_Map_BazResponse_String_MapItemList) ValueType() wire.Type {
	return wire.TBinary
}

func (_Map_BazResponse_String_MapItemList) Close() {}

// ToWire translates a SecondService_EchoStructMap_Args struct into a Thrift-level intermediate
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
func (v *SecondService_EchoStructMap_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Arg == nil {
		return w, errors.New("field Arg of SecondService_EchoStructMap_Args is required")
	}
	w, err = wire.NewValueMap(_Map_BazResponse_String_MapItemList(v.Arg)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Map_BazResponse_String_Read(m wire.MapItemList) ([]struct {
	Key   *base.BazResponse
	Value string
}, error) {
	if m.KeyType() != wire.TStruct {
		return nil, nil
	}

	if m.ValueType() != wire.TBinary {
		return nil, nil
	}

	o := make([]struct {
		Key   *base.BazResponse
		Value string
	}, 0, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := _BazResponse_Read(x.Key)
		if err != nil {
			return err
		}

		v, err := x.Value.GetString(), error(nil)
		if err != nil {
			return err
		}

		o = append(o, struct {
			Key   *base.BazResponse
			Value string
		}{k, v})
		return nil
	})
	m.Close()
	return o, err
}

// FromWire deserializes a SecondService_EchoStructMap_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SecondService_EchoStructMap_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SecondService_EchoStructMap_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SecondService_EchoStructMap_Args) FromWire(w wire.Value) error {
	var err error

	argIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TMap {
				v.Arg, err = _Map_BazResponse_String_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
				argIsSet = true
			}
		}
	}

	if !argIsSet {
		return errors.New("field Arg of SecondService_EchoStructMap_Args is required")
	}

	return nil
}

// String returns a readable string representation of a SecondService_EchoStructMap_Args
// struct.
func (v *SecondService_EchoStructMap_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Arg: %v", v.Arg)
	i++

	return fmt.Sprintf("SecondService_EchoStructMap_Args{%v}", strings.Join(fields[:i], ", "))
}

func _Map_BazResponse_String_Equals(lhs, rhs []struct {
	Key   *base.BazResponse
	Value string
}) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for _, i := range lhs {
		lk := i.Key
		lv := i.Value
		ok := false
		for _, j := range rhs {
			rk := j.Key
			rv := j.Value
			if !lk.Equals(rk) {
				continue
			}

			if !(lv == rv) {
				return false
			}
			ok = true
			break
		}

		if !ok {
			return false
		}
	}
	return true
}

// Equals returns true if all the fields of this SecondService_EchoStructMap_Args match the
// provided SecondService_EchoStructMap_Args.
//
// This function performs a deep comparison.
func (v *SecondService_EchoStructMap_Args) Equals(rhs *SecondService_EchoStructMap_Args) bool {
	if !_Map_BazResponse_String_Equals(v.Arg, rhs.Arg) {
		return false
	}

	return true
}

// GetArg returns the value of Arg if it is set or its
// zero value if it is unset.
func (v *SecondService_EchoStructMap_Args) GetArg() (o []struct {
	Key   *base.BazResponse
	Value string
}) {
	return v.Arg
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "echoStructMap" for this struct.
func (v *SecondService_EchoStructMap_Args) MethodName() string {
	return "echoStructMap"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *SecondService_EchoStructMap_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// SecondService_EchoStructMap_Helper provides functions that aid in handling the
// parameters and return values of the SecondService.echoStructMap
// function.
var SecondService_EchoStructMap_Helper = struct {
	// Args accepts the parameters of echoStructMap in-order and returns
	// the arguments struct for the function.
	Args func(
		arg []struct {
			Key   *base.BazResponse
			Value string
		},
	) *SecondService_EchoStructMap_Args

	// IsException returns true if the given error can be thrown
	// by echoStructMap.
	//
	// An error can be thrown by echoStructMap only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for echoStructMap
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// echoStructMap into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by echoStructMap
	//
	//   value, err := echoStructMap(args)
	//   result, err := SecondService_EchoStructMap_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from echoStructMap: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func([]struct {
		Key   *base.BazResponse
		Value string
	}, error) (*SecondService_EchoStructMap_Result, error)

	// UnwrapResponse takes the result struct for echoStructMap
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if echoStructMap threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := SecondService_EchoStructMap_Helper.UnwrapResponse(result)
	UnwrapResponse func(*SecondService_EchoStructMap_Result) ([]struct {
		Key   *base.BazResponse
		Value string
	}, error)
}{}

func init() {
	SecondService_EchoStructMap_Helper.Args = func(
		arg []struct {
			Key   *base.BazResponse
			Value string
		},
	) *SecondService_EchoStructMap_Args {
		return &SecondService_EchoStructMap_Args{
			Arg: arg,
		}
	}

	SecondService_EchoStructMap_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	SecondService_EchoStructMap_Helper.WrapResponse = func(success []struct {
		Key   *base.BazResponse
		Value string
	}, err error) (*SecondService_EchoStructMap_Result, error) {
		if err == nil {
			return &SecondService_EchoStructMap_Result{Success: success}, nil
		}

		return nil, err
	}
	SecondService_EchoStructMap_Helper.UnwrapResponse = func(result *SecondService_EchoStructMap_Result) (success []struct {
		Key   *base.BazResponse
		Value string
	}, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// SecondService_EchoStructMap_Result represents the result of a SecondService.echoStructMap function call.
//
// The result of a echoStructMap execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type SecondService_EchoStructMap_Result struct {
	// Value returned by echoStructMap after a successful execution.
	Success []struct {
		Key   *base.BazResponse
		Value string
	} `json:"success,omitempty"`
}

// ToWire translates a SecondService_EchoStructMap_Result struct into a Thrift-level intermediate
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
func (v *SecondService_EchoStructMap_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueMap(_Map_BazResponse_String_MapItemList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("SecondService_EchoStructMap_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a SecondService_EchoStructMap_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SecondService_EchoStructMap_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SecondService_EchoStructMap_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SecondService_EchoStructMap_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TMap {
				v.Success, err = _Map_BazResponse_String_Read(field.Value.GetMap())
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
		return fmt.Errorf("SecondService_EchoStructMap_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a SecondService_EchoStructMap_Result
// struct.
func (v *SecondService_EchoStructMap_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("SecondService_EchoStructMap_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SecondService_EchoStructMap_Result match the
// provided SecondService_EchoStructMap_Result.
//
// This function performs a deep comparison.
func (v *SecondService_EchoStructMap_Result) Equals(rhs *SecondService_EchoStructMap_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && _Map_BazResponse_String_Equals(v.Success, rhs.Success))) {
		return false
	}

	return true
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *SecondService_EchoStructMap_Result) GetSuccess() (o []struct {
	Key   *base.BazResponse
	Value string
}) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "echoStructMap" for this struct.
func (v *SecondService_EchoStructMap_Result) MethodName() string {
	return "echoStructMap"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *SecondService_EchoStructMap_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
