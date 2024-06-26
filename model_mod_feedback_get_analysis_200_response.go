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

// checks if the ModFeedbackGetAnalysis200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModFeedbackGetAnalysis200Response{}

// ModFeedbackGetAnalysis200Response struct for ModFeedbackGetAnalysis200Response
type ModFeedbackGetAnalysis200Response struct {
	// Number of completed submissions.
	Completedcount int32 `json:"completedcount"`
	// Number of items (questions).
	Itemscount int32 `json:"itemscount"`
	Itemsdata []ModFeedbackGetAnalysis200ResponseItemsdataInner `json:"itemsdata"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModFeedbackGetAnalysis200Response ModFeedbackGetAnalysis200Response

// NewModFeedbackGetAnalysis200Response instantiates a new ModFeedbackGetAnalysis200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModFeedbackGetAnalysis200Response(completedcount int32, itemscount int32, itemsdata []ModFeedbackGetAnalysis200ResponseItemsdataInner) *ModFeedbackGetAnalysis200Response {
	this := ModFeedbackGetAnalysis200Response{}
	this.Completedcount = completedcount
	this.Itemscount = itemscount
	this.Itemsdata = itemsdata
	return &this
}

// NewModFeedbackGetAnalysis200ResponseWithDefaults instantiates a new ModFeedbackGetAnalysis200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModFeedbackGetAnalysis200ResponseWithDefaults() *ModFeedbackGetAnalysis200Response {
	this := ModFeedbackGetAnalysis200Response{}
	var completedcount int32 = null
	this.Completedcount = completedcount
	var itemscount int32 = null
	this.Itemscount = itemscount
	return &this
}

// GetCompletedcount returns the Completedcount field value
func (o *ModFeedbackGetAnalysis200Response) GetCompletedcount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Completedcount
}

// GetCompletedcountOk returns a tuple with the Completedcount field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200Response) GetCompletedcountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completedcount, true
}

// SetCompletedcount sets field value
func (o *ModFeedbackGetAnalysis200Response) SetCompletedcount(v int32) {
	o.Completedcount = v
}

// GetItemscount returns the Itemscount field value
func (o *ModFeedbackGetAnalysis200Response) GetItemscount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Itemscount
}

// GetItemscountOk returns a tuple with the Itemscount field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200Response) GetItemscountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Itemscount, true
}

// SetItemscount sets field value
func (o *ModFeedbackGetAnalysis200Response) SetItemscount(v int32) {
	o.Itemscount = v
}

// GetItemsdata returns the Itemsdata field value
func (o *ModFeedbackGetAnalysis200Response) GetItemsdata() []ModFeedbackGetAnalysis200ResponseItemsdataInner {
	if o == nil {
		var ret []ModFeedbackGetAnalysis200ResponseItemsdataInner
		return ret
	}

	return o.Itemsdata
}

// GetItemsdataOk returns a tuple with the Itemsdata field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200Response) GetItemsdataOk() ([]ModFeedbackGetAnalysis200ResponseItemsdataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Itemsdata, true
}

// SetItemsdata sets field value
func (o *ModFeedbackGetAnalysis200Response) SetItemsdata(v []ModFeedbackGetAnalysis200ResponseItemsdataInner) {
	o.Itemsdata = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModFeedbackGetAnalysis200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModFeedbackGetAnalysis200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModFeedbackGetAnalysis200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModFeedbackGetAnalysis200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModFeedbackGetAnalysis200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["completedcount"] = o.Completedcount
	toSerialize["itemscount"] = o.Itemscount
	toSerialize["itemsdata"] = o.Itemsdata
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModFeedbackGetAnalysis200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"completedcount",
		"itemscount",
		"itemsdata",
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

	varModFeedbackGetAnalysis200Response := _ModFeedbackGetAnalysis200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModFeedbackGetAnalysis200Response)

	if err != nil {
		return err
	}

	*o = ModFeedbackGetAnalysis200Response(varModFeedbackGetAnalysis200Response)

	return err
}

type NullableModFeedbackGetAnalysis200Response struct {
	value *ModFeedbackGetAnalysis200Response
	isSet bool
}

func (v NullableModFeedbackGetAnalysis200Response) Get() *ModFeedbackGetAnalysis200Response {
	return v.value
}

func (v *NullableModFeedbackGetAnalysis200Response) Set(val *ModFeedbackGetAnalysis200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModFeedbackGetAnalysis200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModFeedbackGetAnalysis200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModFeedbackGetAnalysis200Response(val *ModFeedbackGetAnalysis200Response) *NullableModFeedbackGetAnalysis200Response {
	return &NullableModFeedbackGetAnalysis200Response{value: val, isSet: true}
}

func (v NullableModFeedbackGetAnalysis200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModFeedbackGetAnalysis200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


