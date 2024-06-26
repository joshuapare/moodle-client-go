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

// checks if the CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions{}

// CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions struct for CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions
type CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions struct {
	// exclude empty grades
	Aggregateonlygraded *bool `json:"aggregateonlygraded,omitempty"`
	// aggregate outcomes
	Aggregateoutcomes *bool `json:"aggregateoutcomes,omitempty"`
	// aggregation method
	Aggregation *int32 `json:"aggregation,omitempty"`
	// weight coefficient
	Aggregationcoef2 *string `json:"aggregationcoef2,omitempty"`
	// the decimal count
	Decimals *int32 `json:"decimals,omitempty"`
	// the display type
	Display *int32 `json:"display,omitempty"`
	// drop low grades
	Droplow *int32 `json:"droplow,omitempty"`
	// the grade max
	Grademax *int32 `json:"grademax,omitempty"`
	// the grade min
	Grademin *int32 `json:"grademin,omitempty"`
	// the grade to pass
	Gradepass *int32 `json:"gradepass,omitempty"`
	// the grade type
	Gradetype *int32 `json:"gradetype,omitempty"`
	// grades hidden until
	Hiddenuntil *int32 `json:"hiddenuntil,omitempty"`
	// the category idnumber
	Idnumber *string `json:"idnumber,omitempty"`
	// the category iteminfo
	Iteminfo *string `json:"iteminfo,omitempty"`
	// the category total name
	Itemname *string `json:"itemname,omitempty"`
	// lock grades after
	Locktime *int32 `json:"locktime,omitempty"`
	// The parent category id
	Parentcategoryid *int32 `json:"parentcategoryid,omitempty"`
	// the parent category idnumber
	Parentcategoryidnumber *string `json:"parentcategoryidnumber,omitempty"`
	// weight adjusted
	Weightoverride *bool `json:"weightoverride,omitempty"`
}

// NewCoreGradesCreateGradecategoriesRequestCategoriesInnerOptions instantiates a new CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGradesCreateGradecategoriesRequestCategoriesInnerOptions() *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions {
	this := CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions{}
	var aggregateonlygraded bool = null
	this.Aggregateonlygraded = &aggregateonlygraded
	var aggregateoutcomes bool = null
	this.Aggregateoutcomes = &aggregateoutcomes
	var aggregation int32 = null
	this.Aggregation = &aggregation
	var aggregationcoef2 string = "null"
	this.Aggregationcoef2 = &aggregationcoef2
	var decimals int32 = null
	this.Decimals = &decimals
	var display int32 = null
	this.Display = &display
	var droplow int32 = null
	this.Droplow = &droplow
	var grademax int32 = null
	this.Grademax = &grademax
	var grademin int32 = null
	this.Grademin = &grademin
	var gradepass int32 = null
	this.Gradepass = &gradepass
	var gradetype int32 = null
	this.Gradetype = &gradetype
	var hiddenuntil int32 = null
	this.Hiddenuntil = &hiddenuntil
	var idnumber string = "null"
	this.Idnumber = &idnumber
	var iteminfo string = "null"
	this.Iteminfo = &iteminfo
	var itemname string = "null"
	this.Itemname = &itemname
	var locktime int32 = null
	this.Locktime = &locktime
	var parentcategoryid int32 = null
	this.Parentcategoryid = &parentcategoryid
	var parentcategoryidnumber string = "null"
	this.Parentcategoryidnumber = &parentcategoryidnumber
	var weightoverride bool = null
	this.Weightoverride = &weightoverride
	return &this
}

