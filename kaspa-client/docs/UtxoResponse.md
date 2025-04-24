# UtxoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] [default to "kaspatest:qpqz2vxj23kvh0m73ta2jjn2u4cv4tlufqns2eap8mxyyt0rvrxy6ejkful67"]
**Outpoint** | [**OutpointModel**](OutpointModel.md) |  | 
**UtxoEntry** | [**UtxoModel**](UtxoModel.md) |  | 

## Methods

### NewUtxoResponse

`func NewUtxoResponse(outpoint OutpointModel, utxoEntry UtxoModel, ) *UtxoResponse`

NewUtxoResponse instantiates a new UtxoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtxoResponseWithDefaults

`func NewUtxoResponseWithDefaults() *UtxoResponse`

NewUtxoResponseWithDefaults instantiates a new UtxoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UtxoResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UtxoResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UtxoResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UtxoResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetOutpoint

`func (o *UtxoResponse) GetOutpoint() OutpointModel`

GetOutpoint returns the Outpoint field if non-nil, zero value otherwise.

### GetOutpointOk

`func (o *UtxoResponse) GetOutpointOk() (*OutpointModel, bool)`

GetOutpointOk returns a tuple with the Outpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutpoint

`func (o *UtxoResponse) SetOutpoint(v OutpointModel)`

SetOutpoint sets Outpoint field to given value.


### GetUtxoEntry

`func (o *UtxoResponse) GetUtxoEntry() UtxoModel`

GetUtxoEntry returns the UtxoEntry field if non-nil, zero value otherwise.

### GetUtxoEntryOk

`func (o *UtxoResponse) GetUtxoEntryOk() (*UtxoModel, bool)`

GetUtxoEntryOk returns a tuple with the UtxoEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoEntry

`func (o *UtxoResponse) SetUtxoEntry(v UtxoModel)`

SetUtxoEntry sets UtxoEntry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


