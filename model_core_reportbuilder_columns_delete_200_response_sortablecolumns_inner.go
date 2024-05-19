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

// checks if the CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner{}

// CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner struct for CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner
type CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner struct {
	// heading
	Heading *string `json:"heading,omitempty"`
	// id
	Id *int32 `json:"id,omitempty"`
	// movetitle
	Movetitle *string `json:"movetitle,omitempty"`
	// sortdirection
	Sortdirection *int32 `json:"sortdirection,omitempty"`
	// sortenabled
	Sortenabled *bool `json:"sortenabled,omitempty"`
	// sortenabledtitle
	Sortenabledtitle *string `json:"sortenabledtitle,omitempty"`
	Sorticon *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInnerSorticon `json:"sorticon,omitempty"`
	// sortorder
	Sortorder *int32 `json:"sortorder,omitempty"`
	// title
	Title *string `json:"title,omitempty"`
}

// NewCoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner instantiates a new CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner() *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner {
	this := CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner{}
	return &this
}

// NewCoreReportbuilderColumnsDelete200ResponseSortablecolumnsInnerWithDefaults instantiates a new CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreReportbuilderColumnsDelete200ResponseSortablecolumnsInnerWithDefaults() *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner {
	this := CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner{}
	return &this
}

// GetHeading returns the Heading field value if set, zero value otherwise.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetHeading() string {
	if o == nil || IsNil(o.Heading) {
		var ret string
		return ret
	}
	return *o.Heading
}

// GetHeadingOk returns a tuple with the Heading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetHeadingOk() (*string, bool) {
	if o == nil || IsNil(o.Heading) {
		return nil, false
	}
	return o.Heading, true
}

// HasHeading returns a boolean if a field has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) HasHeading() bool {
	if o != nil && !IsNil(o.Heading) {
		return true
	}

	return false
}

// SetHeading gets a reference to the given string and assigns it to the Heading field.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) SetHeading(v string) {
	o.Heading = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) SetId(v int32) {
	o.Id = &v
}

// GetMovetitle returns the Movetitle field value if set, zero value otherwise.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetMovetitle() string {
	if o == nil || IsNil(o.Movetitle) {
		var ret string
		return ret
	}
	return *o.Movetitle
}

// GetMovetitleOk returns a tuple with the Movetitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetMovetitleOk() (*string, bool) {
	if o == nil || IsNil(o.Movetitle) {
		return nil, false
	}
	return o.Movetitle, true
}

// HasMovetitle returns a boolean if a field has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) HasMovetitle() bool {
	if o != nil && !IsNil(o.Movetitle) {
		return true
	}

	return false
}

// SetMovetitle gets a reference to the given string and assigns it to the Movetitle field.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) SetMovetitle(v string) {
	o.Movetitle = &v
}

// GetSortdirection returns the Sortdirection field value if set, zero value otherwise.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetSortdirection() int32 {
	if o == nil || IsNil(o.Sortdirection) {
		var ret int32
		return ret
	}
	return *o.Sortdirection
}

// GetSortdirectionOk returns a tuple with the Sortdirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetSortdirectionOk() (*int32, bool) {
	if o == nil || IsNil(o.Sortdirection) {
		return nil, false
	}
	return o.Sortdirection, true
}

// HasSortdirection returns a boolean if a field has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) HasSortdirection() bool {
	if o != nil && !IsNil(o.Sortdirection) {
		return true
	}

	return false
}

// SetSortdirection gets a reference to the given int32 and assigns it to the Sortdirection field.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) SetSortdirection(v int32) {
	o.Sortdirection = &v
}

// GetSortenabled returns the Sortenabled field value if set, zero value otherwise.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetSortenabled() bool {
	if o == nil || IsNil(o.Sortenabled) {
		var ret bool
		return ret
	}
	return *o.Sortenabled
}

// GetSortenabledOk returns a tuple with the Sortenabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetSortenabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Sortenabled) {
		return nil, false
	}
	return o.Sortenabled, true
}

