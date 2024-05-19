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

// checks if the CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner{}

// CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner struct for CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner
type CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner struct {
	// courseid
	Courseid *int32 `json:"courseid,omitempty"`
	Events []CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInner `json:"events,omitempty"`
	// firstid
	Firstid *int32 `json:"firstid,omitempty"`
	// lastid
	Lastid *int32 `json:"lastid,omitempty"`
}

// NewCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner instantiates a new CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner() *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner {
	this := CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner{}
	var courseid int32 = null
	this.Courseid = &courseid
	return &this
}

// NewCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerWithDefaults instantiates a new CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerWithDefaults() *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner {
	this := CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner{}
	var courseid int32 = null
	this.Courseid = &courseid
	return &this
}

// GetCourseid returns the Courseid field value if set, zero value otherwise.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) GetCourseid() int32 {
	if o == nil || IsNil(o.Courseid) {
		var ret int32
		return ret
	}
	return *o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) GetCourseidOk() (*int32, bool) {
	if o == nil || IsNil(o.Courseid) {
		return nil, false
	}
	return o.Courseid, true
}

// HasCourseid returns a boolean if a field has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) HasCourseid() bool {
	if o != nil && !IsNil(o.Courseid) {
		return true
	}

	return false
}

// SetCourseid gets a reference to the given int32 and assigns it to the Courseid field.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) SetCourseid(v int32) {
	o.Courseid = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) GetEvents() []CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInner {
	if o == nil || IsNil(o.Events) {
		var ret []CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInner
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) GetEventsOk() ([]CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInner, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInner and assigns it to the Events field.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) SetEvents(v []CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInnerEventsInner) {
	o.Events = v
}

// GetFirstid returns the Firstid field value if set, zero value otherwise.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) GetFirstid() int32 {
	if o == nil || IsNil(o.Firstid) {
		var ret int32
		return ret
	}
	return *o.Firstid
}

// GetFirstidOk returns a tuple with the Firstid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) GetFirstidOk() (*int32, bool) {
	if o == nil || IsNil(o.Firstid) {
		return nil, false
	}
	return o.Firstid, true
}

// HasFirstid returns a boolean if a field has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) HasFirstid() bool {
	if o != nil && !IsNil(o.Firstid) {
		return true
	}

	return false
}

// SetFirstid gets a reference to the given int32 and assigns it to the Firstid field.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) SetFirstid(v int32) {
	o.Firstid = &v
}

// GetLastid returns the Lastid field value if set, zero value otherwise.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) GetLastid() int32 {
	if o == nil || IsNil(o.Lastid) {
		var ret int32
		return ret
	}
	return *o.Lastid
}

// GetLastidOk returns a tuple with the Lastid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) GetLastidOk() (*int32, bool) {
	if o == nil || IsNil(o.Lastid) {
		return nil, false
	}
	return o.Lastid, true
}

// HasLastid returns a boolean if a field has been set.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) HasLastid() bool {
	if o != nil && !IsNil(o.Lastid) {
		return true
	}

	return false
}

// SetLastid gets a reference to the given int32 and assigns it to the Lastid field.
func (o *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) SetLastid(v int32) {
	o.Lastid = &v
}

func (o CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Courseid) {
		toSerialize["courseid"] = o.Courseid
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Firstid) {
		toSerialize["firstid"] = o.Firstid
	}
	if !IsNil(o.Lastid) {
		toSerialize["lastid"] = o.Lastid
	}
	return toSerialize, nil
}

type NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner struct {
	value *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner
	isSet bool
}

func (v NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) Get() *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner {
	return v.value
}

func (v *NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) Set(val *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner(val *CoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) *NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner {
	return &NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner{value: val, isSet: true}
}

func (v NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCalendarGetActionEventsByCourses200ResponseGroupedbycourseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

