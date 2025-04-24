# KaspadInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MempoolSize** | Pointer to **string** |  | [optional] [default to "1"]
**ServerVersion** | Pointer to **string** |  | [optional] [default to "0.12.2"]
**IsUtxoIndexed** | Pointer to **bool** |  | [optional] [default to true]
**IsSynced** | Pointer to **bool** |  | [optional] [default to true]
**P2pIdHashed** | Pointer to **string** |  | [optional] [default to "36a17cd8644eef34fc7fe4719655e06dbdf117008900c46975e66c35acd09b01"]

## Methods

### NewKaspadInfoResponse

`func NewKaspadInfoResponse() *KaspadInfoResponse`

NewKaspadInfoResponse instantiates a new KaspadInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKaspadInfoResponseWithDefaults

`func NewKaspadInfoResponseWithDefaults() *KaspadInfoResponse`

NewKaspadInfoResponseWithDefaults instantiates a new KaspadInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMempoolSize

`func (o *KaspadInfoResponse) GetMempoolSize() string`

GetMempoolSize returns the MempoolSize field if non-nil, zero value otherwise.

### GetMempoolSizeOk

`func (o *KaspadInfoResponse) GetMempoolSizeOk() (*string, bool)`

GetMempoolSizeOk returns a tuple with the MempoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMempoolSize

`func (o *KaspadInfoResponse) SetMempoolSize(v string)`

SetMempoolSize sets MempoolSize field to given value.

### HasMempoolSize

`func (o *KaspadInfoResponse) HasMempoolSize() bool`

HasMempoolSize returns a boolean if a field has been set.

### GetServerVersion

`func (o *KaspadInfoResponse) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *KaspadInfoResponse) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *KaspadInfoResponse) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *KaspadInfoResponse) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetIsUtxoIndexed

`func (o *KaspadInfoResponse) GetIsUtxoIndexed() bool`

GetIsUtxoIndexed returns the IsUtxoIndexed field if non-nil, zero value otherwise.

### GetIsUtxoIndexedOk

`func (o *KaspadInfoResponse) GetIsUtxoIndexedOk() (*bool, bool)`

GetIsUtxoIndexedOk returns a tuple with the IsUtxoIndexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUtxoIndexed

`func (o *KaspadInfoResponse) SetIsUtxoIndexed(v bool)`

SetIsUtxoIndexed sets IsUtxoIndexed field to given value.

### HasIsUtxoIndexed

`func (o *KaspadInfoResponse) HasIsUtxoIndexed() bool`

HasIsUtxoIndexed returns a boolean if a field has been set.

### GetIsSynced

`func (o *KaspadInfoResponse) GetIsSynced() bool`

GetIsSynced returns the IsSynced field if non-nil, zero value otherwise.

### GetIsSyncedOk

`func (o *KaspadInfoResponse) GetIsSyncedOk() (*bool, bool)`

GetIsSyncedOk returns a tuple with the IsSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynced

`func (o *KaspadInfoResponse) SetIsSynced(v bool)`

SetIsSynced sets IsSynced field to given value.

### HasIsSynced

`func (o *KaspadInfoResponse) HasIsSynced() bool`

HasIsSynced returns a boolean if a field has been set.

### GetP2pIdHashed

`func (o *KaspadInfoResponse) GetP2pIdHashed() string`

GetP2pIdHashed returns the P2pIdHashed field if non-nil, zero value otherwise.

### GetP2pIdHashedOk

`func (o *KaspadInfoResponse) GetP2pIdHashedOk() (*string, bool)`

GetP2pIdHashedOk returns a tuple with the P2pIdHashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2pIdHashed

`func (o *KaspadInfoResponse) SetP2pIdHashed(v string)`

SetP2pIdHashed sets P2pIdHashed field to given value.

### HasP2pIdHashed

`func (o *KaspadInfoResponse) HasP2pIdHashed() bool`

HasP2pIdHashed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


