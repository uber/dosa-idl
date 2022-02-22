// Code generated by thriftrw-plugin-yarpc
// @generated

package dosa

import yarpcerrors "go.uber.org/yarpc/yarpcerrors"

// YARPCErrorCode returns nil for BadRequestError.
//
// This is derived from the rpc.code annotation on the Thrift exception.
func (e *BadRequestError) YARPCErrorCode() *yarpcerrors.Code {

	return nil
}

// Name is the error name for BadRequestError.
func (e *BadRequestError) YARPCErrorName() string { return "BadRequestError" }

// YARPCErrorCode returns nil for BadSchemaError.
//
// This is derived from the rpc.code annotation on the Thrift exception.
func (e *BadSchemaError) YARPCErrorCode() *yarpcerrors.Code {

	return nil
}

// Name is the error name for BadSchemaError.
func (e *BadSchemaError) YARPCErrorName() string { return "BadSchemaError" }

// YARPCErrorCode returns nil for InternalServerError.
//
// This is derived from the rpc.code annotation on the Thrift exception.
func (e *InternalServerError) YARPCErrorCode() *yarpcerrors.Code {

	return nil
}

// Name is the error name for InternalServerError.
func (e *InternalServerError) YARPCErrorName() string { return "InternalServerError" }

// YARPCErrorCode returns nil for RateLimitError.
//
// This is derived from the rpc.code annotation on the Thrift exception.
func (e *RateLimitError) YARPCErrorCode() *yarpcerrors.Code {

	return nil
}

// Name is the error name for RateLimitError.
func (e *RateLimitError) YARPCErrorName() string { return "RateLimitError" }
