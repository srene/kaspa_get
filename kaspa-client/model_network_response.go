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

// checks if the NetworkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkResponse{}

// NetworkResponse struct for NetworkResponse
type NetworkResponse struct {
	NetworkName *string `json:"networkName,omitempty"`
	BlockCount *string `json:"blockCount,omitempty"`
	HeaderCount *string `json:"headerCount,omitempty"`
	TipHashes []string `json:"tipHashes,omitempty"`
	Difficulty *float32 `json:"difficulty,omitempty"`
	PastMedianTime *string `json:"pastMedianTime,omitempty"`
	VirtualParentHashes []string `json:"virtualParentHashes,omitempty"`
	PruningPointHash *string `json:"pruningPointHash,omitempty"`
	VirtualDaaScore *string `json:"virtualDaaScore,omitempty"`
}

// NewNetworkResponse instantiates a new NetworkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkResponse() *NetworkResponse {
	this := NetworkResponse{}
	var networkName string = "kaspa-mainnet"
	this.NetworkName = &networkName
	var blockCount string = "261357"
	this.BlockCount = &blockCount
	var headerCount string = "23138783"
	this.HeaderCount = &headerCount
	var difficulty float32 = 3.88707990501409E12
	this.Difficulty = &difficulty
	var pastMedianTime string = "1656456088196"
	this.PastMedianTime = &pastMedianTime
	var pruningPointHash string = "5d32a9403273a34b6551b84340a1459ddde2ae6ba59a47987a6374340ba41d5d"
	this.PruningPointHash = &pruningPointHash
	var virtualDaaScore string = "19989984"
	this.VirtualDaaScore = &virtualDaaScore
	return &this
}

// NewNetworkResponseWithDefaults instantiates a new NetworkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkResponseWithDefaults() *NetworkResponse {
	this := NetworkResponse{}
	var networkName string = "kaspa-mainnet"
	this.NetworkName = &networkName
	var blockCount string = "261357"
	this.BlockCount = &blockCount
	var headerCount string = "23138783"
	this.HeaderCount = &headerCount
	var difficulty float32 = 3.88707990501409E12
	this.Difficulty = &difficulty
	var pastMedianTime string = "1656456088196"
	this.PastMedianTime = &pastMedianTime
	var pruningPointHash string = "5d32a9403273a34b6551b84340a1459ddde2ae6ba59a47987a6374340ba41d5d"
	this.PruningPointHash = &pruningPointHash
	var virtualDaaScore string = "19989984"
	this.VirtualDaaScore = &virtualDaaScore
	return &this
}

// GetNetworkName returns the NetworkName field value if set, zero value otherwise.
func (o *NetworkResponse) GetNetworkName() string {
	if o == nil || IsNil(o.NetworkName) {
		var ret string
		return ret
	}
	return *o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkResponse) GetNetworkNameOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkName) {
		return nil, false
	}
	return o.NetworkName, true
}

// HasNetworkName returns a boolean if a field has been set.
func (o *NetworkResponse) HasNetworkName() bool {
	if o != nil && !IsNil(o.NetworkName) {
		return true
	}

	return false
}

// SetNetworkName gets a reference to the given string and assigns it to the NetworkName field.
func (o *NetworkResponse) SetNetworkName(v string) {
	o.NetworkName = &v
}

// GetBlockCount returns the BlockCount field value if set, zero value otherwise.
func (o *NetworkResponse) GetBlockCount() string {
	if o == nil || IsNil(o.BlockCount) {
		var ret string
		return ret
	}
	return *o.BlockCount
}

// GetBlockCountOk returns a tuple with the BlockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkResponse) GetBlockCountOk() (*string, bool) {
	if o == nil || IsNil(o.BlockCount) {
		return nil, false
	}
	return o.BlockCount, true
}

// HasBlockCount returns a boolean if a field has been set.
func (o *NetworkResponse) HasBlockCount() bool {
	if o != nil && !IsNil(o.BlockCount) {
		return true
	}

	return false
}

// SetBlockCount gets a reference to the given string and assigns it to the BlockCount field.
func (o *NetworkResponse) SetBlockCount(v string) {
	o.BlockCount = &v
}

