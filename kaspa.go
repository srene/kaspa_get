package main

import (
	"context"
	"fmt"

	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/srene/kaspa_get/kaspa_rpc_client_go/client"
	"github.com/srene/kaspa_get/kaspa_rpc_client_go/client/operations"
)

func main() {

	// Create a transport for the API
	transport := openapiclient.New("api-tn10.kaspa.org", "", []string{"https"})
	kaspaClient := client.New(transport, nil)

	// Example call (replace with actual operation name from the spec)
	params := operations.NewSomeOperationParamsWithContext(context.Background())

	resp, err := kaspaClient.Operations.SomeOperation(params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", resp)
}