// HasSortenabled returns a boolean if a field has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) HasSortenabled() bool {
	if o != nil && !IsNil(o.Sortenabled) {
		return true
	}

	return false
}

// SetSortenabled gets a reference to the given bool and assigns it to the Sortenabled field.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) SetSortenabled(v bool) {
	o.Sortenabled = &v
}

// GetSortenabledtitle returns the Sortenabledtitle field value if set, zero value otherwise.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetSortenabledtitle() string {
	if o == nil || IsNil(o.Sortenabledtitle) {
		var ret string
		return ret
	}
	return *o.Sortenabledtitle
}

// GetSortenabledtitleOk returns a tuple with the Sortenabledtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetSortenabledtitleOk() (*string, bool) {
	if o == nil || IsNil(o.Sortenabledtitle) {
		return nil, false
	}
	return o.Sortenabledtitle, true
}

// HasSortenabledtitle returns a boolean if a field has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) HasSortenabledtitle() bool {
	if o != nil && !IsNil(o.Sortenabledtitle) {
		return true
	}

	return false
}

// SetSortenabledtitle gets a reference to the given string and assigns it to the Sortenabledtitle field.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) SetSortenabledtitle(v string) {
	o.Sortenabledtitle = &v
}

// GetSorticon returns the Sorticon field value if set, zero value otherwise.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetSorticon() CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInnerSorticon {
	if o == nil || IsNil(o.Sorticon) {
		var ret CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInnerSorticon
		return ret
	}
	return *o.Sorticon
}

// GetSorticonOk returns a tuple with the Sorticon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetSorticonOk() (*CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInnerSorticon, bool) {
	if o == nil || IsNil(o.Sorticon) {
		return nil, false
	}
	return o.Sorticon, true
}

// HasSorticon returns a boolean if a field has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) HasSorticon() bool {
	if o != nil && !IsNil(o.Sorticon) {
		return true
	}

	return false
}

// SetSorticon gets a reference to the given CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInnerSorticon and assigns it to the Sorticon field.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) SetSorticon(v CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInnerSorticon) {
	o.Sorticon = &v
}

// GetSortorder returns the Sortorder field value if set, zero value otherwise.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetSortorder() int32 {
	if o == nil || IsNil(o.Sortorder) {
		var ret int32
		return ret
	}
	return *o.Sortorder
}

// GetSortorderOk returns a tuple with the Sortorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetSortorderOk() (*int32, bool) {
	if o == nil || IsNil(o.Sortorder) {
		return nil, false
	}
	return o.Sortorder, true
}

// HasSortorder returns a boolean if a field has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) HasSortorder() bool {
	if o != nil && !IsNil(o.Sortorder) {
		return true
	}

	return false
}

// SetSortorder gets a reference to the given int32 and assigns it to the Sortorder field.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) SetSortorder(v int32) {
	o.Sortorder = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) SetTitle(v string) {
	o.Title = &v
}

func (o CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Heading) {
		toSerialize["heading"] = o.Heading
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Movetitle) {
		toSerialize["movetitle"] = o.Movetitle
	}
	if !IsNil(o.Sortdirection) {
		toSerialize["sortdirection"] = o.Sortdirection
	}
	if !IsNil(o.Sortenabled) {
		toSerialize["sortenabled"] = o.Sortenabled
	}
	if !IsNil(o.Sortenabledtitle) {
		toSerialize["sortenabledtitle"] = o.Sortenabledtitle
	}
	if !IsNil(o.Sorticon) {
		toSerialize["sorticon"] = o.Sorticon
	}
	if !IsNil(o.Sortorder) {
		toSerialize["sortorder"] = o.Sortorder
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableCoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner struct {
	value *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner
	isSet bool
}

func (v NullableCoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) Get() *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner {
	return v.value
}

func (v *NullableCoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) Set(val *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner(val *CoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) *NullableCoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner {
	return &NullableCoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner{value: val, isSet: true}
}

func (v NullableCoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreReportbuilderColumnsDelete200ResponseSortablecolumnsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

