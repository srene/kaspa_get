# SubmitTxOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** |  | 
**ScriptPublicKey** | [**SubmitTxScriptPublicKey**](SubmitTxScriptPublicKey.md) |  | 

## Methods

### NewSubmitTxOutput

`func NewSubmitTxOutput(amount int32, scriptPublicKey SubmitTxScriptPublicKey, ) *SubmitTxOutput`

NewSubmitTxOutput instantiates a new SubmitTxOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitTxOutputWithDefaults

`func NewSubmitTxOutputWithDefaults() *SubmitTxOutput`

NewSubmitTxOutputWithDefaults instantiates a new SubmitTxOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SubmitTxOutput) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SubmitTxOutput) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SubmitTxOutput) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetScriptPublicKey

`func (o *SubmitTxOutput) GetScriptPublicKey() SubmitTxScriptPublicKey`

GetScriptPublicKey returns the ScriptPublicKey field if non-nil, zero value otherwise.

### GetScriptPublicKeyOk

`func (o *SubmitTxOutput) GetScriptPublicKeyOk() (*SubmitTxScriptPublicKey, bool)`

GetScriptPublicKeyOk returns a tuple with the ScriptPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPublicKey

`func (o *SubmitTxOutput) SetScriptPublicKey(v SubmitTxScriptPublicKey)`

SetScriptPublicKey sets ScriptPublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


