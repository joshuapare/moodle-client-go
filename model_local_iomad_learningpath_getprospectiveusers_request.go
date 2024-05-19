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

// checks if the LocalIomadLearningpathGetprospectiveusersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalIomadLearningpathGetprospectiveusersRequest{}

// LocalIomadLearningpathGetprospectiveusersRequest struct for LocalIomadLearningpathGetprospectiveusersRequest
type LocalIomadLearningpathGetprospectiveusersRequest struct {
	// ID of Iomad Company
	Companyid int32 `json:"companyid"`
	// Filter user list returned
	Filter *string `json:"filter,omitempty"`
	// ID learning path
	Pathid int32 `json:"pathid"`
	// Filter by user profilefield
	Profilefieldid *int32 `json:"profilefieldid,omitempty"`
}

type _LocalIomadLearningpathGetprospectiveusersRequest LocalIomadLearningpathGetprospectiveusersRequest

// NewLocalIomadLearningpathGetprospectiveusersRequest instantiates a new LocalIomadLearningpathGetprospectiveusersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalIomadLearningpathGetprospectiveusersRequest(companyid int32, pathid int32) *LocalIomadLearningpathGetprospectiveusersRequest {
	this := LocalIomadLearningpathGetprospectiveusersRequest{}
	this.Companyid = companyid
	var filter string = ""
	this.Filter = &filter
	this.Pathid = pathid
	var profilefieldid int32 = 0
	this.Profilefieldid = &profilefieldid
	return &this
}

// NewLocalIomadLearningpathGetprospectiveusersRequestWithDefaults instantiates a new LocalIomadLearningpathGetprospectiveusersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalIomadLearningpathGetprospectiveusersRequestWithDefaults() *LocalIomadLearningpathGetprospectiveusersRequest {
	this := LocalIomadLearningpathGetprospectiveusersRequest{}
	var companyid int32 = null
	this.Companyid = companyid
	var filter string = ""
	this.Filter = &filter
	var pathid int32 = null
	this.Pathid = pathid
	var profilefieldid int32 = 0
	this.Profilefieldid = &profilefieldid
	return &this
}

// GetCompanyid returns the Companyid field value
func (o *LocalIomadLearningpathGetprospectiveusersRequest) GetCompanyid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Companyid
}

// GetCompanyidOk returns a tuple with the Companyid field value
// and a boolean to check if the value has been set.
func (o *LocalIomadLearningpathGetprospectiveusersRequest) GetCompanyidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Companyid, true
}

// SetCompanyid sets field value
func (o *LocalIomadLearningpathGetprospectiveusersRequest) SetCompanyid(v int32) {
	o.Companyid = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *LocalIomadLearningpathGetprospectiveusersRequest) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalIomadLearningpathGetprospectiveusersRequest) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *LocalIomadLearningpathGetprospectiveusersRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *LocalIomadLearningpathGetprospectiveusersRequest) SetFilter(v string) {
	o.Filter = &v
}

// GetPathid returns the Pathid field value
func (o *LocalIomadLearningpathGetprospectiveusersRequest) GetPathid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pathid
}

// GetPathidOk returns a tuple with the Pathid field value
// and a boolean to check if the value has been set.
func (o *LocalIomadLearningpathGetprospectiveusersRequest) GetPathidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pathid, true
}

// SetPathid sets field value
func (o *LocalIomadLearningpathGetprospectiveusersRequest) SetPathid(v int32) {
	o.Pathid = v
}

// GetProfilefieldid returns the Profilefieldid field value if set, zero value otherwise.
func (o *LocalIomadLearningpathGetprospectiveusersRequest) GetProfilefieldid() int32 {
	if o == nil || IsNil(o.Profilefieldid) {
		var ret int32
		return ret
	}
	return *o.Profilefieldid
}

// GetProfilefieldidOk returns a tuple with the Profilefieldid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalIomadLearningpathGetprospectiveusersRequest) GetProfilefieldidOk() (*int32, bool) {
	if o == nil || IsNil(o.Profilefieldid) {
		return nil, false
	}
	return o.Profilefieldid, true
}

// HasProfilefieldid returns a boolean if a field has been set.
func (o *LocalIomadLearningpathGetprospectiveusersRequest) HasProfilefieldid() bool {
	if o != nil && !IsNil(o.Profilefieldid) {
		return true
	}

	return false
}

// SetProfilefieldid gets a reference to the given int32 and assigns it to the Profilefieldid field.
func (o *LocalIomadLearningpathGetprospectiveusersRequest) SetProfilefieldid(v int32) {
	o.Profilefieldid = &v
}

func (o LocalIomadLearningpathGetprospectiveusersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalIomadLearningpathGetprospectiveusersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["companyid"] = o.Companyid
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	toSerialize["pathid"] = o.Pathid
	if !IsNil(o.Profilefieldid) {
		toSerialize["profilefieldid"] = o.Profilefieldid
	}
	return toSerialize, nil
}

func (o *LocalIomadLearningpathGetprospectiveusersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"companyid",
		"pathid",
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

	varLocalIomadLearningpathGetprospectiveusersRequest := _LocalIomadLearningpathGetprospectiveusersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLocalIomadLearningpathGetprospectiveusersRequest)

	if err != nil {
		return err
	}

	*o = LocalIomadLearningpathGetprospectiveusersRequest(varLocalIomadLearningpathGetprospectiveusersRequest)

	return err
}

type NullableLocalIomadLearningpathGetprospectiveusersRequest struct {
	value *LocalIomadLearningpathGetprospectiveusersRequest
	isSet bool
}

func (v NullableLocalIomadLearningpathGetprospectiveusersRequest) Get() *LocalIomadLearningpathGetprospectiveusersRequest {
	return v.value
}

func (v *NullableLocalIomadLearningpathGetprospectiveusersRequest) Set(val *LocalIomadLearningpathGetprospectiveusersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalIomadLearningpathGetprospectiveusersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalIomadLearningpathGetprospectiveusersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalIomadLearningpathGetprospectiveusersRequest(val *LocalIomadLearningpathGetprospectiveusersRequest) *NullableLocalIomadLearningpathGetprospectiveusersRequest {
	return &NullableLocalIomadLearningpathGetprospectiveusersRequest{value: val, isSet: true}
}

func (v NullableLocalIomadLearningpathGetprospectiveusersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalIomadLearningpathGetprospectiveusersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


