# BlockResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHashes** | Pointer to **[]string** |  | [optional] [default to [44edf9bfd32aa154bfad64485882f184372b64bd60565ba121b42fc3cb1238f3, 18c7afdf8f447ca06adb8b4946dc45f5feb1188c7d177da6094dfbc760eca699, 9a822351cd293a653f6721afec1646bd1690da7124b5fbe87001711406010604, 2fda0dad4ec879b4ad02ebb68c757955cab305558998129a7de111ab852e7dcb]]
**Blocks** | Pointer to [**[]BlockModel**](BlockModel.md) |  | [optional] 

## Methods

### NewBlockResponse

`func NewBlockResponse() *BlockResponse`

NewBlockResponse instantiates a new BlockResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockResponseWithDefaults

`func NewBlockResponseWithDefaults() *BlockResponse`

NewBlockResponseWithDefaults instantiates a new BlockResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHashes

`func (o *BlockResponse) GetBlockHashes() []string`

GetBlockHashes returns the BlockHashes field if non-nil, zero value otherwise.

### GetBlockHashesOk

`func (o *BlockResponse) GetBlockHashesOk() (*[]string, bool)`

GetBlockHashesOk returns a tuple with the BlockHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHashes

`func (o *BlockResponse) SetBlockHashes(v []string)`

SetBlockHashes sets BlockHashes field to given value.

### HasBlockHashes

`func (o *BlockResponse) HasBlockHashes() bool`

HasBlockHashes returns a boolean if a field has been set.

### GetBlocks

`func (o *BlockResponse) GetBlocks() []BlockModel`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *BlockResponse) GetBlocksOk() (*[]BlockModel, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *BlockResponse) SetBlocks(v []BlockModel)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *BlockResponse) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


