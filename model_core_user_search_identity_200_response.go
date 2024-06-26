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

// checks if the CoreUserSearchIdentity200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreUserSearchIdentity200Response{}

// CoreUserSearchIdentity200Response struct for CoreUserSearchIdentity200Response
type CoreUserSearchIdentity200Response struct {
	List []CoreUserSearchIdentity200ResponseListInner `json:"list"`
	// Configured maximum users per page.
	Maxusersperpage int32 `json:"maxusersperpage"`
	// Were there more records than maxusersperpage found?
	Overflow bool `json:"overflow"`
}

type _CoreUserSearchIdentity200Response CoreUserSearchIdentity200Response

// NewCoreUserSearchIdentity200Response instantiates a new CoreUserSearchIdentity200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUserSearchIdentity200Response(list []CoreUserSearchIdentity200ResponseListInner, maxusersperpage int32, overflow bool) *CoreUserSearchIdentity200Response {
	this := CoreUserSearchIdentity200Response{}
	this.List = list
	this.Maxusersperpage = maxusersperpage
	this.Overflow = overflow
	return &this
}

// NewCoreUserSearchIdentity200ResponseWithDefaults instantiates a new CoreUserSearchIdentity200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUserSearchIdentity200ResponseWithDefaults() *CoreUserSearchIdentity200Response {
	this := CoreUserSearchIdentity200Response{}
	var maxusersperpage int32 = null
	this.Maxusersperpage = maxusersperpage
	var overflow bool = null
	this.Overflow = overflow
	return &this
}

// GetList returns the List field value
func (o *CoreUserSearchIdentity200Response) GetList() []CoreUserSearchIdentity200ResponseListInner {
	if o == nil {
		var ret []CoreUserSearchIdentity200ResponseListInner
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *CoreUserSearchIdentity200Response) GetListOk() ([]CoreUserSearchIdentity200ResponseListInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.List, true
}

// SetList sets field value
func (o *CoreUserSearchIdentity200Response) SetList(v []CoreUserSearchIdentity200ResponseListInner) {
	o.List = v
}

// GetMaxusersperpage returns the Maxusersperpage field value
func (o *CoreUserSearchIdentity200Response) GetMaxusersperpage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Maxusersperpage
}

// GetMaxusersperpageOk returns a tuple with the Maxusersperpage field value
// and a boolean to check if the value has been set.
func (o *CoreUserSearchIdentity200Response) GetMaxusersperpageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Maxusersperpage, true
}

// SetMaxusersperpage sets field value
func (o *CoreUserSearchIdentity200Response) SetMaxusersperpage(v int32) {
	o.Maxusersperpage = v
}

// GetOverflow returns the Overflow field value
func (o *CoreUserSearchIdentity200Response) GetOverflow() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Overflow
}

// GetOverflowOk returns a tuple with the Overflow field value
// and a boolean to check if the value has been set.
func (o *CoreUserSearchIdentity200Response) GetOverflowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Overflow, true
}

// SetOverflow sets field value
func (o *CoreUserSearchIdentity200Response) SetOverflow(v bool) {
	o.Overflow = v
}

func (o CoreUserSearchIdentity200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreUserSearchIdentity200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["list"] = o.List
	toSerialize["maxusersperpage"] = o.Maxusersperpage
	toSerialize["overflow"] = o.Overflow
	return toSerialize, nil
}

func (o *CoreUserSearchIdentity200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"list",
		"maxusersperpage",
		"overflow",
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

	varCoreUserSearchIdentity200Response := _CoreUserSearchIdentity200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreUserSearchIdentity200Response)

	if err != nil {
		return err
	}

	*o = CoreUserSearchIdentity200Response(varCoreUserSearchIdentity200Response)

	return err
}

type NullableCoreUserSearchIdentity200Response struct {
	value *CoreUserSearchIdentity200Response
	isSet bool
}

func (v NullableCoreUserSearchIdentity200Response) Get() *CoreUserSearchIdentity200Response {
	return v.value
}

func (v *NullableCoreUserSearchIdentity200Response) Set(val *CoreUserSearchIdentity200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUserSearchIdentity200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUserSearchIdentity200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUserSearchIdentity200Response(val *CoreUserSearchIdentity200Response) *NullableCoreUserSearchIdentity200Response {
	return &NullableCoreUserSearchIdentity200Response{value: val, isSet: true}
}

func (v NullableCoreUserSearchIdentity200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUserSearchIdentity200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


