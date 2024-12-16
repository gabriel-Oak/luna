package main

import (
	"luna/src/modules/listening"
	"luna/src/utils"
)

func main() {
	command := listening.ListenForCommand()
	utils.Debug("[command]", command)
}
