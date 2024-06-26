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

// checks if the CoreSearchGetSearchAreasList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreSearchGetSearchAreasList200Response{}

// CoreSearchGetSearchAreasList200Response struct for CoreSearchGetSearchAreasList200Response
type CoreSearchGetSearchAreasList200Response struct {
	Areas []CoreSearchGetSearchAreasList200ResponseAreasInner `json:"areas"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _CoreSearchGetSearchAreasList200Response CoreSearchGetSearchAreasList200Response

// NewCoreSearchGetSearchAreasList200Response instantiates a new CoreSearchGetSearchAreasList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreSearchGetSearchAreasList200Response(areas []CoreSearchGetSearchAreasList200ResponseAreasInner) *CoreSearchGetSearchAreasList200Response {
	this := CoreSearchGetSearchAreasList200Response{}
	this.Areas = areas
	return &this
}

// NewCoreSearchGetSearchAreasList200ResponseWithDefaults instantiates a new CoreSearchGetSearchAreasList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreSearchGetSearchAreasList200ResponseWithDefaults() *CoreSearchGetSearchAreasList200Response {
	this := CoreSearchGetSearchAreasList200Response{}
	return &this
}

// GetAreas returns the Areas field value
func (o *CoreSearchGetSearchAreasList200Response) GetAreas() []CoreSearchGetSearchAreasList200ResponseAreasInner {
	if o == nil {
		var ret []CoreSearchGetSearchAreasList200ResponseAreasInner
		return ret
	}

	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value
// and a boolean to check if the value has been set.
func (o *CoreSearchGetSearchAreasList200Response) GetAreasOk() ([]CoreSearchGetSearchAreasList200ResponseAreasInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Areas, true
}

// SetAreas sets field value
func (o *CoreSearchGetSearchAreasList200Response) SetAreas(v []CoreSearchGetSearchAreasList200ResponseAreasInner) {
	o.Areas = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CoreSearchGetSearchAreasList200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetSearchAreasList200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CoreSearchGetSearchAreasList200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *CoreSearchGetSearchAreasList200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o CoreSearchGetSearchAreasList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreSearchGetSearchAreasList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["areas"] = o.Areas
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *CoreSearchGetSearchAreasList200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"areas",
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

	varCoreSearchGetSearchAreasList200Response := _CoreSearchGetSearchAreasList200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreSearchGetSearchAreasList200Response)

	if err != nil {
		return err
	}

	*o = CoreSearchGetSearchAreasList200Response(varCoreSearchGetSearchAreasList200Response)

	return err
}

type NullableCoreSearchGetSearchAreasList200Response struct {
	value *CoreSearchGetSearchAreasList200Response
	isSet bool
}

func (v NullableCoreSearchGetSearchAreasList200Response) Get() *CoreSearchGetSearchAreasList200Response {
	return v.value
}

func (v *NullableCoreSearchGetSearchAreasList200Response) Set(val *CoreSearchGetSearchAreasList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreSearchGetSearchAreasList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreSearchGetSearchAreasList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreSearchGetSearchAreasList200Response(val *CoreSearchGetSearchAreasList200Response) *NullableCoreSearchGetSearchAreasList200Response {
	return &NullableCoreSearchGetSearchAreasList200Response{value: val, isSet: true}
}

func (v NullableCoreSearchGetSearchAreasList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreSearchGetSearchAreasList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


