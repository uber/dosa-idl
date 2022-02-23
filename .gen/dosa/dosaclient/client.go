// Code generated by thriftrw-plugin-yarpc
// @generated

package dosaclient

import (
	context "context"
	dosa "github.com/uber/dosa-idl/.gen/dosa"
	tchannel "github.com/uber/tchannel-go"
	wire "go.uber.org/thriftrw/wire"
	yarpc "go.uber.org/yarpc"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	reflect "reflect"
)

// Interface is a client for the Dosa service.
type Interface interface {
	CanUpsertSchema(
		ctx context.Context,
		Request *dosa.CanUpsertSchemaRequest,
		opts ...yarpc.CallOption,
	) (*dosa.CanUpsertSchemaResponse, error)

	CheckSchema(
		ctx context.Context,
		Request *dosa.CheckSchemaRequest,
		opts ...yarpc.CallOption,
	) (*dosa.CheckSchemaResponse, error)

	CheckSchemaStatus(
		ctx context.Context,
		Request *dosa.CheckSchemaStatusRequest,
		opts ...yarpc.CallOption,
	) (*dosa.CheckSchemaStatusResponse, error)

	CreateIfNotExists(
		ctx context.Context,
		Request *dosa.CreateRequest,
		opts ...yarpc.CallOption,
	) error

	CreateScope(
		ctx context.Context,
		Request *dosa.CreateScopeRequest,
		opts ...yarpc.CallOption,
	) error

	DropScope(
		ctx context.Context,
		Request *dosa.DropScopeRequest,
		opts ...yarpc.CallOption,
	) error

	MultiRead(
		ctx context.Context,
		Request *dosa.MultiReadRequest,
		opts ...yarpc.CallOption,
	) (*dosa.MultiReadResponse, error)

	MultiRemove(
		ctx context.Context,
		Request *dosa.MultiRemoveRequest,
		opts ...yarpc.CallOption,
	) (*dosa.MultiRemoveResponse, error)

	MultiUpsert(
		ctx context.Context,
		Request *dosa.MultiUpsertRequest,
		opts ...yarpc.CallOption,
	) (*dosa.MultiUpsertResponse, error)

	Range(
		ctx context.Context,
		Request *dosa.RangeRequest,
		opts ...yarpc.CallOption,
	) (*dosa.RangeResponse, error)

	Read(
		ctx context.Context,
		Request *dosa.ReadRequest,
		opts ...yarpc.CallOption,
	) (*dosa.ReadResponse, error)

	Remove(
		ctx context.Context,
		Request *dosa.RemoveRequest,
		opts ...yarpc.CallOption,
	) error

	RemoveRange(
		ctx context.Context,
		Request *dosa.RemoveRangeRequest,
		opts ...yarpc.CallOption,
	) error

	Scan(
		ctx context.Context,
		Request *dosa.ScanRequest,
		opts ...yarpc.CallOption,
	) (*dosa.ScanResponse, error)

	ScopeExists(
		ctx context.Context,
		Request *dosa.ScopeExistsRequest,
		opts ...yarpc.CallOption,
	) (*dosa.ScopeExistsResponse, error)

	Search(
		ctx context.Context,
		Request *dosa.SearchRequest,
		opts ...yarpc.CallOption,
	) (*dosa.SearchResponse, error)

	TruncateScope(
		ctx context.Context,
		Request *dosa.TruncateScopeRequest,
		opts ...yarpc.CallOption,
	) error

	Upsert(
		ctx context.Context,
		Request *dosa.UpsertRequest,
		opts ...yarpc.CallOption,
	) error

	UpsertSchema(
		ctx context.Context,
		Request *dosa.UpsertSchemaRequest,
		opts ...yarpc.CallOption,
	) (*dosa.UpsertSchemaResponse, error)
}

// New builds a new client for the Dosa service.
//
// 	client := dosaclient.New(dispatcher.ClientConfig("dosa"))
func New(c transport.ClientConfig, opts ...thrift.ClientOption) Interface {
	return client{
		c: thrift.New(thrift.Config{
			Service:      "Dosa",
			ClientConfig: c,
		}, opts...),
		nwc: thrift.NewNoWire(thrift.Config{
			Service:      "Dosa",
			ClientConfig: c,
		}, opts...),
	}
}

