// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"github.com/oapi-codegen/runtime"
)

// Counter defines model for Counter.
type Counter struct {
	// Counter The incremented counter value
	Counter int `json:"counter"`

	// Zone Zone data from Redis
	Zone string `json:"zone"`
}

// DeleteCounterResponse defines model for DeleteCounterResponse.
type DeleteCounterResponse = Counter

// Error standard error
type Error struct {
	// Detail Details about the error.
	Detail *string `json:"detail,omitempty"`

	// Instance The portal traceback code
	Instance string `json:"instance"`

	// InvalidParameters TODO
	InvalidParameters *[]InvalidParameters `json:"invalid_parameters,omitempty"`

	// Status The HTTP status code.
	Status int `json:"status"`

	// Title The error response code.
	Title string `json:"title"`

	// Type The error type.
	Type *string `json:"type,omitempty"`
}

// GetCounterResponse defines model for GetCounterResponse.
type GetCounterResponse = Counter

// InvalidParameters defines model for InvalidParameters.
type InvalidParameters struct {
	Choices *[]string `json:"choices,omitempty"`
	Field   *string   `json:"field,omitempty"`
	Reason  *string   `json:"reason,omitempty"`
	Rule    *string   `json:"rule,omitempty"`
}

// KV defines model for KV.
type KV struct {
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
	Value     string     `json:"value"`
}

// KVDeleteResponse defines model for KVDeleteResponse.
type KVDeleteResponse = KV

// KVGetResponse defines model for KVGetResponse.
type KVGetResponse = KV

// KVListResponse defines model for KVListResponse.
type KVListResponse struct {
	Keys []string `json:"keys"`
}

// KVPostRequest defines model for KVPostRequest.
type KVPostRequest struct {
	Expect *string `json:"expect,omitempty"`
	Value  string  `json:"value"`
}

// KVPostResponse defines model for KVPostResponse.
type KVPostResponse = KV

// PostCounterResponse defines model for PostCounterResponse.
type PostCounterResponse = Counter

// VersionResponse defines model for VersionResponse.
type VersionResponse struct {
	// Version Application version
	Version string `json:"version"`
}

