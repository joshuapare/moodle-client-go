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

// checks if the ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner{}

// ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner struct for ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner
type ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner struct {
	// criterion id
	Criterionid *int32 `json:"criterionid,omitempty"`
	// level id
	Levelid *int32 `json:"levelid,omitempty"`
	// remark
	Remark *string `json:"remark,omitempty"`
	// remark format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Remarkformat *int32 `json:"remarkformat,omitempty"`
}

// NewModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner instantiates a new ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner() *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner {
	this := ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner{}
	return &this
}

// NewModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInnerWithDefaults instantiates a new ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInnerWithDefaults() *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner {
	this := ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner{}
	return &this
}

// GetCriterionid returns the Criterionid field value if set, zero value otherwise.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) GetCriterionid() int32 {
	if o == nil || IsNil(o.Criterionid) {
		var ret int32
		return ret
	}
	return *o.Criterionid
}

// GetCriterionidOk returns a tuple with the Criterionid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) GetCriterionidOk() (*int32, bool) {
	if o == nil || IsNil(o.Criterionid) {
		return nil, false
	}
	return o.Criterionid, true
}

// HasCriterionid returns a boolean if a field has been set.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) HasCriterionid() bool {
	if o != nil && !IsNil(o.Criterionid) {
		return true
	}

	return false
}

// SetCriterionid gets a reference to the given int32 and assigns it to the Criterionid field.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) SetCriterionid(v int32) {
	o.Criterionid = &v
}

// GetLevelid returns the Levelid field value if set, zero value otherwise.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) GetLevelid() int32 {
	if o == nil || IsNil(o.Levelid) {
		var ret int32
		return ret
	}
	return *o.Levelid
}

// GetLevelidOk returns a tuple with the Levelid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) GetLevelidOk() (*int32, bool) {
	if o == nil || IsNil(o.Levelid) {
		return nil, false
	}
	return o.Levelid, true
}

// HasLevelid returns a boolean if a field has been set.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) HasLevelid() bool {
	if o != nil && !IsNil(o.Levelid) {
		return true
	}

	return false
}

// SetLevelid gets a reference to the given int32 and assigns it to the Levelid field.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) SetLevelid(v int32) {
	o.Levelid = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) SetRemark(v string) {
	o.Remark = &v
}

// GetRemarkformat returns the Remarkformat field value if set, zero value otherwise.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) GetRemarkformat() int32 {
	if o == nil || IsNil(o.Remarkformat) {
		var ret int32
		return ret
	}
	return *o.Remarkformat
}

// GetRemarkformatOk returns a tuple with the Remarkformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) GetRemarkformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Remarkformat) {
		return nil, false
	}
	return o.Remarkformat, true
}

// HasRemarkformat returns a boolean if a field has been set.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) HasRemarkformat() bool {
	if o != nil && !IsNil(o.Remarkformat) {
		return true
	}

	return false
}

// SetRemarkformat gets a reference to the given int32 and assigns it to the Remarkformat field.
func (o *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) SetRemarkformat(v int32) {
	o.Remarkformat = &v
}

func (o ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criterionid) {
		toSerialize["criterionid"] = o.Criterionid
	}
	if !IsNil(o.Levelid) {
		toSerialize["levelid"] = o.Levelid
	}
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	if !IsNil(o.Remarkformat) {
		toSerialize["remarkformat"] = o.Remarkformat
	}
	return toSerialize, nil
}

type NullableModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner struct {
	value *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner
	isSet bool
}

func (v NullableModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) Get() *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner {
	return v.value
}

func (v *NullableModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) Set(val *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner(val *ModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) *NullableModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner {
	return &NullableModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner{value: val, isSet: true}
}

func (v NullableModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModAssignSaveGradeRequestAdvancedgradingdataRubricCriteriaInnerFillingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


