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

// checks if the ModFeedbackGetAnalysis200ResponseItemsdataInnerItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModFeedbackGetAnalysis200ResponseItemsdataInnerItem{}

// ModFeedbackGetAnalysis200ResponseItemsdataInnerItem struct for ModFeedbackGetAnalysis200ResponseItemsdataInnerItem
type ModFeedbackGetAnalysis200ResponseItemsdataInnerItem struct {
	// The item id this item depend on.
	Dependitem int32 `json:"dependitem"`
	// The depend value.
	Dependvalue string `json:"dependvalue"`
	// The feedback instance id this records belongs to.
	Feedback int32 `json:"feedback"`
	// Whether it has a value or not.
	Hasvalue int32 `json:"hasvalue"`
	// The record id.
	Id int32 `json:"id"`
	Itemfiles []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner `json:"itemfiles"`
	// The item position number
	Itemnumber int32 `json:"itemnumber"`
	// The item label.
	Label string `json:"label"`
	// The item name.
	Name string `json:"name"`
	// Different additional settings for the item (question).
	Options string `json:"options"`
	// Additional data that may be required by external functions
	Otherdata string `json:"otherdata"`
	// The position in the list of questions.
	Position int32 `json:"position"`
	// The text describing the item or the available possible answers.
	Presentation string `json:"presentation"`
	// Whether is a item (question) required or not.
	Required bool `json:"required"`
	// If it belogns to a template, the template id.
	Template int32 `json:"template"`
	// The type of the item.
	Typ string `json:"typ"`
}

type _ModFeedbackGetAnalysis200ResponseItemsdataInnerItem ModFeedbackGetAnalysis200ResponseItemsdataInnerItem

// NewModFeedbackGetAnalysis200ResponseItemsdataInnerItem instantiates a new ModFeedbackGetAnalysis200ResponseItemsdataInnerItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModFeedbackGetAnalysis200ResponseItemsdataInnerItem(dependitem int32, dependvalue string, feedback int32, hasvalue int32, id int32, itemfiles []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner, itemnumber int32, label string, name string, options string, otherdata string, position int32, presentation string, required bool, template int32, typ string) *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem {
	this := ModFeedbackGetAnalysis200ResponseItemsdataInnerItem{}
	this.Dependitem = dependitem
	this.Dependvalue = dependvalue
	this.Feedback = feedback
	this.Hasvalue = hasvalue
	this.Id = id
	this.Itemfiles = itemfiles
	this.Itemnumber = itemnumber
	this.Label = label
	this.Name = name
	this.Options = options
	this.Otherdata = otherdata
	this.Position = position
	this.Presentation = presentation
	this.Required = required
	this.Template = template
	this.Typ = typ
	return &this
}

// NewModFeedbackGetAnalysis200ResponseItemsdataInnerItemWithDefaults instantiates a new ModFeedbackGetAnalysis200ResponseItemsdataInnerItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModFeedbackGetAnalysis200ResponseItemsdataInnerItemWithDefaults() *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem {
	this := ModFeedbackGetAnalysis200ResponseItemsdataInnerItem{}
	var dependitem int32 = 0
	this.Dependitem = dependitem
	var dependvalue string = "null"
	this.Dependvalue = dependvalue
	var feedback int32 = 0
	this.Feedback = feedback
	var hasvalue int32 = 0
	this.Hasvalue = hasvalue
	var id int32 = null
	this.Id = id
	var itemnumber int32 = null
	this.Itemnumber = itemnumber
	var label string = "null"
	this.Label = label
	var name string = "null"
	this.Name = name
	var options string = "null"
	this.Options = options
	var otherdata string = "null"
	this.Otherdata = otherdata
	var position int32 = 0
	this.Position = position
	var presentation string = "null"
	this.Presentation = presentation
	var required bool = 0
	this.Required = required
	var template int32 = 0
	this.Template = template
	var typ string = "null"
	this.Typ = typ
	return &this
}

// GetDependitem returns the Dependitem field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetDependitem() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Dependitem
}

// GetDependitemOk returns a tuple with the Dependitem field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetDependitemOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dependitem, true
}

// SetDependitem sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetDependitem(v int32) {
	o.Dependitem = v
}

// GetDependvalue returns the Dependvalue field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetDependvalue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dependvalue
}

// GetDependvalueOk returns a tuple with the Dependvalue field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetDependvalueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dependvalue, true
}

// SetDependvalue sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetDependvalue(v string) {
	o.Dependvalue = v
}

// GetFeedback returns the Feedback field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetFeedback() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetFeedbackOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Feedback, true
}

// SetFeedback sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetFeedback(v int32) {
	o.Feedback = v
}

// GetHasvalue returns the Hasvalue field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetHasvalue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Hasvalue
}

