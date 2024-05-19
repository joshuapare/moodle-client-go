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

// checks if the ModDataGetEntries200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModDataGetEntries200Response{}

// ModDataGetEntries200Response struct for ModDataGetEntries200Response
type ModDataGetEntries200Response struct {
	Entries []ModDataGetEntries200ResponseEntriesInner `json:"entries"`
	// The list view contents as is rendered in the site.
	Listviewcontents *string `json:"listviewcontents,omitempty"`
	// Total count of records.
	Totalcount int32 `json:"totalcount"`
	// Total size (bytes) of the files included in the records.
	Totalfilesize int32 `json:"totalfilesize"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModDataGetEntries200Response ModDataGetEntries200Response

// NewModDataGetEntries200Response instantiates a new ModDataGetEntries200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModDataGetEntries200Response(entries []ModDataGetEntries200ResponseEntriesInner, totalcount int32, totalfilesize int32) *ModDataGetEntries200Response {
	this := ModDataGetEntries200Response{}
	this.Entries = entries
	var listviewcontents string = "null"
	this.Listviewcontents = &listviewcontents
	this.Totalcount = totalcount
	this.Totalfilesize = totalfilesize
	return &this
}

// NewModDataGetEntries200ResponseWithDefaults instantiates a new ModDataGetEntries200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModDataGetEntries200ResponseWithDefaults() *ModDataGetEntries200Response {
	this := ModDataGetEntries200Response{}
	var listviewcontents string = "null"
	this.Listviewcontents = &listviewcontents
	var totalcount int32 = null
	this.Totalcount = totalcount
	var totalfilesize int32 = null
	this.Totalfilesize = totalfilesize
	return &this
}

// GetEntries returns the Entries field value
func (o *ModDataGetEntries200Response) GetEntries() []ModDataGetEntries200ResponseEntriesInner {
	if o == nil {
		var ret []ModDataGetEntries200ResponseEntriesInner
		return ret
	}

	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value
// and a boolean to check if the value has been set.
func (o *ModDataGetEntries200Response) GetEntriesOk() ([]ModDataGetEntries200ResponseEntriesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entries, true
}

// SetEntries sets field value
func (o *ModDataGetEntries200Response) SetEntries(v []ModDataGetEntries200ResponseEntriesInner) {
	o.Entries = v
}

// GetListviewcontents returns the Listviewcontents field value if set, zero value otherwise.
func (o *ModDataGetEntries200Response) GetListviewcontents() string {
	if o == nil || IsNil(o.Listviewcontents) {
		var ret string
		return ret
	}
	return *o.Listviewcontents
}

// GetListviewcontentsOk returns a tuple with the Listviewcontents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetEntries200Response) GetListviewcontentsOk() (*string, bool) {
	if o == nil || IsNil(o.Listviewcontents) {
		return nil, false
	}
	return o.Listviewcontents, true
}

// HasListviewcontents returns a boolean if a field has been set.
func (o *ModDataGetEntries200Response) HasListviewcontents() bool {
	if o != nil && !IsNil(o.Listviewcontents) {
		return true
	}

	return false
}

// SetListviewcontents gets a reference to the given string and assigns it to the Listviewcontents field.
func (o *ModDataGetEntries200Response) SetListviewcontents(v string) {
	o.Listviewcontents = &v
}

// GetTotalcount returns the Totalcount field value
func (o *ModDataGetEntries200Response) GetTotalcount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Totalcount
}

// GetTotalcountOk returns a tuple with the Totalcount field value
// and a boolean to check if the value has been set.
func (o *ModDataGetEntries200Response) GetTotalcountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Totalcount, true
}

// SetTotalcount sets field value
func (o *ModDataGetEntries200Response) SetTotalcount(v int32) {
	o.Totalcount = v
}

// GetTotalfilesize returns the Totalfilesize field value
func (o *ModDataGetEntries200Response) GetTotalfilesize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Totalfilesize
}

// GetTotalfilesizeOk returns a tuple with the Totalfilesize field value
// and a boolean to check if the value has been set.
func (o *ModDataGetEntries200Response) GetTotalfilesizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Totalfilesize, true
}

// SetTotalfilesize sets field value
func (o *ModDataGetEntries200Response) SetTotalfilesize(v int32) {
	o.Totalfilesize = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModDataGetEntries200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetEntries200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModDataGetEntries200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModDataGetEntries200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModDataGetEntries200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModDataGetEntries200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entries"] = o.Entries
	if !IsNil(o.Listviewcontents) {
		toSerialize["listviewcontents"] = o.Listviewcontents
	}
	toSerialize["totalcount"] = o.Totalcount
	toSerialize["totalfilesize"] = o.Totalfilesize
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModDataGetEntries200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entries",
		"totalcount",
		"totalfilesize",
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

	varModDataGetEntries200Response := _ModDataGetEntries200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModDataGetEntries200Response)

	if err != nil {
		return err
	}

	*o = ModDataGetEntries200Response(varModDataGetEntries200Response)

	return err
}

type NullableModDataGetEntries200Response struct {
	value *ModDataGetEntries200Response
	isSet bool
}

func (v NullableModDataGetEntries200Response) Get() *ModDataGetEntries200Response {
	return v.value
}

func (v *NullableModDataGetEntries200Response) Set(val *ModDataGetEntries200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModDataGetEntries200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModDataGetEntries200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModDataGetEntries200Response(val *ModDataGetEntries200Response) *NullableModDataGetEntries200Response {
	return &NullableModDataGetEntries200Response{value: val, isSet: true}
}

func (v NullableModDataGetEntries200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModDataGetEntries200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

