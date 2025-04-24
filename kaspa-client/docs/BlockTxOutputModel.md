# BlockTxOutputModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**ScriptPublicKey** | Pointer to [**BlockTxOutputScriptPublicKeyModel**](BlockTxOutputScriptPublicKeyModel.md) |  | [optional] 
**VerboseData** | Pointer to [**BlockTxOutputVerboseDataModel**](BlockTxOutputVerboseDataModel.md) |  | [optional] 

## Methods

### NewBlockTxOutputModel

`func NewBlockTxOutputModel() *BlockTxOutputModel`

NewBlockTxOutputModel instantiates a new BlockTxOutputModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockTxOutputModelWithDefaults

`func NewBlockTxOutputModelWithDefaults() *BlockTxOutputModel`

NewBlockTxOutputModelWithDefaults instantiates a new BlockTxOutputModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BlockTxOutputModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BlockTxOutputModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BlockTxOutputModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BlockTxOutputModel) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetScriptPublicKey

`func (o *BlockTxOutputModel) GetScriptPublicKey() BlockTxOutputScriptPublicKeyModel`

GetScriptPublicKey returns the ScriptPublicKey field if non-nil, zero value otherwise.

### GetScriptPublicKeyOk

`func (o *BlockTxOutputModel) GetScriptPublicKeyOk() (*BlockTxOutputScriptPublicKeyModel, bool)`

GetScriptPublicKeyOk returns a tuple with the ScriptPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPublicKey

`func (o *BlockTxOutputModel) SetScriptPublicKey(v BlockTxOutputScriptPublicKeyModel)`

SetScriptPublicKey sets ScriptPublicKey field to given value.

### HasScriptPublicKey

`func (o *BlockTxOutputModel) HasScriptPublicKey() bool`

HasScriptPublicKey returns a boolean if a field has been set.

### GetVerboseData

`func (o *BlockTxOutputModel) GetVerboseData() BlockTxOutputVerboseDataModel`

GetVerboseData returns the VerboseData field if non-nil, zero value otherwise.

### GetVerboseDataOk

`func (o *BlockTxOutputModel) GetVerboseDataOk() (*BlockTxOutputVerboseDataModel, bool)`

GetVerboseDataOk returns a tuple with the VerboseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseData

`func (o *BlockTxOutputModel) SetVerboseData(v BlockTxOutputVerboseDataModel)`

SetVerboseData sets VerboseData field to given value.

### HasVerboseData

`func (o *BlockTxOutputModel) HasVerboseData() bool`

HasVerboseData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


