# TxOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** |  | 
**Index** | **int32** |  | 
**Amount** | **int32** |  | 
**ScriptPublicKey** | Pointer to **string** |  | [optional] 
**ScriptPublicKeyAddress** | Pointer to **string** |  | [optional] 
**ScriptPublicKeyType** | Pointer to **string** |  | [optional] 
**AcceptingBlockHash** | Pointer to **string** |  | [optional] 

## Methods

### NewTxOutput

`func NewTxOutput(transactionId string, index int32, amount int32, ) *TxOutput`

NewTxOutput instantiates a new TxOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxOutputWithDefaults

`func NewTxOutputWithDefaults() *TxOutput`

NewTxOutputWithDefaults instantiates a new TxOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *TxOutput) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TxOutput) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TxOutput) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetIndex

`func (o *TxOutput) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *TxOutput) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *TxOutput) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetAmount

`func (o *TxOutput) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TxOutput) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TxOutput) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetScriptPublicKey

`func (o *TxOutput) GetScriptPublicKey() string`

GetScriptPublicKey returns the ScriptPublicKey field if non-nil, zero value otherwise.

### GetScriptPublicKeyOk

`func (o *TxOutput) GetScriptPublicKeyOk() (*string, bool)`

GetScriptPublicKeyOk returns a tuple with the ScriptPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPublicKey

`func (o *TxOutput) SetScriptPublicKey(v string)`

SetScriptPublicKey sets ScriptPublicKey field to given value.

### HasScriptPublicKey

`func (o *TxOutput) HasScriptPublicKey() bool`

HasScriptPublicKey returns a boolean if a field has been set.

### GetScriptPublicKeyAddress

`func (o *TxOutput) GetScriptPublicKeyAddress() string`

GetScriptPublicKeyAddress returns the ScriptPublicKeyAddress field if non-nil, zero value otherwise.

### GetScriptPublicKeyAddressOk

`func (o *TxOutput) GetScriptPublicKeyAddressOk() (*string, bool)`

GetScriptPublicKeyAddressOk returns a tuple with the ScriptPublicKeyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPublicKeyAddress

`func (o *TxOutput) SetScriptPublicKeyAddress(v string)`

SetScriptPublicKeyAddress sets ScriptPublicKeyAddress field to given value.

### HasScriptPublicKeyAddress

`func (o *TxOutput) HasScriptPublicKeyAddress() bool`

HasScriptPublicKeyAddress returns a boolean if a field has been set.

### GetScriptPublicKeyType

`func (o *TxOutput) GetScriptPublicKeyType() string`

GetScriptPublicKeyType returns the ScriptPublicKeyType field if non-nil, zero value otherwise.

### GetScriptPublicKeyTypeOk

`func (o *TxOutput) GetScriptPublicKeyTypeOk() (*string, bool)`

GetScriptPublicKeyTypeOk returns a tuple with the ScriptPublicKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPublicKeyType

`func (o *TxOutput) SetScriptPublicKeyType(v string)`

SetScriptPublicKeyType sets ScriptPublicKeyType field to given value.

### HasScriptPublicKeyType

`func (o *TxOutput) HasScriptPublicKeyType() bool`

HasScriptPublicKeyType returns a boolean if a field has been set.

### GetAcceptingBlockHash

`func (o *TxOutput) GetAcceptingBlockHash() string`

GetAcceptingBlockHash returns the AcceptingBlockHash field if non-nil, zero value otherwise.

### GetAcceptingBlockHashOk

`func (o *TxOutput) GetAcceptingBlockHashOk() (*string, bool)`

GetAcceptingBlockHashOk returns a tuple with the AcceptingBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptingBlockHash

`func (o *TxOutput) SetAcceptingBlockHash(v string)`

SetAcceptingBlockHash sets AcceptingBlockHash field to given value.

### HasAcceptingBlockHash

`func (o *TxOutput) HasAcceptingBlockHash() bool`

HasAcceptingBlockHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


