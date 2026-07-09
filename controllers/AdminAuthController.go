package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/pkg/helpers"
	"github.com/outsstill/gin-admin/pkg/jwt"
	"github.com/outsstill/gin-admin/pkg/response"
	"github.com/outsstill/gin-admin/requests"
	"github.com/outsstill/gin-admin/setting"
	"github.com/outsstill/go-kit/logger"
)

type AdminAuthController struct {
	*BaseAPIController
}

func NewAdminAuthController(base *BaseAPIController) *AdminAuthController {
	return &AdminAuthController{
		BaseAPIController: base,
	}
}

func (ac *AdminAuthController) Login(c *gin.Context) {
	// 验证
	request := requests.AdminLoginRequest{}
	if ok := requests.ValidateFunc(c, ac.App.DB, &request, requests.VerityAdminLogin); !ok {
		return
	}

	// 如果有验证码
	if ok := ac.App.GetCaptchaService().VerifyCaptcha(request.CaptchaID, request.CaptchaAnswer); !ok {
		response.Fail(c, "验证码错误!")
		return
	}

	//if err := c.ShouldBind(&request); err != nil {
	//	response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
	//	return
	//}

	userModel, err := ac.App.GetAuthService().AttemptAdmin(request.Username, request.Password)

	if err != nil {
		// 失败，显示错误提示
		response.Fail(c, "账号不存在或密码错误")
		return
	} else {
		token := jwt.NewJWT().IssueAdminToken(userModel.GetStringID(), userModel.Username)

		// 获取账号的显示菜单
		//menus, errs := adminUser.GetUserMenus(userModel.ID)
		//
		//if errs != nil {
		//	ac.App.Response.Error(c, err, "获取玩家显示菜单错误")
		//	return
		//}
		//
		//userModel.Menus = menus

		c.Set("current_admin_user_id", userModel.GetStringID())

		response.Data(c, gin.H{
			"token": token,
			"user":  userModel,
		})
	}
}

func (ac *AdminAuthController) Logout(c *gin.Context) {

	response.Success(c)
}

func (ac *AdminAuthController) Current(c *gin.Context) {

	user := ac.App.GetAuthService().CurrentAdminUser(c)

	response.Data(c, user)
}

// RefreshToken 刷新 Access Token
func (ac *AdminAuthController) RefreshToken(c *gin.Context) {

	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.Data(c, gin.H{
			"token": token,
		})
	}
}

// ShowCaptcha 显示图片验证码
func (ac *AdminAuthController) ShowCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, answer, err := ac.App.GetCaptchaService().GenerateCaptcha()

	if setting.IsDebug() {
		fmt.Printf("获取验证码 id:%s answer:%s\n", id, answer)
	}

	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	logger.LogIf(err)
	// 返回给用户
	response.Data(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

func (ac *AdminAuthController) UpdateProfile(c *gin.Context) {

	user := ac.App.GetAuthService().CurrentAdminUser(c)

	// 验证
	request := requests.AdminUserProfileUpdateRequest{}
	if ok := requests.ValidateFunc(c, ac.App.DB, &request, requests.VerityAdminUserProfileUpdate); !ok {
		return
	}

	if !helpers.Empty(request.Name) {
		user.Name = request.Name
	}

	ac.App.GetAdminUserService().Save(user)

	response.Data(c, user)
}

func (ac *AdminAuthController) UpdatePassword(c *gin.Context) {

	user := ac.App.GetAuthService().CurrentAdminUser(c)

	// 验证
	request := requests.AdminUserProfilePasswordUpdateRequest{}
	if ok := requests.ValidateFunc(c, ac.App.DB, &request, requests.VerityAdminUserProfilePasswordUpdate); !ok {
		return
	}

	if !helpers.Empty(request.Password) {
		user.Password = request.Password
	}

	ac.App.GetAdminUserService().Save(user)

	response.Data(c, user)
}
