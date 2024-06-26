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

// checks if the TinyAutosaveResumeSession200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TinyAutosaveResumeSession200Response{}

// TinyAutosaveResumeSession200Response struct for TinyAutosaveResumeSession200Response
type TinyAutosaveResumeSession200Response struct {
	// The draft text
	Drafttext string `json:"drafttext"`
}

type _TinyAutosaveResumeSession200Response TinyAutosaveResumeSession200Response

// NewTinyAutosaveResumeSession200Response instantiates a new TinyAutosaveResumeSession200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTinyAutosaveResumeSession200Response(drafttext string) *TinyAutosaveResumeSession200Response {
	this := TinyAutosaveResumeSession200Response{}
	this.Drafttext = drafttext
	return &this
}

// NewTinyAutosaveResumeSession200ResponseWithDefaults instantiates a new TinyAutosaveResumeSession200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTinyAutosaveResumeSession200ResponseWithDefaults() *TinyAutosaveResumeSession200Response {
	this := TinyAutosaveResumeSession200Response{}
	var drafttext string = "null"
	this.Drafttext = drafttext
	return &this
}

// GetDrafttext returns the Drafttext field value
func (o *TinyAutosaveResumeSession200Response) GetDrafttext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Drafttext
}

// GetDrafttextOk returns a tuple with the Drafttext field value
// and a boolean to check if the value has been set.
func (o *TinyAutosaveResumeSession200Response) GetDrafttextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Drafttext, true
}

// SetDrafttext sets field value
func (o *TinyAutosaveResumeSession200Response) SetDrafttext(v string) {
	o.Drafttext = v
}

func (o TinyAutosaveResumeSession200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TinyAutosaveResumeSession200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["drafttext"] = o.Drafttext
	return toSerialize, nil
}

func (o *TinyAutosaveResumeSession200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"drafttext",
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

	varTinyAutosaveResumeSession200Response := _TinyAutosaveResumeSession200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTinyAutosaveResumeSession200Response)

	if err != nil {
		return err
	}

	*o = TinyAutosaveResumeSession200Response(varTinyAutosaveResumeSession200Response)

	return err
}

type NullableTinyAutosaveResumeSession200Response struct {
	value *TinyAutosaveResumeSession200Response
	isSet bool
}

func (v NullableTinyAutosaveResumeSession200Response) Get() *TinyAutosaveResumeSession200Response {
	return v.value
}

func (v *NullableTinyAutosaveResumeSession200Response) Set(val *TinyAutosaveResumeSession200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTinyAutosaveResumeSession200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTinyAutosaveResumeSession200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTinyAutosaveResumeSession200Response(val *TinyAutosaveResumeSession200Response) *NullableTinyAutosaveResumeSession200Response {
	return &NullableTinyAutosaveResumeSession200Response{value: val, isSet: true}
}

func (v NullableTinyAutosaveResumeSession200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTinyAutosaveResumeSession200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


