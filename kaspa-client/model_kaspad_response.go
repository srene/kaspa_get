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

// checks if the KaspadResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KaspadResponse{}

// KaspadResponse struct for KaspadResponse
type KaspadResponse struct {
	KaspadHost *string `json:"kaspadHost,omitempty"`
	ServerVersion *string `json:"serverVersion,omitempty"`
	IsUtxoIndexed *bool `json:"isUtxoIndexed,omitempty"`
	IsSynced *bool `json:"isSynced,omitempty"`
	P2pId *string `json:"p2pId,omitempty"`
	BlueScore *int32 `json:"blueScore,omitempty"`
}

// NewKaspadResponse instantiates a new KaspadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKaspadResponse() *KaspadResponse {
	this := KaspadResponse{}
	var kaspadHost string = ""
	this.KaspadHost = &kaspadHost
	var serverVersion string = "0.12.6"
	this.ServerVersion = &serverVersion
	var isUtxoIndexed bool = true
	this.IsUtxoIndexed = &isUtxoIndexed
	var isSynced bool = true
	this.IsSynced = &isSynced
	var p2pId string = "1231312"
	this.P2pId = &p2pId
	var blueScore int32 = 101065625
	this.BlueScore = &blueScore
	return &this
}

// NewKaspadResponseWithDefaults instantiates a new KaspadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKaspadResponseWithDefaults() *KaspadResponse {
	this := KaspadResponse{}
	var kaspadHost string = ""
	this.KaspadHost = &kaspadHost
	var serverVersion string = "0.12.6"
	this.ServerVersion = &serverVersion
	var isUtxoIndexed bool = true
	this.IsUtxoIndexed = &isUtxoIndexed
	var isSynced bool = true
	this.IsSynced = &isSynced
	var p2pId string = "1231312"
	this.P2pId = &p2pId
	var blueScore int32 = 101065625
	this.BlueScore = &blueScore
	return &this
}

// GetKaspadHost returns the KaspadHost field value if set, zero value otherwise.
func (o *KaspadResponse) GetKaspadHost() string {
	if o == nil || IsNil(o.KaspadHost) {
		var ret string
		return ret
	}
	return *o.KaspadHost
}

// GetKaspadHostOk returns a tuple with the KaspadHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KaspadResponse) GetKaspadHostOk() (*string, bool) {
	if o == nil || IsNil(o.KaspadHost) {
		return nil, false
	}
	return o.KaspadHost, true
}

// HasKaspadHost returns a boolean if a field has been set.
func (o *KaspadResponse) HasKaspadHost() bool {
	if o != nil && !IsNil(o.KaspadHost) {
		return true
	}

	return false
}

// SetKaspadHost gets a reference to the given string and assigns it to the KaspadHost field.
func (o *KaspadResponse) SetKaspadHost(v string) {
	o.KaspadHost = &v
}

// GetServerVersion returns the ServerVersion field value if set, zero value otherwise.
func (o *KaspadResponse) GetServerVersion() string {
	if o == nil || IsNil(o.ServerVersion) {
		var ret string
		return ret
	}
	return *o.ServerVersion
}

// GetServerVersionOk returns a tuple with the ServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KaspadResponse) GetServerVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ServerVersion) {
		return nil, false
	}
	return o.ServerVersion, true
}

// HasServerVersion returns a boolean if a field has been set.
func (o *KaspadResponse) HasServerVersion() bool {
	if o != nil && !IsNil(o.ServerVersion) {
		return true
	}

	return false
}

// SetServerVersion gets a reference to the given string and assigns it to the ServerVersion field.
func (o *KaspadResponse) SetServerVersion(v string) {
	o.ServerVersion = &v
}

// GetIsUtxoIndexed returns the IsUtxoIndexed field value if set, zero value otherwise.
func (o *KaspadResponse) GetIsUtxoIndexed() bool {
	if o == nil || IsNil(o.IsUtxoIndexed) {
		var ret bool
		return ret
	}
	return *o.IsUtxoIndexed
}

// GetIsUtxoIndexedOk returns a tuple with the IsUtxoIndexed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KaspadResponse) GetIsUtxoIndexedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUtxoIndexed) {
		return nil, false
	}
	return o.IsUtxoIndexed, true
}

