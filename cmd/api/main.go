package main

import "postapp/pkg/server"

// @title post App Majo Test APIs
// @version 1.0
// @description Only Test Purpose, If you Not Authorized, Please Login and input into security header with format Bearer token

// @contact.name API Support
// @contact.email zona.budi11@gmail.com

//@securityDefinitions.apikey Bearer
//@in header
//@name Authorization

// @BasePath /
func main() {
	server.Run()
}
