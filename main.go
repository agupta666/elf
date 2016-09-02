package main

func main() {
	go startDefaultEp("0.0.0.0:8080")
	startShell()
}
