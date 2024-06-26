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

// checks if the ToolDataprivacySetContextlevelForm200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolDataprivacySetContextlevelForm200Response{}

// ToolDataprivacySetContextlevelForm200Response struct for ToolDataprivacySetContextlevelForm200Response
type ToolDataprivacySetContextlevelForm200Response struct {
	// Whether the data was properly set or not
	Result bool `json:"result"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ToolDataprivacySetContextlevelForm200Response ToolDataprivacySetContextlevelForm200Response

// NewToolDataprivacySetContextlevelForm200Response instantiates a new ToolDataprivacySetContextlevelForm200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolDataprivacySetContextlevelForm200Response(result bool) *ToolDataprivacySetContextlevelForm200Response {
	this := ToolDataprivacySetContextlevelForm200Response{}
	this.Result = result
	return &this
}

// NewToolDataprivacySetContextlevelForm200ResponseWithDefaults instantiates a new ToolDataprivacySetContextlevelForm200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolDataprivacySetContextlevelForm200ResponseWithDefaults() *ToolDataprivacySetContextlevelForm200Response {
	this := ToolDataprivacySetContextlevelForm200Response{}
	return &this
}

// GetResult returns the Result field value
func (o *ToolDataprivacySetContextlevelForm200Response) GetResult() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacySetContextlevelForm200Response) GetResultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *ToolDataprivacySetContextlevelForm200Response) SetResult(v bool) {
	o.Result = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ToolDataprivacySetContextlevelForm200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolDataprivacySetContextlevelForm200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ToolDataprivacySetContextlevelForm200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ToolDataprivacySetContextlevelForm200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ToolDataprivacySetContextlevelForm200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolDataprivacySetContextlevelForm200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ToolDataprivacySetContextlevelForm200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
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

	varToolDataprivacySetContextlevelForm200Response := _ToolDataprivacySetContextlevelForm200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolDataprivacySetContextlevelForm200Response)

	if err != nil {
		return err
	}

	*o = ToolDataprivacySetContextlevelForm200Response(varToolDataprivacySetContextlevelForm200Response)

	return err
}

type NullableToolDataprivacySetContextlevelForm200Response struct {
	value *ToolDataprivacySetContextlevelForm200Response
	isSet bool
}

func (v NullableToolDataprivacySetContextlevelForm200Response) Get() *ToolDataprivacySetContextlevelForm200Response {
	return v.value
}

func (v *NullableToolDataprivacySetContextlevelForm200Response) Set(val *ToolDataprivacySetContextlevelForm200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableToolDataprivacySetContextlevelForm200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableToolDataprivacySetContextlevelForm200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolDataprivacySetContextlevelForm200Response(val *ToolDataprivacySetContextlevelForm200Response) *NullableToolDataprivacySetContextlevelForm200Response {
	return &NullableToolDataprivacySetContextlevelForm200Response{value: val, isSet: true}
}

func (v NullableToolDataprivacySetContextlevelForm200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolDataprivacySetContextlevelForm200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


