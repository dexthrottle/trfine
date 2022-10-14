package http

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// const (
// 	authorizationHeader = "Authorization"
// 	userCtx             = "username"
// )

// // func (h *Handler) FetchAuth(ctx gin.Context) (bool, error) {
// // 	cookie, err := ctx.Cookie(usernameCookies)
// // 	if err != nil {
// // 		return false, err
// // 	}
// // 	token, err := h.redisService.GetRefreshToken(&ctx, cookie)
// // 	if err != nil || token == "" {
// // 		return false, err
// // 	}

// // 	return true, nil
// // }

// func (h *Handler) DeserializeUser(ctx *gin.Context) {
// 	header := ctx.GetHeader(authorizationHeader)
// 	if header == "" {
// 		builErrorResponse(ctx, http.StatusUnauthorized, Response{
// 			Status:  statusError,
// 			Message: "empty auth header",
// 			Data:    nil,
// 		})
// 		return
// 	}

// 	headerParts := strings.Split(header, " ")
// 	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
// 		builErrorResponse(ctx, http.StatusBadRequest, Response{
// 			Status:  statusError,
// 			Message: "invalid auth header",
// 			Data:    nil,
// 		})
// 		return
// 	}

// 	if len(headerParts[1]) == 0 {
// 		builErrorResponse(ctx, http.StatusBadRequest, Response{
// 			Status:  statusError,
// 			Message: "token is empty",
// 			Data:    nil,
// 		})
// 		return
// 	}

// 	username, err := h.jwt.ValidateToken(headerParts[1], h.cfg.AppConfig.JWTToken.JwtAccessKey)
// 	if err != nil {
// 		builErrorResponse(ctx, http.StatusUnauthorized, Response{
// 			Status:  statusError,
// 			Message: "unauthorize",
// 			Data:    nil,
// 		})
// 		return
// 	}

// 	ctx.Set(userCtx, username)
// }
