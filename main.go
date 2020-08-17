package main

import(
	"github.com/IamFrost/test5/pkg/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	// utils.CreateConnection()
	router := gin.Default()

	routes.GetProcedureRoutes(router)

	router.Use(cors.Default())
	router.Run(":3000")

}