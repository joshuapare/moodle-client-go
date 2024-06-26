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

// checks if the TinyAutosaveUpdateSessionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TinyAutosaveUpdateSessionRequest{}

// TinyAutosaveUpdateSessionRequest struct for TinyAutosaveUpdateSessionRequest
type TinyAutosaveUpdateSessionRequest struct {
	// The context id that owns the editor
	Contextid int32 `json:"contextid"`
	// The draft text
	Drafttext string `json:"drafttext"`
	// The ID of the element
	Elementid string `json:"elementid"`
	// The page hash
	Pagehash string `json:"pagehash"`
	// The page instance
	Pageinstance string `json:"pageinstance"`
}

type _TinyAutosaveUpdateSessionRequest TinyAutosaveUpdateSessionRequest

// NewTinyAutosaveUpdateSessionRequest instantiates a new TinyAutosaveUpdateSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTinyAutosaveUpdateSessionRequest(contextid int32, drafttext string, elementid string, pagehash string, pageinstance string) *TinyAutosaveUpdateSessionRequest {
	this := TinyAutosaveUpdateSessionRequest{}
	this.Contextid = contextid
	this.Drafttext = drafttext
	this.Elementid = elementid
	this.Pagehash = pagehash
	this.Pageinstance = pageinstance
	return &this
}

// NewTinyAutosaveUpdateSessionRequestWithDefaults instantiates a new TinyAutosaveUpdateSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTinyAutosaveUpdateSessionRequestWithDefaults() *TinyAutosaveUpdateSessionRequest {
	this := TinyAutosaveUpdateSessionRequest{}
	return &this
}

// GetContextid returns the Contextid field value
func (o *TinyAutosaveUpdateSessionRequest) GetContextid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Contextid
}

// GetContextidOk returns a tuple with the Contextid field value
// and a boolean to check if the value has been set.
func (o *TinyAutosaveUpdateSessionRequest) GetContextidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contextid, true
}

// SetContextid sets field value
func (o *TinyAutosaveUpdateSessionRequest) SetContextid(v int32) {
	o.Contextid = v
}

// GetDrafttext returns the Drafttext field value
func (o *TinyAutosaveUpdateSessionRequest) GetDrafttext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Drafttext
}

// GetDrafttextOk returns a tuple with the Drafttext field value
// and a boolean to check if the value has been set.
func (o *TinyAutosaveUpdateSessionRequest) GetDrafttextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Drafttext, true
}

// SetDrafttext sets field value
func (o *TinyAutosaveUpdateSessionRequest) SetDrafttext(v string) {
	o.Drafttext = v
}

// GetElementid returns the Elementid field value
func (o *TinyAutosaveUpdateSessionRequest) GetElementid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Elementid
}

// GetElementidOk returns a tuple with the Elementid field value
// and a boolean to check if the value has been set.
func (o *TinyAutosaveUpdateSessionRequest) GetElementidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Elementid, true
}

// SetElementid sets field value
func (o *TinyAutosaveUpdateSessionRequest) SetElementid(v string) {
	o.Elementid = v
}

// GetPagehash returns the Pagehash field value
func (o *TinyAutosaveUpdateSessionRequest) GetPagehash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pagehash
}

// GetPagehashOk returns a tuple with the Pagehash field value
// and a boolean to check if the value has been set.
func (o *TinyAutosaveUpdateSessionRequest) GetPagehashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagehash, true
}

// SetPagehash sets field value
func (o *TinyAutosaveUpdateSessionRequest) SetPagehash(v string) {
	o.Pagehash = v
}

// GetPageinstance returns the Pageinstance field value
func (o *TinyAutosaveUpdateSessionRequest) GetPageinstance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pageinstance
}

// GetPageinstanceOk returns a tuple with the Pageinstance field value
// and a boolean to check if the value has been set.
func (o *TinyAutosaveUpdateSessionRequest) GetPageinstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pageinstance, true
}

// SetPageinstance sets field value
func (o *TinyAutosaveUpdateSessionRequest) SetPageinstance(v string) {
	o.Pageinstance = v
}

func (o TinyAutosaveUpdateSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TinyAutosaveUpdateSessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contextid"] = o.Contextid
	toSerialize["drafttext"] = o.Drafttext
	toSerialize["elementid"] = o.Elementid
	toSerialize["pagehash"] = o.Pagehash
	toSerialize["pageinstance"] = o.Pageinstance
	return toSerialize, nil
}

func (o *TinyAutosaveUpdateSessionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contextid",
		"drafttext",
		"elementid",
		"pagehash",
		"pageinstance",
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

	varTinyAutosaveUpdateSessionRequest := _TinyAutosaveUpdateSessionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTinyAutosaveUpdateSessionRequest)

	if err != nil {
		return err
	}

	*o = TinyAutosaveUpdateSessionRequest(varTinyAutosaveUpdateSessionRequest)

	return err
}

type NullableTinyAutosaveUpdateSessionRequest struct {
	value *TinyAutosaveUpdateSessionRequest
	isSet bool
}

func (v NullableTinyAutosaveUpdateSessionRequest) Get() *TinyAutosaveUpdateSessionRequest {
	return v.value
}

func (v *NullableTinyAutosaveUpdateSessionRequest) Set(val *TinyAutosaveUpdateSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTinyAutosaveUpdateSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTinyAutosaveUpdateSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTinyAutosaveUpdateSessionRequest(val *TinyAutosaveUpdateSessionRequest) *NullableTinyAutosaveUpdateSessionRequest {
	return &NullableTinyAutosaveUpdateSessionRequest{value: val, isSet: true}
}

func (v NullableTinyAutosaveUpdateSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTinyAutosaveUpdateSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