// KvPostJSONRequestBody defines body for KvPost for application/json ContentType.
type KvPostJSONRequestBody = KVPostRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Reset the counter
	// (DELETE /counter)
	DeleteCounter(w http.ResponseWriter, r *http.Request)
	// Get the current counter value
	// (GET /counter)
	GetCounter(w http.ResponseWriter, r *http.Request)
	// Increment the counter
	// (POST /counter)
	PostCounter(w http.ResponseWriter, r *http.Request)
	// Returns all values currently available
	// (GET /key-value)
	KvList(w http.ResponseWriter, r *http.Request)
	// delete a value in a kv
	// (DELETE /key-value/{key})
	KvDelete(w http.ResponseWriter, r *http.Request, key string)
	// Returns the value for a key or 404 if not found
	// (GET /key-value/{key})
	KvGet(w http.ResponseWriter, r *http.Request, key string)
	// Set a value of a kv
	// (POST /key-value/{key})
	KvPost(w http.ResponseWriter, r *http.Request, key string)
	// Get the application version and color
	// (GET /version)
	GetVersion(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// DeleteCounter operation middleware
func (siw *ServerInterfaceWrapper) DeleteCounter(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteCounter(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetCounter operation middleware
func (siw *ServerInterfaceWrapper) GetCounter(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCounter(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostCounter operation middleware
func (siw *ServerInterfaceWrapper) PostCounter(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostCounter(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// KvList operation middleware
func (siw *ServerInterfaceWrapper) KvList(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.KvList(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// KvDelete operation middleware
func (siw *ServerInterfaceWrapper) KvDelete(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "key" -------------
	var key string

	err = runtime.BindStyledParameterWithOptions("simple", "key", mux.Vars(r)["key"], &key, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "key", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.KvDelete(w, r, key)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// KvGet operation middleware
func (siw *ServerInterfaceWrapper) KvGet(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "key" -------------
	var key string

	err = runtime.BindStyledParameterWithOptions("simple", "key", mux.Vars(r)["key"], &key, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "key", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.KvGet(w, r, key)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// KvPost operation middleware
func (siw *ServerInterfaceWrapper) KvPost(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "key" -------------
	var key string

	err = runtime.BindStyledParameterWithOptions("simple", "key", mux.Vars(r)["key"], &key, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "key", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.KvPost(w, r, key)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetVersion operation middleware
func (siw *ServerInterfaceWrapper) GetVersion(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetVersion(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{})
}

type GorillaServerOptions struct {
	BaseURL          string
	BaseRouter       *mux.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r *mux.Router) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r *mux.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options GorillaServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = mux.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.HandleFunc(options.BaseURL+"/counter", wrapper.DeleteCounter).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/counter", wrapper.GetCounter).Methods("GET")

	r.HandleFunc(options.BaseURL+"/counter", wrapper.PostCounter).Methods("POST")

	r.HandleFunc(options.BaseURL+"/key-value", wrapper.KvList).Methods("GET")

	r.HandleFunc(options.BaseURL+"/key-value/{key}", wrapper.KvDelete).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/key-value/{key}", wrapper.KvGet).Methods("GET")

	r.HandleFunc(options.BaseURL+"/key-value/{key}", wrapper.KvPost).Methods("POST")

	r.HandleFunc(options.BaseURL+"/version", wrapper.GetVersion).Methods("GET")

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xY3W7bOBN9FYLfB+yNbDlOsi10l/5sNmixDdKiF1sEi4k0tllLpJYcuXELv/uCpCzJ",
	"EuMmQA24QO9saUieOWf+qG88VUWpJEoyPPnGTbrAAtzPl6qShNr+LLUqUZNA9yJtX2RoUi1KEkryhH9Y",
	"IBMy1VigJMxYbchWkFfII07rEnnChSSco+abiH9VEof7/K0ksgwI2Eyrgt1gJky73JAWcs43m4hr/LcS",
	"GjOefGpQ1ZveNvbq7jOmxCN+P8J7KMrce/Ha/z7Z8Wh6usXEKzP6gob4xh70CnMkrBm5QVMqafAXM5uI",
	"v9ZaBRw2BDIDnTF076MeURkSiHy47JV7bhjcqYoYLdCvH/OI1whrMq1/aCyVhFCwL2CYVMRmqpLZkJCI",
	"C2kRpRiWplSaIGekIcU7SJcsVRnunOktEmeRTKfPn02fn08nz34/PT+Znk5OTsJnriAX2T8laCiQUJvA",
	"6e9eveMRF4SFe/t/jTOe8P/FbVrGdU7GV36/63a7TXMqaA1r+98QUGXCbv754cM18wbOwx1azyZnoTgk",
	"QfkDrDlpmK6zYbgj/0sR++MhRfyDh/e17x+9XS/iaxK26Dvy3zYe1ZH7yGRgJ92w7YVgptLKJvYgDNuo",
	"e0wAbaXzStQoA04viEqTxPFSyTmUYizkTMXLqoBRna2jDAsVS0UjD8Ql6iXSr/o1rF/DpBqyslAi9T+b",
	"PH0gnNs8nAnMs6ClRjCWjNCryufakI9tONRwWQdvn45NxN98HHpRlRkQZhdk/8yULoB4wu2zEYkCQyla",
	"L3mxDoL12gfRdtXzZrdBlL6tPhyPx4j5EunnAvxWmD2Il7h+UmT3DnbLw+deK3uuq5HDY/G+tIYHctkf",
	"/fOIZPH+Ks7D4vwRtRFKPszJyhsMfbkoy1ykYP+xrdH3XGntUpUrHVDKzleYVlrQ+r0dyjyIFwga9UVF",
	"i+YGZRfduce7bdu7ZRt2SEZh2MX1FSu1WokMDbOeOhcMmynNCpAwF3LOoJEYZMY0kha4si++bhUc87Zl",
	"vOzYOokvrq941FLHT8aT8cSGhSpRQil4wk/Hk/GpndqBFs7HeCf6bNUODe/2uXFj+xZgG0ldrGh2oTaO",
	"XmXNPi+bwNpOmA7IdDLxoSMJpctgaKWOP9ed1Y/M3xuow7c6J9GuZ++rNEVjZlVux10kpmZdLy135z8Q",
	"l59MAzgu2B1kbAYirzT6cKyKAvSaJ/zGAeuhmiMNdbppRHDWldZ2dnXVoueY02yPUu1MeUiZApPr9zRy",
	"DkLed8dXxGNQ63KrVc3+AGGpTEC6q22Z99r5xAoUg1rbPdJ1Ws4htQt1tv3iNZ3sGJOs4X8X2Sbi8RLX",
	"o2YQCOad7dO5MM6xJa6dXm6FFQuI4b0wNBDqzeqtf3wwjXoj4hNyy3pzJMWPKi0NgzzfUlpnVr5msAKR",
	"w12OPaHib0tcb3Y7Wp973yIOy37vFrSffw81sjo4j20clhpXQlWmrR5nk7PDC5KC/I3YvK5k9dk7sniw",
	"DOrmIiQDtlx1GlOf7ks8cKR3b29PCPQlro+I1puO9p5ZOx6CKypKs7PJGROzzoco20463zc+hQqThKLp",
	"/tbb7oe3dsYX1twOhTzidgVPuDdu52jSFUYdDvoz923b2/rq22ZR74WGXqhs/QOV715HA6xfV7T9ojdw",
	"Z3PQkNy5rO6PSUfc8cThe6Qmt9Wszm1bYjvXskdMoDC8q7ne6G5hzCCRkHMTnD4/Nle2gwnUv4M+oWoM",
	"nbE3P104GEc1hu6VwG9qUK/C5eOtSsE2phXmqnTDkbflEa90Xt98kzjOrd1CGUrOJ+eT2N40N7eb/wIA",
	"AP//8EkfIoIcAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
