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

// checks if the CoreUserSearchIdentity200ResponseListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreUserSearchIdentity200ResponseListInner{}

// CoreUserSearchIdentity200ResponseListInner struct for CoreUserSearchIdentity200ResponseListInner
type CoreUserSearchIdentity200ResponseListInner struct {
	Extrafields []CoreUserSearchIdentity200ResponseListInnerExtrafieldsInner `json:"extrafields,omitempty"`
	// The fullname of the user
	Fullname *string `json:"fullname,omitempty"`
	// ID of the user
	Id *int32 `json:"id,omitempty"`
}

// NewCoreUserSearchIdentity200ResponseListInner instantiates a new CoreUserSearchIdentity200ResponseListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUserSearchIdentity200ResponseListInner() *CoreUserSearchIdentity200ResponseListInner {
	this := CoreUserSearchIdentity200ResponseListInner{}
	return &this
}

// NewCoreUserSearchIdentity200ResponseListInnerWithDefaults instantiates a new CoreUserSearchIdentity200ResponseListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUserSearchIdentity200ResponseListInnerWithDefaults() *CoreUserSearchIdentity200ResponseListInner {
	this := CoreUserSearchIdentity200ResponseListInner{}
	return &this
}

// GetExtrafields returns the Extrafields field value if set, zero value otherwise.
func (o *CoreUserSearchIdentity200ResponseListInner) GetExtrafields() []CoreUserSearchIdentity200ResponseListInnerExtrafieldsInner {
	if o == nil || IsNil(o.Extrafields) {
		var ret []CoreUserSearchIdentity200ResponseListInnerExtrafieldsInner
		return ret
	}
	return o.Extrafields
}

// GetExtrafieldsOk returns a tuple with the Extrafields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUserSearchIdentity200ResponseListInner) GetExtrafieldsOk() ([]CoreUserSearchIdentity200ResponseListInnerExtrafieldsInner, bool) {
	if o == nil || IsNil(o.Extrafields) {
		return nil, false
	}
	return o.Extrafields, true
}

// HasExtrafields returns a boolean if a field has been set.
func (o *CoreUserSearchIdentity200ResponseListInner) HasExtrafields() bool {
	if o != nil && !IsNil(o.Extrafields) {
		return true
	}

	return false
}

// SetExtrafields gets a reference to the given []CoreUserSearchIdentity200ResponseListInnerExtrafieldsInner and assigns it to the Extrafields field.
func (o *CoreUserSearchIdentity200ResponseListInner) SetExtrafields(v []CoreUserSearchIdentity200ResponseListInnerExtrafieldsInner) {
	o.Extrafields = v
}

// GetFullname returns the Fullname field value if set, zero value otherwise.
func (o *CoreUserSearchIdentity200ResponseListInner) GetFullname() string {
	if o == nil || IsNil(o.Fullname) {
		var ret string
		return ret
	}
	return *o.Fullname
}

// GetFullnameOk returns a tuple with the Fullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUserSearchIdentity200ResponseListInner) GetFullnameOk() (*string, bool) {
	if o == nil || IsNil(o.Fullname) {
		return nil, false
	}
	return o.Fullname, true
}

// HasFullname returns a boolean if a field has been set.
func (o *CoreUserSearchIdentity200ResponseListInner) HasFullname() bool {
	if o != nil && !IsNil(o.Fullname) {
		return true
	}

	return false
}

// SetFullname gets a reference to the given string and assigns it to the Fullname field.
func (o *CoreUserSearchIdentity200ResponseListInner) SetFullname(v string) {
	o.Fullname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreUserSearchIdentity200ResponseListInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUserSearchIdentity200ResponseListInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreUserSearchIdentity200ResponseListInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreUserSearchIdentity200ResponseListInner) SetId(v int32) {
	o.Id = &v
}

func (o CoreUserSearchIdentity200ResponseListInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreUserSearchIdentity200ResponseListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Extrafields) {
		toSerialize["extrafields"] = o.Extrafields
	}
	if !IsNil(o.Fullname) {
		toSerialize["fullname"] = o.Fullname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCoreUserSearchIdentity200ResponseListInner struct {
	value *CoreUserSearchIdentity200ResponseListInner
	isSet bool
}

func (v NullableCoreUserSearchIdentity200ResponseListInner) Get() *CoreUserSearchIdentity200ResponseListInner {
	return v.value
}

func (v *NullableCoreUserSearchIdentity200ResponseListInner) Set(val *CoreUserSearchIdentity200ResponseListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUserSearchIdentity200ResponseListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUserSearchIdentity200ResponseListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUserSearchIdentity200ResponseListInner(val *CoreUserSearchIdentity200ResponseListInner) *NullableCoreUserSearchIdentity200ResponseListInner {
	return &NullableCoreUserSearchIdentity200ResponseListInner{value: val, isSet: true}
}

func (v NullableCoreUserSearchIdentity200ResponseListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUserSearchIdentity200ResponseListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


