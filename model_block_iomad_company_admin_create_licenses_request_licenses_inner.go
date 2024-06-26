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

// checks if the BlockIomadCompanyAdminCreateLicensesRequestLicensesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockIomadCompanyAdminCreateLicensesRequestLicensesInner{}

// BlockIomadCompanyAdminCreateLicensesRequestLicensesInner one or many licenses
type BlockIomadCompanyAdminCreateLicensesRequestLicensesInner struct {
	// Number of license slots
	Allocation *int32 `json:"allocation,omitempty"`
	// Clear license assignments on expire - 0 = no, 1 = yes
	Clearonexpire *int32 `json:"clearonexpire,omitempty"`
	// Company id
	Companyid *int32 `json:"companyid,omitempty"`
	Courses []BlockIomadCompanyAdminCreateLicensesRequestLicensesInnerCoursesInner `json:"courses,omitempty"`
	// License cut off date (int = timestamp)
	Cutoffdate *int32 `json:"cutoffdate,omitempty"`
	// License expiry date (int = timestamp)
	Expirydate *int32 `json:"expirydate,omitempty"`
	// Instant access - 0 = no, 1 = yes
	Instant *int32 `json:"instant,omitempty"`
	// License name
	Name *string `json:"name,omitempty"`
	// Parent license id
	Parentid *int32 `json:"parentid,omitempty"`
	// Program pf courses 0 = no, 1 = yes
	Program *int32 `json:"program,omitempty"`
	// License reference
	Reference *string `json:"reference,omitempty"`
	// Date from which the liucense is available (int = timestamp) 
	Startdate *int32 `json:"startdate,omitempty"`
	// License type - 0 = standard, 1 = reusable, 2 = standard educator, 3 = reusable educator
	Type *int32 `json:"type,omitempty"`
	// Number how often the lic can be allocated
	Used *int32 `json:"used,omitempty"`
	// Course access length (days)
	Validlength *int32 `json:"validlength,omitempty"`
}

// NewBlockIomadCompanyAdminCreateLicensesRequestLicensesInner instantiates a new BlockIomadCompanyAdminCreateLicensesRequestLicensesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockIomadCompanyAdminCreateLicensesRequestLicensesInner() *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner {
	this := BlockIomadCompanyAdminCreateLicensesRequestLicensesInner{}
	var allocation int32 = null
	this.Allocation = &allocation
	var clearonexpire int32 = null
	this.Clearonexpire = &clearonexpire
	var companyid int32 = null
	this.Companyid = &companyid
	var cutoffdate int32 = null
	this.Cutoffdate = &cutoffdate
	var expirydate int32 = null
	this.Expirydate = &expirydate
	var instant int32 = null
	this.Instant = &instant
	var name string = "null"
	this.Name = &name
	var parentid int32 = null
	this.Parentid = &parentid
	var program int32 = null
	this.Program = &program
	var reference string = "null"
	this.Reference = &reference
	var startdate int32 = null
	this.Startdate = &startdate
	var type_ int32 = null
	this.Type = &type_
	var used int32 = null
	this.Used = &used
	var validlength int32 = null
	this.Validlength = &validlength
	return &this
}

// NewBlockIomadCompanyAdminCreateLicensesRequestLicensesInnerWithDefaults instantiates a new BlockIomadCompanyAdminCreateLicensesRequestLicensesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockIomadCompanyAdminCreateLicensesRequestLicensesInnerWithDefaults() *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner {
	this := BlockIomadCompanyAdminCreateLicensesRequestLicensesInner{}
	var allocation int32 = null
	this.Allocation = &allocation
	var clearonexpire int32 = null
	this.Clearonexpire = &clearonexpire
	var companyid int32 = null
	this.Companyid = &companyid
	var cutoffdate int32 = null
	this.Cutoffdate = &cutoffdate
	var expirydate int32 = null
	this.Expirydate = &expirydate
	var instant int32 = null
	this.Instant = &instant
	var name string = "null"
	this.Name = &name
	var parentid int32 = null
	this.Parentid = &parentid
	var program int32 = null
	this.Program = &program
	var reference string = "null"
	this.Reference = &reference
	var startdate int32 = null
	this.Startdate = &startdate
	var type_ int32 = null
	this.Type = &type_
	var used int32 = null
	this.Used = &used
	var validlength int32 = null
	this.Validlength = &validlength
	return &this
}

// GetAllocation returns the Allocation field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetAllocation() int32 {
	if o == nil || IsNil(o.Allocation) {
		var ret int32
		return ret
	}
	return *o.Allocation
}

// GetAllocationOk returns a tuple with the Allocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetAllocationOk() (*int32, bool) {
	if o == nil || IsNil(o.Allocation) {
		return nil, false
	}
	return o.Allocation, true
}

// HasAllocation returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasAllocation() bool {
	if o != nil && !IsNil(o.Allocation) {
		return true
	}

	return false
}

