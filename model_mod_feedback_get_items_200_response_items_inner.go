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

// checks if the ModFeedbackGetItems200ResponseItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModFeedbackGetItems200ResponseItemsInner{}

// ModFeedbackGetItems200ResponseItemsInner struct for ModFeedbackGetItems200ResponseItemsInner
type ModFeedbackGetItems200ResponseItemsInner struct {
	// The item id this item depend on.
	Dependitem *int32 `json:"dependitem,omitempty"`
	// The depend value.
	Dependvalue *string `json:"dependvalue,omitempty"`
	// The feedback instance id this records belongs to.
	Feedback *int32 `json:"feedback,omitempty"`
	// Whether it has a value or not.
	Hasvalue *int32 `json:"hasvalue,omitempty"`
	// The record id.
	Id *int32 `json:"id,omitempty"`
	Itemfiles []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner `json:"itemfiles,omitempty"`
	// The item position number
	Itemnumber *int32 `json:"itemnumber,omitempty"`
	// The item label.
	Label *string `json:"label,omitempty"`
	// The item name.
	Name *string `json:"name,omitempty"`
	// Different additional settings for the item (question).
	Options *string `json:"options,omitempty"`
	// Additional data that may be required by external functions
	Otherdata *string `json:"otherdata,omitempty"`
	// The position in the list of questions.
	Position *int32 `json:"position,omitempty"`
	// The text describing the item or the available possible answers.
	Presentation *string `json:"presentation,omitempty"`
	// Whether is a item (question) required or not.
	Required *bool `json:"required,omitempty"`
	// If it belogns to a template, the template id.
	Template *int32 `json:"template,omitempty"`
	// The type of the item.
	Typ *string `json:"typ,omitempty"`
}

// NewModFeedbackGetItems200ResponseItemsInner instantiates a new ModFeedbackGetItems200ResponseItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModFeedbackGetItems200ResponseItemsInner() *ModFeedbackGetItems200ResponseItemsInner {
	this := ModFeedbackGetItems200ResponseItemsInner{}
	var dependitem int32 = 0
	this.Dependitem = &dependitem
	var feedback int32 = 0
	this.Feedback = &feedback
	var hasvalue int32 = 0
	this.Hasvalue = &hasvalue
	var position int32 = 0
	this.Position = &position
	var required bool = 0
	this.Required = &required
	var template int32 = 0
	this.Template = &template
	return &this
}

// NewModFeedbackGetItems200ResponseItemsInnerWithDefaults instantiates a new ModFeedbackGetItems200ResponseItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModFeedbackGetItems200ResponseItemsInnerWithDefaults() *ModFeedbackGetItems200ResponseItemsInner {
	this := ModFeedbackGetItems200ResponseItemsInner{}
	var dependitem int32 = 0
	this.Dependitem = &dependitem
	var feedback int32 = 0
	this.Feedback = &feedback
	var hasvalue int32 = 0
	this.Hasvalue = &hasvalue
	var position int32 = 0
	this.Position = &position
	var required bool = 0
	this.Required = &required
	var template int32 = 0
	this.Template = &template
	return &this
}

// GetDependitem returns the Dependitem field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetDependitem() int32 {
	if o == nil || IsNil(o.Dependitem) {
		var ret int32
		return ret
	}
	return *o.Dependitem
}

// GetDependitemOk returns a tuple with the Dependitem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetDependitemOk() (*int32, bool) {
	if o == nil || IsNil(o.Dependitem) {
		return nil, false
	}
	return o.Dependitem, true
}

// HasDependitem returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasDependitem() bool {
	if o != nil && !IsNil(o.Dependitem) {
		return true
	}

	return false
}

// SetDependitem gets a reference to the given int32 and assigns it to the Dependitem field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetDependitem(v int32) {
	o.Dependitem = &v
}

// GetDependvalue returns the Dependvalue field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetDependvalue() string {
	if o == nil || IsNil(o.Dependvalue) {
		var ret string
		return ret
	}
	return *o.Dependvalue
}

