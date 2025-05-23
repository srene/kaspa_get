/*
Kaspa REST-API server

This server is to communicate with kaspa network via REST-API

API version: tbd
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BlockTxVerboseDataModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockTxVerboseDataModel{}

// BlockTxVerboseDataModel struct for BlockTxVerboseDataModel
type BlockTxVerboseDataModel struct {
	TransactionId string `json:"transactionId"`
	Hash *string `json:"hash,omitempty"`
	ComputeMass *int32 `json:"computeMass,omitempty"`
	BlockHash *string `json:"blockHash,omitempty"`
	BlockTime *int32 `json:"blockTime,omitempty"`
}

type _BlockTxVerboseDataModel BlockTxVerboseDataModel

// NewBlockTxVerboseDataModel instantiates a new BlockTxVerboseDataModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockTxVerboseDataModel(transactionId string) *BlockTxVerboseDataModel {
	this := BlockTxVerboseDataModel{}
	this.TransactionId = transactionId
	return &this
}

// NewBlockTxVerboseDataModelWithDefaults instantiates a new BlockTxVerboseDataModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockTxVerboseDataModelWithDefaults() *BlockTxVerboseDataModel {
	this := BlockTxVerboseDataModel{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *BlockTxVerboseDataModel) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *BlockTxVerboseDataModel) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *BlockTxVerboseDataModel) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *BlockTxVerboseDataModel) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxVerboseDataModel) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *BlockTxVerboseDataModel) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *BlockTxVerboseDataModel) SetHash(v string) {
	o.Hash = &v
}

// GetComputeMass returns the ComputeMass field value if set, zero value otherwise.
func (o *BlockTxVerboseDataModel) GetComputeMass() int32 {
	if o == nil || IsNil(o.ComputeMass) {
		var ret int32
		return ret
	}
	return *o.ComputeMass
}

// GetComputeMassOk returns a tuple with the ComputeMass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxVerboseDataModel) GetComputeMassOk() (*int32, bool) {
	if o == nil || IsNil(o.ComputeMass) {
		return nil, false
	}
	return o.ComputeMass, true
}

// HasComputeMass returns a boolean if a field has been set.
func (o *BlockTxVerboseDataModel) HasComputeMass() bool {
	if o != nil && !IsNil(o.ComputeMass) {
		return true
	}

	return false
}

// SetComputeMass gets a reference to the given int32 and assigns it to the ComputeMass field.
func (o *BlockTxVerboseDataModel) SetComputeMass(v int32) {
	o.ComputeMass = &v
}

// GetBlockHash returns the BlockHash field value if set, zero value otherwise.
func (o *BlockTxVerboseDataModel) GetBlockHash() string {
	if o == nil || IsNil(o.BlockHash) {
		var ret string
		return ret
	}
	return *o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxVerboseDataModel) GetBlockHashOk() (*string, bool) {
	if o == nil || IsNil(o.BlockHash) {
		return nil, false
	}
	return o.BlockHash, true
}

// HasBlockHash returns a boolean if a field has been set.
func (o *BlockTxVerboseDataModel) HasBlockHash() bool {
	if o != nil && !IsNil(o.BlockHash) {
		return true
	}

	return false
}

// SetBlockHash gets a reference to the given string and assigns it to the BlockHash field.
func (o *BlockTxVerboseDataModel) SetBlockHash(v string) {
	o.BlockHash = &v
}

// GetBlockTime returns the BlockTime field value if set, zero value otherwise.
func (o *BlockTxVerboseDataModel) GetBlockTime() int32 {
	if o == nil || IsNil(o.BlockTime) {
		var ret int32
		return ret
	}
	return *o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxVerboseDataModel) GetBlockTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.BlockTime) {
		return nil, false
	}
	return o.BlockTime, true
}

// HasBlockTime returns a boolean if a field has been set.
func (o *BlockTxVerboseDataModel) HasBlockTime() bool {
	if o != nil && !IsNil(o.BlockTime) {
		return true
	}

	return false
}

// SetBlockTime gets a reference to the given int32 and assigns it to the BlockTime field.
func (o *BlockTxVerboseDataModel) SetBlockTime(v int32) {
	o.BlockTime = &v
}

func (o BlockTxVerboseDataModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockTxVerboseDataModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactionId"] = o.TransactionId
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.ComputeMass) {
		toSerialize["computeMass"] = o.ComputeMass
	}
	if !IsNil(o.BlockHash) {
		toSerialize["blockHash"] = o.BlockHash
	}
	if !IsNil(o.BlockTime) {
		toSerialize["blockTime"] = o.BlockTime
	}
	return toSerialize, nil
}

func (o *BlockTxVerboseDataModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transactionId",
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

	varBlockTxVerboseDataModel := _BlockTxVerboseDataModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockTxVerboseDataModel)

	if err != nil {
		return err
	}

	*o = BlockTxVerboseDataModel(varBlockTxVerboseDataModel)

	return err
}

type NullableBlockTxVerboseDataModel struct {
	value *BlockTxVerboseDataModel
	isSet bool
}

func (v NullableBlockTxVerboseDataModel) Get() *BlockTxVerboseDataModel {
	return v.value
}

func (v *NullableBlockTxVerboseDataModel) Set(val *BlockTxVerboseDataModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockTxVerboseDataModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockTxVerboseDataModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockTxVerboseDataModel(val *BlockTxVerboseDataModel) *NullableBlockTxVerboseDataModel {
	return &NullableBlockTxVerboseDataModel{value: val, isSet: true}
}

func (v NullableBlockTxVerboseDataModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockTxVerboseDataModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


