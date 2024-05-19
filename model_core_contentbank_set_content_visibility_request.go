/*
Moodle Webservice API

Auto-generated OpenAPI spec for Moodle's Webservice API. Requires installation of the Catalyst RESTful API plugin: https://github.com/catalyst/moodle-webservice_restful.

API version: 4.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moodleclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CoreContentbankSetContentVisibilityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreContentbankSetContentVisibilityRequest{}

// CoreContentbankSetContentVisibilityRequest struct for CoreContentbankSetContentVisibilityRequest
type CoreContentbankSetContentVisibilityRequest struct {
	// The content id to rename
	Contentid int32 `json:"contentid"`
	// The new visibility for the content
	Visibility int32 `json:"visibility"`
}

type _CoreContentbankSetContentVisibilityRequest CoreContentbankSetContentVisibilityRequest

// NewCoreContentbankSetContentVisibilityRequest instantiates a new CoreContentbankSetContentVisibilityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreContentbankSetContentVisibilityRequest(contentid int32, visibility int32) *CoreContentbankSetContentVisibilityRequest {
	this := CoreContentbankSetContentVisibilityRequest{}
	this.Contentid = contentid
	this.Visibility = visibility
	return &this
}

// NewCoreContentbankSetContentVisibilityRequestWithDefaults instantiates a new CoreContentbankSetContentVisibilityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreContentbankSetContentVisibilityRequestWithDefaults() *CoreContentbankSetContentVisibilityRequest {
	this := CoreContentbankSetContentVisibilityRequest{}
	var visibility int32 = null
	this.Visibility = visibility
	return &this
}

// GetContentid returns the Contentid field value
func (o *CoreContentbankSetContentVisibilityRequest) GetContentid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Contentid
}

// GetContentidOk returns a tuple with the Contentid field value
// and a boolean to check if the value has been set.
func (o *CoreContentbankSetContentVisibilityRequest) GetContentidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contentid, true
}

// SetContentid sets field value
func (o *CoreContentbankSetContentVisibilityRequest) SetContentid(v int32) {
	o.Contentid = v
}

// GetVisibility returns the Visibility field value
func (o *CoreContentbankSetContentVisibilityRequest) GetVisibility() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *CoreContentbankSetContentVisibilityRequest) GetVisibilityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *CoreContentbankSetContentVisibilityRequest) SetVisibility(v int32) {
	o.Visibility = v
}

func (o CoreContentbankSetContentVisibilityRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreContentbankSetContentVisibilityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contentid"] = o.Contentid
	toSerialize["visibility"] = o.Visibility
	return toSerialize, nil
}

func (o *CoreContentbankSetContentVisibilityRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contentid",
		"visibility",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCoreContentbankSetContentVisibilityRequest := _CoreContentbankSetContentVisibilityRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreContentbankSetContentVisibilityRequest)

	if err != nil {
		return err
	}

	*o = CoreContentbankSetContentVisibilityRequest(varCoreContentbankSetContentVisibilityRequest)

	return err
}

type NullableCoreContentbankSetContentVisibilityRequest struct {
	value *CoreContentbankSetContentVisibilityRequest
	isSet bool
}

func (v NullableCoreContentbankSetContentVisibilityRequest) Get() *CoreContentbankSetContentVisibilityRequest {
	return v.value
}

func (v *NullableCoreContentbankSetContentVisibilityRequest) Set(val *CoreContentbankSetContentVisibilityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreContentbankSetContentVisibilityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreContentbankSetContentVisibilityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreContentbankSetContentVisibilityRequest(val *CoreContentbankSetContentVisibilityRequest) *NullableCoreContentbankSetContentVisibilityRequest {
	return &NullableCoreContentbankSetContentVisibilityRequest{value: val, isSet: true}
}

func (v NullableCoreContentbankSetContentVisibilityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreContentbankSetContentVisibilityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

