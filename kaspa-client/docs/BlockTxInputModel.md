# BlockTxInputModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviousOutpoint** | Pointer to [**BlockTxInputPreviousOutpointModel**](BlockTxInputPreviousOutpointModel.md) |  | [optional] 
**SignatureScript** | Pointer to **string** |  | [optional] 
**SigOpCount** | Pointer to **int32** |  | [optional] 
**Sequence** | Pointer to **int32** |  | [optional] 

## Methods

### NewBlockTxInputModel

`func NewBlockTxInputModel() *BlockTxInputModel`

NewBlockTxInputModel instantiates a new BlockTxInputModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockTxInputModelWithDefaults

`func NewBlockTxInputModelWithDefaults() *BlockTxInputModel`

NewBlockTxInputModelWithDefaults instantiates a new BlockTxInputModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviousOutpoint

`func (o *BlockTxInputModel) GetPreviousOutpoint() BlockTxInputPreviousOutpointModel`

GetPreviousOutpoint returns the PreviousOutpoint field if non-nil, zero value otherwise.

### GetPreviousOutpointOk

`func (o *BlockTxInputModel) GetPreviousOutpointOk() (*BlockTxInputPreviousOutpointModel, bool)`

GetPreviousOutpointOk returns a tuple with the PreviousOutpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousOutpoint

`func (o *BlockTxInputModel) SetPreviousOutpoint(v BlockTxInputPreviousOutpointModel)`

SetPreviousOutpoint sets PreviousOutpoint field to given value.

### HasPreviousOutpoint

`func (o *BlockTxInputModel) HasPreviousOutpoint() bool`

HasPreviousOutpoint returns a boolean if a field has been set.

### GetSignatureScript

`func (o *BlockTxInputModel) GetSignatureScript() string`

GetSignatureScript returns the SignatureScript field if non-nil, zero value otherwise.

### GetSignatureScriptOk

`func (o *BlockTxInputModel) GetSignatureScriptOk() (*string, bool)`

GetSignatureScriptOk returns a tuple with the SignatureScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureScript

`func (o *BlockTxInputModel) SetSignatureScript(v string)`

SetSignatureScript sets SignatureScript field to given value.

### HasSignatureScript

`func (o *BlockTxInputModel) HasSignatureScript() bool`

HasSignatureScript returns a boolean if a field has been set.

### GetSigOpCount

`func (o *BlockTxInputModel) GetSigOpCount() int32`

GetSigOpCount returns the SigOpCount field if non-nil, zero value otherwise.

### GetSigOpCountOk

`func (o *BlockTxInputModel) GetSigOpCountOk() (*int32, bool)`

GetSigOpCountOk returns a tuple with the SigOpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigOpCount

`func (o *BlockTxInputModel) SetSigOpCount(v int32)`

SetSigOpCount sets SigOpCount field to given value.

### HasSigOpCount

`func (o *BlockTxInputModel) HasSigOpCount() bool`

HasSigOpCount returns a boolean if a field has been set.

### GetSequence

`func (o *BlockTxInputModel) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *BlockTxInputModel) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *BlockTxInputModel) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *BlockTxInputModel) HasSequence() bool`

HasSequence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