// HasIsUtxoIndexed returns a boolean if a field has been set.
func (o *KaspadResponse) HasIsUtxoIndexed() bool {
	if o != nil && !IsNil(o.IsUtxoIndexed) {
		return true
	}

	return false
}

// SetIsUtxoIndexed gets a reference to the given bool and assigns it to the IsUtxoIndexed field.
func (o *KaspadResponse) SetIsUtxoIndexed(v bool) {
	o.IsUtxoIndexed = &v
}

// GetIsSynced returns the IsSynced field value if set, zero value otherwise.
func (o *KaspadResponse) GetIsSynced() bool {
	if o == nil || IsNil(o.IsSynced) {
		var ret bool
		return ret
	}
	return *o.IsSynced
}

// GetIsSyncedOk returns a tuple with the IsSynced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KaspadResponse) GetIsSyncedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSynced) {
		return nil, false
	}
	return o.IsSynced, true
}

// HasIsSynced returns a boolean if a field has been set.
func (o *KaspadResponse) HasIsSynced() bool {
	if o != nil && !IsNil(o.IsSynced) {
		return true
	}

	return false
}

// SetIsSynced gets a reference to the given bool and assigns it to the IsSynced field.
func (o *KaspadResponse) SetIsSynced(v bool) {
	o.IsSynced = &v
}

// GetP2pId returns the P2pId field value if set, zero value otherwise.
func (o *KaspadResponse) GetP2pId() string {
	if o == nil || IsNil(o.P2pId) {
		var ret string
		return ret
	}
	return *o.P2pId
}

// GetP2pIdOk returns a tuple with the P2pId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KaspadResponse) GetP2pIdOk() (*string, bool) {
	if o == nil || IsNil(o.P2pId) {
		return nil, false
	}
	return o.P2pId, true
}

// HasP2pId returns a boolean if a field has been set.
func (o *KaspadResponse) HasP2pId() bool {
	if o != nil && !IsNil(o.P2pId) {
		return true
	}

	return false
}

// SetP2pId gets a reference to the given string and assigns it to the P2pId field.
func (o *KaspadResponse) SetP2pId(v string) {
	o.P2pId = &v
}

// GetBlueScore returns the BlueScore field value if set, zero value otherwise.
func (o *KaspadResponse) GetBlueScore() int32 {
	if o == nil || IsNil(o.BlueScore) {
		var ret int32
		return ret
	}
	return *o.BlueScore
}

// GetBlueScoreOk returns a tuple with the BlueScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KaspadResponse) GetBlueScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.BlueScore) {
		return nil, false
	}
	return o.BlueScore, true
}

// HasBlueScore returns a boolean if a field has been set.
func (o *KaspadResponse) HasBlueScore() bool {
	if o != nil && !IsNil(o.BlueScore) {
		return true
	}

	return false
}

// SetBlueScore gets a reference to the given int32 and assigns it to the BlueScore field.
func (o *KaspadResponse) SetBlueScore(v int32) {
	o.BlueScore = &v
}

func (o KaspadResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KaspadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KaspadHost) {
		toSerialize["kaspadHost"] = o.KaspadHost
	}
	if !IsNil(o.ServerVersion) {
		toSerialize["serverVersion"] = o.ServerVersion
	}
	if !IsNil(o.IsUtxoIndexed) {
		toSerialize["isUtxoIndexed"] = o.IsUtxoIndexed
	}
	if !IsNil(o.IsSynced) {
		toSerialize["isSynced"] = o.IsSynced
	}
	if !IsNil(o.P2pId) {
		toSerialize["p2pId"] = o.P2pId
	}
	if !IsNil(o.BlueScore) {
		toSerialize["blueScore"] = o.BlueScore
	}
	return toSerialize, nil
}

type NullableKaspadResponse struct {
	value *KaspadResponse
	isSet bool
}

func (v NullableKaspadResponse) Get() *KaspadResponse {
	return v.value
}

func (v *NullableKaspadResponse) Set(val *KaspadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKaspadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKaspadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKaspadResponse(val *KaspadResponse) *NullableKaspadResponse {
	return &NullableKaspadResponse{value: val, isSet: true}
}

func (v NullableKaspadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKaspadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


