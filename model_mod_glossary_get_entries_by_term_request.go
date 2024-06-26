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

// checks if the ModGlossaryGetEntriesByTermRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModGlossaryGetEntriesByTermRequest{}

// ModGlossaryGetEntriesByTermRequest struct for ModGlossaryGetEntriesByTermRequest
type ModGlossaryGetEntriesByTermRequest struct {
	// Start returning records from here
	From *int32 `json:"from,omitempty"`
	// Glossary entry ID
	Id int32 `json:"id"`
	// Number of records to return
	Limit *int32 `json:"limit,omitempty"`
	Options *ModGlossaryGetEntriesByAuthorRequestOptions `json:"options,omitempty"`
	// The entry concept, or alias
	Term string `json:"term"`
}

type _ModGlossaryGetEntriesByTermRequest ModGlossaryGetEntriesByTermRequest

// NewModGlossaryGetEntriesByTermRequest instantiates a new ModGlossaryGetEntriesByTermRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModGlossaryGetEntriesByTermRequest(id int32, term string) *ModGlossaryGetEntriesByTermRequest {
	this := ModGlossaryGetEntriesByTermRequest{}
	var from int32 = 0
	this.From = &from
	this.Id = id
	var limit int32 = 20
	this.Limit = &limit
	this.Term = term
	return &this
}

// NewModGlossaryGetEntriesByTermRequestWithDefaults instantiates a new ModGlossaryGetEntriesByTermRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModGlossaryGetEntriesByTermRequestWithDefaults() *ModGlossaryGetEntriesByTermRequest {
	this := ModGlossaryGetEntriesByTermRequest{}
	var from int32 = 0
	this.From = &from
	var limit int32 = 20
	this.Limit = &limit
	var term string = "null"
	this.Term = term
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByTermRequest) GetFrom() int32 {
	if o == nil || IsNil(o.From) {
		var ret int32
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByTermRequest) GetFromOk() (*int32, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByTermRequest) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int32 and assigns it to the From field.
func (o *ModGlossaryGetEntriesByTermRequest) SetFrom(v int32) {
	o.From = &v
}

// GetId returns the Id field value
func (o *ModGlossaryGetEntriesByTermRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByTermRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModGlossaryGetEntriesByTermRequest) SetId(v int32) {
	o.Id = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByTermRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByTermRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByTermRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ModGlossaryGetEntriesByTermRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByTermRequest) GetOptions() ModGlossaryGetEntriesByAuthorRequestOptions {
	if o == nil || IsNil(o.Options) {
		var ret ModGlossaryGetEntriesByAuthorRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByTermRequest) GetOptionsOk() (*ModGlossaryGetEntriesByAuthorRequestOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByTermRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ModGlossaryGetEntriesByAuthorRequestOptions and assigns it to the Options field.
func (o *ModGlossaryGetEntriesByTermRequest) SetOptions(v ModGlossaryGetEntriesByAuthorRequestOptions) {
	o.Options = &v
}

// GetTerm returns the Term field value
func (o *ModGlossaryGetEntriesByTermRequest) GetTerm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Term
}

// GetTermOk returns a tuple with the Term field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByTermRequest) GetTermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Term, true
}

// SetTerm sets field value
func (o *ModGlossaryGetEntriesByTermRequest) SetTerm(v string) {
	o.Term = v
}

func (o ModGlossaryGetEntriesByTermRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModGlossaryGetEntriesByTermRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	toSerialize["term"] = o.Term
	return toSerialize, nil
}

func (o *ModGlossaryGetEntriesByTermRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"term",
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

	varModGlossaryGetEntriesByTermRequest := _ModGlossaryGetEntriesByTermRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModGlossaryGetEntriesByTermRequest)

	if err != nil {
		return err
	}

	*o = ModGlossaryGetEntriesByTermRequest(varModGlossaryGetEntriesByTermRequest)

	return err
}

type NullableModGlossaryGetEntriesByTermRequest struct {
	value *ModGlossaryGetEntriesByTermRequest
	isSet bool
}

func (v NullableModGlossaryGetEntriesByTermRequest) Get() *ModGlossaryGetEntriesByTermRequest {
	return v.value
}

func (v *NullableModGlossaryGetEntriesByTermRequest) Set(val *ModGlossaryGetEntriesByTermRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModGlossaryGetEntriesByTermRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModGlossaryGetEntriesByTermRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModGlossaryGetEntriesByTermRequest(val *ModGlossaryGetEntriesByTermRequest) *NullableModGlossaryGetEntriesByTermRequest {
	return &NullableModGlossaryGetEntriesByTermRequest{value: val, isSet: true}
}

func (v NullableModGlossaryGetEntriesByTermRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModGlossaryGetEntriesByTermRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


