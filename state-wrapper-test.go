package main

import (
	"fmt"
	"github.com/Gurpartap/statemachine-go"
)

type XStateMachine struct{}

type XStateMachineDef struct {
	Id           string
	InitialState string
	States       map[string]StateTransition
	//States []map[string]string
}

func (xStateMachineDef *XStateMachineDef) Transition() {

}

type StateName = string

type EventName = string

type StateTransition struct {
	On map[EventName]TargetStateEventTransition
}

type TargetStateEventTransition struct {
	target string
}

func (xStateMachine *XStateMachine) SetMachineDef(xStateMachineDef *XStateMachineDef) statemachine.Machine {
	fmt.Printf("%+v\n", xStateMachineDef)

	machine := statemachine.BuildNewMachine(func(machineBuilder statemachine.MachineBuilder) {
		// InitialState
		machineBuilder.InitialState(xStateMachineDef.InitialState)

		var stateKeys []string
		for stateKey, stateTransition := range xStateMachineDef.States {
			fmt.Println("stateKey:", stateKey, "=>", "stateTransition:", stateTransition)
			// States
			stateKeys = append(stateKeys, stateKey)
			for eventKey, targetState := range stateTransition.On {
				fmt.Println("eventKey:", eventKey, "=>", "targetState:", targetState.target)
				// Event Transition
				machineBuilder.Event(eventKey).Transition().From(stateKey).To(targetState.target)
			}
		}
		fmt.Println("stateKeys: ", stateKeys)
		machineBuilder.States(stateKeys...)
	})
	return machine
}

type GameProcess struct {
	XStateMachine
	IsLogin      bool
	IsDisconnect bool
}

func (gameProcess *GameProcess) NotifyTriggers() {
	fmt.Println("NotifyTriggers")
}
func (gameProcess *GameProcess) GetIsLogin() bool {
	return gameProcess.IsLogin
}

func GenerateGameProcess() *GameProcess {
	gameProcess := &GameProcess{}
	gameProcess.SetMachineDef(&XStateMachineDef{
		Id:           "mainState",
		InitialState: "loading",
		States: map[StateName]StateTransition{
			"loading": {
				On: map[EventName]TargetStateEventTransition{
					"load": {
						target: "ready",
					},
				},
			},
			"ready": {
				On: map[EventName]TargetStateEventTransition{
					"load": {
						target: "starting",
					},
				},
			},
			"starting": {
				On: map[EventName]TargetStateEventTransition{
					"spin": {
						target: "playing",
					},
					"back": {
						target: "ready",
					},
				},
			},
			"playing": {
				On: map[EventName]TargetStateEventTransition{
					"end": {
						target: "finished",
					},
					"stop": {
						target: "starting",
					},
				},
			},
			"finished": {
				On: map[EventName]TargetStateEventTransition{
					"end": {
						target: "restart",
					},
					"stop": {
						target: "starting",
					},
				},
			},
		},
	})
	//stateMachine :=
	//stateMachine.Transition('TIMER')
	//return gameProcess
}

func main() {
	gameProcess := GenerateGameProcess()
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
