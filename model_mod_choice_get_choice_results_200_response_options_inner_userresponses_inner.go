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

// checks if the ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner{}

// ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner User responses
type ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner struct {
	// answer id
	Answerid *int32 `json:"answerid,omitempty"`
	// user full name
	Fullname *string `json:"fullname,omitempty"`
	// profile user image url
	Profileimageurl *string `json:"profileimageurl,omitempty"`
	// time of modification
	Timemodified *int32 `json:"timemodified,omitempty"`
	// user id
	Userid *int32 `json:"userid,omitempty"`
}

// NewModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner instantiates a new ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner() *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner {
	this := ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner{}
	var answerid int32 = null
	this.Answerid = &answerid
	var profileimageurl string = "null"
	this.Profileimageurl = &profileimageurl
	var timemodified int32 = null
	this.Timemodified = &timemodified
	return &this
}

// NewModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInnerWithDefaults instantiates a new ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInnerWithDefaults() *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner {
	this := ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner{}
	var answerid int32 = null
	this.Answerid = &answerid
	var profileimageurl string = "null"
	this.Profileimageurl = &profileimageurl
	var timemodified int32 = null
	this.Timemodified = &timemodified
	return &this
}

// GetAnswerid returns the Answerid field value if set, zero value otherwise.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) GetAnswerid() int32 {
	if o == nil || IsNil(o.Answerid) {
		var ret int32
		return ret
	}
	return *o.Answerid
}

// GetAnsweridOk returns a tuple with the Answerid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) GetAnsweridOk() (*int32, bool) {
	if o == nil || IsNil(o.Answerid) {
		return nil, false
	}
	return o.Answerid, true
}

// HasAnswerid returns a boolean if a field has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) HasAnswerid() bool {
	if o != nil && !IsNil(o.Answerid) {
		return true
	}

	return false
}

// SetAnswerid gets a reference to the given int32 and assigns it to the Answerid field.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) SetAnswerid(v int32) {
	o.Answerid = &v
}

// GetFullname returns the Fullname field value if set, zero value otherwise.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) GetFullname() string {
	if o == nil || IsNil(o.Fullname) {
		var ret string
		return ret
	}
	return *o.Fullname
}

// GetFullnameOk returns a tuple with the Fullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) GetFullnameOk() (*string, bool) {
	if o == nil || IsNil(o.Fullname) {
		return nil, false
	}
	return o.Fullname, true
}

// HasFullname returns a boolean if a field has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) HasFullname() bool {
	if o != nil && !IsNil(o.Fullname) {
		return true
	}

	return false
}

// SetFullname gets a reference to the given string and assigns it to the Fullname field.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) SetFullname(v string) {
	o.Fullname = &v
}

// GetProfileimageurl returns the Profileimageurl field value if set, zero value otherwise.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) GetProfileimageurl() string {
	if o == nil || IsNil(o.Profileimageurl) {
		var ret string
		return ret
	}
	return *o.Profileimageurl
}

// GetProfileimageurlOk returns a tuple with the Profileimageurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) GetProfileimageurlOk() (*string, bool) {
	if o == nil || IsNil(o.Profileimageurl) {
		return nil, false
	}
	return o.Profileimageurl, true
}

// HasProfileimageurl returns a boolean if a field has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) HasProfileimageurl() bool {
	if o != nil && !IsNil(o.Profileimageurl) {
		return true
	}

	return false
}

// SetProfileimageurl gets a reference to the given string and assigns it to the Profileimageurl field.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) SetProfileimageurl(v string) {
	o.Profileimageurl = &v
}

// GetTimemodified returns the Timemodified field value if set, zero value otherwise.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) GetTimemodified() int32 {
	if o == nil || IsNil(o.Timemodified) {
		var ret int32
		return ret
	}
	return *o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) GetTimemodifiedOk() (*int32, bool) {
	if o == nil || IsNil(o.Timemodified) {
		return nil, false
	}
	return o.Timemodified, true
}

// HasTimemodified returns a boolean if a field has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) HasTimemodified() bool {
	if o != nil && !IsNil(o.Timemodified) {
		return true
	}

	return false
}

// SetTimemodified gets a reference to the given int32 and assigns it to the Timemodified field.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) SetTimemodified(v int32) {
	o.Timemodified = &v
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) GetUserid() int32 {
	if o == nil || IsNil(o.Userid) {
		var ret int32
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) GetUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given int32 and assigns it to the Userid field.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) SetUserid(v int32) {
	o.Userid = &v
}

func (o ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Answerid) {
		toSerialize["answerid"] = o.Answerid
	}
	if !IsNil(o.Fullname) {
		toSerialize["fullname"] = o.Fullname
	}
	if !IsNil(o.Profileimageurl) {
		toSerialize["profileimageurl"] = o.Profileimageurl
	}
	if !IsNil(o.Timemodified) {
		toSerialize["timemodified"] = o.Timemodified
	}
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	return toSerialize, nil
}

type NullableModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner struct {
	value *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner
	isSet bool
}

func (v NullableModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) Get() *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner {
	return v.value
}

func (v *NullableModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) Set(val *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner(val *ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) *NullableModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner {
	return &NullableModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner{value: val, isSet: true}
}

func (v NullableModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


