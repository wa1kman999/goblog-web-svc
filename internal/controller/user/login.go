package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wa1kman999/goblog-web-svc/global"
	"github.com/wa1kman999/goblog-web-svc/internal/http/vs"
	"github.com/wa1kman999/goblog-web-svc/pkg/common/logger"
	"github.com/wa1kman999/goblog-web-svc/pkg/common/utils"
	"github.com/wa1kman999/goblog-web-svc/pkg/user/model"
	"go-micro.dev/v4/client"
)

type LoginParam struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

type LoginResp struct {
	User      model.User `json:"user"`
	Token     string     `json:"token"`
	ExpiresAt int64      `json:"expiresAt"`
}

// Login 用户登陆
func Login(ctx *gin.Context) {
	var l LoginParam
	log := logger.WithContext(ctx.Request.Context())
	if err := ctx.ShouldBindJSON(&l); err != nil {
		log.Errorf(err, "login param is invalid")
		vs.SendParamParseError(ctx)
		return
	}
	// 客户端
	cli := client.NewClient()

	userSvc := accountSVC.NewTaggingAccountSvcService("tagging-account-svc", cli)
	resp, err := userSvc.Login(context.Background(), &accountSVC.LoginRequest{
		Account:  l.Username,
		Password: l.Password,
	})
	fmt.Printf("值：%v", resp)
	fmt.Println(resp.Msg)
	fmt.Println(resp.Code)
	fmt.Println(resp.UserName)
	fmt.Println(resp.Role)
	if err != nil {
		log.Errorf(err, "登陆失败! 用户名不存在或者密码错误!")
		vs.SendBadData(ctx, err, "用户名不存在或者密码错误!")
		return
	} else {
		tokenNext(ctx, model.User{
			Username: resp.UserName,
			Role:     resp.Role,
		})
	}
}

// 登录以后签发jwt
func tokenNext(ctx *gin.Context, user model.User) {
	log := logger.WithContext(ctx.Request.Context())
	j := &utils.JWT{SigningKey: []byte(global.GoBlogWebConfig.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(utils.BaseClaims{
		//ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		log.Errorf(err, "获取token失败!")
		vs.SendBadData(ctx, err, "获取token失败!")
		return
	}
	vs.SendOkData(ctx, LoginResp{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	})
}
