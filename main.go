package main

import (
	datatype "github.com/Scrummyy/scrummyy-api/internal/datatypes"
	"github.com/Scrummyy/scrummyy-api/server"

	"github.com/gin-gonic/gin/binding"
)

func main() {
	conf, err := server.InitConfig()
	if err != nil {
		panic(err)
	}

	// r := gin.Default()

	//Custom form validator
	binding.Validator = new(datatype.DefaultValidator)
	router := server.GetRouter(conf)
	server.RegisterRoutes(router, conf)
	server.Serve(router, conf)

}
