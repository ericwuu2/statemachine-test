package utils

import (
	"fmt"
	"github.com/Gurpartap/statemachine-go"
)

type GameProcess struct {
	statemachine.Machine
	IsLogin   bool
	IsConnect bool
}

func (gameProcess *GameProcess) NotifyTriggers() {
	fmt.Println("NotifyTriggers")
}
func (gameProcess *GameProcess) GetIsLogin() bool {
	return gameProcess.IsLogin
}
func (gameProcess *GameProcess) GetIsConnect() bool {
	return gameProcess.IsConnect
}
func (gameProcess *GameProcess) TestTrue() bool {
	return true
}
func (gameProcess *GameProcess) TestFalse() bool {
	return false
}
func (gameProcess *GameProcess) Echo() {
	fmt.Println("Hello I'm GameProcess")
}
