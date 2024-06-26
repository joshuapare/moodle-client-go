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

// checks if the CoreEnrolSubmitUserEnrolmentForm200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreEnrolSubmitUserEnrolmentForm200Response{}

// CoreEnrolSubmitUserEnrolmentForm200Response struct for CoreEnrolSubmitUserEnrolmentForm200Response
type CoreEnrolSubmitUserEnrolmentForm200Response struct {
	// True if the user's enrolment was successfully updated
	Result bool `json:"result"`
	// Indicates invalid form data
	Validationerror *bool `json:"validationerror,omitempty"`
}

type _CoreEnrolSubmitUserEnrolmentForm200Response CoreEnrolSubmitUserEnrolmentForm200Response

// NewCoreEnrolSubmitUserEnrolmentForm200Response instantiates a new CoreEnrolSubmitUserEnrolmentForm200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreEnrolSubmitUserEnrolmentForm200Response(result bool) *CoreEnrolSubmitUserEnrolmentForm200Response {
	this := CoreEnrolSubmitUserEnrolmentForm200Response{}
	this.Result = result
	var validationerror bool = false
	this.Validationerror = &validationerror
	return &this
}

// NewCoreEnrolSubmitUserEnrolmentForm200ResponseWithDefaults instantiates a new CoreEnrolSubmitUserEnrolmentForm200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreEnrolSubmitUserEnrolmentForm200ResponseWithDefaults() *CoreEnrolSubmitUserEnrolmentForm200Response {
	this := CoreEnrolSubmitUserEnrolmentForm200Response{}
	var result bool = null
	this.Result = result
	var validationerror bool = false
	this.Validationerror = &validationerror
	return &this
}

// GetResult returns the Result field value
func (o *CoreEnrolSubmitUserEnrolmentForm200Response) GetResult() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CoreEnrolSubmitUserEnrolmentForm200Response) GetResultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CoreEnrolSubmitUserEnrolmentForm200Response) SetResult(v bool) {
	o.Result = v
}

// GetValidationerror returns the Validationerror field value if set, zero value otherwise.
func (o *CoreEnrolSubmitUserEnrolmentForm200Response) GetValidationerror() bool {
	if o == nil || IsNil(o.Validationerror) {
		var ret bool
		return ret
	}
	return *o.Validationerror
}

// GetValidationerrorOk returns a tuple with the Validationerror field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreEnrolSubmitUserEnrolmentForm200Response) GetValidationerrorOk() (*bool, bool) {
	if o == nil || IsNil(o.Validationerror) {
		return nil, false
	}
	return o.Validationerror, true
}

// HasValidationerror returns a boolean if a field has been set.
func (o *CoreEnrolSubmitUserEnrolmentForm200Response) HasValidationerror() bool {
	if o != nil && !IsNil(o.Validationerror) {
		return true
	}

	return false
}

// SetValidationerror gets a reference to the given bool and assigns it to the Validationerror field.
func (o *CoreEnrolSubmitUserEnrolmentForm200Response) SetValidationerror(v bool) {
	o.Validationerror = &v
}

func (o CoreEnrolSubmitUserEnrolmentForm200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreEnrolSubmitUserEnrolmentForm200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	if !IsNil(o.Validationerror) {
		toSerialize["validationerror"] = o.Validationerror
	}
	return toSerialize, nil
}

func (o *CoreEnrolSubmitUserEnrolmentForm200Response) UnmarshalJSON(data []byte) (err error) {
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

	varCoreEnrolSubmitUserEnrolmentForm200Response := _CoreEnrolSubmitUserEnrolmentForm200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreEnrolSubmitUserEnrolmentForm200Response)

	if err != nil {
		return err
	}

	*o = CoreEnrolSubmitUserEnrolmentForm200Response(varCoreEnrolSubmitUserEnrolmentForm200Response)

	return err
}

type NullableCoreEnrolSubmitUserEnrolmentForm200Response struct {
	value *CoreEnrolSubmitUserEnrolmentForm200Response
	isSet bool
}

func (v NullableCoreEnrolSubmitUserEnrolmentForm200Response) Get() *CoreEnrolSubmitUserEnrolmentForm200Response {
	return v.value
}

func (v *NullableCoreEnrolSubmitUserEnrolmentForm200Response) Set(val *CoreEnrolSubmitUserEnrolmentForm200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreEnrolSubmitUserEnrolmentForm200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreEnrolSubmitUserEnrolmentForm200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreEnrolSubmitUserEnrolmentForm200Response(val *CoreEnrolSubmitUserEnrolmentForm200Response) *NullableCoreEnrolSubmitUserEnrolmentForm200Response {
	return &NullableCoreEnrolSubmitUserEnrolmentForm200Response{value: val, isSet: true}
}

func (v NullableCoreEnrolSubmitUserEnrolmentForm200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreEnrolSubmitUserEnrolmentForm200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


