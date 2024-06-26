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

// checks if the CoreGradesUpdateGradesRequestItemdetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGradesUpdateGradesRequestItemdetails{}

// CoreGradesUpdateGradesRequestItemdetails struct for CoreGradesUpdateGradesRequestItemdetails
type CoreGradesUpdateGradesRequestItemdetails struct {
	// True if the grade item should be deleted
	Deleted *bool `json:"deleted,omitempty"`
	// Maximum grade allowed
	Grademax *float32 `json:"grademax,omitempty"`
	// Minimum grade allowed
	Grademin *float32 `json:"grademin,omitempty"`
	// The type of grade (0 = none, 1 = value, 2 = scale, 3 = text)
	Gradetype *int32 `json:"gradetype,omitempty"`
	// True if the grade item is hidden
	Hidden *bool `json:"hidden,omitempty"`
	// Arbitrary ID provided by the module responsible for the grade item
	Idnumber *int32 `json:"idnumber,omitempty"`
	// The grade item name
	Itemname *string `json:"itemname,omitempty"`
	// Multiply all grades by this number
	Multfactor *float32 `json:"multfactor,omitempty"`
	// Add this to all grades
	Plusfactor *float32 `json:"plusfactor,omitempty"`
	// The ID of the custom scale being is used
	Scaleid *int32 `json:"scaleid,omitempty"`
}

// NewCoreGradesUpdateGradesRequestItemdetails instantiates a new CoreGradesUpdateGradesRequestItemdetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGradesUpdateGradesRequestItemdetails() *CoreGradesUpdateGradesRequestItemdetails {
	this := CoreGradesUpdateGradesRequestItemdetails{}
	var deleted bool = null
	this.Deleted = &deleted
	var grademax float32 = null
	this.Grademax = &grademax
	var grademin float32 = null
	this.Grademin = &grademin
	var gradetype int32 = null
	this.Gradetype = &gradetype
	var hidden bool = null
	this.Hidden = &hidden
	var idnumber int32 = null
	this.Idnumber = &idnumber
	var itemname string = "null"
	this.Itemname = &itemname
	var multfactor float32 = null
	this.Multfactor = &multfactor
	var plusfactor float32 = null
	this.Plusfactor = &plusfactor
	var scaleid int32 = null
	this.Scaleid = &scaleid
	return &this
}

// NewCoreGradesUpdateGradesRequestItemdetailsWithDefaults instantiates a new CoreGradesUpdateGradesRequestItemdetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGradesUpdateGradesRequestItemdetailsWithDefaults() *CoreGradesUpdateGradesRequestItemdetails {
	this := CoreGradesUpdateGradesRequestItemdetails{}
	var deleted bool = null
	this.Deleted = &deleted
	var grademax float32 = null
	this.Grademax = &grademax
	var grademin float32 = null
	this.Grademin = &grademin
	var gradetype int32 = null
	this.Gradetype = &gradetype
	var hidden bool = null
	this.Hidden = &hidden
	var idnumber int32 = null
	this.Idnumber = &idnumber
	var itemname string = "null"
	this.Itemname = &itemname
	var multfactor float32 = null
	this.Multfactor = &multfactor
	var plusfactor float32 = null
	this.Plusfactor = &plusfactor
	var scaleid int32 = null
	this.Scaleid = &scaleid
	return &this
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *CoreGradesUpdateGradesRequestItemdetails) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetGrademax returns the Grademax field value if set, zero value otherwise.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetGrademax() float32 {
	if o == nil || IsNil(o.Grademax) {
		var ret float32
		return ret
	}
	return *o.Grademax
}

// GetGrademaxOk returns a tuple with the Grademax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetGrademaxOk() (*float32, bool) {
	if o == nil || IsNil(o.Grademax) {
		return nil, false
	}
	return o.Grademax, true
}

// HasGrademax returns a boolean if a field has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) HasGrademax() bool {
	if o != nil && !IsNil(o.Grademax) {
		return true
	}

	return false
}

