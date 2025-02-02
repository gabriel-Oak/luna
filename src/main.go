package main

import (
	"luna/src/modules/listening"
	"luna/src/modules/thinking"
	"luna/src/utils"
)

func main() {
	command := listening.ListenForCommand()
	utils.Debug("[command]", command)
	thinking.ThinkAndRespond(command)
}
