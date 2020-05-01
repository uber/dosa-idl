// Code generated by thriftrw v1.13.0. DO NOT EDIT.
// @generated

package dosa

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// Dosa_CreateIfNotExists_Args represents the arguments for the Dosa.createIfNotExists function.
//
// The arguments for createIfNotExists are sent and received over the wire as this struct.
type Dosa_CreateIfNotExists_Args struct {
	Request *CreateRequest `json:"request,omitempty"`
}

// ToWire translates a Dosa_CreateIfNotExists_Args struct into a Thrift-level intermediate
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
func (v *Dosa_CreateIfNotExists_Args) ToWire() (wire.Value, error) {
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

func _CreateRequest_Read(w wire.Value) (*CreateRequest, error) {
	var v CreateRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Dosa_CreateIfNotExists_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_CreateIfNotExists_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_CreateIfNotExists_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_CreateIfNotExists_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _CreateRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Dosa_CreateIfNotExists_Args
// struct.
func (v *Dosa_CreateIfNotExists_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}

	return fmt.Sprintf("Dosa_CreateIfNotExists_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_CreateIfNotExists_Args match the
// provided Dosa_CreateIfNotExists_Args.
//
// This function performs a deep comparison.
func (v *Dosa_CreateIfNotExists_Args) Equals(rhs *Dosa_CreateIfNotExists_Args) bool {
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
// fast logging of Dosa_CreateIfNotExists_Args.
func (v *Dosa_CreateIfNotExists_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v.Request != nil {
		err = multierr.Append(err, enc.AddObject("request", v.Request))
	}
	return err
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *Dosa_CreateIfNotExists_Args) GetRequest() (o *CreateRequest) {
	if v.Request != nil {
		return v.Request
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "createIfNotExists" for this struct.
func (v *Dosa_CreateIfNotExists_Args) MethodName() string {
	return "createIfNotExists"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Dosa_CreateIfNotExists_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Dosa_CreateIfNotExists_Helper provides functions that aid in handling the
// parameters and return values of the Dosa.createIfNotExists
// function.
var Dosa_CreateIfNotExists_Helper = struct {
	// Args accepts the parameters of createIfNotExists in-order and returns
	// the arguments struct for the function.
	Args func(
		request *CreateRequest,
	) *Dosa_CreateIfNotExists_Args

	// IsException returns true if the given error can be thrown
	// by createIfNotExists.
	//
	// An error can be thrown by createIfNotExists only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for createIfNotExists
	// given the error returned by it. The provided error may
	// be nil if createIfNotExists did not fail.
	//
	// This allows mapping errors returned by createIfNotExists into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// createIfNotExists
	//
	//   err := createIfNotExists(args)
	//   result, err := Dosa_CreateIfNotExists_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from createIfNotExists: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*Dosa_CreateIfNotExists_Result, error)

	// UnwrapResponse takes the result struct for createIfNotExists
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if createIfNotExists threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := Dosa_CreateIfNotExists_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Dosa_CreateIfNotExists_Result) error
}{}

func init() {
	Dosa_CreateIfNotExists_Helper.Args = func(
		request *CreateRequest,
	) *Dosa_CreateIfNotExists_Args {
		return &Dosa_CreateIfNotExists_Args{
			Request: request,
		}
	}

	Dosa_CreateIfNotExists_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BadRequestError:
			return true
		case *InternalServerError:
			return true
		case *RateLimitError:
			return true
		default:
			return false
		}
	}

	Dosa_CreateIfNotExists_Helper.WrapResponse = func(err error) (*Dosa_CreateIfNotExists_Result, error) {
		if err == nil {
			return &Dosa_CreateIfNotExists_Result{}, nil
		}

		switch e := err.(type) {
		case *BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_CreateIfNotExists_Result.ClientError")
			}
			return &Dosa_CreateIfNotExists_Result{ClientError: e}, nil
		case *InternalServerError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_CreateIfNotExists_Result.ServerError")
			}
			return &Dosa_CreateIfNotExists_Result{ServerError: e}, nil
		case *RateLimitError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_CreateIfNotExists_Result.LimitError")
			}
			return &Dosa_CreateIfNotExists_Result{LimitError: e}, nil
		}

		return nil, err
	}
	Dosa_CreateIfNotExists_Helper.UnwrapResponse = func(result *Dosa_CreateIfNotExists_Result) (err error) {
		if result.ClientError != nil {
			err = result.ClientError
			return
		}
		if result.ServerError != nil {
			err = result.ServerError
			return
		}
		if result.LimitError != nil {
			err = result.LimitError
			return
		}
		return
	}

}

// Dosa_CreateIfNotExists_Result represents the result of a Dosa.createIfNotExists function call.
//
// The result of a createIfNotExists execution is sent and received over the wire as this struct.
type Dosa_CreateIfNotExists_Result struct {
	ClientError *BadRequestError     `json:"clientError,omitempty"`
	ServerError *InternalServerError `json:"serverError,omitempty"`
	LimitError  *RateLimitError      `json:"limitError,omitempty"`
}

// ToWire translates a Dosa_CreateIfNotExists_Result struct into a Thrift-level intermediate
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
func (v *Dosa_CreateIfNotExists_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
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
	if v.LimitError != nil {
		w, err = v.LimitError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}

	if i > 1 {
		return wire.Value{}, fmt.Errorf("Dosa_CreateIfNotExists_Result should have at most one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _RateLimitError_Read(w wire.Value) (*RateLimitError, error) {
	var v RateLimitError
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Dosa_CreateIfNotExists_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_CreateIfNotExists_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_CreateIfNotExists_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_CreateIfNotExists_Result) FromWire(w wire.Value) error {
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
		case 3:
			if field.Value.Type() == wire.TStruct {
				v.LimitError, err = _RateLimitError_Read(field.Value)
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
	if v.LimitError != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("Dosa_CreateIfNotExists_Result should have at most one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Dosa_CreateIfNotExists_Result
// struct.
func (v *Dosa_CreateIfNotExists_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	if v.ClientError != nil {
		fields[i] = fmt.Sprintf("ClientError: %v", v.ClientError)
		i++
	}
	if v.ServerError != nil {
		fields[i] = fmt.Sprintf("ServerError: %v", v.ServerError)
		i++
	}
	if v.LimitError != nil {
		fields[i] = fmt.Sprintf("LimitError: %v", v.LimitError)
		i++
	}

	return fmt.Sprintf("Dosa_CreateIfNotExists_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_CreateIfNotExists_Result match the
// provided Dosa_CreateIfNotExists_Result.
//
// This function performs a deep comparison.
func (v *Dosa_CreateIfNotExists_Result) Equals(rhs *Dosa_CreateIfNotExists_Result) bool {
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
	if !((v.LimitError == nil && rhs.LimitError == nil) || (v.LimitError != nil && rhs.LimitError != nil && v.LimitError.Equals(rhs.LimitError))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Dosa_CreateIfNotExists_Result.
func (v *Dosa_CreateIfNotExists_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v.ClientError != nil {
		err = multierr.Append(err, enc.AddObject("clientError", v.ClientError))
	}
	if v.ServerError != nil {
		err = multierr.Append(err, enc.AddObject("serverError", v.ServerError))
	}
	if v.LimitError != nil {
		err = multierr.Append(err, enc.AddObject("limitError", v.LimitError))
	}
	return err
}

// GetClientError returns the value of ClientError if it is set or its
// zero value if it is unset.
func (v *Dosa_CreateIfNotExists_Result) GetClientError() (o *BadRequestError) {
	if v.ClientError != nil {
		return v.ClientError
	}

	return
}

// GetServerError returns the value of ServerError if it is set or its
// zero value if it is unset.
func (v *Dosa_CreateIfNotExists_Result) GetServerError() (o *InternalServerError) {
	if v.ServerError != nil {
		return v.ServerError
	}

	return
}

// GetLimitError returns the value of LimitError if it is set or its
// zero value if it is unset.
func (v *Dosa_CreateIfNotExists_Result) GetLimitError() (o *RateLimitError) {
	if v.LimitError != nil {
		return v.LimitError
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "createIfNotExists" for this struct.
func (v *Dosa_CreateIfNotExists_Result) MethodName() string {
	return "createIfNotExists"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Dosa_CreateIfNotExists_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
