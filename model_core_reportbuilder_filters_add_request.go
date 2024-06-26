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

// checks if the CoreReportbuilderFiltersAddRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreReportbuilderFiltersAddRequest{}

// CoreReportbuilderFiltersAddRequest struct for CoreReportbuilderFiltersAddRequest
type CoreReportbuilderFiltersAddRequest struct {
	// Report ID
	Reportid int32 `json:"reportid"`
	// Unique identifier of the filter
	Uniqueidentifier string `json:"uniqueidentifier"`
}

type _CoreReportbuilderFiltersAddRequest CoreReportbuilderFiltersAddRequest

// NewCoreReportbuilderFiltersAddRequest instantiates a new CoreReportbuilderFiltersAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreReportbuilderFiltersAddRequest(reportid int32, uniqueidentifier string) *CoreReportbuilderFiltersAddRequest {
	this := CoreReportbuilderFiltersAddRequest{}
	this.Reportid = reportid
	this.Uniqueidentifier = uniqueidentifier
	return &this
}

// NewCoreReportbuilderFiltersAddRequestWithDefaults instantiates a new CoreReportbuilderFiltersAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreReportbuilderFiltersAddRequestWithDefaults() *CoreReportbuilderFiltersAddRequest {
	this := CoreReportbuilderFiltersAddRequest{}
	var uniqueidentifier string = "null"
	this.Uniqueidentifier = uniqueidentifier
	return &this
}

// GetReportid returns the Reportid field value
func (o *CoreReportbuilderFiltersAddRequest) GetReportid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Reportid
}

// GetReportidOk returns a tuple with the Reportid field value
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderFiltersAddRequest) GetReportidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reportid, true
}

// SetReportid sets field value
func (o *CoreReportbuilderFiltersAddRequest) SetReportid(v int32) {
	o.Reportid = v
}

// GetUniqueidentifier returns the Uniqueidentifier field value
func (o *CoreReportbuilderFiltersAddRequest) GetUniqueidentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uniqueidentifier
}

// GetUniqueidentifierOk returns a tuple with the Uniqueidentifier field value
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderFiltersAddRequest) GetUniqueidentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uniqueidentifier, true
}

// SetUniqueidentifier sets field value
func (o *CoreReportbuilderFiltersAddRequest) SetUniqueidentifier(v string) {
	o.Uniqueidentifier = v
}

func (o CoreReportbuilderFiltersAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreReportbuilderFiltersAddRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reportid"] = o.Reportid
	toSerialize["uniqueidentifier"] = o.Uniqueidentifier
	return toSerialize, nil
}

func (o *CoreReportbuilderFiltersAddRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reportid",
		"uniqueidentifier",
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

	varCoreReportbuilderFiltersAddRequest := _CoreReportbuilderFiltersAddRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreReportbuilderFiltersAddRequest)

	if err != nil {
		return err
	}

	*o = CoreReportbuilderFiltersAddRequest(varCoreReportbuilderFiltersAddRequest)

	return err
}

type NullableCoreReportbuilderFiltersAddRequest struct {
	value *CoreReportbuilderFiltersAddRequest
	isSet bool
}

func (v NullableCoreReportbuilderFiltersAddRequest) Get() *CoreReportbuilderFiltersAddRequest {
	return v.value
}

func (v *NullableCoreReportbuilderFiltersAddRequest) Set(val *CoreReportbuilderFiltersAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreReportbuilderFiltersAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreReportbuilderFiltersAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreReportbuilderFiltersAddRequest(val *CoreReportbuilderFiltersAddRequest) *NullableCoreReportbuilderFiltersAddRequest {
	return &NullableCoreReportbuilderFiltersAddRequest{value: val, isSet: true}
}

func (v NullableCoreReportbuilderFiltersAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreReportbuilderFiltersAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


