package main

func main() {
	for i := 0; i < 1; i++ {
		createClientsSingleRoutine()
		createClientsMultipleRoutines()
		createClientWithChannels()
	}
}
