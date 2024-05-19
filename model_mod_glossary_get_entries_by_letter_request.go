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

// checks if the ModGlossaryGetEntriesByLetterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModGlossaryGetEntriesByLetterRequest{}

// ModGlossaryGetEntriesByLetterRequest struct for ModGlossaryGetEntriesByLetterRequest
type ModGlossaryGetEntriesByLetterRequest struct {
	// Start returning records from here
	From *int32 `json:"from,omitempty"`
	// Glossary entry ID
	Id int32 `json:"id"`
	// A letter, or either keywords: 'ALL' or 'SPECIAL'.
	Letter string `json:"letter"`
	// Number of records to return
	Limit *int32 `json:"limit,omitempty"`
	Options *ModGlossaryGetEntriesByAuthorRequestOptions `json:"options,omitempty"`
}

type _ModGlossaryGetEntriesByLetterRequest ModGlossaryGetEntriesByLetterRequest

// NewModGlossaryGetEntriesByLetterRequest instantiates a new ModGlossaryGetEntriesByLetterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModGlossaryGetEntriesByLetterRequest(id int32, letter string) *ModGlossaryGetEntriesByLetterRequest {
	this := ModGlossaryGetEntriesByLetterRequest{}
	var from int32 = 0
	this.From = &from
	this.Id = id
	this.Letter = letter
	var limit int32 = 20
	this.Limit = &limit
	return &this
}

// NewModGlossaryGetEntriesByLetterRequestWithDefaults instantiates a new ModGlossaryGetEntriesByLetterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModGlossaryGetEntriesByLetterRequestWithDefaults() *ModGlossaryGetEntriesByLetterRequest {
	this := ModGlossaryGetEntriesByLetterRequest{}
	var from int32 = 0
	this.From = &from
	var letter string = "null"
	this.Letter = letter
	var limit int32 = 20
	this.Limit = &limit
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByLetterRequest) GetFrom() int32 {
	if o == nil || IsNil(o.From) {
		var ret int32
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByLetterRequest) GetFromOk() (*int32, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByLetterRequest) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int32 and assigns it to the From field.
func (o *ModGlossaryGetEntriesByLetterRequest) SetFrom(v int32) {
	o.From = &v
}

// GetId returns the Id field value
func (o *ModGlossaryGetEntriesByLetterRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByLetterRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModGlossaryGetEntriesByLetterRequest) SetId(v int32) {
	o.Id = v
}

// GetLetter returns the Letter field value
func (o *ModGlossaryGetEntriesByLetterRequest) GetLetter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Letter
}

// GetLetterOk returns a tuple with the Letter field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByLetterRequest) GetLetterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Letter, true
}

// SetLetter sets field value
func (o *ModGlossaryGetEntriesByLetterRequest) SetLetter(v string) {
	o.Letter = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByLetterRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByLetterRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByLetterRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ModGlossaryGetEntriesByLetterRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByLetterRequest) GetOptions() ModGlossaryGetEntriesByAuthorRequestOptions {
	if o == nil || IsNil(o.Options) {
		var ret ModGlossaryGetEntriesByAuthorRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByLetterRequest) GetOptionsOk() (*ModGlossaryGetEntriesByAuthorRequestOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByLetterRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ModGlossaryGetEntriesByAuthorRequestOptions and assigns it to the Options field.
func (o *ModGlossaryGetEntriesByLetterRequest) SetOptions(v ModGlossaryGetEntriesByAuthorRequestOptions) {
	o.Options = &v
}

func (o ModGlossaryGetEntriesByLetterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModGlossaryGetEntriesByLetterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	toSerialize["id"] = o.Id
	toSerialize["letter"] = o.Letter
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

func (o *ModGlossaryGetEntriesByLetterRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"letter",
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

	varModGlossaryGetEntriesByLetterRequest := _ModGlossaryGetEntriesByLetterRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModGlossaryGetEntriesByLetterRequest)

	if err != nil {
		return err
	}

	*o = ModGlossaryGetEntriesByLetterRequest(varModGlossaryGetEntriesByLetterRequest)

	return err
}

type NullableModGlossaryGetEntriesByLetterRequest struct {
	value *ModGlossaryGetEntriesByLetterRequest
	isSet bool
}

func (v NullableModGlossaryGetEntriesByLetterRequest) Get() *ModGlossaryGetEntriesByLetterRequest {
	return v.value
}

func (v *NullableModGlossaryGetEntriesByLetterRequest) Set(val *ModGlossaryGetEntriesByLetterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModGlossaryGetEntriesByLetterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModGlossaryGetEntriesByLetterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModGlossaryGetEntriesByLetterRequest(val *ModGlossaryGetEntriesByLetterRequest) *NullableModGlossaryGetEntriesByLetterRequest {
	return &NullableModGlossaryGetEntriesByLetterRequest{value: val, isSet: true}
}

func (v NullableModGlossaryGetEntriesByLetterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModGlossaryGetEntriesByLetterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

