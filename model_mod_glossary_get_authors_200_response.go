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

// checks if the ModGlossaryGetAuthors200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModGlossaryGetAuthors200Response{}

// ModGlossaryGetAuthors200Response struct for ModGlossaryGetAuthors200Response
type ModGlossaryGetAuthors200Response struct {
	Authors []ModGlossaryGetAuthors200ResponseAuthorsInner `json:"authors"`
	// The total number of records.
	Count int32 `json:"count"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModGlossaryGetAuthors200Response ModGlossaryGetAuthors200Response

// NewModGlossaryGetAuthors200Response instantiates a new ModGlossaryGetAuthors200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModGlossaryGetAuthors200Response(authors []ModGlossaryGetAuthors200ResponseAuthorsInner, count int32) *ModGlossaryGetAuthors200Response {
	this := ModGlossaryGetAuthors200Response{}
	this.Authors = authors
	this.Count = count
	return &this
}

// NewModGlossaryGetAuthors200ResponseWithDefaults instantiates a new ModGlossaryGetAuthors200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModGlossaryGetAuthors200ResponseWithDefaults() *ModGlossaryGetAuthors200Response {
	this := ModGlossaryGetAuthors200Response{}
	var count int32 = null
	this.Count = count
	return &this
}

// GetAuthors returns the Authors field value
func (o *ModGlossaryGetAuthors200Response) GetAuthors() []ModGlossaryGetAuthors200ResponseAuthorsInner {
	if o == nil {
		var ret []ModGlossaryGetAuthors200ResponseAuthorsInner
		return ret
	}

	return o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetAuthors200Response) GetAuthorsOk() ([]ModGlossaryGetAuthors200ResponseAuthorsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Authors, true
}

// SetAuthors sets field value
func (o *ModGlossaryGetAuthors200Response) SetAuthors(v []ModGlossaryGetAuthors200ResponseAuthorsInner) {
	o.Authors = v
}

// GetCount returns the Count field value
func (o *ModGlossaryGetAuthors200Response) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetAuthors200Response) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ModGlossaryGetAuthors200Response) SetCount(v int32) {
	o.Count = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModGlossaryGetAuthors200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetAuthors200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModGlossaryGetAuthors200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModGlossaryGetAuthors200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModGlossaryGetAuthors200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModGlossaryGetAuthors200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authors"] = o.Authors
	toSerialize["count"] = o.Count
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModGlossaryGetAuthors200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authors",
		"count",
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

	varModGlossaryGetAuthors200Response := _ModGlossaryGetAuthors200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModGlossaryGetAuthors200Response)

	if err != nil {
		return err
	}

	*o = ModGlossaryGetAuthors200Response(varModGlossaryGetAuthors200Response)

	return err
}

type NullableModGlossaryGetAuthors200Response struct {
	value *ModGlossaryGetAuthors200Response
	isSet bool
}

func (v NullableModGlossaryGetAuthors200Response) Get() *ModGlossaryGetAuthors200Response {
	return v.value
}

func (v *NullableModGlossaryGetAuthors200Response) Set(val *ModGlossaryGetAuthors200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModGlossaryGetAuthors200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModGlossaryGetAuthors200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModGlossaryGetAuthors200Response(val *ModGlossaryGetAuthors200Response) *NullableModGlossaryGetAuthors200Response {
	return &NullableModGlossaryGetAuthors200Response{value: val, isSet: true}
}

func (v NullableModGlossaryGetAuthors200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModGlossaryGetAuthors200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

