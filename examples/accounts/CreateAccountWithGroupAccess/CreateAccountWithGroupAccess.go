package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Assemble the request body for creating an account
	accountDetail := &jamfpro.ResourceAccount{
		Name:                "Barry White",
		DirectoryUser:       false,
		FullName:            "Barry White",
		Email:               "Barry.White@company.com",
		EmailAddress:        "Barry.White@company.com",
		Enabled:             "Enabled",
		ForcePasswordChange: true,
		AccessLevel:         "Group Access", // Full Access / Site Access / Group Access
		PrivilegeSet:        "Custom",       // Administrator / Auditor / Enrollment Only / Custom
		Password:            "this is a really secure password 390423049823409894382092348092348",
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Groups: []jamfpro.AccountsListSubsetGroups{
			{
				ID:   195,
				Name: "tf-localtest-standard-group",
			},
		},
	}

	// Call CreateAccountByID function
	createdAccount, err := client.CreateAccount(accountDetail)

	if err != nil {
		log.Fatal(err)
	}

	// Pretty print the created account details
	accountXML, err := xml.MarshalIndent(createdAccount, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling account data: %v", err)
	}
	fmt.Println("Created Account Details:", string(accountXML))
}
