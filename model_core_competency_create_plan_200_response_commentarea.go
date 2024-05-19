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

// checks if the CoreCompetencyCreatePlan200ResponseCommentarea type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCompetencyCreatePlan200ResponseCommentarea{}

// CoreCompetencyCreatePlan200ResponseCommentarea struct for CoreCompetencyCreatePlan200ResponseCommentarea
type CoreCompetencyCreatePlan200ResponseCommentarea struct {
	// autostart
	Autostart bool `json:"autostart"`
	// canpost
	Canpost bool `json:"canpost"`
	// canpostorhascomments
	Canpostorhascomments bool `json:"canpostorhascomments"`
	// canview
	Canview bool `json:"canview"`
	// cid
	Cid string `json:"cid"`
	// collapsediconkey
	Collapsediconkey string `json:"collapsediconkey"`
	// commentarea
	Commentarea string `json:"commentarea"`
	// component
	Component string `json:"component"`
	// contextid
	Contextid int32 `json:"contextid"`
	// count
	Count int32 `json:"count"`
	// courseid
	Courseid int32 `json:"courseid"`
	// displaycancel
	Displaycancel bool `json:"displaycancel"`
	// displaytotalcount
	Displaytotalcount bool `json:"displaytotalcount"`
	// fullwidth
	Fullwidth bool `json:"fullwidth"`
	// itemid
	Itemid int32 `json:"itemid"`
	// linktext
	Linktext string `json:"linktext"`
	// notoggle
	Notoggle bool `json:"notoggle"`
	// template
	Template string `json:"template"`
}

type _CoreCompetencyCreatePlan200ResponseCommentarea CoreCompetencyCreatePlan200ResponseCommentarea

// NewCoreCompetencyCreatePlan200ResponseCommentarea instantiates a new CoreCompetencyCreatePlan200ResponseCommentarea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCompetencyCreatePlan200ResponseCommentarea(autostart bool, canpost bool, canpostorhascomments bool, canview bool, cid string, collapsediconkey string, commentarea string, component string, contextid int32, count int32, courseid int32, displaycancel bool, displaytotalcount bool, fullwidth bool, itemid int32, linktext string, notoggle bool, template string) *CoreCompetencyCreatePlan200ResponseCommentarea {
	this := CoreCompetencyCreatePlan200ResponseCommentarea{}
	this.Autostart = autostart
	this.Canpost = canpost
	this.Canpostorhascomments = canpostorhascomments
	this.Canview = canview
	this.Cid = cid
	this.Collapsediconkey = collapsediconkey
	this.Commentarea = commentarea
	this.Component = component
	this.Contextid = contextid
	this.Count = count
	this.Courseid = courseid
	this.Displaycancel = displaycancel
	this.Displaytotalcount = displaytotalcount
	this.Fullwidth = fullwidth
	this.Itemid = itemid
	this.Linktext = linktext
	this.Notoggle = notoggle
	this.Template = template
	return &this
}

// NewCoreCompetencyCreatePlan200ResponseCommentareaWithDefaults instantiates a new CoreCompetencyCreatePlan200ResponseCommentarea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCompetencyCreatePlan200ResponseCommentareaWithDefaults() *CoreCompetencyCreatePlan200ResponseCommentarea {
	this := CoreCompetencyCreatePlan200ResponseCommentarea{}
	var autostart bool = null
	this.Autostart = autostart
	var canpost bool = null
	this.Canpost = canpost
	var canpostorhascomments bool = null
	this.Canpostorhascomments = canpostorhascomments
	var canview bool = null
	this.Canview = canview
	var cid string = "null"
	this.Cid = cid
	var collapsediconkey string = "null"
	this.Collapsediconkey = collapsediconkey
	var commentarea string = "null"
	this.Commentarea = commentarea
	var count int32 = null
	this.Count = count
	var displaycancel bool = null
	this.Displaycancel = displaycancel
	var displaytotalcount bool = null
	this.Displaytotalcount = displaytotalcount
	var fullwidth bool = null
	this.Fullwidth = fullwidth
	var itemid int32 = null
	this.Itemid = itemid
	var linktext string = "null"
	this.Linktext = linktext
	var notoggle bool = null
	this.Notoggle = notoggle
	var template string = "null"
	this.Template = template
	return &this
}

// GetAutostart returns the Autostart field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetAutostart() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Autostart
}

// GetAutostartOk returns a tuple with the Autostart field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetAutostartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Autostart, true
}

// SetAutostart sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetAutostart(v bool) {
	o.Autostart = v
}

// GetCanpost returns the Canpost field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCanpost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canpost
}

// GetCanpostOk returns a tuple with the Canpost field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCanpostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canpost, true
}

// SetCanpost sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetCanpost(v bool) {
	o.Canpost = v
}

// GetCanpostorhascomments returns the Canpostorhascomments field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCanpostorhascomments() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canpostorhascomments
}

// GetCanpostorhascommentsOk returns a tuple with the Canpostorhascomments field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCanpostorhascommentsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canpostorhascomments, true
}

// SetCanpostorhascomments sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetCanpostorhascomments(v bool) {
	o.Canpostorhascomments = v
}

// GetCanview returns the Canview field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCanview() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canview
}

// GetCanviewOk returns a tuple with the Canview field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCanviewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canview, true
}

// SetCanview sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetCanview(v bool) {
	o.Canview = v
}

// GetCid returns the Cid field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetCid(v string) {
	o.Cid = v
}

// GetCollapsediconkey returns the Collapsediconkey field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCollapsediconkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collapsediconkey
}

// GetCollapsediconkeyOk returns a tuple with the Collapsediconkey field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCollapsediconkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collapsediconkey, true
}

// SetCollapsediconkey sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetCollapsediconkey(v string) {
	o.Collapsediconkey = v
}

