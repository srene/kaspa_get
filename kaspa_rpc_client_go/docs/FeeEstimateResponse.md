# FeeEstimateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriorityBucket** | [**FeeEstimateBucket**](FeeEstimateBucket.md) |  | 
**NormalBuckets** | [**[]FeeEstimateBucket**](FeeEstimateBucket.md) |  | 
**LowBuckets** | [**[]FeeEstimateBucket**](FeeEstimateBucket.md) |  | 

## Methods

### NewFeeEstimateResponse

`func NewFeeEstimateResponse(priorityBucket FeeEstimateBucket, normalBuckets []FeeEstimateBucket, lowBuckets []FeeEstimateBucket, ) *FeeEstimateResponse`

NewFeeEstimateResponse instantiates a new FeeEstimateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeEstimateResponseWithDefaults

`func NewFeeEstimateResponseWithDefaults() *FeeEstimateResponse`

NewFeeEstimateResponseWithDefaults instantiates a new FeeEstimateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriorityBucket

`func (o *FeeEstimateResponse) GetPriorityBucket() FeeEstimateBucket`

GetPriorityBucket returns the PriorityBucket field if non-nil, zero value otherwise.

### GetPriorityBucketOk

`func (o *FeeEstimateResponse) GetPriorityBucketOk() (*FeeEstimateBucket, bool)`

GetPriorityBucketOk returns a tuple with the PriorityBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityBucket

`func (o *FeeEstimateResponse) SetPriorityBucket(v FeeEstimateBucket)`

SetPriorityBucket sets PriorityBucket field to given value.


### GetNormalBuckets

`func (o *FeeEstimateResponse) GetNormalBuckets() []FeeEstimateBucket`

GetNormalBuckets returns the NormalBuckets field if non-nil, zero value otherwise.

### GetNormalBucketsOk

`func (o *FeeEstimateResponse) GetNormalBucketsOk() (*[]FeeEstimateBucket, bool)`

GetNormalBucketsOk returns a tuple with the NormalBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalBuckets

`func (o *FeeEstimateResponse) SetNormalBuckets(v []FeeEstimateBucket)`

SetNormalBuckets sets NormalBuckets field to given value.


### GetLowBuckets

`func (o *FeeEstimateResponse) GetLowBuckets() []FeeEstimateBucket`

GetLowBuckets returns the LowBuckets field if non-nil, zero value otherwise.

### GetLowBucketsOk

`func (o *FeeEstimateResponse) GetLowBucketsOk() (*[]FeeEstimateBucket, bool)`

GetLowBucketsOk returns a tuple with the LowBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBuckets

`func (o *FeeEstimateResponse) SetLowBuckets(v []FeeEstimateBucket)`

SetLowBuckets sets LowBuckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


