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

// checks if the ModLessonGetLessonAccessInformation200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModLessonGetLessonAccessInformation200Response{}

// ModLessonGetLessonAccessInformation200Response struct for ModLessonGetLessonAccessInformation200Response
type ModLessonGetLessonAccessInformation200Response struct {
	// The number of attempts done by the user.
	Attemptscount int32 `json:"attemptscount"`
	// Whether the user can grade the lesson or not.
	Cangrade bool `json:"cangrade"`
	// Whether the user can manage the lesson or not.
	Canmanage bool `json:"canmanage"`
	// Whether the user can view the lesson reports or not.
	Canviewreports bool `json:"canviewreports"`
	// The lesson first page id.
	Firstpageid int32 `json:"firstpageid"`
	// The last page seen id.
	Lastpageseen int32 `json:"lastpageseen"`
	// Whether the user left during a timed session.
	Leftduringtimedsession bool `json:"leftduringtimedsession"`
	Preventaccessreasons []ModLessonGetLessonAccessInformation200ResponsePreventaccessreasonsInner `json:"preventaccessreasons"`
	// Whether the lesson is in review mode for the current user.
	Reviewmode bool `json:"reviewmode"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModLessonGetLessonAccessInformation200Response ModLessonGetLessonAccessInformation200Response

// NewModLessonGetLessonAccessInformation200Response instantiates a new ModLessonGetLessonAccessInformation200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModLessonGetLessonAccessInformation200Response(attemptscount int32, cangrade bool, canmanage bool, canviewreports bool, firstpageid int32, lastpageseen int32, leftduringtimedsession bool, preventaccessreasons []ModLessonGetLessonAccessInformation200ResponsePreventaccessreasonsInner, reviewmode bool) *ModLessonGetLessonAccessInformation200Response {
	this := ModLessonGetLessonAccessInformation200Response{}
	this.Attemptscount = attemptscount
	this.Cangrade = cangrade
	this.Canmanage = canmanage
	this.Canviewreports = canviewreports
	this.Firstpageid = firstpageid
	this.Lastpageseen = lastpageseen
	this.Leftduringtimedsession = leftduringtimedsession
	this.Preventaccessreasons = preventaccessreasons
	this.Reviewmode = reviewmode
	return &this
}

// NewModLessonGetLessonAccessInformation200ResponseWithDefaults instantiates a new ModLessonGetLessonAccessInformation200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModLessonGetLessonAccessInformation200ResponseWithDefaults() *ModLessonGetLessonAccessInformation200Response {
	this := ModLessonGetLessonAccessInformation200Response{}
	var attemptscount int32 = null
	this.Attemptscount = attemptscount
	var cangrade bool = null
	this.Cangrade = cangrade
	var canmanage bool = null
	this.Canmanage = canmanage
	var canviewreports bool = null
	this.Canviewreports = canviewreports
	var firstpageid int32 = null
	this.Firstpageid = firstpageid
	var lastpageseen int32 = null
	this.Lastpageseen = lastpageseen
	var leftduringtimedsession bool = null
	this.Leftduringtimedsession = leftduringtimedsession
	var reviewmode bool = null
	this.Reviewmode = reviewmode
	return &this
}

// GetAttemptscount returns the Attemptscount field value
func (o *ModLessonGetLessonAccessInformation200Response) GetAttemptscount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Attemptscount
}

// GetAttemptscountOk returns a tuple with the Attemptscount field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetLessonAccessInformation200Response) GetAttemptscountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attemptscount, true
}

// SetAttemptscount sets field value
func (o *ModLessonGetLessonAccessInformation200Response) SetAttemptscount(v int32) {
	o.Attemptscount = v
}

// GetCangrade returns the Cangrade field value
func (o *ModLessonGetLessonAccessInformation200Response) GetCangrade() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Cangrade
}

// GetCangradeOk returns a tuple with the Cangrade field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetLessonAccessInformation200Response) GetCangradeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cangrade, true
}

// SetCangrade sets field value
func (o *ModLessonGetLessonAccessInformation200Response) SetCangrade(v bool) {
	o.Cangrade = v
}

// GetCanmanage returns the Canmanage field value
func (o *ModLessonGetLessonAccessInformation200Response) GetCanmanage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canmanage
}

// GetCanmanageOk returns a tuple with the Canmanage field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetLessonAccessInformation200Response) GetCanmanageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canmanage, true
}

// SetCanmanage sets field value
func (o *ModLessonGetLessonAccessInformation200Response) SetCanmanage(v bool) {
	o.Canmanage = v
}

// GetCanviewreports returns the Canviewreports field value
func (o *ModLessonGetLessonAccessInformation200Response) GetCanviewreports() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canviewreports
}

// GetCanviewreportsOk returns a tuple with the Canviewreports field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetLessonAccessInformation200Response) GetCanviewreportsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canviewreports, true
}

// SetCanviewreports sets field value
func (o *ModLessonGetLessonAccessInformation200Response) SetCanviewreports(v bool) {
	o.Canviewreports = v
}

// GetFirstpageid returns the Firstpageid field value
func (o *ModLessonGetLessonAccessInformation200Response) GetFirstpageid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Firstpageid
}

// GetFirstpageidOk returns a tuple with the Firstpageid field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetLessonAccessInformation200Response) GetFirstpageidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Firstpageid, true
}

// SetFirstpageid sets field value
func (o *ModLessonGetLessonAccessInformation200Response) SetFirstpageid(v int32) {
	o.Firstpageid = v
}

// GetLastpageseen returns the Lastpageseen field value
func (o *ModLessonGetLessonAccessInformation200Response) GetLastpageseen() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Lastpageseen
}

// GetLastpageseenOk returns a tuple with the Lastpageseen field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetLessonAccessInformation200Response) GetLastpageseenOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lastpageseen, true
}

// SetLastpageseen sets field value
func (o *ModLessonGetLessonAccessInformation200Response) SetLastpageseen(v int32) {
	o.Lastpageseen = v
}

// GetLeftduringtimedsession returns the Leftduringtimedsession field value
func (o *ModLessonGetLessonAccessInformation200Response) GetLeftduringtimedsession() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Leftduringtimedsession
}

// GetLeftduringtimedsessionOk returns a tuple with the Leftduringtimedsession field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetLessonAccessInformation200Response) GetLeftduringtimedsessionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Leftduringtimedsession, true
}

// SetLeftduringtimedsession sets field value
func (o *ModLessonGetLessonAccessInformation200Response) SetLeftduringtimedsession(v bool) {
	o.Leftduringtimedsession = v
}

// GetPreventaccessreasons returns the Preventaccessreasons field value
func (o *ModLessonGetLessonAccessInformation200Response) GetPreventaccessreasons() []ModLessonGetLessonAccessInformation200ResponsePreventaccessreasonsInner {
	if o == nil {
		var ret []ModLessonGetLessonAccessInformation200ResponsePreventaccessreasonsInner
		return ret
	}

	return o.Preventaccessreasons
}

// GetPreventaccessreasonsOk returns a tuple with the Preventaccessreasons field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetLessonAccessInformation200Response) GetPreventaccessreasonsOk() ([]ModLessonGetLessonAccessInformation200ResponsePreventaccessreasonsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Preventaccessreasons, true
}

// SetPreventaccessreasons sets field value
func (o *ModLessonGetLessonAccessInformation200Response) SetPreventaccessreasons(v []ModLessonGetLessonAccessInformation200ResponsePreventaccessreasonsInner) {
	o.Preventaccessreasons = v
}

// GetReviewmode returns the Reviewmode field value
func (o *ModLessonGetLessonAccessInformation200Response) GetReviewmode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Reviewmode
}

// GetReviewmodeOk returns a tuple with the Reviewmode field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetLessonAccessInformation200Response) GetReviewmodeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reviewmode, true
}

// SetReviewmode sets field value
func (o *ModLessonGetLessonAccessInformation200Response) SetReviewmode(v bool) {
	o.Reviewmode = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModLessonGetLessonAccessInformation200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModLessonGetLessonAccessInformation200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModLessonGetLessonAccessInformation200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModLessonGetLessonAccessInformation200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModLessonGetLessonAccessInformation200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModLessonGetLessonAccessInformation200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attemptscount"] = o.Attemptscount
	toSerialize["cangrade"] = o.Cangrade
	toSerialize["canmanage"] = o.Canmanage
	toSerialize["canviewreports"] = o.Canviewreports
	toSerialize["firstpageid"] = o.Firstpageid
	toSerialize["lastpageseen"] = o.Lastpageseen
	toSerialize["leftduringtimedsession"] = o.Leftduringtimedsession
	toSerialize["preventaccessreasons"] = o.Preventaccessreasons
	toSerialize["reviewmode"] = o.Reviewmode
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModLessonGetLessonAccessInformation200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attemptscount",
		"cangrade",
		"canmanage",
		"canviewreports",
		"firstpageid",
		"lastpageseen",
		"leftduringtimedsession",
		"preventaccessreasons",
		"reviewmode",
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

	varModLessonGetLessonAccessInformation200Response := _ModLessonGetLessonAccessInformation200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModLessonGetLessonAccessInformation200Response)

	if err != nil {
		return err
	}

	*o = ModLessonGetLessonAccessInformation200Response(varModLessonGetLessonAccessInformation200Response)

	return err
}

type NullableModLessonGetLessonAccessInformation200Response struct {
	value *ModLessonGetLessonAccessInformation200Response
	isSet bool
}

func (v NullableModLessonGetLessonAccessInformation200Response) Get() *ModLessonGetLessonAccessInformation200Response {
	return v.value
}

func (v *NullableModLessonGetLessonAccessInformation200Response) Set(val *ModLessonGetLessonAccessInformation200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModLessonGetLessonAccessInformation200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModLessonGetLessonAccessInformation200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModLessonGetLessonAccessInformation200Response(val *ModLessonGetLessonAccessInformation200Response) *NullableModLessonGetLessonAccessInformation200Response {
	return &NullableModLessonGetLessonAccessInformation200Response{value: val, isSet: true}
}

func (v NullableModLessonGetLessonAccessInformation200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModLessonGetLessonAccessInformation200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

