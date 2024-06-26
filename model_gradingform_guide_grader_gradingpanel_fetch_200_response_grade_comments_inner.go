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

// checks if the GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner{}

// GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner struct for GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner
type GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner struct {
	// The comment value
	Description *string `json:"description,omitempty"`
	// Comment id
	Id *int32 `json:"id,omitempty"`
	// The sortorder of this comment
	Sortorder *int32 `json:"sortorder,omitempty"`
}

// NewGradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner instantiates a new GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner() *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner {
	this := GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner{}
	var description string = "null"
	this.Description = &description
	var id int32 = null
	this.Id = &id
	var sortorder int32 = null
	this.Sortorder = &sortorder
	return &this
}

// NewGradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInnerWithDefaults instantiates a new GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInnerWithDefaults() *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner {
	this := GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner{}
	var description string = "null"
	this.Description = &description
	var id int32 = null
	this.Id = &id
	var sortorder int32 = null
	this.Sortorder = &sortorder
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) SetId(v int32) {
	o.Id = &v
}

// GetSortorder returns the Sortorder field value if set, zero value otherwise.
func (o *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) GetSortorder() int32 {
	if o == nil || IsNil(o.Sortorder) {
		var ret int32
		return ret
	}
	return *o.Sortorder
}

// GetSortorderOk returns a tuple with the Sortorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) GetSortorderOk() (*int32, bool) {
	if o == nil || IsNil(o.Sortorder) {
		return nil, false
	}
	return o.Sortorder, true
}

// HasSortorder returns a boolean if a field has been set.
func (o *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) HasSortorder() bool {
	if o != nil && !IsNil(o.Sortorder) {
		return true
	}

	return false
}

// SetSortorder gets a reference to the given int32 and assigns it to the Sortorder field.
func (o *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) SetSortorder(v int32) {
	o.Sortorder = &v
}

func (o GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Sortorder) {
		toSerialize["sortorder"] = o.Sortorder
	}
	return toSerialize, nil
}

type NullableGradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner struct {
	value *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner
	isSet bool
}

func (v NullableGradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) Get() *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner {
	return v.value
}

func (v *NullableGradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) Set(val *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner(val *GradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) *NullableGradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner {
	return &NullableGradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner{value: val, isSet: true}
}

func (v NullableGradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGradingformGuideGraderGradingpanelFetch200ResponseGradeCommentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


