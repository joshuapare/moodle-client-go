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

// checks if the CoreCourseSearchCoursesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseSearchCoursesRequest{}

// CoreCourseSearchCoursesRequest struct for CoreCourseSearchCoursesRequest
type CoreCourseSearchCoursesRequest struct {
	// criteria name                                                         (search, modulelist (only admins), blocklist (only admins), tagid)
	Criterianame string `json:"criterianame"`
	// criteria value
	Criteriavalue string `json:"criteriavalue"`
	// limit to enrolled courses
	Limittoenrolled *bool `json:"limittoenrolled,omitempty"`
	// limit to courses where completion is enabled
	Onlywithcompletion *bool `json:"onlywithcompletion,omitempty"`
	// page number (0 based)
	Page *int32 `json:"page,omitempty"`
	// items per page
	Perpage *int32 `json:"perpage,omitempty"`
	Requiredcapabilities []map[string]interface{} `json:"requiredcapabilities,omitempty"`
}

type _CoreCourseSearchCoursesRequest CoreCourseSearchCoursesRequest

// NewCoreCourseSearchCoursesRequest instantiates a new CoreCourseSearchCoursesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseSearchCoursesRequest(criterianame string, criteriavalue string) *CoreCourseSearchCoursesRequest {
	this := CoreCourseSearchCoursesRequest{}
	this.Criterianame = criterianame
	this.Criteriavalue = criteriavalue
	var limittoenrolled bool = 0
	this.Limittoenrolled = &limittoenrolled
	var onlywithcompletion bool = 0
	this.Onlywithcompletion = &onlywithcompletion
	var page int32 = 0
	this.Page = &page
	var perpage int32 = 0
	this.Perpage = &perpage
	return &this
}

// NewCoreCourseSearchCoursesRequestWithDefaults instantiates a new CoreCourseSearchCoursesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseSearchCoursesRequestWithDefaults() *CoreCourseSearchCoursesRequest {
	this := CoreCourseSearchCoursesRequest{}
	var criterianame string = "null"
	this.Criterianame = criterianame
	var criteriavalue string = "null"
	this.Criteriavalue = criteriavalue
	var limittoenrolled bool = 0
	this.Limittoenrolled = &limittoenrolled
	var onlywithcompletion bool = 0
	this.Onlywithcompletion = &onlywithcompletion
	var page int32 = 0
	this.Page = &page
	var perpage int32 = 0
	this.Perpage = &perpage
	return &this
}

// GetCriterianame returns the Criterianame field value
func (o *CoreCourseSearchCoursesRequest) GetCriterianame() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Criterianame
}

// GetCriterianameOk returns a tuple with the Criterianame field value
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCoursesRequest) GetCriterianameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Criterianame, true
}

// SetCriterianame sets field value
func (o *CoreCourseSearchCoursesRequest) SetCriterianame(v string) {
	o.Criterianame = v
}

// GetCriteriavalue returns the Criteriavalue field value
func (o *CoreCourseSearchCoursesRequest) GetCriteriavalue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Criteriavalue
}

// GetCriteriavalueOk returns a tuple with the Criteriavalue field value
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCoursesRequest) GetCriteriavalueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Criteriavalue, true
}

// SetCriteriavalue sets field value
func (o *CoreCourseSearchCoursesRequest) SetCriteriavalue(v string) {
	o.Criteriavalue = v
}

// GetLimittoenrolled returns the Limittoenrolled field value if set, zero value otherwise.
func (o *CoreCourseSearchCoursesRequest) GetLimittoenrolled() bool {
	if o == nil || IsNil(o.Limittoenrolled) {
		var ret bool
		return ret
	}
	return *o.Limittoenrolled
}

// GetLimittoenrolledOk returns a tuple with the Limittoenrolled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCoursesRequest) GetLimittoenrolledOk() (*bool, bool) {
	if o == nil || IsNil(o.Limittoenrolled) {
		return nil, false
	}
	return o.Limittoenrolled, true
}

// HasLimittoenrolled returns a boolean if a field has been set.
func (o *CoreCourseSearchCoursesRequest) HasLimittoenrolled() bool {
	if o != nil && !IsNil(o.Limittoenrolled) {
		return true
	}

	return false
}

// SetLimittoenrolled gets a reference to the given bool and assigns it to the Limittoenrolled field.
func (o *CoreCourseSearchCoursesRequest) SetLimittoenrolled(v bool) {
	o.Limittoenrolled = &v
}

// GetOnlywithcompletion returns the Onlywithcompletion field value if set, zero value otherwise.
func (o *CoreCourseSearchCoursesRequest) GetOnlywithcompletion() bool {
	if o == nil || IsNil(o.Onlywithcompletion) {
		var ret bool
		return ret
	}
	return *o.Onlywithcompletion
}

