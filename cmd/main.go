package main

import "github.com/kyimmQ/ielts-writing-golang/internal/server"

//	@title			IeltS API
//	@version		1.0
//	@description	IeltS server APIs.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api

//	@securityDefinitions.basic	BasicAuth
func main() {
	server := server.InitServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}

}