// GetHeaderCount returns the HeaderCount field value if set, zero value otherwise.
func (o *NetworkResponse) GetHeaderCount() string {
	if o == nil || IsNil(o.HeaderCount) {
		var ret string
		return ret
	}
	return *o.HeaderCount
}

// GetHeaderCountOk returns a tuple with the HeaderCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkResponse) GetHeaderCountOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderCount) {
		return nil, false
	}
	return o.HeaderCount, true
}

// HasHeaderCount returns a boolean if a field has been set.
func (o *NetworkResponse) HasHeaderCount() bool {
	if o != nil && !IsNil(o.HeaderCount) {
		return true
	}

	return false
}

// SetHeaderCount gets a reference to the given string and assigns it to the HeaderCount field.
func (o *NetworkResponse) SetHeaderCount(v string) {
	o.HeaderCount = &v
}

// GetTipHashes returns the TipHashes field value if set, zero value otherwise.
func (o *NetworkResponse) GetTipHashes() []string {
	if o == nil || IsNil(o.TipHashes) {
		var ret []string
		return ret
	}
	return o.TipHashes
}

// GetTipHashesOk returns a tuple with the TipHashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkResponse) GetTipHashesOk() ([]string, bool) {
	if o == nil || IsNil(o.TipHashes) {
		return nil, false
	}
	return o.TipHashes, true
}

// HasTipHashes returns a boolean if a field has been set.
func (o *NetworkResponse) HasTipHashes() bool {
	if o != nil && !IsNil(o.TipHashes) {
		return true
	}

	return false
}

// SetTipHashes gets a reference to the given []string and assigns it to the TipHashes field.
func (o *NetworkResponse) SetTipHashes(v []string) {
	o.TipHashes = v
}

// GetDifficulty returns the Difficulty field value if set, zero value otherwise.
func (o *NetworkResponse) GetDifficulty() float32 {
	if o == nil || IsNil(o.Difficulty) {
		var ret float32
		return ret
	}
	return *o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkResponse) GetDifficultyOk() (*float32, bool) {
	if o == nil || IsNil(o.Difficulty) {
		return nil, false
	}
	return o.Difficulty, true
}

// HasDifficulty returns a boolean if a field has been set.
func (o *NetworkResponse) HasDifficulty() bool {
	if o != nil && !IsNil(o.Difficulty) {
		return true
	}

	return false
}

// SetDifficulty gets a reference to the given float32 and assigns it to the Difficulty field.
func (o *NetworkResponse) SetDifficulty(v float32) {
	o.Difficulty = &v
}

// GetPastMedianTime returns the PastMedianTime field value if set, zero value otherwise.
func (o *NetworkResponse) GetPastMedianTime() string {
	if o == nil || IsNil(o.PastMedianTime) {
		var ret string
		return ret
	}
	return *o.PastMedianTime
}

// GetPastMedianTimeOk returns a tuple with the PastMedianTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkResponse) GetPastMedianTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PastMedianTime) {
		return nil, false
	}
	return o.PastMedianTime, true
}

// HasPastMedianTime returns a boolean if a field has been set.
func (o *NetworkResponse) HasPastMedianTime() bool {
	if o != nil && !IsNil(o.PastMedianTime) {
		return true
	}

	return false
}

// SetPastMedianTime gets a reference to the given string and assigns it to the PastMedianTime field.
func (o *NetworkResponse) SetPastMedianTime(v string) {
	o.PastMedianTime = &v
}

// GetVirtualParentHashes returns the VirtualParentHashes field value if set, zero value otherwise.
func (o *NetworkResponse) GetVirtualParentHashes() []string {
	if o == nil || IsNil(o.VirtualParentHashes) {
		var ret []string
		return ret
	}
	return o.VirtualParentHashes
}

// GetVirtualParentHashesOk returns a tuple with the VirtualParentHashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkResponse) GetVirtualParentHashesOk() ([]string, bool) {
	if o == nil || IsNil(o.VirtualParentHashes) {
		return nil, false
	}
	return o.VirtualParentHashes, true
}

