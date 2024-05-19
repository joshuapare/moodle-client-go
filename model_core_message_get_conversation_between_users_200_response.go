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

// checks if the CoreMessageGetConversationBetweenUsers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageGetConversationBetweenUsers200Response{}

// CoreMessageGetConversationBetweenUsers200Response struct for CoreMessageGetConversationBetweenUsers200Response
type CoreMessageGetConversationBetweenUsers200Response struct {
	// If the user can delete messages in the conversation for all users
	Candeletemessagesforallusers *bool `json:"candeletemessagesforallusers,omitempty"`
	// The conversation id
	Id int32 `json:"id"`
	// A link to the conversation picture, if set
	Imageurl *string `json:"imageurl,omitempty"`
	// If the user marked this conversation as a favourite
	Isfavourite bool `json:"isfavourite"`
	// If the user muted this conversation
	Ismuted bool `json:"ismuted"`
	// If the user has read all messages in the conversation
	Isread bool `json:"isread"`
	// Total number of conversation members
	Membercount int32 `json:"membercount"`
	Members []CoreMessageGetConversationBetweenUsers200ResponseMembersInner `json:"members"`
	Messages []CoreMessageGetConversationBetweenUsers200ResponseMessagesInner `json:"messages"`
	// The conversation name, if set
	Name *string `json:"name,omitempty"`
	// A subtitle for the conversation name, if set
	Subname *string `json:"subname,omitempty"`
	// The type of the conversation (1=individual,2=group,3=self)
	Type int32 `json:"type"`
	// The number of unread messages in this conversation
	Unreadcount *int32 `json:"unreadcount,omitempty"`
}

type _CoreMessageGetConversationBetweenUsers200Response CoreMessageGetConversationBetweenUsers200Response

// NewCoreMessageGetConversationBetweenUsers200Response instantiates a new CoreMessageGetConversationBetweenUsers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageGetConversationBetweenUsers200Response(id int32, isfavourite bool, ismuted bool, isread bool, membercount int32, members []CoreMessageGetConversationBetweenUsers200ResponseMembersInner, messages []CoreMessageGetConversationBetweenUsers200ResponseMessagesInner, type_ int32) *CoreMessageGetConversationBetweenUsers200Response {
	this := CoreMessageGetConversationBetweenUsers200Response{}
	var candeletemessagesforallusers bool = false
	this.Candeletemessagesforallusers = &candeletemessagesforallusers
	this.Id = id
	this.Isfavourite = isfavourite
	this.Ismuted = ismuted
	this.Isread = isread
	this.Membercount = membercount
	this.Members = members
	this.Messages = messages
	this.Type = type_
	return &this
}

// NewCoreMessageGetConversationBetweenUsers200ResponseWithDefaults instantiates a new CoreMessageGetConversationBetweenUsers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageGetConversationBetweenUsers200ResponseWithDefaults() *CoreMessageGetConversationBetweenUsers200Response {
	this := CoreMessageGetConversationBetweenUsers200Response{}
	var candeletemessagesforallusers bool = false
	this.Candeletemessagesforallusers = &candeletemessagesforallusers
	return &this
}

// GetCandeletemessagesforallusers returns the Candeletemessagesforallusers field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetCandeletemessagesforallusers() bool {
	if o == nil || IsNil(o.Candeletemessagesforallusers) {
		var ret bool
		return ret
	}
	return *o.Candeletemessagesforallusers
}

// GetCandeletemessagesforallusersOk returns a tuple with the Candeletemessagesforallusers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetCandeletemessagesforallusersOk() (*bool, bool) {
	if o == nil || IsNil(o.Candeletemessagesforallusers) {
		return nil, false
	}
	return o.Candeletemessagesforallusers, true
}

// HasCandeletemessagesforallusers returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) HasCandeletemessagesforallusers() bool {
	if o != nil && !IsNil(o.Candeletemessagesforallusers) {
		return true
	}

	return false
}