// GetOnlywithcompletionOk returns a tuple with the Onlywithcompletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCoursesRequest) GetOnlywithcompletionOk() (*bool, bool) {
	if o == nil || IsNil(o.Onlywithcompletion) {
		return nil, false
	}
	return o.Onlywithcompletion, true
}

// HasOnlywithcompletion returns a boolean if a field has been set.
func (o *CoreCourseSearchCoursesRequest) HasOnlywithcompletion() bool {
	if o != nil && !IsNil(o.Onlywithcompletion) {
		return true
	}

	return false
}

// SetOnlywithcompletion gets a reference to the given bool and assigns it to the Onlywithcompletion field.
func (o *CoreCourseSearchCoursesRequest) SetOnlywithcompletion(v bool) {
	o.Onlywithcompletion = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *CoreCourseSearchCoursesRequest) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCoursesRequest) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *CoreCourseSearchCoursesRequest) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *CoreCourseSearchCoursesRequest) SetPage(v int32) {
	o.Page = &v
}

// GetPerpage returns the Perpage field value if set, zero value otherwise.
func (o *CoreCourseSearchCoursesRequest) GetPerpage() int32 {
	if o == nil || IsNil(o.Perpage) {
		var ret int32
		return ret
	}
	return *o.Perpage
}

// GetPerpageOk returns a tuple with the Perpage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCoursesRequest) GetPerpageOk() (*int32, bool) {
	if o == nil || IsNil(o.Perpage) {
		return nil, false
	}
	return o.Perpage, true
}

// HasPerpage returns a boolean if a field has been set.
func (o *CoreCourseSearchCoursesRequest) HasPerpage() bool {
	if o != nil && !IsNil(o.Perpage) {
		return true
	}

	return false
}

// SetPerpage gets a reference to the given int32 and assigns it to the Perpage field.
func (o *CoreCourseSearchCoursesRequest) SetPerpage(v int32) {
	o.Perpage = &v
}

// GetRequiredcapabilities returns the Requiredcapabilities field value if set, zero value otherwise.
func (o *CoreCourseSearchCoursesRequest) GetRequiredcapabilities() []map[string]interface{} {
	if o == nil || IsNil(o.Requiredcapabilities) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Requiredcapabilities
}

// GetRequiredcapabilitiesOk returns a tuple with the Requiredcapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCoursesRequest) GetRequiredcapabilitiesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Requiredcapabilities) {
		return nil, false
	}
	return o.Requiredcapabilities, true
}

// HasRequiredcapabilities returns a boolean if a field has been set.
func (o *CoreCourseSearchCoursesRequest) HasRequiredcapabilities() bool {
	if o != nil && !IsNil(o.Requiredcapabilities) {
		return true
	}

	return false
}

// SetRequiredcapabilities gets a reference to the given []map[string]interface{} and assigns it to the Requiredcapabilities field.
func (o *CoreCourseSearchCoursesRequest) SetRequiredcapabilities(v []map[string]interface{}) {
	o.Requiredcapabilities = v
}

func (o CoreCourseSearchCoursesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseSearchCoursesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["criterianame"] = o.Criterianame
	toSerialize["criteriavalue"] = o.Criteriavalue
	if !IsNil(o.Limittoenrolled) {
		toSerialize["limittoenrolled"] = o.Limittoenrolled
	}
	if !IsNil(o.Onlywithcompletion) {
		toSerialize["onlywithcompletion"] = o.Onlywithcompletion
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Perpage) {
		toSerialize["perpage"] = o.Perpage
	}
	if !IsNil(o.Requiredcapabilities) {
		toSerialize["requiredcapabilities"] = o.Requiredcapabilities
	}
	return toSerialize, nil
}

func (o *CoreCourseSearchCoursesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"criterianame",
		"criteriavalue",
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

	varCoreCourseSearchCoursesRequest := _CoreCourseSearchCoursesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCourseSearchCoursesRequest)

	if err != nil {
		return err
	}

	*o = CoreCourseSearchCoursesRequest(varCoreCourseSearchCoursesRequest)

	return err
}

type NullableCoreCourseSearchCoursesRequest struct {
	value *CoreCourseSearchCoursesRequest
	isSet bool
}

func (v NullableCoreCourseSearchCoursesRequest) Get() *CoreCourseSearchCoursesRequest {
	return v.value
}

func (v *NullableCoreCourseSearchCoursesRequest) Set(val *CoreCourseSearchCoursesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseSearchCoursesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseSearchCoursesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseSearchCoursesRequest(val *CoreCourseSearchCoursesRequest) *NullableCoreCourseSearchCoursesRequest {
	return &NullableCoreCourseSearchCoursesRequest{value: val, isSet: true}
}

func (v NullableCoreCourseSearchCoursesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseSearchCoursesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