func init() {
	yarpc.RegisterClientBuilder(
		func(c transport.ClientConfig, f reflect.StructField) Interface {
			return New(c, thrift.ClientBuilderOptions(c, f)...)
		},
	)
}

type client struct {
	c   thrift.Client
	nwc thrift.NoWireClient
}

func (c client) CanUpsertSchema(
	ctx context.Context,
	_Request *dosa.CanUpsertSchemaRequest,
	opts ...yarpc.CallOption,
) (success *dosa.CanUpsertSchemaResponse, err error) {

	var result dosa.Dosa_CanUpsertSchema_Result
	args := dosa.Dosa_CanUpsertSchema_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = dosa.Dosa_CanUpsertSchema_Helper.UnwrapResponse(&result)
	return
}

func (c client) CheckSchema(
	ctx context.Context,
	_Request *dosa.CheckSchemaRequest,
	opts ...yarpc.CallOption,
) (success *dosa.CheckSchemaResponse, err error) {

	var result dosa.Dosa_CheckSchema_Result
	args := dosa.Dosa_CheckSchema_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = dosa.Dosa_CheckSchema_Helper.UnwrapResponse(&result)
	return
}

func (c client) CheckSchemaStatus(
	ctx context.Context,
	_Request *dosa.CheckSchemaStatusRequest,
	opts ...yarpc.CallOption,
) (success *dosa.CheckSchemaStatusResponse, err error) {

	var result dosa.Dosa_CheckSchemaStatus_Result
	args := dosa.Dosa_CheckSchemaStatus_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = dosa.Dosa_CheckSchemaStatus_Helper.UnwrapResponse(&result)
	return
}

func (c client) CreateIfNotExists(
	ctx context.Context,
	_Request *dosa.CreateRequest,
	opts ...yarpc.CallOption,
) (err error) {

	var result dosa.Dosa_CreateIfNotExists_Result
	args := dosa.Dosa_CreateIfNotExists_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	err = dosa.Dosa_CreateIfNotExists_Helper.UnwrapResponse(&result)
	return
}

func (c client) CreateScope(
	ctx context.Context,
	_Request *dosa.CreateScopeRequest,
	opts ...yarpc.CallOption,
) (err error) {

	var result dosa.Dosa_CreateScope_Result
	args := dosa.Dosa_CreateScope_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	err = dosa.Dosa_CreateScope_Helper.UnwrapResponse(&result)
	return
}

func (c client) DropScope(
	ctx context.Context,
	_Request *dosa.DropScopeRequest,
	opts ...yarpc.CallOption,
) (err error) {

	var result dosa.Dosa_DropScope_Result
	args := dosa.Dosa_DropScope_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	err = dosa.Dosa_DropScope_Helper.UnwrapResponse(&result)
	return
}

func (c client) MultiRead(
	ctx context.Context,
	_Request *dosa.MultiReadRequest,
	opts ...yarpc.CallOption,
) (success *dosa.MultiReadResponse, err error) {

	var result dosa.Dosa_MultiRead_Result
	args := dosa.Dosa_MultiRead_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = dosa.Dosa_MultiRead_Helper.UnwrapResponse(&result)
	return
}

func (c client) MultiRemove(
	ctx context.Context,
	_Request *dosa.MultiRemoveRequest,
	opts ...yarpc.CallOption,
) (success *dosa.MultiRemoveResponse, err error) {

	var result dosa.Dosa_MultiRemove_Result
	args := dosa.Dosa_MultiRemove_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = dosa.Dosa_MultiRemove_Helper.UnwrapResponse(&result)
	return
}

