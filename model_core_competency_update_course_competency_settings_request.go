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

// checks if the CoreCompetencyUpdateCourseCompetencySettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCompetencyUpdateCourseCompetencySettingsRequest{}

// CoreCompetencyUpdateCourseCompetencySettingsRequest struct for CoreCompetencyUpdateCourseCompetencySettingsRequest
type CoreCompetencyUpdateCourseCompetencySettingsRequest struct {
	// Course id for the course to update
	Courseid int32 `json:"courseid"`
	Settings CoreCompetencyUpdateCourseCompetencySettingsRequestSettings `json:"settings"`
}

type _CoreCompetencyUpdateCourseCompetencySettingsRequest CoreCompetencyUpdateCourseCompetencySettingsRequest

// NewCoreCompetencyUpdateCourseCompetencySettingsRequest instantiates a new CoreCompetencyUpdateCourseCompetencySettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCompetencyUpdateCourseCompetencySettingsRequest(courseid int32, settings CoreCompetencyUpdateCourseCompetencySettingsRequestSettings) *CoreCompetencyUpdateCourseCompetencySettingsRequest {
	this := CoreCompetencyUpdateCourseCompetencySettingsRequest{}
	this.Courseid = courseid
	this.Settings = settings
	return &this
}

// NewCoreCompetencyUpdateCourseCompetencySettingsRequestWithDefaults instantiates a new CoreCompetencyUpdateCourseCompetencySettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCompetencyUpdateCourseCompetencySettingsRequestWithDefaults() *CoreCompetencyUpdateCourseCompetencySettingsRequest {
	this := CoreCompetencyUpdateCourseCompetencySettingsRequest{}
	var courseid int32 = null
	this.Courseid = courseid
	return &this
}

// GetCourseid returns the Courseid field value
func (o *CoreCompetencyUpdateCourseCompetencySettingsRequest) GetCourseid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyUpdateCourseCompetencySettingsRequest) GetCourseidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Courseid, true
}

// SetCourseid sets field value
func (o *CoreCompetencyUpdateCourseCompetencySettingsRequest) SetCourseid(v int32) {
	o.Courseid = v
}

// GetSettings returns the Settings field value
func (o *CoreCompetencyUpdateCourseCompetencySettingsRequest) GetSettings() CoreCompetencyUpdateCourseCompetencySettingsRequestSettings {
	if o == nil {
		var ret CoreCompetencyUpdateCourseCompetencySettingsRequestSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyUpdateCourseCompetencySettingsRequest) GetSettingsOk() (*CoreCompetencyUpdateCourseCompetencySettingsRequestSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *CoreCompetencyUpdateCourseCompetencySettingsRequest) SetSettings(v CoreCompetencyUpdateCourseCompetencySettingsRequestSettings) {
	o.Settings = v
}

func (o CoreCompetencyUpdateCourseCompetencySettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCompetencyUpdateCourseCompetencySettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["courseid"] = o.Courseid
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

func (o *CoreCompetencyUpdateCourseCompetencySettingsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"courseid",
		"settings",
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

	varCoreCompetencyUpdateCourseCompetencySettingsRequest := _CoreCompetencyUpdateCourseCompetencySettingsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCompetencyUpdateCourseCompetencySettingsRequest)

	if err != nil {
		return err
	}

	*o = CoreCompetencyUpdateCourseCompetencySettingsRequest(varCoreCompetencyUpdateCourseCompetencySettingsRequest)

	return err
}

type NullableCoreCompetencyUpdateCourseCompetencySettingsRequest struct {
	value *CoreCompetencyUpdateCourseCompetencySettingsRequest
	isSet bool
}

func (v NullableCoreCompetencyUpdateCourseCompetencySettingsRequest) Get() *CoreCompetencyUpdateCourseCompetencySettingsRequest {
	return v.value
}

func (v *NullableCoreCompetencyUpdateCourseCompetencySettingsRequest) Set(val *CoreCompetencyUpdateCourseCompetencySettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCompetencyUpdateCourseCompetencySettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCompetencyUpdateCourseCompetencySettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCompetencyUpdateCourseCompetencySettingsRequest(val *CoreCompetencyUpdateCourseCompetencySettingsRequest) *NullableCoreCompetencyUpdateCourseCompetencySettingsRequest {
	return &NullableCoreCompetencyUpdateCourseCompetencySettingsRequest{value: val, isSet: true}
}

func (v NullableCoreCompetencyUpdateCourseCompetencySettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCompetencyUpdateCourseCompetencySettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


