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

// checks if the CoreRatingGetItemRatingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreRatingGetItemRatingsRequest{}

// CoreRatingGetItemRatingsRequest struct for CoreRatingGetItemRatingsRequest
type CoreRatingGetItemRatingsRequest struct {
	// component
	Component string `json:"component"`
	// context level: course, module, user, etc...
	Contextlevel string `json:"contextlevel"`
	// the instance id of item associated with the context level
	Instanceid int32 `json:"instanceid"`
	// associated id
	Itemid int32 `json:"itemid"`
	// rating area
	Ratingarea string `json:"ratingarea"`
	// scale id
	Scaleid int32 `json:"scaleid"`
	// sort order (firstname, rating or timemodified)
	Sort string `json:"sort"`
}

type _CoreRatingGetItemRatingsRequest CoreRatingGetItemRatingsRequest

// NewCoreRatingGetItemRatingsRequest instantiates a new CoreRatingGetItemRatingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreRatingGetItemRatingsRequest(component string, contextlevel string, instanceid int32, itemid int32, ratingarea string, scaleid int32, sort string) *CoreRatingGetItemRatingsRequest {
	this := CoreRatingGetItemRatingsRequest{}
	this.Component = component
	this.Contextlevel = contextlevel
	this.Instanceid = instanceid
	this.Itemid = itemid
	this.Ratingarea = ratingarea
	this.Scaleid = scaleid
	this.Sort = sort
	return &this
}

// NewCoreRatingGetItemRatingsRequestWithDefaults instantiates a new CoreRatingGetItemRatingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreRatingGetItemRatingsRequestWithDefaults() *CoreRatingGetItemRatingsRequest {
	this := CoreRatingGetItemRatingsRequest{}
	var sort string = "null"
	this.Sort = sort
	return &this
}

// GetComponent returns the Component field value
func (o *CoreRatingGetItemRatingsRequest) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *CoreRatingGetItemRatingsRequest) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *CoreRatingGetItemRatingsRequest) SetComponent(v string) {
	o.Component = v
}

// GetContextlevel returns the Contextlevel field value
func (o *CoreRatingGetItemRatingsRequest) GetContextlevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contextlevel
}

// GetContextlevelOk returns a tuple with the Contextlevel field value
// and a boolean to check if the value has been set.
func (o *CoreRatingGetItemRatingsRequest) GetContextlevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contextlevel, true
}

// SetContextlevel sets field value
func (o *CoreRatingGetItemRatingsRequest) SetContextlevel(v string) {
	o.Contextlevel = v
}

// GetInstanceid returns the Instanceid field value
func (o *CoreRatingGetItemRatingsRequest) GetInstanceid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Instanceid
}

// GetInstanceidOk returns a tuple with the Instanceid field value
// and a boolean to check if the value has been set.
func (o *CoreRatingGetItemRatingsRequest) GetInstanceidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instanceid, true
}

// SetInstanceid sets field value
func (o *CoreRatingGetItemRatingsRequest) SetInstanceid(v int32) {
	o.Instanceid = v
}

// GetItemid returns the Itemid field value
func (o *CoreRatingGetItemRatingsRequest) GetItemid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Itemid
}

// GetItemidOk returns a tuple with the Itemid field value
// and a boolean to check if the value has been set.
func (o *CoreRatingGetItemRatingsRequest) GetItemidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Itemid, true
}

// SetItemid sets field value
func (o *CoreRatingGetItemRatingsRequest) SetItemid(v int32) {
	o.Itemid = v
}

// GetRatingarea returns the Ratingarea field value
func (o *CoreRatingGetItemRatingsRequest) GetRatingarea() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ratingarea
}

// GetRatingareaOk returns a tuple with the Ratingarea field value
// and a boolean to check if the value has been set.
func (o *CoreRatingGetItemRatingsRequest) GetRatingareaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ratingarea, true
}

// SetRatingarea sets field value
func (o *CoreRatingGetItemRatingsRequest) SetRatingarea(v string) {
	o.Ratingarea = v
}

// GetScaleid returns the Scaleid field value
func (o *CoreRatingGetItemRatingsRequest) GetScaleid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Scaleid
}

// GetScaleidOk returns a tuple with the Scaleid field value
// and a boolean to check if the value has been set.
func (o *CoreRatingGetItemRatingsRequest) GetScaleidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scaleid, true
}

// SetScaleid sets field value
func (o *CoreRatingGetItemRatingsRequest) SetScaleid(v int32) {
	o.Scaleid = v
}

// GetSort returns the Sort field value
func (o *CoreRatingGetItemRatingsRequest) GetSort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value
// and a boolean to check if the value has been set.
func (o *CoreRatingGetItemRatingsRequest) GetSortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sort, true
}

// SetSort sets field value
func (o *CoreRatingGetItemRatingsRequest) SetSort(v string) {
	o.Sort = v
}

func (o CoreRatingGetItemRatingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreRatingGetItemRatingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["component"] = o.Component
	toSerialize["contextlevel"] = o.Contextlevel
	toSerialize["instanceid"] = o.Instanceid
	toSerialize["itemid"] = o.Itemid
	toSerialize["ratingarea"] = o.Ratingarea
	toSerialize["scaleid"] = o.Scaleid
	toSerialize["sort"] = o.Sort
	return toSerialize, nil
}

func (o *CoreRatingGetItemRatingsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"component",
		"contextlevel",
		"instanceid",
		"itemid",
		"ratingarea",
		"scaleid",
		"sort",
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

	varCoreRatingGetItemRatingsRequest := _CoreRatingGetItemRatingsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreRatingGetItemRatingsRequest)

	if err != nil {
		return err
	}

	*o = CoreRatingGetItemRatingsRequest(varCoreRatingGetItemRatingsRequest)

	return err
}

type NullableCoreRatingGetItemRatingsRequest struct {
	value *CoreRatingGetItemRatingsRequest
	isSet bool
}

func (v NullableCoreRatingGetItemRatingsRequest) Get() *CoreRatingGetItemRatingsRequest {
	return v.value
}

func (v *NullableCoreRatingGetItemRatingsRequest) Set(val *CoreRatingGetItemRatingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreRatingGetItemRatingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreRatingGetItemRatingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreRatingGetItemRatingsRequest(val *CoreRatingGetItemRatingsRequest) *NullableCoreRatingGetItemRatingsRequest {
	return &NullableCoreRatingGetItemRatingsRequest{value: val, isSet: true}
}

func (v NullableCoreRatingGetItemRatingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreRatingGetItemRatingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

