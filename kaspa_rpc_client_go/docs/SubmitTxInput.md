# SubmitTxInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviousOutpoint** | [**SubmitTxOutpoint**](SubmitTxOutpoint.md) |  | 
**SignatureScript** | **string** |  | 
**Sequence** | **int32** |  | 
**SigOpCount** | **int32** |  | 

## Methods

### NewSubmitTxInput

`func NewSubmitTxInput(previousOutpoint SubmitTxOutpoint, signatureScript string, sequence int32, sigOpCount int32, ) *SubmitTxInput`

NewSubmitTxInput instantiates a new SubmitTxInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitTxInputWithDefaults

`func NewSubmitTxInputWithDefaults() *SubmitTxInput`

NewSubmitTxInputWithDefaults instantiates a new SubmitTxInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviousOutpoint

`func (o *SubmitTxInput) GetPreviousOutpoint() SubmitTxOutpoint`

GetPreviousOutpoint returns the PreviousOutpoint field if non-nil, zero value otherwise.

### GetPreviousOutpointOk

`func (o *SubmitTxInput) GetPreviousOutpointOk() (*SubmitTxOutpoint, bool)`

GetPreviousOutpointOk returns a tuple with the PreviousOutpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousOutpoint

`func (o *SubmitTxInput) SetPreviousOutpoint(v SubmitTxOutpoint)`

SetPreviousOutpoint sets PreviousOutpoint field to given value.


### GetSignatureScript

`func (o *SubmitTxInput) GetSignatureScript() string`

GetSignatureScript returns the SignatureScript field if non-nil, zero value otherwise.

### GetSignatureScriptOk

`func (o *SubmitTxInput) GetSignatureScriptOk() (*string, bool)`

GetSignatureScriptOk returns a tuple with the SignatureScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureScript

`func (o *SubmitTxInput) SetSignatureScript(v string)`

SetSignatureScript sets SignatureScript field to given value.


### GetSequence

`func (o *SubmitTxInput) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *SubmitTxInput) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *SubmitTxInput) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetSigOpCount

`func (o *SubmitTxInput) GetSigOpCount() int32`

GetSigOpCount returns the SigOpCount field if non-nil, zero value otherwise.

### GetSigOpCountOk

`func (o *SubmitTxInput) GetSigOpCountOk() (*int32, bool)`

GetSigOpCountOk returns a tuple with the SigOpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigOpCount

`func (o *SubmitTxInput) SetSigOpCount(v int32)`

SetSigOpCount sets SigOpCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


