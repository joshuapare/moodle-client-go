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

// checks if the ReportCompetencyDataForReportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportCompetencyDataForReportRequest{}

// ReportCompetencyDataForReportRequest struct for ReportCompetencyDataForReportRequest
type ReportCompetencyDataForReportRequest struct {
	// The course id
	Courseid int32 `json:"courseid"`
	// The module id
	Moduleid int32 `json:"moduleid"`
	// The user id
	Userid int32 `json:"userid"`
}

type _ReportCompetencyDataForReportRequest ReportCompetencyDataForReportRequest

// NewReportCompetencyDataForReportRequest instantiates a new ReportCompetencyDataForReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportCompetencyDataForReportRequest(courseid int32, moduleid int32, userid int32) *ReportCompetencyDataForReportRequest {
	this := ReportCompetencyDataForReportRequest{}
	this.Courseid = courseid
	this.Moduleid = moduleid
	this.Userid = userid
	return &this
}

// NewReportCompetencyDataForReportRequestWithDefaults instantiates a new ReportCompetencyDataForReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportCompetencyDataForReportRequestWithDefaults() *ReportCompetencyDataForReportRequest {
	this := ReportCompetencyDataForReportRequest{}
	var moduleid int32 = null
	this.Moduleid = moduleid
	return &this
}

// GetCourseid returns the Courseid field value
func (o *ReportCompetencyDataForReportRequest) GetCourseid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value
// and a boolean to check if the value has been set.
func (o *ReportCompetencyDataForReportRequest) GetCourseidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Courseid, true
}

// SetCourseid sets field value
func (o *ReportCompetencyDataForReportRequest) SetCourseid(v int32) {
	o.Courseid = v
}

// GetModuleid returns the Moduleid field value
func (o *ReportCompetencyDataForReportRequest) GetModuleid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Moduleid
}

// GetModuleidOk returns a tuple with the Moduleid field value
// and a boolean to check if the value has been set.
func (o *ReportCompetencyDataForReportRequest) GetModuleidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Moduleid, true
}

// SetModuleid sets field value
func (o *ReportCompetencyDataForReportRequest) SetModuleid(v int32) {
	o.Moduleid = v
}

// GetUserid returns the Userid field value
func (o *ReportCompetencyDataForReportRequest) GetUserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Userid
}

// GetUseridOk returns a tuple with the Userid field value
// and a boolean to check if the value has been set.
func (o *ReportCompetencyDataForReportRequest) GetUseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userid, true
}

// SetUserid sets field value
func (o *ReportCompetencyDataForReportRequest) SetUserid(v int32) {
	o.Userid = v
}

func (o ReportCompetencyDataForReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportCompetencyDataForReportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["courseid"] = o.Courseid
	toSerialize["moduleid"] = o.Moduleid
	toSerialize["userid"] = o.Userid
	return toSerialize, nil
}

func (o *ReportCompetencyDataForReportRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"courseid",
		"moduleid",
		"userid",
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

	varReportCompetencyDataForReportRequest := _ReportCompetencyDataForReportRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportCompetencyDataForReportRequest)

	if err != nil {
		return err
	}

	*o = ReportCompetencyDataForReportRequest(varReportCompetencyDataForReportRequest)

	return err
}

type NullableReportCompetencyDataForReportRequest struct {
	value *ReportCompetencyDataForReportRequest
	isSet bool
}

func (v NullableReportCompetencyDataForReportRequest) Get() *ReportCompetencyDataForReportRequest {
	return v.value
}

func (v *NullableReportCompetencyDataForReportRequest) Set(val *ReportCompetencyDataForReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReportCompetencyDataForReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReportCompetencyDataForReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportCompetencyDataForReportRequest(val *ReportCompetencyDataForReportRequest) *NullableReportCompetencyDataForReportRequest {
	return &NullableReportCompetencyDataForReportRequest{value: val, isSet: true}
}

func (v NullableReportCompetencyDataForReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportCompetencyDataForReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


