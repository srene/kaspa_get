# VerboseDataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] [default to "18c7afdf8f447ca06adb8b4946dc45f5feb1188c7d177da6094dfbc760eca699"]
**Difficulty** | Pointer to **float32** |  | [optional] [default to [4.10220452325294E12]]
**SelectedParentHash** | Pointer to **string** |  | [optional] [default to "580f65c8da9d436480817f6bd7c13eecd9223b37f0d34ae42fb17e1e9fda397e"]
**TransactionIds** | Pointer to **[]string** |  | [optional] [default to [533f8314bf772259fe517f53507a79ebe61c8c6a11748d93a0835551233b3311]]
**BlueScore** | Pointer to **string** |  | [optional] [default to "18483232"]
**ChildrenHashes** | Pointer to **[]string** |  | [optional] 
**MergeSetBluesHashes** | Pointer to **[]string** |  | [optional] [default to []]
**MergeSetRedsHashes** | Pointer to **[]string** |  | [optional] [default to []]
**IsChainBlock** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewVerboseDataModel

`func NewVerboseDataModel() *VerboseDataModel`

NewVerboseDataModel instantiates a new VerboseDataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerboseDataModelWithDefaults

`func NewVerboseDataModelWithDefaults() *VerboseDataModel`

NewVerboseDataModelWithDefaults instantiates a new VerboseDataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *VerboseDataModel) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *VerboseDataModel) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *VerboseDataModel) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *VerboseDataModel) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetDifficulty

`func (o *VerboseDataModel) GetDifficulty() float32`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *VerboseDataModel) GetDifficultyOk() (*float32, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *VerboseDataModel) SetDifficulty(v float32)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *VerboseDataModel) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetSelectedParentHash

`func (o *VerboseDataModel) GetSelectedParentHash() string`

GetSelectedParentHash returns the SelectedParentHash field if non-nil, zero value otherwise.

### GetSelectedParentHashOk

`func (o *VerboseDataModel) GetSelectedParentHashOk() (*string, bool)`

GetSelectedParentHashOk returns a tuple with the SelectedParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedParentHash

`func (o *VerboseDataModel) SetSelectedParentHash(v string)`

SetSelectedParentHash sets SelectedParentHash field to given value.

### HasSelectedParentHash

`func (o *VerboseDataModel) HasSelectedParentHash() bool`

HasSelectedParentHash returns a boolean if a field has been set.

### GetTransactionIds

`func (o *VerboseDataModel) GetTransactionIds() []string`

GetTransactionIds returns the TransactionIds field if non-nil, zero value otherwise.

### GetTransactionIdsOk

`func (o *VerboseDataModel) GetTransactionIdsOk() (*[]string, bool)`

GetTransactionIdsOk returns a tuple with the TransactionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIds

`func (o *VerboseDataModel) SetTransactionIds(v []string)`

SetTransactionIds sets TransactionIds field to given value.

### HasTransactionIds

`func (o *VerboseDataModel) HasTransactionIds() bool`

HasTransactionIds returns a boolean if a field has been set.

### GetBlueScore

`func (o *VerboseDataModel) GetBlueScore() string`

GetBlueScore returns the BlueScore field if non-nil, zero value otherwise.

### GetBlueScoreOk

`func (o *VerboseDataModel) GetBlueScoreOk() (*string, bool)`

GetBlueScoreOk returns a tuple with the BlueScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueScore

`func (o *VerboseDataModel) SetBlueScore(v string)`

SetBlueScore sets BlueScore field to given value.

### HasBlueScore

`func (o *VerboseDataModel) HasBlueScore() bool`

HasBlueScore returns a boolean if a field has been set.

### GetChildrenHashes

`func (o *VerboseDataModel) GetChildrenHashes() []string`

GetChildrenHashes returns the ChildrenHashes field if non-nil, zero value otherwise.

### GetChildrenHashesOk

`func (o *VerboseDataModel) GetChildrenHashesOk() (*[]string, bool)`

GetChildrenHashesOk returns a tuple with the ChildrenHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenHashes

`func (o *VerboseDataModel) SetChildrenHashes(v []string)`

SetChildrenHashes sets ChildrenHashes field to given value.

### HasChildrenHashes

`func (o *VerboseDataModel) HasChildrenHashes() bool`

HasChildrenHashes returns a boolean if a field has been set.

### GetMergeSetBluesHashes

`func (o *VerboseDataModel) GetMergeSetBluesHashes() []string`

GetMergeSetBluesHashes returns the MergeSetBluesHashes field if non-nil, zero value otherwise.

### GetMergeSetBluesHashesOk

`func (o *VerboseDataModel) GetMergeSetBluesHashesOk() (*[]string, bool)`

GetMergeSetBluesHashesOk returns a tuple with the MergeSetBluesHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeSetBluesHashes

`func (o *VerboseDataModel) SetMergeSetBluesHashes(v []string)`

SetMergeSetBluesHashes sets MergeSetBluesHashes field to given value.

### HasMergeSetBluesHashes

`func (o *VerboseDataModel) HasMergeSetBluesHashes() bool`

HasMergeSetBluesHashes returns a boolean if a field has been set.

### GetMergeSetRedsHashes

`func (o *VerboseDataModel) GetMergeSetRedsHashes() []string`

GetMergeSetRedsHashes returns the MergeSetRedsHashes field if non-nil, zero value otherwise.

### GetMergeSetRedsHashesOk

`func (o *VerboseDataModel) GetMergeSetRedsHashesOk() (*[]string, bool)`

GetMergeSetRedsHashesOk returns a tuple with the MergeSetRedsHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeSetRedsHashes

`func (o *VerboseDataModel) SetMergeSetRedsHashes(v []string)`

SetMergeSetRedsHashes sets MergeSetRedsHashes field to given value.

### HasMergeSetRedsHashes

`func (o *VerboseDataModel) HasMergeSetRedsHashes() bool`

HasMergeSetRedsHashes returns a boolean if a field has been set.

### GetIsChainBlock

`func (o *VerboseDataModel) GetIsChainBlock() bool`

GetIsChainBlock returns the IsChainBlock field if non-nil, zero value otherwise.

### GetIsChainBlockOk

`func (o *VerboseDataModel) GetIsChainBlockOk() (*bool, bool)`

GetIsChainBlockOk returns a tuple with the IsChainBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChainBlock

`func (o *VerboseDataModel) SetIsChainBlock(v bool)`

SetIsChainBlock sets IsChainBlock field to given value.

### HasIsChainBlock

`func (o *VerboseDataModel) HasIsChainBlock() bool`

HasIsChainBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