// GetCommentarea returns the Commentarea field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCommentarea() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Commentarea
}

// GetCommentareaOk returns a tuple with the Commentarea field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCommentareaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commentarea, true
}

// SetCommentarea sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetCommentarea(v string) {
	o.Commentarea = v
}

// GetComponent returns the Component field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetComponent(v string) {
	o.Component = v
}

// GetContextid returns the Contextid field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetContextid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Contextid
}

// GetContextidOk returns a tuple with the Contextid field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetContextidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contextid, true
}

// SetContextid sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetContextid(v int32) {
	o.Contextid = v
}

// GetCount returns the Count field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetCount(v int32) {
	o.Count = v
}

// GetCourseid returns the Courseid field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCourseid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetCourseidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Courseid, true
}

// SetCourseid sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetCourseid(v int32) {
	o.Courseid = v
}

// GetDisplaycancel returns the Displaycancel field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetDisplaycancel() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Displaycancel
}

// GetDisplaycancelOk returns a tuple with the Displaycancel field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetDisplaycancelOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Displaycancel, true
}

// SetDisplaycancel sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetDisplaycancel(v bool) {
	o.Displaycancel = v
}

// GetDisplaytotalcount returns the Displaytotalcount field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetDisplaytotalcount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Displaytotalcount
}

// GetDisplaytotalcountOk returns a tuple with the Displaytotalcount field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetDisplaytotalcountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Displaytotalcount, true
}

// SetDisplaytotalcount sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetDisplaytotalcount(v bool) {
	o.Displaytotalcount = v
}

// GetFullwidth returns the Fullwidth field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetFullwidth() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Fullwidth
}

// GetFullwidthOk returns a tuple with the Fullwidth field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetFullwidthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fullwidth, true
}

// SetFullwidth sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetFullwidth(v bool) {
	o.Fullwidth = v
}

// GetItemid returns the Itemid field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetItemid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Itemid
}

// GetItemidOk returns a tuple with the Itemid field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetItemidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Itemid, true
}

// SetItemid sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetItemid(v int32) {
	o.Itemid = v
}

// GetLinktext returns the Linktext field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetLinktext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Linktext
}

// GetLinktextOk returns a tuple with the Linktext field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetLinktextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Linktext, true
}

// SetLinktext sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetLinktext(v string) {
	o.Linktext = v
}

// GetNotoggle returns the Notoggle field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetNotoggle() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Notoggle
}

// GetNotoggleOk returns a tuple with the Notoggle field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetNotoggleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notoggle, true
}

// SetNotoggle sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetNotoggle(v bool) {
	o.Notoggle = v
}

// GetTemplate returns the Template field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *CoreCompetencyCreatePlan200ResponseCommentarea) SetTemplate(v string) {
	o.Template = v
}

func (o CoreCompetencyCreatePlan200ResponseCommentarea) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCompetencyCreatePlan200ResponseCommentarea) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["autostart"] = o.Autostart
	toSerialize["canpost"] = o.Canpost
	toSerialize["canpostorhascomments"] = o.Canpostorhascomments
	toSerialize["canview"] = o.Canview
	toSerialize["cid"] = o.Cid
	toSerialize["collapsediconkey"] = o.Collapsediconkey
	toSerialize["commentarea"] = o.Commentarea
	toSerialize["component"] = o.Component
	toSerialize["contextid"] = o.Contextid
	toSerialize["count"] = o.Count
	toSerialize["courseid"] = o.Courseid
	toSerialize["displaycancel"] = o.Displaycancel
	toSerialize["displaytotalcount"] = o.Displaytotalcount
	toSerialize["fullwidth"] = o.Fullwidth
	toSerialize["itemid"] = o.Itemid
	toSerialize["linktext"] = o.Linktext
	toSerialize["notoggle"] = o.Notoggle
	toSerialize["template"] = o.Template
	return toSerialize, nil
}

func (o *CoreCompetencyCreatePlan200ResponseCommentarea) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"autostart",
		"canpost",
		"canpostorhascomments",
		"canview",
		"cid",
		"collapsediconkey",
		"commentarea",
		"component",
		"contextid",
		"count",
		"courseid",
		"displaycancel",
		"displaytotalcount",
		"fullwidth",
		"itemid",
		"linktext",
		"notoggle",
		"template",
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

	varCoreCompetencyCreatePlan200ResponseCommentarea := _CoreCompetencyCreatePlan200ResponseCommentarea{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCompetencyCreatePlan200ResponseCommentarea)

	if err != nil {
		return err
	}

	*o = CoreCompetencyCreatePlan200ResponseCommentarea(varCoreCompetencyCreatePlan200ResponseCommentarea)

	return err
}

type NullableCoreCompetencyCreatePlan200ResponseCommentarea struct {
	value *CoreCompetencyCreatePlan200ResponseCommentarea
	isSet bool
}

func (v NullableCoreCompetencyCreatePlan200ResponseCommentarea) Get() *CoreCompetencyCreatePlan200ResponseCommentarea {
	return v.value
}

func (v *NullableCoreCompetencyCreatePlan200ResponseCommentarea) Set(val *CoreCompetencyCreatePlan200ResponseCommentarea) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCompetencyCreatePlan200ResponseCommentarea) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCompetencyCreatePlan200ResponseCommentarea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCompetencyCreatePlan200ResponseCommentarea(val *CoreCompetencyCreatePlan200ResponseCommentarea) *NullableCoreCompetencyCreatePlan200ResponseCommentarea {
	return &NullableCoreCompetencyCreatePlan200ResponseCommentarea{value: val, isSet: true}
}

func (v NullableCoreCompetencyCreatePlan200ResponseCommentarea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCompetencyCreatePlan200ResponseCommentarea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

