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

// checks if the ToolMobileCallExternalFunctions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolMobileCallExternalFunctions200Response{}

// ToolMobileCallExternalFunctions200Response struct for ToolMobileCallExternalFunctions200Response
type ToolMobileCallExternalFunctions200Response struct {
	Responses []ToolMobileCallExternalFunctions200ResponseResponsesInner `json:"responses"`
}

type _ToolMobileCallExternalFunctions200Response ToolMobileCallExternalFunctions200Response

// NewToolMobileCallExternalFunctions200Response instantiates a new ToolMobileCallExternalFunctions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolMobileCallExternalFunctions200Response(responses []ToolMobileCallExternalFunctions200ResponseResponsesInner) *ToolMobileCallExternalFunctions200Response {
	this := ToolMobileCallExternalFunctions200Response{}
	this.Responses = responses
	return &this
}

// NewToolMobileCallExternalFunctions200ResponseWithDefaults instantiates a new ToolMobileCallExternalFunctions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolMobileCallExternalFunctions200ResponseWithDefaults() *ToolMobileCallExternalFunctions200Response {
	this := ToolMobileCallExternalFunctions200Response{}
	return &this
}

// GetResponses returns the Responses field value
func (o *ToolMobileCallExternalFunctions200Response) GetResponses() []ToolMobileCallExternalFunctions200ResponseResponsesInner {
	if o == nil {
		var ret []ToolMobileCallExternalFunctions200ResponseResponsesInner
		return ret
	}

	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value
// and a boolean to check if the value has been set.
func (o *ToolMobileCallExternalFunctions200Response) GetResponsesOk() ([]ToolMobileCallExternalFunctions200ResponseResponsesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Responses, true
}

// SetResponses sets field value
func (o *ToolMobileCallExternalFunctions200Response) SetResponses(v []ToolMobileCallExternalFunctions200ResponseResponsesInner) {
	o.Responses = v
}

func (o ToolMobileCallExternalFunctions200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolMobileCallExternalFunctions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["responses"] = o.Responses
	return toSerialize, nil
}

func (o *ToolMobileCallExternalFunctions200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"responses",
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

	varToolMobileCallExternalFunctions200Response := _ToolMobileCallExternalFunctions200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolMobileCallExternalFunctions200Response)

	if err != nil {
		return err
	}

	*o = ToolMobileCallExternalFunctions200Response(varToolMobileCallExternalFunctions200Response)

	return err
}

type NullableToolMobileCallExternalFunctions200Response struct {
	value *ToolMobileCallExternalFunctions200Response
	isSet bool
}

func (v NullableToolMobileCallExternalFunctions200Response) Get() *ToolMobileCallExternalFunctions200Response {
	return v.value
}

func (v *NullableToolMobileCallExternalFunctions200Response) Set(val *ToolMobileCallExternalFunctions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableToolMobileCallExternalFunctions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableToolMobileCallExternalFunctions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolMobileCallExternalFunctions200Response(val *ToolMobileCallExternalFunctions200Response) *NullableToolMobileCallExternalFunctions200Response {
	return &NullableToolMobileCallExternalFunctions200Response{value: val, isSet: true}
}

func (v NullableToolMobileCallExternalFunctions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolMobileCallExternalFunctions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


