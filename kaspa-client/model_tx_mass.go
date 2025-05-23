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

// checks if the TxMass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TxMass{}

// TxMass struct for TxMass
type TxMass struct {
	Mass int32 `json:"mass"`
	StorageMass int32 `json:"storage_mass"`
	ComputeMass int32 `json:"compute_mass"`
}

type _TxMass TxMass

// NewTxMass instantiates a new TxMass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxMass(mass int32, storageMass int32, computeMass int32) *TxMass {
	this := TxMass{}
	this.Mass = mass
	this.StorageMass = storageMass
	this.ComputeMass = computeMass
	return &this
}

// NewTxMassWithDefaults instantiates a new TxMass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxMassWithDefaults() *TxMass {
	this := TxMass{}
	return &this
}

// GetMass returns the Mass field value
func (o *TxMass) GetMass() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Mass
}

// GetMassOk returns a tuple with the Mass field value
// and a boolean to check if the value has been set.
func (o *TxMass) GetMassOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mass, true
}

// SetMass sets field value
func (o *TxMass) SetMass(v int32) {
	o.Mass = v
}

// GetStorageMass returns the StorageMass field value
func (o *TxMass) GetStorageMass() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StorageMass
}

// GetStorageMassOk returns a tuple with the StorageMass field value
// and a boolean to check if the value has been set.
func (o *TxMass) GetStorageMassOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageMass, true
}

// SetStorageMass sets field value
func (o *TxMass) SetStorageMass(v int32) {
	o.StorageMass = v
}

// GetComputeMass returns the ComputeMass field value
func (o *TxMass) GetComputeMass() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ComputeMass
}

// GetComputeMassOk returns a tuple with the ComputeMass field value
// and a boolean to check if the value has been set.
func (o *TxMass) GetComputeMassOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputeMass, true
}

// SetComputeMass sets field value
func (o *TxMass) SetComputeMass(v int32) {
	o.ComputeMass = v
}

func (o TxMass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TxMass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mass"] = o.Mass
	toSerialize["storage_mass"] = o.StorageMass
	toSerialize["compute_mass"] = o.ComputeMass
	return toSerialize, nil
}

func (o *TxMass) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mass",
		"storage_mass",
		"compute_mass",
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

	varTxMass := _TxMass{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTxMass)

	if err != nil {
		return err
	}

	*o = TxMass(varTxMass)

	return err
}

type NullableTxMass struct {
	value *TxMass
	isSet bool
}

func (v NullableTxMass) Get() *TxMass {
	return v.value
}

func (v *NullableTxMass) Set(val *TxMass) {
	v.value = val
	v.isSet = true
}

func (v NullableTxMass) IsSet() bool {
	return v.isSet
}

func (v *NullableTxMass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxMass(val *TxMass) *NullableTxMass {
	return &NullableTxMass{value: val, isSet: true}
}

func (v NullableTxMass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxMass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


