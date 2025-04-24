# BlockTxModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | Pointer to [**[]BlockTxInputModel**](BlockTxInputModel.md) |  | [optional] 
**Outputs** | Pointer to [**[]BlockTxOutputModel**](BlockTxOutputModel.md) |  | [optional] 
**SubnetworkId** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**VerboseData** | [**BlockTxVerboseDataModel**](BlockTxVerboseDataModel.md) |  | 
**LockTime** | Pointer to **int32** |  | [optional] 
**Gas** | Pointer to **int32** |  | [optional] 
**Mass** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewBlockTxModel

`func NewBlockTxModel(verboseData BlockTxVerboseDataModel, ) *BlockTxModel`

NewBlockTxModel instantiates a new BlockTxModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockTxModelWithDefaults

`func NewBlockTxModelWithDefaults() *BlockTxModel`

NewBlockTxModelWithDefaults instantiates a new BlockTxModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputs

`func (o *BlockTxModel) GetInputs() []BlockTxInputModel`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *BlockTxModel) GetInputsOk() (*[]BlockTxInputModel, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *BlockTxModel) SetInputs(v []BlockTxInputModel)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *BlockTxModel) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetOutputs

`func (o *BlockTxModel) GetOutputs() []BlockTxOutputModel`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *BlockTxModel) GetOutputsOk() (*[]BlockTxOutputModel, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *BlockTxModel) SetOutputs(v []BlockTxOutputModel)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *BlockTxModel) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetSubnetworkId

`func (o *BlockTxModel) GetSubnetworkId() string`

GetSubnetworkId returns the SubnetworkId field if non-nil, zero value otherwise.

### GetSubnetworkIdOk

`func (o *BlockTxModel) GetSubnetworkIdOk() (*string, bool)`

GetSubnetworkIdOk returns a tuple with the SubnetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetworkId

`func (o *BlockTxModel) SetSubnetworkId(v string)`

SetSubnetworkId sets SubnetworkId field to given value.

### HasSubnetworkId

`func (o *BlockTxModel) HasSubnetworkId() bool`

HasSubnetworkId returns a boolean if a field has been set.

### GetPayload

`func (o *BlockTxModel) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *BlockTxModel) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *BlockTxModel) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *BlockTxModel) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetVerboseData

`func (o *BlockTxModel) GetVerboseData() BlockTxVerboseDataModel`

GetVerboseData returns the VerboseData field if non-nil, zero value otherwise.

### GetVerboseDataOk

`func (o *BlockTxModel) GetVerboseDataOk() (*BlockTxVerboseDataModel, bool)`

GetVerboseDataOk returns a tuple with the VerboseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseData

`func (o *BlockTxModel) SetVerboseData(v BlockTxVerboseDataModel)`

SetVerboseData sets VerboseData field to given value.


### GetLockTime

`func (o *BlockTxModel) GetLockTime() int32`

GetLockTime returns the LockTime field if non-nil, zero value otherwise.

### GetLockTimeOk

`func (o *BlockTxModel) GetLockTimeOk() (*int32, bool)`

GetLockTimeOk returns a tuple with the LockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTime

`func (o *BlockTxModel) SetLockTime(v int32)`

SetLockTime sets LockTime field to given value.

### HasLockTime

`func (o *BlockTxModel) HasLockTime() bool`

HasLockTime returns a boolean if a field has been set.

### GetGas

`func (o *BlockTxModel) GetGas() int32`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *BlockTxModel) GetGasOk() (*int32, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *BlockTxModel) SetGas(v int32)`

SetGas sets Gas field to given value.

### HasGas

`func (o *BlockTxModel) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetMass

`func (o *BlockTxModel) GetMass() int32`

GetMass returns the Mass field if non-nil, zero value otherwise.

### GetMassOk

`func (o *BlockTxModel) GetMassOk() (*int32, bool)`

GetMassOk returns a tuple with the Mass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMass

`func (o *BlockTxModel) SetMass(v int32)`

SetMass sets Mass field to given value.

### HasMass

`func (o *BlockTxModel) HasMass() bool`

HasMass returns a boolean if a field has been set.

### GetVersion

`func (o *BlockTxModel) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlockTxModel) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlockTxModel) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BlockTxModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


