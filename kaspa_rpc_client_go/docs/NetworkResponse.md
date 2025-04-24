# NetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkName** | Pointer to **string** |  | [optional] [default to "kaspa-mainnet"]
**BlockCount** | Pointer to **string** |  | [optional] [default to "261357"]
**HeaderCount** | Pointer to **string** |  | [optional] [default to "23138783"]
**TipHashes** | Pointer to **[]string** |  | [optional] [default to [efdbe104c6275cf881583fba77834c8528fd1ab059f6b4737c42564d0d9fedbc, 6affbe62baef0f1a562f166b9857844b03b51a8ec9b8417ceb308d53fdc239a2]]
**Difficulty** | Pointer to **float32** |  | [optional] [default to 3.88707990501409E12]
**PastMedianTime** | Pointer to **string** |  | [optional] [default to "1656456088196"]
**VirtualParentHashes** | Pointer to **[]string** |  | [optional] [default to [6affbe62baef0f1a562f166b9857844b03b51a8ec9b8417ceb308d53fdc239a2, efdbe104c6275cf881583fba77834c8528fd1ab059f6b4737c42564d0d9fedbc]]
**PruningPointHash** | Pointer to **string** |  | [optional] [default to "5d32a9403273a34b6551b84340a1459ddde2ae6ba59a47987a6374340ba41d5d"]
**VirtualDaaScore** | Pointer to **string** |  | [optional] [default to "19989984"]

## Methods

### NewNetworkResponse

`func NewNetworkResponse() *NetworkResponse`

NewNetworkResponse instantiates a new NetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkResponseWithDefaults

`func NewNetworkResponseWithDefaults() *NetworkResponse`

NewNetworkResponseWithDefaults instantiates a new NetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkName

`func (o *NetworkResponse) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *NetworkResponse) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *NetworkResponse) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *NetworkResponse) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetBlockCount

`func (o *NetworkResponse) GetBlockCount() string`

GetBlockCount returns the BlockCount field if non-nil, zero value otherwise.

### GetBlockCountOk

`func (o *NetworkResponse) GetBlockCountOk() (*string, bool)`

GetBlockCountOk returns a tuple with the BlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCount

`func (o *NetworkResponse) SetBlockCount(v string)`

SetBlockCount sets BlockCount field to given value.

### HasBlockCount

`func (o *NetworkResponse) HasBlockCount() bool`

HasBlockCount returns a boolean if a field has been set.

### GetHeaderCount

`func (o *NetworkResponse) GetHeaderCount() string`

GetHeaderCount returns the HeaderCount field if non-nil, zero value otherwise.

### GetHeaderCountOk

`func (o *NetworkResponse) GetHeaderCountOk() (*string, bool)`

GetHeaderCountOk returns a tuple with the HeaderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderCount

`func (o *NetworkResponse) SetHeaderCount(v string)`

SetHeaderCount sets HeaderCount field to given value.

### HasHeaderCount

`func (o *NetworkResponse) HasHeaderCount() bool`

HasHeaderCount returns a boolean if a field has been set.

### GetTipHashes

`func (o *NetworkResponse) GetTipHashes() []string`

GetTipHashes returns the TipHashes field if non-nil, zero value otherwise.

### GetTipHashesOk

`func (o *NetworkResponse) GetTipHashesOk() (*[]string, bool)`

GetTipHashesOk returns a tuple with the TipHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipHashes

`func (o *NetworkResponse) SetTipHashes(v []string)`

SetTipHashes sets TipHashes field to given value.

### HasTipHashes

`func (o *NetworkResponse) HasTipHashes() bool`

HasTipHashes returns a boolean if a field has been set.

### GetDifficulty

`func (o *NetworkResponse) GetDifficulty() float32`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *NetworkResponse) GetDifficultyOk() (*float32, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *NetworkResponse) SetDifficulty(v float32)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *NetworkResponse) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetPastMedianTime

`func (o *NetworkResponse) GetPastMedianTime() string`

GetPastMedianTime returns the PastMedianTime field if non-nil, zero value otherwise.

### GetPastMedianTimeOk

`func (o *NetworkResponse) GetPastMedianTimeOk() (*string, bool)`

GetPastMedianTimeOk returns a tuple with the PastMedianTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastMedianTime

`func (o *NetworkResponse) SetPastMedianTime(v string)`

SetPastMedianTime sets PastMedianTime field to given value.

### HasPastMedianTime

`func (o *NetworkResponse) HasPastMedianTime() bool`

HasPastMedianTime returns a boolean if a field has been set.

### GetVirtualParentHashes

`func (o *NetworkResponse) GetVirtualParentHashes() []string`

GetVirtualParentHashes returns the VirtualParentHashes field if non-nil, zero value otherwise.

### GetVirtualParentHashesOk

`func (o *NetworkResponse) GetVirtualParentHashesOk() (*[]string, bool)`

GetVirtualParentHashesOk returns a tuple with the VirtualParentHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualParentHashes

`func (o *NetworkResponse) SetVirtualParentHashes(v []string)`

SetVirtualParentHashes sets VirtualParentHashes field to given value.

### HasVirtualParentHashes

`func (o *NetworkResponse) HasVirtualParentHashes() bool`

HasVirtualParentHashes returns a boolean if a field has been set.

### GetPruningPointHash

`func (o *NetworkResponse) GetPruningPointHash() string`

GetPruningPointHash returns the PruningPointHash field if non-nil, zero value otherwise.

### GetPruningPointHashOk

`func (o *NetworkResponse) GetPruningPointHashOk() (*string, bool)`

GetPruningPointHashOk returns a tuple with the PruningPointHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningPointHash

`func (o *NetworkResponse) SetPruningPointHash(v string)`

SetPruningPointHash sets PruningPointHash field to given value.

### HasPruningPointHash

`func (o *NetworkResponse) HasPruningPointHash() bool`

HasPruningPointHash returns a boolean if a field has been set.

### GetVirtualDaaScore

`func (o *NetworkResponse) GetVirtualDaaScore() string`

GetVirtualDaaScore returns the VirtualDaaScore field if non-nil, zero value otherwise.

### GetVirtualDaaScoreOk

`func (o *NetworkResponse) GetVirtualDaaScoreOk() (*string, bool)`

GetVirtualDaaScoreOk returns a tuple with the VirtualDaaScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDaaScore

`func (o *NetworkResponse) SetVirtualDaaScore(v string)`

SetVirtualDaaScore sets VirtualDaaScore field to given value.

### HasVirtualDaaScore

`func (o *NetworkResponse) HasVirtualDaaScore() bool`

HasVirtualDaaScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


