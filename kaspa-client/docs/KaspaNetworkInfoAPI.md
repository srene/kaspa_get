# \KaspaNetworkInfoAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlockdagInfoBlockdagGet**](KaspaNetworkInfoAPI.md#GetBlockdagInfoBlockdagGet) | **Get** /info/blockdag | Get Blockdag
[**GetBlockrewardInfoBlockrewardGet**](KaspaNetworkInfoAPI.md#GetBlockrewardInfoBlockrewardGet) | **Get** /info/blockreward | Get Blockreward
[**GetCirculatingCoinsInfoCoinsupplyCirculatingGet**](KaspaNetworkInfoAPI.md#GetCirculatingCoinsInfoCoinsupplyCirculatingGet) | **Get** /info/coinsupply/circulating | Get Circulating Coins
[**GetCoinsupplyInfoCoinsupplyGet**](KaspaNetworkInfoAPI.md#GetCoinsupplyInfoCoinsupplyGet) | **Get** /info/coinsupply | Get Coinsupply
[**GetFeeEstimateInfoFeeEstimateGet**](KaspaNetworkInfoAPI.md#GetFeeEstimateInfoFeeEstimateGet) | **Get** /info/fee-estimate | Get Fee Estimate
[**GetHalvingInfoHalvingGet**](KaspaNetworkInfoAPI.md#GetHalvingInfoHalvingGet) | **Get** /info/halving | Get Halving
[**GetHashrateInfoHashrateGet**](KaspaNetworkInfoAPI.md#GetHashrateInfoHashrateGet) | **Get** /info/hashrate | Get Hashrate
[**GetKaspadInfoInfoKaspadGet**](KaspaNetworkInfoAPI.md#GetKaspadInfoInfoKaspadGet) | **Get** /info/kaspad | Get Kaspad Info
[**GetMarketcapInfoMarketcapGet**](KaspaNetworkInfoAPI.md#GetMarketcapInfoMarketcapGet) | **Get** /info/marketcap | Get Marketcap
[**GetMaxHashrateInfoHashrateMaxGet**](KaspaNetworkInfoAPI.md#GetMaxHashrateInfoHashrateMaxGet) | **Get** /info/hashrate/max | Get Max Hashrate
[**GetNetworkInfoNetworkGet**](KaspaNetworkInfoAPI.md#GetNetworkInfoNetworkGet) | **Get** /info/network | Get Network
[**GetPriceInfoPriceGet**](KaspaNetworkInfoAPI.md#GetPriceInfoPriceGet) | **Get** /info/price | Get Price
[**GetTotalCoinsInfoCoinsupplyTotalGet**](KaspaNetworkInfoAPI.md#GetTotalCoinsInfoCoinsupplyTotalGet) | **Get** /info/coinsupply/total | Get Total Coins
[**GetVirtualSelectedParentBlueScoreInfoVirtualChainBlueScoreGet**](KaspaNetworkInfoAPI.md#GetVirtualSelectedParentBlueScoreInfoVirtualChainBlueScoreGet) | **Get** /info/virtual-chain-blue-score | Get Virtual Selected Parent Blue Score
[**HealthStateInfoHealthGet**](KaspaNetworkInfoAPI.md#HealthStateInfoHealthGet) | **Get** /info/health | Health State



## GetBlockdagInfoBlockdagGet

> EndpointsGetBlockdagBlockdagResponse GetBlockdagInfoBlockdagGet(ctx).Execute()

Get Blockdag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.GetBlockdagInfoBlockdagGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.GetBlockdagInfoBlockdagGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockdagInfoBlockdagGet`: EndpointsGetBlockdagBlockdagResponse
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.GetBlockdagInfoBlockdagGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockdagInfoBlockdagGetRequest struct via the builder pattern


### Return type

[**EndpointsGetBlockdagBlockdagResponse**](EndpointsGetBlockdagBlockdagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockrewardInfoBlockrewardGet

> ResponseGetBlockrewardInfoBlockrewardGet GetBlockrewardInfoBlockrewardGet(ctx).StringOnly(stringOnly).Execute()

Get Blockreward



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	stringOnly := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.GetBlockrewardInfoBlockrewardGet(context.Background()).StringOnly(stringOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.GetBlockrewardInfoBlockrewardGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockrewardInfoBlockrewardGet`: ResponseGetBlockrewardInfoBlockrewardGet
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.GetBlockrewardInfoBlockrewardGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockrewardInfoBlockrewardGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stringOnly** | **bool** |  | [default to false]

### Return type

[**ResponseGetBlockrewardInfoBlockrewardGet**](ResponseGetBlockrewardInfoBlockrewardGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCirculatingCoinsInfoCoinsupplyCirculatingGet

> string GetCirculatingCoinsInfoCoinsupplyCirculatingGet(ctx).InBillion(inBillion).Execute()

Get Circulating Coins



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	inBillion := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.GetCirculatingCoinsInfoCoinsupplyCirculatingGet(context.Background()).InBillion(inBillion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.GetCirculatingCoinsInfoCoinsupplyCirculatingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCirculatingCoinsInfoCoinsupplyCirculatingGet`: string
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.GetCirculatingCoinsInfoCoinsupplyCirculatingGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCirculatingCoinsInfoCoinsupplyCirculatingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inBillion** | **bool** |  | [default to false]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCoinsupplyInfoCoinsupplyGet

> CoinSupplyResponse GetCoinsupplyInfoCoinsupplyGet(ctx).Execute()

Get Coinsupply



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.GetCoinsupplyInfoCoinsupplyGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.GetCoinsupplyInfoCoinsupplyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCoinsupplyInfoCoinsupplyGet`: CoinSupplyResponse
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.GetCoinsupplyInfoCoinsupplyGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCoinsupplyInfoCoinsupplyGetRequest struct via the builder pattern


### Return type

[**CoinSupplyResponse**](CoinSupplyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeeEstimateInfoFeeEstimateGet

> FeeEstimateResponse GetFeeEstimateInfoFeeEstimateGet(ctx).Execute()

Get Fee Estimate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.GetFeeEstimateInfoFeeEstimateGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.GetFeeEstimateInfoFeeEstimateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeeEstimateInfoFeeEstimateGet`: FeeEstimateResponse
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.GetFeeEstimateInfoFeeEstimateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeeEstimateInfoFeeEstimateGetRequest struct via the builder pattern


### Return type

[**FeeEstimateResponse**](FeeEstimateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHalvingInfoHalvingGet

> ResponseGetHalvingInfoHalvingGet GetHalvingInfoHalvingGet(ctx).Field(field).Execute()

Get Halving



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	field := "field_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.GetHalvingInfoHalvingGet(context.Background()).Field(field).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.GetHalvingInfoHalvingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHalvingInfoHalvingGet`: ResponseGetHalvingInfoHalvingGet
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.GetHalvingInfoHalvingGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHalvingInfoHalvingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **field** | **string** |  | 

### Return type

[**ResponseGetHalvingInfoHalvingGet**](ResponseGetHalvingInfoHalvingGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHashrateInfoHashrateGet

> ResponseGetHashrateInfoHashrateGet GetHashrateInfoHashrateGet(ctx).StringOnly(stringOnly).Execute()

Get Hashrate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	stringOnly := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.GetHashrateInfoHashrateGet(context.Background()).StringOnly(stringOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.GetHashrateInfoHashrateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHashrateInfoHashrateGet`: ResponseGetHashrateInfoHashrateGet
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.GetHashrateInfoHashrateGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHashrateInfoHashrateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stringOnly** | **bool** |  | [default to false]

### Return type

[**ResponseGetHashrateInfoHashrateGet**](ResponseGetHashrateInfoHashrateGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKaspadInfoInfoKaspadGet

> KaspadInfoResponse GetKaspadInfoInfoKaspadGet(ctx).Execute()

Get Kaspad Info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.GetKaspadInfoInfoKaspadGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.GetKaspadInfoInfoKaspadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKaspadInfoInfoKaspadGet`: KaspadInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.GetKaspadInfoInfoKaspadGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKaspadInfoInfoKaspadGetRequest struct via the builder pattern


### Return type

[**KaspadInfoResponse**](KaspadInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketcapInfoMarketcapGet

> ResponseGetMarketcapInfoMarketcapGet GetMarketcapInfoMarketcapGet(ctx).StringOnly(stringOnly).Execute()

Get Marketcap



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	stringOnly := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.GetMarketcapInfoMarketcapGet(context.Background()).StringOnly(stringOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.GetMarketcapInfoMarketcapGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketcapInfoMarketcapGet`: ResponseGetMarketcapInfoMarketcapGet
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.GetMarketcapInfoMarketcapGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketcapInfoMarketcapGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stringOnly** | **bool** |  | [default to false]

### Return type

[**ResponseGetMarketcapInfoMarketcapGet**](ResponseGetMarketcapInfoMarketcapGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxHashrateInfoHashrateMaxGet

> MaxHashrateResponse GetMaxHashrateInfoHashrateMaxGet(ctx).Execute()

Get Max Hashrate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.GetMaxHashrateInfoHashrateMaxGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.GetMaxHashrateInfoHashrateMaxGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaxHashrateInfoHashrateMaxGet`: MaxHashrateResponse
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.GetMaxHashrateInfoHashrateMaxGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxHashrateInfoHashrateMaxGetRequest struct via the builder pattern


### Return type

[**MaxHashrateResponse**](MaxHashrateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkInfoNetworkGet

> NetworkResponse GetNetworkInfoNetworkGet(ctx).Execute()

Get Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.GetNetworkInfoNetworkGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.GetNetworkInfoNetworkGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkInfoNetworkGet`: NetworkResponse
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.GetNetworkInfoNetworkGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInfoNetworkGetRequest struct via the builder pattern


### Return type

[**NetworkResponse**](NetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceInfoPriceGet

> ResponseGetPriceInfoPriceGet GetPriceInfoPriceGet(ctx).StringOnly(stringOnly).Execute()

Get Price



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	stringOnly := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.GetPriceInfoPriceGet(context.Background()).StringOnly(stringOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.GetPriceInfoPriceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceInfoPriceGet`: ResponseGetPriceInfoPriceGet
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.GetPriceInfoPriceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceInfoPriceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stringOnly** | **bool** |  | [default to false]

### Return type

[**ResponseGetPriceInfoPriceGet**](ResponseGetPriceInfoPriceGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTotalCoinsInfoCoinsupplyTotalGet

> string GetTotalCoinsInfoCoinsupplyTotalGet(ctx).Execute()

Get Total Coins



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.GetTotalCoinsInfoCoinsupplyTotalGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.GetTotalCoinsInfoCoinsupplyTotalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTotalCoinsInfoCoinsupplyTotalGet`: string
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.GetTotalCoinsInfoCoinsupplyTotalGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTotalCoinsInfoCoinsupplyTotalGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualSelectedParentBlueScoreInfoVirtualChainBlueScoreGet

> EndpointsGetVirtualChainBlueScoreBlockdagResponse GetVirtualSelectedParentBlueScoreInfoVirtualChainBlueScoreGet(ctx).Execute()

Get Virtual Selected Parent Blue Score



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.GetVirtualSelectedParentBlueScoreInfoVirtualChainBlueScoreGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.GetVirtualSelectedParentBlueScoreInfoVirtualChainBlueScoreGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualSelectedParentBlueScoreInfoVirtualChainBlueScoreGet`: EndpointsGetVirtualChainBlueScoreBlockdagResponse
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.GetVirtualSelectedParentBlueScoreInfoVirtualChainBlueScoreGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualSelectedParentBlueScoreInfoVirtualChainBlueScoreGetRequest struct via the builder pattern


### Return type

[**EndpointsGetVirtualChainBlueScoreBlockdagResponse**](EndpointsGetVirtualChainBlueScoreBlockdagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthStateInfoHealthGet

> HealthResponse HealthStateInfoHealthGet(ctx).Execute()

Health State



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaNetworkInfoAPI.HealthStateInfoHealthGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaNetworkInfoAPI.HealthStateInfoHealthGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthStateInfoHealthGet`: HealthResponse
	fmt.Fprintf(os.Stdout, "Response from `KaspaNetworkInfoAPI.HealthStateInfoHealthGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthStateInfoHealthGetRequest struct via the builder pattern


### Return type

[**HealthResponse**](HealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

