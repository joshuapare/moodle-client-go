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

// checks if the CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide{}

// CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide struct for CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide
type CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide struct {
	GuideComments []CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuideGuideCommentsInner `json:"guide_comments,omitempty"`
	GuideCriteria []CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuideGuideCriteriaInner `json:"guide_criteria,omitempty"`
}

// NewCoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide instantiates a new CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide() *CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide {
	this := CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide{}
	return &this
}

// NewCoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuideWithDefaults instantiates a new CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuideWithDefaults() *CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide {
	this := CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide{}
	return &this
}

// GetGuideComments returns the GuideComments field value if set, zero value otherwise.
func (o *CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) GetGuideComments() []CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuideGuideCommentsInner {
	if o == nil || IsNil(o.GuideComments) {
		var ret []CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuideGuideCommentsInner
		return ret
	}
	return o.GuideComments
}

// GetGuideCommentsOk returns a tuple with the GuideComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) GetGuideCommentsOk() ([]CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuideGuideCommentsInner, bool) {
	if o == nil || IsNil(o.GuideComments) {
		return nil, false
	}
	return o.GuideComments, true
}

// HasGuideComments returns a boolean if a field has been set.
func (o *CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) HasGuideComments() bool {
	if o != nil && !IsNil(o.GuideComments) {
		return true
	}

	return false
}

// SetGuideComments gets a reference to the given []CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuideGuideCommentsInner and assigns it to the GuideComments field.
func (o *CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) SetGuideComments(v []CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuideGuideCommentsInner) {
	o.GuideComments = v
}

// GetGuideCriteria returns the GuideCriteria field value if set, zero value otherwise.
func (o *CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) GetGuideCriteria() []CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuideGuideCriteriaInner {
	if o == nil || IsNil(o.GuideCriteria) {
		var ret []CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuideGuideCriteriaInner
		return ret
	}
	return o.GuideCriteria
}

// GetGuideCriteriaOk returns a tuple with the GuideCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) GetGuideCriteriaOk() ([]CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuideGuideCriteriaInner, bool) {
	if o == nil || IsNil(o.GuideCriteria) {
		return nil, false
	}
	return o.GuideCriteria, true
}

// HasGuideCriteria returns a boolean if a field has been set.
func (o *CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) HasGuideCriteria() bool {
	if o != nil && !IsNil(o.GuideCriteria) {
		return true
	}

	return false
}

// SetGuideCriteria gets a reference to the given []CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuideGuideCriteriaInner and assigns it to the GuideCriteria field.
func (o *CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) SetGuideCriteria(v []CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuideGuideCriteriaInner) {
	o.GuideCriteria = v
}

func (o CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GuideComments) {
		toSerialize["guide_comments"] = o.GuideComments
	}
	if !IsNil(o.GuideCriteria) {
		toSerialize["guide_criteria"] = o.GuideCriteria
	}
	return toSerialize, nil
}

type NullableCoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide struct {
	value *CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide
	isSet bool
}

func (v NullableCoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) Get() *CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide {
	return v.value
}

func (v *NullableCoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) Set(val *CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide(val *CoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) *NullableCoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide {
	return &NullableCoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide{value: val, isSet: true}
}

func (v NullableCoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGradingSaveDefinitionsRequestAreasInnerDefinitionsInnerGuide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


