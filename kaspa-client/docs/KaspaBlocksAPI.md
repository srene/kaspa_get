# \KaspaBlocksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlockBlocksBlockIdGet**](KaspaBlocksAPI.md#GetBlockBlocksBlockIdGet) | **Get** /blocks/{blockId} | Get Block
[**GetBlocksBlocksGet**](KaspaBlocksAPI.md#GetBlocksBlocksGet) | **Get** /blocks | Get Blocks
[**GetBlocksFromBluescoreBlocksFromBluescoreGet**](KaspaBlocksAPI.md#GetBlocksFromBluescoreBlocksFromBluescoreGet) | **Get** /blocks-from-bluescore | Get Blocks From Bluescore



## GetBlockBlocksBlockIdGet

> BlockModel GetBlockBlocksBlockIdGet(ctx, blockId).IncludeTransactions(includeTransactions).IncludeColor(includeColor).Execute()

Get Block



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
	blockId := "blockId_example" // string | 
	includeTransactions := true // bool |  (optional) (default to true)
	includeColor := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaBlocksAPI.GetBlockBlocksBlockIdGet(context.Background(), blockId).IncludeTransactions(includeTransactions).IncludeColor(includeColor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaBlocksAPI.GetBlockBlocksBlockIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockBlocksBlockIdGet`: BlockModel
	fmt.Fprintf(os.Stdout, "Response from `KaspaBlocksAPI.GetBlockBlocksBlockIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockBlocksBlockIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTransactions** | **bool** |  | [default to true]
 **includeColor** | **bool** |  | [default to false]

### Return type

[**BlockModel**](BlockModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlocksBlocksGet

> BlockResponse GetBlocksBlocksGet(ctx).LowHash(lowHash).IncludeBlocks(includeBlocks).IncludeTransactions(includeTransactions).Execute()

Get Blocks



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
	lowHash := "lowHash_example" // string | 
	includeBlocks := true // bool |  (optional) (default to false)
	includeTransactions := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaBlocksAPI.GetBlocksBlocksGet(context.Background()).LowHash(lowHash).IncludeBlocks(includeBlocks).IncludeTransactions(includeTransactions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaBlocksAPI.GetBlocksBlocksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlocksBlocksGet`: BlockResponse
	fmt.Fprintf(os.Stdout, "Response from `KaspaBlocksAPI.GetBlocksBlocksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlocksBlocksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lowHash** | **string** |  | 
 **includeBlocks** | **bool** |  | [default to false]
 **includeTransactions** | **bool** |  | [default to false]

### Return type

[**BlockResponse**](BlockResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlocksFromBluescoreBlocksFromBluescoreGet

> []BlockModel GetBlocksFromBluescoreBlocksFromBluescoreGet(ctx).BlueScore(blueScore).IncludeTransactions(includeTransactions).Execute()

Get Blocks From Bluescore



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
	blueScore := int32(56) // int32 |  (optional) (default to 43679173)
	includeTransactions := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KaspaBlocksAPI.GetBlocksFromBluescoreBlocksFromBluescoreGet(context.Background()).BlueScore(blueScore).IncludeTransactions(includeTransactions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaspaBlocksAPI.GetBlocksFromBluescoreBlocksFromBluescoreGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlocksFromBluescoreBlocksFromBluescoreGet`: []BlockModel
	fmt.Fprintf(os.Stdout, "Response from `KaspaBlocksAPI.GetBlocksFromBluescoreBlocksFromBluescoreGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlocksFromBluescoreBlocksFromBluescoreGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueScore** | **int32** |  | [default to 43679173]
 **includeTransactions** | **bool** |  | [default to false]

### Return type

[**[]BlockModel**](BlockModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

