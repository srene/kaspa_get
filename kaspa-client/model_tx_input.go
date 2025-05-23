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

// checks if the TxInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TxInput{}

// TxInput struct for TxInput
type TxInput struct {
	TransactionId string `json:"transaction_id"`
	Index int32 `json:"index"`
	PreviousOutpointHash string `json:"previous_outpoint_hash"`
	PreviousOutpointIndex string `json:"previous_outpoint_index"`
	PreviousOutpointResolved *TxOutput `json:"previous_outpoint_resolved,omitempty"`
	PreviousOutpointAddress *string `json:"previous_outpoint_address,omitempty"`
	PreviousOutpointAmount *int32 `json:"previous_outpoint_amount,omitempty"`
	SignatureScript *string `json:"signature_script,omitempty"`
	SigOpCount *int32 `json:"sig_op_count,omitempty"`
}

type _TxInput TxInput

// NewTxInput instantiates a new TxInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxInput(transactionId string, index int32, previousOutpointHash string, previousOutpointIndex string) *TxInput {
	this := TxInput{}
	this.TransactionId = transactionId
	this.Index = index
	this.PreviousOutpointHash = previousOutpointHash
	this.PreviousOutpointIndex = previousOutpointIndex
	return &this
}

// NewTxInputWithDefaults instantiates a new TxInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxInputWithDefaults() *TxInput {
	this := TxInput{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *TxInput) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *TxInput) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *TxInput) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetIndex returns the Index field value
func (o *TxInput) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *TxInput) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *TxInput) SetIndex(v int32) {
	o.Index = v
}

// GetPreviousOutpointHash returns the PreviousOutpointHash field value
func (o *TxInput) GetPreviousOutpointHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousOutpointHash
}

// GetPreviousOutpointHashOk returns a tuple with the PreviousOutpointHash field value
// and a boolean to check if the value has been set.
func (o *TxInput) GetPreviousOutpointHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousOutpointHash, true
}

// SetPreviousOutpointHash sets field value
func (o *TxInput) SetPreviousOutpointHash(v string) {
	o.PreviousOutpointHash = v
}

// GetPreviousOutpointIndex returns the PreviousOutpointIndex field value
func (o *TxInput) GetPreviousOutpointIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousOutpointIndex
}

// GetPreviousOutpointIndexOk returns a tuple with the PreviousOutpointIndex field value
// and a boolean to check if the value has been set.
func (o *TxInput) GetPreviousOutpointIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousOutpointIndex, true
}

// SetPreviousOutpointIndex sets field value
func (o *TxInput) SetPreviousOutpointIndex(v string) {
	o.PreviousOutpointIndex = v
}

// GetPreviousOutpointResolved returns the PreviousOutpointResolved field value if set, zero value otherwise.
func (o *TxInput) GetPreviousOutpointResolved() TxOutput {
	if o == nil || IsNil(o.PreviousOutpointResolved) {
		var ret TxOutput
		return ret
	}
	return *o.PreviousOutpointResolved
}

// GetPreviousOutpointResolvedOk returns a tuple with the PreviousOutpointResolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxInput) GetPreviousOutpointResolvedOk() (*TxOutput, bool) {
	if o == nil || IsNil(o.PreviousOutpointResolved) {
		return nil, false
	}
	return o.PreviousOutpointResolved, true
}

// HasPreviousOutpointResolved returns a boolean if a field has been set.
func (o *TxInput) HasPreviousOutpointResolved() bool {
	if o != nil && !IsNil(o.PreviousOutpointResolved) {
		return true
	}

	return false
}

// SetPreviousOutpointResolved gets a reference to the given TxOutput and assigns it to the PreviousOutpointResolved field.
func (o *TxInput) SetPreviousOutpointResolved(v TxOutput) {
	o.PreviousOutpointResolved = &v
}

// GetPreviousOutpointAddress returns the PreviousOutpointAddress field value if set, zero value otherwise.
func (o *TxInput) GetPreviousOutpointAddress() string {
	if o == nil || IsNil(o.PreviousOutpointAddress) {
		var ret string
		return ret
	}
	return *o.PreviousOutpointAddress
}

// GetPreviousOutpointAddressOk returns a tuple with the PreviousOutpointAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxInput) GetPreviousOutpointAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousOutpointAddress) {
		return nil, false
	}
	return o.PreviousOutpointAddress, true
}

// HasPreviousOutpointAddress returns a boolean if a field has been set.
func (o *TxInput) HasPreviousOutpointAddress() bool {
	if o != nil && !IsNil(o.PreviousOutpointAddress) {
		return true
	}

	return false
}

