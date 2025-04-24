# BlockModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | [**EndpointsGetBlocksBlockHeader**](EndpointsGetBlocksBlockHeader.md) |  | 
**Transactions** | Pointer to [**[]BlockTxModel**](BlockTxModel.md) |  | [optional] 
**VerboseData** | [**VerboseDataModel**](VerboseDataModel.md) |  | 
**Extra** | Pointer to [**ExtraModel**](ExtraModel.md) |  | [optional] 

## Methods

### NewBlockModel

`func NewBlockModel(header EndpointsGetBlocksBlockHeader, verboseData VerboseDataModel, ) *BlockModel`

NewBlockModel instantiates a new BlockModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockModelWithDefaults

`func NewBlockModelWithDefaults() *BlockModel`

NewBlockModelWithDefaults instantiates a new BlockModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *BlockModel) GetHeader() EndpointsGetBlocksBlockHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *BlockModel) GetHeaderOk() (*EndpointsGetBlocksBlockHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *BlockModel) SetHeader(v EndpointsGetBlocksBlockHeader)`

SetHeader sets Header field to given value.


### GetTransactions

`func (o *BlockModel) GetTransactions() []BlockTxModel`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *BlockModel) GetTransactionsOk() (*[]BlockTxModel, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *BlockModel) SetTransactions(v []BlockTxModel)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *BlockModel) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetVerboseData

`func (o *BlockModel) GetVerboseData() VerboseDataModel`

GetVerboseData returns the VerboseData field if non-nil, zero value otherwise.

### GetVerboseDataOk

`func (o *BlockModel) GetVerboseDataOk() (*VerboseDataModel, bool)`

GetVerboseDataOk returns a tuple with the VerboseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseData

`func (o *BlockModel) SetVerboseData(v VerboseDataModel)`

SetVerboseData sets VerboseData field to given value.


### GetExtra

`func (o *BlockModel) GetExtra() ExtraModel`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *BlockModel) GetExtraOk() (*ExtraModel, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *BlockModel) SetExtra(v ExtraModel)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *BlockModel) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