// SetGrademax gets a reference to the given float32 and assigns it to the Grademax field.
func (o *CoreGradesUpdateGradesRequestItemdetails) SetGrademax(v float32) {
	o.Grademax = &v
}

// GetGrademin returns the Grademin field value if set, zero value otherwise.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetGrademin() float32 {
	if o == nil || IsNil(o.Grademin) {
		var ret float32
		return ret
	}
	return *o.Grademin
}

// GetGrademinOk returns a tuple with the Grademin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetGrademinOk() (*float32, bool) {
	if o == nil || IsNil(o.Grademin) {
		return nil, false
	}
	return o.Grademin, true
}

// HasGrademin returns a boolean if a field has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) HasGrademin() bool {
	if o != nil && !IsNil(o.Grademin) {
		return true
	}

	return false
}

// SetGrademin gets a reference to the given float32 and assigns it to the Grademin field.
func (o *CoreGradesUpdateGradesRequestItemdetails) SetGrademin(v float32) {
	o.Grademin = &v
}

// GetGradetype returns the Gradetype field value if set, zero value otherwise.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetGradetype() int32 {
	if o == nil || IsNil(o.Gradetype) {
		var ret int32
		return ret
	}
	return *o.Gradetype
}

// GetGradetypeOk returns a tuple with the Gradetype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetGradetypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Gradetype) {
		return nil, false
	}
	return o.Gradetype, true
}

// HasGradetype returns a boolean if a field has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) HasGradetype() bool {
	if o != nil && !IsNil(o.Gradetype) {
		return true
	}

	return false
}

// SetGradetype gets a reference to the given int32 and assigns it to the Gradetype field.
func (o *CoreGradesUpdateGradesRequestItemdetails) SetGradetype(v int32) {
	o.Gradetype = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *CoreGradesUpdateGradesRequestItemdetails) SetHidden(v bool) {
	o.Hidden = &v
}

// GetIdnumber returns the Idnumber field value if set, zero value otherwise.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetIdnumber() int32 {
	if o == nil || IsNil(o.Idnumber) {
		var ret int32
		return ret
	}
	return *o.Idnumber
}

// GetIdnumberOk returns a tuple with the Idnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetIdnumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Idnumber) {
		return nil, false
	}
	return o.Idnumber, true
}

// HasIdnumber returns a boolean if a field has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) HasIdnumber() bool {
	if o != nil && !IsNil(o.Idnumber) {
		return true
	}

	return false
}

// SetIdnumber gets a reference to the given int32 and assigns it to the Idnumber field.
func (o *CoreGradesUpdateGradesRequestItemdetails) SetIdnumber(v int32) {
	o.Idnumber = &v
}

// GetItemname returns the Itemname field value if set, zero value otherwise.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetItemname() string {
	if o == nil || IsNil(o.Itemname) {
		var ret string
		return ret
	}
	return *o.Itemname
}

// GetItemnameOk returns a tuple with the Itemname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetItemnameOk() (*string, bool) {
	if o == nil || IsNil(o.Itemname) {
		return nil, false
	}
	return o.Itemname, true
}

// HasItemname returns a boolean if a field has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) HasItemname() bool {
	if o != nil && !IsNil(o.Itemname) {
		return true
	}

	return false
}

// SetItemname gets a reference to the given string and assigns it to the Itemname field.
func (o *CoreGradesUpdateGradesRequestItemdetails) SetItemname(v string) {
	o.Itemname = &v
}

// GetMultfactor returns the Multfactor field value if set, zero value otherwise.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetMultfactor() float32 {
	if o == nil || IsNil(o.Multfactor) {
		var ret float32
		return ret
	}
	return *o.Multfactor
}

// GetMultfactorOk returns a tuple with the Multfactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetMultfactorOk() (*float32, bool) {
	if o == nil || IsNil(o.Multfactor) {
		return nil, false
	}
	return o.Multfactor, true
}

// HasMultfactor returns a boolean if a field has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) HasMultfactor() bool {
	if o != nil && !IsNil(o.Multfactor) {
		return true
	}

	return false
}

