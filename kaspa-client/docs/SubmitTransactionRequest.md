# SubmitTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transaction** | [**SubmitTxModel**](SubmitTxModel.md) |  | 
**AllowOrphan** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSubmitTransactionRequest

`func NewSubmitTransactionRequest(transaction SubmitTxModel, ) *SubmitTransactionRequest`

NewSubmitTransactionRequest instantiates a new SubmitTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitTransactionRequestWithDefaults

`func NewSubmitTransactionRequestWithDefaults() *SubmitTransactionRequest`

NewSubmitTransactionRequestWithDefaults instantiates a new SubmitTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransaction

`func (o *SubmitTransactionRequest) GetTransaction() SubmitTxModel`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *SubmitTransactionRequest) GetTransactionOk() (*SubmitTxModel, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *SubmitTransactionRequest) SetTransaction(v SubmitTxModel)`

SetTransaction sets Transaction field to given value.


### GetAllowOrphan

`func (o *SubmitTransactionRequest) GetAllowOrphan() bool`

GetAllowOrphan returns the AllowOrphan field if non-nil, zero value otherwise.

### GetAllowOrphanOk

`func (o *SubmitTransactionRequest) GetAllowOrphanOk() (*bool, bool)`

GetAllowOrphanOk returns a tuple with the AllowOrphan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOrphan

`func (o *SubmitTransactionRequest) SetAllowOrphan(v bool)`

SetAllowOrphan sets AllowOrphan field to given value.

### HasAllowOrphan

`func (o *SubmitTransactionRequest) HasAllowOrphan() bool`

HasAllowOrphan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