// NewCoreGradesCreateGradecategoriesRequestCategoriesInnerOptionsWithDefaults instantiates a new CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGradesCreateGradecategoriesRequestCategoriesInnerOptionsWithDefaults() *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions {
	this := CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions{}
	var aggregateonlygraded bool = null
	this.Aggregateonlygraded = &aggregateonlygraded
	var aggregateoutcomes bool = null
	this.Aggregateoutcomes = &aggregateoutcomes
	var aggregation int32 = null
	this.Aggregation = &aggregation
	var aggregationcoef2 string = "null"
	this.Aggregationcoef2 = &aggregationcoef2
	var decimals int32 = null
	this.Decimals = &decimals
	var display int32 = null
	this.Display = &display
	var droplow int32 = null
	this.Droplow = &droplow
	var grademax int32 = null
	this.Grademax = &grademax
	var grademin int32 = null
	this.Grademin = &grademin
	var gradepass int32 = null
	this.Gradepass = &gradepass
	var gradetype int32 = null
	this.Gradetype = &gradetype
	var hiddenuntil int32 = null
	this.Hiddenuntil = &hiddenuntil
	var idnumber string = "null"
	this.Idnumber = &idnumber
	var iteminfo string = "null"
	this.Iteminfo = &iteminfo
	var itemname string = "null"
	this.Itemname = &itemname
	var locktime int32 = null
	this.Locktime = &locktime
	var parentcategoryid int32 = null
	this.Parentcategoryid = &parentcategoryid
	var parentcategoryidnumber string = "null"
	this.Parentcategoryidnumber = &parentcategoryidnumber
	var weightoverride bool = null
	this.Weightoverride = &weightoverride
	return &this
}

// GetAggregateonlygraded returns the Aggregateonlygraded field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetAggregateonlygraded() bool {
	if o == nil || IsNil(o.Aggregateonlygraded) {
		var ret bool
		return ret
	}
	return *o.Aggregateonlygraded
}

// GetAggregateonlygradedOk returns a tuple with the Aggregateonlygraded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetAggregateonlygradedOk() (*bool, bool) {
	if o == nil || IsNil(o.Aggregateonlygraded) {
		return nil, false
	}
	return o.Aggregateonlygraded, true
}

// HasAggregateonlygraded returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasAggregateonlygraded() bool {
	if o != nil && !IsNil(o.Aggregateonlygraded) {
		return true
	}

	return false
}

// SetAggregateonlygraded gets a reference to the given bool and assigns it to the Aggregateonlygraded field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetAggregateonlygraded(v bool) {
	o.Aggregateonlygraded = &v
}

// GetAggregateoutcomes returns the Aggregateoutcomes field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetAggregateoutcomes() bool {
	if o == nil || IsNil(o.Aggregateoutcomes) {
		var ret bool
		return ret
	}
	return *o.Aggregateoutcomes
}

// GetAggregateoutcomesOk returns a tuple with the Aggregateoutcomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetAggregateoutcomesOk() (*bool, bool) {
	if o == nil || IsNil(o.Aggregateoutcomes) {
		return nil, false
	}
	return o.Aggregateoutcomes, true
}

// HasAggregateoutcomes returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasAggregateoutcomes() bool {
	if o != nil && !IsNil(o.Aggregateoutcomes) {
		return true
	}

	return false
}

