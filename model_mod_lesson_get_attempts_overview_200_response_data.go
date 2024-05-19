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

// checks if the ModLessonGetAttemptsOverview200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModLessonGetAttemptsOverview200ResponseData{}

// ModLessonGetAttemptsOverview200ResponseData struct for ModLessonGetAttemptsOverview200ResponseData
type ModLessonGetAttemptsOverview200ResponseData struct {
	// Average score.
	Avescore float32 `json:"avescore"`
	// Average time (spent in taking the lesson).
	Avetime int32 `json:"avetime"`
	// High score.
	Highscore float32 `json:"highscore"`
	// High time.
	Hightime int32 `json:"hightime"`
	// True if the lesson was scored.
	Lessonscored bool `json:"lessonscored"`
	// Low score.
	Lowscore float32 `json:"lowscore"`
	// Low time.
	Lowtime int32 `json:"lowtime"`
	// Number of attempts.
	Numofattempts int32 `json:"numofattempts"`
	Students []ModLessonGetAttemptsOverview200ResponseDataStudentsInner `json:"students,omitempty"`
}

type _ModLessonGetAttemptsOverview200ResponseData ModLessonGetAttemptsOverview200ResponseData

// NewModLessonGetAttemptsOverview200ResponseData instantiates a new ModLessonGetAttemptsOverview200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModLessonGetAttemptsOverview200ResponseData(avescore float32, avetime int32, highscore float32, hightime int32, lessonscored bool, lowscore float32, lowtime int32, numofattempts int32) *ModLessonGetAttemptsOverview200ResponseData {
	this := ModLessonGetAttemptsOverview200ResponseData{}
	this.Avescore = avescore
	this.Avetime = avetime
	this.Highscore = highscore
	this.Hightime = hightime
	this.Lessonscored = lessonscored
	this.Lowscore = lowscore
	this.Lowtime = lowtime
	this.Numofattempts = numofattempts
	return &this
}

// NewModLessonGetAttemptsOverview200ResponseDataWithDefaults instantiates a new ModLessonGetAttemptsOverview200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModLessonGetAttemptsOverview200ResponseDataWithDefaults() *ModLessonGetAttemptsOverview200ResponseData {
	this := ModLessonGetAttemptsOverview200ResponseData{}
	var avescore float32 = null
	this.Avescore = avescore
	var avetime int32 = null
	this.Avetime = avetime
	var highscore float32 = null
	this.Highscore = highscore
	var hightime int32 = null
	this.Hightime = hightime
	var lessonscored bool = null
	this.Lessonscored = lessonscored
	var lowscore float32 = null
	this.Lowscore = lowscore
	var lowtime int32 = null
	this.Lowtime = lowtime
	var numofattempts int32 = null
	this.Numofattempts = numofattempts
	return &this
}

// GetAvescore returns the Avescore field value
func (o *ModLessonGetAttemptsOverview200ResponseData) GetAvescore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Avescore
}

// GetAvescoreOk returns a tuple with the Avescore field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetAttemptsOverview200ResponseData) GetAvescoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avescore, true
}

// SetAvescore sets field value
func (o *ModLessonGetAttemptsOverview200ResponseData) SetAvescore(v float32) {
	o.Avescore = v
}

// GetAvetime returns the Avetime field value
func (o *ModLessonGetAttemptsOverview200ResponseData) GetAvetime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Avetime
}

// GetAvetimeOk returns a tuple with the Avetime field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetAttemptsOverview200ResponseData) GetAvetimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avetime, true
}

// SetAvetime sets field value
func (o *ModLessonGetAttemptsOverview200ResponseData) SetAvetime(v int32) {
	o.Avetime = v
}

// GetHighscore returns the Highscore field value
func (o *ModLessonGetAttemptsOverview200ResponseData) GetHighscore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Highscore
}

// GetHighscoreOk returns a tuple with the Highscore field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetAttemptsOverview200ResponseData) GetHighscoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Highscore, true
}

// SetHighscore sets field value
func (o *ModLessonGetAttemptsOverview200ResponseData) SetHighscore(v float32) {
	o.Highscore = v
}

// GetHightime returns the Hightime field value
func (o *ModLessonGetAttemptsOverview200ResponseData) GetHightime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Hightime
}

// GetHightimeOk returns a tuple with the Hightime field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetAttemptsOverview200ResponseData) GetHightimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hightime, true
}

// SetHightime sets field value
func (o *ModLessonGetAttemptsOverview200ResponseData) SetHightime(v int32) {
	o.Hightime = v
}

// GetLessonscored returns the Lessonscored field value
func (o *ModLessonGetAttemptsOverview200ResponseData) GetLessonscored() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Lessonscored
}

// GetLessonscoredOk returns a tuple with the Lessonscored field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetAttemptsOverview200ResponseData) GetLessonscoredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lessonscored, true
}

