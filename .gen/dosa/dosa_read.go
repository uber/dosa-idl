// Code generated by thriftrw v1.1.0
// @generated

package dosa

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Dosa_Read_Args struct {
	Request *ReadRequest `json:"request,omitempty"`
}

func (v *Dosa_Read_Args) ToWire() (wire.Value, error) {
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

func _ReadRequest_Read(w wire.Value) (*ReadRequest, error) {
	var v ReadRequest
	err := v.FromWire(w)
	return &v, err
}

func (v *Dosa_Read_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _ReadRequest_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *Dosa_Read_Args) String() string {
	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}
	return fmt.Sprintf("Dosa_Read_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Dosa_Read_Args) MethodName() string {
	return "read"
}

func (v *Dosa_Read_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Dosa_Read_Helper = struct {
	Args           func(request *ReadRequest) *Dosa_Read_Args
	IsException    func(error) bool
	WrapResponse   func(*ReadResponse, error) (*Dosa_Read_Result, error)
	UnwrapResponse func(*Dosa_Read_Result) (*ReadResponse, error)
}{}

func init() {
	Dosa_Read_Helper.Args = func(request *ReadRequest) *Dosa_Read_Args {
		return &Dosa_Read_Args{Request: request}
	}
	Dosa_Read_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BadRequestError:
			return true
		case *InternalServerError:
			return true
		default:
			return false
		}
	}
	Dosa_Read_Helper.WrapResponse = func(success *ReadResponse, err error) (*Dosa_Read_Result, error) {
		if err == nil {
			return &Dosa_Read_Result{Success: success}, nil
		}
		switch e := err.(type) {
		case *BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_Read_Result.ClientError")
			}
			return &Dosa_Read_Result{ClientError: e}, nil
		case *InternalServerError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_Read_Result.ServerError")
			}
			return &Dosa_Read_Result{ServerError: e}, nil
		}
		return nil, err
	}
	Dosa_Read_Helper.UnwrapResponse = func(result *Dosa_Read_Result) (success *ReadResponse, err error) {
		if result.ClientError != nil {
			err = result.ClientError
			return
		}
		if result.ServerError != nil {
			err = result.ServerError
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

type Dosa_Read_Result struct {
	Success     *ReadResponse        `json:"success,omitempty"`
	ClientError *BadRequestError     `json:"clientError,omitempty"`
	ServerError *InternalServerError `json:"serverError,omitempty"`
}

func (v *Dosa_Read_Result) ToWire() (wire.Value, error) {
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
	if i != 1 {
		return wire.Value{}, fmt.Errorf("Dosa_Read_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ReadResponse_Read(w wire.Value) (*ReadResponse, error) {
	var v ReadResponse
	err := v.FromWire(w)
	return &v, err
}

func (v *Dosa_Read_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _ReadResponse_Read(field.Value)
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
	if count != 1 {
		return fmt.Errorf("Dosa_Read_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *Dosa_Read_Result) String() string {
	var fields [3]string
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
	return fmt.Sprintf("Dosa_Read_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Dosa_Read_Result) MethodName() string {
	return "read"
}

func (v *Dosa_Read_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
