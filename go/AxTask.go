package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Action types and builders
type ActionType struct{}

var Action = ActionType{}

// PauseAction creates a pause action
func (a ActionType) PauseAction(duration int) map[string]interface{} {
	return map[string]interface{}{
		"type": 18,
		"data": map[string]interface{}{
			"pauseTime": duration,
		},
	}
}

// PlayAudioAction creates a play audio action
func (a ActionType) PlayAudioAction(audioId string) map[string]interface{} {
	return map[string]interface{}{
		"type": 5,
		"data": map[string]interface{}{
			"mode":     1,
			"url":      "",
			"audioId":  audioId,
			"interval": -1,
			"num":      1,
			"volume":   100,
			"channel":  1,
			"duration": -1,
		},
	}
}

// WaitAction creates a wait action
func (a ActionType) WaitAction(userData interface{}) map[string]interface{} {
	return map[string]interface{}{
		"type": 40,
		"data": map[string]interface{}{
			"userData": userData,
		},
	}
}

// POI represents a point of interest
type POI struct {
	AreaID     string    `json:"areaId"`
	Coordinate []float64 `json:"coordinate"`
	Name       string    `json:"name"`
	Yaw        float64   `json:"yaw"`
}

// TaskPoint represents a point in the task
type TaskPoint struct {
	pt map[string]interface{}
}

// NewTaskPoint creates a new task point
func NewTaskPoint(poi POI, ignoreYaw bool) *TaskPoint {
	pt := map[string]interface{}{
		"areaId":     poi.AreaID,
		"x":          poi.Coordinate[0],
		"y":          poi.Coordinate[1],
		"type":       0,
		"stopRadius": 1,
		"ext": map[string]interface{}{
			"name": poi.Name,
		},
		"stepActs": []interface{}{},
	}

	if !ignoreYaw {
		pt["yaw"] = poi.Yaw
	}

	return &TaskPoint{pt: pt}
}

// AddStepActs adds a step action to the task point
func (tp *TaskPoint) AddStepActs(stepAct map[string]interface{}) *TaskPoint {
	tp.pt["stepActs"] = append(tp.pt["stepActs"].([]interface{}), stepAct)
	return tp
}

// TaskBuilder helps build a task
type TaskBuilder struct {
	task map[string]interface{}
}

// NewTaskBuilder creates a new task builder
func NewTaskBuilder(name, robotId string) *TaskBuilder {
	return &TaskBuilder{
		task: map[string]interface{}{
			"name":             name,
			"robotId":          robotId,
			"routeMode":        1,
			"runMode":          1,
			"runNum":           1,
			"taskType":         4,
			"runType":          21,
			"sourceType":       6,
			"ignorePublicSite": false,
			"speed":            1.0,
			"taskPts":          []interface{}{},
		},
	}
}

// AddTaskPt adds a task point to the task
func (tb *TaskBuilder) AddTaskPt(tp *TaskPoint) *TaskBuilder {
	tb.task["taskPts"] = append(tb.task["taskPts"].([]interface{}), tp.pt)
	return tb
}

// SetBackPt sets the back point for the task
func (tb *TaskBuilder) SetBackPt(pt *TaskPoint) *TaskBuilder {
	tb.task["backPt"] = pt.pt
	return tb
}

// GetTask returns the complete task
func (tb *TaskBuilder) GetTask() map[string]interface{} {
	return tb.task
}

// TaskManager handles task operations
type TaskManager struct {
	token     string
	URLPrefix string
}

// NewTaskManager creates a new task manager
func NewTaskManager(token string, urlPrefix string) *TaskManager {
	return &TaskManager{token: token, URLPrefix: urlPrefix}
}

// GetTaskInfo retrieves task information
func (tm *TaskManager) GetTaskInfo(taskId string) (bool, map[string]interface{}) {
	url := fmt.Sprintf("%s/task/v1.1/%s", tm.URLPrefix, taskId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false, nil
	}

	req.Header.Set("X-Token", tm.token)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return false, nil
	}

	var response struct {
		Status int                    `json:"status"`
		Data   map[string]interface{} `json:"data"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error parsing response:", err)
		return false, nil
	}

	if response.Status == 200 {
		return true, response.Data
	}

	return false, nil
}

// ExecuteTask executes a task
func (tm *TaskManager) ExecuteTask(taskId string) bool {
	url := fmt.Sprintf("%s/task/v1.1/%s/execute", tm.URLPrefix, taskId)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	req.Header.Set("X-Token", tm.token)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false
	}
	defer resp.Body.Close()

	var response struct {
		Status int `json:"status"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Error parsing response:", err)
		return false
	}

	return response.Status == 200
}

// NewTask creates a new task
func (tm *TaskManager) NewTask(taskData map[string]interface{}) (bool, string) {
	url := tm.URLPrefix + "/task/v1.1"

	jsonData, err := json.Marshal(taskData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return false, ""
	}
	fmt.Println("JSON:", string(jsonData))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false, ""
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Token", tm.token)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false, ""
	}
	defer resp.Body.Close()

	var response struct {
		Status int `json:"status"`
		Data   struct {
			TaskId string `json:"taskId"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Error parsing response:", err)
		return false, ""
	}

	fmt.Println("response:", response)

	if response.Status == 200 {
		return true, response.Data.TaskId
	}

	return false, ""
}
