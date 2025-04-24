# MaxHashrateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hashrate** | Pointer to **float32** |  | [optional] [default to 12000132]
**Blockheader** | [**EndpointsGetHashrateBlockHeader**](EndpointsGetHashrateBlockHeader.md) |  | 

## Methods

### NewMaxHashrateResponse

`func NewMaxHashrateResponse(blockheader EndpointsGetHashrateBlockHeader, ) *MaxHashrateResponse`

NewMaxHashrateResponse instantiates a new MaxHashrateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxHashrateResponseWithDefaults

`func NewMaxHashrateResponseWithDefaults() *MaxHashrateResponse`

NewMaxHashrateResponseWithDefaults instantiates a new MaxHashrateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashrate

`func (o *MaxHashrateResponse) GetHashrate() float32`

GetHashrate returns the Hashrate field if non-nil, zero value otherwise.

### GetHashrateOk

`func (o *MaxHashrateResponse) GetHashrateOk() (*float32, bool)`

GetHashrateOk returns a tuple with the Hashrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashrate

`func (o *MaxHashrateResponse) SetHashrate(v float32)`

SetHashrate sets Hashrate field to given value.

### HasHashrate

`func (o *MaxHashrateResponse) HasHashrate() bool`

HasHashrate returns a boolean if a field has been set.

### GetBlockheader

`func (o *MaxHashrateResponse) GetBlockheader() EndpointsGetHashrateBlockHeader`

GetBlockheader returns the Blockheader field if non-nil, zero value otherwise.

### GetBlockheaderOk

`func (o *MaxHashrateResponse) GetBlockheaderOk() (*EndpointsGetHashrateBlockHeader, bool)`

GetBlockheaderOk returns a tuple with the Blockheader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockheader

`func (o *MaxHashrateResponse) SetBlockheader(v EndpointsGetHashrateBlockHeader)`

SetBlockheader sets Blockheader field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


