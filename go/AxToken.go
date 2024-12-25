package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Config struct to hold configuration
type Config struct {
	URLPrefix     string
	APPID         string
	APPSecret     string
	Authorization string
	RobotID       string
}

// TokenResponse represents the response from the server
type TokenResponse struct {
	Status int `json:"status"`
	Data   struct {
		Key        string `json:"key"`
		Token      string `json:"token"`
		ExpireTime int64  `json:"expireTime"`
	} `json:"data"`
}

// TokenManager handles token operations
type TokenManager struct {
	token      string
	expireTime int64
	key        string
	timestamp  int64
	ok         bool
}

// NewTokenManager creates a new instance of TokenManager
func NewTokenManager() *TokenManager {
	return &TokenManager{
		ok: false,
	}
}

// GetToken retrieves a valid token
func (tm *TokenManager) GetToken(config *Config) (bool, string) {
	if tm.ok {
		currentTime := time.Now().UnixMilli()
		if currentTime < tm.timestamp+tm.expireTime*1000 {
			return true, tm.token
		}
	}
	return tm.getTokenFromServer(config)
}

// getTokenFromServer fetches a new token from the server
func (tm *TokenManager) getTokenFromServer(config *Config) (bool, string) {
	url := config.URLPrefix + "/auth/v1.1/token"
	timestamp := time.Now().UnixMilli()

	// Create request data
	data := map[string]interface{}{
		"appId":     config.APPID,
		"timestamp": timestamp,
		"sign":      "",
	}

	// Calculate sign
	signStr := fmt.Sprintf("%s%d%s", config.APPID, timestamp, config.APPSecret)
	hasher := md5.New()
	hasher.Write([]byte(signStr))
	data["sign"] = hex.EncodeToString(hasher.Sum(nil))

	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		tm.ok = false
		return false, ""
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		tm.ok = false
		return false, ""
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", config.Authorization)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		tm.ok = false
		return false, ""
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		tm.ok = false
		return false, ""
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		tm.ok = false
		return false, ""
	}

	// Parse response
	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		fmt.Println("Error parsing response:", err)
		tm.ok = false
		return false, ""
	}

	// Check response status
	if tokenResp.Status == 200 {
		tm.key = tokenResp.Data.Key
		tm.token = tokenResp.Data.Token
		tm.expireTime = tokenResp.Data.ExpireTime
		tm.timestamp = timestamp
		tm.ok = true
		return true, tm.token
	}

	tm.ok = false
	return false, ""
}
