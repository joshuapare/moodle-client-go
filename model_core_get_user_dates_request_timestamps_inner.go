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

// checks if the CoreGetUserDatesRequestTimestampsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGetUserDatesRequestTimestampsInner{}

// CoreGetUserDatesRequestTimestampsInner struct for CoreGetUserDatesRequestTimestampsInner
type CoreGetUserDatesRequestTimestampsInner struct {
	// Remove leading zero for day
	Fixday *int32 `json:"fixday,omitempty"`
	// Remove leading zero for hour
	Fixhour *int32 `json:"fixhour,omitempty"`
	// format string
	Format *string `json:"format,omitempty"`
	// unix timestamp
	Timestamp *int32 `json:"timestamp,omitempty"`
	// The calendar type
	Type *string `json:"type,omitempty"`
}

// NewCoreGetUserDatesRequestTimestampsInner instantiates a new CoreGetUserDatesRequestTimestampsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGetUserDatesRequestTimestampsInner() *CoreGetUserDatesRequestTimestampsInner {
	this := CoreGetUserDatesRequestTimestampsInner{}
	var fixday int32 = 1
	this.Fixday = &fixday
	var fixhour int32 = 1
	this.Fixhour = &fixhour
	var format string = "null"
	this.Format = &format
	var timestamp int32 = null
	this.Timestamp = &timestamp
	var type_ string = "null"
	this.Type = &type_
	return &this
}

// NewCoreGetUserDatesRequestTimestampsInnerWithDefaults instantiates a new CoreGetUserDatesRequestTimestampsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGetUserDatesRequestTimestampsInnerWithDefaults() *CoreGetUserDatesRequestTimestampsInner {
	this := CoreGetUserDatesRequestTimestampsInner{}
	var fixday int32 = 1
	this.Fixday = &fixday
	var fixhour int32 = 1
	this.Fixhour = &fixhour
	var format string = "null"
	this.Format = &format
	var timestamp int32 = null
	this.Timestamp = &timestamp
	var type_ string = "null"
	this.Type = &type_
	return &this
}

// GetFixday returns the Fixday field value if set, zero value otherwise.
func (o *CoreGetUserDatesRequestTimestampsInner) GetFixday() int32 {
	if o == nil || IsNil(o.Fixday) {
		var ret int32
		return ret
	}
	return *o.Fixday
}

// GetFixdayOk returns a tuple with the Fixday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGetUserDatesRequestTimestampsInner) GetFixdayOk() (*int32, bool) {
	if o == nil || IsNil(o.Fixday) {
		return nil, false
	}
	return o.Fixday, true
}

// HasFixday returns a boolean if a field has been set.
func (o *CoreGetUserDatesRequestTimestampsInner) HasFixday() bool {
	if o != nil && !IsNil(o.Fixday) {
		return true
	}

	return false
}

// SetFixday gets a reference to the given int32 and assigns it to the Fixday field.
func (o *CoreGetUserDatesRequestTimestampsInner) SetFixday(v int32) {
	o.Fixday = &v
}

// GetFixhour returns the Fixhour field value if set, zero value otherwise.
func (o *CoreGetUserDatesRequestTimestampsInner) GetFixhour() int32 {
	if o == nil || IsNil(o.Fixhour) {
		var ret int32
		return ret
	}
	return *o.Fixhour
}

// GetFixhourOk returns a tuple with the Fixhour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGetUserDatesRequestTimestampsInner) GetFixhourOk() (*int32, bool) {
	if o == nil || IsNil(o.Fixhour) {
		return nil, false
	}
	return o.Fixhour, true
}

// HasFixhour returns a boolean if a field has been set.
func (o *CoreGetUserDatesRequestTimestampsInner) HasFixhour() bool {
	if o != nil && !IsNil(o.Fixhour) {
		return true
	}

	return false
}

// SetFixhour gets a reference to the given int32 and assigns it to the Fixhour field.
func (o *CoreGetUserDatesRequestTimestampsInner) SetFixhour(v int32) {
	o.Fixhour = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CoreGetUserDatesRequestTimestampsInner) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGetUserDatesRequestTimestampsInner) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CoreGetUserDatesRequestTimestampsInner) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *CoreGetUserDatesRequestTimestampsInner) SetFormat(v string) {
	o.Format = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *CoreGetUserDatesRequestTimestampsInner) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGetUserDatesRequestTimestampsInner) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *CoreGetUserDatesRequestTimestampsInner) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *CoreGetUserDatesRequestTimestampsInner) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CoreGetUserDatesRequestTimestampsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGetUserDatesRequestTimestampsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CoreGetUserDatesRequestTimestampsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CoreGetUserDatesRequestTimestampsInner) SetType(v string) {
	o.Type = &v
}

func (o CoreGetUserDatesRequestTimestampsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGetUserDatesRequestTimestampsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fixday) {
		toSerialize["fixday"] = o.Fixday
	}
	if !IsNil(o.Fixhour) {
		toSerialize["fixhour"] = o.Fixhour
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCoreGetUserDatesRequestTimestampsInner struct {
	value *CoreGetUserDatesRequestTimestampsInner
	isSet bool
}

func (v NullableCoreGetUserDatesRequestTimestampsInner) Get() *CoreGetUserDatesRequestTimestampsInner {
	return v.value
}

func (v *NullableCoreGetUserDatesRequestTimestampsInner) Set(val *CoreGetUserDatesRequestTimestampsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGetUserDatesRequestTimestampsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGetUserDatesRequestTimestampsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGetUserDatesRequestTimestampsInner(val *CoreGetUserDatesRequestTimestampsInner) *NullableCoreGetUserDatesRequestTimestampsInner {
	return &NullableCoreGetUserDatesRequestTimestampsInner{value: val, isSet: true}
}

func (v NullableCoreGetUserDatesRequestTimestampsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGetUserDatesRequestTimestampsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


