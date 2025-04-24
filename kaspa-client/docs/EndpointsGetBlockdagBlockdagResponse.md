# EndpointsGetBlockdagBlockdagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkName** | Pointer to **string** |  | [optional] [default to "kaspa-mainnet"]
**BlockCount** | Pointer to **string** |  | [optional] [default to "260890"]
**HeaderCount** | Pointer to **string** |  | [optional] [default to "2131312"]
**TipHashes** | Pointer to **[]string** |  | [optional] [default to [78273854a739e3e379dfd34a262bbe922400d8e360e30e3f31228519a334350a]]
**Difficulty** | Pointer to **float32** |  | [optional] [default to 3.8706776777772E12]
**PastMedianTime** | Pointer to **string** |  | [optional] [default to "1656455670700"]
**VirtualParentHashes** | Pointer to **[]string** |  | [optional] [default to [78273854a739e3e379dfd34a262bbe922400d8e360e30e3f31228519a334350a]]
**PruningPointHash** | Pointer to **string** |  | [optional] [default to "[5d32a9403273a34b6551b84340a1459ddde2ae6ba59a47987a6374340ba41d5d]"]
**VirtualDaaScore** | Pointer to **string** |  | [optional] [default to "19989141"]

## Methods

### NewEndpointsGetBlockdagBlockdagResponse

`func NewEndpointsGetBlockdagBlockdagResponse() *EndpointsGetBlockdagBlockdagResponse`

NewEndpointsGetBlockdagBlockdagResponse instantiates a new EndpointsGetBlockdagBlockdagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointsGetBlockdagBlockdagResponseWithDefaults

`func NewEndpointsGetBlockdagBlockdagResponseWithDefaults() *EndpointsGetBlockdagBlockdagResponse`

NewEndpointsGetBlockdagBlockdagResponseWithDefaults instantiates a new EndpointsGetBlockdagBlockdagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkName

