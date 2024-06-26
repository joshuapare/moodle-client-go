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

// checks if the ToolLpDataForTemplateCompetenciesPage200ResponseStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolLpDataForTemplateCompetenciesPage200ResponseStatistics{}

// ToolLpDataForTemplateCompetenciesPage200ResponseStatistics struct for ToolLpDataForTemplateCompetenciesPage200ResponseStatistics
type ToolLpDataForTemplateCompetenciesPage200ResponseStatistics struct {
	// competencycount
	Competencycount int32 `json:"competencycount"`
	// completedplancount
	Completedplancount int32 `json:"completedplancount"`
	// completedplanpercentage
	Completedplanpercentage float32 `json:"completedplanpercentage"`
	// completedplanpercentageformatted
	Completedplanpercentageformatted string `json:"completedplanpercentageformatted"`
	Leastproficient []CoreCompetencyReadUserEvidence200ResponseCompetenciesInner `json:"leastproficient"`
	// leastproficientcount
	Leastproficientcount int32 `json:"leastproficientcount"`
	// linkedcompetencycount
	Linkedcompetencycount int32 `json:"linkedcompetencycount"`
	// linkedcompetencypercentage
	Linkedcompetencypercentage float32 `json:"linkedcompetencypercentage"`
	// linkedcompetencypercentageformatted
	Linkedcompetencypercentageformatted string `json:"linkedcompetencypercentageformatted"`
	// plancount
	Plancount int32 `json:"plancount"`
	// proficientusercompetencyplancount
	Proficientusercompetencyplancount int32 `json:"proficientusercompetencyplancount"`
	// proficientusercompetencyplanpercentage
	Proficientusercompetencyplanpercentage float32 `json:"proficientusercompetencyplanpercentage"`
	// proficientusercompetencyplanpercentageformatted
	Proficientusercompetencyplanpercentageformatted string `json:"proficientusercompetencyplanpercentageformatted"`
	// unlinkedcompetencycount
	Unlinkedcompetencycount int32 `json:"unlinkedcompetencycount"`
	// usercompetencyplancount
	Usercompetencyplancount int32 `json:"usercompetencyplancount"`
}

type _ToolLpDataForTemplateCompetenciesPage200ResponseStatistics ToolLpDataForTemplateCompetenciesPage200ResponseStatistics

// NewToolLpDataForTemplateCompetenciesPage200ResponseStatistics instantiates a new ToolLpDataForTemplateCompetenciesPage200ResponseStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolLpDataForTemplateCompetenciesPage200ResponseStatistics(competencycount int32, completedplancount int32, completedplanpercentage float32, completedplanpercentageformatted string, leastproficient []CoreCompetencyReadUserEvidence200ResponseCompetenciesInner, leastproficientcount int32, linkedcompetencycount int32, linkedcompetencypercentage float32, linkedcompetencypercentageformatted string, plancount int32, proficientusercompetencyplancount int32, proficientusercompetencyplanpercentage float32, proficientusercompetencyplanpercentageformatted string, unlinkedcompetencycount int32, usercompetencyplancount int32) *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics {
	this := ToolLpDataForTemplateCompetenciesPage200ResponseStatistics{}
	this.Competencycount = competencycount
	this.Completedplancount = completedplancount
	this.Completedplanpercentage = completedplanpercentage
	this.Completedplanpercentageformatted = completedplanpercentageformatted
	this.Leastproficient = leastproficient
	this.Leastproficientcount = leastproficientcount
	this.Linkedcompetencycount = linkedcompetencycount
	this.Linkedcompetencypercentage = linkedcompetencypercentage
	this.Linkedcompetencypercentageformatted = linkedcompetencypercentageformatted
	this.Plancount = plancount
	this.Proficientusercompetencyplancount = proficientusercompetencyplancount
	this.Proficientusercompetencyplanpercentage = proficientusercompetencyplanpercentage
	this.Proficientusercompetencyplanpercentageformatted = proficientusercompetencyplanpercentageformatted
	this.Unlinkedcompetencycount = unlinkedcompetencycount
	this.Usercompetencyplancount = usercompetencyplancount
	return &this
}

