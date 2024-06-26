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

// checks if the CoreMessageGetMemberInfoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageGetMemberInfoRequest{}

// CoreMessageGetMemberInfoRequest struct for CoreMessageGetMemberInfoRequest
type CoreMessageGetMemberInfoRequest struct {
	// include contact requests in response
	Includecontactrequests *bool `json:"includecontactrequests,omitempty"`
	// include privacy info in response
	Includeprivacyinfo *bool `json:"includeprivacyinfo,omitempty"`
	// id of the user
	Referenceuserid int32 `json:"referenceuserid"`
	Userids []map[string]interface{} `json:"userids"`
}

type _CoreMessageGetMemberInfoRequest CoreMessageGetMemberInfoRequest

// NewCoreMessageGetMemberInfoRequest instantiates a new CoreMessageGetMemberInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageGetMemberInfoRequest(referenceuserid int32, userids []map[string]interface{}) *CoreMessageGetMemberInfoRequest {
	this := CoreMessageGetMemberInfoRequest{}
	var includecontactrequests bool = false
	this.Includecontactrequests = &includecontactrequests
	var includeprivacyinfo bool = false
	this.Includeprivacyinfo = &includeprivacyinfo
	this.Referenceuserid = referenceuserid
	this.Userids = userids
	return &this
}

// NewCoreMessageGetMemberInfoRequestWithDefaults instantiates a new CoreMessageGetMemberInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageGetMemberInfoRequestWithDefaults() *CoreMessageGetMemberInfoRequest {
	this := CoreMessageGetMemberInfoRequest{}
	var includecontactrequests bool = false
	this.Includecontactrequests = &includecontactrequests
	var includeprivacyinfo bool = false
	this.Includeprivacyinfo = &includeprivacyinfo
	var referenceuserid int32 = null
	this.Referenceuserid = referenceuserid
	return &this
}

// GetIncludecontactrequests returns the Includecontactrequests field value if set, zero value otherwise.
func (o *CoreMessageGetMemberInfoRequest) GetIncludecontactrequests() bool {
	if o == nil || IsNil(o.Includecontactrequests) {
		var ret bool
		return ret
	}
	return *o.Includecontactrequests
}

// GetIncludecontactrequestsOk returns a tuple with the Includecontactrequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetMemberInfoRequest) GetIncludecontactrequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.Includecontactrequests) {
		return nil, false
	}
	return o.Includecontactrequests, true
}

// HasIncludecontactrequests returns a boolean if a field has been set.
func (o *CoreMessageGetMemberInfoRequest) HasIncludecontactrequests() bool {
	if o != nil && !IsNil(o.Includecontactrequests) {
		return true
	}

	return false
}

// SetIncludecontactrequests gets a reference to the given bool and assigns it to the Includecontactrequests field.
func (o *CoreMessageGetMemberInfoRequest) SetIncludecontactrequests(v bool) {
	o.Includecontactrequests = &v
}

// GetIncludeprivacyinfo returns the Includeprivacyinfo field value if set, zero value otherwise.
func (o *CoreMessageGetMemberInfoRequest) GetIncludeprivacyinfo() bool {
	if o == nil || IsNil(o.Includeprivacyinfo) {
		var ret bool
		return ret
	}
	return *o.Includeprivacyinfo
}

// GetIncludeprivacyinfoOk returns a tuple with the Includeprivacyinfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetMemberInfoRequest) GetIncludeprivacyinfoOk() (*bool, bool) {
	if o == nil || IsNil(o.Includeprivacyinfo) {
		return nil, false
	}
	return o.Includeprivacyinfo, true
}

// HasIncludeprivacyinfo returns a boolean if a field has been set.
func (o *CoreMessageGetMemberInfoRequest) HasIncludeprivacyinfo() bool {
	if o != nil && !IsNil(o.Includeprivacyinfo) {
		return true
	}

	return false
}

// SetIncludeprivacyinfo gets a reference to the given bool and assigns it to the Includeprivacyinfo field.
func (o *CoreMessageGetMemberInfoRequest) SetIncludeprivacyinfo(v bool) {
	o.Includeprivacyinfo = &v
}

// GetReferenceuserid returns the Referenceuserid field value
func (o *CoreMessageGetMemberInfoRequest) GetReferenceuserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Referenceuserid
}

// GetReferenceuseridOk returns a tuple with the Referenceuserid field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetMemberInfoRequest) GetReferenceuseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Referenceuserid, true
}

// SetReferenceuserid sets field value
func (o *CoreMessageGetMemberInfoRequest) SetReferenceuserid(v int32) {
	o.Referenceuserid = v
}

// GetUserids returns the Userids field value
func (o *CoreMessageGetMemberInfoRequest) GetUserids() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Userids
}

// GetUseridsOk returns a tuple with the Userids field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetMemberInfoRequest) GetUseridsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Userids, true
}

// SetUserids sets field value
func (o *CoreMessageGetMemberInfoRequest) SetUserids(v []map[string]interface{}) {
	o.Userids = v
}

func (o CoreMessageGetMemberInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageGetMemberInfoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Includecontactrequests) {
		toSerialize["includecontactrequests"] = o.Includecontactrequests
	}
	if !IsNil(o.Includeprivacyinfo) {
		toSerialize["includeprivacyinfo"] = o.Includeprivacyinfo
	}
	toSerialize["referenceuserid"] = o.Referenceuserid
	toSerialize["userids"] = o.Userids
	return toSerialize, nil
}

func (o *CoreMessageGetMemberInfoRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"referenceuserid",
		"userids",
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

	varCoreMessageGetMemberInfoRequest := _CoreMessageGetMemberInfoRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageGetMemberInfoRequest)

	if err != nil {
		return err
	}

	*o = CoreMessageGetMemberInfoRequest(varCoreMessageGetMemberInfoRequest)

	return err
}

type NullableCoreMessageGetMemberInfoRequest struct {
	value *CoreMessageGetMemberInfoRequest
	isSet bool
}

func (v NullableCoreMessageGetMemberInfoRequest) Get() *CoreMessageGetMemberInfoRequest {
	return v.value
}

func (v *NullableCoreMessageGetMemberInfoRequest) Set(val *CoreMessageGetMemberInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageGetMemberInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageGetMemberInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageGetMemberInfoRequest(val *CoreMessageGetMemberInfoRequest) *NullableCoreMessageGetMemberInfoRequest {
	return &NullableCoreMessageGetMemberInfoRequest{value: val, isSet: true}
}

func (v NullableCoreMessageGetMemberInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageGetMemberInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


