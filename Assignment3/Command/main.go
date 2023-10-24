package main

func main() {
	car := &Car{}

	startEngineCommand := &StartEngineCommand{car}
	stopEngineCommand := &StopEngineCommand{car}
	unlockDoorsCommand := &UnlockDoorsCommand{car}
	lockDoorsCommand := &LockDoorsCommand{car}

	startButton := &Button{startEngineCommand}
	stopButton := &Button{stopEngineCommand}
	lockButton := &Button{lockDoorsCommand}
	unlockButton := &Button{unlockDoorsCommand}

	lockButton.press()
	startButton.press()
	stopButton.press()
	unlockButton.press()
}
