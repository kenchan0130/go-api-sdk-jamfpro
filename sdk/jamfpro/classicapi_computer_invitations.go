// classicapi_computer_invitations.go
// Jamf Pro Classic Api - Computer Invitations
// api reference: https://developer.jamf.com/jamf-pro/reference/computerinvitations
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriComputerInvitations = "/JSSResource/computerinvitations"

type ResponseComputerInvitationsList struct {
	Size               int                        `xml:"size"`
	ComputerInvitation []ComputerInvitationDetail `xml:"computer_invitation"`
}

type ComputerInvitationDetail struct {
	ID                  int    `xml:"id,omitempty"`
	Invitation          int64  `xml:"invitation,omitempty"`
	InvitationType      string `xml:"invitation_type,omitempty"`
	ExpirationDate      string `xml:"expiration_date,omitempty"`
	ExpirationDateUTC   string `xml:"expiration_date_utc,omitempty"`
	ExpirationDateEpoch int64  `xml:"expiration_date_epoch,omitempty"`
}

type ResponseComputerInvitation struct {
	ID                          int                    `xml:"id,omitempty"`
	Invitation                  string                 `xml:"invitation,omitempty"`
	InvitationStatus            string                 `xml:"invitation_status,omitempty"`
	InvitationType              string                 `xml:"invitation_type,omitempty"`
	ExpirationDate              string                 `xml:"expiration_date,omitempty"`
	ExpirationDateUTC           string                 `xml:"expiration_date_utc,omitempty"`
	ExpirationDateEpoch         int64                  `xml:"expiration_date_epoch,omitempty"`
	SSHUsername                 string                 `xml:"ssh_username,omitempty"`
	SSHPassword                 string                 `xml:"ssh_password,omitempty"`
	MultipleUsersAllowed        bool                   `xml:"multiple_users_allowed,omitempty"`
	TimesUsed                   int                    `xml:"times_used,omitempty"`
	CreateAccountIfDoesNotExist bool                   `xml:"create_account_if_does_not_exist,omitempty"`
	HideAccount                 bool                   `xml:"hide_account,omitempty"`
	LockDownSSH                 bool                   `xml:"lock_down_ssh,omitempty"`
	InvitedUserUUID             string                 `xml:"invited_user_uuid,omitempty"`
	EnrollIntoSite              ComputerInvitationSite `xml:"enroll_into_site,omitempty"`
	KeepExistingSiteMembership  bool                   `xml:"keep_existing_site_membership,omitempty"`
	Site                        ComputerInvitationSite `xml:"site,omitempty"`
}

type ComputerInvitationSite struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// GetComputerInvitations retrieves a list of all computer invitations.
func (c *Client) GetComputerInvitations() (*ResponseComputerInvitationsList, error) {
	endpoint := uriComputerInvitations

	var invitations ResponseComputerInvitationsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &invitations)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Computer Invitations: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &invitations, nil
}

// GetComputerInvitationByID retrieves a computer invitation by its ID.
func (c *Client) GetComputerInvitationByID(invitationID int) (*ResponseComputerInvitation, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerInvitations, invitationID)

	var invitation ResponseComputerInvitation
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &invitation)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Computer Invitation by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &invitation, nil
}

// GetComputerInvitationsByName retrieves a computer invitation by its invitation Name.
func (c *Client) GetComputerInvitationsByInvitationID(invitationID int) (*ResponseComputerInvitation, error) {
	endpoint := fmt.Sprintf("%s/invitation/%d", uriComputerInvitations, invitationID)

	var invitation ResponseComputerInvitation
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &invitation)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Computer Invitation by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &invitation, nil
}

// CreateComputerInvitation creates a new computer invitation.
func (c *Client) CreateComputerInvitation(invitation *ResponseComputerInvitation) (*ResponseComputerInvitation, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriComputerInvitations)

	// Check if site is not provided and set default values
	if invitation.Site.ID == 0 && invitation.Site.Name == "" {
		invitation.Site = ComputerInvitationSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Wrap the invitation request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"computer_invitation"`
		*ResponseComputerInvitation
	}{
		ResponseComputerInvitation: invitation,
	}

	var createdInvitation ResponseComputerInvitation
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdInvitation)
	if err != nil {
		return nil, fmt.Errorf("failed to create Computer Invitation: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdInvitation, nil
}

// DeleteComputerInvitationByID deletes a computer invitation by its ID.
func (c *Client) DeleteComputerInvitationByID(invitationID int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerInvitations, invitationID)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Computer Invitation by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