// GetDependvalueOk returns a tuple with the Dependvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetDependvalueOk() (*string, bool) {
	if o == nil || IsNil(o.Dependvalue) {
		return nil, false
	}
	return o.Dependvalue, true
}

// HasDependvalue returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasDependvalue() bool {
	if o != nil && !IsNil(o.Dependvalue) {
		return true
	}

	return false
}

// SetDependvalue gets a reference to the given string and assigns it to the Dependvalue field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetDependvalue(v string) {
	o.Dependvalue = &v
}

// GetFeedback returns the Feedback field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetFeedback() int32 {
	if o == nil || IsNil(o.Feedback) {
		var ret int32
		return ret
	}
	return *o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetFeedbackOk() (*int32, bool) {
	if o == nil || IsNil(o.Feedback) {
		return nil, false
	}
	return o.Feedback, true
}

// HasFeedback returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasFeedback() bool {
	if o != nil && !IsNil(o.Feedback) {
		return true
	}

	return false
}

// SetFeedback gets a reference to the given int32 and assigns it to the Feedback field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetFeedback(v int32) {
	o.Feedback = &v
}

// GetHasvalue returns the Hasvalue field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetHasvalue() int32 {
	if o == nil || IsNil(o.Hasvalue) {
		var ret int32
		return ret
	}
	return *o.Hasvalue
}

// GetHasvalueOk returns a tuple with the Hasvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetHasvalueOk() (*int32, bool) {
	if o == nil || IsNil(o.Hasvalue) {
		return nil, false
	}
	return o.Hasvalue, true
}

// HasHasvalue returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasHasvalue() bool {
	if o != nil && !IsNil(o.Hasvalue) {
		return true
	}

	return false
}

// SetHasvalue gets a reference to the given int32 and assigns it to the Hasvalue field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetHasvalue(v int32) {
	o.Hasvalue = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetId(v int32) {
	o.Id = &v
}

// GetItemfiles returns the Itemfiles field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetItemfiles() []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner {
	if o == nil || IsNil(o.Itemfiles) {
		var ret []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner
		return ret
	}
	return o.Itemfiles
}

// GetItemfilesOk returns a tuple with the Itemfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetItemfilesOk() ([]ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner, bool) {
	if o == nil || IsNil(o.Itemfiles) {
		return nil, false
	}
	return o.Itemfiles, true
}

// HasItemfiles returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasItemfiles() bool {
	if o != nil && !IsNil(o.Itemfiles) {
		return true
	}

	return false
}

// SetItemfiles gets a reference to the given []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner and assigns it to the Itemfiles field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetItemfiles(v []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner) {
	o.Itemfiles = v
}

// GetItemnumber returns the Itemnumber field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetItemnumber() int32 {
	if o == nil || IsNil(o.Itemnumber) {
		var ret int32
		return ret
	}
	return *o.Itemnumber
}

// GetItemnumberOk returns a tuple with the Itemnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetItemnumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Itemnumber) {
		return nil, false
	}
	return o.Itemnumber, true
}

// HasItemnumber returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasItemnumber() bool {
	if o != nil && !IsNil(o.Itemnumber) {
		return true
	}

	return false
}

// SetItemnumber gets a reference to the given int32 and assigns it to the Itemnumber field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetItemnumber(v int32) {
	o.Itemnumber = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetOptions() string {
	if o == nil || IsNil(o.Options) {
		var ret string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetOptionsOk() (*string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given string and assigns it to the Options field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetOptions(v string) {
	o.Options = &v
}

// GetOtherdata returns the Otherdata field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetOtherdata() string {
	if o == nil || IsNil(o.Otherdata) {
		var ret string
		return ret
	}
	return *o.Otherdata
}

// GetOtherdataOk returns a tuple with the Otherdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetOtherdataOk() (*string, bool) {
	if o == nil || IsNil(o.Otherdata) {
		return nil, false
	}
	return o.Otherdata, true
}

// HasOtherdata returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasOtherdata() bool {
	if o != nil && !IsNil(o.Otherdata) {
		return true
	}

	return false
}

