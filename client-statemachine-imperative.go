package main

import (
	"fmt"
	"github.com/Gurpartap/statemachine-go"
	statemachine_common "github.com/ericwuu2/statemachine-common"
)

// NewClientGameProcess ImperativeParadigm
func NewClientGameProcess() *statemachine_common.GameProcess {
	clientGameProcess := &statemachine_common.GameProcess{}

	clientGameProcess.Machine = statemachine.BuildNewMachine(func(machineBuilder statemachine.MachineBuilder) {
		machineBuilder.States("Launch", "MainMenu", "ChooseGame", "PlayGame")
		machineBuilder.InitialState("Launch")

		machineBuilder.Event("Launch_Next").Transition().From("Launch").To("MainMenu")
		machineBuilder.Event("MainMenu_Next").Transition().From("MainMenu").To("ChooseGame")
		machineBuilder.Event("ChooseGame_Next").Transition().From("ChooseGame").To("PlayGame")
		machineBuilder.Event("ChooseGame_Previous").Transition().From("ChooseGame").To("MainMenu")
		machineBuilder.Event("PlayGame_Previous").Transition().From("PlayGame").To("ChooseGame")

		machineBuilder.AfterTransition().FromAny().ToAny().Do(func() {
			fmt.Println("[gameProcess] state:", clientGameProcess.Machine.GetState())
		})
		machineBuilder.AfterFailure().OnAnyEvent().Do(func(e statemachine.Event, err error) {
			fmt.Println(e.Event())
			fmt.Println(err)
		})

	})
	return clientGameProcess
}

func main() {
	gameProcess := NewClientGameProcess()
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
