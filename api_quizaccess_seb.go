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


// QuizaccessSebAPIService QuizaccessSebAPI service
type QuizaccessSebAPIService service

type ApiQuizaccessSebValidateQuizKeysRequest struct {
	ctx context.Context
	ApiService *QuizaccessSebAPIService
	quizaccessSebValidateQuizKeysRequest *QuizaccessSebValidateQuizKeysRequest
}

func (r ApiQuizaccessSebValidateQuizKeysRequest) QuizaccessSebValidateQuizKeysRequest(quizaccessSebValidateQuizKeysRequest QuizaccessSebValidateQuizKeysRequest) ApiQuizaccessSebValidateQuizKeysRequest {
	r.quizaccessSebValidateQuizKeysRequest = &quizaccessSebValidateQuizKeysRequest
	return r
}

func (r ApiQuizaccessSebValidateQuizKeysRequest) Execute() (*QuizaccessSebValidateQuizKeys200Response, *http.Response, error) {
	return r.ApiService.QuizaccessSebValidateQuizKeysExecute(r)
}

/*
QuizaccessSebValidateQuizKeys Validate a Safe Exam Browser config key or a browser exam key.

Validate a Safe Exam Browser config key or a browser exam key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQuizaccessSebValidateQuizKeysRequest
*/
func (a *QuizaccessSebAPIService) QuizaccessSebValidateQuizKeys(ctx context.Context) ApiQuizaccessSebValidateQuizKeysRequest {
	return ApiQuizaccessSebValidateQuizKeysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QuizaccessSebValidateQuizKeys200Response
func (a *QuizaccessSebAPIService) QuizaccessSebValidateQuizKeysExecute(r ApiQuizaccessSebValidateQuizKeysRequest) (*QuizaccessSebValidateQuizKeys200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuizaccessSebValidateQuizKeys200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuizaccessSebAPIService.QuizaccessSebValidateQuizKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/quizaccess_seb_validate_quiz_keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.quizaccessSebValidateQuizKeysRequest == nil {
		return localVarReturnValue, nil, reportError("quizaccessSebValidateQuizKeysRequest is required and must be specified")
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
	localVarPostBody = r.quizaccessSebValidateQuizKeysRequest
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
