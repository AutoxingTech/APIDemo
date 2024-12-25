package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// RobotListResponse represents the response for robot list
type RobotListResponse struct {
	Status int `json:"status"`
	Data   struct {
		List []Robot `json:"list"`
	} `json:"data"`
}

// RobotStateResponse represents the response for robot state
type RobotStateResponse struct {
	Status int        `json:"status"`
	Data   RobotState `json:"data"`
}

// Robot represents a single robot's data
type Robot struct {
	RobotID  string `json:"robotId"`
	IsOnLine bool   `json:"isOnLine"`
	// Add other robot fields as needed
}

// RobotState represents the state of a robot
type RobotState struct {
	// Add specific state fields as needed
	// Example:
	// Battery    int    `json:"battery"`
	// Status     string `json:"status"`
}

// RobotManager handles robot-related operations
type RobotManager struct {
	token     string
	URLPrefix string
}

// NewRobotManager creates a new instance of RobotManager
func NewRobotManager(token string, urlPrefix string) *RobotManager {
	return &RobotManager{
		token:     token,
		URLPrefix: urlPrefix,
	}
}

// GetRobotList retrieves the list of robots
func (rm *RobotManager) GetRobotList() (bool, []Robot) {
	url := rm.URLPrefix + "/robot/v1.1/list"

	data := map[string]interface{}{
		"pageSize": 10,
		"pageNum":  1,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return false, nil
	}

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false, nil
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Token", rm.token)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false, nil
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return false, nil
	}

	// Parse response
	var listResp RobotListResponse
	if err := json.Unmarshal(body, &listResp); err != nil {
		fmt.Println("Error parsing response:", err)
		return false, nil
	}

	// Check response status
	if listResp.Status == 200 {
		return true, listResp.Data.List
	}

	return false, nil
}

// GetRobotState retrieves the state of a specific robot
func (rm *RobotManager) GetRobotState(robotID string) (bool, RobotState) {
	url := fmt.Sprintf("%s/robot/v1.1/%s/state", rm.URLPrefix, robotID)

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false, RobotState{}
	}

	// Set headers
	req.Header.Set("X-Token", rm.token)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false, RobotState{}
	}
	defer resp.Body.Close()

	// Read and print response for debugging
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return false, RobotState{}
	}
	fmt.Println(string(body))

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return false, RobotState{}
	}

	// Parse response
	var stateResp RobotStateResponse
	if err := json.Unmarshal(body, &stateResp); err != nil {
		fmt.Println("Error parsing response:", err)
		return false, RobotState{}
	}

	// Check response status
	if stateResp.Status == 200 {
		return true, stateResp.Data
	}

	return false, RobotState{}
}
