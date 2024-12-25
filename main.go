package main

import (
	_ "github.com/sms2sakthivel/user-manager/docs"
	"github.com/sms2sakthivel/user-manager/users"
)

func main() {
	// Step 1: Create a New User Service Application
	app := users.NewApp()

	// Step 2: Start Server and Listen on the Port 8001
	app.Listen(":8001")
}
