# \KaspaAddressesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddressesActiveAddressesActivePost**](KaspaAddressesAPI.md#GetAddressesActiveAddressesActivePost) | **Post** /addresses/active | Get Addresses Active
[**GetBalanceFromKaspaAddressAddressesKaspaAddressBalanceGet**](KaspaAddressesAPI.md#GetBalanceFromKaspaAddressAddressesKaspaAddressBalanceGet) | **Get** /addresses/{kaspaAddress}/balance | Get Balance From Kaspa Address
[**GetBalancesFromKaspaAddressesAddressesBalancesPost**](KaspaAddressesAPI.md#GetBalancesFromKaspaAddressesAddressesBalancesPost) | **Post** /addresses/balances | Get Balances From Kaspa Addresses
[**GetFullTransactionsForAddressAddressesKaspaAddressFullTransactionsGet**](KaspaAddressesAPI.md#GetFullTransactionsForAddressAddressesKaspaAddressFullTransactionsGet) | **Get** /addresses/{kaspaAddress}/full-transactions | Get Full Transactions For Address
[**GetFullTransactionsForAddressPageAddressesKaspaAddressFullTransactionsPageGet**](KaspaAddressesAPI.md#GetFullTransactionsForAddressPageAddressesKaspaAddressFullTransactionsPageGet) | **Get** /addresses/{kaspaAddress}/full-transactions-page | Get Full Transactions For Address Page
[**GetNameForAddressAddressesKaspaAddressNameGet**](KaspaAddressesAPI.md#GetNameForAddressAddressesKaspaAddressNameGet) | **Get** /addresses/{kaspaAddress}/name | Get Name For Address
[**GetTransactionCountForAddressAddressesKaspaAddressTransactionsCountGet**](KaspaAddressesAPI.md#GetTransactionCountForAddressAddressesKaspaAddressTransactionsCountGet) | **Get** /addresses/{kaspaAddress}/transactions-count | Get Transaction Count For Address
[**GetUtxosForAddressAddressesKaspaAddressUtxosGet**](KaspaAddressesAPI.md#GetUtxosForAddressAddressesKaspaAddressUtxosGet) | **Get** /addresses/{kaspaAddress}/utxos | Get Utxos For Address
[**GetUtxosForAddressesAddressesUtxosPost**](KaspaAddressesAPI.md#GetUtxosForAddressesAddressesUtxosPost) | **Post** /addresses/utxos | Get Utxos For Addresses



## GetAddressesActiveAddressesActivePost

> []TxIdResponse GetAddressesActiveAddressesActivePost(ctx).AddressesActiveRequest(addressesActiveRequest).Execute()

Get Addresses Active



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
	addressesActiveRequest := *openapiclient.NewAddressesActiveRequest() // AddressesActiveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaAddressesAPI.GetAddressesActiveAddressesActivePost(context.Background()).AddressesActiveRequest(addressesActiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaAddressesAPI.GetAddressesActiveAddressesActivePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressesActiveAddressesActivePost`: []TxIdResponse
	fmt.Fprintf(os.Stdout, "Response from `KaspaAddressesAPI.GetAddressesActiveAddressesActivePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressesActiveAddressesActivePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addressesActiveRequest** | [**AddressesActiveRequest**](AddressesActiveRequest.md) |  | 

### Return type

[**[]TxIdResponse**](TxIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalanceFromKaspaAddressAddressesKaspaAddressBalanceGet

> BalanceResponse GetBalanceFromKaspaAddressAddressesKaspaAddressBalanceGet(ctx, kaspaAddress).Execute()

Get Balance From Kaspa Address



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
	kaspaAddress := "kaspaAddress_example" // string | Kaspa address as string e.g. kaspatest:qpqz2vxj23kvh0m73ta2jjn2u4cv4tlufqns2eap8mxyyt0rvrxy6ejkful67

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaAddressesAPI.GetBalanceFromKaspaAddressAddressesKaspaAddressBalanceGet(context.Background(), kaspaAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaAddressesAPI.GetBalanceFromKaspaAddressAddressesKaspaAddressBalanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalanceFromKaspaAddressAddressesKaspaAddressBalanceGet`: BalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `KaspaAddressesAPI.GetBalanceFromKaspaAddressAddressesKaspaAddressBalanceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kaspaAddress** | **string** | Kaspa address as string e.g. kaspatest:qpqz2vxj23kvh0m73ta2jjn2u4cv4tlufqns2eap8mxyyt0rvrxy6ejkful67 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceFromKaspaAddressAddressesKaspaAddressBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BalanceResponse**](BalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalancesFromKaspaAddressesAddressesBalancesPost

> []BalancesByAddressEntry GetBalancesFromKaspaAddressesAddressesBalancesPost(ctx).BalanceRequest(balanceRequest).Execute()

Get Balances From Kaspa Addresses



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
	balanceRequest := *openapiclient.NewBalanceRequest() // BalanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaAddressesAPI.GetBalancesFromKaspaAddressesAddressesBalancesPost(context.Background()).BalanceRequest(balanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaAddressesAPI.GetBalancesFromKaspaAddressesAddressesBalancesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalancesFromKaspaAddressesAddressesBalancesPost`: []BalancesByAddressEntry
	fmt.Fprintf(os.Stdout, "Response from `KaspaAddressesAPI.GetBalancesFromKaspaAddressesAddressesBalancesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBalancesFromKaspaAddressesAddressesBalancesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **balanceRequest** | [**BalanceRequest**](BalanceRequest.md) |  | 

### Return type

[**[]BalancesByAddressEntry**](BalancesByAddressEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFullTransactionsForAddressAddressesKaspaAddressFullTransactionsGet

> []TxModel GetFullTransactionsForAddressAddressesKaspaAddressFullTransactionsGet(ctx, kaspaAddress).Limit(limit).Offset(offset).Fields(fields).ResolvePreviousOutpoints(resolvePreviousOutpoints).Execute()

Get Full Transactions For Address



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
	kaspaAddress := "kaspaAddress_example" // string | Kaspa address as string e.g. kaspatest:qpqz2vxj23kvh0m73ta2jjn2u4cv4tlufqns2eap8mxyyt0rvrxy6ejkful67
	limit := int32(56) // int32 | The number of records to get (optional) (default to 50)
	offset := int32(56) // int32 | The offset from which to get records (optional) (default to 0)
	fields := "fields_example" // string |  (optional) (default to "")
	resolvePreviousOutpoints := openapiclient.endpoints__get_address_transactions__PreviousOutpointLookupMode("no") // EndpointsGetAddressTransactionsPreviousOutpointLookupMode | Use this parameter if you want to fetch the TransactionInput previous outpoint details. Light fetches only the adress and amount. Full fetches the whole TransactionOutput and adds it into each TxInput. (optional) (default to "no")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaAddressesAPI.GetFullTransactionsForAddressAddressesKaspaAddressFullTransactionsGet(context.Background(), kaspaAddress).Limit(limit).Offset(offset).Fields(fields).ResolvePreviousOutpoints(resolvePreviousOutpoints).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaAddressesAPI.GetFullTransactionsForAddressAddressesKaspaAddressFullTransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFullTransactionsForAddressAddressesKaspaAddressFullTransactionsGet`: []TxModel
	fmt.Fprintf(os.Stdout, "Response from `KaspaAddressesAPI.GetFullTransactionsForAddressAddressesKaspaAddressFullTransactionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kaspaAddress** | **string** | Kaspa address as string e.g. kaspatest:qpqz2vxj23kvh0m73ta2jjn2u4cv4tlufqns2eap8mxyyt0rvrxy6ejkful67 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFullTransactionsForAddressAddressesKaspaAddressFullTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of records to get | [default to 50]
 **offset** | **int32** | The offset from which to get records | [default to 0]
 **fields** | **string** |  | [default to &quot;&quot;]
 **resolvePreviousOutpoints** | [**EndpointsGetAddressTransactionsPreviousOutpointLookupMode**](EndpointsGetAddressTransactionsPreviousOutpointLookupMode.md) | Use this parameter if you want to fetch the TransactionInput previous outpoint details. Light fetches only the adress and amount. Full fetches the whole TransactionOutput and adds it into each TxInput. | [default to &quot;no&quot;]

### Return type

[**[]TxModel**](TxModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFullTransactionsForAddressPageAddressesKaspaAddressFullTransactionsPageGet

> []TxModel GetFullTransactionsForAddressPageAddressesKaspaAddressFullTransactionsPageGet(ctx, kaspaAddress).Limit(limit).Before(before).After(after).Fields(fields).ResolvePreviousOutpoints(resolvePreviousOutpoints).Execute()

Get Full Transactions For Address Page



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
	kaspaAddress := "kaspaAddress_example" // string | Kaspa address as string e.g. kaspatest:qpqz2vxj23kvh0m73ta2jjn2u4cv4tlufqns2eap8mxyyt0rvrxy6ejkful67
	limit := int32(56) // int32 | The max number of records to get. For paging combine with using 'before/after' from oldest previous result. Use value of X-Next-Page-Before/-After as long as header is present to continue paging. The actual number of transactions returned for each page can be > limit. (optional) (default to 50)
	before := int32(56) // int32 | Only include transactions with block time before this (epoch-millis) (optional) (default to 0)
	after := int32(56) // int32 | Only include transactions with block time after this (epoch-millis) (optional) (default to 0)
	fields := "fields_example" // string |  (optional) (default to "")
	resolvePreviousOutpoints := openapiclient.endpoints__get_address_transactions__PreviousOutpointLookupMode("no") // EndpointsGetAddressTransactionsPreviousOutpointLookupMode | Use this parameter if you want to fetch the TransactionInput previous outpoint details. Light fetches only the adress and amount. Full fetches the whole TransactionOutput and adds it into each TxInput. (optional) (default to "no")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaAddressesAPI.GetFullTransactionsForAddressPageAddressesKaspaAddressFullTransactionsPageGet(context.Background(), kaspaAddress).Limit(limit).Before(before).After(after).Fields(fields).ResolvePreviousOutpoints(resolvePreviousOutpoints).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaAddressesAPI.GetFullTransactionsForAddressPageAddressesKaspaAddressFullTransactionsPageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFullTransactionsForAddressPageAddressesKaspaAddressFullTransactionsPageGet`: []TxModel
	fmt.Fprintf(os.Stdout, "Response from `KaspaAddressesAPI.GetFullTransactionsForAddressPageAddressesKaspaAddressFullTransactionsPageGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kaspaAddress** | **string** | Kaspa address as string e.g. kaspatest:qpqz2vxj23kvh0m73ta2jjn2u4cv4tlufqns2eap8mxyyt0rvrxy6ejkful67 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFullTransactionsForAddressPageAddressesKaspaAddressFullTransactionsPageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The max number of records to get. For paging combine with using &#39;before/after&#39; from oldest previous result. Use value of X-Next-Page-Before/-After as long as header is present to continue paging. The actual number of transactions returned for each page can be &gt; limit. | [default to 50]
 **before** | **int32** | Only include transactions with block time before this (epoch-millis) | [default to 0]
 **after** | **int32** | Only include transactions with block time after this (epoch-millis) | [default to 0]
 **fields** | **string** |  | [default to &quot;&quot;]
 **resolvePreviousOutpoints** | [**EndpointsGetAddressTransactionsPreviousOutpointLookupMode**](EndpointsGetAddressTransactionsPreviousOutpointLookupMode.md) | Use this parameter if you want to fetch the TransactionInput previous outpoint details. Light fetches only the adress and amount. Full fetches the whole TransactionOutput and adds it into each TxInput. | [default to &quot;no&quot;]

### Return type

[**[]TxModel**](TxModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNameForAddressAddressesKaspaAddressNameGet

> AddressName GetNameForAddressAddressesKaspaAddressNameGet(ctx, kaspaAddress).Execute()

Get Name For Address



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
	kaspaAddress := "kaspaAddress_example" // string | Kaspa address as string e.g. kaspa:qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqkx9awp4e

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaAddressesAPI.GetNameForAddressAddressesKaspaAddressNameGet(context.Background(), kaspaAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaAddressesAPI.GetNameForAddressAddressesKaspaAddressNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNameForAddressAddressesKaspaAddressNameGet`: AddressName
	fmt.Fprintf(os.Stdout, "Response from `KaspaAddressesAPI.GetNameForAddressAddressesKaspaAddressNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kaspaAddress** | **string** | Kaspa address as string e.g. kaspa:qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqkx9awp4e | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNameForAddressAddressesKaspaAddressNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressName**](AddressName.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionCountForAddressAddressesKaspaAddressTransactionsCountGet

> TransactionCount GetTransactionCountForAddressAddressesKaspaAddressTransactionsCountGet(ctx, kaspaAddress).Execute()

Get Transaction Count For Address



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
	kaspaAddress := "kaspaAddress_example" // string | Kaspa address as string e.g. kaspatest:qpqz2vxj23kvh0m73ta2jjn2u4cv4tlufqns2eap8mxyyt0rvrxy6ejkful67

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaAddressesAPI.GetTransactionCountForAddressAddressesKaspaAddressTransactionsCountGet(context.Background(), kaspaAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaAddressesAPI.GetTransactionCountForAddressAddressesKaspaAddressTransactionsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionCountForAddressAddressesKaspaAddressTransactionsCountGet`: TransactionCount
	fmt.Fprintf(os.Stdout, "Response from `KaspaAddressesAPI.GetTransactionCountForAddressAddressesKaspaAddressTransactionsCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kaspaAddress** | **string** | Kaspa address as string e.g. kaspatest:qpqz2vxj23kvh0m73ta2jjn2u4cv4tlufqns2eap8mxyyt0rvrxy6ejkful67 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionCountForAddressAddressesKaspaAddressTransactionsCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransactionCount**](TransactionCount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUtxosForAddressAddressesKaspaAddressUtxosGet

> []UtxoResponse GetUtxosForAddressAddressesKaspaAddressUtxosGet(ctx, kaspaAddress).Execute()

Get Utxos For Address



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
	kaspaAddress := "kaspaAddress_example" // string | Kaspa address as string e.g. kaspatest:qpqz2vxj23kvh0m73ta2jjn2u4cv4tlufqns2eap8mxyyt0rvrxy6ejkful67

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaAddressesAPI.GetUtxosForAddressAddressesKaspaAddressUtxosGet(context.Background(), kaspaAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaAddressesAPI.GetUtxosForAddressAddressesKaspaAddressUtxosGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUtxosForAddressAddressesKaspaAddressUtxosGet`: []UtxoResponse
	fmt.Fprintf(os.Stdout, "Response from `KaspaAddressesAPI.GetUtxosForAddressAddressesKaspaAddressUtxosGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kaspaAddress** | **string** | Kaspa address as string e.g. kaspatest:qpqz2vxj23kvh0m73ta2jjn2u4cv4tlufqns2eap8mxyyt0rvrxy6ejkful67 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUtxosForAddressAddressesKaspaAddressUtxosGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UtxoResponse**](UtxoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUtxosForAddressesAddressesUtxosPost

> []UtxoResponse GetUtxosForAddressesAddressesUtxosPost(ctx).UtxoRequest(utxoRequest).Execute()

Get Utxos For Addresses



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
	utxoRequest := *openapiclient.NewUtxoRequest() // UtxoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaAddressesAPI.GetUtxosForAddressesAddressesUtxosPost(context.Background()).UtxoRequest(utxoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaAddressesAPI.GetUtxosForAddressesAddressesUtxosPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUtxosForAddressesAddressesUtxosPost`: []UtxoResponse
	fmt.Fprintf(os.Stdout, "Response from `KaspaAddressesAPI.GetUtxosForAddressesAddressesUtxosPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUtxosForAddressesAddressesUtxosPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **utxoRequest** | [**UtxoRequest**](UtxoRequest.md) |  | 

### Return type

[**[]UtxoResponse**](UtxoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