// NewToolLpDataForTemplateCompetenciesPage200ResponseStatisticsWithDefaults instantiates a new ToolLpDataForTemplateCompetenciesPage200ResponseStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolLpDataForTemplateCompetenciesPage200ResponseStatisticsWithDefaults() *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics {
	this := ToolLpDataForTemplateCompetenciesPage200ResponseStatistics{}
	var completedplancount int32 = null
	this.Completedplancount = completedplancount
	var completedplanpercentage float32 = null
	this.Completedplanpercentage = completedplanpercentage
	var completedplanpercentageformatted string = "null"
	this.Completedplanpercentageformatted = completedplanpercentageformatted
	var linkedcompetencycount int32 = null
	this.Linkedcompetencycount = linkedcompetencycount
	var linkedcompetencypercentage float32 = null
	this.Linkedcompetencypercentage = linkedcompetencypercentage
	var linkedcompetencypercentageformatted string = "null"
	this.Linkedcompetencypercentageformatted = linkedcompetencypercentageformatted
	var plancount int32 = null
	this.Plancount = plancount
	var proficientusercompetencyplancount int32 = null
	this.Proficientusercompetencyplancount = proficientusercompetencyplancount
	var proficientusercompetencyplanpercentage float32 = null
	this.Proficientusercompetencyplanpercentage = proficientusercompetencyplanpercentage
	var proficientusercompetencyplanpercentageformatted string = "null"
	this.Proficientusercompetencyplanpercentageformatted = proficientusercompetencyplanpercentageformatted
	var unlinkedcompetencycount int32 = null
	this.Unlinkedcompetencycount = unlinkedcompetencycount
	var usercompetencyplancount int32 = null
	this.Usercompetencyplancount = usercompetencyplancount
	return &this
}

// GetCompetencycount returns the Competencycount field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetCompetencycount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Competencycount
}

// GetCompetencycountOk returns a tuple with the Competencycount field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetCompetencycountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Competencycount, true
}

// SetCompetencycount sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetCompetencycount(v int32) {
	o.Competencycount = v
}

// GetCompletedplancount returns the Completedplancount field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetCompletedplancount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Completedplancount
}

// GetCompletedplancountOk returns a tuple with the Completedplancount field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetCompletedplancountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completedplancount, true
}

// SetCompletedplancount sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetCompletedplancount(v int32) {
	o.Completedplancount = v
}

// GetCompletedplanpercentage returns the Completedplanpercentage field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetCompletedplanpercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Completedplanpercentage
}

// GetCompletedplanpercentageOk returns a tuple with the Completedplanpercentage field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetCompletedplanpercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completedplanpercentage, true
}

// SetCompletedplanpercentage sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetCompletedplanpercentage(v float32) {
	o.Completedplanpercentage = v
}

// GetCompletedplanpercentageformatted returns the Completedplanpercentageformatted field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetCompletedplanpercentageformatted() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Completedplanpercentageformatted
}

// GetCompletedplanpercentageformattedOk returns a tuple with the Completedplanpercentageformatted field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetCompletedplanpercentageformattedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completedplanpercentageformatted, true
}

// SetCompletedplanpercentageformatted sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetCompletedplanpercentageformatted(v string) {
	o.Completedplanpercentageformatted = v
}

// GetLeastproficient returns the Leastproficient field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetLeastproficient() []CoreCompetencyReadUserEvidence200ResponseCompetenciesInner {
	if o == nil {
		var ret []CoreCompetencyReadUserEvidence200ResponseCompetenciesInner
		return ret
	}

	return o.Leastproficient
}

// GetLeastproficientOk returns a tuple with the Leastproficient field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetLeastproficientOk() ([]CoreCompetencyReadUserEvidence200ResponseCompetenciesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Leastproficient, true
}

// SetLeastproficient sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetLeastproficient(v []CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) {
	o.Leastproficient = v
}

