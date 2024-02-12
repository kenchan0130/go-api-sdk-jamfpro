package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Fetch all printers
	printers, err := client.GetPrinters()
	if err != nil {
		log.Fatalf("Error fetching printers: %v", err)
	}

	fmt.Println("Printers fetched. Starting deletion process:")

	// Iterate over each printer and delete
	for _, printer := range printers.Printer {
		fmt.Printf("deleting printer ID: %d, Name: %s\n", printer.ID, printer.Name)

		err = client.DeletePrinterByID(printer.ID)
		if err != nil {
			log.Printf("error deleting printer ID %d: %v\n", printer.ID, err)
			continue // Move to the next printer if there's an error
		}

		fmt.Printf("printer ID %d deleted successfully.\n", printer.ID)
	}

	fmt.Println("Printer deletion process completed.")
}
