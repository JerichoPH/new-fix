package middlewares

import (
	"fmt"
	"new-fix/models"
	"new-fix/tools"
	"new-fix/types"
	"new-fix/wrongs"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CheckAuth 检查Jwt是否合法
func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取令牌
		split := strings.Split(tools.GetJwtFromHeader(ctx), " ")
		if len(split) != 2 {
			wrongs.ThrowUnLogin("令牌格式错误")
		}
		tokenType := split[0]
		token := split[1]

		var (
			account models.AccountMdl
			ret     *gorm.DB
		)
		account = models.AccountMdl{}
		if token == "" {
			wrongs.ThrowUnLogin("令牌不存在")
		} else {
			switch tokenType {
			case "JWT":
				claims, err := tools.ParseJwt(token)

				// 判断令牌是否有效
				if err != nil {
					wrongs.ThrowUnLogin("令牌解析失败")
				} else if time.Now().Unix() > claims.ExpiresAt {
					wrongs.ThrowUnLogin("令牌过期")
				}

				// 判断用户是否存在
				if reflect.DeepEqual(claims, tools.Claims{}) {
					wrongs.ThrowUnLogin("令牌解析失败：用户不存在")
				}

				// 获取用户信息
				ret = models.NewMySqlMdl().SetModel(models.AccountMdl{}).GetDb("").Where("uuid", claims.Uuid).First(&account)
				if wrongs.ThrowWhenEmpty(ret, "") {
					wrongs.ThrowUnLogin(fmt.Sprintf("令牌指向用户(JWT) %s %v ", token, claims))
				}
			case "AU":
				ret = models.NewMySqlMdl().SetModel(models.AccountMdl{}).SetWheres(map[string]any{"uuid": token}).GetDb("").First(&account)
				if wrongs.ThrowWhenEmpty(ret, "") {
					wrongs.ThrowUnLogin(fmt.Sprintf("令牌指向用户(AU) %s", token))
				}
			default:
				wrongs.ThrowForbidden("权鉴认证方式不支持")
			}

			ctx.Set(string(types.ACCOUNT_ID), account.Id)             // 设置用户编号
			ctx.Set(string(types.ACCOUNT_ACCOUNT), account.Username)  // 设置用户账号
			ctx.Set(string(types.ACCOUNT_NICKNAME), account.Nickname) // 设置用户昵称
			ctx.Set(string(types.ACCOUNT_AUTH), account)              // 设置用户信息
		}

		ctx.Next()
	}
}