// HasVirtualParentHashes returns a boolean if a field has been set.
func (o *NetworkResponse) HasVirtualParentHashes() bool {
	if o != nil && !IsNil(o.VirtualParentHashes) {
		return true
	}

	return false
}

// SetVirtualParentHashes gets a reference to the given []string and assigns it to the VirtualParentHashes field.
func (o *NetworkResponse) SetVirtualParentHashes(v []string) {
	o.VirtualParentHashes = v
}

// GetPruningPointHash returns the PruningPointHash field value if set, zero value otherwise.
func (o *NetworkResponse) GetPruningPointHash() string {
	if o == nil || IsNil(o.PruningPointHash) {
		var ret string
		return ret
	}
	return *o.PruningPointHash
}

// GetPruningPointHashOk returns a tuple with the PruningPointHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkResponse) GetPruningPointHashOk() (*string, bool) {
	if o == nil || IsNil(o.PruningPointHash) {
		return nil, false
	}
	return o.PruningPointHash, true
}

// HasPruningPointHash returns a boolean if a field has been set.
func (o *NetworkResponse) HasPruningPointHash() bool {
	if o != nil && !IsNil(o.PruningPointHash) {
		return true
	}

	return false
}

// SetPruningPointHash gets a reference to the given string and assigns it to the PruningPointHash field.
func (o *NetworkResponse) SetPruningPointHash(v string) {
	o.PruningPointHash = &v
}

// GetVirtualDaaScore returns the VirtualDaaScore field value if set, zero value otherwise.
func (o *NetworkResponse) GetVirtualDaaScore() string {
	if o == nil || IsNil(o.VirtualDaaScore) {
		var ret string
		return ret
	}
	return *o.VirtualDaaScore
}

// GetVirtualDaaScoreOk returns a tuple with the VirtualDaaScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkResponse) GetVirtualDaaScoreOk() (*string, bool) {
	if o == nil || IsNil(o.VirtualDaaScore) {
		return nil, false
	}
	return o.VirtualDaaScore, true
}

// HasVirtualDaaScore returns a boolean if a field has been set.
func (o *NetworkResponse) HasVirtualDaaScore() bool {
	if o != nil && !IsNil(o.VirtualDaaScore) {
		return true
	}

	return false
}

// SetVirtualDaaScore gets a reference to the given string and assigns it to the VirtualDaaScore field.
func (o *NetworkResponse) SetVirtualDaaScore(v string) {
	o.VirtualDaaScore = &v
}

func (o NetworkResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkName) {
		toSerialize["networkName"] = o.NetworkName
	}
	if !IsNil(o.BlockCount) {
		toSerialize["blockCount"] = o.BlockCount
	}
	if !IsNil(o.HeaderCount) {
		toSerialize["headerCount"] = o.HeaderCount
	}
	if !IsNil(o.TipHashes) {
		toSerialize["tipHashes"] = o.TipHashes
	}
	if !IsNil(o.Difficulty) {
		toSerialize["difficulty"] = o.Difficulty
	}
	if !IsNil(o.PastMedianTime) {
		toSerialize["pastMedianTime"] = o.PastMedianTime
	}
	if !IsNil(o.VirtualParentHashes) {
		toSerialize["virtualParentHashes"] = o.VirtualParentHashes
	}
	if !IsNil(o.PruningPointHash) {
		toSerialize["pruningPointHash"] = o.PruningPointHash
	}
	if !IsNil(o.VirtualDaaScore) {
		toSerialize["virtualDaaScore"] = o.VirtualDaaScore
	}
	return toSerialize, nil
}

type NullableNetworkResponse struct {
	value *NetworkResponse
	isSet bool
}

func (v NullableNetworkResponse) Get() *NetworkResponse {
	return v.value
}

func (v *NullableNetworkResponse) Set(val *NetworkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkResponse(val *NetworkResponse) *NullableNetworkResponse {
	return &NullableNetworkResponse{value: val, isSet: true}
}

func (v NullableNetworkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


