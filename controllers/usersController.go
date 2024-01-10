package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smonkeymonkey/marathon_runners/metrics"
	"github.com/smonkeymonkey/marathon_runners/services"
)

type UsersController struct {
	usersService *services.UsersService
}

func NewUsersController(usersService *services.UsersService) *UsersController {
	return &UsersController{
		usersService: usersService,
	}
}
func (uc UsersController) Login(ctx *gin.Context) {
	metrics.HttpRequestsCounter.Inc()

	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		log.Println("Error while reading creditials")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	accessToken, responseErr := uc.usersService.Login(username, password)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, accessToken)
}
func (uc UsersController) Logout(ctx *gin.Context) {
	metrics.HttpRequestsCounter.Inc()

	accessToken := ctx.Request.Header.Get("Token")
	responseErr := uc.usersService.Logout(accessToken)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}
