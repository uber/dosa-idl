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

// Dosa_CheckSchema_Args represents the arguments for the Dosa.checkSchema function.
//
// The arguments for checkSchema are sent and received over the wire as this struct.
type Dosa_CheckSchema_Args struct {
	Request *CheckSchemaRequest `json:"request,omitempty"`
}

// ToWire translates a Dosa_CheckSchema_Args struct into a Thrift-level intermediate
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
func (v *Dosa_CheckSchema_Args) ToWire() (wire.Value, error) {
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

func _CheckSchemaRequest_Read(w wire.Value) (*CheckSchemaRequest, error) {
	var v CheckSchemaRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Dosa_CheckSchema_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_CheckSchema_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_CheckSchema_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_CheckSchema_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _CheckSchemaRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Dosa_CheckSchema_Args
// struct.
func (v *Dosa_CheckSchema_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}

	return fmt.Sprintf("Dosa_CheckSchema_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_CheckSchema_Args match the
// provided Dosa_CheckSchema_Args.
//
// This function performs a deep comparison.
func (v *Dosa_CheckSchema_Args) Equals(rhs *Dosa_CheckSchema_Args) bool {
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
// fast logging of Dosa_CheckSchema_Args.
func (v *Dosa_CheckSchema_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v.Request != nil {
		err = multierr.Append(err, enc.AddObject("request", v.Request))
	}
	return err
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *Dosa_CheckSchema_Args) GetRequest() (o *CheckSchemaRequest) {
	if v.Request != nil {
		return v.Request
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "checkSchema" for this struct.
func (v *Dosa_CheckSchema_Args) MethodName() string {
	return "checkSchema"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Dosa_CheckSchema_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Dosa_CheckSchema_Helper provides functions that aid in handling the
// parameters and return values of the Dosa.checkSchema
// function.
var Dosa_CheckSchema_Helper = struct {
	// Args accepts the parameters of checkSchema in-order and returns
	// the arguments struct for the function.
	Args func(
		request *CheckSchemaRequest,
	) *Dosa_CheckSchema_Args

	// IsException returns true if the given error can be thrown
	// by checkSchema.
	//
	// An error can be thrown by checkSchema only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for checkSchema
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// checkSchema into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by checkSchema
	//
	//   value, err := checkSchema(args)
	//   result, err := Dosa_CheckSchema_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from checkSchema: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*CheckSchemaResponse, error) (*Dosa_CheckSchema_Result, error)

	// UnwrapResponse takes the result struct for checkSchema
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if checkSchema threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Dosa_CheckSchema_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Dosa_CheckSchema_Result) (*CheckSchemaResponse, error)
}{}

func init() {
	Dosa_CheckSchema_Helper.Args = func(
		request *CheckSchemaRequest,
	) *Dosa_CheckSchema_Args {
		return &Dosa_CheckSchema_Args{
			Request: request,
		}
	}

	Dosa_CheckSchema_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BadRequestError:
			return true
		case *InternalServerError:
			return true
		case *BadSchemaError:
			return true
		case *RateLimitError:
			return true
		default:
			return false
		}
	}

	Dosa_CheckSchema_Helper.WrapResponse = func(success *CheckSchemaResponse, err error) (*Dosa_CheckSchema_Result, error) {
		if err == nil {
			return &Dosa_CheckSchema_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_CheckSchema_Result.ClientError")
			}
			return &Dosa_CheckSchema_Result{ClientError: e}, nil
		case *InternalServerError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_CheckSchema_Result.ServerError")
			}
			return &Dosa_CheckSchema_Result{ServerError: e}, nil
		case *BadSchemaError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_CheckSchema_Result.SchemaError")
			}
			return &Dosa_CheckSchema_Result{SchemaError: e}, nil
		case *RateLimitError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_CheckSchema_Result.LimitError")
			}
			return &Dosa_CheckSchema_Result{LimitError: e}, nil
		}

		return nil, err
	}
	Dosa_CheckSchema_Helper.UnwrapResponse = func(result *Dosa_CheckSchema_Result) (success *CheckSchemaResponse, err error) {
		if result.ClientError != nil {
			err = result.ClientError
			return
		}
		if result.ServerError != nil {
			err = result.ServerError
			return
		}
		if result.SchemaError != nil {
			err = result.SchemaError
			return
		}
		if result.LimitError != nil {
			err = result.LimitError
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

// Dosa_CheckSchema_Result represents the result of a Dosa.checkSchema function call.
//
// The result of a checkSchema execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Dosa_CheckSchema_Result struct {
	// Value returned by checkSchema after a successful execution.
	Success     *CheckSchemaResponse `json:"success,omitempty"`
	ClientError *BadRequestError     `json:"clientError,omitempty"`
	ServerError *InternalServerError `json:"serverError,omitempty"`
	SchemaError *BadSchemaError      `json:"schemaError,omitempty"`
	LimitError  *RateLimitError      `json:"limitError,omitempty"`
}

// ToWire translates a Dosa_CheckSchema_Result struct into a Thrift-level intermediate
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
func (v *Dosa_CheckSchema_Result) ToWire() (wire.Value, error) {
	var (
		fields [5]wire.Field
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
	if v.SchemaError != nil {
		w, err = v.SchemaError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.LimitError != nil {
		w, err = v.LimitError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Dosa_CheckSchema_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _CheckSchemaResponse_Read(w wire.Value) (*CheckSchemaResponse, error) {
	var v CheckSchemaResponse
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Dosa_CheckSchema_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_CheckSchema_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_CheckSchema_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_CheckSchema_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _CheckSchemaResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
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
				v.SchemaError, err = _BadSchemaError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 4:
			if field.Value.Type() == wire.TStruct {
				v.LimitError, err = _RateLimitError_Read(field.Value)
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
	if v.ClientError != nil {
		count++
	}
	if v.ServerError != nil {
		count++
	}
	if v.SchemaError != nil {
		count++
	}
	if v.LimitError != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Dosa_CheckSchema_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Dosa_CheckSchema_Result
// struct.
func (v *Dosa_CheckSchema_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [5]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.ClientError != nil {
		fields[i] = fmt.Sprintf("ClientError: %v", v.ClientError)
		i++
	}
	if v.ServerError != nil {
		fields[i] = fmt.Sprintf("ServerError: %v", v.ServerError)
		i++
	}
	if v.SchemaError != nil {
		fields[i] = fmt.Sprintf("SchemaError: %v", v.SchemaError)
		i++
	}
	if v.LimitError != nil {
		fields[i] = fmt.Sprintf("LimitError: %v", v.LimitError)
		i++
	}

	return fmt.Sprintf("Dosa_CheckSchema_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_CheckSchema_Result match the
// provided Dosa_CheckSchema_Result.
//
// This function performs a deep comparison.
func (v *Dosa_CheckSchema_Result) Equals(rhs *Dosa_CheckSchema_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.ClientError == nil && rhs.ClientError == nil) || (v.ClientError != nil && rhs.ClientError != nil && v.ClientError.Equals(rhs.ClientError))) {
		return false
	}
	if !((v.ServerError == nil && rhs.ServerError == nil) || (v.ServerError != nil && rhs.ServerError != nil && v.ServerError.Equals(rhs.ServerError))) {
		return false
	}
	if !((v.SchemaError == nil && rhs.SchemaError == nil) || (v.SchemaError != nil && rhs.SchemaError != nil && v.SchemaError.Equals(rhs.SchemaError))) {
		return false
	}
	if !((v.LimitError == nil && rhs.LimitError == nil) || (v.LimitError != nil && rhs.LimitError != nil && v.LimitError.Equals(rhs.LimitError))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Dosa_CheckSchema_Result.
func (v *Dosa_CheckSchema_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	if v.ClientError != nil {
		err = multierr.Append(err, enc.AddObject("clientError", v.ClientError))
	}
	if v.ServerError != nil {
		err = multierr.Append(err, enc.AddObject("serverError", v.ServerError))
	}
	if v.SchemaError != nil {
		err = multierr.Append(err, enc.AddObject("schemaError", v.SchemaError))
	}
	if v.LimitError != nil {
		err = multierr.Append(err, enc.AddObject("limitError", v.LimitError))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Dosa_CheckSchema_Result) GetSuccess() (o *CheckSchemaResponse) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// GetClientError returns the value of ClientError if it is set or its
// zero value if it is unset.
func (v *Dosa_CheckSchema_Result) GetClientError() (o *BadRequestError) {
	if v.ClientError != nil {
		return v.ClientError
	}

	return
}

// GetServerError returns the value of ServerError if it is set or its
// zero value if it is unset.
func (v *Dosa_CheckSchema_Result) GetServerError() (o *InternalServerError) {
	if v.ServerError != nil {
		return v.ServerError
	}

	return
}

// GetSchemaError returns the value of SchemaError if it is set or its
// zero value if it is unset.
func (v *Dosa_CheckSchema_Result) GetSchemaError() (o *BadSchemaError) {
	if v.SchemaError != nil {
		return v.SchemaError
	}

	return
}

// GetLimitError returns the value of LimitError if it is set or its
// zero value if it is unset.
func (v *Dosa_CheckSchema_Result) GetLimitError() (o *RateLimitError) {
	if v.LimitError != nil {
		return v.LimitError
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "checkSchema" for this struct.
func (v *Dosa_CheckSchema_Result) MethodName() string {
	return "checkSchema"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Dosa_CheckSchema_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
