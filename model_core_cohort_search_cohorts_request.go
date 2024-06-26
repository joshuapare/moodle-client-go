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

// checks if the CoreCohortSearchCohortsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCohortSearchCohortsRequest{}

// CoreCohortSearchCohortsRequest struct for CoreCohortSearchCohortsRequest
type CoreCohortSearchCohortsRequest struct {
	Context CoreCohortSearchCohortsRequestContext `json:"context"`
	// What other contexts to fetch the frameworks from. (all, parents, self)
	Includes *string `json:"includes,omitempty"`
	// limitfrom we are fetching the records from
	Limitfrom *int32 `json:"limitfrom,omitempty"`
	// Number of records to fetch
	Limitnum *int32 `json:"limitnum,omitempty"`
	// Query string
	Query string `json:"query"`
}

type _CoreCohortSearchCohortsRequest CoreCohortSearchCohortsRequest

// NewCoreCohortSearchCohortsRequest instantiates a new CoreCohortSearchCohortsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCohortSearchCohortsRequest(context CoreCohortSearchCohortsRequestContext, query string) *CoreCohortSearchCohortsRequest {
	this := CoreCohortSearchCohortsRequest{}
	this.Context = context
	var includes string = "parents"
	this.Includes = &includes
	var limitfrom int32 = 0
	this.Limitfrom = &limitfrom
	var limitnum int32 = 25
	this.Limitnum = &limitnum
	this.Query = query
	return &this
}

// NewCoreCohortSearchCohortsRequestWithDefaults instantiates a new CoreCohortSearchCohortsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCohortSearchCohortsRequestWithDefaults() *CoreCohortSearchCohortsRequest {
	this := CoreCohortSearchCohortsRequest{}
	var includes string = "parents"
	this.Includes = &includes
	var limitfrom int32 = 0
	this.Limitfrom = &limitfrom
	var limitnum int32 = 25
	this.Limitnum = &limitnum
	var query string = "null"
	this.Query = query
	return &this
}

// GetContext returns the Context field value
func (o *CoreCohortSearchCohortsRequest) GetContext() CoreCohortSearchCohortsRequestContext {
	if o == nil {
		var ret CoreCohortSearchCohortsRequestContext
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *CoreCohortSearchCohortsRequest) GetContextOk() (*CoreCohortSearchCohortsRequestContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *CoreCohortSearchCohortsRequest) SetContext(v CoreCohortSearchCohortsRequestContext) {
	o.Context = v
}

// GetIncludes returns the Includes field value if set, zero value otherwise.
func (o *CoreCohortSearchCohortsRequest) GetIncludes() string {
	if o == nil || IsNil(o.Includes) {
		var ret string
		return ret
	}
	return *o.Includes
}

// GetIncludesOk returns a tuple with the Includes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCohortSearchCohortsRequest) GetIncludesOk() (*string, bool) {
	if o == nil || IsNil(o.Includes) {
		return nil, false
	}
	return o.Includes, true
}

// HasIncludes returns a boolean if a field has been set.
func (o *CoreCohortSearchCohortsRequest) HasIncludes() bool {
	if o != nil && !IsNil(o.Includes) {
		return true
	}

	return false
}

// SetIncludes gets a reference to the given string and assigns it to the Includes field.
func (o *CoreCohortSearchCohortsRequest) SetIncludes(v string) {
	o.Includes = &v
}

// GetLimitfrom returns the Limitfrom field value if set, zero value otherwise.
func (o *CoreCohortSearchCohortsRequest) GetLimitfrom() int32 {
	if o == nil || IsNil(o.Limitfrom) {
		var ret int32
		return ret
	}
	return *o.Limitfrom
}

// GetLimitfromOk returns a tuple with the Limitfrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCohortSearchCohortsRequest) GetLimitfromOk() (*int32, bool) {
	if o == nil || IsNil(o.Limitfrom) {
		return nil, false
	}
	return o.Limitfrom, true
}

// HasLimitfrom returns a boolean if a field has been set.
func (o *CoreCohortSearchCohortsRequest) HasLimitfrom() bool {
	if o != nil && !IsNil(o.Limitfrom) {
		return true
	}

	return false
}

// SetLimitfrom gets a reference to the given int32 and assigns it to the Limitfrom field.
func (o *CoreCohortSearchCohortsRequest) SetLimitfrom(v int32) {
	o.Limitfrom = &v
}

// GetLimitnum returns the Limitnum field value if set, zero value otherwise.
func (o *CoreCohortSearchCohortsRequest) GetLimitnum() int32 {
	if o == nil || IsNil(o.Limitnum) {
		var ret int32
		return ret
	}
	return *o.Limitnum
}

// GetLimitnumOk returns a tuple with the Limitnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCohortSearchCohortsRequest) GetLimitnumOk() (*int32, bool) {
	if o == nil || IsNil(o.Limitnum) {
		return nil, false
	}
	return o.Limitnum, true
}

// HasLimitnum returns a boolean if a field has been set.
func (o *CoreCohortSearchCohortsRequest) HasLimitnum() bool {
	if o != nil && !IsNil(o.Limitnum) {
		return true
	}

	return false
}

// SetLimitnum gets a reference to the given int32 and assigns it to the Limitnum field.
func (o *CoreCohortSearchCohortsRequest) SetLimitnum(v int32) {
	o.Limitnum = &v
}

// GetQuery returns the Query field value
func (o *CoreCohortSearchCohortsRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CoreCohortSearchCohortsRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *CoreCohortSearchCohortsRequest) SetQuery(v string) {
	o.Query = v
}

func (o CoreCohortSearchCohortsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCohortSearchCohortsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["context"] = o.Context
	if !IsNil(o.Includes) {
		toSerialize["includes"] = o.Includes
	}
	if !IsNil(o.Limitfrom) {
		toSerialize["limitfrom"] = o.Limitfrom
	}
	if !IsNil(o.Limitnum) {
		toSerialize["limitnum"] = o.Limitnum
	}
	toSerialize["query"] = o.Query
	return toSerialize, nil
}

func (o *CoreCohortSearchCohortsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"context",
		"query",
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

	varCoreCohortSearchCohortsRequest := _CoreCohortSearchCohortsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCohortSearchCohortsRequest)

	if err != nil {
		return err
	}

	*o = CoreCohortSearchCohortsRequest(varCoreCohortSearchCohortsRequest)

	return err
}

type NullableCoreCohortSearchCohortsRequest struct {
	value *CoreCohortSearchCohortsRequest
	isSet bool
}

func (v NullableCoreCohortSearchCohortsRequest) Get() *CoreCohortSearchCohortsRequest {
	return v.value
}

func (v *NullableCoreCohortSearchCohortsRequest) Set(val *CoreCohortSearchCohortsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCohortSearchCohortsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCohortSearchCohortsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCohortSearchCohortsRequest(val *CoreCohortSearchCohortsRequest) *NullableCoreCohortSearchCohortsRequest {
	return &NullableCoreCohortSearchCohortsRequest{value: val, isSet: true}
}

func (v NullableCoreCohortSearchCohortsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCohortSearchCohortsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


