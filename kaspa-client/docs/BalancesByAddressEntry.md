# BalancesByAddressEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] [default to "kaspatest:qpqz2vxj23kvh0m73ta2jjn2u4cv4tlufqns2eap8mxyyt0rvrxy6ejkful67"]
**Balance** | Pointer to **int32** |  | [optional] [default to 12451591699]

## Methods

### NewBalancesByAddressEntry

`func NewBalancesByAddressEntry() *BalancesByAddressEntry`

NewBalancesByAddressEntry instantiates a new BalancesByAddressEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalancesByAddressEntryWithDefaults

`func NewBalancesByAddressEntryWithDefaults() *BalancesByAddressEntry`

NewBalancesByAddressEntryWithDefaults instantiates a new BalancesByAddressEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BalancesByAddressEntry) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BalancesByAddressEntry) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BalancesByAddressEntry) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BalancesByAddressEntry) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBalance

`func (o *BalancesByAddressEntry) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BalancesByAddressEntry) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BalancesByAddressEntry) SetBalance(v int32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *BalancesByAddressEntry) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


