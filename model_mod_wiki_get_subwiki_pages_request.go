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

// checks if the ModWikiGetSubwikiPagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModWikiGetSubwikiPagesRequest{}

// ModWikiGetSubwikiPagesRequest struct for ModWikiGetSubwikiPagesRequest
type ModWikiGetSubwikiPagesRequest struct {
	// Subwiki's group ID, -1 means current group. It will be ignored if the wiki doesn't use groups.
	Groupid *int32 `json:"groupid,omitempty"`
	Options *ModWikiGetSubwikiPagesRequestOptions `json:"options,omitempty"`
	// Subwiki's user ID, 0 means current user. It will be ignored in collaborative wikis.
	Userid *int32 `json:"userid,omitempty"`
	// Wiki instance ID.
	Wikiid int32 `json:"wikiid"`
}

type _ModWikiGetSubwikiPagesRequest ModWikiGetSubwikiPagesRequest

// NewModWikiGetSubwikiPagesRequest instantiates a new ModWikiGetSubwikiPagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModWikiGetSubwikiPagesRequest(wikiid int32) *ModWikiGetSubwikiPagesRequest {
	this := ModWikiGetSubwikiPagesRequest{}
	var groupid int32 = -1
	this.Groupid = &groupid
	var userid int32 = 0
	this.Userid = &userid
	this.Wikiid = wikiid
	return &this
}

// NewModWikiGetSubwikiPagesRequestWithDefaults instantiates a new ModWikiGetSubwikiPagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModWikiGetSubwikiPagesRequestWithDefaults() *ModWikiGetSubwikiPagesRequest {
	this := ModWikiGetSubwikiPagesRequest{}
	var groupid int32 = -1
	this.Groupid = &groupid
	var userid int32 = 0
	this.Userid = &userid
	return &this
}

// GetGroupid returns the Groupid field value if set, zero value otherwise.
func (o *ModWikiGetSubwikiPagesRequest) GetGroupid() int32 {
	if o == nil || IsNil(o.Groupid) {
		var ret int32
		return ret
	}
	return *o.Groupid
}

// GetGroupidOk returns a tuple with the Groupid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWikiGetSubwikiPagesRequest) GetGroupidOk() (*int32, bool) {
	if o == nil || IsNil(o.Groupid) {
		return nil, false
	}
	return o.Groupid, true
}

// HasGroupid returns a boolean if a field has been set.
func (o *ModWikiGetSubwikiPagesRequest) HasGroupid() bool {
	if o != nil && !IsNil(o.Groupid) {
		return true
	}

	return false
}

// SetGroupid gets a reference to the given int32 and assigns it to the Groupid field.
func (o *ModWikiGetSubwikiPagesRequest) SetGroupid(v int32) {
	o.Groupid = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ModWikiGetSubwikiPagesRequest) GetOptions() ModWikiGetSubwikiPagesRequestOptions {
	if o == nil || IsNil(o.Options) {
		var ret ModWikiGetSubwikiPagesRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWikiGetSubwikiPagesRequest) GetOptionsOk() (*ModWikiGetSubwikiPagesRequestOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ModWikiGetSubwikiPagesRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ModWikiGetSubwikiPagesRequestOptions and assigns it to the Options field.
func (o *ModWikiGetSubwikiPagesRequest) SetOptions(v ModWikiGetSubwikiPagesRequestOptions) {
	o.Options = &v
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *ModWikiGetSubwikiPagesRequest) GetUserid() int32 {
	if o == nil || IsNil(o.Userid) {
		var ret int32
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWikiGetSubwikiPagesRequest) GetUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *ModWikiGetSubwikiPagesRequest) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given int32 and assigns it to the Userid field.
func (o *ModWikiGetSubwikiPagesRequest) SetUserid(v int32) {
	o.Userid = &v
}

// GetWikiid returns the Wikiid field value
func (o *ModWikiGetSubwikiPagesRequest) GetWikiid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Wikiid
}

// GetWikiidOk returns a tuple with the Wikiid field value
// and a boolean to check if the value has been set.
func (o *ModWikiGetSubwikiPagesRequest) GetWikiidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wikiid, true
}

// SetWikiid sets field value
func (o *ModWikiGetSubwikiPagesRequest) SetWikiid(v int32) {
	o.Wikiid = v
}

func (o ModWikiGetSubwikiPagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModWikiGetSubwikiPagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Groupid) {
		toSerialize["groupid"] = o.Groupid
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	toSerialize["wikiid"] = o.Wikiid
	return toSerialize, nil
}

func (o *ModWikiGetSubwikiPagesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wikiid",
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

	varModWikiGetSubwikiPagesRequest := _ModWikiGetSubwikiPagesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModWikiGetSubwikiPagesRequest)

	if err != nil {
		return err
	}

	*o = ModWikiGetSubwikiPagesRequest(varModWikiGetSubwikiPagesRequest)

	return err
}

type NullableModWikiGetSubwikiPagesRequest struct {
	value *ModWikiGetSubwikiPagesRequest
	isSet bool
}

func (v NullableModWikiGetSubwikiPagesRequest) Get() *ModWikiGetSubwikiPagesRequest {
	return v.value
}

func (v *NullableModWikiGetSubwikiPagesRequest) Set(val *ModWikiGetSubwikiPagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModWikiGetSubwikiPagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModWikiGetSubwikiPagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModWikiGetSubwikiPagesRequest(val *ModWikiGetSubwikiPagesRequest) *NullableModWikiGetSubwikiPagesRequest {
	return &NullableModWikiGetSubwikiPagesRequest{value: val, isSet: true}
}

func (v NullableModWikiGetSubwikiPagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModWikiGetSubwikiPagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


