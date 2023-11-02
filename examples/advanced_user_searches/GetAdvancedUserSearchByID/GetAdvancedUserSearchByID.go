package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for Jamf Pro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new Jamf Pro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// ID of the advanced user search to retrieve
	advancedUserSearchID := 1 // Replace with the actual ID

	// Call GetAdvancedUserSearchByID function
	advancedUserSearch, err := client.GetAdvancedUserSearchByID(advancedUserSearchID)
	if err != nil {
		log.Fatalf("Error fetching advanced user search by ID: %v", err)
	}

	// Pretty print the advanced user search in XML
	advancedUserSearchXML, err := xml.MarshalIndent(advancedUserSearch, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling advanced user search data: %v", err)
	}
	fmt.Println("Fetched Advanced User Search:\n", string(advancedUserSearchXML))
}