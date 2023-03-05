package controllers

import (
	"net-flashbang/models"
	"net-flashbang/models/ping"
	"net-flashbang/views/console"
)

const CommandPing string = "1"

// Run does the running of the console application
func Run(enablePersistence bool) {
	if enablePersistence {
		models.EnableFilePersistence()
	} else {
		models.DisableFilePersistence()
	}

	/*err := models.Initialize()
	checkAndHandleErrorWithTermination(err)*/

	console.Clear()
	console.PrintMenu()

	for true {
		executeCommand()
	}
}

/*unc checkAndHandleErrorWithTermination(err error) {
	if err != nil {
		console.PrintError(err)
		log.Fatal(err)
	}
}*/

func executeCommand() {
	command := console.AskForInput()
	parseAndExecuteCommand(command)
}

func parseAndExecuteCommand(input string) {
	switch {
	case input == CommandPing:
		console.Clear()
		NewIpAddress := console.AskForIP()
		ping.Ping(NewIpAddress)
	}
}