// GetLeastproficientcount returns the Leastproficientcount field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetLeastproficientcount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Leastproficientcount
}

// GetLeastproficientcountOk returns a tuple with the Leastproficientcount field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetLeastproficientcountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Leastproficientcount, true
}

// SetLeastproficientcount sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetLeastproficientcount(v int32) {
	o.Leastproficientcount = v
}

// GetLinkedcompetencycount returns the Linkedcompetencycount field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetLinkedcompetencycount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Linkedcompetencycount
}

// GetLinkedcompetencycountOk returns a tuple with the Linkedcompetencycount field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetLinkedcompetencycountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Linkedcompetencycount, true
}

// SetLinkedcompetencycount sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetLinkedcompetencycount(v int32) {
	o.Linkedcompetencycount = v
}

// GetLinkedcompetencypercentage returns the Linkedcompetencypercentage field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetLinkedcompetencypercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Linkedcompetencypercentage
}

// GetLinkedcompetencypercentageOk returns a tuple with the Linkedcompetencypercentage field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetLinkedcompetencypercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Linkedcompetencypercentage, true
}

// SetLinkedcompetencypercentage sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetLinkedcompetencypercentage(v float32) {
	o.Linkedcompetencypercentage = v
}

// GetLinkedcompetencypercentageformatted returns the Linkedcompetencypercentageformatted field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetLinkedcompetencypercentageformatted() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Linkedcompetencypercentageformatted
}

// GetLinkedcompetencypercentageformattedOk returns a tuple with the Linkedcompetencypercentageformatted field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetLinkedcompetencypercentageformattedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Linkedcompetencypercentageformatted, true
}

// SetLinkedcompetencypercentageformatted sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetLinkedcompetencypercentageformatted(v string) {
	o.Linkedcompetencypercentageformatted = v
}

// GetPlancount returns the Plancount field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetPlancount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Plancount
}

// GetPlancountOk returns a tuple with the Plancount field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetPlancountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plancount, true
}

// SetPlancount sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetPlancount(v int32) {
	o.Plancount = v
}

// GetProficientusercompetencyplancount returns the Proficientusercompetencyplancount field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetProficientusercompetencyplancount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Proficientusercompetencyplancount
}

// GetProficientusercompetencyplancountOk returns a tuple with the Proficientusercompetencyplancount field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetProficientusercompetencyplancountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proficientusercompetencyplancount, true
}

// SetProficientusercompetencyplancount sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetProficientusercompetencyplancount(v int32) {
	o.Proficientusercompetencyplancount = v
}

// GetProficientusercompetencyplanpercentage returns the Proficientusercompetencyplanpercentage field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetProficientusercompetencyplanpercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Proficientusercompetencyplanpercentage
}

// GetProficientusercompetencyplanpercentageOk returns a tuple with the Proficientusercompetencyplanpercentage field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetProficientusercompetencyplanpercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proficientusercompetencyplanpercentage, true
}

// SetProficientusercompetencyplanpercentage sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetProficientusercompetencyplanpercentage(v float32) {
	o.Proficientusercompetencyplanpercentage = v
}

// GetProficientusercompetencyplanpercentageformatted returns the Proficientusercompetencyplanpercentageformatted field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetProficientusercompetencyplanpercentageformatted() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Proficientusercompetencyplanpercentageformatted
}

// GetProficientusercompetencyplanpercentageformattedOk returns a tuple with the Proficientusercompetencyplanpercentageformatted field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetProficientusercompetencyplanpercentageformattedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proficientusercompetencyplanpercentageformatted, true
}

// SetProficientusercompetencyplanpercentageformatted sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetProficientusercompetencyplanpercentageformatted(v string) {
	o.Proficientusercompetencyplanpercentageformatted = v
}

// GetUnlinkedcompetencycount returns the Unlinkedcompetencycount field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetUnlinkedcompetencycount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Unlinkedcompetencycount
}

