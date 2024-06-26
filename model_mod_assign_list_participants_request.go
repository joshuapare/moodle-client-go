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

// checks if the ModAssignListParticipantsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModAssignListParticipantsRequest{}

// ModAssignListParticipantsRequest struct for ModAssignListParticipantsRequest
type ModAssignListParticipantsRequest struct {
	// assign instance id
	Assignid int32 `json:"assignid"`
	// search string to filter the results
	Filter string `json:"filter"`
	// group id
	Groupid int32 `json:"groupid"`
	// Do return courses where the user is enrolled
	Includeenrolments *bool `json:"includeenrolments,omitempty"`
	// maximum number of records to return
	Limit *int32 `json:"limit,omitempty"`
	// Do not return all user fields
	Onlyids *bool `json:"onlyids,omitempty"`
	// number of records to skip
	Skip *int32 `json:"skip,omitempty"`
	// Apply current user table sorting preferences.
	Tablesort *bool `json:"tablesort,omitempty"`
}

type _ModAssignListParticipantsRequest ModAssignListParticipantsRequest

// NewModAssignListParticipantsRequest instantiates a new ModAssignListParticipantsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModAssignListParticipantsRequest(assignid int32, filter string, groupid int32) *ModAssignListParticipantsRequest {
	this := ModAssignListParticipantsRequest{}
	this.Assignid = assignid
	this.Filter = filter
	this.Groupid = groupid
	var includeenrolments bool = true
	this.Includeenrolments = &includeenrolments
	var limit int32 = 0
	this.Limit = &limit
	var onlyids bool = false
	this.Onlyids = &onlyids
	var skip int32 = 0
	this.Skip = &skip
	var tablesort bool = false
	this.Tablesort = &tablesort
	return &this
}

// NewModAssignListParticipantsRequestWithDefaults instantiates a new ModAssignListParticipantsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModAssignListParticipantsRequestWithDefaults() *ModAssignListParticipantsRequest {
	this := ModAssignListParticipantsRequest{}
	var filter string = "null"
	this.Filter = filter
	var includeenrolments bool = true
	this.Includeenrolments = &includeenrolments
	var limit int32 = 0
	this.Limit = &limit
	var onlyids bool = false
	this.Onlyids = &onlyids
	var skip int32 = 0
	this.Skip = &skip
	var tablesort bool = false
	this.Tablesort = &tablesort
	return &this
}

// GetAssignid returns the Assignid field value
func (o *ModAssignListParticipantsRequest) GetAssignid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Assignid
}

// GetAssignidOk returns a tuple with the Assignid field value
// and a boolean to check if the value has been set.
func (o *ModAssignListParticipantsRequest) GetAssignidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assignid, true
}

// SetAssignid sets field value
func (o *ModAssignListParticipantsRequest) SetAssignid(v int32) {
	o.Assignid = v
}

// GetFilter returns the Filter field value
func (o *ModAssignListParticipantsRequest) GetFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *ModAssignListParticipantsRequest) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *ModAssignListParticipantsRequest) SetFilter(v string) {
	o.Filter = v
}

// GetGroupid returns the Groupid field value
func (o *ModAssignListParticipantsRequest) GetGroupid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Groupid
}

// GetGroupidOk returns a tuple with the Groupid field value
// and a boolean to check if the value has been set.
func (o *ModAssignListParticipantsRequest) GetGroupidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Groupid, true
}

// SetGroupid sets field value
func (o *ModAssignListParticipantsRequest) SetGroupid(v int32) {
	o.Groupid = v
}

// GetIncludeenrolments returns the Includeenrolments field value if set, zero value otherwise.
func (o *ModAssignListParticipantsRequest) GetIncludeenrolments() bool {
	if o == nil || IsNil(o.Includeenrolments) {
		var ret bool
		return ret
	}
	return *o.Includeenrolments
}

// GetIncludeenrolmentsOk returns a tuple with the Includeenrolments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignListParticipantsRequest) GetIncludeenrolmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.Includeenrolments) {
		return nil, false
	}
	return o.Includeenrolments, true
}

// HasIncludeenrolments returns a boolean if a field has been set.
func (o *ModAssignListParticipantsRequest) HasIncludeenrolments() bool {
	if o != nil && !IsNil(o.Includeenrolments) {
		return true
	}

	return false
}

