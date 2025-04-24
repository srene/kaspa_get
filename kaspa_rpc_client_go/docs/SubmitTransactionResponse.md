# SubmitTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewSubmitTransactionResponse

`func NewSubmitTransactionResponse() *SubmitTransactionResponse`

NewSubmitTransactionResponse instantiates a new SubmitTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitTransactionResponseWithDefaults

`func NewSubmitTransactionResponseWithDefaults() *SubmitTransactionResponse`

NewSubmitTransactionResponseWithDefaults instantiates a new SubmitTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *SubmitTransactionResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *SubmitTransactionResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *SubmitTransactionResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *SubmitTransactionResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetError

`func (o *SubmitTransactionResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SubmitTransactionResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SubmitTransactionResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SubmitTransactionResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