// SetCandeletemessagesforallusers gets a reference to the given bool and assigns it to the Candeletemessagesforallusers field.
func (o *CoreMessageGetConversationBetweenUsers200Response) SetCandeletemessagesforallusers(v bool) {
	o.Candeletemessagesforallusers = &v
}

// GetId returns the Id field value
func (o *CoreMessageGetConversationBetweenUsers200Response) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CoreMessageGetConversationBetweenUsers200Response) SetId(v int32) {
	o.Id = v
}

// GetImageurl returns the Imageurl field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetImageurl() string {
	if o == nil || IsNil(o.Imageurl) {
		var ret string
		return ret
	}
	return *o.Imageurl
}

// GetImageurlOk returns a tuple with the Imageurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetImageurlOk() (*string, bool) {
	if o == nil || IsNil(o.Imageurl) {
		return nil, false
	}
	return o.Imageurl, true
}

// HasImageurl returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) HasImageurl() bool {
	if o != nil && !IsNil(o.Imageurl) {
		return true
	}

	return false
}

// SetImageurl gets a reference to the given string and assigns it to the Imageurl field.
func (o *CoreMessageGetConversationBetweenUsers200Response) SetImageurl(v string) {
	o.Imageurl = &v
}

// GetIsfavourite returns the Isfavourite field value
func (o *CoreMessageGetConversationBetweenUsers200Response) GetIsfavourite() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Isfavourite
}

// GetIsfavouriteOk returns a tuple with the Isfavourite field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetIsfavouriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isfavourite, true
}

// SetIsfavourite sets field value
func (o *CoreMessageGetConversationBetweenUsers200Response) SetIsfavourite(v bool) {
	o.Isfavourite = v
}

// GetIsmuted returns the Ismuted field value
func (o *CoreMessageGetConversationBetweenUsers200Response) GetIsmuted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ismuted
}

// GetIsmutedOk returns a tuple with the Ismuted field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetIsmutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ismuted, true
}

// SetIsmuted sets field value
func (o *CoreMessageGetConversationBetweenUsers200Response) SetIsmuted(v bool) {
	o.Ismuted = v
}

// GetIsread returns the Isread field value
func (o *CoreMessageGetConversationBetweenUsers200Response) GetIsread() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Isread
}

// GetIsreadOk returns a tuple with the Isread field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetIsreadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isread, true
}

// SetIsread sets field value
func (o *CoreMessageGetConversationBetweenUsers200Response) SetIsread(v bool) {
	o.Isread = v
}

// GetMembercount returns the Membercount field value
func (o *CoreMessageGetConversationBetweenUsers200Response) GetMembercount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Membercount
}

// GetMembercountOk returns a tuple with the Membercount field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetMembercountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Membercount, true
}

// SetMembercount sets field value
func (o *CoreMessageGetConversationBetweenUsers200Response) SetMembercount(v int32) {
	o.Membercount = v
}

// GetMembers returns the Members field value
func (o *CoreMessageGetConversationBetweenUsers200Response) GetMembers() []CoreMessageGetConversationBetweenUsers200ResponseMembersInner {
	if o == nil {
		var ret []CoreMessageGetConversationBetweenUsers200ResponseMembersInner
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetMembersOk() ([]CoreMessageGetConversationBetweenUsers200ResponseMembersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *CoreMessageGetConversationBetweenUsers200Response) SetMembers(v []CoreMessageGetConversationBetweenUsers200ResponseMembersInner) {
	o.Members = v
}

// GetMessages returns the Messages field value
func (o *CoreMessageGetConversationBetweenUsers200Response) GetMessages() []CoreMessageGetConversationBetweenUsers200ResponseMessagesInner {
	if o == nil {
		var ret []CoreMessageGetConversationBetweenUsers200ResponseMessagesInner
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetMessagesOk() ([]CoreMessageGetConversationBetweenUsers200ResponseMessagesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *CoreMessageGetConversationBetweenUsers200Response) SetMessages(v []CoreMessageGetConversationBetweenUsers200ResponseMessagesInner) {
	o.Messages = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreMessageGetConversationBetweenUsers200Response) SetName(v string) {
	o.Name = &v
}

// GetSubname returns the Subname field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetSubname() string {
	if o == nil || IsNil(o.Subname) {
		var ret string
		return ret
	}
	return *o.Subname
}

// GetSubnameOk returns a tuple with the Subname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetSubnameOk() (*string, bool) {
	if o == nil || IsNil(o.Subname) {
		return nil, false
	}
	return o.Subname, true
}

// HasSubname returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) HasSubname() bool {
	if o != nil && !IsNil(o.Subname) {
		return true
	}

	return false
}

// SetSubname gets a reference to the given string and assigns it to the Subname field.
func (o *CoreMessageGetConversationBetweenUsers200Response) SetSubname(v string) {
	o.Subname = &v
}

// GetType returns the Type field value
func (o *CoreMessageGetConversationBetweenUsers200Response) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CoreMessageGetConversationBetweenUsers200Response) SetType(v int32) {
	o.Type = v
}