// SetOtherdata gets a reference to the given string and assigns it to the Otherdata field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetOtherdata(v string) {
	o.Otherdata = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetPosition(v int32) {
	o.Position = &v
}

// GetPresentation returns the Presentation field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetPresentation() string {
	if o == nil || IsNil(o.Presentation) {
		var ret string
		return ret
	}
	return *o.Presentation
}

// GetPresentationOk returns a tuple with the Presentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetPresentationOk() (*string, bool) {
	if o == nil || IsNil(o.Presentation) {
		return nil, false
	}
	return o.Presentation, true
}

// HasPresentation returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasPresentation() bool {
	if o != nil && !IsNil(o.Presentation) {
		return true
	}

	return false
}

// SetPresentation gets a reference to the given string and assigns it to the Presentation field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetPresentation(v string) {
	o.Presentation = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetRequired(v bool) {
	o.Required = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetTemplate() int32 {
	if o == nil || IsNil(o.Template) {
		var ret int32
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetTemplateOk() (*int32, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given int32 and assigns it to the Template field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetTemplate(v int32) {
	o.Template = &v
}

// GetTyp returns the Typ field value if set, zero value otherwise.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetTyp() string {
	if o == nil || IsNil(o.Typ) {
		var ret string
		return ret
	}
	return *o.Typ
}

// GetTypOk returns a tuple with the Typ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) GetTypOk() (*string, bool) {
	if o == nil || IsNil(o.Typ) {
		return nil, false
	}
	return o.Typ, true
}

// HasTyp returns a boolean if a field has been set.
func (o *ModFeedbackGetItems200ResponseItemsInner) HasTyp() bool {
	if o != nil && !IsNil(o.Typ) {
		return true
	}

	return false
}

// SetTyp gets a reference to the given string and assigns it to the Typ field.
func (o *ModFeedbackGetItems200ResponseItemsInner) SetTyp(v string) {
	o.Typ = &v
}

func (o ModFeedbackGetItems200ResponseItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModFeedbackGetItems200ResponseItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dependitem) {
		toSerialize["dependitem"] = o.Dependitem
	}
	if !IsNil(o.Dependvalue) {
		toSerialize["dependvalue"] = o.Dependvalue
	}
	if !IsNil(o.Feedback) {
		toSerialize["feedback"] = o.Feedback
	}
	if !IsNil(o.Hasvalue) {
		toSerialize["hasvalue"] = o.Hasvalue
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Itemfiles) {
		toSerialize["itemfiles"] = o.Itemfiles
	}
	if !IsNil(o.Itemnumber) {
		toSerialize["itemnumber"] = o.Itemnumber
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Otherdata) {
		toSerialize["otherdata"] = o.Otherdata
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Presentation) {
		toSerialize["presentation"] = o.Presentation
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Typ) {
		toSerialize["typ"] = o.Typ
	}
	return toSerialize, nil
}

type NullableModFeedbackGetItems200ResponseItemsInner struct {
	value *ModFeedbackGetItems200ResponseItemsInner
	isSet bool
}

func (v NullableModFeedbackGetItems200ResponseItemsInner) Get() *ModFeedbackGetItems200ResponseItemsInner {
	return v.value
}

func (v *NullableModFeedbackGetItems200ResponseItemsInner) Set(val *ModFeedbackGetItems200ResponseItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModFeedbackGetItems200ResponseItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModFeedbackGetItems200ResponseItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModFeedbackGetItems200ResponseItemsInner(val *ModFeedbackGetItems200ResponseItemsInner) *NullableModFeedbackGetItems200ResponseItemsInner {
	return &NullableModFeedbackGetItems200ResponseItemsInner{value: val, isSet: true}
}

func (v NullableModFeedbackGetItems200ResponseItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModFeedbackGetItems200ResponseItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


