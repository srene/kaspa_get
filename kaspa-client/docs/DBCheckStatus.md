# DBCheckStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSynced** | Pointer to **bool** |  | [optional] [default to true]
**BlueScore** | Pointer to **int32** |  | [optional] 
**BlueScoreDiff** | Pointer to **int32** |  | [optional] 
**AcceptedTxBlockTime** | Pointer to **int32** |  | [optional] 
**AcceptedTxBlockTimeDiff** | Pointer to **int32** |  | [optional] 

## Methods

### NewDBCheckStatus

`func NewDBCheckStatus() *DBCheckStatus`

NewDBCheckStatus instantiates a new DBCheckStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDBCheckStatusWithDefaults

`func NewDBCheckStatusWithDefaults() *DBCheckStatus`

NewDBCheckStatusWithDefaults instantiates a new DBCheckStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSynced

`func (o *DBCheckStatus) GetIsSynced() bool`

GetIsSynced returns the IsSynced field if non-nil, zero value otherwise.

### GetIsSyncedOk

`func (o *DBCheckStatus) GetIsSyncedOk() (*bool, bool)`

GetIsSyncedOk returns a tuple with the IsSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynced

`func (o *DBCheckStatus) SetIsSynced(v bool)`

SetIsSynced sets IsSynced field to given value.

### HasIsSynced

`func (o *DBCheckStatus) HasIsSynced() bool`

HasIsSynced returns a boolean if a field has been set.

### GetBlueScore

`func (o *DBCheckStatus) GetBlueScore() int32`

GetBlueScore returns the BlueScore field if non-nil, zero value otherwise.

### GetBlueScoreOk

`func (o *DBCheckStatus) GetBlueScoreOk() (*int32, bool)`

GetBlueScoreOk returns a tuple with the BlueScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueScore

`func (o *DBCheckStatus) SetBlueScore(v int32)`

SetBlueScore sets BlueScore field to given value.

### HasBlueScore

`func (o *DBCheckStatus) HasBlueScore() bool`

HasBlueScore returns a boolean if a field has been set.

### GetBlueScoreDiff

`func (o *DBCheckStatus) GetBlueScoreDiff() int32`

GetBlueScoreDiff returns the BlueScoreDiff field if non-nil, zero value otherwise.

### GetBlueScoreDiffOk

`func (o *DBCheckStatus) GetBlueScoreDiffOk() (*int32, bool)`

GetBlueScoreDiffOk returns a tuple with the BlueScoreDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueScoreDiff

`func (o *DBCheckStatus) SetBlueScoreDiff(v int32)`

SetBlueScoreDiff sets BlueScoreDiff field to given value.

### HasBlueScoreDiff

`func (o *DBCheckStatus) HasBlueScoreDiff() bool`

HasBlueScoreDiff returns a boolean if a field has been set.

### GetAcceptedTxBlockTime

`func (o *DBCheckStatus) GetAcceptedTxBlockTime() int32`

GetAcceptedTxBlockTime returns the AcceptedTxBlockTime field if non-nil, zero value otherwise.

### GetAcceptedTxBlockTimeOk

`func (o *DBCheckStatus) GetAcceptedTxBlockTimeOk() (*int32, bool)`

GetAcceptedTxBlockTimeOk returns a tuple with the AcceptedTxBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTxBlockTime

`func (o *DBCheckStatus) SetAcceptedTxBlockTime(v int32)`

SetAcceptedTxBlockTime sets AcceptedTxBlockTime field to given value.

### HasAcceptedTxBlockTime

`func (o *DBCheckStatus) HasAcceptedTxBlockTime() bool`

HasAcceptedTxBlockTime returns a boolean if a field has been set.

### GetAcceptedTxBlockTimeDiff

`func (o *DBCheckStatus) GetAcceptedTxBlockTimeDiff() int32`

GetAcceptedTxBlockTimeDiff returns the AcceptedTxBlockTimeDiff field if non-nil, zero value otherwise.

### GetAcceptedTxBlockTimeDiffOk

`func (o *DBCheckStatus) GetAcceptedTxBlockTimeDiffOk() (*int32, bool)`

GetAcceptedTxBlockTimeDiffOk returns a tuple with the AcceptedTxBlockTimeDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTxBlockTimeDiff

`func (o *DBCheckStatus) SetAcceptedTxBlockTimeDiff(v int32)`

SetAcceptedTxBlockTimeDiff sets AcceptedTxBlockTimeDiff field to given value.

### HasAcceptedTxBlockTimeDiff

`func (o *DBCheckStatus) HasAcceptedTxBlockTimeDiff() bool`

HasAcceptedTxBlockTimeDiff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


