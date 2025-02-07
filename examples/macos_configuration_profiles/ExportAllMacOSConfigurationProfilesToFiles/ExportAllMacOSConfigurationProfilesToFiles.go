package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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

	profilesList, err := client.GetMacOSConfigurationProfiles()
	if err != nil {
		log.Fatalf("Failed to fetch macOS Configuration Profiles: %v", err)
	}

	fmt.Println("Found the following macOS Configuration Profiles:")
	for _, profile := range profilesList.Results {
		fmt.Printf("ID: %d, Name: %s\n", profile.ID, profile.Name)
	}
	fmt.Println("These profiles will be exported.")

	exportDir := "/Users/joseph/github/go-api-sdk-jamfpro"
	if err := os.MkdirAll(exportDir, 0750); err != nil {
		log.Fatalf("Failed to create export directory: %v", err)
	}

	for _, profile := range profilesList.Results {
		respProfile, err := client.GetMacOSConfigurationProfileByID(profile.ID)
		if err != nil {
			log.Printf("Failed to fetch profile with ID %d: %v", profile.ID, err)
			continue
		}

		xmlData, err := xml.MarshalIndent(respProfile, "", "  ")
		if err != nil {
			log.Printf("Failed to convert profile with ID %d to XML: %v", profile.ID, err)
			continue
		}

		payloadsContent := extractPayloads(string(xmlData))
		if payloadsContent == "" {
			log.Printf("No <payloads> content found for profile ID %d", profile.ID)
			continue
		}

		filename := filepath.Join(exportDir, profile.Name+".mobileconfig")
		file, err := os.Create(filename)
		if err != nil {
			log.Printf("Failed to create file for profile with ID %d: %v", profile.ID, err)
			continue
		}
		defer file.Close()

		if _, err := file.WriteString(payloadsContent); err != nil {
			log.Printf("Failed to write to file for profile with ID %d: %v", profile.ID, err)
			continue
		}

		fmt.Printf("Exported profile with ID %d to %s\n", profile.ID, filename)
	}

	fmt.Println("Export completed!")
}

func extractPayloads(xmlData string) string {
	startTag := "<payloads>"
	endTag := "</payloads>"
	startIndex := strings.Index(xmlData, startTag)
	endIndex := strings.Index(xmlData, endTag)

	if startIndex == -1 || endIndex == -1 {
		return ""
	}

	return xmlData[startIndex+len(startTag) : endIndex]
}