// SetAggregateoutcomes gets a reference to the given bool and assigns it to the Aggregateoutcomes field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetAggregateoutcomes(v bool) {
	o.Aggregateoutcomes = &v
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetAggregation() int32 {
	if o == nil || IsNil(o.Aggregation) {
		var ret int32
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetAggregationOk() (*int32, bool) {
	if o == nil || IsNil(o.Aggregation) {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasAggregation() bool {
	if o != nil && !IsNil(o.Aggregation) {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given int32 and assigns it to the Aggregation field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetAggregation(v int32) {
	o.Aggregation = &v
}

// GetAggregationcoef2 returns the Aggregationcoef2 field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetAggregationcoef2() string {
	if o == nil || IsNil(o.Aggregationcoef2) {
		var ret string
		return ret
	}
	return *o.Aggregationcoef2
}

// GetAggregationcoef2Ok returns a tuple with the Aggregationcoef2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetAggregationcoef2Ok() (*string, bool) {
	if o == nil || IsNil(o.Aggregationcoef2) {
		return nil, false
	}
	return o.Aggregationcoef2, true
}

// HasAggregationcoef2 returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasAggregationcoef2() bool {
	if o != nil && !IsNil(o.Aggregationcoef2) {
		return true
	}

	return false
}

// SetAggregationcoef2 gets a reference to the given string and assigns it to the Aggregationcoef2 field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetAggregationcoef2(v string) {
	o.Aggregationcoef2 = &v
}

// GetDecimals returns the Decimals field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetDecimals() int32 {
	if o == nil || IsNil(o.Decimals) {
		var ret int32
		return ret
	}
	return *o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetDecimalsOk() (*int32, bool) {
	if o == nil || IsNil(o.Decimals) {
		return nil, false
	}
	return o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasDecimals() bool {
	if o != nil && !IsNil(o.Decimals) {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given int32 and assigns it to the Decimals field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetDecimals(v int32) {
	o.Decimals = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetDisplay() int32 {
	if o == nil || IsNil(o.Display) {
		var ret int32
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetDisplayOk() (*int32, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given int32 and assigns it to the Display field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetDisplay(v int32) {
	o.Display = &v
}

// GetDroplow returns the Droplow field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetDroplow() int32 {
	if o == nil || IsNil(o.Droplow) {
		var ret int32
		return ret
	}
	return *o.Droplow
}

// GetDroplowOk returns a tuple with the Droplow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetDroplowOk() (*int32, bool) {
	if o == nil || IsNil(o.Droplow) {
		return nil, false
	}
	return o.Droplow, true
}

// HasDroplow returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasDroplow() bool {
	if o != nil && !IsNil(o.Droplow) {
		return true
	}

	return false
}

// SetDroplow gets a reference to the given int32 and assigns it to the Droplow field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetDroplow(v int32) {
	o.Droplow = &v
}

// GetGrademax returns the Grademax field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetGrademax() int32 {
	if o == nil || IsNil(o.Grademax) {
		var ret int32
		return ret
	}
	return *o.Grademax
}

// GetGrademaxOk returns a tuple with the Grademax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetGrademaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Grademax) {
		return nil, false
	}
	return o.Grademax, true
}

// HasGrademax returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasGrademax() bool {
	if o != nil && !IsNil(o.Grademax) {
		return true
	}

	return false
}

// SetGrademax gets a reference to the given int32 and assigns it to the Grademax field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetGrademax(v int32) {
	o.Grademax = &v
}

// GetGrademin returns the Grademin field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetGrademin() int32 {
	if o == nil || IsNil(o.Grademin) {
		var ret int32
		return ret
	}
	return *o.Grademin
}

// GetGrademinOk returns a tuple with the Grademin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetGrademinOk() (*int32, bool) {
	if o == nil || IsNil(o.Grademin) {
		return nil, false
	}
	return o.Grademin, true
}

// HasGrademin returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasGrademin() bool {
	if o != nil && !IsNil(o.Grademin) {
		return true
	}

	return false
}

// SetGrademin gets a reference to the given int32 and assigns it to the Grademin field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetGrademin(v int32) {
	o.Grademin = &v
}

// GetGradepass returns the Gradepass field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetGradepass() int32 {
	if o == nil || IsNil(o.Gradepass) {
		var ret int32
		return ret
	}
	return *o.Gradepass
}

// GetGradepassOk returns a tuple with the Gradepass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetGradepassOk() (*int32, bool) {
	if o == nil || IsNil(o.Gradepass) {
		return nil, false
	}
	return o.Gradepass, true
}

// HasGradepass returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasGradepass() bool {
	if o != nil && !IsNil(o.Gradepass) {
		return true
	}

	return false
}

// SetGradepass gets a reference to the given int32 and assigns it to the Gradepass field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetGradepass(v int32) {
	o.Gradepass = &v
}

// GetGradetype returns the Gradetype field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetGradetype() int32 {
	if o == nil || IsNil(o.Gradetype) {
		var ret int32
		return ret
	}
	return *o.Gradetype
}

// GetGradetypeOk returns a tuple with the Gradetype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetGradetypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Gradetype) {
		return nil, false
	}
	return o.Gradetype, true
}

// HasGradetype returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasGradetype() bool {
	if o != nil && !IsNil(o.Gradetype) {
		return true
	}

	return false
}

// SetGradetype gets a reference to the given int32 and assigns it to the Gradetype field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetGradetype(v int32) {
	o.Gradetype = &v
}

// GetHiddenuntil returns the Hiddenuntil field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetHiddenuntil() int32 {
	if o == nil || IsNil(o.Hiddenuntil) {
		var ret int32
		return ret
	}
	return *o.Hiddenuntil
}

