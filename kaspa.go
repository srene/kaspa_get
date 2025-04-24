package main

import (
	"context"
	"fmt"
	"log"

	openapiclient "github.com/srene/kaspa_get/kaspa-client"
)

// TransactionResponse represents the structure of the response from the Kaspa API
/*type TransactionResponse struct {
	Payload string `json:"payload"`
	// Add other relevant fields based on the full response structure
}

func getTransaction(txID string) (*TransactionResponse, error) {
	url := fmt.Sprintf("https://api-tn10.kaspa.org/transactions/%s", txID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching transaction: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error response from server: %s", string(bodyBytes))
	}

	var txResponse TransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&txResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return &txResponse, nil
}

func main() {
	txID := "fc57832f3fac15e61f320c2c39781680d17f8fb2f22c88c6ede08438d64387f4"
	tx, err := getTransaction(txID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Fetched transaction: %+v\n", tx.Payload)
}*/

func main() {
	cfg := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(cfg)

	resp, _, err := apiClient.YourAPIService.YourEndpointMethod(context.Background(), "someParam")
	if err != nil {
		log.Fatalf("API call failed: %v", err)
	}
	fmt.Printf("Result: %+v\n", resp)
}
