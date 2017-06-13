// Code generated by thriftrw v1.0.0
// @generated

package dosa

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Dosa_UpsertSchemaDryRun_Args struct {
	Request *UpsertSchemaDryRunRequest `json:"request,omitempty"`
}

func (v *Dosa_UpsertSchemaDryRun_Args) ToWire() (wire.Value, error) {
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

func _UpsertSchemaDryRunRequest_Read(w wire.Value) (*UpsertSchemaDryRunRequest, error) {
	var v UpsertSchemaDryRunRequest
	err := v.FromWire(w)
	return &v, err
}

func (v *Dosa_UpsertSchemaDryRun_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _UpsertSchemaDryRunRequest_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *Dosa_UpsertSchemaDryRun_Args) String() string {
	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}
	return fmt.Sprintf("Dosa_UpsertSchemaDryRun_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Dosa_UpsertSchemaDryRun_Args) MethodName() string {
	return "upsertSchemaDryRun"
}

func (v *Dosa_UpsertSchemaDryRun_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Dosa_UpsertSchemaDryRun_Helper = struct {
	Args           func(request *UpsertSchemaDryRunRequest) *Dosa_UpsertSchemaDryRun_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*Dosa_UpsertSchemaDryRun_Result, error)
	UnwrapResponse func(*Dosa_UpsertSchemaDryRun_Result) error
}{}

func init() {
	Dosa_UpsertSchemaDryRun_Helper.Args = func(request *UpsertSchemaDryRunRequest) *Dosa_UpsertSchemaDryRun_Args {
		return &Dosa_UpsertSchemaDryRun_Args{Request: request}
	}
	Dosa_UpsertSchemaDryRun_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BadRequestError:
			return true
		case *InternalServerError:
			return true
		case *BadSchemaError:
			return true
		default:
			return false
		}
	}
	Dosa_UpsertSchemaDryRun_Helper.WrapResponse = func(err error) (*Dosa_UpsertSchemaDryRun_Result, error) {
		if err == nil {
			return &Dosa_UpsertSchemaDryRun_Result{}, nil
		}
		switch e := err.(type) {
		case *BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_UpsertSchemaDryRun_Result.ClientError")
			}
			return &Dosa_UpsertSchemaDryRun_Result{ClientError: e}, nil
		case *InternalServerError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_UpsertSchemaDryRun_Result.ServerError")
			}
			return &Dosa_UpsertSchemaDryRun_Result{ServerError: e}, nil
		case *BadSchemaError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_UpsertSchemaDryRun_Result.SchemaError")
			}
			return &Dosa_UpsertSchemaDryRun_Result{SchemaError: e}, nil
		}
		return nil, err
	}
	Dosa_UpsertSchemaDryRun_Helper.UnwrapResponse = func(result *Dosa_UpsertSchemaDryRun_Result) (err error) {
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
		return
	}
}

type Dosa_UpsertSchemaDryRun_Result struct {
	ClientError *BadRequestError     `json:"clientError,omitempty"`
	ServerError *InternalServerError `json:"serverError,omitempty"`
	SchemaError *BadSchemaError      `json:"schemaError,omitempty"`
}

func (v *Dosa_UpsertSchemaDryRun_Result) ToWire() (wire.Value, error) {
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
	if v.SchemaError != nil {
		w, err = v.SchemaError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if i > 1 {
		return wire.Value{}, fmt.Errorf("Dosa_UpsertSchemaDryRun_Result should have at most one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Dosa_UpsertSchemaDryRun_Result) FromWire(w wire.Value) error {
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
				v.SchemaError, err = _BadSchemaError_Read(field.Value)
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
	if v.SchemaError != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("Dosa_UpsertSchemaDryRun_Result should have at most one field: got %v fields", count)
	}
	return nil
}

func (v *Dosa_UpsertSchemaDryRun_Result) String() string {
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
	if v.SchemaError != nil {
		fields[i] = fmt.Sprintf("SchemaError: %v", v.SchemaError)
		i++
	}
	return fmt.Sprintf("Dosa_UpsertSchemaDryRun_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Dosa_UpsertSchemaDryRun_Result) MethodName() string {
	return "upsertSchemaDryRun"
}

func (v *Dosa_UpsertSchemaDryRun_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}