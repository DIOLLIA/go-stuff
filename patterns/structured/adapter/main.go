package main

func main() {

	client := Client{}

	mac := Mac{}
	client.InsertLightningConnectorIntoComputer(&mac)

	win := Windows{}
	adaptedWin := WindowsAdapter{windowMachine: &win}
	client.InsertLightningConnectorIntoComputer(&adaptedWin)
}
