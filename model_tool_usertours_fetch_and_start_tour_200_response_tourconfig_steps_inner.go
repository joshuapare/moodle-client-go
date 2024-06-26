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

// checks if the ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner{}

// ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner struct for ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner
type ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner struct {
	// Whether a backdrop should be used
	Backdrop *bool `json:"backdrop,omitempty"`
	// Step Content
	Content *string `json:"content,omitempty"`
	// Delay before showing the step (ms)
	Delay *int32 `json:"delay,omitempty"`
	// Step Target
	Element *string `json:"element,omitempty"`
	// Whether to display the step even if it could not be found
	Orphan *bool `json:"orphan,omitempty"`
	// Step Placement
	Placement *string `json:"placement,omitempty"`
	// Whether to move to the next step when the target element is clicked
	Reflex *bool `json:"reflex,omitempty"`
	// The actual ID of the step
	Stepid *int32 `json:"stepid,omitempty"`
	// Step Title
	Title *string `json:"title,omitempty"`
}

// NewToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner instantiates a new ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner() *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner {
	this := ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner{}
	var backdrop bool = null
	this.Backdrop = &backdrop
	var content string = "null"
	this.Content = &content
	var delay int32 = null
	this.Delay = &delay
	var element string = "null"
	this.Element = &element
	var orphan bool = null
	this.Orphan = &orphan
	var placement string = "null"
	this.Placement = &placement
	var reflex bool = null
	this.Reflex = &reflex
	var stepid int32 = null
	this.Stepid = &stepid
	var title string = "null"
	this.Title = &title
	return &this
}

// NewToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInnerWithDefaults instantiates a new ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInnerWithDefaults() *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner {
	this := ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner{}
	var backdrop bool = null
	this.Backdrop = &backdrop
	var content string = "null"
	this.Content = &content
	var delay int32 = null
	this.Delay = &delay
	var element string = "null"
	this.Element = &element
	var orphan bool = null
	this.Orphan = &orphan
	var placement string = "null"
	this.Placement = &placement
	var reflex bool = null
	this.Reflex = &reflex
	var stepid int32 = null
	this.Stepid = &stepid
	var title string = "null"
	this.Title = &title
	return &this
}

// GetBackdrop returns the Backdrop field value if set, zero value otherwise.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetBackdrop() bool {
	if o == nil || IsNil(o.Backdrop) {
		var ret bool
		return ret
	}
	return *o.Backdrop
}

// GetBackdropOk returns a tuple with the Backdrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetBackdropOk() (*bool, bool) {
	if o == nil || IsNil(o.Backdrop) {
		return nil, false
	}
	return o.Backdrop, true
}

// HasBackdrop returns a boolean if a field has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) HasBackdrop() bool {
	if o != nil && !IsNil(o.Backdrop) {
		return true
	}

	return false
}

// SetBackdrop gets a reference to the given bool and assigns it to the Backdrop field.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) SetBackdrop(v bool) {
	o.Backdrop = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) SetContent(v string) {
	o.Content = &v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetDelay() int32 {
	if o == nil || IsNil(o.Delay) {
		var ret int32
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.Delay) {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) HasDelay() bool {
	if o != nil && !IsNil(o.Delay) {
		return true
	}

	return false
}

// SetDelay gets a reference to the given int32 and assigns it to the Delay field.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) SetDelay(v int32) {
	o.Delay = &v
}

// GetElement returns the Element field value if set, zero value otherwise.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetElement() string {
	if o == nil || IsNil(o.Element) {
		var ret string
		return ret
	}
	return *o.Element
}

// GetElementOk returns a tuple with the Element field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetElementOk() (*string, bool) {
	if o == nil || IsNil(o.Element) {
		return nil, false
	}
	return o.Element, true
}

// HasElement returns a boolean if a field has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) HasElement() bool {
	if o != nil && !IsNil(o.Element) {
		return true
	}

	return false
}

// SetElement gets a reference to the given string and assigns it to the Element field.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) SetElement(v string) {
	o.Element = &v
}

// GetOrphan returns the Orphan field value if set, zero value otherwise.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetOrphan() bool {
	if o == nil || IsNil(o.Orphan) {
		var ret bool
		return ret
	}
	return *o.Orphan
}

// GetOrphanOk returns a tuple with the Orphan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetOrphanOk() (*bool, bool) {
	if o == nil || IsNil(o.Orphan) {
		return nil, false
	}
	return o.Orphan, true
}

// HasOrphan returns a boolean if a field has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) HasOrphan() bool {
	if o != nil && !IsNil(o.Orphan) {
		return true
	}

	return false
}

// SetOrphan gets a reference to the given bool and assigns it to the Orphan field.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) SetOrphan(v bool) {
	o.Orphan = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetPlacement() string {
	if o == nil || IsNil(o.Placement) {
		var ret string
		return ret
	}
	return *o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetPlacementOk() (*string, bool) {
	if o == nil || IsNil(o.Placement) {
		return nil, false
	}
	return o.Placement, true
}

// HasPlacement returns a boolean if a field has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) HasPlacement() bool {
	if o != nil && !IsNil(o.Placement) {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given string and assigns it to the Placement field.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) SetPlacement(v string) {
	o.Placement = &v
}

// GetReflex returns the Reflex field value if set, zero value otherwise.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetReflex() bool {
	if o == nil || IsNil(o.Reflex) {
		var ret bool
		return ret
	}
	return *o.Reflex
}

// GetReflexOk returns a tuple with the Reflex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetReflexOk() (*bool, bool) {
	if o == nil || IsNil(o.Reflex) {
		return nil, false
	}
	return o.Reflex, true
}

// HasReflex returns a boolean if a field has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) HasReflex() bool {
	if o != nil && !IsNil(o.Reflex) {
		return true
	}

	return false
}

// SetReflex gets a reference to the given bool and assigns it to the Reflex field.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) SetReflex(v bool) {
	o.Reflex = &v
}

// GetStepid returns the Stepid field value if set, zero value otherwise.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetStepid() int32 {
	if o == nil || IsNil(o.Stepid) {
		var ret int32
		return ret
	}
	return *o.Stepid
}

// GetStepidOk returns a tuple with the Stepid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetStepidOk() (*int32, bool) {
	if o == nil || IsNil(o.Stepid) {
		return nil, false
	}
	return o.Stepid, true
}

// HasStepid returns a boolean if a field has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) HasStepid() bool {
	if o != nil && !IsNil(o.Stepid) {
		return true
	}

	return false
}

// SetStepid gets a reference to the given int32 and assigns it to the Stepid field.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) SetStepid(v int32) {
	o.Stepid = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) SetTitle(v string) {
	o.Title = &v
}

func (o ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Backdrop) {
		toSerialize["backdrop"] = o.Backdrop
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Delay) {
		toSerialize["delay"] = o.Delay
	}
	if !IsNil(o.Element) {
		toSerialize["element"] = o.Element
	}
	if !IsNil(o.Orphan) {
		toSerialize["orphan"] = o.Orphan
	}
	if !IsNil(o.Placement) {
		toSerialize["placement"] = o.Placement
	}
	if !IsNil(o.Reflex) {
		toSerialize["reflex"] = o.Reflex
	}
	if !IsNil(o.Stepid) {
		toSerialize["stepid"] = o.Stepid
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner struct {
	value *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner
	isSet bool
}

func (v NullableToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) Get() *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner {
	return v.value
}

func (v *NullableToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) Set(val *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner(val *ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) *NullableToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner {
	return &NullableToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner{value: val, isSet: true}
}

func (v NullableToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


