# UtxoModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] [default to "[11501593788]"]
**ScriptPublicKey** | [**ScriptPublicKeyModel**](ScriptPublicKeyModel.md) |  | 
**BlockDaaScore** | Pointer to **string** |  | [optional] [default to "18867232"]
**IsCoinbase** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewUtxoModel

`func NewUtxoModel(scriptPublicKey ScriptPublicKeyModel, ) *UtxoModel`

NewUtxoModel instantiates a new UtxoModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtxoModelWithDefaults

`func NewUtxoModelWithDefaults() *UtxoModel`

NewUtxoModelWithDefaults instantiates a new UtxoModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UtxoModel) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UtxoModel) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UtxoModel) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UtxoModel) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetScriptPublicKey

`func (o *UtxoModel) GetScriptPublicKey() ScriptPublicKeyModel`

GetScriptPublicKey returns the ScriptPublicKey field if non-nil, zero value otherwise.

### GetScriptPublicKeyOk

`func (o *UtxoModel) GetScriptPublicKeyOk() (*ScriptPublicKeyModel, bool)`

GetScriptPublicKeyOk returns a tuple with the ScriptPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPublicKey

`func (o *UtxoModel) SetScriptPublicKey(v ScriptPublicKeyModel)`

SetScriptPublicKey sets ScriptPublicKey field to given value.


### GetBlockDaaScore

`func (o *UtxoModel) GetBlockDaaScore() string`

GetBlockDaaScore returns the BlockDaaScore field if non-nil, zero value otherwise.

### GetBlockDaaScoreOk

`func (o *UtxoModel) GetBlockDaaScoreOk() (*string, bool)`

GetBlockDaaScoreOk returns a tuple with the BlockDaaScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDaaScore

`func (o *UtxoModel) SetBlockDaaScore(v string)`

SetBlockDaaScore sets BlockDaaScore field to given value.

### HasBlockDaaScore

`func (o *UtxoModel) HasBlockDaaScore() bool`

HasBlockDaaScore returns a boolean if a field has been set.

### GetIsCoinbase

`func (o *UtxoModel) GetIsCoinbase() bool`

GetIsCoinbase returns the IsCoinbase field if non-nil, zero value otherwise.

### GetIsCoinbaseOk

`func (o *UtxoModel) GetIsCoinbaseOk() (*bool, bool)`

GetIsCoinbaseOk returns a tuple with the IsCoinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCoinbase

`func (o *UtxoModel) SetIsCoinbase(v bool)`

SetIsCoinbase sets IsCoinbase field to given value.

### HasIsCoinbase

`func (o *UtxoModel) HasIsCoinbase() bool`

HasIsCoinbase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