// SetPreviousOutpointAddress gets a reference to the given string and assigns it to the PreviousOutpointAddress field.
func (o *TxInput) SetPreviousOutpointAddress(v string) {
	o.PreviousOutpointAddress = &v
}

// GetPreviousOutpointAmount returns the PreviousOutpointAmount field value if set, zero value otherwise.
func (o *TxInput) GetPreviousOutpointAmount() int32 {
	if o == nil || IsNil(o.PreviousOutpointAmount) {
		var ret int32
		return ret
	}
	return *o.PreviousOutpointAmount
}

// GetPreviousOutpointAmountOk returns a tuple with the PreviousOutpointAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxInput) GetPreviousOutpointAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.PreviousOutpointAmount) {
		return nil, false
	}
	return o.PreviousOutpointAmount, true
}

// HasPreviousOutpointAmount returns a boolean if a field has been set.
func (o *TxInput) HasPreviousOutpointAmount() bool {
	if o != nil && !IsNil(o.PreviousOutpointAmount) {
		return true
	}

	return false
}

// SetPreviousOutpointAmount gets a reference to the given int32 and assigns it to the PreviousOutpointAmount field.
func (o *TxInput) SetPreviousOutpointAmount(v int32) {
	o.PreviousOutpointAmount = &v
}

// GetSignatureScript returns the SignatureScript field value if set, zero value otherwise.
func (o *TxInput) GetSignatureScript() string {
	if o == nil || IsNil(o.SignatureScript) {
		var ret string
		return ret
	}
	return *o.SignatureScript
}

// GetSignatureScriptOk returns a tuple with the SignatureScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxInput) GetSignatureScriptOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureScript) {
		return nil, false
	}
	return o.SignatureScript, true
}

// HasSignatureScript returns a boolean if a field has been set.
func (o *TxInput) HasSignatureScript() bool {
	if o != nil && !IsNil(o.SignatureScript) {
		return true
	}

	return false
}

// SetSignatureScript gets a reference to the given string and assigns it to the SignatureScript field.
func (o *TxInput) SetSignatureScript(v string) {
	o.SignatureScript = &v
}

// GetSigOpCount returns the SigOpCount field value if set, zero value otherwise.
func (o *TxInput) GetSigOpCount() int32 {
	if o == nil || IsNil(o.SigOpCount) {
		var ret int32
		return ret
	}
	return *o.SigOpCount
}

// GetSigOpCountOk returns a tuple with the SigOpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxInput) GetSigOpCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SigOpCount) {
		return nil, false
	}
	return o.SigOpCount, true
}

// HasSigOpCount returns a boolean if a field has been set.
func (o *TxInput) HasSigOpCount() bool {
	if o != nil && !IsNil(o.SigOpCount) {
		return true
	}

	return false
}

// SetSigOpCount gets a reference to the given int32 and assigns it to the SigOpCount field.
func (o *TxInput) SetSigOpCount(v int32) {
	o.SigOpCount = &v
}

func (o TxInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TxInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction_id"] = o.TransactionId
	toSerialize["index"] = o.Index
	toSerialize["previous_outpoint_hash"] = o.PreviousOutpointHash
	toSerialize["previous_outpoint_index"] = o.PreviousOutpointIndex
	if !IsNil(o.PreviousOutpointResolved) {
		toSerialize["previous_outpoint_resolved"] = o.PreviousOutpointResolved
	}
	if !IsNil(o.PreviousOutpointAddress) {
		toSerialize["previous_outpoint_address"] = o.PreviousOutpointAddress
	}
	if !IsNil(o.PreviousOutpointAmount) {
		toSerialize["previous_outpoint_amount"] = o.PreviousOutpointAmount
	}
	if !IsNil(o.SignatureScript) {
		toSerialize["signature_script"] = o.SignatureScript
	}
	if !IsNil(o.SigOpCount) {
		toSerialize["sig_op_count"] = o.SigOpCount
	}
	return toSerialize, nil
}

func (o *TxInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction_id",
		"index",
		"previous_outpoint_hash",
		"previous_outpoint_index",
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

	varTxInput := _TxInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTxInput)

	if err != nil {
		return err
	}

	*o = TxInput(varTxInput)

	return err
}

type NullableTxInput struct {
	value *TxInput
	isSet bool
}

func (v NullableTxInput) Get() *TxInput {
	return v.value
}

func (v *NullableTxInput) Set(val *TxInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTxInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTxInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxInput(val *TxInput) *NullableTxInput {
	return &NullableTxInput{value: val, isSet: true}
}

func (v NullableTxInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


