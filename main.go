package main

import (
	_ "sharit-backend/routers"

	"github.com/astaxie/beego"
)

func main() {
	// CORS for https://foo.* origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// // - Credentials share
	//chat.Run()
	beego.Run()

}
