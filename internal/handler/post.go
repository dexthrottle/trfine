package handler

import (
	"net/http"
	"strconv"

	"github.com/dexthrottle/trfine/internal/dto"
	"github.com/dexthrottle/trfine/internal/helper"
	"github.com/dexthrottle/trfine/internal/model"
	"github.com/gin-gonic/gin"
)

func (c *Handler) AllPost(ctx *gin.Context) {
	userTgId := ctx.GetHeader("user_tg_id")
	convUserTgId, err := strconv.Atoi(userTgId)
	if err != nil {
		c.log.Error(err)
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, res)
	}

	posts, err := c.service.Post.All(ctx, convUserTgId)
	if err != nil {
		c.log.Errorf("get all posts error: %v", err)
		response := helper.BuildErrorResponse("get all posts error", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res := helper.BuildResponse(true, "OK", posts)
	ctx.JSON(http.StatusOK, res)
}

// Создание Post
func (c *Handler) InsertPost(ctx *gin.Context) {
	var postCreateDTO dto.PostCreateDTO
	errDTO := ctx.ShouldBind(&postCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse(
			"Failed to process request",
			errDTO.Error(),
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	} else {
		user, err := c.service.FindUserByTgUserId(ctx, postCreateDTO.UserTgId)
		if err != nil {
			c.log.Errorf("FindUserByTgUserId category error: %v", err)
			res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusForbidden, res)
			return
		}
		postCreateDTO.UserID = user.ID

		post, err := c.service.Post.Insert(ctx, postCreateDTO)
		if err != nil {
			c.log.Errorf("insert post error: %v", err)
			res := helper.BuildErrorResponse("insert post error", err.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		response := helper.BuildResponse(true, "OK", post)
		ctx.JSON(http.StatusCreated, response)
	}
}

// Удаление Post
func (c *Handler) DeletePost(ctx *gin.Context) {
	var post model.Post

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse(
			"Failed tou get id",
			"No param id were found",
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	post.ID = id
	userTgID := ctx.GetHeader("user_tg_id")
	convUserTgId, err := strconv.Atoi(userTgID)
	if err != nil {
		c.log.Errorf("is allowed to edit error: %v", err)
		response := helper.BuildErrorResponse("is allowed to edit error", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	user, err := c.service.FindUserByTgUserId(ctx, convUserTgId)
	if err != nil {
		c.log.Errorf("is allowed to edit error: %v", err)
		response := helper.BuildErrorResponse(
			"You dont have permission",
			"You are not the owner",
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusForbidden, response)
		return
	}
	c.service.Post.Delete(ctx, post, int(user.ID))
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusNoContent, res)

}
