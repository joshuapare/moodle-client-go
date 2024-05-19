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

// checks if the CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner{}

// CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner struct for CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner
type CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner struct {
	// Whether this item is currently selected
	Selected *bool `json:"selected,omitempty"`
	// The description fo the option
	Title *string `json:"title,omitempty"`
	// The grade value
	Value *float32 `json:"value,omitempty"`
}

// NewCoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner instantiates a new CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner() *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner {
	this := CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner{}
	return &this
}

// NewCoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInnerWithDefaults instantiates a new CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInnerWithDefaults() *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner {
	this := CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner{}
	return &this
}

// GetSelected returns the Selected field value if set, zero value otherwise.
func (o *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) GetSelected() bool {
	if o == nil || IsNil(o.Selected) {
		var ret bool
		return ret
	}
	return *o.Selected
}

// GetSelectedOk returns a tuple with the Selected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) GetSelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Selected) {
		return nil, false
	}
	return o.Selected, true
}

// HasSelected returns a boolean if a field has been set.
func (o *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) HasSelected() bool {
	if o != nil && !IsNil(o.Selected) {
		return true
	}

	return false
}

// SetSelected gets a reference to the given bool and assigns it to the Selected field.
func (o *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) SetSelected(v bool) {
	o.Selected = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) SetTitle(v string) {
	o.Title = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) SetValue(v float32) {
	o.Value = &v
}

func (o CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Selected) {
		toSerialize["selected"] = o.Selected
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner struct {
	value *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner
	isSet bool
}

func (v NullableCoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) Get() *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner {
	return v.value
}

func (v *NullableCoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) Set(val *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner(val *CoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) *NullableCoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner {
	return &NullableCoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner{value: val, isSet: true}
}

func (v NullableCoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGradesGraderGradingpanelScaleStore200ResponseGradeOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

