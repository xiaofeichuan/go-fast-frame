package middleware

import (
	"fmt"
	"go-fast-frame/common/dto/response"
	"go-fast-frame/common/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JwtAuth 验证token
func JwtAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.Request.Header.Get("Authorization")
		tokenStr = strings.Replace(tokenStr, "Bearer ", "", -1)

		if tokenStr == "" {
			response.FailWithCode(http.StatusUnauthorized, "未授权", context)
			context.Abort() //结束后续操作
			return
		}

		// 解析token
		token, err := utils.ParseToken(tokenStr)
		if err != nil {
			response.FailWithCode(http.StatusUnauthorized, "授权失败，错误Token", context)
			context.Abort() //结束后续操作
			return
		}
		//是否过期

		// 获取 token 中的 claims
		claims, ok := token.Claims.(*utils.UserAuthClaims)
		if !ok {
			response.FailWithCode(http.StatusUnauthorized, "授权失败，解析失败", context)
			context.Abort() //结束后续操作
			return
		}
		fmt.Println("获取 token 中的 claims")
		fmt.Println(claims)
		context.Set("claims", claims)
		context.Next() //继续操作
	}
}
