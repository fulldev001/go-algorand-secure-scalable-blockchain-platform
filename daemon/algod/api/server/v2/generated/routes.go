// Package generated provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get account information.
	// (GET /v2/accounts/{address})
	AccountInformation(ctx echo.Context, address string, params AccountInformationParams) error
	// Get a list of unconfirmed transactions currently in the transaction pool by address.
	// (GET /v2/accounts/{address}/transactions/pending)
	GetPendingTransactionsByAddress(ctx echo.Context, address string, params GetPendingTransactionsByAddressParams) error
	// Get application information.
	// (GET /v2/applications/{application-id})
	GetApplicationByID(ctx echo.Context, applicationId uint64) error
	// Get asset information.
	// (GET /v2/assets/{asset-id})
	GetAssetByID(ctx echo.Context, assetId uint64) error
	// Get the block for the given round.
	// (GET /v2/blocks/{round})
	GetBlock(ctx echo.Context, round uint64, params GetBlockParams) error
	// Get a Merkle proof for a transaction in a block.
	// (GET /v2/blocks/{round}/transactions/{txid}/proof)
	GetProof(ctx echo.Context, round uint64, txid string, params GetProofParams) error
	// Get the current supply reported by the ledger.
	// (GET /v2/ledger/supply)
	GetSupply(ctx echo.Context) error
	// Gets the current node status.
	// (GET /v2/status)
	GetStatus(ctx echo.Context) error
	// Gets the node status after waiting for the given round.
	// (GET /v2/status/wait-for-block-after/{round})
	WaitForBlock(ctx echo.Context, round uint64) error
	// Compile TEAL source code to binary, produce its hash
	// (POST /v2/teal/compile)
	TealCompile(ctx echo.Context) error
	// Provide debugging information for a transaction (or group).
	// (POST /v2/teal/dryrun)
	TealDryrun(ctx echo.Context) error
	// Broadcasts a raw transaction to the network.
	// (POST /v2/transactions)
	RawTransaction(ctx echo.Context) error
	// Get parameters for constructing a new transaction
	// (GET /v2/transactions/params)
	TransactionParams(ctx echo.Context) error
	// Get a list of unconfirmed transactions currently in the transaction pool.
	// (GET /v2/transactions/pending)
	GetPendingTransactions(ctx echo.Context, params GetPendingTransactionsParams) error
	// Get a specific pending transaction.
	// (GET /v2/transactions/pending/{txid})
	PendingTransactionInformation(ctx echo.Context, txid string, params PendingTransactionInformationParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AccountInformation converts echo context to params.
func (w *ServerInterfaceWrapper) AccountInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params AccountInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AccountInformation(ctx, address, params)
	return err
}

// GetPendingTransactionsByAddress converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactionsByAddress(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsByAddressParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactionsByAddress(ctx, address, params)
	return err
}

// GetApplicationByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetApplicationByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "application-id" -------------
	var applicationId uint64

	err = runtime.BindStyledParameter("simple", false, "application-id", ctx.Param("application-id"), &applicationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter application-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetApplicationByID(ctx, applicationId)
	return err
}

// GetAssetByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetAssetByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "asset-id" -------------
	var assetId uint64

	err = runtime.BindStyledParameter("simple", false, "asset-id", ctx.Param("asset-id"), &assetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAssetByID(ctx, assetId)
	return err
}

// GetBlock converts echo context to params.
func (w *ServerInterfaceWrapper) GetBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBlockParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBlock(ctx, round, params)
	return err
}

// GetProof converts echo context to params.
func (w *ServerInterfaceWrapper) GetProof(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameter("simple", false, "txid", ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetProofParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetProof(ctx, round, txid, params)
	return err
}

// GetSupply converts echo context to params.
func (w *ServerInterfaceWrapper) GetSupply(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSupply(ctx)
	return err
}

// GetStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetStatus(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStatus(ctx)
	return err
}

// WaitForBlock converts echo context to params.
func (w *ServerInterfaceWrapper) WaitForBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.WaitForBlock(ctx, round)
	return err
}

// TealCompile converts echo context to params.
func (w *ServerInterfaceWrapper) TealCompile(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TealCompile(ctx)
	return err
}

// TealDryrun converts echo context to params.
func (w *ServerInterfaceWrapper) TealDryrun(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TealDryrun(ctx)
	return err
}

// RawTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) RawTransaction(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RawTransaction(ctx)
	return err
}

// TransactionParams converts echo context to params.
func (w *ServerInterfaceWrapper) TransactionParams(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TransactionParams(ctx)
	return err
}

// GetPendingTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactions(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactions(ctx, params)
	return err
}

// PendingTransactionInformation converts echo context to params.
func (w *ServerInterfaceWrapper) PendingTransactionInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameter("simple", false, "txid", ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PendingTransactionInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PendingTransactionInformation(ctx, txid, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/v2/accounts/:address", wrapper.AccountInformation, m...)
	router.GET("/v2/accounts/:address/transactions/pending", wrapper.GetPendingTransactionsByAddress, m...)
	router.GET("/v2/applications/:application-id", wrapper.GetApplicationByID, m...)
	router.GET("/v2/assets/:asset-id", wrapper.GetAssetByID, m...)
	router.GET("/v2/blocks/:round", wrapper.GetBlock, m...)
	router.GET("/v2/blocks/:round/transactions/:txid/proof", wrapper.GetProof, m...)
	router.GET("/v2/ledger/supply", wrapper.GetSupply, m...)
	router.GET("/v2/status", wrapper.GetStatus, m...)
	router.GET("/v2/status/wait-for-block-after/:round", wrapper.WaitForBlock, m...)
	router.POST("/v2/teal/compile", wrapper.TealCompile, m...)
	router.POST("/v2/teal/dryrun", wrapper.TealDryrun, m...)
	router.POST("/v2/transactions", wrapper.RawTransaction, m...)
	router.GET("/v2/transactions/params", wrapper.TransactionParams, m...)
	router.GET("/v2/transactions/pending", wrapper.GetPendingTransactions, m...)
	router.GET("/v2/transactions/pending/:txid", wrapper.PendingTransactionInformation, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"pZTsHHMnfrj8fwEAAP///pnN5DLyAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
