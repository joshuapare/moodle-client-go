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

// checks if the ToolDataprivacySetContextDefaults200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolDataprivacySetContextDefaults200Response{}

// ToolDataprivacySetContextDefaults200Response struct for ToolDataprivacySetContextDefaults200Response
type ToolDataprivacySetContextDefaults200Response struct {
	// Whether the context defaults were successfully set or not
	Result bool `json:"result"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ToolDataprivacySetContextDefaults200Response ToolDataprivacySetContextDefaults200Response

// NewToolDataprivacySetContextDefaults200Response instantiates a new ToolDataprivacySetContextDefaults200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolDataprivacySetContextDefaults200Response(result bool) *ToolDataprivacySetContextDefaults200Response {
	this := ToolDataprivacySetContextDefaults200Response{}
	this.Result = result
	return &this
}

// NewToolDataprivacySetContextDefaults200ResponseWithDefaults instantiates a new ToolDataprivacySetContextDefaults200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolDataprivacySetContextDefaults200ResponseWithDefaults() *ToolDataprivacySetContextDefaults200Response {
	this := ToolDataprivacySetContextDefaults200Response{}
	var result bool = null
	this.Result = result
	return &this
}

// GetResult returns the Result field value
func (o *ToolDataprivacySetContextDefaults200Response) GetResult() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacySetContextDefaults200Response) GetResultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *ToolDataprivacySetContextDefaults200Response) SetResult(v bool) {
	o.Result = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ToolDataprivacySetContextDefaults200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolDataprivacySetContextDefaults200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ToolDataprivacySetContextDefaults200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ToolDataprivacySetContextDefaults200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ToolDataprivacySetContextDefaults200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolDataprivacySetContextDefaults200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ToolDataprivacySetContextDefaults200Response) UnmarshalJSON(data []byte) (err error) {
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

	varToolDataprivacySetContextDefaults200Response := _ToolDataprivacySetContextDefaults200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolDataprivacySetContextDefaults200Response)

	if err != nil {
		return err
	}

	*o = ToolDataprivacySetContextDefaults200Response(varToolDataprivacySetContextDefaults200Response)

	return err
}

type NullableToolDataprivacySetContextDefaults200Response struct {
	value *ToolDataprivacySetContextDefaults200Response
	isSet bool
}

func (v NullableToolDataprivacySetContextDefaults200Response) Get() *ToolDataprivacySetContextDefaults200Response {
	return v.value
}

func (v *NullableToolDataprivacySetContextDefaults200Response) Set(val *ToolDataprivacySetContextDefaults200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableToolDataprivacySetContextDefaults200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableToolDataprivacySetContextDefaults200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolDataprivacySetContextDefaults200Response(val *ToolDataprivacySetContextDefaults200Response) *NullableToolDataprivacySetContextDefaults200Response {
	return &NullableToolDataprivacySetContextDefaults200Response{value: val, isSet: true}
}

func (v NullableToolDataprivacySetContextDefaults200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolDataprivacySetContextDefaults200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


