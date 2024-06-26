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

// checks if the CoreMessageSearchContactsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageSearchContactsRequest{}

// CoreMessageSearchContactsRequest struct for CoreMessageSearchContactsRequest
type CoreMessageSearchContactsRequest struct {
	// Limit search to the user's courses
	Onlymycourses *bool `json:"onlymycourses,omitempty"`
	// String the user's fullname has to match to be found
	Searchtext string `json:"searchtext"`
}

type _CoreMessageSearchContactsRequest CoreMessageSearchContactsRequest

// NewCoreMessageSearchContactsRequest instantiates a new CoreMessageSearchContactsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageSearchContactsRequest(searchtext string) *CoreMessageSearchContactsRequest {
	this := CoreMessageSearchContactsRequest{}
	var onlymycourses bool = false
	this.Onlymycourses = &onlymycourses
	this.Searchtext = searchtext
	return &this
}

// NewCoreMessageSearchContactsRequestWithDefaults instantiates a new CoreMessageSearchContactsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageSearchContactsRequestWithDefaults() *CoreMessageSearchContactsRequest {
	this := CoreMessageSearchContactsRequest{}
	var onlymycourses bool = false
	this.Onlymycourses = &onlymycourses
	var searchtext string = "null"
	this.Searchtext = searchtext
	return &this
}

// GetOnlymycourses returns the Onlymycourses field value if set, zero value otherwise.
func (o *CoreMessageSearchContactsRequest) GetOnlymycourses() bool {
	if o == nil || IsNil(o.Onlymycourses) {
		var ret bool
		return ret
	}
	return *o.Onlymycourses
}

// GetOnlymycoursesOk returns a tuple with the Onlymycourses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageSearchContactsRequest) GetOnlymycoursesOk() (*bool, bool) {
	if o == nil || IsNil(o.Onlymycourses) {
		return nil, false
	}
	return o.Onlymycourses, true
}

// HasOnlymycourses returns a boolean if a field has been set.
func (o *CoreMessageSearchContactsRequest) HasOnlymycourses() bool {
	if o != nil && !IsNil(o.Onlymycourses) {
		return true
	}

	return false
}

// SetOnlymycourses gets a reference to the given bool and assigns it to the Onlymycourses field.
func (o *CoreMessageSearchContactsRequest) SetOnlymycourses(v bool) {
	o.Onlymycourses = &v
}

// GetSearchtext returns the Searchtext field value
func (o *CoreMessageSearchContactsRequest) GetSearchtext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Searchtext
}

// GetSearchtextOk returns a tuple with the Searchtext field value
// and a boolean to check if the value has been set.
func (o *CoreMessageSearchContactsRequest) GetSearchtextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Searchtext, true
}

// SetSearchtext sets field value
func (o *CoreMessageSearchContactsRequest) SetSearchtext(v string) {
	o.Searchtext = v
}

func (o CoreMessageSearchContactsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageSearchContactsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Onlymycourses) {
		toSerialize["onlymycourses"] = o.Onlymycourses
	}
	toSerialize["searchtext"] = o.Searchtext
	return toSerialize, nil
}

func (o *CoreMessageSearchContactsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"searchtext",
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

	varCoreMessageSearchContactsRequest := _CoreMessageSearchContactsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageSearchContactsRequest)

	if err != nil {
		return err
	}

	*o = CoreMessageSearchContactsRequest(varCoreMessageSearchContactsRequest)

	return err
}

type NullableCoreMessageSearchContactsRequest struct {
	value *CoreMessageSearchContactsRequest
	isSet bool
}

func (v NullableCoreMessageSearchContactsRequest) Get() *CoreMessageSearchContactsRequest {
	return v.value
}

func (v *NullableCoreMessageSearchContactsRequest) Set(val *CoreMessageSearchContactsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageSearchContactsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageSearchContactsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageSearchContactsRequest(val *CoreMessageSearchContactsRequest) *NullableCoreMessageSearchContactsRequest {
	return &NullableCoreMessageSearchContactsRequest{value: val, isSet: true}
}

func (v NullableCoreMessageSearchContactsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageSearchContactsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


