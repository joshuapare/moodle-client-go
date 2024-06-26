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

// checks if the ModQuizAddRandomQuestionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModQuizAddRandomQuestionsRequest{}

// ModQuizAddRandomQuestionsRequest struct for ModQuizAddRandomQuestionsRequest
type ModQuizAddRandomQuestionsRequest struct {
	// The page where random questions will be added to
	Addonpage int32 `json:"addonpage"`
	// The cmid of the quiz
	Cmid int32 `json:"cmid"`
	// (Optional) The filter condition used when adding random questions from an existing category.                     Not required if adding random questions from a new category.
	Filtercondition *string `json:"filtercondition,omitempty"`
	// (Optional) The name of a new question category to create and use for the random questions.
	Newcategory *string `json:"newcategory,omitempty"`
	// (Optional) The parent of the new question category, if creating one.
	Parentcategory *string `json:"parentcategory,omitempty"`
	// Number of random questions
	Randomcount int32 `json:"randomcount"`
}

type _ModQuizAddRandomQuestionsRequest ModQuizAddRandomQuestionsRequest

// NewModQuizAddRandomQuestionsRequest instantiates a new ModQuizAddRandomQuestionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModQuizAddRandomQuestionsRequest(addonpage int32, cmid int32, randomcount int32) *ModQuizAddRandomQuestionsRequest {
	this := ModQuizAddRandomQuestionsRequest{}
	this.Addonpage = addonpage
	this.Cmid = cmid
	var filtercondition string = ""
	this.Filtercondition = &filtercondition
	var newcategory string = ""
	this.Newcategory = &newcategory
	var parentcategory string = "0"
	this.Parentcategory = &parentcategory
	this.Randomcount = randomcount
	return &this
}

// NewModQuizAddRandomQuestionsRequestWithDefaults instantiates a new ModQuizAddRandomQuestionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModQuizAddRandomQuestionsRequestWithDefaults() *ModQuizAddRandomQuestionsRequest {
	this := ModQuizAddRandomQuestionsRequest{}
	var addonpage int32 = null
	this.Addonpage = addonpage
	var cmid int32 = null
	this.Cmid = cmid
	var filtercondition string = ""
	this.Filtercondition = &filtercondition
	var newcategory string = ""
	this.Newcategory = &newcategory
	var parentcategory string = "0"
	this.Parentcategory = &parentcategory
	var randomcount int32 = null
	this.Randomcount = randomcount
	return &this
}

// GetAddonpage returns the Addonpage field value
func (o *ModQuizAddRandomQuestionsRequest) GetAddonpage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Addonpage
}

// GetAddonpageOk returns a tuple with the Addonpage field value
// and a boolean to check if the value has been set.
func (o *ModQuizAddRandomQuestionsRequest) GetAddonpageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addonpage, true
}

// SetAddonpage sets field value
func (o *ModQuizAddRandomQuestionsRequest) SetAddonpage(v int32) {
	o.Addonpage = v
}

// GetCmid returns the Cmid field value
func (o *ModQuizAddRandomQuestionsRequest) GetCmid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cmid
}

// GetCmidOk returns a tuple with the Cmid field value
// and a boolean to check if the value has been set.
func (o *ModQuizAddRandomQuestionsRequest) GetCmidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cmid, true
}

// SetCmid sets field value
func (o *ModQuizAddRandomQuestionsRequest) SetCmid(v int32) {
	o.Cmid = v
}

// GetFiltercondition returns the Filtercondition field value if set, zero value otherwise.
func (o *ModQuizAddRandomQuestionsRequest) GetFiltercondition() string {
	if o == nil || IsNil(o.Filtercondition) {
		var ret string
		return ret
	}
	return *o.Filtercondition
}

// GetFilterconditionOk returns a tuple with the Filtercondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModQuizAddRandomQuestionsRequest) GetFilterconditionOk() (*string, bool) {
	if o == nil || IsNil(o.Filtercondition) {
		return nil, false
	}
	return o.Filtercondition, true
}

// HasFiltercondition returns a boolean if a field has been set.
func (o *ModQuizAddRandomQuestionsRequest) HasFiltercondition() bool {
	if o != nil && !IsNil(o.Filtercondition) {
		return true
	}

	return false
}

