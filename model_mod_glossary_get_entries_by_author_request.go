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

// checks if the ModGlossaryGetEntriesByAuthorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModGlossaryGetEntriesByAuthorRequest{}

// ModGlossaryGetEntriesByAuthorRequest struct for ModGlossaryGetEntriesByAuthorRequest
type ModGlossaryGetEntriesByAuthorRequest struct {
	// Search and order using: 'FIRSTNAME' or 'LASTNAME'
	Field *string `json:"field,omitempty"`
	// Start returning records from here
	From *int32 `json:"from,omitempty"`
	// Glossary entry ID
	Id int32 `json:"id"`
	// First letter of firstname or lastname, or either keywords: 'ALL' or 'SPECIAL'.
	Letter string `json:"letter"`
	// Number of records to return
	Limit *int32 `json:"limit,omitempty"`
	Options *ModGlossaryGetEntriesByAuthorRequestOptions `json:"options,omitempty"`
	// The direction of the order: 'ASC' or 'DESC'
	Sort *string `json:"sort,omitempty"`
}

type _ModGlossaryGetEntriesByAuthorRequest ModGlossaryGetEntriesByAuthorRequest

// NewModGlossaryGetEntriesByAuthorRequest instantiates a new ModGlossaryGetEntriesByAuthorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModGlossaryGetEntriesByAuthorRequest(id int32, letter string) *ModGlossaryGetEntriesByAuthorRequest {
	this := ModGlossaryGetEntriesByAuthorRequest{}
	var field string = "LASTNAME"
	this.Field = &field
	var from int32 = 0
	this.From = &from
	this.Id = id
	this.Letter = letter
	var limit int32 = 20
	this.Limit = &limit
	var sort string = "ASC"
	this.Sort = &sort
	return &this
}

// NewModGlossaryGetEntriesByAuthorRequestWithDefaults instantiates a new ModGlossaryGetEntriesByAuthorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModGlossaryGetEntriesByAuthorRequestWithDefaults() *ModGlossaryGetEntriesByAuthorRequest {
	this := ModGlossaryGetEntriesByAuthorRequest{}
	var field string = "LASTNAME"
	this.Field = &field
	var from int32 = 0
	this.From = &from
	var letter string = "null"
	this.Letter = letter
	var limit int32 = 20
	this.Limit = &limit
	var sort string = "ASC"
	this.Sort = &sort
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByAuthorRequest) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByAuthorRequest) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByAuthorRequest) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *ModGlossaryGetEntriesByAuthorRequest) SetField(v string) {
	o.Field = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByAuthorRequest) GetFrom() int32 {
	if o == nil || IsNil(o.From) {
		var ret int32
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByAuthorRequest) GetFromOk() (*int32, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByAuthorRequest) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int32 and assigns it to the From field.
func (o *ModGlossaryGetEntriesByAuthorRequest) SetFrom(v int32) {
	o.From = &v
}

// GetId returns the Id field value
func (o *ModGlossaryGetEntriesByAuthorRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByAuthorRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModGlossaryGetEntriesByAuthorRequest) SetId(v int32) {
	o.Id = v
}

// GetLetter returns the Letter field value
func (o *ModGlossaryGetEntriesByAuthorRequest) GetLetter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Letter
}

// GetLetterOk returns a tuple with the Letter field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByAuthorRequest) GetLetterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Letter, true
}

// SetLetter sets field value
func (o *ModGlossaryGetEntriesByAuthorRequest) SetLetter(v string) {
	o.Letter = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByAuthorRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByAuthorRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByAuthorRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ModGlossaryGetEntriesByAuthorRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByAuthorRequest) GetOptions() ModGlossaryGetEntriesByAuthorRequestOptions {
	if o == nil || IsNil(o.Options) {
		var ret ModGlossaryGetEntriesByAuthorRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByAuthorRequest) GetOptionsOk() (*ModGlossaryGetEntriesByAuthorRequestOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByAuthorRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ModGlossaryGetEntriesByAuthorRequestOptions and assigns it to the Options field.
func (o *ModGlossaryGetEntriesByAuthorRequest) SetOptions(v ModGlossaryGetEntriesByAuthorRequestOptions) {
	o.Options = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByAuthorRequest) GetSort() string {
	if o == nil || IsNil(o.Sort) {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByAuthorRequest) GetSortOk() (*string, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByAuthorRequest) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *ModGlossaryGetEntriesByAuthorRequest) SetSort(v string) {
	o.Sort = &v
}

func (o ModGlossaryGetEntriesByAuthorRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModGlossaryGetEntriesByAuthorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
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
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	return toSerialize, nil
}

func (o *ModGlossaryGetEntriesByAuthorRequest) UnmarshalJSON(data []byte) (err error) {
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

	varModGlossaryGetEntriesByAuthorRequest := _ModGlossaryGetEntriesByAuthorRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModGlossaryGetEntriesByAuthorRequest)

	if err != nil {
		return err
	}

	*o = ModGlossaryGetEntriesByAuthorRequest(varModGlossaryGetEntriesByAuthorRequest)

	return err
}

type NullableModGlossaryGetEntriesByAuthorRequest struct {
	value *ModGlossaryGetEntriesByAuthorRequest
	isSet bool
}

func (v NullableModGlossaryGetEntriesByAuthorRequest) Get() *ModGlossaryGetEntriesByAuthorRequest {
	return v.value
}

func (v *NullableModGlossaryGetEntriesByAuthorRequest) Set(val *ModGlossaryGetEntriesByAuthorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModGlossaryGetEntriesByAuthorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModGlossaryGetEntriesByAuthorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModGlossaryGetEntriesByAuthorRequest(val *ModGlossaryGetEntriesByAuthorRequest) *NullableModGlossaryGetEntriesByAuthorRequest {
	return &NullableModGlossaryGetEntriesByAuthorRequest{value: val, isSet: true}
}

func (v NullableModGlossaryGetEntriesByAuthorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModGlossaryGetEntriesByAuthorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


