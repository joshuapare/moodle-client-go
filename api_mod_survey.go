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


// ModSurveyAPIService ModSurveyAPI service
type ModSurveyAPIService service

type ApiModSurveyGetQuestionsRequest struct {
	ctx context.Context
	ApiService *ModSurveyAPIService
	modSurveyGetQuestionsRequest *ModSurveyGetQuestionsRequest
}

func (r ApiModSurveyGetQuestionsRequest) ModSurveyGetQuestionsRequest(modSurveyGetQuestionsRequest ModSurveyGetQuestionsRequest) ApiModSurveyGetQuestionsRequest {
	r.modSurveyGetQuestionsRequest = &modSurveyGetQuestionsRequest
	return r
}

func (r ApiModSurveyGetQuestionsRequest) Execute() (*ModSurveyGetQuestions200Response, *http.Response, error) {
	return r.ApiService.ModSurveyGetQuestionsExecute(r)
}

/*
ModSurveyGetQuestions Get the complete list of questions for the survey, including subquestions.

Get the complete list of questions for the survey, including subquestions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiModSurveyGetQuestionsRequest
*/
func (a *ModSurveyAPIService) ModSurveyGetQuestions(ctx context.Context) ApiModSurveyGetQuestionsRequest {
	return ApiModSurveyGetQuestionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModSurveyGetQuestions200Response
func (a *ModSurveyAPIService) ModSurveyGetQuestionsExecute(r ApiModSurveyGetQuestionsRequest) (*ModSurveyGetQuestions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModSurveyGetQuestions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ModSurveyAPIService.ModSurveyGetQuestions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mod_survey_get_questions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.modSurveyGetQuestionsRequest == nil {
		return localVarReturnValue, nil, reportError("modSurveyGetQuestionsRequest is required and must be specified")
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
	localVarPostBody = r.modSurveyGetQuestionsRequest
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

type ApiModSurveyGetSurveysByCoursesRequest struct {
	ctx context.Context
	ApiService *ModSurveyAPIService
	modChatGetChatsByCoursesRequest *ModChatGetChatsByCoursesRequest
}

func (r ApiModSurveyGetSurveysByCoursesRequest) ModChatGetChatsByCoursesRequest(modChatGetChatsByCoursesRequest ModChatGetChatsByCoursesRequest) ApiModSurveyGetSurveysByCoursesRequest {
	r.modChatGetChatsByCoursesRequest = &modChatGetChatsByCoursesRequest
	return r
}

func (r ApiModSurveyGetSurveysByCoursesRequest) Execute() (*ModSurveyGetSurveysByCourses200Response, *http.Response, error) {
	return r.ApiService.ModSurveyGetSurveysByCoursesExecute(r)
}

/*
ModSurveyGetSurveysByCourses Returns a list of survey instances in a provided set of courses,                             if no courses are provided then all the survey instances the user has access to will be returned.

Returns a list of survey instances in a provided set of courses,
                            if no courses are provided then all the survey instances the user has access to will be returned.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiModSurveyGetSurveysByCoursesRequest
*/
func (a *ModSurveyAPIService) ModSurveyGetSurveysByCourses(ctx context.Context) ApiModSurveyGetSurveysByCoursesRequest {
	return ApiModSurveyGetSurveysByCoursesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModSurveyGetSurveysByCourses200Response
func (a *ModSurveyAPIService) ModSurveyGetSurveysByCoursesExecute(r ApiModSurveyGetSurveysByCoursesRequest) (*ModSurveyGetSurveysByCourses200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModSurveyGetSurveysByCourses200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ModSurveyAPIService.ModSurveyGetSurveysByCourses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mod_survey_get_surveys_by_courses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.modChatGetChatsByCoursesRequest == nil {
		return localVarReturnValue, nil, reportError("modChatGetChatsByCoursesRequest is required and must be specified")
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
	localVarPostBody = r.modChatGetChatsByCoursesRequest
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

type ApiModSurveySubmitAnswersRequest struct {
	ctx context.Context
	ApiService *ModSurveyAPIService
	modSurveySubmitAnswersRequest *ModSurveySubmitAnswersRequest
}

func (r ApiModSurveySubmitAnswersRequest) ModSurveySubmitAnswersRequest(modSurveySubmitAnswersRequest ModSurveySubmitAnswersRequest) ApiModSurveySubmitAnswersRequest {
	r.modSurveySubmitAnswersRequest = &modSurveySubmitAnswersRequest
	return r
}

func (r ApiModSurveySubmitAnswersRequest) Execute() (*CoreCalendarDeleteSubscription200Response, *http.Response, error) {
	return r.ApiService.ModSurveySubmitAnswersExecute(r)
}

/*
ModSurveySubmitAnswers Submit the answers for a given survey.

Submit the answers for a given survey.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiModSurveySubmitAnswersRequest
*/
func (a *ModSurveyAPIService) ModSurveySubmitAnswers(ctx context.Context) ApiModSurveySubmitAnswersRequest {
	return ApiModSurveySubmitAnswersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreCalendarDeleteSubscription200Response
func (a *ModSurveyAPIService) ModSurveySubmitAnswersExecute(r ApiModSurveySubmitAnswersRequest) (*CoreCalendarDeleteSubscription200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreCalendarDeleteSubscription200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ModSurveyAPIService.ModSurveySubmitAnswers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mod_survey_submit_answers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.modSurveySubmitAnswersRequest == nil {
		return localVarReturnValue, nil, reportError("modSurveySubmitAnswersRequest is required and must be specified")
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
	localVarPostBody = r.modSurveySubmitAnswersRequest
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

type ApiModSurveyViewSurveyRequest struct {
	ctx context.Context
	ApiService *ModSurveyAPIService
	modSurveyViewSurveyRequest *ModSurveyViewSurveyRequest
}

func (r ApiModSurveyViewSurveyRequest) ModSurveyViewSurveyRequest(modSurveyViewSurveyRequest ModSurveyViewSurveyRequest) ApiModSurveyViewSurveyRequest {
	r.modSurveyViewSurveyRequest = &modSurveyViewSurveyRequest
	return r
}

func (r ApiModSurveyViewSurveyRequest) Execute() (*CoreCalendarDeleteSubscription200Response, *http.Response, error) {
	return r.ApiService.ModSurveyViewSurveyExecute(r)
}

/*
ModSurveyViewSurvey Trigger the course module viewed event and update the module completion status.

Trigger the course module viewed event and update the module completion status.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiModSurveyViewSurveyRequest
*/
func (a *ModSurveyAPIService) ModSurveyViewSurvey(ctx context.Context) ApiModSurveyViewSurveyRequest {
	return ApiModSurveyViewSurveyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreCalendarDeleteSubscription200Response
func (a *ModSurveyAPIService) ModSurveyViewSurveyExecute(r ApiModSurveyViewSurveyRequest) (*CoreCalendarDeleteSubscription200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreCalendarDeleteSubscription200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ModSurveyAPIService.ModSurveyViewSurvey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mod_survey_view_survey"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.modSurveyViewSurveyRequest == nil {
		return localVarReturnValue, nil, reportError("modSurveyViewSurveyRequest is required and must be specified")
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
	localVarPostBody = r.modSurveyViewSurveyRequest
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