// SetAllocation gets a reference to the given int32 and assigns it to the Allocation field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetAllocation(v int32) {
	o.Allocation = &v
}

// GetClearonexpire returns the Clearonexpire field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetClearonexpire() int32 {
	if o == nil || IsNil(o.Clearonexpire) {
		var ret int32
		return ret
	}
	return *o.Clearonexpire
}

// GetClearonexpireOk returns a tuple with the Clearonexpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetClearonexpireOk() (*int32, bool) {
	if o == nil || IsNil(o.Clearonexpire) {
		return nil, false
	}
	return o.Clearonexpire, true
}

// HasClearonexpire returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasClearonexpire() bool {
	if o != nil && !IsNil(o.Clearonexpire) {
		return true
	}

	return false
}

// SetClearonexpire gets a reference to the given int32 and assigns it to the Clearonexpire field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetClearonexpire(v int32) {
	o.Clearonexpire = &v
}

// GetCompanyid returns the Companyid field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetCompanyid() int32 {
	if o == nil || IsNil(o.Companyid) {
		var ret int32
		return ret
	}
	return *o.Companyid
}

// GetCompanyidOk returns a tuple with the Companyid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetCompanyidOk() (*int32, bool) {
	if o == nil || IsNil(o.Companyid) {
		return nil, false
	}
	return o.Companyid, true
}

// HasCompanyid returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasCompanyid() bool {
	if o != nil && !IsNil(o.Companyid) {
		return true
	}

	return false
}

// SetCompanyid gets a reference to the given int32 and assigns it to the Companyid field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetCompanyid(v int32) {
	o.Companyid = &v
}

// GetCourses returns the Courses field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetCourses() []BlockIomadCompanyAdminCreateLicensesRequestLicensesInnerCoursesInner {
	if o == nil || IsNil(o.Courses) {
		var ret []BlockIomadCompanyAdminCreateLicensesRequestLicensesInnerCoursesInner
		return ret
	}
	return o.Courses
}

// GetCoursesOk returns a tuple with the Courses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetCoursesOk() ([]BlockIomadCompanyAdminCreateLicensesRequestLicensesInnerCoursesInner, bool) {
	if o == nil || IsNil(o.Courses) {
		return nil, false
	}
	return o.Courses, true
}

// HasCourses returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasCourses() bool {
	if o != nil && !IsNil(o.Courses) {
		return true
	}

	return false
}

// SetCourses gets a reference to the given []BlockIomadCompanyAdminCreateLicensesRequestLicensesInnerCoursesInner and assigns it to the Courses field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetCourses(v []BlockIomadCompanyAdminCreateLicensesRequestLicensesInnerCoursesInner) {
	o.Courses = v
}

// GetCutoffdate returns the Cutoffdate field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetCutoffdate() int32 {
	if o == nil || IsNil(o.Cutoffdate) {
		var ret int32
		return ret
	}
	return *o.Cutoffdate
}

// GetCutoffdateOk returns a tuple with the Cutoffdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetCutoffdateOk() (*int32, bool) {
	if o == nil || IsNil(o.Cutoffdate) {
		return nil, false
	}
	return o.Cutoffdate, true
}

// HasCutoffdate returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasCutoffdate() bool {
	if o != nil && !IsNil(o.Cutoffdate) {
		return true
	}

	return false
}

// SetCutoffdate gets a reference to the given int32 and assigns it to the Cutoffdate field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetCutoffdate(v int32) {
	o.Cutoffdate = &v
}

// GetExpirydate returns the Expirydate field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetExpirydate() int32 {
	if o == nil || IsNil(o.Expirydate) {
		var ret int32
		return ret
	}
	return *o.Expirydate
}

// GetExpirydateOk returns a tuple with the Expirydate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetExpirydateOk() (*int32, bool) {
	if o == nil || IsNil(o.Expirydate) {
		return nil, false
	}
	return o.Expirydate, true
}

// HasExpirydate returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasExpirydate() bool {
	if o != nil && !IsNil(o.Expirydate) {
		return true
	}

	return false
}

// SetExpirydate gets a reference to the given int32 and assigns it to the Expirydate field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetExpirydate(v int32) {
	o.Expirydate = &v
}

// GetInstant returns the Instant field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetInstant() int32 {
	if o == nil || IsNil(o.Instant) {
		var ret int32
		return ret
	}
	return *o.Instant
}

// GetInstantOk returns a tuple with the Instant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetInstantOk() (*int32, bool) {
	if o == nil || IsNil(o.Instant) {
		return nil, false
	}
	return o.Instant, true
}

// HasInstant returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasInstant() bool {
	if o != nil && !IsNil(o.Instant) {
		return true
	}

	return false
}

// SetInstant gets a reference to the given int32 and assigns it to the Instant field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetInstant(v int32) {
	o.Instant = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetName(v string) {
	o.Name = &v
}