// GetUnreadcount returns the Unreadcount field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetUnreadcount() int32 {
	if o == nil || IsNil(o.Unreadcount) {
		var ret int32
		return ret
	}
	return *o.Unreadcount
}

// GetUnreadcountOk returns a tuple with the Unreadcount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) GetUnreadcountOk() (*int32, bool) {
	if o == nil || IsNil(o.Unreadcount) {
		return nil, false
	}
	return o.Unreadcount, true
}

// HasUnreadcount returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200Response) HasUnreadcount() bool {
	if o != nil && !IsNil(o.Unreadcount) {
		return true
	}

	return false
}

// SetUnreadcount gets a reference to the given int32 and assigns it to the Unreadcount field.
func (o *CoreMessageGetConversationBetweenUsers200Response) SetUnreadcount(v int32) {
	o.Unreadcount = &v
}

func (o CoreMessageGetConversationBetweenUsers200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageGetConversationBetweenUsers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Candeletemessagesforallusers) {
		toSerialize["candeletemessagesforallusers"] = o.Candeletemessagesforallusers
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Imageurl) {
		toSerialize["imageurl"] = o.Imageurl
	}
	toSerialize["isfavourite"] = o.Isfavourite
	toSerialize["ismuted"] = o.Ismuted
	toSerialize["isread"] = o.Isread
	toSerialize["membercount"] = o.Membercount
	toSerialize["members"] = o.Members
	toSerialize["messages"] = o.Messages
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Subname) {
		toSerialize["subname"] = o.Subname
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Unreadcount) {
		toSerialize["unreadcount"] = o.Unreadcount
	}
	return toSerialize, nil
}

func (o *CoreMessageGetConversationBetweenUsers200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"isfavourite",
		"ismuted",
		"isread",
		"membercount",
		"members",
		"messages",
		"type",
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

	varCoreMessageGetConversationBetweenUsers200Response := _CoreMessageGetConversationBetweenUsers200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageGetConversationBetweenUsers200Response)

	if err != nil {
		return err
	}

	*o = CoreMessageGetConversationBetweenUsers200Response(varCoreMessageGetConversationBetweenUsers200Response)

	return err
}

type NullableCoreMessageGetConversationBetweenUsers200Response struct {
	value *CoreMessageGetConversationBetweenUsers200Response
	isSet bool
}

func (v NullableCoreMessageGetConversationBetweenUsers200Response) Get() *CoreMessageGetConversationBetweenUsers200Response {
	return v.value
}

func (v *NullableCoreMessageGetConversationBetweenUsers200Response) Set(val *CoreMessageGetConversationBetweenUsers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageGetConversationBetweenUsers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageGetConversationBetweenUsers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageGetConversationBetweenUsers200Response(val *CoreMessageGetConversationBetweenUsers200Response) *NullableCoreMessageGetConversationBetweenUsers200Response {
	return &NullableCoreMessageGetConversationBetweenUsers200Response{value: val, isSet: true}
}

func (v NullableCoreMessageGetConversationBetweenUsers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageGetConversationBetweenUsers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


