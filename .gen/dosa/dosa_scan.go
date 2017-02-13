// Code generated by thriftrw v1.0.0
// @generated

package dosa

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Dosa_Scan_Args struct {
	Request *ScanRequest `json:"request,omitempty"`
}

func (v *Dosa_Scan_Args) ToWire() (wire.Value, error) {
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

func _ScanRequest_Read(w wire.Value) (*ScanRequest, error) {
	var v ScanRequest
	err := v.FromWire(w)
	return &v, err
}

func (v *Dosa_Scan_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _ScanRequest_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *Dosa_Scan_Args) String() string {
	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}
	return fmt.Sprintf("Dosa_Scan_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Dosa_Scan_Args) MethodName() string {
	return "scan"
}

func (v *Dosa_Scan_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Dosa_Scan_Helper = struct {
	Args           func(request *ScanRequest) *Dosa_Scan_Args
	IsException    func(error) bool
	WrapResponse   func(*ScanResponse, error) (*Dosa_Scan_Result, error)
	UnwrapResponse func(*Dosa_Scan_Result) (*ScanResponse, error)
}{}

func init() {
	Dosa_Scan_Helper.Args = func(request *ScanRequest) *Dosa_Scan_Args {
		return &Dosa_Scan_Args{Request: request}
	}
	Dosa_Scan_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BadRequestError:
			return true
		case *InternalServerError:
			return true
		default:
			return false
		}
	}
	Dosa_Scan_Helper.WrapResponse = func(success *ScanResponse, err error) (*Dosa_Scan_Result, error) {
		if err == nil {
			return &Dosa_Scan_Result{Success: success}, nil
		}
		switch e := err.(type) {
		case *BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_Scan_Result.ClientError")
			}
			return &Dosa_Scan_Result{ClientError: e}, nil
		case *InternalServerError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_Scan_Result.ServerError")
			}
			return &Dosa_Scan_Result{ServerError: e}, nil
		}
		return nil, err
	}
	Dosa_Scan_Helper.UnwrapResponse = func(result *Dosa_Scan_Result) (success *ScanResponse, err error) {
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

type Dosa_Scan_Result struct {
	Success     *ScanResponse        `json:"success,omitempty"`
	ClientError *BadRequestError     `json:"clientError,omitempty"`
	ServerError *InternalServerError `json:"serverError,omitempty"`
}

func (v *Dosa_Scan_Result) ToWire() (wire.Value, error) {
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
		return wire.Value{}, fmt.Errorf("Dosa_Scan_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ScanResponse_Read(w wire.Value) (*ScanResponse, error) {
	var v ScanResponse
	err := v.FromWire(w)
	return &v, err
}

func (v *Dosa_Scan_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _ScanResponse_Read(field.Value)
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
		return fmt.Errorf("Dosa_Scan_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *Dosa_Scan_Result) String() string {
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
	return fmt.Sprintf("Dosa_Scan_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Dosa_Scan_Result) MethodName() string {
	return "scan"
}

func (v *Dosa_Scan_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}