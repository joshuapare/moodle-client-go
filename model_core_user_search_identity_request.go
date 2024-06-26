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

// checks if the CoreUserSearchIdentityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreUserSearchIdentityRequest{}

// CoreUserSearchIdentityRequest struct for CoreUserSearchIdentityRequest
type CoreUserSearchIdentityRequest struct {
	// The search query
	Query string `json:"query"`
}

type _CoreUserSearchIdentityRequest CoreUserSearchIdentityRequest

// NewCoreUserSearchIdentityRequest instantiates a new CoreUserSearchIdentityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUserSearchIdentityRequest(query string) *CoreUserSearchIdentityRequest {
	this := CoreUserSearchIdentityRequest{}
	this.Query = query
	return &this
}

// NewCoreUserSearchIdentityRequestWithDefaults instantiates a new CoreUserSearchIdentityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUserSearchIdentityRequestWithDefaults() *CoreUserSearchIdentityRequest {
	this := CoreUserSearchIdentityRequest{}
	var query string = "null"
	this.Query = query
	return &this
}

// GetQuery returns the Query field value
func (o *CoreUserSearchIdentityRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CoreUserSearchIdentityRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *CoreUserSearchIdentityRequest) SetQuery(v string) {
	o.Query = v
}

func (o CoreUserSearchIdentityRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreUserSearchIdentityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	return toSerialize, nil
}

func (o *CoreUserSearchIdentityRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"query",
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

	varCoreUserSearchIdentityRequest := _CoreUserSearchIdentityRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreUserSearchIdentityRequest)

	if err != nil {
		return err
	}

	*o = CoreUserSearchIdentityRequest(varCoreUserSearchIdentityRequest)

	return err
}

type NullableCoreUserSearchIdentityRequest struct {
	value *CoreUserSearchIdentityRequest
	isSet bool
}

func (v NullableCoreUserSearchIdentityRequest) Get() *CoreUserSearchIdentityRequest {
	return v.value
}

func (v *NullableCoreUserSearchIdentityRequest) Set(val *CoreUserSearchIdentityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUserSearchIdentityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUserSearchIdentityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUserSearchIdentityRequest(val *CoreUserSearchIdentityRequest) *NullableCoreUserSearchIdentityRequest {
	return &NullableCoreUserSearchIdentityRequest{value: val, isSet: true}
}

func (v NullableCoreUserSearchIdentityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUserSearchIdentityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


