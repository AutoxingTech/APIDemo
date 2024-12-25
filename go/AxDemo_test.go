package main

import (
	"fmt"
	"os"
	"testing"
)

func TestAxToken(t *testing.T) {
	// Example configuration
	config := &Config{
		URLPrefix:     os.Getenv("URL_PREFIX"),
		APPID:         os.Getenv("APP_ID"),
		APPSecret:     os.Getenv("APP_SECRET"),
		Authorization: "APPCODE " + os.Getenv("Authorization"),
	}

	manager := NewTokenManager()
	success, token := manager.GetToken(config)
	fmt.Printf("Success: %v, Token: %s\n", success, token)
	if success {
		t.Log("TestAxToken passed")
		return
	}
	t.Error("TestAxToken failed")
}

func TestAxRobot(t *testing.T) {
	// Example configuration
	config := &Config{
		URLPrefix:     os.Getenv("URL_PREFIX"),
		APPID:         os.Getenv("APP_ID"),
		APPSecret:     os.Getenv("APP_SECRET"),
		Authorization: "APPCODE " + os.Getenv("Authorization"),
	}

	tokenManager := NewTokenManager()
	success, token := tokenManager.GetToken(config)

	if success {
		// Example usage
		manager := NewRobotManager(token, config.URLPrefix)

		// Get robot list
		success, robots := manager.GetRobotList()
		fmt.Printf("GetRobotList result: %v\n", success)
		if success {
			for _, robot := range robots {
				fmt.Printf("Robot ID: %s, Online: %v\n", robot.RobotID, robot.IsOnLine)
			}
		} else {
			fmt.Println("Get Robot List Failed")
		}

		// Get specific robot state
		success, state := manager.GetRobotState("xxxxxxxxxxxx")
		fmt.Printf("GetRobotState result: %v, State: %+v\n", success, state)
		return
	}
	t.Error("TestAxToken failed")
}

func TestAxTask(t *testing.T) {

	config := &Config{
		URLPrefix:     os.Getenv("URL_PREFIX"),
		APPID:         os.Getenv("APP_ID"),
		APPSecret:     os.Getenv("APP_SECRET"),
		Authorization: "APPCODE " + os.Getenv("Authorization"),
		RobotID:       os.Getenv("RobotID"),
	}

	tokenManager := NewTokenManager()
	success, token := tokenManager.GetToken(config)
	if !success {
		t.Error("Failed to get token")
		return
	}
	// Example usage
	poi1 := POI{
		AreaID:     "66ea87fe6cb0037e92ba0ac4",
		Coordinate: []float64{-0.22222543918815063, 1.6403502840489637},
		Name:       "m1",
		Yaw:        0,
	}

	poi2 := POI{
		AreaID:     "66ea87fe6cb0037e92ba0ac4",
		Coordinate: []float64{-0.16790582975545476, 3.853874768537935},
		Name:       "m2",
		Yaw:        0,
	}

	task := NewTaskBuilder("Task1", config.RobotID)

	tp1 := NewTaskPoint(poi1, true)
	task.AddTaskPt(tp1)

	tp2 := NewTaskPoint(poi2, true)
	tp2.AddStepActs(Action.PlayAudioAction("3111002")).
		AddStepActs(Action.PauseAction(10)).
		AddStepActs(Action.PlayAudioAction("3111012"))
	task.AddTaskPt(tp2)

	park := NewTaskPoint(poi1, true)
	park.AddStepActs(Action.WaitAction(map[string]string{"cmd": "test"}))
	task.SetBackPt(park)

	fmt.Printf("Task: %+v\n", task.GetTask())

	manager := NewTaskManager(token, config.URLPrefix)

	ok, taskID := manager.NewTask(task.GetTask())
	fmt.Printf("New task created: %v, ID: %s\n", ok, taskID)

	if ok {
		ok = manager.ExecuteTask(taskID)
		if ok {
			// for {
			// 	time.Sleep(time.Second)
			// 	ok, data := manager.GetTaskInfo(taskID)
			// 	if ok {
			// 		fmt.Printf("isCancel:%v isFinish:%v isExcute:%v\n",
			// 			data["isCancel"], data["isFinish"], data["isExcute"])
			// 	} else {
			// 		break
			// 	}
			// }

			t.Log("TestAxToken passed")
			return
		}
	}

	t.Error("TestAxToken failed")
}