`func (o *EndpointsGetBlockdagBlockdagResponse) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *EndpointsGetBlockdagBlockdagResponse) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *EndpointsGetBlockdagBlockdagResponse) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *EndpointsGetBlockdagBlockdagResponse) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetBlockCount

`func (o *EndpointsGetBlockdagBlockdagResponse) GetBlockCount() string`

GetBlockCount returns the BlockCount field if non-nil, zero value otherwise.

### GetBlockCountOk

`func (o *EndpointsGetBlockdagBlockdagResponse) GetBlockCountOk() (*string, bool)`

GetBlockCountOk returns a tuple with the BlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCount

`func (o *EndpointsGetBlockdagBlockdagResponse) SetBlockCount(v string)`

SetBlockCount sets BlockCount field to given value.

### HasBlockCount

`func (o *EndpointsGetBlockdagBlockdagResponse) HasBlockCount() bool`

HasBlockCount returns a boolean if a field has been set.

### GetHeaderCount

`func (o *EndpointsGetBlockdagBlockdagResponse) GetHeaderCount() string`

GetHeaderCount returns the HeaderCount field if non-nil, zero value otherwise.

### GetHeaderCountOk

`func (o *EndpointsGetBlockdagBlockdagResponse) GetHeaderCountOk() (*string, bool)`

GetHeaderCountOk returns a tuple with the HeaderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderCount

`func (o *EndpointsGetBlockdagBlockdagResponse) SetHeaderCount(v string)`

SetHeaderCount sets HeaderCount field to given value.

### HasHeaderCount

`func (o *EndpointsGetBlockdagBlockdagResponse) HasHeaderCount() bool`

HasHeaderCount returns a boolean if a field has been set.

### GetTipHashes

`func (o *EndpointsGetBlockdagBlockdagResponse) GetTipHashes() []string`

GetTipHashes returns the TipHashes field if non-nil, zero value otherwise.

### GetTipHashesOk

`func (o *EndpointsGetBlockdagBlockdagResponse) GetTipHashesOk() (*[]string, bool)`

GetTipHashesOk returns a tuple with the TipHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipHashes

`func (o *EndpointsGetBlockdagBlockdagResponse) SetTipHashes(v []string)`

SetTipHashes sets TipHashes field to given value.

### HasTipHashes

`func (o *EndpointsGetBlockdagBlockdagResponse) HasTipHashes() bool`

HasTipHashes returns a boolean if a field has been set.

### GetDifficulty

`func (o *EndpointsGetBlockdagBlockdagResponse) GetDifficulty() float32`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *EndpointsGetBlockdagBlockdagResponse) GetDifficultyOk() (*float32, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *EndpointsGetBlockdagBlockdagResponse) SetDifficulty(v float32)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *EndpointsGetBlockdagBlockdagResponse) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetPastMedianTime

`func (o *EndpointsGetBlockdagBlockdagResponse) GetPastMedianTime() string`

GetPastMedianTime returns the PastMedianTime field if non-nil, zero value otherwise.

### GetPastMedianTimeOk

`func (o *EndpointsGetBlockdagBlockdagResponse) GetPastMedianTimeOk() (*string, bool)`

GetPastMedianTimeOk returns a tuple with the PastMedianTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastMedianTime

`func (o *EndpointsGetBlockdagBlockdagResponse) SetPastMedianTime(v string)`

SetPastMedianTime sets PastMedianTime field to given value.

### HasPastMedianTime

`func (o *EndpointsGetBlockdagBlockdagResponse) HasPastMedianTime() bool`

HasPastMedianTime returns a boolean if a field has been set.

### GetVirtualParentHashes

`func (o *EndpointsGetBlockdagBlockdagResponse) GetVirtualParentHashes() []string`

GetVirtualParentHashes returns the VirtualParentHashes field if non-nil, zero value otherwise.

### GetVirtualParentHashesOk

`func (o *EndpointsGetBlockdagBlockdagResponse) GetVirtualParentHashesOk() (*[]string, bool)`

GetVirtualParentHashesOk returns a tuple with the VirtualParentHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualParentHashes

`func (o *EndpointsGetBlockdagBlockdagResponse) SetVirtualParentHashes(v []string)`

SetVirtualParentHashes sets VirtualParentHashes field to given value.

### HasVirtualParentHashes

`func (o *EndpointsGetBlockdagBlockdagResponse) HasVirtualParentHashes() bool`

HasVirtualParentHashes returns a boolean if a field has been set.

### GetPruningPointHash

`func (o *EndpointsGetBlockdagBlockdagResponse) GetPruningPointHash() string`

GetPruningPointHash returns the PruningPointHash field if non-nil, zero value otherwise.

### GetPruningPointHashOk

`func (o *EndpointsGetBlockdagBlockdagResponse) GetPruningPointHashOk() (*string, bool)`

GetPruningPointHashOk returns a tuple with the PruningPointHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningPointHash

`func (o *EndpointsGetBlockdagBlockdagResponse) SetPruningPointHash(v string)`

SetPruningPointHash sets PruningPointHash field to given value.

### HasPruningPointHash

`func (o *EndpointsGetBlockdagBlockdagResponse) HasPruningPointHash() bool`

HasPruningPointHash returns a boolean if a field has been set.

### GetVirtualDaaScore

`func (o *EndpointsGetBlockdagBlockdagResponse) GetVirtualDaaScore() string`

GetVirtualDaaScore returns the VirtualDaaScore field if non-nil, zero value otherwise.

### GetVirtualDaaScoreOk

`func (o *EndpointsGetBlockdagBlockdagResponse) GetVirtualDaaScoreOk() (*string, bool)`

GetVirtualDaaScoreOk returns a tuple with the VirtualDaaScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDaaScore

`func (o *EndpointsGetBlockdagBlockdagResponse) SetVirtualDaaScore(v string)`

SetVirtualDaaScore sets VirtualDaaScore field to given value.

### HasVirtualDaaScore

`func (o *EndpointsGetBlockdagBlockdagResponse) HasVirtualDaaScore() bool`

HasVirtualDaaScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


