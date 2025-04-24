# SubmitTxModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  | 
**Inputs** | [**[]SubmitTxInput**](SubmitTxInput.md) |  | 
**Outputs** | [**[]SubmitTxOutput**](SubmitTxOutput.md) |  | 
**LockTime** | Pointer to **int32** |  | [optional] [default to 0]
**SubnetworkId** | Pointer to **string** |  | [optional] 

## Methods

### NewSubmitTxModel

`func NewSubmitTxModel(version int32, inputs []SubmitTxInput, outputs []SubmitTxOutput, ) *SubmitTxModel`

NewSubmitTxModel instantiates a new SubmitTxModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitTxModelWithDefaults

`func NewSubmitTxModelWithDefaults() *SubmitTxModel`

NewSubmitTxModelWithDefaults instantiates a new SubmitTxModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SubmitTxModel) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SubmitTxModel) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SubmitTxModel) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetInputs

`func (o *SubmitTxModel) GetInputs() []SubmitTxInput`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *SubmitTxModel) GetInputsOk() (*[]SubmitTxInput, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *SubmitTxModel) SetInputs(v []SubmitTxInput)`

SetInputs sets Inputs field to given value.


### GetOutputs

`func (o *SubmitTxModel) GetOutputs() []SubmitTxOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *SubmitTxModel) GetOutputsOk() (*[]SubmitTxOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *SubmitTxModel) SetOutputs(v []SubmitTxOutput)`

SetOutputs sets Outputs field to given value.


### GetLockTime

`func (o *SubmitTxModel) GetLockTime() int32`

GetLockTime returns the LockTime field if non-nil, zero value otherwise.

### GetLockTimeOk

`func (o *SubmitTxModel) GetLockTimeOk() (*int32, bool)`

GetLockTimeOk returns a tuple with the LockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTime

`func (o *SubmitTxModel) SetLockTime(v int32)`

SetLockTime sets LockTime field to given value.

### HasLockTime

`func (o *SubmitTxModel) HasLockTime() bool`

HasLockTime returns a boolean if a field has been set.

### GetSubnetworkId

`func (o *SubmitTxModel) GetSubnetworkId() string`

GetSubnetworkId returns the SubnetworkId field if non-nil, zero value otherwise.

### GetSubnetworkIdOk

`func (o *SubmitTxModel) GetSubnetworkIdOk() (*string, bool)`

GetSubnetworkIdOk returns a tuple with the SubnetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetworkId

`func (o *SubmitTxModel) SetSubnetworkId(v string)`

SetSubnetworkId sets SubnetworkId field to given value.

### HasSubnetworkId

`func (o *SubmitTxModel) HasSubnetworkId() bool`

HasSubnetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


