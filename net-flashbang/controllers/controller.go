package controllers

import (
	"log"
	"net-flashbang/models"
	"net-flashbang/view/console"
)

const CommandTest string = "1"

// Run does the running of the console application
func Run(enablePersistence bool) {
	if enablePersistence {
		models.EnableFilePersistence()
	} else {
		models.DisableFilePersistence()
	}

	err := models.Initialize()
	checkAndHandleErrorWithTermination(err)

	console.Clear()
	console.PrintMenu()

	for true {
		executeCommand()
	}
}

func checkAndHandleErrorWithTermination(err error) {
	if err != nil {
		console.PrintError(err)
		log.Fatal(err)
	}
}

func executeCommand() {
	command := console.AskForInput()
	parseAndExecuteCommand(command)
}

func parseAndExecuteCommand(input string) {
	switch {
	case input == CommandTest:
		console.Test()
		break

	default:
		console.PrintMessage("Command not defined. Check menu (c) for available commands.")
	}
}
