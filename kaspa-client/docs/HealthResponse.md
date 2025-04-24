# HealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KaspadServers** | [**[]KaspadResponse**](KaspadResponse.md) |  | 
**Database** | [**DBCheckStatus**](DBCheckStatus.md) |  | 

## Methods

### NewHealthResponse

`func NewHealthResponse(kaspadServers []KaspadResponse, database DBCheckStatus, ) *HealthResponse`

NewHealthResponse instantiates a new HealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthResponseWithDefaults

`func NewHealthResponseWithDefaults() *HealthResponse`

NewHealthResponseWithDefaults instantiates a new HealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKaspadServers

`func (o *HealthResponse) GetKaspadServers() []KaspadResponse`

GetKaspadServers returns the KaspadServers field if non-nil, zero value otherwise.

### GetKaspadServersOk

`func (o *HealthResponse) GetKaspadServersOk() (*[]KaspadResponse, bool)`

GetKaspadServersOk returns a tuple with the KaspadServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKaspadServers

`func (o *HealthResponse) SetKaspadServers(v []KaspadResponse)`

SetKaspadServers sets KaspadServers field to given value.


### GetDatabase

`func (o *HealthResponse) GetDatabase() DBCheckStatus`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *HealthResponse) GetDatabaseOk() (*DBCheckStatus, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *HealthResponse) SetDatabase(v DBCheckStatus)`

SetDatabase sets Database field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


