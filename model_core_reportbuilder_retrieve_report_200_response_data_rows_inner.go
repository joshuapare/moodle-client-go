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

// checks if the CoreReportbuilderRetrieveReport200ResponseDataRowsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreReportbuilderRetrieveReport200ResponseDataRowsInner{}

// CoreReportbuilderRetrieveReport200ResponseDataRowsInner struct for CoreReportbuilderRetrieveReport200ResponseDataRowsInner
type CoreReportbuilderRetrieveReport200ResponseDataRowsInner struct {
	Columns []map[string]interface{} `json:"columns,omitempty"`
}

// NewCoreReportbuilderRetrieveReport200ResponseDataRowsInner instantiates a new CoreReportbuilderRetrieveReport200ResponseDataRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreReportbuilderRetrieveReport200ResponseDataRowsInner() *CoreReportbuilderRetrieveReport200ResponseDataRowsInner {
	this := CoreReportbuilderRetrieveReport200ResponseDataRowsInner{}
	return &this
}

// NewCoreReportbuilderRetrieveReport200ResponseDataRowsInnerWithDefaults instantiates a new CoreReportbuilderRetrieveReport200ResponseDataRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreReportbuilderRetrieveReport200ResponseDataRowsInnerWithDefaults() *CoreReportbuilderRetrieveReport200ResponseDataRowsInner {
	this := CoreReportbuilderRetrieveReport200ResponseDataRowsInner{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *CoreReportbuilderRetrieveReport200ResponseDataRowsInner) GetColumns() []map[string]interface{} {
	if o == nil || IsNil(o.Columns) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderRetrieveReport200ResponseDataRowsInner) GetColumnsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *CoreReportbuilderRetrieveReport200ResponseDataRowsInner) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []map[string]interface{} and assigns it to the Columns field.
func (o *CoreReportbuilderRetrieveReport200ResponseDataRowsInner) SetColumns(v []map[string]interface{}) {
	o.Columns = v
}

func (o CoreReportbuilderRetrieveReport200ResponseDataRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreReportbuilderRetrieveReport200ResponseDataRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	return toSerialize, nil
}

type NullableCoreReportbuilderRetrieveReport200ResponseDataRowsInner struct {
	value *CoreReportbuilderRetrieveReport200ResponseDataRowsInner
	isSet bool
}

func (v NullableCoreReportbuilderRetrieveReport200ResponseDataRowsInner) Get() *CoreReportbuilderRetrieveReport200ResponseDataRowsInner {
	return v.value
}

func (v *NullableCoreReportbuilderRetrieveReport200ResponseDataRowsInner) Set(val *CoreReportbuilderRetrieveReport200ResponseDataRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreReportbuilderRetrieveReport200ResponseDataRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreReportbuilderRetrieveReport200ResponseDataRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreReportbuilderRetrieveReport200ResponseDataRowsInner(val *CoreReportbuilderRetrieveReport200ResponseDataRowsInner) *NullableCoreReportbuilderRetrieveReport200ResponseDataRowsInner {
	return &NullableCoreReportbuilderRetrieveReport200ResponseDataRowsInner{value: val, isSet: true}
}

func (v NullableCoreReportbuilderRetrieveReport200ResponseDataRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreReportbuilderRetrieveReport200ResponseDataRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


