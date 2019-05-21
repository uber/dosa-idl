// Code generated by thriftrw v1.19.1. DO NOT EDIT.
// @generated

package dosa

import (
	errors "errors"
	fmt "fmt"
	strings "strings"

	multierr "go.uber.org/multierr"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
)

// Dosa_RemoveRange_Args represents the arguments for the Dosa.removeRange function.
//
// The arguments for removeRange are sent and received over the wire as this struct.
type Dosa_RemoveRange_Args struct {
	Request *RemoveRangeRequest `json:"request,omitempty"`
}

// ToWire translates a Dosa_RemoveRange_Args struct into a Thrift-level intermediate
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
func (v *Dosa_RemoveRange_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Request != nil {
		w, err = v.Request.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _RemoveRangeRequest_Read(w wire.Value) (*RemoveRangeRequest, error) {
	var v RemoveRangeRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Dosa_RemoveRange_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_RemoveRange_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_RemoveRange_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_RemoveRange_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _RemoveRangeRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Dosa_RemoveRange_Args
// struct.
func (v *Dosa_RemoveRange_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}

	return fmt.Sprintf("Dosa_RemoveRange_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_RemoveRange_Args match the
// provided Dosa_RemoveRange_Args.
//
// This function performs a deep comparison.
func (v *Dosa_RemoveRange_Args) Equals(rhs *Dosa_RemoveRange_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Dosa_RemoveRange_Args.
func (v *Dosa_RemoveRange_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Request != nil {
		err = multierr.Append(err, enc.AddObject("request", v.Request))
	}
	return err
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *Dosa_RemoveRange_Args) GetRequest() (o *RemoveRangeRequest) {
	if v != nil && v.Request != nil {
		return v.Request
	}

	return
}

// IsSetRequest returns true if Request is not nil.
func (v *Dosa_RemoveRange_Args) IsSetRequest() bool {
	return v != nil && v.Request != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "removeRange" for this struct.
func (v *Dosa_RemoveRange_Args) MethodName() string {
	return "removeRange"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Dosa_RemoveRange_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Dosa_RemoveRange_Helper provides functions that aid in handling the
// parameters and return values of the Dosa.removeRange
// function.
var Dosa_RemoveRange_Helper = struct {
	// Args accepts the parameters of removeRange in-order and returns
	// the arguments struct for the function.
	Args func(
		request *RemoveRangeRequest,
	) *Dosa_RemoveRange_Args

	// IsException returns true if the given error can be thrown
	// by removeRange.
	//
	// An error can be thrown by removeRange only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for removeRange
	// given the error returned by it. The provided error may
	// be nil if removeRange did not fail.
	//
	// This allows mapping errors returned by removeRange into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// removeRange
	//
	//   err := removeRange(args)
	//   result, err := Dosa_RemoveRange_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from removeRange: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*Dosa_RemoveRange_Result, error)

	// UnwrapResponse takes the result struct for removeRange
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if removeRange threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := Dosa_RemoveRange_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Dosa_RemoveRange_Result) error
}{}

func init() {
	Dosa_RemoveRange_Helper.Args = func(
		request *RemoveRangeRequest,
	) *Dosa_RemoveRange_Args {
		return &Dosa_RemoveRange_Args{
			Request: request,
		}
	}

	Dosa_RemoveRange_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BadRequestError:
			return true
		case *InternalServerError:
			return true
		default:
			return false
		}
	}

	Dosa_RemoveRange_Helper.WrapResponse = func(err error) (*Dosa_RemoveRange_Result, error) {
		if err == nil {
			return &Dosa_RemoveRange_Result{}, nil
		}

		switch e := err.(type) {
		case *BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_RemoveRange_Result.ClientError")
			}
			return &Dosa_RemoveRange_Result{ClientError: e}, nil
		case *InternalServerError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_RemoveRange_Result.ServerError")
			}
			return &Dosa_RemoveRange_Result{ServerError: e}, nil
		}

		return nil, err
	}
	Dosa_RemoveRange_Helper.UnwrapResponse = func(result *Dosa_RemoveRange_Result) (err error) {
		if result.ClientError != nil {
			err = result.ClientError
			return
		}
		if result.ServerError != nil {
			err = result.ServerError
			return
		}
		return
	}

}

// Dosa_RemoveRange_Result represents the result of a Dosa.removeRange function call.
//
// The result of a removeRange execution is sent and received over the wire as this struct.
type Dosa_RemoveRange_Result struct {
	ClientError *BadRequestError     `json:"clientError,omitempty"`
	ServerError *InternalServerError `json:"serverError,omitempty"`
}

// ToWire translates a Dosa_RemoveRange_Result struct into a Thrift-level intermediate
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
func (v *Dosa_RemoveRange_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.ClientError != nil {
		w, err = v.ClientError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.ServerError != nil {
		w, err = v.ServerError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	if i > 1 {
		return wire.Value{}, fmt.Errorf("Dosa_RemoveRange_Result should have at most one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Dosa_RemoveRange_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_RemoveRange_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_RemoveRange_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_RemoveRange_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.ClientError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.ServerError, err = _InternalServerError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.ClientError != nil {
		count++
	}
	if v.ServerError != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("Dosa_RemoveRange_Result should have at most one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Dosa_RemoveRange_Result
// struct.
func (v *Dosa_RemoveRange_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.ClientError != nil {
		fields[i] = fmt.Sprintf("ClientError: %v", v.ClientError)
		i++
	}
	if v.ServerError != nil {
		fields[i] = fmt.Sprintf("ServerError: %v", v.ServerError)
		i++
	}

	return fmt.Sprintf("Dosa_RemoveRange_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_RemoveRange_Result match the
// provided Dosa_RemoveRange_Result.
//
// This function performs a deep comparison.
func (v *Dosa_RemoveRange_Result) Equals(rhs *Dosa_RemoveRange_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.ClientError == nil && rhs.ClientError == nil) || (v.ClientError != nil && rhs.ClientError != nil && v.ClientError.Equals(rhs.ClientError))) {
		return false
	}
	if !((v.ServerError == nil && rhs.ServerError == nil) || (v.ServerError != nil && rhs.ServerError != nil && v.ServerError.Equals(rhs.ServerError))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Dosa_RemoveRange_Result.
func (v *Dosa_RemoveRange_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.ClientError != nil {
		err = multierr.Append(err, enc.AddObject("clientError", v.ClientError))
	}
	if v.ServerError != nil {
		err = multierr.Append(err, enc.AddObject("serverError", v.ServerError))
	}
	return err
}

// GetClientError returns the value of ClientError if it is set or its
// zero value if it is unset.
func (v *Dosa_RemoveRange_Result) GetClientError() (o *BadRequestError) {
	if v != nil && v.ClientError != nil {
		return v.ClientError
	}

	return
}

// IsSetClientError returns true if ClientError is not nil.
func (v *Dosa_RemoveRange_Result) IsSetClientError() bool {
	return v != nil && v.ClientError != nil
}

// GetServerError returns the value of ServerError if it is set or its
// zero value if it is unset.
func (v *Dosa_RemoveRange_Result) GetServerError() (o *InternalServerError) {
	if v != nil && v.ServerError != nil {
		return v.ServerError
	}

	return
}

// IsSetServerError returns true if ServerError is not nil.
func (v *Dosa_RemoveRange_Result) IsSetServerError() bool {
	return v != nil && v.ServerError != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "removeRange" for this struct.
func (v *Dosa_RemoveRange_Result) MethodName() string {
	return "removeRange"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Dosa_RemoveRange_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
