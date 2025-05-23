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

// checks if the BlockTxOutputScriptPublicKeyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockTxOutputScriptPublicKeyModel{}

// BlockTxOutputScriptPublicKeyModel struct for BlockTxOutputScriptPublicKeyModel
type BlockTxOutputScriptPublicKeyModel struct {
	ScriptPublicKey *string `json:"scriptPublicKey,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewBlockTxOutputScriptPublicKeyModel instantiates a new BlockTxOutputScriptPublicKeyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockTxOutputScriptPublicKeyModel() *BlockTxOutputScriptPublicKeyModel {
	this := BlockTxOutputScriptPublicKeyModel{}
	return &this
}

// NewBlockTxOutputScriptPublicKeyModelWithDefaults instantiates a new BlockTxOutputScriptPublicKeyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockTxOutputScriptPublicKeyModelWithDefaults() *BlockTxOutputScriptPublicKeyModel {
	this := BlockTxOutputScriptPublicKeyModel{}
	return &this
}

// GetScriptPublicKey returns the ScriptPublicKey field value if set, zero value otherwise.
func (o *BlockTxOutputScriptPublicKeyModel) GetScriptPublicKey() string {
	if o == nil || IsNil(o.ScriptPublicKey) {
		var ret string
		return ret
	}
	return *o.ScriptPublicKey
}

// GetScriptPublicKeyOk returns a tuple with the ScriptPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxOutputScriptPublicKeyModel) GetScriptPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ScriptPublicKey) {
		return nil, false
	}
	return o.ScriptPublicKey, true
}

// HasScriptPublicKey returns a boolean if a field has been set.
func (o *BlockTxOutputScriptPublicKeyModel) HasScriptPublicKey() bool {
	if o != nil && !IsNil(o.ScriptPublicKey) {
		return true
	}

	return false
}

// SetScriptPublicKey gets a reference to the given string and assigns it to the ScriptPublicKey field.
func (o *BlockTxOutputScriptPublicKeyModel) SetScriptPublicKey(v string) {
	o.ScriptPublicKey = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BlockTxOutputScriptPublicKeyModel) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxOutputScriptPublicKeyModel) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BlockTxOutputScriptPublicKeyModel) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *BlockTxOutputScriptPublicKeyModel) SetVersion(v int32) {
	o.Version = &v
}

func (o BlockTxOutputScriptPublicKeyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockTxOutputScriptPublicKeyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScriptPublicKey) {
		toSerialize["scriptPublicKey"] = o.ScriptPublicKey
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableBlockTxOutputScriptPublicKeyModel struct {
	value *BlockTxOutputScriptPublicKeyModel
	isSet bool
}

func (v NullableBlockTxOutputScriptPublicKeyModel) Get() *BlockTxOutputScriptPublicKeyModel {
	return v.value
}

func (v *NullableBlockTxOutputScriptPublicKeyModel) Set(val *BlockTxOutputScriptPublicKeyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockTxOutputScriptPublicKeyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockTxOutputScriptPublicKeyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockTxOutputScriptPublicKeyModel(val *BlockTxOutputScriptPublicKeyModel) *NullableBlockTxOutputScriptPublicKeyModel {
	return &NullableBlockTxOutputScriptPublicKeyModel{value: val, isSet: true}
}

func (v NullableBlockTxOutputScriptPublicKeyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockTxOutputScriptPublicKeyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


