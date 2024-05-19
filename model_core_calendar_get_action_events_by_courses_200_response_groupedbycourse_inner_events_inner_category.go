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

// checks if the CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory{}

// CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory struct for CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory
type CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory struct {
	// coursecount
	Coursecount int32 `json:"coursecount"`
	// depth
	Depth int32 `json:"depth"`
	// description
	Description *string `json:"description,omitempty"`
	// id
	Id int32 `json:"id"`
	// idnumber
	Idnumber string `json:"idnumber"`
	// name
	Name string `json:"name"`
	// nestedname
	Nestedname string `json:"nestedname"`
	// parent
	Parent int32 `json:"parent"`
	// timemodified
	Timemodified int32 `json:"timemodified"`
	// url
	Url string `json:"url"`
	// visible
	Visible int32 `json:"visible"`
}

type _CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory

// NewCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory instantiates a new CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory(coursecount int32, depth int32, id int32, idnumber string, name string, nestedname string, parent int32, timemodified int32, url string, visible int32) *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory {
	this := CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory{}
	this.Coursecount = coursecount
	this.Depth = depth
	this.Id = id
	this.Idnumber = idnumber
	this.Name = name
	this.Nestedname = nestedname
	this.Parent = parent
	this.Timemodified = timemodified
	this.Url = url
	this.Visible = visible
	return &this
}

// NewCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategoryWithDefaults instantiates a new CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategoryWithDefaults() *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory {
	this := CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory{}
	var coursecount int32 = 0
	this.Coursecount = coursecount
	var depth int32 = 0
	this.Depth = depth
	var name string = ""
	this.Name = name
	var timemodified int32 = 0
	this.Timemodified = timemodified
	var visible int32 = 1
	this.Visible = visible
	return &this
}

// GetCoursecount returns the Coursecount field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetCoursecount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Coursecount
}

// GetCoursecountOk returns a tuple with the Coursecount field value
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetCoursecountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coursecount, true
}

// SetCoursecount sets field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) SetCoursecount(v int32) {
	o.Coursecount = v
}

// GetDepth returns the Depth field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) SetDepth(v int32) {
	o.Depth = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) SetId(v int32) {
	o.Id = v
}

// GetIdnumber returns the Idnumber field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetIdnumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Idnumber
}

// GetIdnumberOk returns a tuple with the Idnumber field value
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetIdnumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Idnumber, true
}

// SetIdnumber sets field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) SetIdnumber(v string) {
	o.Idnumber = v
}

// GetName returns the Name field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) SetName(v string) {
	o.Name = v
}

// GetNestedname returns the Nestedname field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetNestedname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nestedname
}

// GetNestednameOk returns a tuple with the Nestedname field value
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetNestednameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nestedname, true
}

// SetNestedname sets field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) SetNestedname(v string) {
	o.Nestedname = v
}

// GetParent returns the Parent field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetParent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetParentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) SetParent(v int32) {
	o.Parent = v
}

// GetTimemodified returns the Timemodified field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetTimemodified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetTimemodifiedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timemodified, true
}

// SetTimemodified sets field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) SetTimemodified(v int32) {
	o.Timemodified = v
}

// GetUrl returns the Url field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) SetUrl(v string) {
	o.Url = v
}

// GetVisible returns the Visible field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetVisible() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) GetVisibleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visible, true
}

// SetVisible sets field value
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) SetVisible(v int32) {
	o.Visible = v
}

func (o CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["coursecount"] = o.Coursecount
	toSerialize["depth"] = o.Depth
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	toSerialize["idnumber"] = o.Idnumber
	toSerialize["name"] = o.Name
	toSerialize["nestedname"] = o.Nestedname
	toSerialize["parent"] = o.Parent
	toSerialize["timemodified"] = o.Timemodified
	toSerialize["url"] = o.Url
	toSerialize["visible"] = o.Visible
	return toSerialize, nil
}

func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"coursecount",
		"depth",
		"id",
		"idnumber",
		"name",
		"nestedname",
		"parent",
		"timemodified",
		"url",
		"visible",
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

	varCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory := _CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory)

	if err != nil {
		return err
	}

	*o = CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory(varCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory)

	return err
}

type NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory struct {
	value *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory
	isSet bool
}

func (v NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) Get() *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory {
	return v.value
}

func (v *NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) Set(val *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory(val *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) *NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory {
	return &NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory{value: val, isSet: true}
}

func (v NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInnerCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


