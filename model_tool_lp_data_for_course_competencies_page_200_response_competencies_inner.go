/*
Moodle Webservice API

Auto-generated OpenAPI spec for Moodle's Webservice API. Requires installation of the Catalyst RESTful API plugin: https://github.com/catalyst/moodle-webservice_restful.

API version: 4.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moodleclient

import (
	"encoding/json"
)

// checks if the ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner{}

// ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner struct for ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner
type ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner struct {
	Competency *CoreCompetencyCreateCompetency200Response `json:"competency,omitempty"`
	Comppath *ToolLpDataForCompetencySummary200ResponseComppath `json:"comppath,omitempty"`
	Coursecompetency *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerCoursecompetency `json:"coursecompetency,omitempty"`
	Coursemodules []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerCoursemodulesInner `json:"coursemodules,omitempty"`
	Plans []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerPlansInner `json:"plans,omitempty"`
	Ruleoutcomeoptions []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerRuleoutcomeoptionsInner `json:"ruleoutcomeoptions,omitempty"`
	Usercompetencycourse *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse `json:"usercompetencycourse,omitempty"`
}

// NewToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner instantiates a new ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner() *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner {
	this := ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner{}
	return &this
}

// NewToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerWithDefaults instantiates a new ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerWithDefaults() *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner {
	this := ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner{}
	return &this
}

// GetCompetency returns the Competency field value if set, zero value otherwise.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) GetCompetency() CoreCompetencyCreateCompetency200Response {
	if o == nil || IsNil(o.Competency) {
		var ret CoreCompetencyCreateCompetency200Response
		return ret
	}
	return *o.Competency
}

// GetCompetencyOk returns a tuple with the Competency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) GetCompetencyOk() (*CoreCompetencyCreateCompetency200Response, bool) {
	if o == nil || IsNil(o.Competency) {
		return nil, false
	}
	return o.Competency, true
}

// HasCompetency returns a boolean if a field has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) HasCompetency() bool {
	if o != nil && !IsNil(o.Competency) {
		return true
	}

	return false
}

// SetCompetency gets a reference to the given CoreCompetencyCreateCompetency200Response and assigns it to the Competency field.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) SetCompetency(v CoreCompetencyCreateCompetency200Response) {
	o.Competency = &v
}

// GetComppath returns the Comppath field value if set, zero value otherwise.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) GetComppath() ToolLpDataForCompetencySummary200ResponseComppath {
	if o == nil || IsNil(o.Comppath) {
		var ret ToolLpDataForCompetencySummary200ResponseComppath
		return ret
	}
	return *o.Comppath
}

// GetComppathOk returns a tuple with the Comppath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) GetComppathOk() (*ToolLpDataForCompetencySummary200ResponseComppath, bool) {
	if o == nil || IsNil(o.Comppath) {
		return nil, false
	}
	return o.Comppath, true
}

// HasComppath returns a boolean if a field has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) HasComppath() bool {
	if o != nil && !IsNil(o.Comppath) {
		return true
	}

	return false
}

// SetComppath gets a reference to the given ToolLpDataForCompetencySummary200ResponseComppath and assigns it to the Comppath field.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) SetComppath(v ToolLpDataForCompetencySummary200ResponseComppath) {
	o.Comppath = &v
}

// GetCoursecompetency returns the Coursecompetency field value if set, zero value otherwise.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) GetCoursecompetency() ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerCoursecompetency {
	if o == nil || IsNil(o.Coursecompetency) {
		var ret ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerCoursecompetency
		return ret
	}
	return *o.Coursecompetency
}

// GetCoursecompetencyOk returns a tuple with the Coursecompetency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) GetCoursecompetencyOk() (*ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerCoursecompetency, bool) {
	if o == nil || IsNil(o.Coursecompetency) {
		return nil, false
	}
	return o.Coursecompetency, true
}

// HasCoursecompetency returns a boolean if a field has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) HasCoursecompetency() bool {
	if o != nil && !IsNil(o.Coursecompetency) {
		return true
	}

	return false
}

// SetCoursecompetency gets a reference to the given ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerCoursecompetency and assigns it to the Coursecompetency field.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) SetCoursecompetency(v ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerCoursecompetency) {
	o.Coursecompetency = &v
}

// GetCoursemodules returns the Coursemodules field value if set, zero value otherwise.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) GetCoursemodules() []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerCoursemodulesInner {
	if o == nil || IsNil(o.Coursemodules) {
		var ret []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerCoursemodulesInner
		return ret
	}
	return o.Coursemodules
}

// GetCoursemodulesOk returns a tuple with the Coursemodules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) GetCoursemodulesOk() ([]ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerCoursemodulesInner, bool) {
	if o == nil || IsNil(o.Coursemodules) {
		return nil, false
	}
	return o.Coursemodules, true
}

// HasCoursemodules returns a boolean if a field has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) HasCoursemodules() bool {
	if o != nil && !IsNil(o.Coursemodules) {
		return true
	}

	return false
}

// SetCoursemodules gets a reference to the given []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerCoursemodulesInner and assigns it to the Coursemodules field.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) SetCoursemodules(v []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerCoursemodulesInner) {
	o.Coursemodules = v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) GetPlans() []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerPlansInner {
	if o == nil || IsNil(o.Plans) {
		var ret []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerPlansInner
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) GetPlansOk() ([]ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerPlansInner, bool) {
	if o == nil || IsNil(o.Plans) {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) HasPlans() bool {
	if o != nil && !IsNil(o.Plans) {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerPlansInner and assigns it to the Plans field.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) SetPlans(v []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerPlansInner) {
	o.Plans = v
}

// GetRuleoutcomeoptions returns the Ruleoutcomeoptions field value if set, zero value otherwise.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) GetRuleoutcomeoptions() []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerRuleoutcomeoptionsInner {
	if o == nil || IsNil(o.Ruleoutcomeoptions) {
		var ret []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerRuleoutcomeoptionsInner
		return ret
	}
	return o.Ruleoutcomeoptions
}

// GetRuleoutcomeoptionsOk returns a tuple with the Ruleoutcomeoptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) GetRuleoutcomeoptionsOk() ([]ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerRuleoutcomeoptionsInner, bool) {
	if o == nil || IsNil(o.Ruleoutcomeoptions) {
		return nil, false
	}
	return o.Ruleoutcomeoptions, true
}

// HasRuleoutcomeoptions returns a boolean if a field has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) HasRuleoutcomeoptions() bool {
	if o != nil && !IsNil(o.Ruleoutcomeoptions) {
		return true
	}

	return false
}

// SetRuleoutcomeoptions gets a reference to the given []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerRuleoutcomeoptionsInner and assigns it to the Ruleoutcomeoptions field.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) SetRuleoutcomeoptions(v []ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerRuleoutcomeoptionsInner) {
	o.Ruleoutcomeoptions = v
}

// GetUsercompetencycourse returns the Usercompetencycourse field value if set, zero value otherwise.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) GetUsercompetencycourse() ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse {
	if o == nil || IsNil(o.Usercompetencycourse) {
		var ret ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse
		return ret
	}
	return *o.Usercompetencycourse
}

// GetUsercompetencycourseOk returns a tuple with the Usercompetencycourse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) GetUsercompetencycourseOk() (*ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse, bool) {
	if o == nil || IsNil(o.Usercompetencycourse) {
		return nil, false
	}
	return o.Usercompetencycourse, true
}

// HasUsercompetencycourse returns a boolean if a field has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) HasUsercompetencycourse() bool {
	if o != nil && !IsNil(o.Usercompetencycourse) {
		return true
	}

	return false
}

// SetUsercompetencycourse gets a reference to the given ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse and assigns it to the Usercompetencycourse field.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) SetUsercompetencycourse(v ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) {
	o.Usercompetencycourse = &v
}

func (o ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Competency) {
		toSerialize["competency"] = o.Competency
	}
	if !IsNil(o.Comppath) {
		toSerialize["comppath"] = o.Comppath
	}
	if !IsNil(o.Coursecompetency) {
		toSerialize["coursecompetency"] = o.Coursecompetency
	}
	if !IsNil(o.Coursemodules) {
		toSerialize["coursemodules"] = o.Coursemodules
	}
	if !IsNil(o.Plans) {
		toSerialize["plans"] = o.Plans
	}
	if !IsNil(o.Ruleoutcomeoptions) {
		toSerialize["ruleoutcomeoptions"] = o.Ruleoutcomeoptions
	}
	if !IsNil(o.Usercompetencycourse) {
		toSerialize["usercompetencycourse"] = o.Usercompetencycourse
	}
	return toSerialize, nil
}

type NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner struct {
	value *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner
	isSet bool
}

func (v NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) Get() *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner {
	return v.value
}

func (v *NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) Set(val *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner(val *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) *NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner {
	return &NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner{value: val, isSet: true}
}

func (v NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