// GetHasvalueOk returns a tuple with the Hasvalue field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetHasvalueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hasvalue, true
}

// SetHasvalue sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetHasvalue(v int32) {
	o.Hasvalue = v
}

// GetId returns the Id field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetId(v int32) {
	o.Id = v
}

// GetItemfiles returns the Itemfiles field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetItemfiles() []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner {
	if o == nil {
		var ret []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner
		return ret
	}

	return o.Itemfiles
}

// GetItemfilesOk returns a tuple with the Itemfiles field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetItemfilesOk() ([]ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Itemfiles, true
}

// SetItemfiles sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetItemfiles(v []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner) {
	o.Itemfiles = v
}

// GetItemnumber returns the Itemnumber field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetItemnumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Itemnumber
}

// GetItemnumberOk returns a tuple with the Itemnumber field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetItemnumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Itemnumber, true
}

// SetItemnumber sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetItemnumber(v int32) {
	o.Itemnumber = v
}

// GetLabel returns the Label field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetLabel(v string) {
	o.Label = v
}

// GetName returns the Name field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetOptions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetOptionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetOptions(v string) {
	o.Options = v
}

// GetOtherdata returns the Otherdata field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetOtherdata() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Otherdata
}

// GetOtherdataOk returns a tuple with the Otherdata field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetOtherdataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Otherdata, true
}

// SetOtherdata sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetOtherdata(v string) {
	o.Otherdata = v
}

// GetPosition returns the Position field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetPosition(v int32) {
	o.Position = v
}

// GetPresentation returns the Presentation field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetPresentation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Presentation
}

// GetPresentationOk returns a tuple with the Presentation field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetPresentationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Presentation, true
}

// SetPresentation sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetPresentation(v string) {
	o.Presentation = v
}

// GetRequired returns the Required field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetRequired(v bool) {
	o.Required = v
}

// GetTemplate returns the Template field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetTemplate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetTemplateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetTemplate(v int32) {
	o.Template = v
}

// GetTyp returns the Typ field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetTyp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Typ
}

// GetTypOk returns a tuple with the Typ field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) GetTypOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typ, true
}

// SetTyp sets field value
func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) SetTyp(v string) {
	o.Typ = v
}

func (o ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dependitem"] = o.Dependitem
	toSerialize["dependvalue"] = o.Dependvalue
	toSerialize["feedback"] = o.Feedback
	toSerialize["hasvalue"] = o.Hasvalue
	toSerialize["id"] = o.Id
	toSerialize["itemfiles"] = o.Itemfiles
	toSerialize["itemnumber"] = o.Itemnumber
	toSerialize["label"] = o.Label
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	toSerialize["otherdata"] = o.Otherdata
	toSerialize["position"] = o.Position
	toSerialize["presentation"] = o.Presentation
	toSerialize["required"] = o.Required
	toSerialize["template"] = o.Template
	toSerialize["typ"] = o.Typ
	return toSerialize, nil
}

func (o *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dependitem",
		"dependvalue",
		"feedback",
		"hasvalue",
		"id",
		"itemfiles",
		"itemnumber",
		"label",
		"name",
		"options",
		"otherdata",
		"position",
		"presentation",
		"required",
		"template",
		"typ",
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

	varModFeedbackGetAnalysis200ResponseItemsdataInnerItem := _ModFeedbackGetAnalysis200ResponseItemsdataInnerItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModFeedbackGetAnalysis200ResponseItemsdataInnerItem)

	if err != nil {
		return err
	}

	*o = ModFeedbackGetAnalysis200ResponseItemsdataInnerItem(varModFeedbackGetAnalysis200ResponseItemsdataInnerItem)

	return err
}

type NullableModFeedbackGetAnalysis200ResponseItemsdataInnerItem struct {
	value *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem
	isSet bool
}

func (v NullableModFeedbackGetAnalysis200ResponseItemsdataInnerItem) Get() *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem {
	return v.value
}

func (v *NullableModFeedbackGetAnalysis200ResponseItemsdataInnerItem) Set(val *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) {
	v.value = val
	v.isSet = true
}

func (v NullableModFeedbackGetAnalysis200ResponseItemsdataInnerItem) IsSet() bool {
	return v.isSet
}

func (v *NullableModFeedbackGetAnalysis200ResponseItemsdataInnerItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModFeedbackGetAnalysis200ResponseItemsdataInnerItem(val *ModFeedbackGetAnalysis200ResponseItemsdataInnerItem) *NullableModFeedbackGetAnalysis200ResponseItemsdataInnerItem {
	return &NullableModFeedbackGetAnalysis200ResponseItemsdataInnerItem{value: val, isSet: true}
}

func (v NullableModFeedbackGetAnalysis200ResponseItemsdataInnerItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModFeedbackGetAnalysis200ResponseItemsdataInnerItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


