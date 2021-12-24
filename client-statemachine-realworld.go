package main

import (
	"fmt"
	"github.com/Gurpartap/statemachine-go"
	statemachine_common "github.com/ericwuu2/statemachine-common"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// NewRealWorldClientGameProcess Structure
func NewRealWorldClientGameProcess() *statemachine_common.GameProcess {
	clientGameProcess := &statemachine_common.GameProcess{}
	clientGameProcess.Machine = statemachine.NewMachine()
	clientGameProcessMachineDef := &statemachine.MachineDef{
		States:       []string{"Launch", "MainMenu", "ChooseGame", "PlayGame", "Disconnect"},
		InitialState: "Launch",
		Events: map[string]*statemachine.EventDef{
			"AutoLoading": {
				Transitions: []*statemachine.TransitionDef{{
					From: []string{"Launch"},
					To:   "MainMenu",
				}},
			},
			"TapScreen": {
				Transitions: []*statemachine.TransitionDef{{
					From: []string{"MainMenu"},
					To:   "ChooseGame",
				}},
			},
			"SelectedGame": {
				Transitions: []*statemachine.TransitionDef{{
					From: []string{"ChooseGame"},
					To:   "PlayGame",
				}},
			},
			"GoBackMainMenu": {
				Transitions: []*statemachine.TransitionDef{{
					From: []string{"ChooseGame"},
					To:   "MainMenu",
				}},
			},
			"Heartbeat": {
				TimedEvery: 10 * time.Second,
				Transitions: []*statemachine.TransitionDef{
					{
						IfGuards: []*statemachine.TransitionGuardDef{{
							Label:          "",
							RegisteredFunc: "",
							Guard:          clientGameProcess.GetIsConnect,
						}},
					},
					{
						UnlessGuards: []*statemachine.TransitionGuardDef{{
							Label:          "",
							RegisteredFunc: "",
							Guard:          clientGameProcess.GetIsConnect,
						}},
						To: "Disconnect",
					},
				},
			},
		},
		BeforeCallbacks: []*statemachine.TransitionCallbackDef{
			{
				//ExitToState: "Disconnect",
				To: []string{"Disconnect"},
				Do: []*statemachine.TransitionCallbackFuncDef{{
					Func: func() {
						fmt.Println("[Fire] Heartbeat Event")
					},
				}},
			},
		},
		AfterCallbacks: []*statemachine.TransitionCallbackDef{
			{
				Do: []*statemachine.TransitionCallbackFuncDef{{
					Func: func() {
						fmt.Println("[gameProcess] Current state:", clientGameProcess.Machine.GetState())
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
	gameProcess := NewRealWorldClientGameProcess()
	gameProcess.IsConnect = true

	gameLoopChannel := make(chan string, 3)

	fmt.Println("===================")
	fmt.Println("[gameProcess] IsConnect: ", gameProcess.GetIsConnect())
	fmt.Println("[gameProcess] IsLogin: ", gameProcess.GetIsLogin())
	fmt.Println("[gameProcess] state:", gameProcess.GetState())

	fmt.Println("===================")

	time.AfterFunc(12*time.Second, func() {
		fmt.Println("[Error] 遊戲中斷連線了 !!")
		gameProcess.IsConnect = false
	})

	time.AfterFunc(30*time.Second, func() {
		fmt.Println("[Error] 遊戲連線恢復 !!")
		gameProcess.IsConnect = true
		gameLoopChannel <- "Restart"
	})

	//myTimer.Stop()

	if gameProcess.IsConnect {
		//RunningGameLoop(gameProcess)
		fmt.Println("===================")
		fmt.Println("[Fire] AutoLoading Event")
		gameProcess.Fire("AutoLoading")

		fmt.Println("===================")
		fmt.Println("[Fire] TapScreen Event")
		gameProcess.Fire("TapScreen")

		fmt.Println("===================")
		fmt.Println("[Fire] SelectedGame Event")
		gameProcess.Fire("SelectedGame")

		//fmt.Println("===================")
		//fmt.Println("[Fire] GoBackMainMenu Event")
		//gameProcess.Fire("GoBackMainMenu")
	}

	select {
	case event := <-gameLoopChannel:
		fmt.Println("[GameLoopChannel] Received Event:", event)
		if event == "Restart" {
			fmt.Println("[GameLoop] 遊戲恢復連線")

			fmt.Println("===================")
			fmt.Println("[gameProcess] SetCurrentState: Launch")
			gameProcess.SetCurrentState("Launch")

			fmt.Println("===================")
			fmt.Println("[Fire] AutoLoading Event")
			gameProcess.Fire("AutoLoading")

			fmt.Println("===================")
			fmt.Println("[Fire] TapScreen Event")
			gameProcess.Fire("TapScreen")

			fmt.Println("===================")
			fmt.Println("[Fire] SelectedGame Event")
			gameProcess.Fire("SelectedGame")

			//fmt.Println("===================")
			//fmt.Println("[Fire] GoBackMainMenu Event")
			//gameProcess.Fire("GoBackMainMenu")
		}
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
