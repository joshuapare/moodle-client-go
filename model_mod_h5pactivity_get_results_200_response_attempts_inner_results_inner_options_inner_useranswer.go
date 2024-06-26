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

// checks if the ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer{}

// ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer struct for ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer
type ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer struct {
	// Option text value
	Answer *string `json:"answer,omitempty"`
	// If has to be displayed as a checked option
	Checked *bool `json:"checked,omitempty"`
	// If has to be displayed as correct
	Correct *bool `json:"correct,omitempty"`
	// If has to be displayed as failed
	Fail *bool `json:"fail,omitempty"`
	// If has to be displayed as incorrect
	Incorrect *bool `json:"incorrect,omitempty"`
	// If has to be displayed as passed
	Pass *bool `json:"pass,omitempty"`
	// If has to be displayed as simple text
	Text *bool `json:"text,omitempty"`
	// If has to be displayed as a unchecked option
	Unchecked *bool `json:"unchecked,omitempty"`
}

// NewModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer instantiates a new ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer() *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer {
	this := ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer{}
	return &this
}

// NewModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswerWithDefaults instantiates a new ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswerWithDefaults() *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer {
	this := ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer{}
	return &this
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetAnswer() string {
	if o == nil || IsNil(o.Answer) {
		var ret string
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetAnswerOk() (*string, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given string and assigns it to the Answer field.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) SetAnswer(v string) {
	o.Answer = &v
}

// GetChecked returns the Checked field value if set, zero value otherwise.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetChecked() bool {
	if o == nil || IsNil(o.Checked) {
		var ret bool
		return ret
	}
	return *o.Checked
}

// GetCheckedOk returns a tuple with the Checked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetCheckedOk() (*bool, bool) {
	if o == nil || IsNil(o.Checked) {
		return nil, false
	}
	return o.Checked, true
}

// HasChecked returns a boolean if a field has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) HasChecked() bool {
	if o != nil && !IsNil(o.Checked) {
		return true
	}

	return false
}

// SetChecked gets a reference to the given bool and assigns it to the Checked field.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) SetChecked(v bool) {
	o.Checked = &v
}

// GetCorrect returns the Correct field value if set, zero value otherwise.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetCorrect() bool {
	if o == nil || IsNil(o.Correct) {
		var ret bool
		return ret
	}
	return *o.Correct
}

// GetCorrectOk returns a tuple with the Correct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetCorrectOk() (*bool, bool) {
	if o == nil || IsNil(o.Correct) {
		return nil, false
	}
	return o.Correct, true
}

// HasCorrect returns a boolean if a field has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) HasCorrect() bool {
	if o != nil && !IsNil(o.Correct) {
		return true
	}

	return false
}

// SetCorrect gets a reference to the given bool and assigns it to the Correct field.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) SetCorrect(v bool) {
	o.Correct = &v
}

// GetFail returns the Fail field value if set, zero value otherwise.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetFail() bool {
	if o == nil || IsNil(o.Fail) {
		var ret bool
		return ret
	}
	return *o.Fail
}

// GetFailOk returns a tuple with the Fail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetFailOk() (*bool, bool) {
	if o == nil || IsNil(o.Fail) {
		return nil, false
	}
	return o.Fail, true
}

// HasFail returns a boolean if a field has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) HasFail() bool {
	if o != nil && !IsNil(o.Fail) {
		return true
	}

	return false
}

// SetFail gets a reference to the given bool and assigns it to the Fail field.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) SetFail(v bool) {
	o.Fail = &v
}

// GetIncorrect returns the Incorrect field value if set, zero value otherwise.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetIncorrect() bool {
	if o == nil || IsNil(o.Incorrect) {
		var ret bool
		return ret
	}
	return *o.Incorrect
}

// GetIncorrectOk returns a tuple with the Incorrect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetIncorrectOk() (*bool, bool) {
	if o == nil || IsNil(o.Incorrect) {
		return nil, false
	}
	return o.Incorrect, true
}

// HasIncorrect returns a boolean if a field has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) HasIncorrect() bool {
	if o != nil && !IsNil(o.Incorrect) {
		return true
	}

	return false
}

// SetIncorrect gets a reference to the given bool and assigns it to the Incorrect field.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) SetIncorrect(v bool) {
	o.Incorrect = &v
}

// GetPass returns the Pass field value if set, zero value otherwise.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetPass() bool {
	if o == nil || IsNil(o.Pass) {
		var ret bool
		return ret
	}
	return *o.Pass
}

// GetPassOk returns a tuple with the Pass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetPassOk() (*bool, bool) {
	if o == nil || IsNil(o.Pass) {
		return nil, false
	}
	return o.Pass, true
}

// HasPass returns a boolean if a field has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) HasPass() bool {
	if o != nil && !IsNil(o.Pass) {
		return true
	}

	return false
}

// SetPass gets a reference to the given bool and assigns it to the Pass field.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) SetPass(v bool) {
	o.Pass = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetText() bool {
	if o == nil || IsNil(o.Text) {
		var ret bool
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetTextOk() (*bool, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given bool and assigns it to the Text field.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) SetText(v bool) {
	o.Text = &v
}

// GetUnchecked returns the Unchecked field value if set, zero value otherwise.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetUnchecked() bool {
	if o == nil || IsNil(o.Unchecked) {
		var ret bool
		return ret
	}
	return *o.Unchecked
}

// GetUncheckedOk returns a tuple with the Unchecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) GetUncheckedOk() (*bool, bool) {
	if o == nil || IsNil(o.Unchecked) {
		return nil, false
	}
	return o.Unchecked, true
}

// HasUnchecked returns a boolean if a field has been set.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) HasUnchecked() bool {
	if o != nil && !IsNil(o.Unchecked) {
		return true
	}

	return false
}

// SetUnchecked gets a reference to the given bool and assigns it to the Unchecked field.
func (o *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) SetUnchecked(v bool) {
	o.Unchecked = &v
}

func (o ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	if !IsNil(o.Checked) {
		toSerialize["checked"] = o.Checked
	}
	if !IsNil(o.Correct) {
		toSerialize["correct"] = o.Correct
	}
	if !IsNil(o.Fail) {
		toSerialize["fail"] = o.Fail
	}
	if !IsNil(o.Incorrect) {
		toSerialize["incorrect"] = o.Incorrect
	}
	if !IsNil(o.Pass) {
		toSerialize["pass"] = o.Pass
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Unchecked) {
		toSerialize["unchecked"] = o.Unchecked
	}
	return toSerialize, nil
}

type NullableModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer struct {
	value *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer
	isSet bool
}

func (v NullableModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) Get() *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer {
	return v.value
}

func (v *NullableModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) Set(val *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) {
	v.value = val
	v.isSet = true
}

func (v NullableModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) IsSet() bool {
	return v.isSet
}

func (v *NullableModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer(val *ModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) *NullableModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer {
	return &NullableModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer{value: val, isSet: true}
}

func (v NullableModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModH5pactivityGetResults200ResponseAttemptsInnerResultsInnerOptionsInnerUseranswer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


