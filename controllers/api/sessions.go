package controllerapi

import (
	"context"
	"gin_blog/lib/global"
	"gin_blog/lib/utils"
	"gin_blog/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginUser struct {
	Phone     string `json:"phone"`
	LoginCode string `json:"login_code"`
}

var ctx = context.Background()

//登录
//params
// phone
// login_code
func Login(c *gin.Context) {
	var user models.User
	//声明接收的变量
	var loginParams LoginUser
	//将request的body中数据，自动按照json格式解析到结构体中
	if err := c.ShouldBindJSON(&loginParams); err != nil {
		//返回错误信息
		c.JSON(200, gin.H{"status": 500, "msg": "参数解析失败", "data": gin.H{"error": err.Error()}})
		return
	}
	phone := loginParams.Phone
	loginCode := loginParams.LoginCode
	rcode, rerr := global.REDIS.Get(ctx, phone).Result()
	if rerr != nil {
		c.JSON(200, gin.H{"status": 500, "msg": "验证码获取失败", "data": gin.H{}})
		return
	}
	if loginCode != rcode {
		c.JSON(200, gin.H{"status": 500, "msg": "验证码错误", "data": gin.H{}})
		return
	}
	global.DB.Where("Phone = ?", phone).Find(&user)
	if user == (models.User{}) {
		c.JSON(200, gin.H{
			"status": 500,
			"msg":    "登录失败,请核实信息",
			"data":   gin.H{},
		})
		return
	}
	//jwt鉴权
	j := &utils.JWT{}
	jwtinfo := utils.JWTCustomInfo{
		ID:       user.ID,
		Nickname: user.Nickname,
	}
	token, err := j.CreateToken(jwtinfo)
	if err != nil {
		c.JSON(200, gin.H{"status": 500, "msg": "获取token失败,请核实信息", "data": gin.H{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "token获取成功",
		"data":   token,
	})
}

//发送短信验证码
//@params
//phone: 电话
func SendLoginCode(c *gin.Context) {
	//生成四位随机数
	code := 1234
	var loginParams LoginUser
	if err := c.ShouldBindJSON(&loginParams); err != nil {
		//返回错误信息
		c.JSON(200, gin.H{"status": 500, "msg": "参数解析失败", "data": gin.H{}})
		return
	}

	rerr := global.REDIS.Set(ctx, loginParams.Phone, code, 5*time.Minute).Err()
	if rerr != nil {
		panic(rerr)
	}

	xcode, e1 := global.REDIS.Get(ctx, loginParams.Phone).Result()
	if e1 != nil {
		panic(e1)
	}
	c.JSON(200, gin.H{
		"status": 200,
		"msg":    "获取手机登录验证码成功",
		"data": gin.H{
			"code": xcode,
		},
	})

}
