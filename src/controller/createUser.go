package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/guirl-dev/basic-crud-go/src/config/restErr"
	"github.com/guirl-dev/basic-crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		rErr := restErr.NewBadRequestError(
			fmt.Sprintf("invalid json body: %s", err.Error()))

		c.JSON(rErr.Code, rErr)
		return
	}

	fmt.Println(userRequest)
}
