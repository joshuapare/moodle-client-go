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

// checks if the ModDataGetEntry200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModDataGetEntry200Response{}

// ModDataGetEntry200Response struct for ModDataGetEntry200Response
type ModDataGetEntry200Response struct {
	Entry ModDataGetEntry200ResponseEntry `json:"entry"`
	// The entry as is rendered in the site.
	Entryviewcontents *string `json:"entryviewcontents,omitempty"`
	Ratinginfo *ModDataGetEntry200ResponseRatinginfo `json:"ratinginfo,omitempty"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModDataGetEntry200Response ModDataGetEntry200Response

// NewModDataGetEntry200Response instantiates a new ModDataGetEntry200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModDataGetEntry200Response(entry ModDataGetEntry200ResponseEntry) *ModDataGetEntry200Response {
	this := ModDataGetEntry200Response{}
	this.Entry = entry
	var entryviewcontents string = "null"
	this.Entryviewcontents = &entryviewcontents
	return &this
}

// NewModDataGetEntry200ResponseWithDefaults instantiates a new ModDataGetEntry200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModDataGetEntry200ResponseWithDefaults() *ModDataGetEntry200Response {
	this := ModDataGetEntry200Response{}
	var entryviewcontents string = "null"
	this.Entryviewcontents = &entryviewcontents
	return &this
}

// GetEntry returns the Entry field value
func (o *ModDataGetEntry200Response) GetEntry() ModDataGetEntry200ResponseEntry {
	if o == nil {
		var ret ModDataGetEntry200ResponseEntry
		return ret
	}

	return o.Entry
}

// GetEntryOk returns a tuple with the Entry field value
// and a boolean to check if the value has been set.
func (o *ModDataGetEntry200Response) GetEntryOk() (*ModDataGetEntry200ResponseEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entry, true
}

// SetEntry sets field value
func (o *ModDataGetEntry200Response) SetEntry(v ModDataGetEntry200ResponseEntry) {
	o.Entry = v
}

// GetEntryviewcontents returns the Entryviewcontents field value if set, zero value otherwise.
func (o *ModDataGetEntry200Response) GetEntryviewcontents() string {
	if o == nil || IsNil(o.Entryviewcontents) {
		var ret string
		return ret
	}
	return *o.Entryviewcontents
}

// GetEntryviewcontentsOk returns a tuple with the Entryviewcontents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetEntry200Response) GetEntryviewcontentsOk() (*string, bool) {
	if o == nil || IsNil(o.Entryviewcontents) {
		return nil, false
	}
	return o.Entryviewcontents, true
}

// HasEntryviewcontents returns a boolean if a field has been set.
func (o *ModDataGetEntry200Response) HasEntryviewcontents() bool {
	if o != nil && !IsNil(o.Entryviewcontents) {
		return true
	}

	return false
}

// SetEntryviewcontents gets a reference to the given string and assigns it to the Entryviewcontents field.
func (o *ModDataGetEntry200Response) SetEntryviewcontents(v string) {
	o.Entryviewcontents = &v
}

// GetRatinginfo returns the Ratinginfo field value if set, zero value otherwise.
func (o *ModDataGetEntry200Response) GetRatinginfo() ModDataGetEntry200ResponseRatinginfo {
	if o == nil || IsNil(o.Ratinginfo) {
		var ret ModDataGetEntry200ResponseRatinginfo
		return ret
	}
	return *o.Ratinginfo
}

// GetRatinginfoOk returns a tuple with the Ratinginfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetEntry200Response) GetRatinginfoOk() (*ModDataGetEntry200ResponseRatinginfo, bool) {
	if o == nil || IsNil(o.Ratinginfo) {
		return nil, false
	}
	return o.Ratinginfo, true
}

// HasRatinginfo returns a boolean if a field has been set.
func (o *ModDataGetEntry200Response) HasRatinginfo() bool {
	if o != nil && !IsNil(o.Ratinginfo) {
		return true
	}

	return false
}

// SetRatinginfo gets a reference to the given ModDataGetEntry200ResponseRatinginfo and assigns it to the Ratinginfo field.
func (o *ModDataGetEntry200Response) SetRatinginfo(v ModDataGetEntry200ResponseRatinginfo) {
	o.Ratinginfo = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModDataGetEntry200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetEntry200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModDataGetEntry200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModDataGetEntry200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModDataGetEntry200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModDataGetEntry200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entry"] = o.Entry
	if !IsNil(o.Entryviewcontents) {
		toSerialize["entryviewcontents"] = o.Entryviewcontents
	}
	if !IsNil(o.Ratinginfo) {
		toSerialize["ratinginfo"] = o.Ratinginfo
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModDataGetEntry200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entry",
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

	varModDataGetEntry200Response := _ModDataGetEntry200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModDataGetEntry200Response)

	if err != nil {
		return err
	}

	*o = ModDataGetEntry200Response(varModDataGetEntry200Response)

	return err
}

type NullableModDataGetEntry200Response struct {
	value *ModDataGetEntry200Response
	isSet bool
}

func (v NullableModDataGetEntry200Response) Get() *ModDataGetEntry200Response {
	return v.value
}

func (v *NullableModDataGetEntry200Response) Set(val *ModDataGetEntry200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModDataGetEntry200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModDataGetEntry200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModDataGetEntry200Response(val *ModDataGetEntry200Response) *NullableModDataGetEntry200Response {
	return &NullableModDataGetEntry200Response{value: val, isSet: true}
}

func (v NullableModDataGetEntry200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModDataGetEntry200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

