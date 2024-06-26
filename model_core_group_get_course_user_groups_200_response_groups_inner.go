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

// checks if the CoreGroupGetCourseUserGroups200ResponseGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGroupGetCourseUserGroups200ResponseGroupsInner{}

// CoreGroupGetCourseUserGroups200ResponseGroupsInner struct for CoreGroupGetCourseUserGroups200ResponseGroupsInner
type CoreGroupGetCourseUserGroups200ResponseGroupsInner struct {
	// course id
	Courseid *int32 `json:"courseid,omitempty"`
	// group description text
	Description *string `json:"description,omitempty"`
	// description format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Descriptionformat *int32 `json:"descriptionformat,omitempty"`
	// group record id
	Id *int32 `json:"id,omitempty"`
	// id number
	Idnumber *string `json:"idnumber,omitempty"`
	// group name
	Name *string `json:"name,omitempty"`
}

// NewCoreGroupGetCourseUserGroups200ResponseGroupsInner instantiates a new CoreGroupGetCourseUserGroups200ResponseGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGroupGetCourseUserGroups200ResponseGroupsInner() *CoreGroupGetCourseUserGroups200ResponseGroupsInner {
	this := CoreGroupGetCourseUserGroups200ResponseGroupsInner{}
	return &this
}

// NewCoreGroupGetCourseUserGroups200ResponseGroupsInnerWithDefaults instantiates a new CoreGroupGetCourseUserGroups200ResponseGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGroupGetCourseUserGroups200ResponseGroupsInnerWithDefaults() *CoreGroupGetCourseUserGroups200ResponseGroupsInner {
	this := CoreGroupGetCourseUserGroups200ResponseGroupsInner{}
	return &this
}

// GetCourseid returns the Courseid field value if set, zero value otherwise.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) GetCourseid() int32 {
	if o == nil || IsNil(o.Courseid) {
		var ret int32
		return ret
	}
	return *o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) GetCourseidOk() (*int32, bool) {
	if o == nil || IsNil(o.Courseid) {
		return nil, false
	}
	return o.Courseid, true
}

// HasCourseid returns a boolean if a field has been set.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) HasCourseid() bool {
	if o != nil && !IsNil(o.Courseid) {
		return true
	}

	return false
}

// SetCourseid gets a reference to the given int32 and assigns it to the Courseid field.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) SetCourseid(v int32) {
	o.Courseid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptionformat returns the Descriptionformat field value if set, zero value otherwise.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) GetDescriptionformat() int32 {
	if o == nil || IsNil(o.Descriptionformat) {
		var ret int32
		return ret
	}
	return *o.Descriptionformat
}

// GetDescriptionformatOk returns a tuple with the Descriptionformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) GetDescriptionformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Descriptionformat) {
		return nil, false
	}
	return o.Descriptionformat, true
}

// HasDescriptionformat returns a boolean if a field has been set.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) HasDescriptionformat() bool {
	if o != nil && !IsNil(o.Descriptionformat) {
		return true
	}

	return false
}

// SetDescriptionformat gets a reference to the given int32 and assigns it to the Descriptionformat field.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) SetDescriptionformat(v int32) {
	o.Descriptionformat = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) SetId(v int32) {
	o.Id = &v
}

// GetIdnumber returns the Idnumber field value if set, zero value otherwise.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) GetIdnumber() string {
	if o == nil || IsNil(o.Idnumber) {
		var ret string
		return ret
	}
	return *o.Idnumber
}

// GetIdnumberOk returns a tuple with the Idnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) GetIdnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Idnumber) {
		return nil, false
	}
	return o.Idnumber, true
}

// HasIdnumber returns a boolean if a field has been set.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) HasIdnumber() bool {
	if o != nil && !IsNil(o.Idnumber) {
		return true
	}

	return false
}

// SetIdnumber gets a reference to the given string and assigns it to the Idnumber field.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) SetIdnumber(v string) {
	o.Idnumber = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreGroupGetCourseUserGroups200ResponseGroupsInner) SetName(v string) {
	o.Name = &v
}

func (o CoreGroupGetCourseUserGroups200ResponseGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGroupGetCourseUserGroups200ResponseGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Courseid) {
		toSerialize["courseid"] = o.Courseid
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Descriptionformat) {
		toSerialize["descriptionformat"] = o.Descriptionformat
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Idnumber) {
		toSerialize["idnumber"] = o.Idnumber
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCoreGroupGetCourseUserGroups200ResponseGroupsInner struct {
	value *CoreGroupGetCourseUserGroups200ResponseGroupsInner
	isSet bool
}

func (v NullableCoreGroupGetCourseUserGroups200ResponseGroupsInner) Get() *CoreGroupGetCourseUserGroups200ResponseGroupsInner {
	return v.value
}

func (v *NullableCoreGroupGetCourseUserGroups200ResponseGroupsInner) Set(val *CoreGroupGetCourseUserGroups200ResponseGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGroupGetCourseUserGroups200ResponseGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGroupGetCourseUserGroups200ResponseGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGroupGetCourseUserGroups200ResponseGroupsInner(val *CoreGroupGetCourseUserGroups200ResponseGroupsInner) *NullableCoreGroupGetCourseUserGroups200ResponseGroupsInner {
	return &NullableCoreGroupGetCourseUserGroups200ResponseGroupsInner{value: val, isSet: true}
}

func (v NullableCoreGroupGetCourseUserGroups200ResponseGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGroupGetCourseUserGroups200ResponseGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


