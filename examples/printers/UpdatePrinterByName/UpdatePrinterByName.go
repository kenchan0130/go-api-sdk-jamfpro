package main

import (
	"encoding/xml"
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

	// Instantiate the default logger and set the desired log level
	//logLevel := logger.LogLevelInfo // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

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

	// Printer details to create
	updatedPrinter := &jamfpro.ResourcePrinter{
		Name:        "HP 9th Floor",
		Category:    "",
		URI:         "lpd://10.1.20.204/",
		CUPSName:    "HP_9th_Floor",
		Location:    "string",
		Model:       "HP LaserJet 500 color MFP M575",
		Info:        "string",
		Notes:       "string",
		MakeDefault: true,
		UseGeneric:  true,
		PPD:         "9th_Floor_HP.ppd",
		PPDPath:     "/System/Library/Frameworks/ApplicationServices.framework/Versions/A/Frameworks/PrintCore.framework/Resources/Generic.ppd",
		PPDContents: "string",
	}

	name := "HP 9th Floor" // Replace with the actual printer name

	response, err := client.UpdatePrinterByName(name, updatedPrinter)
	if err != nil {
		fmt.Println("Error updating printer:", err)
		return
	}

	fmt.Printf("Printer updated successfully, ID: %d\n", response.ID)

	// Fetch the full details of the updated printer
	updatedPrinterDetails, err := client.GetPrinterByID(response.ID)
	if err != nil {
		fmt.Println("Error fetching updated printer details:", err)
		return
	}

	// Marshal the updated printer details to XML for display
	printerXML, err := xml.MarshalIndent(updatedPrinterDetails, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated printer to XML: %v", err)
	}

	fmt.Printf("Updated Printer Details:\n%s\n", printerXML)
}
