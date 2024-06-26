/*
Moodle Webservice API

Auto-generated OpenAPI spec for Moodle's Webservice API. Requires installation of the Catalyst RESTful API plugin: https://github.com/catalyst/moodle-webservice_restful.

API version: 4.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moodleclient

import (
	"encoding/json"
)

// checks if the CoreBlogViewEntriesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreBlogViewEntriesRequest{}

// CoreBlogViewEntriesRequest struct for CoreBlogViewEntriesRequest
type CoreBlogViewEntriesRequest struct {
	Filters []CoreBlogViewEntriesRequestFiltersInner `json:"filters,omitempty"`
}

// NewCoreBlogViewEntriesRequest instantiates a new CoreBlogViewEntriesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreBlogViewEntriesRequest() *CoreBlogViewEntriesRequest {
	this := CoreBlogViewEntriesRequest{}
	return &this
}

// NewCoreBlogViewEntriesRequestWithDefaults instantiates a new CoreBlogViewEntriesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreBlogViewEntriesRequestWithDefaults() *CoreBlogViewEntriesRequest {
	this := CoreBlogViewEntriesRequest{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *CoreBlogViewEntriesRequest) GetFilters() []CoreBlogViewEntriesRequestFiltersInner {
	if o == nil || IsNil(o.Filters) {
		var ret []CoreBlogViewEntriesRequestFiltersInner
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreBlogViewEntriesRequest) GetFiltersOk() ([]CoreBlogViewEntriesRequestFiltersInner, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *CoreBlogViewEntriesRequest) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []CoreBlogViewEntriesRequestFiltersInner and assigns it to the Filters field.
func (o *CoreBlogViewEntriesRequest) SetFilters(v []CoreBlogViewEntriesRequestFiltersInner) {
	o.Filters = v
}

func (o CoreBlogViewEntriesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreBlogViewEntriesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	return toSerialize, nil
}

type NullableCoreBlogViewEntriesRequest struct {
	value *CoreBlogViewEntriesRequest
	isSet bool
}

func (v NullableCoreBlogViewEntriesRequest) Get() *CoreBlogViewEntriesRequest {
	return v.value
}

func (v *NullableCoreBlogViewEntriesRequest) Set(val *CoreBlogViewEntriesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreBlogViewEntriesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreBlogViewEntriesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreBlogViewEntriesRequest(val *CoreBlogViewEntriesRequest) *NullableCoreBlogViewEntriesRequest {
	return &NullableCoreBlogViewEntriesRequest{value: val, isSet: true}
}

func (v NullableCoreBlogViewEntriesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreBlogViewEntriesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


