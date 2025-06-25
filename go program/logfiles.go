package main

import (
	"os"
)

func createFile(name string, content string) {
	os.WriteFile(name, []byte(content), 0644)
}

func main() {
	createFile("server1.log", `INFO: Server started successfully
ERROR: Failed to connect to database
INFO: Listening on port 8080
ERROR: Disk space low`)

	createFile("server2.log", `INFO: Health check passed
ERROR: Could not load configuration file
INFO: Request processed
INFO: Shutdown signal received
`)

	createFile("server3.log", `ERROR: Timeout occurred while processing request
INFO: Restart completed
ERROR: Unauthorized access attempt detected
INFO: Cleanup done
`)
}
