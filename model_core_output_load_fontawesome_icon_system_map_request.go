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

// checks if the CoreOutputLoadFontawesomeIconSystemMapRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreOutputLoadFontawesomeIconSystemMapRequest{}

// CoreOutputLoadFontawesomeIconSystemMapRequest struct for CoreOutputLoadFontawesomeIconSystemMapRequest
type CoreOutputLoadFontawesomeIconSystemMapRequest struct {
	// The theme to fetch the map for
	Themename string `json:"themename"`
}

type _CoreOutputLoadFontawesomeIconSystemMapRequest CoreOutputLoadFontawesomeIconSystemMapRequest

// NewCoreOutputLoadFontawesomeIconSystemMapRequest instantiates a new CoreOutputLoadFontawesomeIconSystemMapRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreOutputLoadFontawesomeIconSystemMapRequest(themename string) *CoreOutputLoadFontawesomeIconSystemMapRequest {
	this := CoreOutputLoadFontawesomeIconSystemMapRequest{}
	this.Themename = themename
	return &this
}

// NewCoreOutputLoadFontawesomeIconSystemMapRequestWithDefaults instantiates a new CoreOutputLoadFontawesomeIconSystemMapRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreOutputLoadFontawesomeIconSystemMapRequestWithDefaults() *CoreOutputLoadFontawesomeIconSystemMapRequest {
	this := CoreOutputLoadFontawesomeIconSystemMapRequest{}
	var themename string = "null"
	this.Themename = themename
	return &this
}

// GetThemename returns the Themename field value
func (o *CoreOutputLoadFontawesomeIconSystemMapRequest) GetThemename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Themename
}

// GetThemenameOk returns a tuple with the Themename field value
// and a boolean to check if the value has been set.
func (o *CoreOutputLoadFontawesomeIconSystemMapRequest) GetThemenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Themename, true
}

// SetThemename sets field value
func (o *CoreOutputLoadFontawesomeIconSystemMapRequest) SetThemename(v string) {
	o.Themename = v
}

func (o CoreOutputLoadFontawesomeIconSystemMapRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreOutputLoadFontawesomeIconSystemMapRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["themename"] = o.Themename
	return toSerialize, nil
}

func (o *CoreOutputLoadFontawesomeIconSystemMapRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"themename",
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

	varCoreOutputLoadFontawesomeIconSystemMapRequest := _CoreOutputLoadFontawesomeIconSystemMapRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreOutputLoadFontawesomeIconSystemMapRequest)

	if err != nil {
		return err
	}

	*o = CoreOutputLoadFontawesomeIconSystemMapRequest(varCoreOutputLoadFontawesomeIconSystemMapRequest)

	return err
}

type NullableCoreOutputLoadFontawesomeIconSystemMapRequest struct {
	value *CoreOutputLoadFontawesomeIconSystemMapRequest
	isSet bool
}

func (v NullableCoreOutputLoadFontawesomeIconSystemMapRequest) Get() *CoreOutputLoadFontawesomeIconSystemMapRequest {
	return v.value
}

func (v *NullableCoreOutputLoadFontawesomeIconSystemMapRequest) Set(val *CoreOutputLoadFontawesomeIconSystemMapRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreOutputLoadFontawesomeIconSystemMapRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreOutputLoadFontawesomeIconSystemMapRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreOutputLoadFontawesomeIconSystemMapRequest(val *CoreOutputLoadFontawesomeIconSystemMapRequest) *NullableCoreOutputLoadFontawesomeIconSystemMapRequest {
	return &NullableCoreOutputLoadFontawesomeIconSystemMapRequest{value: val, isSet: true}
}

func (v NullableCoreOutputLoadFontawesomeIconSystemMapRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreOutputLoadFontawesomeIconSystemMapRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


