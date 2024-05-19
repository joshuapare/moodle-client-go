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

// checks if the ModAssignGetSubmissionStatus200ResponseLastattempt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModAssignGetSubmissionStatus200ResponseLastattempt{}

// ModAssignGetSubmissionStatus200ResponseLastattempt struct for ModAssignGetSubmissionStatus200ResponseLastattempt
type ModAssignGetSubmissionStatus200ResponseLastattempt struct {
	// Whether blind marking is enabled.
	Blindmarking bool `json:"blindmarking"`
	// Whether the user can edit the current submission.
	Canedit bool `json:"canedit"`
	// Whether the owner of the submission can edit it.
	Caneditowner bool `json:"caneditowner"`
	// Whether the user can submit.
	Cansubmit bool `json:"cansubmit"`
	// Extension due date.
	Extensionduedate int32 `json:"extensionduedate"`
	// Whether the submission is graded.
	Graded bool `json:"graded"`
	// Grading status.
	Gradingstatus string `json:"gradingstatus"`
	// Whether new submissions are locked.
	Locked bool `json:"locked"`
	Submission *ModAssignGetSubmissionStatus200ResponseLastattemptSubmission `json:"submission,omitempty"`
	// The submission group id (for group submissions only).
	Submissiongroup *int32 `json:"submissiongroup,omitempty"`
	Submissiongroupmemberswhoneedtosubmit []map[string]interface{} `json:"submissiongroupmemberswhoneedtosubmit,omitempty"`
	// Whether submissions are enabled or not.
	Submissionsenabled bool `json:"submissionsenabled"`
	Teamsubmission *ModAssignGetSubmissionStatus200ResponseLastattemptTeamsubmission `json:"teamsubmission,omitempty"`
	// Time limit for submission.
	Timelimit *int32 `json:"timelimit,omitempty"`
	Usergroups []map[string]interface{} `json:"usergroups"`
}

type _ModAssignGetSubmissionStatus200ResponseLastattempt ModAssignGetSubmissionStatus200ResponseLastattempt

// NewModAssignGetSubmissionStatus200ResponseLastattempt instantiates a new ModAssignGetSubmissionStatus200ResponseLastattempt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModAssignGetSubmissionStatus200ResponseLastattempt(blindmarking bool, canedit bool, caneditowner bool, cansubmit bool, extensionduedate int32, graded bool, gradingstatus string, locked bool, submissionsenabled bool, usergroups []map[string]interface{}) *ModAssignGetSubmissionStatus200ResponseLastattempt {
	this := ModAssignGetSubmissionStatus200ResponseLastattempt{}
	this.Blindmarking = blindmarking
	this.Canedit = canedit
	this.Caneditowner = caneditowner
	this.Cansubmit = cansubmit
	this.Extensionduedate = extensionduedate
	this.Graded = graded
	this.Gradingstatus = gradingstatus
	this.Locked = locked
	var submissiongroup int32 = null
	this.Submissiongroup = &submissiongroup
	this.Submissionsenabled = submissionsenabled
	var timelimit int32 = null
	this.Timelimit = &timelimit
	this.Usergroups = usergroups
	return &this
}

// NewModAssignGetSubmissionStatus200ResponseLastattemptWithDefaults instantiates a new ModAssignGetSubmissionStatus200ResponseLastattempt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModAssignGetSubmissionStatus200ResponseLastattemptWithDefaults() *ModAssignGetSubmissionStatus200ResponseLastattempt {
	this := ModAssignGetSubmissionStatus200ResponseLastattempt{}
	var blindmarking bool = null
	this.Blindmarking = blindmarking
	var canedit bool = null
	this.Canedit = canedit
	var caneditowner bool = null
	this.Caneditowner = caneditowner
	var cansubmit bool = null
	this.Cansubmit = cansubmit
	var extensionduedate int32 = null
	this.Extensionduedate = extensionduedate
	var graded bool = null
	this.Graded = graded
	var gradingstatus string = "null"
	this.Gradingstatus = gradingstatus
	var locked bool = null
	this.Locked = locked
	var submissiongroup int32 = null
	this.Submissiongroup = &submissiongroup
	var timelimit int32 = null
	this.Timelimit = &timelimit
	return &this
}

// GetBlindmarking returns the Blindmarking field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetBlindmarking() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Blindmarking
}

// GetBlindmarkingOk returns a tuple with the Blindmarking field value
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetBlindmarkingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blindmarking, true
}

// SetBlindmarking sets field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetBlindmarking(v bool) {
	o.Blindmarking = v
}

// GetCanedit returns the Canedit field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetCanedit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canedit
}

// GetCaneditOk returns a tuple with the Canedit field value
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetCaneditOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canedit, true
}

// SetCanedit sets field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetCanedit(v bool) {
	o.Canedit = v
}

// GetCaneditowner returns the Caneditowner field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetCaneditowner() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Caneditowner
}

// GetCaneditownerOk returns a tuple with the Caneditowner field value
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetCaneditownerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Caneditowner, true
}

// SetCaneditowner sets field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetCaneditowner(v bool) {
	o.Caneditowner = v
}

