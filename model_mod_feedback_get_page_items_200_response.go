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

// checks if the ModFeedbackGetPageItems200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModFeedbackGetPageItems200Response{}

// ModFeedbackGetPageItems200Response struct for ModFeedbackGetPageItems200Response
type ModFeedbackGetPageItems200Response struct {
	// Whether there are more pages.
	Hasnextpage bool `json:"hasnextpage"`
	// Whether is a previous page.
	Hasprevpage bool `json:"hasprevpage"`
	Items []ModFeedbackGetItems200ResponseItemsInner `json:"items"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModFeedbackGetPageItems200Response ModFeedbackGetPageItems200Response

// NewModFeedbackGetPageItems200Response instantiates a new ModFeedbackGetPageItems200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModFeedbackGetPageItems200Response(hasnextpage bool, hasprevpage bool, items []ModFeedbackGetItems200ResponseItemsInner) *ModFeedbackGetPageItems200Response {
	this := ModFeedbackGetPageItems200Response{}
	this.Hasnextpage = hasnextpage
	this.Hasprevpage = hasprevpage
	this.Items = items
	return &this
}

// NewModFeedbackGetPageItems200ResponseWithDefaults instantiates a new ModFeedbackGetPageItems200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModFeedbackGetPageItems200ResponseWithDefaults() *ModFeedbackGetPageItems200Response {
	this := ModFeedbackGetPageItems200Response{}
	var hasnextpage bool = null
	this.Hasnextpage = hasnextpage
	var hasprevpage bool = null
	this.Hasprevpage = hasprevpage
	return &this
}

// GetHasnextpage returns the Hasnextpage field value
func (o *ModFeedbackGetPageItems200Response) GetHasnextpage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hasnextpage
}

// GetHasnextpageOk returns a tuple with the Hasnextpage field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetPageItems200Response) GetHasnextpageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hasnextpage, true
}

// SetHasnextpage sets field value
func (o *ModFeedbackGetPageItems200Response) SetHasnextpage(v bool) {
	o.Hasnextpage = v
}

// GetHasprevpage returns the Hasprevpage field value
func (o *ModFeedbackGetPageItems200Response) GetHasprevpage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hasprevpage
}

// GetHasprevpageOk returns a tuple with the Hasprevpage field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetPageItems200Response) GetHasprevpageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hasprevpage, true
}

// SetHasprevpage sets field value
func (o *ModFeedbackGetPageItems200Response) SetHasprevpage(v bool) {
	o.Hasprevpage = v
}

// GetItems returns the Items field value
func (o *ModFeedbackGetPageItems200Response) GetItems() []ModFeedbackGetItems200ResponseItemsInner {
	if o == nil {
		var ret []ModFeedbackGetItems200ResponseItemsInner
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetPageItems200Response) GetItemsOk() ([]ModFeedbackGetItems200ResponseItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ModFeedbackGetPageItems200Response) SetItems(v []ModFeedbackGetItems200ResponseItemsInner) {
	o.Items = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModFeedbackGetPageItems200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetPageItems200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModFeedbackGetPageItems200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModFeedbackGetPageItems200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModFeedbackGetPageItems200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModFeedbackGetPageItems200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hasnextpage"] = o.Hasnextpage
	toSerialize["hasprevpage"] = o.Hasprevpage
	toSerialize["items"] = o.Items
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModFeedbackGetPageItems200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hasnextpage",
		"hasprevpage",
		"items",
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

	varModFeedbackGetPageItems200Response := _ModFeedbackGetPageItems200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModFeedbackGetPageItems200Response)

	if err != nil {
		return err
	}

	*o = ModFeedbackGetPageItems200Response(varModFeedbackGetPageItems200Response)

	return err
}

type NullableModFeedbackGetPageItems200Response struct {
	value *ModFeedbackGetPageItems200Response
	isSet bool
}

func (v NullableModFeedbackGetPageItems200Response) Get() *ModFeedbackGetPageItems200Response {
	return v.value
}

func (v *NullableModFeedbackGetPageItems200Response) Set(val *ModFeedbackGetPageItems200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModFeedbackGetPageItems200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModFeedbackGetPageItems200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModFeedbackGetPageItems200Response(val *ModFeedbackGetPageItems200Response) *NullableModFeedbackGetPageItems200Response {
	return &NullableModFeedbackGetPageItems200Response{value: val, isSet: true}
}

func (v NullableModFeedbackGetPageItems200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModFeedbackGetPageItems200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


