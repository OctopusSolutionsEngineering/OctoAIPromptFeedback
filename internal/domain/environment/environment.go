package environment

import "os"

// GetPort gets the port from the environment variable
func GetPort() string {
	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if port == "" {
		port = "8080" // Default port
	}
	return port
}
