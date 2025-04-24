# \KaspaTransactionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateTransactionMassTransactionsMassPost**](KaspaTransactionsAPI.md#CalculateTransactionMassTransactionsMassPost) | **Post** /transactions/mass | Calculate Transaction Mass
[**GetTransactionTransactionsTransactionIdGet**](KaspaTransactionsAPI.md#GetTransactionTransactionsTransactionIdGet) | **Get** /transactions/{transactionId} | Get Transaction
[**SearchForTransactionsTransactionsSearchPost**](KaspaTransactionsAPI.md#SearchForTransactionsTransactionsSearchPost) | **Post** /transactions/search | Search For Transactions
[**SubmitANewTransactionTransactionsPost**](KaspaTransactionsAPI.md#SubmitANewTransactionTransactionsPost) | **Post** /transactions | Submit A New Transaction



## CalculateTransactionMassTransactionsMassPost

> TxMass CalculateTransactionMassTransactionsMassPost(ctx).SubmitTxModel(submitTxModel).Execute()

Calculate Transaction Mass



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
	submitTxModel := *openapiclient.NewSubmitTxModel(int32(123), []openapiclient.SubmitTxInput{*openapiclient.NewSubmitTxInput(*openapiclient.NewSubmitTxOutpoint("TransactionId_example", int32(123)), "SignatureScript_example", int32(123), int32(123))}, []openapiclient.SubmitTxOutput{*openapiclient.NewSubmitTxOutput(int32(123), *openapiclient.NewSubmitTxScriptPublicKey(int32(123), "ScriptPublicKey_example"))}) // SubmitTxModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaTransactionsAPI.CalculateTransactionMassTransactionsMassPost(context.Background()).SubmitTxModel(submitTxModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaTransactionsAPI.CalculateTransactionMassTransactionsMassPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculateTransactionMassTransactionsMassPost`: TxMass
	fmt.Fprintf(os.Stdout, "Response from `KaspaTransactionsAPI.CalculateTransactionMassTransactionsMassPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalculateTransactionMassTransactionsMassPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitTxModel** | [**SubmitTxModel**](SubmitTxModel.md) |  | 

### Return type

[**TxMass**](TxMass.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionTransactionsTransactionIdGet

> TxModel GetTransactionTransactionsTransactionIdGet(ctx, transactionId).BlockHash(blockHash).Inputs(inputs).Outputs(outputs).ResolvePreviousOutpoints(resolvePreviousOutpoints).Execute()

Get Transaction



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
	transactionId := "transactionId_example" // string | 
	blockHash := "blockHash_example" // string | Specify a containing block (if known) for faster lookup (optional)
	inputs := true // bool |  (optional) (default to true)
	outputs := true // bool |  (optional) (default to true)
	resolvePreviousOutpoints := openapiclient.endpoints__get_transactions__PreviousOutpointLookupMode("no") // EndpointsGetTransactionsPreviousOutpointLookupMode | Use this parameter if you want to fetch the TransactionInput previous outpoint details. Light fetches only the address and amount. Full fetches the whole TransactionOutput and adds it into each TxInput. (optional) (default to "no")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaTransactionsAPI.GetTransactionTransactionsTransactionIdGet(context.Background(), transactionId).BlockHash(blockHash).Inputs(inputs).Outputs(outputs).ResolvePreviousOutpoints(resolvePreviousOutpoints).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaTransactionsAPI.GetTransactionTransactionsTransactionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionTransactionsTransactionIdGet`: TxModel
	fmt.Fprintf(os.Stdout, "Response from `KaspaTransactionsAPI.GetTransactionTransactionsTransactionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionTransactionsTransactionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blockHash** | **string** | Specify a containing block (if known) for faster lookup | 
 **inputs** | **bool** |  | [default to true]
 **outputs** | **bool** |  | [default to true]
 **resolvePreviousOutpoints** | [**EndpointsGetTransactionsPreviousOutpointLookupMode**](EndpointsGetTransactionsPreviousOutpointLookupMode.md) | Use this parameter if you want to fetch the TransactionInput previous outpoint details. Light fetches only the address and amount. Full fetches the whole TransactionOutput and adds it into each TxInput. | [default to &quot;no&quot;]

### Return type

[**TxModel**](TxModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchForTransactionsTransactionsSearchPost

> []TxModel SearchForTransactionsTransactionsSearchPost(ctx).TxSearch(txSearch).Fields(fields).ResolvePreviousOutpoints(resolvePreviousOutpoints).Execute()

Search For Transactions



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
	txSearch := *openapiclient.NewTxSearch() // TxSearch | 
	fields := "fields_example" // string |  (optional) (default to "")
	resolvePreviousOutpoints := openapiclient.endpoints__get_transactions__PreviousOutpointLookupMode("no") // EndpointsGetTransactionsPreviousOutpointLookupMode | Use this parameter if you want to fetch the TransactionInput previous outpoint details. Light fetches only the address and amount. Full fetches the whole TransactionOutput and adds it into each TxInput. (optional) (default to "no")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaTransactionsAPI.SearchForTransactionsTransactionsSearchPost(context.Background()).TxSearch(txSearch).Fields(fields).ResolvePreviousOutpoints(resolvePreviousOutpoints).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaTransactionsAPI.SearchForTransactionsTransactionsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchForTransactionsTransactionsSearchPost`: []TxModel
	fmt.Fprintf(os.Stdout, "Response from `KaspaTransactionsAPI.SearchForTransactionsTransactionsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchForTransactionsTransactionsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **txSearch** | [**TxSearch**](TxSearch.md) |  | 
 **fields** | **string** |  | [default to &quot;&quot;]
 **resolvePreviousOutpoints** | [**EndpointsGetTransactionsPreviousOutpointLookupMode**](EndpointsGetTransactionsPreviousOutpointLookupMode.md) | Use this parameter if you want to fetch the TransactionInput previous outpoint details. Light fetches only the address and amount. Full fetches the whole TransactionOutput and adds it into each TxInput. | [default to &quot;no&quot;]

### Return type

[**[]TxModel**](TxModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitANewTransactionTransactionsPost

> SubmitTransactionResponse SubmitANewTransactionTransactionsPost(ctx).SubmitTransactionRequest(submitTransactionRequest).ReplaceByFee(replaceByFee).Execute()

Submit A New Transaction

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
	submitTransactionRequest := *openapiclient.NewSubmitTransactionRequest(*openapiclient.NewSubmitTxModel(int32(123), []openapiclient.SubmitTxInput{*openapiclient.NewSubmitTxInput(*openapiclient.NewSubmitTxOutpoint("TransactionId_example", int32(123)), "SignatureScript_example", int32(123), int32(123))}, []openapiclient.SubmitTxOutput{*openapiclient.NewSubmitTxOutput(int32(123), *openapiclient.NewSubmitTxScriptPublicKey(int32(123), "ScriptPublicKey_example"))})) // SubmitTransactionRequest | 
	replaceByFee := true // bool | Replace an existing transaction in the mempool (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaTransactionsAPI.SubmitANewTransactionTransactionsPost(context.Background()).SubmitTransactionRequest(submitTransactionRequest).ReplaceByFee(replaceByFee).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaTransactionsAPI.SubmitANewTransactionTransactionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitANewTransactionTransactionsPost`: SubmitTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `KaspaTransactionsAPI.SubmitANewTransactionTransactionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitANewTransactionTransactionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitTransactionRequest** | [**SubmitTransactionRequest**](SubmitTransactionRequest.md) |  | 
 **replaceByFee** | **bool** | Replace an existing transaction in the mempool | [default to false]

### Return type

[**SubmitTransactionResponse**](SubmitTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

