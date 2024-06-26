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

// checks if the CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner{}

// CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner struct for CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner
type CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner struct {
	// The first name(s) of the user
	Firstname *string `json:"firstname,omitempty"`
	// The full name of the user
	Fullname *string `json:"fullname,omitempty"`
	// ID of the user
	Id *int32 `json:"id,omitempty"`
	// The family name of the user
	Lastname *string `json:"lastname,omitempty"`
	// The location of the users larger image
	Profileimage *string `json:"profileimage,omitempty"`
}

// NewCoreCourseGetEnrolledUsersByCmid200ResponseUsersInner instantiates a new CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseGetEnrolledUsersByCmid200ResponseUsersInner() *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner {
	this := CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner{}
	var fullname string = "null"
	this.Fullname = &fullname
	var id int32 = null
	this.Id = &id
	var profileimage string = "null"
	this.Profileimage = &profileimage
	return &this
}

// NewCoreCourseGetEnrolledUsersByCmid200ResponseUsersInnerWithDefaults instantiates a new CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseGetEnrolledUsersByCmid200ResponseUsersInnerWithDefaults() *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner {
	this := CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner{}
	var fullname string = "null"
	this.Fullname = &fullname
	var id int32 = null
	this.Id = &id
	var profileimage string = "null"
	this.Profileimage = &profileimage
	return &this
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) GetFirstname() string {
	if o == nil || IsNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) GetFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) HasFirstname() bool {
	if o != nil && !IsNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) SetFirstname(v string) {
	o.Firstname = &v
}

// GetFullname returns the Fullname field value if set, zero value otherwise.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) GetFullname() string {
	if o == nil || IsNil(o.Fullname) {
		var ret string
		return ret
	}
	return *o.Fullname
}

// GetFullnameOk returns a tuple with the Fullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) GetFullnameOk() (*string, bool) {
	if o == nil || IsNil(o.Fullname) {
		return nil, false
	}
	return o.Fullname, true
}

// HasFullname returns a boolean if a field has been set.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) HasFullname() bool {
	if o != nil && !IsNil(o.Fullname) {
		return true
	}

	return false
}

// SetFullname gets a reference to the given string and assigns it to the Fullname field.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) SetFullname(v string) {
	o.Fullname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) SetId(v int32) {
	o.Id = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) GetLastname() string {
	if o == nil || IsNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) GetLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) HasLastname() bool {
	if o != nil && !IsNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) SetLastname(v string) {
	o.Lastname = &v
}

// GetProfileimage returns the Profileimage field value if set, zero value otherwise.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) GetProfileimage() string {
	if o == nil || IsNil(o.Profileimage) {
		var ret string
		return ret
	}
	return *o.Profileimage
}

// GetProfileimageOk returns a tuple with the Profileimage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) GetProfileimageOk() (*string, bool) {
	if o == nil || IsNil(o.Profileimage) {
		return nil, false
	}
	return o.Profileimage, true
}

// HasProfileimage returns a boolean if a field has been set.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) HasProfileimage() bool {
	if o != nil && !IsNil(o.Profileimage) {
		return true
	}

	return false
}

// SetProfileimage gets a reference to the given string and assigns it to the Profileimage field.
func (o *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) SetProfileimage(v string) {
	o.Profileimage = &v
}

func (o CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !IsNil(o.Fullname) {
		toSerialize["fullname"] = o.Fullname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !IsNil(o.Profileimage) {
		toSerialize["profileimage"] = o.Profileimage
	}
	return toSerialize, nil
}

type NullableCoreCourseGetEnrolledUsersByCmid200ResponseUsersInner struct {
	value *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner
	isSet bool
}

func (v NullableCoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) Get() *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner {
	return v.value
}

func (v *NullableCoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) Set(val *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseGetEnrolledUsersByCmid200ResponseUsersInner(val *CoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) *NullableCoreCourseGetEnrolledUsersByCmid200ResponseUsersInner {
	return &NullableCoreCourseGetEnrolledUsersByCmid200ResponseUsersInner{value: val, isSet: true}
}

func (v NullableCoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseGetEnrolledUsersByCmid200ResponseUsersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


