/*
Moodle Webservice API

Auto-generated OpenAPI spec for Moodle's Webservice API. Requires installation of the Catalyst RESTful API plugin: https://github.com/catalyst/moodle-webservice_restful.

API version: 4.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moodleclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ToolMobileAPIService ToolMobileAPI service
type ToolMobileAPIService service

type ApiToolMobileCallExternalFunctionsRequest struct {
	ctx context.Context
	ApiService *ToolMobileAPIService
	toolMobileCallExternalFunctionsRequest *ToolMobileCallExternalFunctionsRequest
}

func (r ApiToolMobileCallExternalFunctionsRequest) ToolMobileCallExternalFunctionsRequest(toolMobileCallExternalFunctionsRequest ToolMobileCallExternalFunctionsRequest) ApiToolMobileCallExternalFunctionsRequest {
	r.toolMobileCallExternalFunctionsRequest = &toolMobileCallExternalFunctionsRequest
	return r
}

func (r ApiToolMobileCallExternalFunctionsRequest) Execute() (*ToolMobileCallExternalFunctions200Response, *http.Response, error) {
	return r.ApiService.ToolMobileCallExternalFunctionsExecute(r)
}

/*
ToolMobileCallExternalFunctions Call multiple external functions and return all responses.

Call multiple external functions and return all responses.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiToolMobileCallExternalFunctionsRequest
*/
func (a *ToolMobileAPIService) ToolMobileCallExternalFunctions(ctx context.Context) ApiToolMobileCallExternalFunctionsRequest {
	return ApiToolMobileCallExternalFunctionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ToolMobileCallExternalFunctions200Response
func (a *ToolMobileAPIService) ToolMobileCallExternalFunctionsExecute(r ApiToolMobileCallExternalFunctionsRequest) (*ToolMobileCallExternalFunctions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolMobileCallExternalFunctions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolMobileAPIService.ToolMobileCallExternalFunctions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tool_mobile_call_external_functions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.toolMobileCallExternalFunctionsRequest == nil {
		return localVarReturnValue, nil, reportError("toolMobileCallExternalFunctionsRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.toolMobileCallExternalFunctionsRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiToolMobileGetAutologinKeyRequest struct {
	ctx context.Context
	ApiService *ToolMobileAPIService
	toolMobileGetAutologinKeyRequest *ToolMobileGetAutologinKeyRequest
}

func (r ApiToolMobileGetAutologinKeyRequest) ToolMobileGetAutologinKeyRequest(toolMobileGetAutologinKeyRequest ToolMobileGetAutologinKeyRequest) ApiToolMobileGetAutologinKeyRequest {
	r.toolMobileGetAutologinKeyRequest = &toolMobileGetAutologinKeyRequest
	return r
}

func (r ApiToolMobileGetAutologinKeyRequest) Execute() (*ToolMobileGetAutologinKey200Response, *http.Response, error) {
	return r.ApiService.ToolMobileGetAutologinKeyExecute(r)
}

/*
ToolMobileGetAutologinKey Creates an auto-login key for the current user.                             Is created only in https sites and is restricted by time, ip address and only works if the request                             comes from the Moodle mobile or desktop app.

Creates an auto-login key for the current user.
                            Is created only in https sites and is restricted by time, ip address and only works if the request
                            comes from the Moodle mobile or desktop app.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiToolMobileGetAutologinKeyRequest
*/
func (a *ToolMobileAPIService) ToolMobileGetAutologinKey(ctx context.Context) ApiToolMobileGetAutologinKeyRequest {
	return ApiToolMobileGetAutologinKeyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ToolMobileGetAutologinKey200Response
func (a *ToolMobileAPIService) ToolMobileGetAutologinKeyExecute(r ApiToolMobileGetAutologinKeyRequest) (*ToolMobileGetAutologinKey200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolMobileGetAutologinKey200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolMobileAPIService.ToolMobileGetAutologinKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tool_mobile_get_autologin_key"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.toolMobileGetAutologinKeyRequest == nil {
		return localVarReturnValue, nil, reportError("toolMobileGetAutologinKeyRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.toolMobileGetAutologinKeyRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiToolMobileGetConfigRequest struct {
	ctx context.Context
	ApiService *ToolMobileAPIService
	toolMobileGetConfigRequest *ToolMobileGetConfigRequest
}

func (r ApiToolMobileGetConfigRequest) ToolMobileGetConfigRequest(toolMobileGetConfigRequest ToolMobileGetConfigRequest) ApiToolMobileGetConfigRequest {
	r.toolMobileGetConfigRequest = &toolMobileGetConfigRequest
	return r
}

func (r ApiToolMobileGetConfigRequest) Execute() (*ToolMobileGetConfig200Response, *http.Response, error) {
	return r.ApiService.ToolMobileGetConfigExecute(r)
}

/*
ToolMobileGetConfig Returns a list of the site configurations, filtering by section.

Returns a list of the site configurations, filtering by section.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiToolMobileGetConfigRequest
*/
func (a *ToolMobileAPIService) ToolMobileGetConfig(ctx context.Context) ApiToolMobileGetConfigRequest {
	return ApiToolMobileGetConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ToolMobileGetConfig200Response
func (a *ToolMobileAPIService) ToolMobileGetConfigExecute(r ApiToolMobileGetConfigRequest) (*ToolMobileGetConfig200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolMobileGetConfig200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolMobileAPIService.ToolMobileGetConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tool_mobile_get_config"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.toolMobileGetConfigRequest == nil {
		return localVarReturnValue, nil, reportError("toolMobileGetConfigRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.toolMobileGetConfigRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiToolMobileGetContentRequest struct {
	ctx context.Context
	ApiService *ToolMobileAPIService
	toolMobileGetContentRequest *ToolMobileGetContentRequest
}

func (r ApiToolMobileGetContentRequest) ToolMobileGetContentRequest(toolMobileGetContentRequest ToolMobileGetContentRequest) ApiToolMobileGetContentRequest {
	r.toolMobileGetContentRequest = &toolMobileGetContentRequest
	return r
}

func (r ApiToolMobileGetContentRequest) Execute() (*ToolMobileGetContent200Response, *http.Response, error) {
	return r.ApiService.ToolMobileGetContentExecute(r)
}

/*
ToolMobileGetContent Returns a piece of content to be displayed in the Mobile app.

Returns a piece of content to be displayed in the Mobile app.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiToolMobileGetContentRequest
*/
func (a *ToolMobileAPIService) ToolMobileGetContent(ctx context.Context) ApiToolMobileGetContentRequest {
	return ApiToolMobileGetContentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ToolMobileGetContent200Response
func (a *ToolMobileAPIService) ToolMobileGetContentExecute(r ApiToolMobileGetContentRequest) (*ToolMobileGetContent200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolMobileGetContent200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolMobileAPIService.ToolMobileGetContent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tool_mobile_get_content"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.toolMobileGetContentRequest == nil {
		return localVarReturnValue, nil, reportError("toolMobileGetContentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.toolMobileGetContentRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiToolMobileGetPluginsSupportingMobileRequest struct {
	ctx context.Context
	ApiService *ToolMobileAPIService
}

func (r ApiToolMobileGetPluginsSupportingMobileRequest) Execute() (*ToolMobileGetPluginsSupportingMobile200Response, *http.Response, error) {
	return r.ApiService.ToolMobileGetPluginsSupportingMobileExecute(r)
}

/*
ToolMobileGetPluginsSupportingMobile Returns a list of Moodle plugins supporting the mobile app.

Returns a list of Moodle plugins supporting the mobile app.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiToolMobileGetPluginsSupportingMobileRequest
*/
func (a *ToolMobileAPIService) ToolMobileGetPluginsSupportingMobile(ctx context.Context) ApiToolMobileGetPluginsSupportingMobileRequest {
	return ApiToolMobileGetPluginsSupportingMobileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ToolMobileGetPluginsSupportingMobile200Response
func (a *ToolMobileAPIService) ToolMobileGetPluginsSupportingMobileExecute(r ApiToolMobileGetPluginsSupportingMobileRequest) (*ToolMobileGetPluginsSupportingMobile200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolMobileGetPluginsSupportingMobile200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolMobileAPIService.ToolMobileGetPluginsSupportingMobile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tool_mobile_get_plugins_supporting_mobile"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiToolMobileGetPublicConfigRequest struct {
	ctx context.Context
	ApiService *ToolMobileAPIService
}

func (r ApiToolMobileGetPublicConfigRequest) Execute() (*ToolMobileGetPublicConfig200Response, *http.Response, error) {
	return r.ApiService.ToolMobileGetPublicConfigExecute(r)
}

/*
ToolMobileGetPublicConfig Returns a list of the site public settings, those not requiring authentication.

Returns a list of the site public settings, those not requiring authentication.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiToolMobileGetPublicConfigRequest
*/
func (a *ToolMobileAPIService) ToolMobileGetPublicConfig(ctx context.Context) ApiToolMobileGetPublicConfigRequest {
	return ApiToolMobileGetPublicConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ToolMobileGetPublicConfig200Response
func (a *ToolMobileAPIService) ToolMobileGetPublicConfigExecute(r ApiToolMobileGetPublicConfigRequest) (*ToolMobileGetPublicConfig200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolMobileGetPublicConfig200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolMobileAPIService.ToolMobileGetPublicConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tool_mobile_get_public_config"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiToolMobileGetTokensForQrLoginRequest struct {
	ctx context.Context
	ApiService *ToolMobileAPIService
	toolMobileGetTokensForQrLoginRequest *ToolMobileGetTokensForQrLoginRequest
}

func (r ApiToolMobileGetTokensForQrLoginRequest) ToolMobileGetTokensForQrLoginRequest(toolMobileGetTokensForQrLoginRequest ToolMobileGetTokensForQrLoginRequest) ApiToolMobileGetTokensForQrLoginRequest {
	r.toolMobileGetTokensForQrLoginRequest = &toolMobileGetTokensForQrLoginRequest
	return r
}

func (r ApiToolMobileGetTokensForQrLoginRequest) Execute() (*ToolMobileGetTokensForQrLogin200Response, *http.Response, error) {
	return r.ApiService.ToolMobileGetTokensForQrLoginExecute(r)
}

/*
ToolMobileGetTokensForQrLogin Returns a WebService token (and private token) for QR login.

Returns a WebService token (and private token) for QR login.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiToolMobileGetTokensForQrLoginRequest
*/
func (a *ToolMobileAPIService) ToolMobileGetTokensForQrLogin(ctx context.Context) ApiToolMobileGetTokensForQrLoginRequest {
	return ApiToolMobileGetTokensForQrLoginRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ToolMobileGetTokensForQrLogin200Response
func (a *ToolMobileAPIService) ToolMobileGetTokensForQrLoginExecute(r ApiToolMobileGetTokensForQrLoginRequest) (*ToolMobileGetTokensForQrLogin200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolMobileGetTokensForQrLogin200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolMobileAPIService.ToolMobileGetTokensForQrLogin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tool_mobile_get_tokens_for_qr_login"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.toolMobileGetTokensForQrLoginRequest == nil {
		return localVarReturnValue, nil, reportError("toolMobileGetTokensForQrLoginRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.toolMobileGetTokensForQrLoginRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiToolMobileValidateSubscriptionKeyRequest struct {
	ctx context.Context
	ApiService *ToolMobileAPIService
	toolMobileValidateSubscriptionKeyRequest *ToolMobileValidateSubscriptionKeyRequest
}

func (r ApiToolMobileValidateSubscriptionKeyRequest) ToolMobileValidateSubscriptionKeyRequest(toolMobileValidateSubscriptionKeyRequest ToolMobileValidateSubscriptionKeyRequest) ApiToolMobileValidateSubscriptionKeyRequest {
	r.toolMobileValidateSubscriptionKeyRequest = &toolMobileValidateSubscriptionKeyRequest
	return r
}

func (r ApiToolMobileValidateSubscriptionKeyRequest) Execute() (*ToolMobileValidateSubscriptionKey200Response, *http.Response, error) {
	return r.ApiService.ToolMobileValidateSubscriptionKeyExecute(r)
}

/*
ToolMobileValidateSubscriptionKey Check if the given site subscription key is valid.

Check if the given site subscription key is valid.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiToolMobileValidateSubscriptionKeyRequest
*/
func (a *ToolMobileAPIService) ToolMobileValidateSubscriptionKey(ctx context.Context) ApiToolMobileValidateSubscriptionKeyRequest {
	return ApiToolMobileValidateSubscriptionKeyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ToolMobileValidateSubscriptionKey200Response
func (a *ToolMobileAPIService) ToolMobileValidateSubscriptionKeyExecute(r ApiToolMobileValidateSubscriptionKeyRequest) (*ToolMobileValidateSubscriptionKey200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolMobileValidateSubscriptionKey200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolMobileAPIService.ToolMobileValidateSubscriptionKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tool_mobile_validate_subscription_key"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.toolMobileValidateSubscriptionKeyRequest == nil {
		return localVarReturnValue, nil, reportError("toolMobileValidateSubscriptionKeyRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.toolMobileValidateSubscriptionKeyRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
