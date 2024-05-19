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

// checks if the ModGlossaryGetCategories200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModGlossaryGetCategories200Response{}

// ModGlossaryGetCategories200Response struct for ModGlossaryGetCategories200Response
type ModGlossaryGetCategories200Response struct {
	Categories []ModGlossaryGetCategories200ResponseCategoriesInner `json:"categories"`
	// The total number of records.
	Count int32 `json:"count"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModGlossaryGetCategories200Response ModGlossaryGetCategories200Response

// NewModGlossaryGetCategories200Response instantiates a new ModGlossaryGetCategories200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModGlossaryGetCategories200Response(categories []ModGlossaryGetCategories200ResponseCategoriesInner, count int32) *ModGlossaryGetCategories200Response {
	this := ModGlossaryGetCategories200Response{}
	this.Categories = categories
	this.Count = count
	return &this
}

// NewModGlossaryGetCategories200ResponseWithDefaults instantiates a new ModGlossaryGetCategories200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModGlossaryGetCategories200ResponseWithDefaults() *ModGlossaryGetCategories200Response {
	this := ModGlossaryGetCategories200Response{}
	return &this
}

// GetCategories returns the Categories field value
func (o *ModGlossaryGetCategories200Response) GetCategories() []ModGlossaryGetCategories200ResponseCategoriesInner {
	if o == nil {
		var ret []ModGlossaryGetCategories200ResponseCategoriesInner
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetCategories200Response) GetCategoriesOk() ([]ModGlossaryGetCategories200ResponseCategoriesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Categories, true
}

// SetCategories sets field value
func (o *ModGlossaryGetCategories200Response) SetCategories(v []ModGlossaryGetCategories200ResponseCategoriesInner) {
	o.Categories = v
}

// GetCount returns the Count field value
func (o *ModGlossaryGetCategories200Response) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetCategories200Response) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ModGlossaryGetCategories200Response) SetCount(v int32) {
	o.Count = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModGlossaryGetCategories200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetCategories200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModGlossaryGetCategories200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModGlossaryGetCategories200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModGlossaryGetCategories200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModGlossaryGetCategories200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["categories"] = o.Categories
	toSerialize["count"] = o.Count
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModGlossaryGetCategories200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"categories",
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

	varModGlossaryGetCategories200Response := _ModGlossaryGetCategories200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModGlossaryGetCategories200Response)

	if err != nil {
		return err
	}

	*o = ModGlossaryGetCategories200Response(varModGlossaryGetCategories200Response)

	return err
}

type NullableModGlossaryGetCategories200Response struct {
	value *ModGlossaryGetCategories200Response
	isSet bool
}

func (v NullableModGlossaryGetCategories200Response) Get() *ModGlossaryGetCategories200Response {
	return v.value
}

func (v *NullableModGlossaryGetCategories200Response) Set(val *ModGlossaryGetCategories200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModGlossaryGetCategories200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModGlossaryGetCategories200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModGlossaryGetCategories200Response(val *ModGlossaryGetCategories200Response) *NullableModGlossaryGetCategories200Response {
	return &NullableModGlossaryGetCategories200Response{value: val, isSet: true}
}

func (v NullableModGlossaryGetCategories200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModGlossaryGetCategories200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