// GetCansubmit returns the Cansubmit field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetCansubmit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Cansubmit
}

// GetCansubmitOk returns a tuple with the Cansubmit field value
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetCansubmitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cansubmit, true
}

// SetCansubmit sets field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetCansubmit(v bool) {
	o.Cansubmit = v
}

// GetExtensionduedate returns the Extensionduedate field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetExtensionduedate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Extensionduedate
}

// GetExtensionduedateOk returns a tuple with the Extensionduedate field value
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetExtensionduedateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Extensionduedate, true
}

// SetExtensionduedate sets field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetExtensionduedate(v int32) {
	o.Extensionduedate = v
}

// GetGraded returns the Graded field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetGraded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Graded
}

// GetGradedOk returns a tuple with the Graded field value
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetGradedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Graded, true
}

// SetGraded sets field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetGraded(v bool) {
	o.Graded = v
}

// GetGradingstatus returns the Gradingstatus field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetGradingstatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gradingstatus
}

// GetGradingstatusOk returns a tuple with the Gradingstatus field value
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetGradingstatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gradingstatus, true
}

// SetGradingstatus sets field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetGradingstatus(v string) {
	o.Gradingstatus = v
}

// GetLocked returns the Locked field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetLocked(v bool) {
	o.Locked = v
}

// GetSubmission returns the Submission field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetSubmission() ModAssignGetSubmissionStatus200ResponseLastattemptSubmission {
	if o == nil || IsNil(o.Submission) {
		var ret ModAssignGetSubmissionStatus200ResponseLastattemptSubmission
		return ret
	}
	return *o.Submission
}

// GetSubmissionOk returns a tuple with the Submission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetSubmissionOk() (*ModAssignGetSubmissionStatus200ResponseLastattemptSubmission, bool) {
	if o == nil || IsNil(o.Submission) {
		return nil, false
	}
	return o.Submission, true
}

// HasSubmission returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) HasSubmission() bool {
	if o != nil && !IsNil(o.Submission) {
		return true
	}

	return false
}

// SetSubmission gets a reference to the given ModAssignGetSubmissionStatus200ResponseLastattemptSubmission and assigns it to the Submission field.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetSubmission(v ModAssignGetSubmissionStatus200ResponseLastattemptSubmission) {
	o.Submission = &v
}

// GetSubmissiongroup returns the Submissiongroup field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetSubmissiongroup() int32 {
	if o == nil || IsNil(o.Submissiongroup) {
		var ret int32
		return ret
	}
	return *o.Submissiongroup
}

// GetSubmissiongroupOk returns a tuple with the Submissiongroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetSubmissiongroupOk() (*int32, bool) {
	if o == nil || IsNil(o.Submissiongroup) {
		return nil, false
	}
	return o.Submissiongroup, true
}

// HasSubmissiongroup returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) HasSubmissiongroup() bool {
	if o != nil && !IsNil(o.Submissiongroup) {
		return true
	}

	return false
}

// SetSubmissiongroup gets a reference to the given int32 and assigns it to the Submissiongroup field.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetSubmissiongroup(v int32) {
	o.Submissiongroup = &v
}

// GetSubmissiongroupmemberswhoneedtosubmit returns the Submissiongroupmemberswhoneedtosubmit field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetSubmissiongroupmemberswhoneedtosubmit() []map[string]interface{} {
	if o == nil || IsNil(o.Submissiongroupmemberswhoneedtosubmit) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Submissiongroupmemberswhoneedtosubmit
}

// GetSubmissiongroupmemberswhoneedtosubmitOk returns a tuple with the Submissiongroupmemberswhoneedtosubmit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetSubmissiongroupmemberswhoneedtosubmitOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Submissiongroupmemberswhoneedtosubmit) {
		return nil, false
	}
	return o.Submissiongroupmemberswhoneedtosubmit, true
}

// HasSubmissiongroupmemberswhoneedtosubmit returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) HasSubmissiongroupmemberswhoneedtosubmit() bool {
	if o != nil && !IsNil(o.Submissiongroupmemberswhoneedtosubmit) {
		return true
	}

	return false
}

// SetSubmissiongroupmemberswhoneedtosubmit gets a reference to the given []map[string]interface{} and assigns it to the Submissiongroupmemberswhoneedtosubmit field.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetSubmissiongroupmemberswhoneedtosubmit(v []map[string]interface{}) {
	o.Submissiongroupmemberswhoneedtosubmit = v
}

// GetSubmissionsenabled returns the Submissionsenabled field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetSubmissionsenabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Submissionsenabled
}

// GetSubmissionsenabledOk returns a tuple with the Submissionsenabled field value
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetSubmissionsenabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Submissionsenabled, true
}

// SetSubmissionsenabled sets field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetSubmissionsenabled(v bool) {
	o.Submissionsenabled = v
}

// GetTeamsubmission returns the Teamsubmission field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetTeamsubmission() ModAssignGetSubmissionStatus200ResponseLastattemptTeamsubmission {
	if o == nil || IsNil(o.Teamsubmission) {
		var ret ModAssignGetSubmissionStatus200ResponseLastattemptTeamsubmission
		return ret
	}
	return *o.Teamsubmission
}

