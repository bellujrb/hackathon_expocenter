package external

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func FetchUsersDetails() ([]Customer, error) {
	urlBase := os.Getenv("URL")
	if urlBase == "" {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
	resp, err := http.Get(urlBase + "/api/customers/details")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var customer []Customer
	if err := json.Unmarshal(body, &customer); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return customer, nil
}

func FetchStoreDetails() ([]Store, error) {
	urlBase := os.Getenv("URL")
	if urlBase == "" {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
	resp, err := http.Get(urlBase + "/api/stores/details")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var stores []Store
	if err := json.Unmarshal(body, &stores); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return stores, nil
}
