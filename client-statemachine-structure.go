package main

import (
	"fmt"
	"github.com/Gurpartap/statemachine-go"
	statemachine_common "github.com/ericwuu2/statemachine-common"
	"time"
)

// NewStructureClientGameProcess Structure
func NewStructureClientGameProcess() *statemachine_common.GameProcess {
	clientGameProcess := &statemachine_common.GameProcess{}
	clientGameProcess.Machine = statemachine.NewMachine()
	clientGameProcessMachineDef := &statemachine.MachineDef{
		States:       []string{"Launch", "MainMenu", "ChooseGame", "PlayGame"},
		InitialState: "Launch",
		Events: map[string]*statemachine.EventDef{
			"Launch_Next": {
				Transitions: []*statemachine.TransitionDef{{
					From: []string{"Launch"},
					To:   "MainMenu",
				}},
			},
			"MainMenu_Next": {
				Transitions: []*statemachine.TransitionDef{{
					From: []string{"MainMenu"},
					To:   "ChooseGame",
				}},
			},
			"ChooseGame_Next": {
				Transitions: []*statemachine.TransitionDef{{
					From: []string{"ChooseGame"},
					To:   "PlayGame",
				}},
			},
			"ChooseGame_Previous": {
				Transitions: []*statemachine.TransitionDef{{
					From: []string{"ChooseGame"},
					To:   "MainMenu",
				}},
			},
			"PlayGame_Previous": {
				TimedEvery: 5 * time.Second,
				Transitions: []*statemachine.TransitionDef{{
					From: []string{"PlayGame"},
					To:   "ChooseGame",
				}},
			},
		},
		AfterCallbacks: []*statemachine.TransitionCallbackDef{
			{
				Do: []*statemachine.TransitionCallbackFuncDef{{
					Func: func() {
						fmt.Println("[gameProcess] state:", clientGameProcess.Machine.GetState())
					},
				}},
			},
		},
		FailureCallbacks: []*statemachine.EventCallbackDef{
			{
				Do: []*statemachine.EventCallbackFuncDef{{
					Func: func(e statemachine.Event, err error) {
						fmt.Println(e.Event())
						fmt.Println(err)
					},
				}},
			},
		},
	}
	clientGameProcess.SetMachineDef(clientGameProcessMachineDef)

	return clientGameProcess
}

func main() {
	gameProcess := NewStructureClientGameProcess()
	fmt.Println("===================")
	fmt.Println("[gameProcess] state:", gameProcess.GetState())

	fmt.Println("===================")
	fmt.Println("[Fire] Launch_Next Event")
	gameProcess.Fire("Launch_Next")

	fmt.Println("===================")
	fmt.Println("[Fire] MainMenu_Next Event")
	gameProcess.Fire("MainMenu_Next")

	fmt.Println("===================")
	fmt.Println("[Fire] ChooseGame_Next Event")
	gameProcess.Fire("ChooseGame_Next")

	fmt.Println("===================")
	fmt.Println("[Fire] PlayGame_Previous Event")
	gameProcess.Fire("PlayGame_Previous")

	fmt.Println("===================")
	fmt.Println("[Fire] ChooseGame_Previous Event")
	gameProcess.Fire("ChooseGame_Previous")

	fmt.Println("===================")
	fmt.Println("[Fire] Other State Event")
	gameProcess.Fire("ChooseGame_Next")

	fmt.Println("===================")
	fmt.Println("[Fire] Unknown Event")
	gameProcess.Fire("Unknown Event")

}
