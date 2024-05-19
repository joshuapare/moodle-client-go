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

// checks if the ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner{}

// ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner struct for ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner
type ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner struct {
	// first
	First *bool `json:"first,omitempty"`
	// id
	Id *int32 `json:"id,omitempty"`
	// last
	Last *bool `json:"last,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
	// position
	Position *int32 `json:"position,omitempty"`
}

// NewReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner instantiates a new ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner() *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner {
	this := ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner{}
	var first bool = null
	this.First = &first
	var last bool = null
	this.Last = &last
	var position int32 = null
	this.Position = &position
	return &this
}

// NewReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInnerWithDefaults instantiates a new ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInnerWithDefaults() *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner {
	this := ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner{}
	var first bool = null
	this.First = &first
	var last bool = null
	this.Last = &last
	var position int32 = null
	this.Position = &position
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) GetFirst() bool {
	if o == nil || IsNil(o.First) {
		var ret bool
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) GetFirstOk() (*bool, bool) {
	if o == nil || IsNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) HasFirst() bool {
	if o != nil && !IsNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given bool and assigns it to the First field.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) SetFirst(v bool) {
	o.First = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) SetId(v int32) {
	o.Id = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) GetLast() bool {
	if o == nil || IsNil(o.Last) {
		var ret bool
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) GetLastOk() (*bool, bool) {
	if o == nil || IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given bool and assigns it to the Last field.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) SetLast(v bool) {
	o.Last = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) SetName(v string) {
	o.Name = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) SetPosition(v int32) {
	o.Position = &v
}

func (o ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.First) {
		toSerialize["first"] = o.First
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	return toSerialize, nil
}

type NullableReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner struct {
	value *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner
	isSet bool
}

func (v NullableReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) Get() *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner {
	return v.value
}

func (v *NullableReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) Set(val *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner(val *ReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) *NullableReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner {
	return &NullableReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner{value: val, isSet: true}
}

func (v NullableReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportCompetencyDataForReport200ResponseUsercompetenciesInnerCompetencyComppathAncestorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