// GetTeamsubmissionOk returns a tuple with the Teamsubmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetTeamsubmissionOk() (*ModAssignGetSubmissionStatus200ResponseLastattemptTeamsubmission, bool) {
	if o == nil || IsNil(o.Teamsubmission) {
		return nil, false
	}
	return o.Teamsubmission, true
}

// HasTeamsubmission returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) HasTeamsubmission() bool {
	if o != nil && !IsNil(o.Teamsubmission) {
		return true
	}

	return false
}

// SetTeamsubmission gets a reference to the given ModAssignGetSubmissionStatus200ResponseLastattemptTeamsubmission and assigns it to the Teamsubmission field.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetTeamsubmission(v ModAssignGetSubmissionStatus200ResponseLastattemptTeamsubmission) {
	o.Teamsubmission = &v
}

// GetTimelimit returns the Timelimit field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetTimelimit() int32 {
	if o == nil || IsNil(o.Timelimit) {
		var ret int32
		return ret
	}
	return *o.Timelimit
}

// GetTimelimitOk returns a tuple with the Timelimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetTimelimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Timelimit) {
		return nil, false
	}
	return o.Timelimit, true
}

// HasTimelimit returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) HasTimelimit() bool {
	if o != nil && !IsNil(o.Timelimit) {
		return true
	}

	return false
}

// SetTimelimit gets a reference to the given int32 and assigns it to the Timelimit field.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetTimelimit(v int32) {
	o.Timelimit = &v
}

// GetUsergroups returns the Usergroups field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetUsergroups() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Usergroups
}

// GetUsergroupsOk returns a tuple with the Usergroups field value
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) GetUsergroupsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usergroups, true
}

// SetUsergroups sets field value
func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) SetUsergroups(v []map[string]interface{}) {
	o.Usergroups = v
}

func (o ModAssignGetSubmissionStatus200ResponseLastattempt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModAssignGetSubmissionStatus200ResponseLastattempt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["blindmarking"] = o.Blindmarking
	toSerialize["canedit"] = o.Canedit
	toSerialize["caneditowner"] = o.Caneditowner
	toSerialize["cansubmit"] = o.Cansubmit
	toSerialize["extensionduedate"] = o.Extensionduedate
	toSerialize["graded"] = o.Graded
	toSerialize["gradingstatus"] = o.Gradingstatus
	toSerialize["locked"] = o.Locked
	if !IsNil(o.Submission) {
		toSerialize["submission"] = o.Submission
	}
	if !IsNil(o.Submissiongroup) {
		toSerialize["submissiongroup"] = o.Submissiongroup
	}
	if !IsNil(o.Submissiongroupmemberswhoneedtosubmit) {
		toSerialize["submissiongroupmemberswhoneedtosubmit"] = o.Submissiongroupmemberswhoneedtosubmit
	}
	toSerialize["submissionsenabled"] = o.Submissionsenabled
	if !IsNil(o.Teamsubmission) {
		toSerialize["teamsubmission"] = o.Teamsubmission
	}
	if !IsNil(o.Timelimit) {
		toSerialize["timelimit"] = o.Timelimit
	}
	toSerialize["usergroups"] = o.Usergroups
	return toSerialize, nil
}

func (o *ModAssignGetSubmissionStatus200ResponseLastattempt) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"blindmarking",
		"canedit",
		"caneditowner",
		"cansubmit",
		"extensionduedate",
		"graded",
		"gradingstatus",
		"locked",
		"submissionsenabled",
		"usergroups",
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

	varModAssignGetSubmissionStatus200ResponseLastattempt := _ModAssignGetSubmissionStatus200ResponseLastattempt{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModAssignGetSubmissionStatus200ResponseLastattempt)

	if err != nil {
		return err
	}

	*o = ModAssignGetSubmissionStatus200ResponseLastattempt(varModAssignGetSubmissionStatus200ResponseLastattempt)

	return err
}

type NullableModAssignGetSubmissionStatus200ResponseLastattempt struct {
	value *ModAssignGetSubmissionStatus200ResponseLastattempt
	isSet bool
}

func (v NullableModAssignGetSubmissionStatus200ResponseLastattempt) Get() *ModAssignGetSubmissionStatus200ResponseLastattempt {
	return v.value
}

func (v *NullableModAssignGetSubmissionStatus200ResponseLastattempt) Set(val *ModAssignGetSubmissionStatus200ResponseLastattempt) {
	v.value = val
	v.isSet = true
}

func (v NullableModAssignGetSubmissionStatus200ResponseLastattempt) IsSet() bool {
	return v.isSet
}

func (v *NullableModAssignGetSubmissionStatus200ResponseLastattempt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModAssignGetSubmissionStatus200ResponseLastattempt(val *ModAssignGetSubmissionStatus200ResponseLastattempt) *NullableModAssignGetSubmissionStatus200ResponseLastattempt {
	return &NullableModAssignGetSubmissionStatus200ResponseLastattempt{value: val, isSet: true}
}

func (v NullableModAssignGetSubmissionStatus200ResponseLastattempt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModAssignGetSubmissionStatus200ResponseLastattempt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


