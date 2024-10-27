package campaign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mp-cli/config"
	"net/http"

	"github.com/spf13/cobra"
)

func newCreateCommand() *cobra.Command {
	var name, description, template string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new campaign",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.LoadConfig()
			if err != nil {
				fmt.Printf("Error loading config: %v\n", err)
				return
			}

			// Authenticate user
			token, err := authenticateUser(cfg)
			if err != nil {
				fmt.Printf("Error authenticating user: %v\n", err)
				return
			}

			// Create campaign
			campaign, err := createCampaign(cfg, token, name, description, template)
			if err != nil {
				fmt.Printf("Error creating campaign: %v\n", err)
				return
			}

			fmt.Printf("Campaign created successfully: %+v\n", campaign)
		},
	}

	// Add flags for campaign details
	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the campaign")
	cmd.Flags().StringVarP(&description, "description", "d", "", "Description of the campaign")
	cmd.Flags().StringVarP(&template, "template", "t", "", "Template for the campaign")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("description")
	cmd.MarkFlagRequired("template")

	return cmd
}

func authenticateUser(cfg *config.Config) (string, error) {
	url := fmt.Sprintf("%s/api/user/login", cfg.APIBaseURL)
	payload := map[string]string{
		"username": cfg.Username,
		"password": cfg.Password,
	}

	resp, err := makeRequest("POST", url, payload, "")
	if err != nil {
		return "", err
	}

	var loginResp struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(resp, &loginResp); err != nil {
		return "", fmt.Errorf("failed to parse login response: %v", err)
	}

	return loginResp.Token, nil
}

func createCampaign(cfg *config.Config, token, name, description, template string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/campaign", cfg.APIBaseURL)
	payload := map[string]string{
		"name":        name,
		"description": description,
		"template":    template,
		"owner_id":    cfg.UserID,
	}

	resp, err := makeRequest("POST", url, payload, token)
	if err != nil {
		return nil, err
	}

	var campaign map[string]interface{}
	if err := json.Unmarshal(resp, &campaign); err != nil {
		return nil, fmt.Errorf("failed to parse campaign response: %v", err)
	}

	return campaign, nil
}

func makeRequest(method, url string, payload interface{}, token string) ([]byte, error) {
	client := &http.Client{}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}