// SetFiltercondition gets a reference to the given string and assigns it to the Filtercondition field.
func (o *ModQuizAddRandomQuestionsRequest) SetFiltercondition(v string) {
	o.Filtercondition = &v
}

// GetNewcategory returns the Newcategory field value if set, zero value otherwise.
func (o *ModQuizAddRandomQuestionsRequest) GetNewcategory() string {
	if o == nil || IsNil(o.Newcategory) {
		var ret string
		return ret
	}
	return *o.Newcategory
}

// GetNewcategoryOk returns a tuple with the Newcategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModQuizAddRandomQuestionsRequest) GetNewcategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Newcategory) {
		return nil, false
	}
	return o.Newcategory, true
}

// HasNewcategory returns a boolean if a field has been set.
func (o *ModQuizAddRandomQuestionsRequest) HasNewcategory() bool {
	if o != nil && !IsNil(o.Newcategory) {
		return true
	}

	return false
}

// SetNewcategory gets a reference to the given string and assigns it to the Newcategory field.
func (o *ModQuizAddRandomQuestionsRequest) SetNewcategory(v string) {
	o.Newcategory = &v
}

// GetParentcategory returns the Parentcategory field value if set, zero value otherwise.
func (o *ModQuizAddRandomQuestionsRequest) GetParentcategory() string {
	if o == nil || IsNil(o.Parentcategory) {
		var ret string
		return ret
	}
	return *o.Parentcategory
}

// GetParentcategoryOk returns a tuple with the Parentcategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModQuizAddRandomQuestionsRequest) GetParentcategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Parentcategory) {
		return nil, false
	}
	return o.Parentcategory, true
}

// HasParentcategory returns a boolean if a field has been set.
func (o *ModQuizAddRandomQuestionsRequest) HasParentcategory() bool {
	if o != nil && !IsNil(o.Parentcategory) {
		return true
	}

	return false
}

// SetParentcategory gets a reference to the given string and assigns it to the Parentcategory field.
func (o *ModQuizAddRandomQuestionsRequest) SetParentcategory(v string) {
	o.Parentcategory = &v
}

// GetRandomcount returns the Randomcount field value
func (o *ModQuizAddRandomQuestionsRequest) GetRandomcount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Randomcount
}

// GetRandomcountOk returns a tuple with the Randomcount field value
// and a boolean to check if the value has been set.
func (o *ModQuizAddRandomQuestionsRequest) GetRandomcountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Randomcount, true
}

// SetRandomcount sets field value
func (o *ModQuizAddRandomQuestionsRequest) SetRandomcount(v int32) {
	o.Randomcount = v
}

func (o ModQuizAddRandomQuestionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModQuizAddRandomQuestionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addonpage"] = o.Addonpage
	toSerialize["cmid"] = o.Cmid
	if !IsNil(o.Filtercondition) {
		toSerialize["filtercondition"] = o.Filtercondition
	}
	if !IsNil(o.Newcategory) {
		toSerialize["newcategory"] = o.Newcategory
	}
	if !IsNil(o.Parentcategory) {
		toSerialize["parentcategory"] = o.Parentcategory
	}
	toSerialize["randomcount"] = o.Randomcount
	return toSerialize, nil
}

func (o *ModQuizAddRandomQuestionsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addonpage",
		"cmid",
		"randomcount",
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

	varModQuizAddRandomQuestionsRequest := _ModQuizAddRandomQuestionsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModQuizAddRandomQuestionsRequest)

	if err != nil {
		return err
	}

	*o = ModQuizAddRandomQuestionsRequest(varModQuizAddRandomQuestionsRequest)

	return err
}

type NullableModQuizAddRandomQuestionsRequest struct {
	value *ModQuizAddRandomQuestionsRequest
	isSet bool
}

func (v NullableModQuizAddRandomQuestionsRequest) Get() *ModQuizAddRandomQuestionsRequest {
	return v.value
}

func (v *NullableModQuizAddRandomQuestionsRequest) Set(val *ModQuizAddRandomQuestionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModQuizAddRandomQuestionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModQuizAddRandomQuestionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModQuizAddRandomQuestionsRequest(val *ModQuizAddRandomQuestionsRequest) *NullableModQuizAddRandomQuestionsRequest {
	return &NullableModQuizAddRandomQuestionsRequest{value: val, isSet: true}
}

func (v NullableModQuizAddRandomQuestionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModQuizAddRandomQuestionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


