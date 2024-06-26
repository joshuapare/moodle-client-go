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

// checks if the ModGlossaryGetEntriesByCategoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModGlossaryGetEntriesByCategoryRequest{}

// ModGlossaryGetEntriesByCategoryRequest struct for ModGlossaryGetEntriesByCategoryRequest
type ModGlossaryGetEntriesByCategoryRequest struct {
	// The category ID. Use '0' for all categories, or '-1' for uncategorised entries.
	Categoryid int32 `json:"categoryid"`
	// Start returning records from here
	From *int32 `json:"from,omitempty"`
	// The glossary ID.
	Id int32 `json:"id"`
	// Number of records to return
	Limit *int32 `json:"limit,omitempty"`
	Options *ModGlossaryGetEntriesByAuthorRequestOptions `json:"options,omitempty"`
}

type _ModGlossaryGetEntriesByCategoryRequest ModGlossaryGetEntriesByCategoryRequest

// NewModGlossaryGetEntriesByCategoryRequest instantiates a new ModGlossaryGetEntriesByCategoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModGlossaryGetEntriesByCategoryRequest(categoryid int32, id int32) *ModGlossaryGetEntriesByCategoryRequest {
	this := ModGlossaryGetEntriesByCategoryRequest{}
	this.Categoryid = categoryid
	var from int32 = 0
	this.From = &from
	this.Id = id
	var limit int32 = 20
	this.Limit = &limit
	return &this
}

// NewModGlossaryGetEntriesByCategoryRequestWithDefaults instantiates a new ModGlossaryGetEntriesByCategoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModGlossaryGetEntriesByCategoryRequestWithDefaults() *ModGlossaryGetEntriesByCategoryRequest {
	this := ModGlossaryGetEntriesByCategoryRequest{}
	var categoryid int32 = null
	this.Categoryid = categoryid
	var from int32 = 0
	this.From = &from
	var id int32 = null
	this.Id = id
	var limit int32 = 20
	this.Limit = &limit
	return &this
}

// GetCategoryid returns the Categoryid field value
func (o *ModGlossaryGetEntriesByCategoryRequest) GetCategoryid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Categoryid
}

// GetCategoryidOk returns a tuple with the Categoryid field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByCategoryRequest) GetCategoryidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Categoryid, true
}

// SetCategoryid sets field value
func (o *ModGlossaryGetEntriesByCategoryRequest) SetCategoryid(v int32) {
	o.Categoryid = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByCategoryRequest) GetFrom() int32 {
	if o == nil || IsNil(o.From) {
		var ret int32
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByCategoryRequest) GetFromOk() (*int32, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByCategoryRequest) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int32 and assigns it to the From field.
func (o *ModGlossaryGetEntriesByCategoryRequest) SetFrom(v int32) {
	o.From = &v
}

// GetId returns the Id field value
func (o *ModGlossaryGetEntriesByCategoryRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByCategoryRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModGlossaryGetEntriesByCategoryRequest) SetId(v int32) {
	o.Id = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByCategoryRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByCategoryRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByCategoryRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ModGlossaryGetEntriesByCategoryRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByCategoryRequest) GetOptions() ModGlossaryGetEntriesByAuthorRequestOptions {
	if o == nil || IsNil(o.Options) {
		var ret ModGlossaryGetEntriesByAuthorRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByCategoryRequest) GetOptionsOk() (*ModGlossaryGetEntriesByAuthorRequestOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByCategoryRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ModGlossaryGetEntriesByAuthorRequestOptions and assigns it to the Options field.
func (o *ModGlossaryGetEntriesByCategoryRequest) SetOptions(v ModGlossaryGetEntriesByAuthorRequestOptions) {
	o.Options = &v
}

func (o ModGlossaryGetEntriesByCategoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModGlossaryGetEntriesByCategoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["categoryid"] = o.Categoryid
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
	return toSerialize, nil
}

func (o *ModGlossaryGetEntriesByCategoryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"categoryid",
		"id",
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

	varModGlossaryGetEntriesByCategoryRequest := _ModGlossaryGetEntriesByCategoryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModGlossaryGetEntriesByCategoryRequest)

	if err != nil {
		return err
	}

	*o = ModGlossaryGetEntriesByCategoryRequest(varModGlossaryGetEntriesByCategoryRequest)

	return err
}

type NullableModGlossaryGetEntriesByCategoryRequest struct {
	value *ModGlossaryGetEntriesByCategoryRequest
	isSet bool
}

func (v NullableModGlossaryGetEntriesByCategoryRequest) Get() *ModGlossaryGetEntriesByCategoryRequest {
	return v.value
}

func (v *NullableModGlossaryGetEntriesByCategoryRequest) Set(val *ModGlossaryGetEntriesByCategoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModGlossaryGetEntriesByCategoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModGlossaryGetEntriesByCategoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModGlossaryGetEntriesByCategoryRequest(val *ModGlossaryGetEntriesByCategoryRequest) *NullableModGlossaryGetEntriesByCategoryRequest {
	return &NullableModGlossaryGetEntriesByCategoryRequest{value: val, isSet: true}
}

func (v NullableModGlossaryGetEntriesByCategoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModGlossaryGetEntriesByCategoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


