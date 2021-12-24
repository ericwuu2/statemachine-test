package main

import (
	"fmt"
	"github.com/Gurpartap/statemachine-go"
	"github.com/ericwuu2/statemachine-common"
	"time"
)

func NewGameProcessWithStruct() *statemachine_common.GameProcess {
	gameProcess := &statemachine_common.GameProcess{}
	gameProcess.Machine = statemachine.NewMachine()
	gameProcessMachineDef := &statemachine.MachineDef{
		States:       []string{"loading", "ready", "starting", "spinning", "finished"},
		InitialState: "loading",
		Events: map[string]*statemachine.EventDef{
			"ready": {
				Transitions: []*statemachine.TransitionDef{{
					From: []string{"loading"},
					To:   "ready",
					//IfGuards: []*statemachine.TransitionGuardDef{{
					//	Label: "READY",
					//}},
					//UnlessGuards: []*statemachine.TransitionGuardDef{{
					//	Label: "UnlessGuards",
					//}},
				}},
			},
			"start": {

				Transitions: []*statemachine.TransitionDef{{
					From: []string{"ready"},
					To:   "starting",
				}},

				Choice: &statemachine.ChoiceDef{
					Condition: &statemachine.ChoiceConditionDef{
						Label:          "",
						RegisteredFunc: "",
						Condition:      nil,
					},
					UnlessGuard: &statemachine.TransitionGuardDef{
						Label:          "",
						RegisteredFunc: "",
						Guard:          nil,
					},
					OnTrue: &statemachine.EventDef{
						//TimedEvery:  0,
						//Choice:      nil,
						Transitions: nil,
					},
					OnFalse: &statemachine.EventDef{
						//TimedEvery:  0,
						//Choice:      nil,
						Transitions: nil,
					},
				},
			},
			"play": {
				Transitions: []*statemachine.TransitionDef{{
					From: []string{"starting"},
					To:   "spinning",
				}},
			},
			"end": {
				Transitions: []*statemachine.TransitionDef{{
					From: []string{"spinning"},
					To:   "finished",
				}},
			},
			"tick": {
				TimedEvery: 5 * time.Second,
				Transitions: []*statemachine.TransitionDef{{
					From: []string{"finished"},
					To:   "ready",
				}},
			},
		},
		BeforeCallbacks: []*statemachine.TransitionCallbackDef{
			{
				From: []string{"loading"},
				To:   []string{"ready"},
				Do: []*statemachine.TransitionCallbackFuncDef{{
					Func: func() {
						fmt.Println("[Before] Loading state to Ready state")
					},
				}},
			},
			{
				From: []string{"ready"},
				To:   []string{"starting"},
				Do: []*statemachine.TransitionCallbackFuncDef{{
					Func: func() {
						fmt.Println("[Before] Ready state to Starting state")
					},
				}},
			},
		},
		AfterCallbacks: []*statemachine.TransitionCallbackDef{
			{
				Do: []*statemachine.TransitionCallbackFuncDef{{
					Func: func() {
						//fmt.Println("AfterCallbacks")
					},
				}},
			},
		},
		FailureCallbacks: []*statemachine.EventCallbackDef{
			{
				Do: []*statemachine.EventCallbackFuncDef{{
					Func: func() {
						fmt.Println("NotifyTriggers")
					},
				}},
			},
		},
	}

	gameProcess.SetMachineDef(gameProcessMachineDef)

	return gameProcess
}

func main() {
	gameProcess := NewGameProcessWithStruct()
	fmt.Println("IsLogin", gameProcess.IsLogin)
	gameProcess.IsLogin = true
	fmt.Println("IsLogin", gameProcess.GetIsLogin())

	fmt.Println("===================")
	fmt.Println("[gameProcess] state:", gameProcess.GetState())

	fmt.Println("===================")
	fmt.Println("[Fire] ready")
	gameProcess.Fire("ready")
	fmt.Println("[gameProcess] state:", gameProcess.GetState())

	fmt.Println("===================")
	fmt.Println("[Fire] start")
	gameProcess.Fire("start")
	fmt.Println("[gameProcess] state:", gameProcess.GetState())

	fmt.Println("===================")
	fmt.Println("[Fire] play")
	gameProcess.Fire("play")
	fmt.Println("[gameProcess] state:", gameProcess.GetState())

	fmt.Println("===================")
	fmt.Println("[Fire] end")
	gameProcess.Fire("end")
	fmt.Println("[gameProcess] state:", gameProcess.GetState())

}
