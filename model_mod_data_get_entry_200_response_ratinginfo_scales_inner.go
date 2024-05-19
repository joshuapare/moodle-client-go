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

// checks if the ModDataGetEntry200ResponseRatinginfoScalesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModDataGetEntry200ResponseRatinginfoScalesInner{}

// ModDataGetEntry200ResponseRatinginfoScalesInner Scale information
type ModDataGetEntry200ResponseRatinginfoScalesInner struct {
	// Course id.
	Courseid *int32 `json:"courseid,omitempty"`
	// Scale id.
	Id *int32 `json:"id,omitempty"`
	// Whether is a numeric scale.
	Isnumeric *bool `json:"isnumeric,omitempty"`
	Items []ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner `json:"items,omitempty"`
	// Max value for the scale.
	Max *int32 `json:"max,omitempty"`
	// Scale name (when a real scale is used).
	Name *string `json:"name,omitempty"`
}

// NewModDataGetEntry200ResponseRatinginfoScalesInner instantiates a new ModDataGetEntry200ResponseRatinginfoScalesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModDataGetEntry200ResponseRatinginfoScalesInner() *ModDataGetEntry200ResponseRatinginfoScalesInner {
	this := ModDataGetEntry200ResponseRatinginfoScalesInner{}
	var courseid int32 = null
	this.Courseid = &courseid
	var isnumeric bool = null
	this.Isnumeric = &isnumeric
	var max int32 = null
	this.Max = &max
	var name string = "null"
	this.Name = &name
	return &this
}

// NewModDataGetEntry200ResponseRatinginfoScalesInnerWithDefaults instantiates a new ModDataGetEntry200ResponseRatinginfoScalesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModDataGetEntry200ResponseRatinginfoScalesInnerWithDefaults() *ModDataGetEntry200ResponseRatinginfoScalesInner {
	this := ModDataGetEntry200ResponseRatinginfoScalesInner{}
	var courseid int32 = null
	this.Courseid = &courseid
	var isnumeric bool = null
	this.Isnumeric = &isnumeric
	var max int32 = null
	this.Max = &max
	var name string = "null"
	this.Name = &name
	return &this
}

// GetCourseid returns the Courseid field value if set, zero value otherwise.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) GetCourseid() int32 {
	if o == nil || IsNil(o.Courseid) {
		var ret int32
		return ret
	}
	return *o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) GetCourseidOk() (*int32, bool) {
	if o == nil || IsNil(o.Courseid) {
		return nil, false
	}
	return o.Courseid, true
}

// HasCourseid returns a boolean if a field has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) HasCourseid() bool {
	if o != nil && !IsNil(o.Courseid) {
		return true
	}

	return false
}

// SetCourseid gets a reference to the given int32 and assigns it to the Courseid field.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) SetCourseid(v int32) {
	o.Courseid = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) SetId(v int32) {
	o.Id = &v
}

// GetIsnumeric returns the Isnumeric field value if set, zero value otherwise.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) GetIsnumeric() bool {
	if o == nil || IsNil(o.Isnumeric) {
		var ret bool
		return ret
	}
	return *o.Isnumeric
}

// GetIsnumericOk returns a tuple with the Isnumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) GetIsnumericOk() (*bool, bool) {
	if o == nil || IsNil(o.Isnumeric) {
		return nil, false
	}
	return o.Isnumeric, true
}

// HasIsnumeric returns a boolean if a field has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) HasIsnumeric() bool {
	if o != nil && !IsNil(o.Isnumeric) {
		return true
	}

	return false
}

// SetIsnumeric gets a reference to the given bool and assigns it to the Isnumeric field.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) SetIsnumeric(v bool) {
	o.Isnumeric = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) GetItems() []ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner {
	if o == nil || IsNil(o.Items) {
		var ret []ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) GetItemsOk() ([]ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner and assigns it to the Items field.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) SetItems(v []ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) {
	o.Items = v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) SetMax(v int32) {
	o.Max = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInner) SetName(v string) {
	o.Name = &v
}

func (o ModDataGetEntry200ResponseRatinginfoScalesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModDataGetEntry200ResponseRatinginfoScalesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Courseid) {
		toSerialize["courseid"] = o.Courseid
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Isnumeric) {
		toSerialize["isnumeric"] = o.Isnumeric
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableModDataGetEntry200ResponseRatinginfoScalesInner struct {
	value *ModDataGetEntry200ResponseRatinginfoScalesInner
	isSet bool
}

func (v NullableModDataGetEntry200ResponseRatinginfoScalesInner) Get() *ModDataGetEntry200ResponseRatinginfoScalesInner {
	return v.value
}

func (v *NullableModDataGetEntry200ResponseRatinginfoScalesInner) Set(val *ModDataGetEntry200ResponseRatinginfoScalesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModDataGetEntry200ResponseRatinginfoScalesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModDataGetEntry200ResponseRatinginfoScalesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModDataGetEntry200ResponseRatinginfoScalesInner(val *ModDataGetEntry200ResponseRatinginfoScalesInner) *NullableModDataGetEntry200ResponseRatinginfoScalesInner {
	return &NullableModDataGetEntry200ResponseRatinginfoScalesInner{value: val, isSet: true}
}

func (v NullableModDataGetEntry200ResponseRatinginfoScalesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModDataGetEntry200ResponseRatinginfoScalesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