// GetHiddenuntilOk returns a tuple with the Hiddenuntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetHiddenuntilOk() (*int32, bool) {
	if o == nil || IsNil(o.Hiddenuntil) {
		return nil, false
	}
	return o.Hiddenuntil, true
}

// HasHiddenuntil returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasHiddenuntil() bool {
	if o != nil && !IsNil(o.Hiddenuntil) {
		return true
	}

	return false
}

// SetHiddenuntil gets a reference to the given int32 and assigns it to the Hiddenuntil field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetHiddenuntil(v int32) {
	o.Hiddenuntil = &v
}

// GetIdnumber returns the Idnumber field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetIdnumber() string {
	if o == nil || IsNil(o.Idnumber) {
		var ret string
		return ret
	}
	return *o.Idnumber
}

// GetIdnumberOk returns a tuple with the Idnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetIdnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Idnumber) {
		return nil, false
	}
	return o.Idnumber, true
}

// HasIdnumber returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasIdnumber() bool {
	if o != nil && !IsNil(o.Idnumber) {
		return true
	}

	return false
}

// SetIdnumber gets a reference to the given string and assigns it to the Idnumber field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetIdnumber(v string) {
	o.Idnumber = &v
}

// GetIteminfo returns the Iteminfo field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetIteminfo() string {
	if o == nil || IsNil(o.Iteminfo) {
		var ret string
		return ret
	}
	return *o.Iteminfo
}

// GetIteminfoOk returns a tuple with the Iteminfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetIteminfoOk() (*string, bool) {
	if o == nil || IsNil(o.Iteminfo) {
		return nil, false
	}
	return o.Iteminfo, true
}

// HasIteminfo returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasIteminfo() bool {
	if o != nil && !IsNil(o.Iteminfo) {
		return true
	}

	return false
}

// SetIteminfo gets a reference to the given string and assigns it to the Iteminfo field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetIteminfo(v string) {
	o.Iteminfo = &v
}

// GetItemname returns the Itemname field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetItemname() string {
	if o == nil || IsNil(o.Itemname) {
		var ret string
		return ret
	}
	return *o.Itemname
}

// GetItemnameOk returns a tuple with the Itemname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetItemnameOk() (*string, bool) {
	if o == nil || IsNil(o.Itemname) {
		return nil, false
	}
	return o.Itemname, true
}

// HasItemname returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasItemname() bool {
	if o != nil && !IsNil(o.Itemname) {
		return true
	}

	return false
}

// SetItemname gets a reference to the given string and assigns it to the Itemname field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetItemname(v string) {
	o.Itemname = &v
}

// GetLocktime returns the Locktime field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetLocktime() int32 {
	if o == nil || IsNil(o.Locktime) {
		var ret int32
		return ret
	}
	return *o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetLocktimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Locktime) {
		return nil, false
	}
	return o.Locktime, true
}

// HasLocktime returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasLocktime() bool {
	if o != nil && !IsNil(o.Locktime) {
		return true
	}

	return false
}

// SetLocktime gets a reference to the given int32 and assigns it to the Locktime field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetLocktime(v int32) {
	o.Locktime = &v
}

// GetParentcategoryid returns the Parentcategoryid field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetParentcategoryid() int32 {
	if o == nil || IsNil(o.Parentcategoryid) {
		var ret int32
		return ret
	}
	return *o.Parentcategoryid
}

// GetParentcategoryidOk returns a tuple with the Parentcategoryid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetParentcategoryidOk() (*int32, bool) {
	if o == nil || IsNil(o.Parentcategoryid) {
		return nil, false
	}
	return o.Parentcategoryid, true
}

// HasParentcategoryid returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasParentcategoryid() bool {
	if o != nil && !IsNil(o.Parentcategoryid) {
		return true
	}

	return false
}

// SetParentcategoryid gets a reference to the given int32 and assigns it to the Parentcategoryid field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetParentcategoryid(v int32) {
	o.Parentcategoryid = &v
}

// GetParentcategoryidnumber returns the Parentcategoryidnumber field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetParentcategoryidnumber() string {
	if o == nil || IsNil(o.Parentcategoryidnumber) {
		var ret string
		return ret
	}
	return *o.Parentcategoryidnumber
}

