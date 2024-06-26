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

// checks if the CoreBlogGetEntries200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreBlogGetEntries200Response{}

// CoreBlogGetEntries200Response struct for CoreBlogGetEntries200Response
type CoreBlogGetEntries200Response struct {
	Entries []CoreBlogGetEntries200ResponseEntriesInner `json:"entries"`
	// The total number of entries found.
	Totalentries int32 `json:"totalentries"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _CoreBlogGetEntries200Response CoreBlogGetEntries200Response

// NewCoreBlogGetEntries200Response instantiates a new CoreBlogGetEntries200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreBlogGetEntries200Response(entries []CoreBlogGetEntries200ResponseEntriesInner, totalentries int32) *CoreBlogGetEntries200Response {
	this := CoreBlogGetEntries200Response{}
	this.Entries = entries
	this.Totalentries = totalentries
	return &this
}

// NewCoreBlogGetEntries200ResponseWithDefaults instantiates a new CoreBlogGetEntries200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreBlogGetEntries200ResponseWithDefaults() *CoreBlogGetEntries200Response {
	this := CoreBlogGetEntries200Response{}
	var totalentries int32 = null
	this.Totalentries = totalentries
	return &this
}

// GetEntries returns the Entries field value
func (o *CoreBlogGetEntries200Response) GetEntries() []CoreBlogGetEntries200ResponseEntriesInner {
	if o == nil {
		var ret []CoreBlogGetEntries200ResponseEntriesInner
		return ret
	}

	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value
// and a boolean to check if the value has been set.
func (o *CoreBlogGetEntries200Response) GetEntriesOk() ([]CoreBlogGetEntries200ResponseEntriesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entries, true
}

// SetEntries sets field value
func (o *CoreBlogGetEntries200Response) SetEntries(v []CoreBlogGetEntries200ResponseEntriesInner) {
	o.Entries = v
}

// GetTotalentries returns the Totalentries field value
func (o *CoreBlogGetEntries200Response) GetTotalentries() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Totalentries
}

// GetTotalentriesOk returns a tuple with the Totalentries field value
// and a boolean to check if the value has been set.
func (o *CoreBlogGetEntries200Response) GetTotalentriesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Totalentries, true
}

// SetTotalentries sets field value
func (o *CoreBlogGetEntries200Response) SetTotalentries(v int32) {
	o.Totalentries = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CoreBlogGetEntries200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreBlogGetEntries200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CoreBlogGetEntries200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *CoreBlogGetEntries200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o CoreBlogGetEntries200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreBlogGetEntries200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entries"] = o.Entries
	toSerialize["totalentries"] = o.Totalentries
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *CoreBlogGetEntries200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entries",
		"totalentries",
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

	varCoreBlogGetEntries200Response := _CoreBlogGetEntries200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreBlogGetEntries200Response)

	if err != nil {
		return err
	}

	*o = CoreBlogGetEntries200Response(varCoreBlogGetEntries200Response)

	return err
}

type NullableCoreBlogGetEntries200Response struct {
	value *CoreBlogGetEntries200Response
	isSet bool
}

func (v NullableCoreBlogGetEntries200Response) Get() *CoreBlogGetEntries200Response {
	return v.value
}

func (v *NullableCoreBlogGetEntries200Response) Set(val *CoreBlogGetEntries200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreBlogGetEntries200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreBlogGetEntries200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreBlogGetEntries200Response(val *CoreBlogGetEntries200Response) *NullableCoreBlogGetEntries200Response {
	return &NullableCoreBlogGetEntries200Response{value: val, isSet: true}
}

func (v NullableCoreBlogGetEntries200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreBlogGetEntries200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


