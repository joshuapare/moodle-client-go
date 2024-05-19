/*
Moodle Webservice API

Auto-generated OpenAPI spec for Moodle's Webservice API. Requires installation of the Catalyst RESTful API plugin: https://github.com/catalyst/moodle-webservice_restful.

API version: 4.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moodleclient

import (
	"encoding/json"
)

// checks if the ModChoiceGetChoiceResults200ResponseOptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModChoiceGetChoiceResults200ResponseOptionsInner{}

// ModChoiceGetChoiceResults200ResponseOptionsInner Options
type ModChoiceGetChoiceResults200ResponseOptionsInner struct {
	// choice instance id
	Id *int32 `json:"id,omitempty"`
	// maximum number of answers
	Maxanswer *int32 `json:"maxanswer,omitempty"`
	// number of users answers
	Numberofuser *int32 `json:"numberofuser,omitempty"`
	// percentage of users answers
	Percentageamount *float32 `json:"percentageamount,omitempty"`
	// text of the choice
	Text *string `json:"text,omitempty"`
	Userresponses []ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner `json:"userresponses,omitempty"`
}

// NewModChoiceGetChoiceResults200ResponseOptionsInner instantiates a new ModChoiceGetChoiceResults200ResponseOptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModChoiceGetChoiceResults200ResponseOptionsInner() *ModChoiceGetChoiceResults200ResponseOptionsInner {
	this := ModChoiceGetChoiceResults200ResponseOptionsInner{}
	var numberofuser int32 = null
	this.Numberofuser = &numberofuser
	var percentageamount float32 = null
	this.Percentageamount = &percentageamount
	return &this
}

// NewModChoiceGetChoiceResults200ResponseOptionsInnerWithDefaults instantiates a new ModChoiceGetChoiceResults200ResponseOptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModChoiceGetChoiceResults200ResponseOptionsInnerWithDefaults() *ModChoiceGetChoiceResults200ResponseOptionsInner {
	this := ModChoiceGetChoiceResults200ResponseOptionsInner{}
	var numberofuser int32 = null
	this.Numberofuser = &numberofuser
	var percentageamount float32 = null
	this.Percentageamount = &percentageamount
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) SetId(v int32) {
	o.Id = &v
}

// GetMaxanswer returns the Maxanswer field value if set, zero value otherwise.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) GetMaxanswer() int32 {
	if o == nil || IsNil(o.Maxanswer) {
		var ret int32
		return ret
	}
	return *o.Maxanswer
}

// GetMaxanswerOk returns a tuple with the Maxanswer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) GetMaxanswerOk() (*int32, bool) {
	if o == nil || IsNil(o.Maxanswer) {
		return nil, false
	}
	return o.Maxanswer, true
}

// HasMaxanswer returns a boolean if a field has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) HasMaxanswer() bool {
	if o != nil && !IsNil(o.Maxanswer) {
		return true
	}

	return false
}

// SetMaxanswer gets a reference to the given int32 and assigns it to the Maxanswer field.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) SetMaxanswer(v int32) {
	o.Maxanswer = &v
}

// GetNumberofuser returns the Numberofuser field value if set, zero value otherwise.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) GetNumberofuser() int32 {
	if o == nil || IsNil(o.Numberofuser) {
		var ret int32
		return ret
	}
	return *o.Numberofuser
}

// GetNumberofuserOk returns a tuple with the Numberofuser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) GetNumberofuserOk() (*int32, bool) {
	if o == nil || IsNil(o.Numberofuser) {
		return nil, false
	}
	return o.Numberofuser, true
}

// HasNumberofuser returns a boolean if a field has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) HasNumberofuser() bool {
	if o != nil && !IsNil(o.Numberofuser) {
		return true
	}

	return false
}

// SetNumberofuser gets a reference to the given int32 and assigns it to the Numberofuser field.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) SetNumberofuser(v int32) {
	o.Numberofuser = &v
}

// GetPercentageamount returns the Percentageamount field value if set, zero value otherwise.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) GetPercentageamount() float32 {
	if o == nil || IsNil(o.Percentageamount) {
		var ret float32
		return ret
	}
	return *o.Percentageamount
}

// GetPercentageamountOk returns a tuple with the Percentageamount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) GetPercentageamountOk() (*float32, bool) {
	if o == nil || IsNil(o.Percentageamount) {
		return nil, false
	}
	return o.Percentageamount, true
}

// HasPercentageamount returns a boolean if a field has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) HasPercentageamount() bool {
	if o != nil && !IsNil(o.Percentageamount) {
		return true
	}

	return false
}

// SetPercentageamount gets a reference to the given float32 and assigns it to the Percentageamount field.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) SetPercentageamount(v float32) {
	o.Percentageamount = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) SetText(v string) {
	o.Text = &v
}

// GetUserresponses returns the Userresponses field value if set, zero value otherwise.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) GetUserresponses() []ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner {
	if o == nil || IsNil(o.Userresponses) {
		var ret []ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner
		return ret
	}
	return o.Userresponses
}

// GetUserresponsesOk returns a tuple with the Userresponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) GetUserresponsesOk() ([]ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner, bool) {
	if o == nil || IsNil(o.Userresponses) {
		return nil, false
	}
	return o.Userresponses, true
}

// HasUserresponses returns a boolean if a field has been set.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) HasUserresponses() bool {
	if o != nil && !IsNil(o.Userresponses) {
		return true
	}

	return false
}

// SetUserresponses gets a reference to the given []ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner and assigns it to the Userresponses field.
func (o *ModChoiceGetChoiceResults200ResponseOptionsInner) SetUserresponses(v []ModChoiceGetChoiceResults200ResponseOptionsInnerUserresponsesInner) {
	o.Userresponses = v
}

func (o ModChoiceGetChoiceResults200ResponseOptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModChoiceGetChoiceResults200ResponseOptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Maxanswer) {
		toSerialize["maxanswer"] = o.Maxanswer
	}
	if !IsNil(o.Numberofuser) {
		toSerialize["numberofuser"] = o.Numberofuser
	}
	if !IsNil(o.Percentageamount) {
		toSerialize["percentageamount"] = o.Percentageamount
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Userresponses) {
		toSerialize["userresponses"] = o.Userresponses
	}
	return toSerialize, nil
}

type NullableModChoiceGetChoiceResults200ResponseOptionsInner struct {
	value *ModChoiceGetChoiceResults200ResponseOptionsInner
	isSet bool
}

func (v NullableModChoiceGetChoiceResults200ResponseOptionsInner) Get() *ModChoiceGetChoiceResults200ResponseOptionsInner {
	return v.value
}

func (v *NullableModChoiceGetChoiceResults200ResponseOptionsInner) Set(val *ModChoiceGetChoiceResults200ResponseOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModChoiceGetChoiceResults200ResponseOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModChoiceGetChoiceResults200ResponseOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModChoiceGetChoiceResults200ResponseOptionsInner(val *ModChoiceGetChoiceResults200ResponseOptionsInner) *NullableModChoiceGetChoiceResults200ResponseOptionsInner {
	return &NullableModChoiceGetChoiceResults200ResponseOptionsInner{value: val, isSet: true}
}

func (v NullableModChoiceGetChoiceResults200ResponseOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModChoiceGetChoiceResults200ResponseOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

