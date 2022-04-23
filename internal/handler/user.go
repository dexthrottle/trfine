package handler

import (
	"net/http"

	"github.com/dexthrottle/trfine/internal/dto"
	"github.com/dexthrottle/trfine/internal/helper"
	"github.com/gin-gonic/gin"
)

func (c *Handler) CreateUser(ctx *gin.Context) {
	var createDTO dto.CreateUserDTO
	errDTO := ctx.ShouldBind(&createDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userStatus, err := c.service.User.IsDuplicateUserTGID(ctx, createDTO.UserTGId)
	c.log.Errorln(err)
	if err == nil || userStatus {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate user_tg_id", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	} else {
		createdUser, err := c.service.User.Insert(ctx, createDTO)
		if err != nil {
			c.log.Errorf("create user failed: %v", err)
			response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusConflict, response)
			return
		}
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *Handler) ProfileUser(ctx *gin.Context) {

	id := ctx.GetHeader("user_tg_id")
	user, err := c.service.User.Profile(ctx, id)
	if err != nil {
		c.log.Errorf("profile user error: %v", err)
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	if user.UserTGId == 0 {
		response := helper.BuildErrorResponse("Failed to process request", "Такого пользователя нет", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	res := helper.BuildResponse(true, "OK", user)
	ctx.JSON(http.StatusOK, res)

}
