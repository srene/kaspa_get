# FeeEstimateBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feerate** | Pointer to **int32** |  | [optional] [default to 1]
**EstimatedSeconds** | Pointer to **float32** |  | [optional] [default to 0.004]

## Methods

### NewFeeEstimateBucket

`func NewFeeEstimateBucket() *FeeEstimateBucket`

NewFeeEstimateBucket instantiates a new FeeEstimateBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeEstimateBucketWithDefaults

`func NewFeeEstimateBucketWithDefaults() *FeeEstimateBucket`

NewFeeEstimateBucketWithDefaults instantiates a new FeeEstimateBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeerate

`func (o *FeeEstimateBucket) GetFeerate() int32`

GetFeerate returns the Feerate field if non-nil, zero value otherwise.

### GetFeerateOk

`func (o *FeeEstimateBucket) GetFeerateOk() (*int32, bool)`

GetFeerateOk returns a tuple with the Feerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeerate

`func (o *FeeEstimateBucket) SetFeerate(v int32)`

SetFeerate sets Feerate field to given value.

### HasFeerate

`func (o *FeeEstimateBucket) HasFeerate() bool`

HasFeerate returns a boolean if a field has been set.

### GetEstimatedSeconds

`func (o *FeeEstimateBucket) GetEstimatedSeconds() float32`

GetEstimatedSeconds returns the EstimatedSeconds field if non-nil, zero value otherwise.

### GetEstimatedSecondsOk

`func (o *FeeEstimateBucket) GetEstimatedSecondsOk() (*float32, bool)`

GetEstimatedSecondsOk returns a tuple with the EstimatedSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedSeconds

`func (o *FeeEstimateBucket) SetEstimatedSeconds(v float32)`

SetEstimatedSeconds sets EstimatedSeconds field to given value.

### HasEstimatedSeconds

`func (o *FeeEstimateBucket) HasEstimatedSeconds() bool`

HasEstimatedSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


