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

// checks if the ModGlossaryGetEntriesByAuthorIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModGlossaryGetEntriesByAuthorIdRequest{}

// ModGlossaryGetEntriesByAuthorIdRequest struct for ModGlossaryGetEntriesByAuthorIdRequest
type ModGlossaryGetEntriesByAuthorIdRequest struct {
	// The author ID
	Authorid int32 `json:"authorid"`
	// Start returning records from here
	From *int32 `json:"from,omitempty"`
	// Glossary entry ID
	Id int32 `json:"id"`
	// Number of records to return
	Limit *int32 `json:"limit,omitempty"`
	Options *ModGlossaryGetEntriesByAuthorRequestOptions `json:"options,omitempty"`
	// Order by: 'CONCEPT', 'CREATION' or 'UPDATE'
	Order *string `json:"order,omitempty"`
	// The direction of the order: 'ASC' or 'DESC'
	Sort *string `json:"sort,omitempty"`
}

type _ModGlossaryGetEntriesByAuthorIdRequest ModGlossaryGetEntriesByAuthorIdRequest

// NewModGlossaryGetEntriesByAuthorIdRequest instantiates a new ModGlossaryGetEntriesByAuthorIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModGlossaryGetEntriesByAuthorIdRequest(authorid int32, id int32) *ModGlossaryGetEntriesByAuthorIdRequest {
	this := ModGlossaryGetEntriesByAuthorIdRequest{}
	this.Authorid = authorid
	var from int32 = 0
	this.From = &from
	this.Id = id
	var limit int32 = 20
	this.Limit = &limit
	var order string = "CONCEPT"
	this.Order = &order
	var sort string = "ASC"
	this.Sort = &sort
	return &this
}

// NewModGlossaryGetEntriesByAuthorIdRequestWithDefaults instantiates a new ModGlossaryGetEntriesByAuthorIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModGlossaryGetEntriesByAuthorIdRequestWithDefaults() *ModGlossaryGetEntriesByAuthorIdRequest {
	this := ModGlossaryGetEntriesByAuthorIdRequest{}
	var authorid int32 = null
	this.Authorid = authorid
	var from int32 = 0
	this.From = &from
	var limit int32 = 20
	this.Limit = &limit
	var order string = "CONCEPT"
	this.Order = &order
	var sort string = "ASC"
	this.Sort = &sort
	return &this
}

// GetAuthorid returns the Authorid field value
func (o *ModGlossaryGetEntriesByAuthorIdRequest) GetAuthorid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Authorid
}

// GetAuthoridOk returns a tuple with the Authorid field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) GetAuthoridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authorid, true
}

// SetAuthorid sets field value
func (o *ModGlossaryGetEntriesByAuthorIdRequest) SetAuthorid(v int32) {
	o.Authorid = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) GetFrom() int32 {
	if o == nil || IsNil(o.From) {
		var ret int32
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) GetFromOk() (*int32, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int32 and assigns it to the From field.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) SetFrom(v int32) {
	o.From = &v
}

// GetId returns the Id field value
func (o *ModGlossaryGetEntriesByAuthorIdRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModGlossaryGetEntriesByAuthorIdRequest) SetId(v int32) {
	o.Id = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) GetOptions() ModGlossaryGetEntriesByAuthorRequestOptions {
	if o == nil || IsNil(o.Options) {
		var ret ModGlossaryGetEntriesByAuthorRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) GetOptionsOk() (*ModGlossaryGetEntriesByAuthorRequestOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ModGlossaryGetEntriesByAuthorRequestOptions and assigns it to the Options field.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) SetOptions(v ModGlossaryGetEntriesByAuthorRequestOptions) {
	o.Options = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) SetOrder(v string) {
	o.Order = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) GetSort() string {
	if o == nil || IsNil(o.Sort) {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) GetSortOk() (*string, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *ModGlossaryGetEntriesByAuthorIdRequest) SetSort(v string) {
	o.Sort = &v
}

func (o ModGlossaryGetEntriesByAuthorIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModGlossaryGetEntriesByAuthorIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorid"] = o.Authorid
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
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	return toSerialize, nil
}

func (o *ModGlossaryGetEntriesByAuthorIdRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorid",
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

	varModGlossaryGetEntriesByAuthorIdRequest := _ModGlossaryGetEntriesByAuthorIdRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModGlossaryGetEntriesByAuthorIdRequest)

	if err != nil {
		return err
	}

	*o = ModGlossaryGetEntriesByAuthorIdRequest(varModGlossaryGetEntriesByAuthorIdRequest)

	return err
}

type NullableModGlossaryGetEntriesByAuthorIdRequest struct {
	value *ModGlossaryGetEntriesByAuthorIdRequest
	isSet bool
}

func (v NullableModGlossaryGetEntriesByAuthorIdRequest) Get() *ModGlossaryGetEntriesByAuthorIdRequest {
	return v.value
}

func (v *NullableModGlossaryGetEntriesByAuthorIdRequest) Set(val *ModGlossaryGetEntriesByAuthorIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModGlossaryGetEntriesByAuthorIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModGlossaryGetEntriesByAuthorIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModGlossaryGetEntriesByAuthorIdRequest(val *ModGlossaryGetEntriesByAuthorIdRequest) *NullableModGlossaryGetEntriesByAuthorIdRequest {
	return &NullableModGlossaryGetEntriesByAuthorIdRequest{value: val, isSet: true}
}

func (v NullableModGlossaryGetEntriesByAuthorIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModGlossaryGetEntriesByAuthorIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


