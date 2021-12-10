package main

import (
	"fmt"
	"github.com/Gurpartap/statemachine-go"
	"github.com/ericwuu2/statemachine-common"
	"time"
)

func NewGameProcess() *statemachine_common.GameProcess {
	gameProcess := &statemachine_common.GameProcess{}

	gameProcess.Machine = statemachine.BuildNewMachine(func(machineBuilder statemachine.MachineBuilder) {
		machineBuilder.States("loading", "ready", "starting", "spinning", "finished")
		machineBuilder.InitialState("loading")

		//machineBuilder.Event("ready").Transition().From("loading").To("ready").If(&gameProcess.IsLogin).AndUnless(&gameProcess.IsDisconnect).Label("ReadyLabel")
		machineBuilder.Event("ready").
			Choice(&gameProcess.IsLogin).
			Label("READY").
			Label("LoadingToReady").
			Unless(&gameProcess.IsDisconnect).
			OnTrue(func(eventBuilder statemachine.EventBuilder) {
				eventBuilder.Transition().From("loading").To("ready")
			}).
			OnFalse(func(eventBuilder statemachine.EventBuilder) {
				eventBuilder.Transition().From("loading").To("finished")
			})

		machineBuilder.Event("start").Transition().From("ready").To("starting").If(func() bool {
			return true
		})
		machineBuilder.Event("play").Transition().From("starting").To("spinning")
		machineBuilder.Event("end").Transition().From("spinning").To("finished")
		machineBuilder.Event("tiktik").
			TimedEvery(5 * time.Second).
			Transition().
			From("finished").
			To("ready")

		machineBuilder.AfterFailure().OnAnyEvent().Do(func(e statemachine.Event, err error) {
			fmt.Println(e.Event())
			fmt.Println(err)
			//log.Println("could not transition with event='%s' err=%+v\n",
			//	e.Event(),
			//	err)
		})

		machineBuilder.BeforeTransition().Any().Do(gameProcess.NotifyTriggers)

		machineBuilder.BeforeTransition().From("loading").To("ready").Do(func() {
			fmt.Println("[Before] Loading state to Ready state")
		})
		machineBuilder.AroundTransition().From("loading").To("ready").Do(func(next func()) {
			fmt.Println("[Around] Loading state to Ready state")
			next()
		})
		machineBuilder.AfterTransition().From("loading").To("ready").Do(func() {
			fmt.Println("[After] Loading state to Ready state")
		})
	})
	return gameProcess
}

func main() {
	gameProcess := NewGameProcess()
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
