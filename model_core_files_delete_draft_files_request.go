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

// checks if the CoreFilesDeleteDraftFilesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreFilesDeleteDraftFilesRequest{}

// CoreFilesDeleteDraftFilesRequest struct for CoreFilesDeleteDraftFilesRequest
type CoreFilesDeleteDraftFilesRequest struct {
	// Item id of the draft file area
	Draftitemid int32 `json:"draftitemid"`
	Files []CoreFilesDeleteDraftFilesRequestFilesInner `json:"files"`
}

type _CoreFilesDeleteDraftFilesRequest CoreFilesDeleteDraftFilesRequest

// NewCoreFilesDeleteDraftFilesRequest instantiates a new CoreFilesDeleteDraftFilesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreFilesDeleteDraftFilesRequest(draftitemid int32, files []CoreFilesDeleteDraftFilesRequestFilesInner) *CoreFilesDeleteDraftFilesRequest {
	this := CoreFilesDeleteDraftFilesRequest{}
	this.Draftitemid = draftitemid
	this.Files = files
	return &this
}

// NewCoreFilesDeleteDraftFilesRequestWithDefaults instantiates a new CoreFilesDeleteDraftFilesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreFilesDeleteDraftFilesRequestWithDefaults() *CoreFilesDeleteDraftFilesRequest {
	this := CoreFilesDeleteDraftFilesRequest{}
	var draftitemid int32 = null
	this.Draftitemid = draftitemid
	return &this
}

// GetDraftitemid returns the Draftitemid field value
func (o *CoreFilesDeleteDraftFilesRequest) GetDraftitemid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Draftitemid
}

// GetDraftitemidOk returns a tuple with the Draftitemid field value
// and a boolean to check if the value has been set.
func (o *CoreFilesDeleteDraftFilesRequest) GetDraftitemidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Draftitemid, true
}

// SetDraftitemid sets field value
func (o *CoreFilesDeleteDraftFilesRequest) SetDraftitemid(v int32) {
	o.Draftitemid = v
}

// GetFiles returns the Files field value
func (o *CoreFilesDeleteDraftFilesRequest) GetFiles() []CoreFilesDeleteDraftFilesRequestFilesInner {
	if o == nil {
		var ret []CoreFilesDeleteDraftFilesRequestFilesInner
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *CoreFilesDeleteDraftFilesRequest) GetFilesOk() ([]CoreFilesDeleteDraftFilesRequestFilesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *CoreFilesDeleteDraftFilesRequest) SetFiles(v []CoreFilesDeleteDraftFilesRequestFilesInner) {
	o.Files = v
}

func (o CoreFilesDeleteDraftFilesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreFilesDeleteDraftFilesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["draftitemid"] = o.Draftitemid
	toSerialize["files"] = o.Files
	return toSerialize, nil
}

func (o *CoreFilesDeleteDraftFilesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"draftitemid",
		"files",
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

	varCoreFilesDeleteDraftFilesRequest := _CoreFilesDeleteDraftFilesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreFilesDeleteDraftFilesRequest)

	if err != nil {
		return err
	}

	*o = CoreFilesDeleteDraftFilesRequest(varCoreFilesDeleteDraftFilesRequest)

	return err
}

type NullableCoreFilesDeleteDraftFilesRequest struct {
	value *CoreFilesDeleteDraftFilesRequest
	isSet bool
}

func (v NullableCoreFilesDeleteDraftFilesRequest) Get() *CoreFilesDeleteDraftFilesRequest {
	return v.value
}

func (v *NullableCoreFilesDeleteDraftFilesRequest) Set(val *CoreFilesDeleteDraftFilesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreFilesDeleteDraftFilesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreFilesDeleteDraftFilesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreFilesDeleteDraftFilesRequest(val *CoreFilesDeleteDraftFilesRequest) *NullableCoreFilesDeleteDraftFilesRequest {
	return &NullableCoreFilesDeleteDraftFilesRequest{value: val, isSet: true}
}

func (v NullableCoreFilesDeleteDraftFilesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreFilesDeleteDraftFilesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

