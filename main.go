package main

// Web application which generates random lotto values: 6 out of 49.
// Based on http://random.org micro service.

import ()

func main() {
	initLogger()
	loadConfig()
	startServer()
}