// SetIncludeenrolments gets a reference to the given bool and assigns it to the Includeenrolments field.
func (o *ModAssignListParticipantsRequest) SetIncludeenrolments(v bool) {
	o.Includeenrolments = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ModAssignListParticipantsRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignListParticipantsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ModAssignListParticipantsRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ModAssignListParticipantsRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetOnlyids returns the Onlyids field value if set, zero value otherwise.
func (o *ModAssignListParticipantsRequest) GetOnlyids() bool {
	if o == nil || IsNil(o.Onlyids) {
		var ret bool
		return ret
	}
	return *o.Onlyids
}

// GetOnlyidsOk returns a tuple with the Onlyids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignListParticipantsRequest) GetOnlyidsOk() (*bool, bool) {
	if o == nil || IsNil(o.Onlyids) {
		return nil, false
	}
	return o.Onlyids, true
}

// HasOnlyids returns a boolean if a field has been set.
func (o *ModAssignListParticipantsRequest) HasOnlyids() bool {
	if o != nil && !IsNil(o.Onlyids) {
		return true
	}

	return false
}

// SetOnlyids gets a reference to the given bool and assigns it to the Onlyids field.
func (o *ModAssignListParticipantsRequest) SetOnlyids(v bool) {
	o.Onlyids = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *ModAssignListParticipantsRequest) GetSkip() int32 {
	if o == nil || IsNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignListParticipantsRequest) GetSkipOk() (*int32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *ModAssignListParticipantsRequest) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *ModAssignListParticipantsRequest) SetSkip(v int32) {
	o.Skip = &v
}

// GetTablesort returns the Tablesort field value if set, zero value otherwise.
func (o *ModAssignListParticipantsRequest) GetTablesort() bool {
	if o == nil || IsNil(o.Tablesort) {
		var ret bool
		return ret
	}
	return *o.Tablesort
}

// GetTablesortOk returns a tuple with the Tablesort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignListParticipantsRequest) GetTablesortOk() (*bool, bool) {
	if o == nil || IsNil(o.Tablesort) {
		return nil, false
	}
	return o.Tablesort, true
}

// HasTablesort returns a boolean if a field has been set.
func (o *ModAssignListParticipantsRequest) HasTablesort() bool {
	if o != nil && !IsNil(o.Tablesort) {
		return true
	}

	return false
}

// SetTablesort gets a reference to the given bool and assigns it to the Tablesort field.
func (o *ModAssignListParticipantsRequest) SetTablesort(v bool) {
	o.Tablesort = &v
}

func (o ModAssignListParticipantsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModAssignListParticipantsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assignid"] = o.Assignid
	toSerialize["filter"] = o.Filter
	toSerialize["groupid"] = o.Groupid
	if !IsNil(o.Includeenrolments) {
		toSerialize["includeenrolments"] = o.Includeenrolments
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Onlyids) {
		toSerialize["onlyids"] = o.Onlyids
	}
	if !IsNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !IsNil(o.Tablesort) {
		toSerialize["tablesort"] = o.Tablesort
	}
	return toSerialize, nil
}

func (o *ModAssignListParticipantsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assignid",
		"filter",
		"groupid",
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

	varModAssignListParticipantsRequest := _ModAssignListParticipantsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModAssignListParticipantsRequest)

	if err != nil {
		return err
	}

	*o = ModAssignListParticipantsRequest(varModAssignListParticipantsRequest)

	return err
}

type NullableModAssignListParticipantsRequest struct {
	value *ModAssignListParticipantsRequest
	isSet bool
}

func (v NullableModAssignListParticipantsRequest) Get() *ModAssignListParticipantsRequest {
	return v.value
}

func (v *NullableModAssignListParticipantsRequest) Set(val *ModAssignListParticipantsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModAssignListParticipantsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModAssignListParticipantsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModAssignListParticipantsRequest(val *ModAssignListParticipantsRequest) *NullableModAssignListParticipantsRequest {
	return &NullableModAssignListParticipantsRequest{value: val, isSet: true}
}

func (v NullableModAssignListParticipantsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModAssignListParticipantsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


