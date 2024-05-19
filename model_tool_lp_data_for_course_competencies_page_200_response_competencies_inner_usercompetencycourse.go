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

// checks if the ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse{}

// ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse struct for ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse
type ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse struct {
	// competencyid
	Competencyid int32 `json:"competencyid"`
	// courseid
	Courseid int32 `json:"courseid"`
	// grade
	Grade int32 `json:"grade"`
	// gradename
	Gradename string `json:"gradename"`
	// id
	Id int32 `json:"id"`
	// proficiency
	Proficiency bool `json:"proficiency"`
	// proficiencyname
	Proficiencyname string `json:"proficiencyname"`
	// timecreated
	Timecreated int32 `json:"timecreated"`
	// timemodified
	Timemodified int32 `json:"timemodified"`
	// userid
	Userid int32 `json:"userid"`
	// usermodified
	Usermodified int32 `json:"usermodified"`
}

type _ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse

// NewToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse instantiates a new ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse(competencyid int32, courseid int32, grade int32, gradename string, id int32, proficiency bool, proficiencyname string, timecreated int32, timemodified int32, userid int32, usermodified int32) *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse {
	this := ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse{}
	this.Competencyid = competencyid
	this.Courseid = courseid
	this.Grade = grade
	this.Gradename = gradename
	this.Id = id
	this.Proficiency = proficiency
	this.Proficiencyname = proficiencyname
	this.Timecreated = timecreated
	this.Timemodified = timemodified
	this.Userid = userid
	this.Usermodified = usermodified
	return &this
}

// NewToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourseWithDefaults instantiates a new ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourseWithDefaults() *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse {
	this := ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse{}
	var id int32 = 0
	this.Id = id
	var timecreated int32 = 0
	this.Timecreated = timecreated
	var timemodified int32 = 0
	this.Timemodified = timemodified
	var usermodified int32 = 0
	this.Usermodified = usermodified
	return &this
}

// GetCompetencyid returns the Competencyid field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetCompetencyid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Competencyid
}

// GetCompetencyidOk returns a tuple with the Competencyid field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetCompetencyidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Competencyid, true
}

// SetCompetencyid sets field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) SetCompetencyid(v int32) {
	o.Competencyid = v
}

// GetCourseid returns the Courseid field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetCourseid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetCourseidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Courseid, true
}

// SetCourseid sets field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) SetCourseid(v int32) {
	o.Courseid = v
}

// GetGrade returns the Grade field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetGrade() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Grade
}

// GetGradeOk returns a tuple with the Grade field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetGradeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grade, true
}

// SetGrade sets field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) SetGrade(v int32) {
	o.Grade = v
}

// GetGradename returns the Gradename field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetGradename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gradename
}

// GetGradenameOk returns a tuple with the Gradename field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetGradenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gradename, true
}

// SetGradename sets field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) SetGradename(v string) {
	o.Gradename = v
}

// GetId returns the Id field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) SetId(v int32) {
	o.Id = v
}

// GetProficiency returns the Proficiency field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetProficiency() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Proficiency
}

// GetProficiencyOk returns a tuple with the Proficiency field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetProficiencyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proficiency, true
}

// SetProficiency sets field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) SetProficiency(v bool) {
	o.Proficiency = v
}

// GetProficiencyname returns the Proficiencyname field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetProficiencyname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Proficiencyname
}

// GetProficiencynameOk returns a tuple with the Proficiencyname field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetProficiencynameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proficiencyname, true
}

// SetProficiencyname sets field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) SetProficiencyname(v string) {
	o.Proficiencyname = v
}

// GetTimecreated returns the Timecreated field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetTimecreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timecreated
}

// GetTimecreatedOk returns a tuple with the Timecreated field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetTimecreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timecreated, true
}

// SetTimecreated sets field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) SetTimecreated(v int32) {
	o.Timecreated = v
}

// GetTimemodified returns the Timemodified field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetTimemodified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetTimemodifiedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timemodified, true
}

// SetTimemodified sets field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) SetTimemodified(v int32) {
	o.Timemodified = v
}

// GetUserid returns the Userid field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetUserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Userid
}

// GetUseridOk returns a tuple with the Userid field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetUseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userid, true
}

// SetUserid sets field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) SetUserid(v int32) {
	o.Userid = v
}

// GetUsermodified returns the Usermodified field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetUsermodified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Usermodified
}

// GetUsermodifiedOk returns a tuple with the Usermodified field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) GetUsermodifiedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usermodified, true
}

// SetUsermodified sets field value
func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) SetUsermodified(v int32) {
	o.Usermodified = v
}

func (o ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["competencyid"] = o.Competencyid
	toSerialize["courseid"] = o.Courseid
	toSerialize["grade"] = o.Grade
	toSerialize["gradename"] = o.Gradename
	toSerialize["id"] = o.Id
	toSerialize["proficiency"] = o.Proficiency
	toSerialize["proficiencyname"] = o.Proficiencyname
	toSerialize["timecreated"] = o.Timecreated
	toSerialize["timemodified"] = o.Timemodified
	toSerialize["userid"] = o.Userid
	toSerialize["usermodified"] = o.Usermodified
	return toSerialize, nil
}

func (o *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"competencyid",
		"courseid",
		"grade",
		"gradename",
		"id",
		"proficiency",
		"proficiencyname",
		"timecreated",
		"timemodified",
		"userid",
		"usermodified",
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

	varToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse := _ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse)

	if err != nil {
		return err
	}

	*o = ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse(varToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse)

	return err
}

type NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse struct {
	value *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse
	isSet bool
}

func (v NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) Get() *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse {
	return v.value
}

func (v *NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) Set(val *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) {
	v.value = val
	v.isSet = true
}

func (v NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) IsSet() bool {
	return v.isSet
}

func (v *NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse(val *ToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) *NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse {
	return &NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse{value: val, isSet: true}
}

func (v NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolLpDataForCourseCompetenciesPage200ResponseCompetenciesInnerUsercompetencycourse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


