/*
Kaspa REST-API server

This server is to communicate with kaspa network via REST-API

API version: tbd
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ParentHashModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParentHashModel{}

// ParentHashModel struct for ParentHashModel
type ParentHashModel struct {
	ParentHashes []string `json:"parentHashes,omitempty"`
}

// NewParentHashModel instantiates a new ParentHashModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParentHashModel() *ParentHashModel {
	this := ParentHashModel{}
	return &this
}

// NewParentHashModelWithDefaults instantiates a new ParentHashModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParentHashModelWithDefaults() *ParentHashModel {
	this := ParentHashModel{}
	return &this
}

// GetParentHashes returns the ParentHashes field value if set, zero value otherwise.
func (o *ParentHashModel) GetParentHashes() []string {
	if o == nil || IsNil(o.ParentHashes) {
		var ret []string
		return ret
	}
	return o.ParentHashes
}

// GetParentHashesOk returns a tuple with the ParentHashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentHashModel) GetParentHashesOk() ([]string, bool) {
	if o == nil || IsNil(o.ParentHashes) {
		return nil, false
	}
	return o.ParentHashes, true
}

// HasParentHashes returns a boolean if a field has been set.
func (o *ParentHashModel) HasParentHashes() bool {
	if o != nil && !IsNil(o.ParentHashes) {
		return true
	}

	return false
}

// SetParentHashes gets a reference to the given []string and assigns it to the ParentHashes field.
func (o *ParentHashModel) SetParentHashes(v []string) {
	o.ParentHashes = v
}

func (o ParentHashModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParentHashModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParentHashes) {
		toSerialize["parentHashes"] = o.ParentHashes
	}
	return toSerialize, nil
}

type NullableParentHashModel struct {
	value *ParentHashModel
	isSet bool
}

func (v NullableParentHashModel) Get() *ParentHashModel {
	return v.value
}

func (v *NullableParentHashModel) Set(val *ParentHashModel) {
	v.value = val
	v.isSet = true
}

func (v NullableParentHashModel) IsSet() bool {
	return v.isSet
}

func (v *NullableParentHashModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParentHashModel(val *ParentHashModel) *NullableParentHashModel {
	return &NullableParentHashModel{value: val, isSet: true}
}

func (v NullableParentHashModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParentHashModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


