package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestActionType_PauseAction(t *testing.T) {
	type args struct {
		duration int
	}

	tests := []struct {
		name string
		a    ActionType
		args args
		want ActionType
	}{
		// TODO: Add test cases.
		{
			name: "Test PauseAction with duration 10",
			a:    ActionType{},
			args: args{duration: 10},
			want: ActionType{18, map[string]interface{}{"pauseTime": 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActionType{}
			if got := a.PauseAction(tt.args.duration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ActionType.PauseAction() = %v, want %v", got, tt.want)

			} else {
				jsonData, err := json.Marshal(got)
				if err != nil {
					fmt.Println("Error marshaling JSON:", err)

				}
				t.Logf("JSON: %v", string(jsonData))
			}
		})
	}
}

func TestActionType_LiftUp(t *testing.T) {
	type fields struct {
		Type int
		Data map[string]interface{}
	}
	type args struct {
		useAreaId *string
	}

	aid := "test"

	tests := []struct {
		name   string
		fields fields
		args   args
		want   ActionType
	}{
		// TODO: Add test cases.
		{name: "test for liftup",

			args: args{
				useAreaId: &aid,
			},
			want: ActionType{
				Type: 47,
				Data: map[string]interface{}{
					"useAreaId": aid,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActionType{}
			if got := a.LiftUp(tt.args.useAreaId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ActionType.LiftUp() = %v, want %v", got, tt.want)
			} else {
				jsonData, err := json.Marshal(got)
				if err != nil {
					fmt.Println("Error marshaling JSON:", err)
				}
				fmt.Println("JSON:", string(jsonData))
			}
		})
	}
}

func TestActionType_LiftDown(t *testing.T) {
	type fields struct {
		Type int
		Data map[string]interface{}
	}
	type args struct {
		useAreaId *string
	}

	aid := "test"

	tests := []struct {
		name   string
		fields fields
		args   args
		want   ActionType
	}{
		// TODO: Add test cases.
		{name: "test for liftdown",

			args: args{
				useAreaId: &aid,
			},
			want: ActionType{
				Type: 48,
				Data: map[string]interface{}{
					"useAreaId": aid,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActionType{
				Type: tt.fields.Type,
				Data: tt.fields.Data,
			}
			if got := a.LiftDown(tt.args.useAreaId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ActionType.LiftDown() = %v, want %v", got, tt.want)
			} else {
				jsonData, err := json.Marshal(got)
				if err != nil {
					fmt.Println("Error marshaling JSON:", err)
				}
				fmt.Println("JSON:", string(jsonData))
			}
		})

	}
}
