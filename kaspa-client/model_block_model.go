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

// checks if the BlockModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockModel{}

// BlockModel struct for BlockModel
type BlockModel struct {
	Header EndpointsGetBlocksBlockHeader `json:"header"`
	Transactions []BlockTxModel `json:"transactions,omitempty"`
	VerboseData VerboseDataModel `json:"verboseData"`
	Extra *ExtraModel `json:"extra,omitempty"`
}

type _BlockModel BlockModel

// NewBlockModel instantiates a new BlockModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockModel(header EndpointsGetBlocksBlockHeader, verboseData VerboseDataModel) *BlockModel {
	this := BlockModel{}
	this.Header = header
	this.VerboseData = verboseData
	return &this
}

// NewBlockModelWithDefaults instantiates a new BlockModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockModelWithDefaults() *BlockModel {
	this := BlockModel{}
	return &this
}

// GetHeader returns the Header field value
func (o *BlockModel) GetHeader() EndpointsGetBlocksBlockHeader {
	if o == nil {
		var ret EndpointsGetBlocksBlockHeader
		return ret
	}

	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value
// and a boolean to check if the value has been set.
func (o *BlockModel) GetHeaderOk() (*EndpointsGetBlocksBlockHeader, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Header, true
}

// SetHeader sets field value
func (o *BlockModel) SetHeader(v EndpointsGetBlocksBlockHeader) {
	o.Header = v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *BlockModel) GetTransactions() []BlockTxModel {
	if o == nil || IsNil(o.Transactions) {
		var ret []BlockTxModel
		return ret
	}
	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockModel) GetTransactionsOk() ([]BlockTxModel, bool) {
	if o == nil || IsNil(o.Transactions) {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *BlockModel) HasTransactions() bool {
	if o != nil && !IsNil(o.Transactions) {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []BlockTxModel and assigns it to the Transactions field.
func (o *BlockModel) SetTransactions(v []BlockTxModel) {
	o.Transactions = v
}

// GetVerboseData returns the VerboseData field value
func (o *BlockModel) GetVerboseData() VerboseDataModel {
	if o == nil {
		var ret VerboseDataModel
		return ret
	}

	return o.VerboseData
}

// GetVerboseDataOk returns a tuple with the VerboseData field value
// and a boolean to check if the value has been set.
func (o *BlockModel) GetVerboseDataOk() (*VerboseDataModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseData, true
}

// SetVerboseData sets field value
func (o *BlockModel) SetVerboseData(v VerboseDataModel) {
	o.VerboseData = v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *BlockModel) GetExtra() ExtraModel {
	if o == nil || IsNil(o.Extra) {
		var ret ExtraModel
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockModel) GetExtraOk() (*ExtraModel, bool) {
	if o == nil || IsNil(o.Extra) {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *BlockModel) HasExtra() bool {
	if o != nil && !IsNil(o.Extra) {
		return true
	}

	return false
}

// SetExtra gets a reference to the given ExtraModel and assigns it to the Extra field.
func (o *BlockModel) SetExtra(v ExtraModel) {
	o.Extra = &v
}

func (o BlockModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["header"] = o.Header
	if !IsNil(o.Transactions) {
		toSerialize["transactions"] = o.Transactions
	}
	toSerialize["verboseData"] = o.VerboseData
	if !IsNil(o.Extra) {
		toSerialize["extra"] = o.Extra
	}
	return toSerialize, nil
}

func (o *BlockModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"header",
		"verboseData",
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

	varBlockModel := _BlockModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockModel)

	if err != nil {
		return err
	}

	*o = BlockModel(varBlockModel)

	return err
}

type NullableBlockModel struct {
	value *BlockModel
	isSet bool
}

func (v NullableBlockModel) Get() *BlockModel {
	return v.value
}

func (v *NullableBlockModel) Set(val *BlockModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockModel(val *BlockModel) *NullableBlockModel {
	return &NullableBlockModel{value: val, isSet: true}
}

func (v NullableBlockModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


