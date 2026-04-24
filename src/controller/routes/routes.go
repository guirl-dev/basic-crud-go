package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guirl-dev/basic-crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/user/getuserbyid/:userId", controller.FindUserById)
	r.GET("/user/getuserbyemail/:userEmail", controller.FindUserByEmail)
	r.POST("/user/createuser", controller.CreateUser)
	r.PUT("/user/updateuser/:userId", controller.UpdateUser)
	r.DELETE("/user/deleteuser/:userId", controller.DeleteUser)
}
