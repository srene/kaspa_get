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

// checks if the BlockTxModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockTxModel{}

// BlockTxModel struct for BlockTxModel
type BlockTxModel struct {
	Inputs []BlockTxInputModel `json:"inputs,omitempty"`
	Outputs []BlockTxOutputModel `json:"outputs,omitempty"`
	SubnetworkId *string `json:"subnetworkId,omitempty"`
	Payload *string `json:"payload,omitempty"`
	VerboseData BlockTxVerboseDataModel `json:"verboseData"`
	LockTime *int32 `json:"lockTime,omitempty"`
	Gas *int32 `json:"gas,omitempty"`
	Mass *int32 `json:"mass,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

type _BlockTxModel BlockTxModel

// NewBlockTxModel instantiates a new BlockTxModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockTxModel(verboseData BlockTxVerboseDataModel) *BlockTxModel {
	this := BlockTxModel{}
	this.VerboseData = verboseData
	return &this
}

// NewBlockTxModelWithDefaults instantiates a new BlockTxModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockTxModelWithDefaults() *BlockTxModel {
	this := BlockTxModel{}
	return &this
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *BlockTxModel) GetInputs() []BlockTxInputModel {
	if o == nil || IsNil(o.Inputs) {
		var ret []BlockTxInputModel
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxModel) GetInputsOk() ([]BlockTxInputModel, bool) {
	if o == nil || IsNil(o.Inputs) {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *BlockTxModel) HasInputs() bool {
	if o != nil && !IsNil(o.Inputs) {
		return true
	}

	return false
}

// SetInputs gets a reference to the given []BlockTxInputModel and assigns it to the Inputs field.
func (o *BlockTxModel) SetInputs(v []BlockTxInputModel) {
	o.Inputs = v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *BlockTxModel) GetOutputs() []BlockTxOutputModel {
	if o == nil || IsNil(o.Outputs) {
		var ret []BlockTxOutputModel
		return ret
	}
	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxModel) GetOutputsOk() ([]BlockTxOutputModel, bool) {
	if o == nil || IsNil(o.Outputs) {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *BlockTxModel) HasOutputs() bool {
	if o != nil && !IsNil(o.Outputs) {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given []BlockTxOutputModel and assigns it to the Outputs field.
func (o *BlockTxModel) SetOutputs(v []BlockTxOutputModel) {
	o.Outputs = v
}

// GetSubnetworkId returns the SubnetworkId field value if set, zero value otherwise.
func (o *BlockTxModel) GetSubnetworkId() string {
	if o == nil || IsNil(o.SubnetworkId) {
		var ret string
		return ret
	}
	return *o.SubnetworkId
}

// GetSubnetworkIdOk returns a tuple with the SubnetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxModel) GetSubnetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubnetworkId) {
		return nil, false
	}
	return o.SubnetworkId, true
}

// HasSubnetworkId returns a boolean if a field has been set.
func (o *BlockTxModel) HasSubnetworkId() bool {
	if o != nil && !IsNil(o.SubnetworkId) {
		return true
	}

	return false
}

// SetSubnetworkId gets a reference to the given string and assigns it to the SubnetworkId field.
func (o *BlockTxModel) SetSubnetworkId(v string) {
	o.SubnetworkId = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *BlockTxModel) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxModel) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *BlockTxModel) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *BlockTxModel) SetPayload(v string) {
	o.Payload = &v
}

// GetVerboseData returns the VerboseData field value
func (o *BlockTxModel) GetVerboseData() BlockTxVerboseDataModel {
	if o == nil {
		var ret BlockTxVerboseDataModel
		return ret
	}

	return o.VerboseData
}

// GetVerboseDataOk returns a tuple with the VerboseData field value
// and a boolean to check if the value has been set.
func (o *BlockTxModel) GetVerboseDataOk() (*BlockTxVerboseDataModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseData, true
}

// SetVerboseData sets field value
func (o *BlockTxModel) SetVerboseData(v BlockTxVerboseDataModel) {
	o.VerboseData = v
}

// GetLockTime returns the LockTime field value if set, zero value otherwise.
func (o *BlockTxModel) GetLockTime() int32 {
	if o == nil || IsNil(o.LockTime) {
		var ret int32
		return ret
	}
	return *o.LockTime
}

// GetLockTimeOk returns a tuple with the LockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxModel) GetLockTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.LockTime) {
		return nil, false
	}
	return o.LockTime, true
}

// HasLockTime returns a boolean if a field has been set.
func (o *BlockTxModel) HasLockTime() bool {
	if o != nil && !IsNil(o.LockTime) {
		return true
	}

	return false
}

// SetLockTime gets a reference to the given int32 and assigns it to the LockTime field.
func (o *BlockTxModel) SetLockTime(v int32) {
	o.LockTime = &v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *BlockTxModel) GetGas() int32 {
	if o == nil || IsNil(o.Gas) {
		var ret int32
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxModel) GetGasOk() (*int32, bool) {
	if o == nil || IsNil(o.Gas) {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *BlockTxModel) HasGas() bool {
	if o != nil && !IsNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given int32 and assigns it to the Gas field.
func (o *BlockTxModel) SetGas(v int32) {
	o.Gas = &v
}

// GetMass returns the Mass field value if set, zero value otherwise.
func (o *BlockTxModel) GetMass() int32 {
	if o == nil || IsNil(o.Mass) {
		var ret int32
		return ret
	}
	return *o.Mass
}

// GetMassOk returns a tuple with the Mass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxModel) GetMassOk() (*int32, bool) {
	if o == nil || IsNil(o.Mass) {
		return nil, false
	}
	return o.Mass, true
}

// HasMass returns a boolean if a field has been set.
func (o *BlockTxModel) HasMass() bool {
	if o != nil && !IsNil(o.Mass) {
		return true
	}

	return false
}

// SetMass gets a reference to the given int32 and assigns it to the Mass field.
func (o *BlockTxModel) SetMass(v int32) {
	o.Mass = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BlockTxModel) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxModel) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BlockTxModel) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *BlockTxModel) SetVersion(v int32) {
	o.Version = &v
}

func (o BlockTxModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockTxModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Inputs) {
		toSerialize["inputs"] = o.Inputs
	}
	if !IsNil(o.Outputs) {
		toSerialize["outputs"] = o.Outputs
	}
	if !IsNil(o.SubnetworkId) {
		toSerialize["subnetworkId"] = o.SubnetworkId
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	toSerialize["verboseData"] = o.VerboseData
	if !IsNil(o.LockTime) {
		toSerialize["lockTime"] = o.LockTime
	}
	if !IsNil(o.Gas) {
		toSerialize["gas"] = o.Gas
	}
	if !IsNil(o.Mass) {
		toSerialize["mass"] = o.Mass
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

func (o *BlockTxModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varBlockTxModel := _BlockTxModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockTxModel)

	if err != nil {
		return err
	}

	*o = BlockTxModel(varBlockTxModel)

	return err
}

type NullableBlockTxModel struct {
	value *BlockTxModel
	isSet bool
}

func (v NullableBlockTxModel) Get() *BlockTxModel {
	return v.value
}

func (v *NullableBlockTxModel) Set(val *BlockTxModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockTxModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockTxModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockTxModel(val *BlockTxModel) *NullableBlockTxModel {
	return &NullableBlockTxModel{value: val, isSet: true}
}

func (v NullableBlockTxModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockTxModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


