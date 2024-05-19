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

// checks if the BlockIomadCompanyAdminRestrictCapabilityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockIomadCompanyAdminRestrictCapabilityRequest{}

// BlockIomadCompanyAdminRestrictCapabilityRequest struct for BlockIomadCompanyAdminRestrictCapabilityRequest
type BlockIomadCompanyAdminRestrictCapabilityRequest struct {
	// Set capability?
	Allow bool `json:"allow"`
	// The capability
	Capability string `json:"capability"`
	// Company ID. Ignored if templateid is non-zero
	Companyid int32 `json:"companyid"`
	// Role ID
	Roleid int32 `json:"roleid"`
	// Template ID. Set to 0 if company restriction
	Templateid *int32 `json:"templateid,omitempty"`
}

type _BlockIomadCompanyAdminRestrictCapabilityRequest BlockIomadCompanyAdminRestrictCapabilityRequest

// NewBlockIomadCompanyAdminRestrictCapabilityRequest instantiates a new BlockIomadCompanyAdminRestrictCapabilityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockIomadCompanyAdminRestrictCapabilityRequest(allow bool, capability string, companyid int32, roleid int32) *BlockIomadCompanyAdminRestrictCapabilityRequest {
	this := BlockIomadCompanyAdminRestrictCapabilityRequest{}
	this.Allow = allow
	this.Capability = capability
	this.Companyid = companyid
	this.Roleid = roleid
	var templateid int32 = 0
	this.Templateid = &templateid
	return &this
}

// NewBlockIomadCompanyAdminRestrictCapabilityRequestWithDefaults instantiates a new BlockIomadCompanyAdminRestrictCapabilityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockIomadCompanyAdminRestrictCapabilityRequestWithDefaults() *BlockIomadCompanyAdminRestrictCapabilityRequest {
	this := BlockIomadCompanyAdminRestrictCapabilityRequest{}
	var allow bool = null
	this.Allow = allow
	var capability string = "null"
	this.Capability = capability
	var companyid int32 = null
	this.Companyid = companyid
	var roleid int32 = null
	this.Roleid = roleid
	var templateid int32 = 0
	this.Templateid = &templateid
	return &this
}

// GetAllow returns the Allow field value
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) GetAllow() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) GetAllowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allow, true
}

// SetAllow sets field value
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) SetAllow(v bool) {
	o.Allow = v
}

// GetCapability returns the Capability field value
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) GetCapability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) GetCapabilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capability, true
}

// SetCapability sets field value
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) SetCapability(v string) {
	o.Capability = v
}

// GetCompanyid returns the Companyid field value
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) GetCompanyid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Companyid
}

// GetCompanyidOk returns a tuple with the Companyid field value
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) GetCompanyidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Companyid, true
}

// SetCompanyid sets field value
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) SetCompanyid(v int32) {
	o.Companyid = v
}

// GetRoleid returns the Roleid field value
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) GetRoleid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Roleid
}

// GetRoleidOk returns a tuple with the Roleid field value
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) GetRoleidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roleid, true
}

// SetRoleid sets field value
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) SetRoleid(v int32) {
	o.Roleid = v
}

// GetTemplateid returns the Templateid field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) GetTemplateid() int32 {
	if o == nil || IsNil(o.Templateid) {
		var ret int32
		return ret
	}
	return *o.Templateid
}

// GetTemplateidOk returns a tuple with the Templateid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) GetTemplateidOk() (*int32, bool) {
	if o == nil || IsNil(o.Templateid) {
		return nil, false
	}
	return o.Templateid, true
}

// HasTemplateid returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) HasTemplateid() bool {
	if o != nil && !IsNil(o.Templateid) {
		return true
	}

	return false
}

// SetTemplateid gets a reference to the given int32 and assigns it to the Templateid field.
func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) SetTemplateid(v int32) {
	o.Templateid = &v
}

func (o BlockIomadCompanyAdminRestrictCapabilityRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockIomadCompanyAdminRestrictCapabilityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allow"] = o.Allow
	toSerialize["capability"] = o.Capability
	toSerialize["companyid"] = o.Companyid
	toSerialize["roleid"] = o.Roleid
	if !IsNil(o.Templateid) {
		toSerialize["templateid"] = o.Templateid
	}
	return toSerialize, nil
}

func (o *BlockIomadCompanyAdminRestrictCapabilityRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allow",
		"capability",
		"companyid",
		"roleid",
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

	varBlockIomadCompanyAdminRestrictCapabilityRequest := _BlockIomadCompanyAdminRestrictCapabilityRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockIomadCompanyAdminRestrictCapabilityRequest)

	if err != nil {
		return err
	}

	*o = BlockIomadCompanyAdminRestrictCapabilityRequest(varBlockIomadCompanyAdminRestrictCapabilityRequest)

	return err
}

type NullableBlockIomadCompanyAdminRestrictCapabilityRequest struct {
	value *BlockIomadCompanyAdminRestrictCapabilityRequest
	isSet bool
}

func (v NullableBlockIomadCompanyAdminRestrictCapabilityRequest) Get() *BlockIomadCompanyAdminRestrictCapabilityRequest {
	return v.value
}

func (v *NullableBlockIomadCompanyAdminRestrictCapabilityRequest) Set(val *BlockIomadCompanyAdminRestrictCapabilityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockIomadCompanyAdminRestrictCapabilityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockIomadCompanyAdminRestrictCapabilityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockIomadCompanyAdminRestrictCapabilityRequest(val *BlockIomadCompanyAdminRestrictCapabilityRequest) *NullableBlockIomadCompanyAdminRestrictCapabilityRequest {
	return &NullableBlockIomadCompanyAdminRestrictCapabilityRequest{value: val, isSet: true}
}

func (v NullableBlockIomadCompanyAdminRestrictCapabilityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockIomadCompanyAdminRestrictCapabilityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

