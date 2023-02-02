package middleware

import (
	"LibraryManagementSys/constant"
	"LibraryManagementSys/global"
	"LibraryManagementSys/model"
	"LibraryManagementSys/model/response"
	"LibraryManagementSys/utils"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"net/http"
)

func JwtAuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/*		if ctx.FullPath() == global.LMS_CONFIG.System.RouterPrefix+"/login" {
				ctx.Next()
			} else {*/
		token := ctx.GetHeader(constant.HeaderAuthorization)
		if utils.IsEmpty(token) {
			response.FailWithMessage("您无权访问，请登录！", ctx)
			ctx.Abort()
			return
		}
		token = token[7:]
		rDb := global.LMS_REDIS
		// redis验证
		res, err := rDb.Get(context.TODO(), token).Bytes()
		if err == redis.Nil {
			response.FailWithMessage("您无权访问，请登录！", ctx)
			ctx.Abort()
			return
		}
		// token验证
		customClaims, err := utils.ParseToken(token)
		if err != nil {
			response.FailWithMessage(err.Error(), ctx)
			ctx.Abort()
			return
		}
		// 将customClaims放入到context中,后续可取用户信息
		ctx.Set(constant.CustomClaimsKey, customClaims)
		var user model.User
		_ = json.Unmarshal(res, &user)
		if user.ID == customClaims.BaseClaims.ID {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusForbidden, "Token信息被篡改，您无权访问，请登录！")
			ctx.Abort()
		}
		//}
	}
}
