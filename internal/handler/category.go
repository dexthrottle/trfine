package handler

import (
	"net/http"
	"strconv"

	"github.com/dexthrottle/trfine/internal/dto"
	"github.com/dexthrottle/trfine/internal/helper"
	"github.com/dexthrottle/trfine/internal/model"
	"github.com/gin-gonic/gin"
)

func (c *Handler) AllCategory(ctx *gin.Context) {
	userTgId := ctx.GetHeader("user_tg_id")
	convUserTgId, err := strconv.Atoi(userTgId)
	if err != nil {
		c.log.Error(err)
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	categories, err := c.service.Category.All(ctx, convUserTgId)
	if err != nil {
		c.log.Errorf("get all categories error: %v", err)
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", categories)
	ctx.JSON(http.StatusOK, res)
}

func (c *Handler) InsertCategory(ctx *gin.Context) {
	var categoryCreateDTO dto.CreateCategoryDTO
	errDTO := ctx.ShouldBind(&categoryCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	} else {
		user, err := c.service.FindUserByTgUserId(ctx, categoryCreateDTO.UserTgId)
		if err != nil {
			c.log.Errorf("FindUserByTgUserId category error: %v", err)
			res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusConflict, res)
			return
		}

		categoryCreateDTO.UserID = user.ID
		c.log.Infof("%+v", categoryCreateDTO)
		category, err := c.service.Category.Insert(ctx, categoryCreateDTO)
		if err != nil {
			c.log.Errorf("insert category error: %v", err)
			res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusConflict, res)
			return
		}
		response := helper.BuildResponse(true, "OK", category)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *Handler) DeleteCategory(ctx *gin.Context) {
	var category model.Category

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	category.ID = id

	userTgID := ctx.GetHeader("user_tg_id")
	convUserTgId, err := strconv.Atoi(userTgID)
	if err != nil {
		c.log.Errorf("is allowed to edit error: %v", err)
		response := helper.BuildErrorResponse("is allowed to edit error", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	user, err := c.service.FindUserByTgUserId(ctx, convUserTgId)
	c.log.Printf("is handler %+v", category)
	if err != nil {
		c.log.Errorf("is allowed to edit error: %v", err)
		response := helper.BuildErrorResponse("is allowed to edit error", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = c.service.Category.Delete(ctx, category, int(user.ID))
	if err != nil {
		c.log.Errorf("delete category error: %v", err)
		response := helper.BuildErrorResponse("delete category error", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}

	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusNoContent, res)

}