// GetParentcategoryidnumberOk returns a tuple with the Parentcategoryidnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetParentcategoryidnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Parentcategoryidnumber) {
		return nil, false
	}
	return o.Parentcategoryidnumber, true
}

// HasParentcategoryidnumber returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasParentcategoryidnumber() bool {
	if o != nil && !IsNil(o.Parentcategoryidnumber) {
		return true
	}

	return false
}

// SetParentcategoryidnumber gets a reference to the given string and assigns it to the Parentcategoryidnumber field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetParentcategoryidnumber(v string) {
	o.Parentcategoryidnumber = &v
}

// GetWeightoverride returns the Weightoverride field value if set, zero value otherwise.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetWeightoverride() bool {
	if o == nil || IsNil(o.Weightoverride) {
		var ret bool
		return ret
	}
	return *o.Weightoverride
}

// GetWeightoverrideOk returns a tuple with the Weightoverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) GetWeightoverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.Weightoverride) {
		return nil, false
	}
	return o.Weightoverride, true
}

// HasWeightoverride returns a boolean if a field has been set.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) HasWeightoverride() bool {
	if o != nil && !IsNil(o.Weightoverride) {
		return true
	}

	return false
}

// SetWeightoverride gets a reference to the given bool and assigns it to the Weightoverride field.
func (o *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) SetWeightoverride(v bool) {
	o.Weightoverride = &v
}

func (o CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Aggregateonlygraded) {
		toSerialize["aggregateonlygraded"] = o.Aggregateonlygraded
	}
	if !IsNil(o.Aggregateoutcomes) {
		toSerialize["aggregateoutcomes"] = o.Aggregateoutcomes
	}
	if !IsNil(o.Aggregation) {
		toSerialize["aggregation"] = o.Aggregation
	}
	if !IsNil(o.Aggregationcoef2) {
		toSerialize["aggregationcoef2"] = o.Aggregationcoef2
	}
	if !IsNil(o.Decimals) {
		toSerialize["decimals"] = o.Decimals
	}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.Droplow) {
		toSerialize["droplow"] = o.Droplow
	}
	if !IsNil(o.Grademax) {
		toSerialize["grademax"] = o.Grademax
	}
	if !IsNil(o.Grademin) {
		toSerialize["grademin"] = o.Grademin
	}
	if !IsNil(o.Gradepass) {
		toSerialize["gradepass"] = o.Gradepass
	}
	if !IsNil(o.Gradetype) {
		toSerialize["gradetype"] = o.Gradetype
	}
	if !IsNil(o.Hiddenuntil) {
		toSerialize["hiddenuntil"] = o.Hiddenuntil
	}
	if !IsNil(o.Idnumber) {
		toSerialize["idnumber"] = o.Idnumber
	}
	if !IsNil(o.Iteminfo) {
		toSerialize["iteminfo"] = o.Iteminfo
	}
	if !IsNil(o.Itemname) {
		toSerialize["itemname"] = o.Itemname
	}
	if !IsNil(o.Locktime) {
		toSerialize["locktime"] = o.Locktime
	}
	if !IsNil(o.Parentcategoryid) {
		toSerialize["parentcategoryid"] = o.Parentcategoryid
	}
	if !IsNil(o.Parentcategoryidnumber) {
		toSerialize["parentcategoryidnumber"] = o.Parentcategoryidnumber
	}
	if !IsNil(o.Weightoverride) {
		toSerialize["weightoverride"] = o.Weightoverride
	}
	return toSerialize, nil
}

type NullableCoreGradesCreateGradecategoriesRequestCategoriesInnerOptions struct {
	value *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions
	isSet bool
}

func (v NullableCoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) Get() *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions {
	return v.value
}

func (v *NullableCoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) Set(val *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGradesCreateGradecategoriesRequestCategoriesInnerOptions(val *CoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) *NullableCoreGradesCreateGradecategoriesRequestCategoriesInnerOptions {
	return &NullableCoreGradesCreateGradecategoriesRequestCategoriesInnerOptions{value: val, isSet: true}
}

func (v NullableCoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGradesCreateGradecategoriesRequestCategoriesInnerOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


