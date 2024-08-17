package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Issue struct {
	ID      int     `json:"id"`
	Subject string  `json:"subject"`
	Tracker Tracker `json:"tracker"`
}

type Tracker struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type IssuesResponse struct {
	Issues []Issue `json:"issues"`
}

func main() {
	// Get the Redmine API key from environment variable
	apiKey := os.Getenv("REDMINE_API_KEY")
	if apiKey == "" {
		fmt.Println("REDMINE_API_KEY environment variable not set")
		os.Exit(1)
	}

	// Get the command-line arguments
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s [url]\n", os.Args[0])
		os.Exit(1)
	}
	inputURL := os.Args[1]

	// Parse and modify the URL to use `issues.json`
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		fmt.Printf("Invalid URL: %v\n", err)
		os.Exit(1)
	}
	parsedURL.Path = "redmine/issues.json"

	// Fetch issues
	issues, err := fetchIssues(parsedURL.String(), apiKey)
	if err != nil {
		fmt.Printf("Error fetching issues: %v\n", err)
		os.Exit(1)
	}

	// Display issues
	for _, issue := range issues {
		fmt.Printf("[%s #%d] %s\n", issue.Tracker.Name, issue.ID, issue.Subject)
	}
}

func fetchIssues(url string, apiKey string) ([]Issue, error) {
	// Create an HTTP client with TLS verification disabled
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Redmine-API-Key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var issuesResponse IssuesResponse
	if err := json.Unmarshal(body, &issuesResponse); err != nil {
		return nil, err
	}

	return issuesResponse.Issues, nil
}
