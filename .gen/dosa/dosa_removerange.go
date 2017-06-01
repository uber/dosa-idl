// Code generated by thriftrw v1.3.0
// @generated

package dosa

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Dosa_RemoveRange_Args struct {
	Request *RemoveRangeRequest `json:"request,omitempty"`
}

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

func (v *Dosa_RemoveRange_Args) Equals(rhs *Dosa_RemoveRange_Args) bool {
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}
	return true
}

func (v *Dosa_RemoveRange_Args) MethodName() string {
	return "removeRange"
}

func (v *Dosa_RemoveRange_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Dosa_RemoveRange_Helper = struct {
	Args           func(request *RemoveRangeRequest) *Dosa_RemoveRange_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*Dosa_RemoveRange_Result, error)
	UnwrapResponse func(*Dosa_RemoveRange_Result) error
}{}

func init() {
	Dosa_RemoveRange_Helper.Args = func(request *RemoveRangeRequest) *Dosa_RemoveRange_Args {
		return &Dosa_RemoveRange_Args{Request: request}
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

type Dosa_RemoveRange_Result struct {
	ClientError *BadRequestError     `json:"clientError,omitempty"`
	ServerError *InternalServerError `json:"serverError,omitempty"`
}

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

func (v *Dosa_RemoveRange_Result) Equals(rhs *Dosa_RemoveRange_Result) bool {
	if !((v.ClientError == nil && rhs.ClientError == nil) || (v.ClientError != nil && rhs.ClientError != nil && v.ClientError.Equals(rhs.ClientError))) {
		return false
	}
	if !((v.ServerError == nil && rhs.ServerError == nil) || (v.ServerError != nil && rhs.ServerError != nil && v.ServerError.Equals(rhs.ServerError))) {
		return false
	}
	return true
}

func (v *Dosa_RemoveRange_Result) MethodName() string {
	return "removeRange"
}

func (v *Dosa_RemoveRange_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