// GetUnlinkedcompetencycountOk returns a tuple with the Unlinkedcompetencycount field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetUnlinkedcompetencycountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unlinkedcompetencycount, true
}

// SetUnlinkedcompetencycount sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetUnlinkedcompetencycount(v int32) {
	o.Unlinkedcompetencycount = v
}

// GetUsercompetencyplancount returns the Usercompetencyplancount field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetUsercompetencyplancount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Usercompetencyplancount
}

// GetUsercompetencyplancountOk returns a tuple with the Usercompetencyplancount field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) GetUsercompetencyplancountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usercompetencyplancount, true
}

// SetUsercompetencyplancount sets field value
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) SetUsercompetencyplancount(v int32) {
	o.Usercompetencyplancount = v
}

func (o ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["competencycount"] = o.Competencycount
	toSerialize["completedplancount"] = o.Completedplancount
	toSerialize["completedplanpercentage"] = o.Completedplanpercentage
	toSerialize["completedplanpercentageformatted"] = o.Completedplanpercentageformatted
	toSerialize["leastproficient"] = o.Leastproficient
	toSerialize["leastproficientcount"] = o.Leastproficientcount
	toSerialize["linkedcompetencycount"] = o.Linkedcompetencycount
	toSerialize["linkedcompetencypercentage"] = o.Linkedcompetencypercentage
	toSerialize["linkedcompetencypercentageformatted"] = o.Linkedcompetencypercentageformatted
	toSerialize["plancount"] = o.Plancount
	toSerialize["proficientusercompetencyplancount"] = o.Proficientusercompetencyplancount
	toSerialize["proficientusercompetencyplanpercentage"] = o.Proficientusercompetencyplanpercentage
	toSerialize["proficientusercompetencyplanpercentageformatted"] = o.Proficientusercompetencyplanpercentageformatted
	toSerialize["unlinkedcompetencycount"] = o.Unlinkedcompetencycount
	toSerialize["usercompetencyplancount"] = o.Usercompetencyplancount
	return toSerialize, nil
}

func (o *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"competencycount",
		"completedplancount",
		"completedplanpercentage",
		"completedplanpercentageformatted",
		"leastproficient",
		"leastproficientcount",
		"linkedcompetencycount",
		"linkedcompetencypercentage",
		"linkedcompetencypercentageformatted",
		"plancount",
		"proficientusercompetencyplancount",
		"proficientusercompetencyplanpercentage",
		"proficientusercompetencyplanpercentageformatted",
		"unlinkedcompetencycount",
		"usercompetencyplancount",
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

	varToolLpDataForTemplateCompetenciesPage200ResponseStatistics := _ToolLpDataForTemplateCompetenciesPage200ResponseStatistics{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolLpDataForTemplateCompetenciesPage200ResponseStatistics)

	if err != nil {
		return err
	}

	*o = ToolLpDataForTemplateCompetenciesPage200ResponseStatistics(varToolLpDataForTemplateCompetenciesPage200ResponseStatistics)

	return err
}

type NullableToolLpDataForTemplateCompetenciesPage200ResponseStatistics struct {
	value *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics
	isSet bool
}

func (v NullableToolLpDataForTemplateCompetenciesPage200ResponseStatistics) Get() *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics {
	return v.value
}

func (v *NullableToolLpDataForTemplateCompetenciesPage200ResponseStatistics) Set(val *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableToolLpDataForTemplateCompetenciesPage200ResponseStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableToolLpDataForTemplateCompetenciesPage200ResponseStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolLpDataForTemplateCompetenciesPage200ResponseStatistics(val *ToolLpDataForTemplateCompetenciesPage200ResponseStatistics) *NullableToolLpDataForTemplateCompetenciesPage200ResponseStatistics {
	return &NullableToolLpDataForTemplateCompetenciesPage200ResponseStatistics{value: val, isSet: true}
}

func (v NullableToolLpDataForTemplateCompetenciesPage200ResponseStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolLpDataForTemplateCompetenciesPage200ResponseStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


