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

// checks if the GradingformGuideGraderGradingpanelFetch200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GradingformGuideGraderGradingpanelFetch200Response{}

// GradingformGuideGraderGradingpanelFetch200Response struct for GradingformGuideGraderGradingpanelFetch200Response
type GradingformGuideGraderGradingpanelFetch200Response struct {
	Grade GradingformGuideGraderGradingpanelFetch200ResponseGrade `json:"grade"`
	// Does the user have a grade?
	Hasgrade bool `json:"hasgrade"`
	// The template to use when rendering this data
	Templatename string `json:"templatename"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _GradingformGuideGraderGradingpanelFetch200Response GradingformGuideGraderGradingpanelFetch200Response

// NewGradingformGuideGraderGradingpanelFetch200Response instantiates a new GradingformGuideGraderGradingpanelFetch200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGradingformGuideGraderGradingpanelFetch200Response(grade GradingformGuideGraderGradingpanelFetch200ResponseGrade, hasgrade bool, templatename string) *GradingformGuideGraderGradingpanelFetch200Response {
	this := GradingformGuideGraderGradingpanelFetch200Response{}
	this.Grade = grade
	this.Hasgrade = hasgrade
	this.Templatename = templatename
	return &this
}

// NewGradingformGuideGraderGradingpanelFetch200ResponseWithDefaults instantiates a new GradingformGuideGraderGradingpanelFetch200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGradingformGuideGraderGradingpanelFetch200ResponseWithDefaults() *GradingformGuideGraderGradingpanelFetch200Response {
	this := GradingformGuideGraderGradingpanelFetch200Response{}
	return &this
}

// GetGrade returns the Grade field value
func (o *GradingformGuideGraderGradingpanelFetch200Response) GetGrade() GradingformGuideGraderGradingpanelFetch200ResponseGrade {
	if o == nil {
		var ret GradingformGuideGraderGradingpanelFetch200ResponseGrade
		return ret
	}

	return o.Grade
}

// GetGradeOk returns a tuple with the Grade field value
// and a boolean to check if the value has been set.
func (o *GradingformGuideGraderGradingpanelFetch200Response) GetGradeOk() (*GradingformGuideGraderGradingpanelFetch200ResponseGrade, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grade, true
}

// SetGrade sets field value
func (o *GradingformGuideGraderGradingpanelFetch200Response) SetGrade(v GradingformGuideGraderGradingpanelFetch200ResponseGrade) {
	o.Grade = v
}

// GetHasgrade returns the Hasgrade field value
func (o *GradingformGuideGraderGradingpanelFetch200Response) GetHasgrade() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hasgrade
}

// GetHasgradeOk returns a tuple with the Hasgrade field value
// and a boolean to check if the value has been set.
func (o *GradingformGuideGraderGradingpanelFetch200Response) GetHasgradeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hasgrade, true
}

// SetHasgrade sets field value
func (o *GradingformGuideGraderGradingpanelFetch200Response) SetHasgrade(v bool) {
	o.Hasgrade = v
}

// GetTemplatename returns the Templatename field value
func (o *GradingformGuideGraderGradingpanelFetch200Response) GetTemplatename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Templatename
}

// GetTemplatenameOk returns a tuple with the Templatename field value
// and a boolean to check if the value has been set.
func (o *GradingformGuideGraderGradingpanelFetch200Response) GetTemplatenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Templatename, true
}

// SetTemplatename sets field value
func (o *GradingformGuideGraderGradingpanelFetch200Response) SetTemplatename(v string) {
	o.Templatename = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *GradingformGuideGraderGradingpanelFetch200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GradingformGuideGraderGradingpanelFetch200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *GradingformGuideGraderGradingpanelFetch200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *GradingformGuideGraderGradingpanelFetch200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o GradingformGuideGraderGradingpanelFetch200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GradingformGuideGraderGradingpanelFetch200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grade"] = o.Grade
	toSerialize["hasgrade"] = o.Hasgrade
	toSerialize["templatename"] = o.Templatename
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *GradingformGuideGraderGradingpanelFetch200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"grade",
		"hasgrade",
		"templatename",
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

	varGradingformGuideGraderGradingpanelFetch200Response := _GradingformGuideGraderGradingpanelFetch200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGradingformGuideGraderGradingpanelFetch200Response)

	if err != nil {
		return err
	}

	*o = GradingformGuideGraderGradingpanelFetch200Response(varGradingformGuideGraderGradingpanelFetch200Response)

	return err
}

type NullableGradingformGuideGraderGradingpanelFetch200Response struct {
	value *GradingformGuideGraderGradingpanelFetch200Response
	isSet bool
}

func (v NullableGradingformGuideGraderGradingpanelFetch200Response) Get() *GradingformGuideGraderGradingpanelFetch200Response {
	return v.value
}

func (v *NullableGradingformGuideGraderGradingpanelFetch200Response) Set(val *GradingformGuideGraderGradingpanelFetch200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGradingformGuideGraderGradingpanelFetch200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGradingformGuideGraderGradingpanelFetch200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGradingformGuideGraderGradingpanelFetch200Response(val *GradingformGuideGraderGradingpanelFetch200Response) *NullableGradingformGuideGraderGradingpanelFetch200Response {
	return &NullableGradingformGuideGraderGradingpanelFetch200Response{value: val, isSet: true}
}

func (v NullableGradingformGuideGraderGradingpanelFetch200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGradingformGuideGraderGradingpanelFetch200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

