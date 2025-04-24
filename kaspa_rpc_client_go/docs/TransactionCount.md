# TransactionCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**LimitExceeded** | **bool** |  | 

## Methods

### NewTransactionCount

`func NewTransactionCount(total int32, limitExceeded bool, ) *TransactionCount`

NewTransactionCount instantiates a new TransactionCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCountWithDefaults

`func NewTransactionCountWithDefaults() *TransactionCount`

NewTransactionCountWithDefaults instantiates a new TransactionCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *TransactionCount) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TransactionCount) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TransactionCount) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetLimitExceeded

`func (o *TransactionCount) GetLimitExceeded() bool`

GetLimitExceeded returns the LimitExceeded field if non-nil, zero value otherwise.

### GetLimitExceededOk

`func (o *TransactionCount) GetLimitExceededOk() (*bool, bool)`

GetLimitExceededOk returns a tuple with the LimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitExceeded

`func (o *TransactionCount) SetLimitExceeded(v bool)`

SetLimitExceeded sets LimitExceeded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