// SetMultfactor gets a reference to the given float32 and assigns it to the Multfactor field.
func (o *CoreGradesUpdateGradesRequestItemdetails) SetMultfactor(v float32) {
	o.Multfactor = &v
}

// GetPlusfactor returns the Plusfactor field value if set, zero value otherwise.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetPlusfactor() float32 {
	if o == nil || IsNil(o.Plusfactor) {
		var ret float32
		return ret
	}
	return *o.Plusfactor
}

// GetPlusfactorOk returns a tuple with the Plusfactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetPlusfactorOk() (*float32, bool) {
	if o == nil || IsNil(o.Plusfactor) {
		return nil, false
	}
	return o.Plusfactor, true
}

// HasPlusfactor returns a boolean if a field has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) HasPlusfactor() bool {
	if o != nil && !IsNil(o.Plusfactor) {
		return true
	}

	return false
}

// SetPlusfactor gets a reference to the given float32 and assigns it to the Plusfactor field.
func (o *CoreGradesUpdateGradesRequestItemdetails) SetPlusfactor(v float32) {
	o.Plusfactor = &v
}

// GetScaleid returns the Scaleid field value if set, zero value otherwise.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetScaleid() int32 {
	if o == nil || IsNil(o.Scaleid) {
		var ret int32
		return ret
	}
	return *o.Scaleid
}

// GetScaleidOk returns a tuple with the Scaleid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) GetScaleidOk() (*int32, bool) {
	if o == nil || IsNil(o.Scaleid) {
		return nil, false
	}
	return o.Scaleid, true
}

// HasScaleid returns a boolean if a field has been set.
func (o *CoreGradesUpdateGradesRequestItemdetails) HasScaleid() bool {
	if o != nil && !IsNil(o.Scaleid) {
		return true
	}

	return false
}

// SetScaleid gets a reference to the given int32 and assigns it to the Scaleid field.
func (o *CoreGradesUpdateGradesRequestItemdetails) SetScaleid(v int32) {
	o.Scaleid = &v
}

func (o CoreGradesUpdateGradesRequestItemdetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGradesUpdateGradesRequestItemdetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !IsNil(o.Grademax) {
		toSerialize["grademax"] = o.Grademax
	}
	if !IsNil(o.Grademin) {
		toSerialize["grademin"] = o.Grademin
	}
	if !IsNil(o.Gradetype) {
		toSerialize["gradetype"] = o.Gradetype
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.Idnumber) {
		toSerialize["idnumber"] = o.Idnumber
	}
	if !IsNil(o.Itemname) {
		toSerialize["itemname"] = o.Itemname
	}
	if !IsNil(o.Multfactor) {
		toSerialize["multfactor"] = o.Multfactor
	}
	if !IsNil(o.Plusfactor) {
		toSerialize["plusfactor"] = o.Plusfactor
	}
	if !IsNil(o.Scaleid) {
		toSerialize["scaleid"] = o.Scaleid
	}
	return toSerialize, nil
}

type NullableCoreGradesUpdateGradesRequestItemdetails struct {
	value *CoreGradesUpdateGradesRequestItemdetails
	isSet bool
}

func (v NullableCoreGradesUpdateGradesRequestItemdetails) Get() *CoreGradesUpdateGradesRequestItemdetails {
	return v.value
}

func (v *NullableCoreGradesUpdateGradesRequestItemdetails) Set(val *CoreGradesUpdateGradesRequestItemdetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGradesUpdateGradesRequestItemdetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGradesUpdateGradesRequestItemdetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGradesUpdateGradesRequestItemdetails(val *CoreGradesUpdateGradesRequestItemdetails) *NullableCoreGradesUpdateGradesRequestItemdetails {
	return &NullableCoreGradesUpdateGradesRequestItemdetails{value: val, isSet: true}
}

func (v NullableCoreGradesUpdateGradesRequestItemdetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGradesUpdateGradesRequestItemdetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