// SetLessonscored sets field value
func (o *ModLessonGetAttemptsOverview200ResponseData) SetLessonscored(v bool) {
	o.Lessonscored = v
}

// GetLowscore returns the Lowscore field value
func (o *ModLessonGetAttemptsOverview200ResponseData) GetLowscore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Lowscore
}

// GetLowscoreOk returns a tuple with the Lowscore field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetAttemptsOverview200ResponseData) GetLowscoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lowscore, true
}

// SetLowscore sets field value
func (o *ModLessonGetAttemptsOverview200ResponseData) SetLowscore(v float32) {
	o.Lowscore = v
}

// GetLowtime returns the Lowtime field value
func (o *ModLessonGetAttemptsOverview200ResponseData) GetLowtime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Lowtime
}

// GetLowtimeOk returns a tuple with the Lowtime field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetAttemptsOverview200ResponseData) GetLowtimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lowtime, true
}

// SetLowtime sets field value
func (o *ModLessonGetAttemptsOverview200ResponseData) SetLowtime(v int32) {
	o.Lowtime = v
}

// GetNumofattempts returns the Numofattempts field value
func (o *ModLessonGetAttemptsOverview200ResponseData) GetNumofattempts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Numofattempts
}

// GetNumofattemptsOk returns a tuple with the Numofattempts field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetAttemptsOverview200ResponseData) GetNumofattemptsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Numofattempts, true
}

// SetNumofattempts sets field value
func (o *ModLessonGetAttemptsOverview200ResponseData) SetNumofattempts(v int32) {
	o.Numofattempts = v
}

// GetStudents returns the Students field value if set, zero value otherwise.
func (o *ModLessonGetAttemptsOverview200ResponseData) GetStudents() []ModLessonGetAttemptsOverview200ResponseDataStudentsInner {
	if o == nil || IsNil(o.Students) {
		var ret []ModLessonGetAttemptsOverview200ResponseDataStudentsInner
		return ret
	}
	return o.Students
}

// GetStudentsOk returns a tuple with the Students field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModLessonGetAttemptsOverview200ResponseData) GetStudentsOk() ([]ModLessonGetAttemptsOverview200ResponseDataStudentsInner, bool) {
	if o == nil || IsNil(o.Students) {
		return nil, false
	}
	return o.Students, true
}

// HasStudents returns a boolean if a field has been set.
func (o *ModLessonGetAttemptsOverview200ResponseData) HasStudents() bool {
	if o != nil && !IsNil(o.Students) {
		return true
	}

	return false
}

// SetStudents gets a reference to the given []ModLessonGetAttemptsOverview200ResponseDataStudentsInner and assigns it to the Students field.
func (o *ModLessonGetAttemptsOverview200ResponseData) SetStudents(v []ModLessonGetAttemptsOverview200ResponseDataStudentsInner) {
	o.Students = v
}

func (o ModLessonGetAttemptsOverview200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModLessonGetAttemptsOverview200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["avescore"] = o.Avescore
	toSerialize["avetime"] = o.Avetime
	toSerialize["highscore"] = o.Highscore
	toSerialize["hightime"] = o.Hightime
	toSerialize["lessonscored"] = o.Lessonscored
	toSerialize["lowscore"] = o.Lowscore
	toSerialize["lowtime"] = o.Lowtime
	toSerialize["numofattempts"] = o.Numofattempts
	if !IsNil(o.Students) {
		toSerialize["students"] = o.Students
	}
	return toSerialize, nil
}

func (o *ModLessonGetAttemptsOverview200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"avescore",
		"avetime",
		"highscore",
		"hightime",
		"lessonscored",
		"lowscore",
		"lowtime",
		"numofattempts",
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

	varModLessonGetAttemptsOverview200ResponseData := _ModLessonGetAttemptsOverview200ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModLessonGetAttemptsOverview200ResponseData)

	if err != nil {
		return err
	}

	*o = ModLessonGetAttemptsOverview200ResponseData(varModLessonGetAttemptsOverview200ResponseData)

	return err
}

type NullableModLessonGetAttemptsOverview200ResponseData struct {
	value *ModLessonGetAttemptsOverview200ResponseData
	isSet bool
}

func (v NullableModLessonGetAttemptsOverview200ResponseData) Get() *ModLessonGetAttemptsOverview200ResponseData {
	return v.value
}

func (v *NullableModLessonGetAttemptsOverview200ResponseData) Set(val *ModLessonGetAttemptsOverview200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableModLessonGetAttemptsOverview200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableModLessonGetAttemptsOverview200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModLessonGetAttemptsOverview200ResponseData(val *ModLessonGetAttemptsOverview200ResponseData) *NullableModLessonGetAttemptsOverview200ResponseData {
	return &NullableModLessonGetAttemptsOverview200ResponseData{value: val, isSet: true}
}

func (v NullableModLessonGetAttemptsOverview200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModLessonGetAttemptsOverview200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


