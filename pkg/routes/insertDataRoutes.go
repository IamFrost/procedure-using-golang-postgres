package routes

import (
	"github.com/IamFrost/test5/pkg/middlewares/procedure"

	"github.com/gin-gonic/gin"
)

// GetProcedureRoutes returns routes for executing procedure
func GetProcedureRoutes(router *gin.Engine) *gin.Engine {

	procedureRoutes := router.Group("/procedure")
	{
		functionRoutes := procedureRoutes.Group("/functions")
		{
			// getRoutes := areaRoutes.Group("/get")
			// {
			// 	getRoutes.GET("/all", procedure.InsertDataProCall)
			// 	// getRoutes.GET("/maxareacode", advarea.GetMaxAdvertiserAreaCode)
			// 	// getRoutes.GET("/one/:code", advarea.GetOneAdvertiserInfo)
			// }
			postRoutes := functionRoutes.Group("/post")
			{
				//postRoutes.POST("/one", advarea.CreateOneAdvertiserArea)
				postRoutes.POST("/inserttwo/:integer1/:integer2", procedure.InsertDataProCall)
			}

			// putRoutes := areaRoutes.Group("/put")
			// {
			// 	putRoutes.PUT("one/:code", advarea.UpdateOneAdvertiserArea)
			// }

			// deleteRoutes := areaRoutes.Group("/delete")
			// {
			// 	deleteRoutes.DELETE("/one/:code", advarea.DeleteOneAdvertiserArea)
			// }
		}
	}

	return router
}