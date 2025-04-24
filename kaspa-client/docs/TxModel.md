# TxModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetworkId** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Mass** | Pointer to **int32** |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**BlockHash** | Pointer to **[]string** |  | [optional] 
**BlockTime** | Pointer to **int32** |  | [optional] 
**IsAccepted** | Pointer to **bool** |  | [optional] 
**AcceptingBlockHash** | Pointer to **string** |  | [optional] 
**AcceptingBlockBlueScore** | Pointer to **int32** |  | [optional] 
**AcceptingBlockTime** | Pointer to **int32** |  | [optional] 
**Inputs** | Pointer to [**[]TxInput**](TxInput.md) |  | [optional] 
**Outputs** | Pointer to [**[]TxOutput**](TxOutput.md) |  | [optional] 

## Methods

### NewTxModel

`func NewTxModel() *TxModel`

NewTxModel instantiates a new TxModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxModelWithDefaults

`func NewTxModelWithDefaults() *TxModel`

NewTxModelWithDefaults instantiates a new TxModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetworkId

`func (o *TxModel) GetSubnetworkId() string`

GetSubnetworkId returns the SubnetworkId field if non-nil, zero value otherwise.

### GetSubnetworkIdOk

`func (o *TxModel) GetSubnetworkIdOk() (*string, bool)`

GetSubnetworkIdOk returns a tuple with the SubnetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetworkId

`func (o *TxModel) SetSubnetworkId(v string)`

SetSubnetworkId sets SubnetworkId field to given value.

### HasSubnetworkId

`func (o *TxModel) HasSubnetworkId() bool`

HasSubnetworkId returns a boolean if a field has been set.

### GetTransactionId

`func (o *TxModel) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TxModel) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TxModel) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *TxModel) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetHash

`func (o *TxModel) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *TxModel) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *TxModel) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *TxModel) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetMass

`func (o *TxModel) GetMass() int32`

GetMass returns the Mass field if non-nil, zero value otherwise.

### GetMassOk

`func (o *TxModel) GetMassOk() (*int32, bool)`

GetMassOk returns a tuple with the Mass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMass

`func (o *TxModel) SetMass(v int32)`

SetMass sets Mass field to given value.

### HasMass

`func (o *TxModel) HasMass() bool`

HasMass returns a boolean if a field has been set.

### GetPayload

`func (o *TxModel) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *TxModel) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *TxModel) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *TxModel) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetBlockHash

`func (o *TxModel) GetBlockHash() []string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TxModel) GetBlockHashOk() (*[]string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TxModel) SetBlockHash(v []string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *TxModel) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetBlockTime

`func (o *TxModel) GetBlockTime() int32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *TxModel) GetBlockTimeOk() (*int32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *TxModel) SetBlockTime(v int32)`

SetBlockTime sets BlockTime field to given value.

### HasBlockTime

`func (o *TxModel) HasBlockTime() bool`

HasBlockTime returns a boolean if a field has been set.

### GetIsAccepted

`func (o *TxModel) GetIsAccepted() bool`

GetIsAccepted returns the IsAccepted field if non-nil, zero value otherwise.

### GetIsAcceptedOk

`func (o *TxModel) GetIsAcceptedOk() (*bool, bool)`

GetIsAcceptedOk returns a tuple with the IsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccepted

`func (o *TxModel) SetIsAccepted(v bool)`

SetIsAccepted sets IsAccepted field to given value.

### HasIsAccepted

`func (o *TxModel) HasIsAccepted() bool`

HasIsAccepted returns a boolean if a field has been set.

### GetAcceptingBlockHash

`func (o *TxModel) GetAcceptingBlockHash() string`

GetAcceptingBlockHash returns the AcceptingBlockHash field if non-nil, zero value otherwise.

### GetAcceptingBlockHashOk

`func (o *TxModel) GetAcceptingBlockHashOk() (*string, bool)`

GetAcceptingBlockHashOk returns a tuple with the AcceptingBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptingBlockHash

`func (o *TxModel) SetAcceptingBlockHash(v string)`

SetAcceptingBlockHash sets AcceptingBlockHash field to given value.

### HasAcceptingBlockHash

`func (o *TxModel) HasAcceptingBlockHash() bool`

HasAcceptingBlockHash returns a boolean if a field has been set.

### GetAcceptingBlockBlueScore

`func (o *TxModel) GetAcceptingBlockBlueScore() int32`

GetAcceptingBlockBlueScore returns the AcceptingBlockBlueScore field if non-nil, zero value otherwise.

### GetAcceptingBlockBlueScoreOk

`func (o *TxModel) GetAcceptingBlockBlueScoreOk() (*int32, bool)`

GetAcceptingBlockBlueScoreOk returns a tuple with the AcceptingBlockBlueScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptingBlockBlueScore

`func (o *TxModel) SetAcceptingBlockBlueScore(v int32)`

SetAcceptingBlockBlueScore sets AcceptingBlockBlueScore field to given value.

### HasAcceptingBlockBlueScore

`func (o *TxModel) HasAcceptingBlockBlueScore() bool`

HasAcceptingBlockBlueScore returns a boolean if a field has been set.

### GetAcceptingBlockTime

`func (o *TxModel) GetAcceptingBlockTime() int32`

GetAcceptingBlockTime returns the AcceptingBlockTime field if non-nil, zero value otherwise.

### GetAcceptingBlockTimeOk

`func (o *TxModel) GetAcceptingBlockTimeOk() (*int32, bool)`

GetAcceptingBlockTimeOk returns a tuple with the AcceptingBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptingBlockTime

`func (o *TxModel) SetAcceptingBlockTime(v int32)`

SetAcceptingBlockTime sets AcceptingBlockTime field to given value.

### HasAcceptingBlockTime

`func (o *TxModel) HasAcceptingBlockTime() bool`

HasAcceptingBlockTime returns a boolean if a field has been set.

### GetInputs

`func (o *TxModel) GetInputs() []TxInput`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *TxModel) GetInputsOk() (*[]TxInput, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *TxModel) SetInputs(v []TxInput)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *TxModel) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetOutputs

`func (o *TxModel) GetOutputs() []TxOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *TxModel) GetOutputsOk() (*[]TxOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *TxModel) SetOutputs(v []TxOutput)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *TxModel) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


