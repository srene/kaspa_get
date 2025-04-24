# EndpointsGetBlocksBlockHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** |  | [optional] [default to 1]
**HashMerkleRoot** | Pointer to **string** |  | [optional] [default to "e6641454e16cff4f232b899564eeaa6e480b66069d87bee6a2b2476e63fcd887"]
**AcceptedIdMerkleRoot** | Pointer to **string** |  | [optional] [default to "9bab45b027a0b2b47135b6f6f866e5e4040fc1fdf2fe56eb0c90a603ce86092b"]
**UtxoCommitment** | Pointer to **string** |  | [optional] [default to "236d5f9ffd19b317a97693322c3e2ae11a44b5df803d71f1ccf6c2393bc6143c"]
**Timestamp** | Pointer to **string** |  | [optional] [default to "1656450648874"]
**Bits** | Pointer to **int32** |  | [optional] [default to 455233226]
**Nonce** | Pointer to **string** |  | [optional] [default to "14797571275553019490"]
**DaaScore** | Pointer to **string** |  | [optional] [default to "19984482"]
**BlueWork** | Pointer to **string** |  | [optional] [default to "2d1b3f04f8a0dcd31"]
**Parents** | Pointer to [**[]ParentHashModel**](ParentHashModel.md) |  | [optional] 
**BlueScore** | Pointer to **string** |  | [optional] [default to "18483232"]
**PruningPoint** | Pointer to **string** |  | [optional] [default to "5d32a9403273a34b6551b84340a1459ddde2ae6ba59a47987a6374340ba41d5d"]

## Methods

### NewEndpointsGetBlocksBlockHeader

`func NewEndpointsGetBlocksBlockHeader() *EndpointsGetBlocksBlockHeader`

NewEndpointsGetBlocksBlockHeader instantiates a new EndpointsGetBlocksBlockHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointsGetBlocksBlockHeaderWithDefaults

`func NewEndpointsGetBlocksBlockHeaderWithDefaults() *EndpointsGetBlocksBlockHeader`

NewEndpointsGetBlocksBlockHeaderWithDefaults instantiates a new EndpointsGetBlocksBlockHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *EndpointsGetBlocksBlockHeader) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EndpointsGetBlocksBlockHeader) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EndpointsGetBlocksBlockHeader) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EndpointsGetBlocksBlockHeader) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetHashMerkleRoot

`func (o *EndpointsGetBlocksBlockHeader) GetHashMerkleRoot() string`

GetHashMerkleRoot returns the HashMerkleRoot field if non-nil, zero value otherwise.

### GetHashMerkleRootOk

`func (o *EndpointsGetBlocksBlockHeader) GetHashMerkleRootOk() (*string, bool)`

GetHashMerkleRootOk returns a tuple with the HashMerkleRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashMerkleRoot

`func (o *EndpointsGetBlocksBlockHeader) SetHashMerkleRoot(v string)`

SetHashMerkleRoot sets HashMerkleRoot field to given value.

### HasHashMerkleRoot

`func (o *EndpointsGetBlocksBlockHeader) HasHashMerkleRoot() bool`

HasHashMerkleRoot returns a boolean if a field has been set.

### GetAcceptedIdMerkleRoot

`func (o *EndpointsGetBlocksBlockHeader) GetAcceptedIdMerkleRoot() string`

GetAcceptedIdMerkleRoot returns the AcceptedIdMerkleRoot field if non-nil, zero value otherwise.

### GetAcceptedIdMerkleRootOk

`func (o *EndpointsGetBlocksBlockHeader) GetAcceptedIdMerkleRootOk() (*string, bool)`

GetAcceptedIdMerkleRootOk returns a tuple with the AcceptedIdMerkleRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedIdMerkleRoot

`func (o *EndpointsGetBlocksBlockHeader) SetAcceptedIdMerkleRoot(v string)`

SetAcceptedIdMerkleRoot sets AcceptedIdMerkleRoot field to given value.

### HasAcceptedIdMerkleRoot

`func (o *EndpointsGetBlocksBlockHeader) HasAcceptedIdMerkleRoot() bool`

HasAcceptedIdMerkleRoot returns a boolean if a field has been set.

### GetUtxoCommitment

`func (o *EndpointsGetBlocksBlockHeader) GetUtxoCommitment() string`

GetUtxoCommitment returns the UtxoCommitment field if non-nil, zero value otherwise.

### GetUtxoCommitmentOk

`func (o *EndpointsGetBlocksBlockHeader) GetUtxoCommitmentOk() (*string, bool)`

