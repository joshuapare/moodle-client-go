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

// checks if the ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan{}

// ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan struct for ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan
type ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan struct {
	// competencyid
	Competencyid int32 `json:"competencyid"`
	// grade
	Grade int32 `json:"grade"`
	// gradename
	Gradename string `json:"gradename"`
	// id
	Id int32 `json:"id"`
	// planid
	Planid int32 `json:"planid"`
	// proficiency
	Proficiency bool `json:"proficiency"`
	// proficiencyname
	Proficiencyname string `json:"proficiencyname"`
	// sortorder
	Sortorder int32 `json:"sortorder"`
	// timecreated
	Timecreated int32 `json:"timecreated"`
	// timemodified
	Timemodified int32 `json:"timemodified"`
	// userid
	Userid int32 `json:"userid"`
	// usermodified
	Usermodified int32 `json:"usermodified"`
}

type _ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan

// NewToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan instantiates a new ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan(competencyid int32, grade int32, gradename string, id int32, planid int32, proficiency bool, proficiencyname string, sortorder int32, timecreated int32, timemodified int32, userid int32, usermodified int32) *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan {
	this := ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan{}
	this.Competencyid = competencyid
	this.Grade = grade
	this.Gradename = gradename
	this.Id = id
	this.Planid = planid
	this.Proficiency = proficiency
	this.Proficiencyname = proficiencyname
	this.Sortorder = sortorder
	this.Timecreated = timecreated
	this.Timemodified = timemodified
	this.Userid = userid
	this.Usermodified = usermodified
	return &this
}

// NewToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplanWithDefaults instantiates a new ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplanWithDefaults() *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan {
	this := ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan{}
	var id int32 = 0
	this.Id = id
	var planid int32 = null
	this.Planid = planid
	var sortorder int32 = 0
	this.Sortorder = sortorder
	var timecreated int32 = 0
	this.Timecreated = timecreated
	var timemodified int32 = 0
	this.Timemodified = timemodified
	var usermodified int32 = 0
	this.Usermodified = usermodified
	return &this
}

// GetCompetencyid returns the Competencyid field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetCompetencyid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Competencyid
}

// GetCompetencyidOk returns a tuple with the Competencyid field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetCompetencyidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Competencyid, true
}

// SetCompetencyid sets field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) SetCompetencyid(v int32) {
	o.Competencyid = v
}

// GetGrade returns the Grade field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetGrade() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Grade
}

// GetGradeOk returns a tuple with the Grade field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetGradeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grade, true
}

// SetGrade sets field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) SetGrade(v int32) {
	o.Grade = v
}

// GetGradename returns the Gradename field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetGradename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gradename
}

// GetGradenameOk returns a tuple with the Gradename field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetGradenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gradename, true
}

// SetGradename sets field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) SetGradename(v string) {
	o.Gradename = v
}

// GetId returns the Id field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) SetId(v int32) {
	o.Id = v
}

// GetPlanid returns the Planid field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetPlanid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Planid
}

// GetPlanidOk returns a tuple with the Planid field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetPlanidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Planid, true
}

// SetPlanid sets field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) SetPlanid(v int32) {
	o.Planid = v
}

// GetProficiency returns the Proficiency field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetProficiency() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Proficiency
}

// GetProficiencyOk returns a tuple with the Proficiency field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetProficiencyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proficiency, true
}

// SetProficiency sets field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) SetProficiency(v bool) {
	o.Proficiency = v
}

// GetProficiencyname returns the Proficiencyname field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetProficiencyname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Proficiencyname
}

// GetProficiencynameOk returns a tuple with the Proficiencyname field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetProficiencynameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proficiencyname, true
}

// SetProficiencyname sets field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) SetProficiencyname(v string) {
	o.Proficiencyname = v
}

// GetSortorder returns the Sortorder field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetSortorder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sortorder
}

// GetSortorderOk returns a tuple with the Sortorder field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetSortorderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sortorder, true
}

// SetSortorder sets field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) SetSortorder(v int32) {
	o.Sortorder = v
}

// GetTimecreated returns the Timecreated field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetTimecreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timecreated
}

// GetTimecreatedOk returns a tuple with the Timecreated field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetTimecreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timecreated, true
}

// SetTimecreated sets field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) SetTimecreated(v int32) {
	o.Timecreated = v
}

// GetTimemodified returns the Timemodified field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetTimemodified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetTimemodifiedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timemodified, true
}

// SetTimemodified sets field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) SetTimemodified(v int32) {
	o.Timemodified = v
}

// GetUserid returns the Userid field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetUserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Userid
}

// GetUseridOk returns a tuple with the Userid field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetUseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userid, true
}

// SetUserid sets field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) SetUserid(v int32) {
	o.Userid = v
}

// GetUsermodified returns the Usermodified field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetUsermodified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Usermodified
}

// GetUsermodifiedOk returns a tuple with the Usermodified field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) GetUsermodifiedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usermodified, true
}

// SetUsermodified sets field value
func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) SetUsermodified(v int32) {
	o.Usermodified = v
}

func (o ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["competencyid"] = o.Competencyid
	toSerialize["grade"] = o.Grade
	toSerialize["gradename"] = o.Gradename
	toSerialize["id"] = o.Id
	toSerialize["planid"] = o.Planid
	toSerialize["proficiency"] = o.Proficiency
	toSerialize["proficiencyname"] = o.Proficiencyname
	toSerialize["sortorder"] = o.Sortorder
	toSerialize["timecreated"] = o.Timecreated
	toSerialize["timemodified"] = o.Timemodified
	toSerialize["userid"] = o.Userid
	toSerialize["usermodified"] = o.Usermodified
	return toSerialize, nil
}

func (o *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"competencyid",
		"grade",
		"gradename",
		"id",
		"planid",
		"proficiency",
		"proficiencyname",
		"sortorder",
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

	varToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan := _ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan)

	if err != nil {
		return err
	}

	*o = ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan(varToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan)

	return err
}

type NullableToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan struct {
	value *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan
	isSet bool
}

func (v NullableToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) Get() *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan {
	return v.value
}

func (v *NullableToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) Set(val *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) {
	v.value = val
	v.isSet = true
}

func (v NullableToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) IsSet() bool {
	return v.isSet
}

func (v *NullableToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan(val *ToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) *NullableToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan {
	return &NullableToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan{value: val, isSet: true}
}

func (v NullableToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolLpDataForPlanPage200ResponseCompetenciesInnerUsercompetencyplan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


