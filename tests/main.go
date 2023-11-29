package main

import dev "github.com/pjsoftware/go-dev"

func main() {
	dl := dev.InitLogging("./tests/log/tests.log")
	dl.EnableLogging()

	dl.Print("Logging output enabled")
	dl.Printf("Printf also enabled: %d, %d, %d", 1,2,3)

	dl.DisableLogging()
	dl.Print("This message should not appear!")
}