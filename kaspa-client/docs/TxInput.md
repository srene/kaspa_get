# TxInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** |  | 
**Index** | **int32** |  | 
**PreviousOutpointHash** | **string** |  | 
**PreviousOutpointIndex** | **string** |  | 
**PreviousOutpointResolved** | Pointer to [**TxOutput**](TxOutput.md) |  | [optional] 
**PreviousOutpointAddress** | Pointer to **string** |  | [optional] 
**PreviousOutpointAmount** | Pointer to **int32** |  | [optional] 
**SignatureScript** | Pointer to **string** |  | [optional] 
**SigOpCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewTxInput

`func NewTxInput(transactionId string, index int32, previousOutpointHash string, previousOutpointIndex string, ) *TxInput`

NewTxInput instantiates a new TxInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxInputWithDefaults

`func NewTxInputWithDefaults() *TxInput`

NewTxInputWithDefaults instantiates a new TxInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *TxInput) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TxInput) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TxInput) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetIndex

`func (o *TxInput) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *TxInput) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *TxInput) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetPreviousOutpointHash

`func (o *TxInput) GetPreviousOutpointHash() string`

GetPreviousOutpointHash returns the PreviousOutpointHash field if non-nil, zero value otherwise.

### GetPreviousOutpointHashOk

`func (o *TxInput) GetPreviousOutpointHashOk() (*string, bool)`

GetPreviousOutpointHashOk returns a tuple with the PreviousOutpointHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousOutpointHash

`func (o *TxInput) SetPreviousOutpointHash(v string)`

SetPreviousOutpointHash sets PreviousOutpointHash field to given value.


### GetPreviousOutpointIndex

`func (o *TxInput) GetPreviousOutpointIndex() string`

GetPreviousOutpointIndex returns the PreviousOutpointIndex field if non-nil, zero value otherwise.

### GetPreviousOutpointIndexOk

`func (o *TxInput) GetPreviousOutpointIndexOk() (*string, bool)`

GetPreviousOutpointIndexOk returns a tuple with the PreviousOutpointIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousOutpointIndex

`func (o *TxInput) SetPreviousOutpointIndex(v string)`

SetPreviousOutpointIndex sets PreviousOutpointIndex field to given value.


### GetPreviousOutpointResolved

`func (o *TxInput) GetPreviousOutpointResolved() TxOutput`

GetPreviousOutpointResolved returns the PreviousOutpointResolved field if non-nil, zero value otherwise.

### GetPreviousOutpointResolvedOk

`func (o *TxInput) GetPreviousOutpointResolvedOk() (*TxOutput, bool)`

GetPreviousOutpointResolvedOk returns a tuple with the PreviousOutpointResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousOutpointResolved

`func (o *TxInput) SetPreviousOutpointResolved(v TxOutput)`

SetPreviousOutpointResolved sets PreviousOutpointResolved field to given value.

### HasPreviousOutpointResolved

`func (o *TxInput) HasPreviousOutpointResolved() bool`

HasPreviousOutpointResolved returns a boolean if a field has been set.

### GetPreviousOutpointAddress

`func (o *TxInput) GetPreviousOutpointAddress() string`

GetPreviousOutpointAddress returns the PreviousOutpointAddress field if non-nil, zero value otherwise.

### GetPreviousOutpointAddressOk

`func (o *TxInput) GetPreviousOutpointAddressOk() (*string, bool)`

GetPreviousOutpointAddressOk returns a tuple with the PreviousOutpointAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousOutpointAddress

`func (o *TxInput) SetPreviousOutpointAddress(v string)`

SetPreviousOutpointAddress sets PreviousOutpointAddress field to given value.

### HasPreviousOutpointAddress

`func (o *TxInput) HasPreviousOutpointAddress() bool`

HasPreviousOutpointAddress returns a boolean if a field has been set.

### GetPreviousOutpointAmount

`func (o *TxInput) GetPreviousOutpointAmount() int32`

GetPreviousOutpointAmount returns the PreviousOutpointAmount field if non-nil, zero value otherwise.

### GetPreviousOutpointAmountOk

`func (o *TxInput) GetPreviousOutpointAmountOk() (*int32, bool)`

GetPreviousOutpointAmountOk returns a tuple with the PreviousOutpointAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousOutpointAmount

`func (o *TxInput) SetPreviousOutpointAmount(v int32)`

SetPreviousOutpointAmount sets PreviousOutpointAmount field to given value.

### HasPreviousOutpointAmount

`func (o *TxInput) HasPreviousOutpointAmount() bool`

HasPreviousOutpointAmount returns a boolean if a field has been set.

### GetSignatureScript

`func (o *TxInput) GetSignatureScript() string`

GetSignatureScript returns the SignatureScript field if non-nil, zero value otherwise.

### GetSignatureScriptOk

`func (o *TxInput) GetSignatureScriptOk() (*string, bool)`

GetSignatureScriptOk returns a tuple with the SignatureScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureScript

`func (o *TxInput) SetSignatureScript(v string)`

SetSignatureScript sets SignatureScript field to given value.

### HasSignatureScript

`func (o *TxInput) HasSignatureScript() bool`

HasSignatureScript returns a boolean if a field has been set.

### GetSigOpCount

`func (o *TxInput) GetSigOpCount() int32`

GetSigOpCount returns the SigOpCount field if non-nil, zero value otherwise.

### GetSigOpCountOk

`func (o *TxInput) GetSigOpCountOk() (*int32, bool)`

GetSigOpCountOk returns a tuple with the SigOpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigOpCount

`func (o *TxInput) SetSigOpCount(v int32)`

SetSigOpCount sets SigOpCount field to given value.

### HasSigOpCount

`func (o *TxInput) HasSigOpCount() bool`

HasSigOpCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