// GetParentid returns the Parentid field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetParentid() int32 {
	if o == nil || IsNil(o.Parentid) {
		var ret int32
		return ret
	}
	return *o.Parentid
}

// GetParentidOk returns a tuple with the Parentid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetParentidOk() (*int32, bool) {
	if o == nil || IsNil(o.Parentid) {
		return nil, false
	}
	return o.Parentid, true
}

// HasParentid returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasParentid() bool {
	if o != nil && !IsNil(o.Parentid) {
		return true
	}

	return false
}

// SetParentid gets a reference to the given int32 and assigns it to the Parentid field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetParentid(v int32) {
	o.Parentid = &v
}

// GetProgram returns the Program field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetProgram() int32 {
	if o == nil || IsNil(o.Program) {
		var ret int32
		return ret
	}
	return *o.Program
}

// GetProgramOk returns a tuple with the Program field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetProgramOk() (*int32, bool) {
	if o == nil || IsNil(o.Program) {
		return nil, false
	}
	return o.Program, true
}

// HasProgram returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasProgram() bool {
	if o != nil && !IsNil(o.Program) {
		return true
	}

	return false
}

// SetProgram gets a reference to the given int32 and assigns it to the Program field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetProgram(v int32) {
	o.Program = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetReference(v string) {
	o.Reference = &v
}

// GetStartdate returns the Startdate field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetStartdate() int32 {
	if o == nil || IsNil(o.Startdate) {
		var ret int32
		return ret
	}
	return *o.Startdate
}

// GetStartdateOk returns a tuple with the Startdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetStartdateOk() (*int32, bool) {
	if o == nil || IsNil(o.Startdate) {
		return nil, false
	}
	return o.Startdate, true
}

// HasStartdate returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasStartdate() bool {
	if o != nil && !IsNil(o.Startdate) {
		return true
	}

	return false
}

// SetStartdate gets a reference to the given int32 and assigns it to the Startdate field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetStartdate(v int32) {
	o.Startdate = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetType(v int32) {
	o.Type = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetUsed() int32 {
	if o == nil || IsNil(o.Used) {
		var ret int32
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetUsedOk() (*int32, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int32 and assigns it to the Used field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetUsed(v int32) {
	o.Used = &v
}

// GetValidlength returns the Validlength field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetValidlength() int32 {
	if o == nil || IsNil(o.Validlength) {
		var ret int32
		return ret
	}
	return *o.Validlength
}

// GetValidlengthOk returns a tuple with the Validlength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) GetValidlengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Validlength) {
		return nil, false
	}
	return o.Validlength, true
}

// HasValidlength returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) HasValidlength() bool {
	if o != nil && !IsNil(o.Validlength) {
		return true
	}

	return false
}

// SetValidlength gets a reference to the given int32 and assigns it to the Validlength field.
func (o *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) SetValidlength(v int32) {
	o.Validlength = &v
}

func (o BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Allocation) {
		toSerialize["allocation"] = o.Allocation
	}
	if !IsNil(o.Clearonexpire) {
		toSerialize["clearonexpire"] = o.Clearonexpire
	}
	if !IsNil(o.Companyid) {
		toSerialize["companyid"] = o.Companyid
	}
	if !IsNil(o.Courses) {
		toSerialize["courses"] = o.Courses
	}
	if !IsNil(o.Cutoffdate) {
		toSerialize["cutoffdate"] = o.Cutoffdate
	}
	if !IsNil(o.Expirydate) {
		toSerialize["expirydate"] = o.Expirydate
	}
	if !IsNil(o.Instant) {
		toSerialize["instant"] = o.Instant
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Parentid) {
		toSerialize["parentid"] = o.Parentid
	}
	if !IsNil(o.Program) {
		toSerialize["program"] = o.Program
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.Startdate) {
		toSerialize["startdate"] = o.Startdate
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	if !IsNil(o.Validlength) {
		toSerialize["validlength"] = o.Validlength
	}
	return toSerialize, nil
}

type NullableBlockIomadCompanyAdminCreateLicensesRequestLicensesInner struct {
	value *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner
	isSet bool
}

func (v NullableBlockIomadCompanyAdminCreateLicensesRequestLicensesInner) Get() *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner {
	return v.value
}

func (v *NullableBlockIomadCompanyAdminCreateLicensesRequestLicensesInner) Set(val *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockIomadCompanyAdminCreateLicensesRequestLicensesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockIomadCompanyAdminCreateLicensesRequestLicensesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockIomadCompanyAdminCreateLicensesRequestLicensesInner(val *BlockIomadCompanyAdminCreateLicensesRequestLicensesInner) *NullableBlockIomadCompanyAdminCreateLicensesRequestLicensesInner {
	return &NullableBlockIomadCompanyAdminCreateLicensesRequestLicensesInner{value: val, isSet: true}
}

func (v NullableBlockIomadCompanyAdminCreateLicensesRequestLicensesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockIomadCompanyAdminCreateLicensesRequestLicensesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