GetUtxoCommitmentOk returns a tuple with the UtxoCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoCommitment

`func (o *EndpointsGetBlocksBlockHeader) SetUtxoCommitment(v string)`

SetUtxoCommitment sets UtxoCommitment field to given value.

### HasUtxoCommitment

`func (o *EndpointsGetBlocksBlockHeader) HasUtxoCommitment() bool`

HasUtxoCommitment returns a boolean if a field has been set.

### GetTimestamp

`func (o *EndpointsGetBlocksBlockHeader) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EndpointsGetBlocksBlockHeader) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EndpointsGetBlocksBlockHeader) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EndpointsGetBlocksBlockHeader) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBits

`func (o *EndpointsGetBlocksBlockHeader) GetBits() int32`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *EndpointsGetBlocksBlockHeader) GetBitsOk() (*int32, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *EndpointsGetBlocksBlockHeader) SetBits(v int32)`

SetBits sets Bits field to given value.

### HasBits

`func (o *EndpointsGetBlocksBlockHeader) HasBits() bool`

HasBits returns a boolean if a field has been set.

### GetNonce

`func (o *EndpointsGetBlocksBlockHeader) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *EndpointsGetBlocksBlockHeader) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *EndpointsGetBlocksBlockHeader) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *EndpointsGetBlocksBlockHeader) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetDaaScore

`func (o *EndpointsGetBlocksBlockHeader) GetDaaScore() string`

GetDaaScore returns the DaaScore field if non-nil, zero value otherwise.

### GetDaaScoreOk

`func (o *EndpointsGetBlocksBlockHeader) GetDaaScoreOk() (*string, bool)`

GetDaaScoreOk returns a tuple with the DaaScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaaScore

`func (o *EndpointsGetBlocksBlockHeader) SetDaaScore(v string)`

SetDaaScore sets DaaScore field to given value.

### HasDaaScore

`func (o *EndpointsGetBlocksBlockHeader) HasDaaScore() bool`

HasDaaScore returns a boolean if a field has been set.

### GetBlueWork

`func (o *EndpointsGetBlocksBlockHeader) GetBlueWork() string`

GetBlueWork returns the BlueWork field if non-nil, zero value otherwise.

### GetBlueWorkOk

`func (o *EndpointsGetBlocksBlockHeader) GetBlueWorkOk() (*string, bool)`

GetBlueWorkOk returns a tuple with the BlueWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueWork

`func (o *EndpointsGetBlocksBlockHeader) SetBlueWork(v string)`

SetBlueWork sets BlueWork field to given value.

### HasBlueWork

`func (o *EndpointsGetBlocksBlockHeader) HasBlueWork() bool`

HasBlueWork returns a boolean if a field has been set.

### GetParents

`func (o *EndpointsGetBlocksBlockHeader) GetParents() []ParentHashModel`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *EndpointsGetBlocksBlockHeader) GetParentsOk() (*[]ParentHashModel, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *EndpointsGetBlocksBlockHeader) SetParents(v []ParentHashModel)`

SetParents sets Parents field to given value.

### HasParents

`func (o *EndpointsGetBlocksBlockHeader) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetBlueScore

`func (o *EndpointsGetBlocksBlockHeader) GetBlueScore() string`

GetBlueScore returns the BlueScore field if non-nil, zero value otherwise.

### GetBlueScoreOk

`func (o *EndpointsGetBlocksBlockHeader) GetBlueScoreOk() (*string, bool)`

GetBlueScoreOk returns a tuple with the BlueScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueScore

`func (o *EndpointsGetBlocksBlockHeader) SetBlueScore(v string)`

SetBlueScore sets BlueScore field to given value.

### HasBlueScore

`func (o *EndpointsGetBlocksBlockHeader) HasBlueScore() bool`

HasBlueScore returns a boolean if a field has been set.

### GetPruningPoint

`func (o *EndpointsGetBlocksBlockHeader) GetPruningPoint() string`

GetPruningPoint returns the PruningPoint field if non-nil, zero value otherwise.

### GetPruningPointOk

`func (o *EndpointsGetBlocksBlockHeader) GetPruningPointOk() (*string, bool)`

GetPruningPointOk returns a tuple with the PruningPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningPoint

`func (o *EndpointsGetBlocksBlockHeader) SetPruningPoint(v string)`

SetPruningPoint sets PruningPoint field to given value.

### HasPruningPoint

`func (o *EndpointsGetBlocksBlockHeader) HasPruningPoint() bool`

HasPruningPoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


