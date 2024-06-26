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

// checks if the CoreCompetencySearchCompetenciesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCompetencySearchCompetenciesRequest{}

// CoreCompetencySearchCompetenciesRequest struct for CoreCompetencySearchCompetenciesRequest
type CoreCompetencySearchCompetenciesRequest struct {
	// Competency framework id
	Competencyframeworkid int32 `json:"competencyframeworkid"`
	// Text to search for
	Searchtext string `json:"searchtext"`
}

type _CoreCompetencySearchCompetenciesRequest CoreCompetencySearchCompetenciesRequest

// NewCoreCompetencySearchCompetenciesRequest instantiates a new CoreCompetencySearchCompetenciesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCompetencySearchCompetenciesRequest(competencyframeworkid int32, searchtext string) *CoreCompetencySearchCompetenciesRequest {
	this := CoreCompetencySearchCompetenciesRequest{}
	this.Competencyframeworkid = competencyframeworkid
	this.Searchtext = searchtext
	return &this
}

// NewCoreCompetencySearchCompetenciesRequestWithDefaults instantiates a new CoreCompetencySearchCompetenciesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCompetencySearchCompetenciesRequestWithDefaults() *CoreCompetencySearchCompetenciesRequest {
	this := CoreCompetencySearchCompetenciesRequest{}
	var competencyframeworkid int32 = null
	this.Competencyframeworkid = competencyframeworkid
	var searchtext string = "null"
	this.Searchtext = searchtext
	return &this
}

// GetCompetencyframeworkid returns the Competencyframeworkid field value
func (o *CoreCompetencySearchCompetenciesRequest) GetCompetencyframeworkid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Competencyframeworkid
}

// GetCompetencyframeworkidOk returns a tuple with the Competencyframeworkid field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencySearchCompetenciesRequest) GetCompetencyframeworkidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Competencyframeworkid, true
}

// SetCompetencyframeworkid sets field value
func (o *CoreCompetencySearchCompetenciesRequest) SetCompetencyframeworkid(v int32) {
	o.Competencyframeworkid = v
}

// GetSearchtext returns the Searchtext field value
func (o *CoreCompetencySearchCompetenciesRequest) GetSearchtext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Searchtext
}

// GetSearchtextOk returns a tuple with the Searchtext field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencySearchCompetenciesRequest) GetSearchtextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Searchtext, true
}

// SetSearchtext sets field value
func (o *CoreCompetencySearchCompetenciesRequest) SetSearchtext(v string) {
	o.Searchtext = v
}

func (o CoreCompetencySearchCompetenciesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCompetencySearchCompetenciesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["competencyframeworkid"] = o.Competencyframeworkid
	toSerialize["searchtext"] = o.Searchtext
	return toSerialize, nil
}

func (o *CoreCompetencySearchCompetenciesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"competencyframeworkid",
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

	varCoreCompetencySearchCompetenciesRequest := _CoreCompetencySearchCompetenciesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCompetencySearchCompetenciesRequest)

	if err != nil {
		return err
	}

	*o = CoreCompetencySearchCompetenciesRequest(varCoreCompetencySearchCompetenciesRequest)

	return err
}

type NullableCoreCompetencySearchCompetenciesRequest struct {
	value *CoreCompetencySearchCompetenciesRequest
	isSet bool
}

func (v NullableCoreCompetencySearchCompetenciesRequest) Get() *CoreCompetencySearchCompetenciesRequest {
	return v.value
}

func (v *NullableCoreCompetencySearchCompetenciesRequest) Set(val *CoreCompetencySearchCompetenciesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCompetencySearchCompetenciesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCompetencySearchCompetenciesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCompetencySearchCompetenciesRequest(val *CoreCompetencySearchCompetenciesRequest) *NullableCoreCompetencySearchCompetenciesRequest {
	return &NullableCoreCompetencySearchCompetenciesRequest{value: val, isSet: true}
}

func (v NullableCoreCompetencySearchCompetenciesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCompetencySearchCompetenciesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


