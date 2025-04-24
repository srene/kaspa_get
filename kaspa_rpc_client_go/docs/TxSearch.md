# TxSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionIds** | Pointer to **[]string** |  | [optional] 
**AcceptingBlueScores** | Pointer to [**TxSearchAcceptingBlueScores**](TxSearchAcceptingBlueScores.md) |  | [optional] 

## Methods

### NewTxSearch

`func NewTxSearch() *TxSearch`

NewTxSearch instantiates a new TxSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxSearchWithDefaults

`func NewTxSearchWithDefaults() *TxSearch`

NewTxSearchWithDefaults instantiates a new TxSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionIds

`func (o *TxSearch) GetTransactionIds() []string`

GetTransactionIds returns the TransactionIds field if non-nil, zero value otherwise.

### GetTransactionIdsOk

`func (o *TxSearch) GetTransactionIdsOk() (*[]string, bool)`

GetTransactionIdsOk returns a tuple with the TransactionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIds

`func (o *TxSearch) SetTransactionIds(v []string)`

SetTransactionIds sets TransactionIds field to given value.

### HasTransactionIds

`func (o *TxSearch) HasTransactionIds() bool`

HasTransactionIds returns a boolean if a field has been set.

### GetAcceptingBlueScores

`func (o *TxSearch) GetAcceptingBlueScores() TxSearchAcceptingBlueScores`

GetAcceptingBlueScores returns the AcceptingBlueScores field if non-nil, zero value otherwise.

### GetAcceptingBlueScoresOk

`func (o *TxSearch) GetAcceptingBlueScoresOk() (*TxSearchAcceptingBlueScores, bool)`

GetAcceptingBlueScoresOk returns a tuple with the AcceptingBlueScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptingBlueScores

`func (o *TxSearch) SetAcceptingBlueScores(v TxSearchAcceptingBlueScores)`

SetAcceptingBlueScores sets AcceptingBlueScores field to given value.

### HasAcceptingBlueScores

`func (o *TxSearch) HasAcceptingBlueScores() bool`

HasAcceptingBlueScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