func (c client) MultiUpsert(
	ctx context.Context,
	_Request *dosa.MultiUpsertRequest,
	opts ...yarpc.CallOption,
) (success *dosa.MultiUpsertResponse, err error) {

	var result dosa.Dosa_MultiUpsert_Result
	args := dosa.Dosa_MultiUpsert_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = dosa.Dosa_MultiUpsert_Helper.UnwrapResponse(&result)
	return
}

func (c client) Range(
	ctx context.Context,
	_Request *dosa.RangeRequest,
	opts ...yarpc.CallOption,
) (success *dosa.RangeResponse, err error) {

	var result dosa.Dosa_Range_Result
	args := dosa.Dosa_Range_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = dosa.Dosa_Range_Helper.UnwrapResponse(&result)
	return
}

func (c client) Read(
	ctx context.Context,
	_Request *dosa.ReadRequest,
	opts ...yarpc.CallOption,
) (success *dosa.ReadResponse, err error) {

	var result dosa.Dosa_Read_Result
	args := dosa.Dosa_Read_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = dosa.Dosa_Read_Helper.UnwrapResponse(&result)
	return
}

func (c client) Remove(
	ctx context.Context,
	_Request *dosa.RemoveRequest,
	opts ...yarpc.CallOption,
) (err error) {

	var result dosa.Dosa_Remove_Result
	args := dosa.Dosa_Remove_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	err = dosa.Dosa_Remove_Helper.UnwrapResponse(&result)
	return
}

func (c client) RemoveRange(
	ctx context.Context,
	_Request *dosa.RemoveRangeRequest,
	opts ...yarpc.CallOption,
) (err error) {

	var result dosa.Dosa_RemoveRange_Result
	args := dosa.Dosa_RemoveRange_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	err = dosa.Dosa_RemoveRange_Helper.UnwrapResponse(&result)
	return
}

func (c client) Scan(
	ctx context.Context,
	_Request *dosa.ScanRequest,
	opts ...yarpc.CallOption,
) (success *dosa.ScanResponse, err error) {

	var result dosa.Dosa_Scan_Result
	args := dosa.Dosa_Scan_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = dosa.Dosa_Scan_Helper.UnwrapResponse(&result)
	return
}

func (c client) ScopeExists(
	ctx context.Context,
	_Request *dosa.ScopeExistsRequest,
	opts ...yarpc.CallOption,
) (success *dosa.ScopeExistsResponse, err error) {

	var result dosa.Dosa_ScopeExists_Result
	args := dosa.Dosa_ScopeExists_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = dosa.Dosa_ScopeExists_Helper.UnwrapResponse(&result)
	return
}

func (c client) Search(
	ctx context.Context,
	_Request *dosa.SearchRequest,
	opts ...yarpc.CallOption,
) (success *dosa.SearchResponse, err error) {

	var result dosa.Dosa_Search_Result
	args := dosa.Dosa_Search_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = dosa.Dosa_Search_Helper.UnwrapResponse(&result)
	return
}

func (c client) TruncateScope(
	ctx context.Context,
	_Request *dosa.TruncateScopeRequest,
	opts ...yarpc.CallOption,
) (err error) {

	var result dosa.Dosa_TruncateScope_Result
	args := dosa.Dosa_TruncateScope_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	err = dosa.Dosa_TruncateScope_Helper.UnwrapResponse(&result)
	return
}

func (c client) Upsert(
	ctx context.Context,
	_Request *dosa.UpsertRequest,
	opts ...yarpc.CallOption,
) (err error) {

	var result dosa.Dosa_Upsert_Result
	args := dosa.Dosa_Upsert_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	err = dosa.Dosa_Upsert_Helper.UnwrapResponse(&result)
	return
}

func (c client) UpsertSchema(
	ctx context.Context,
	_Request *dosa.UpsertSchemaRequest,
	opts ...yarpc.CallOption,
) (success *dosa.UpsertSchemaResponse, err error) {

	var result dosa.Dosa_UpsertSchema_Result
	args := dosa.Dosa_UpsertSchema_Helper.Args(_Request)

	ctx = tchannel.WithoutHeaders(ctx)
	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = dosa.Dosa_UpsertSchema_Helper.UnwrapResponse(&result)
	return
}
