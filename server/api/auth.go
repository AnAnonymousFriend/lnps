package api

import (
	"LNPS/server/models"
	"github.com/gin-gonic/gin"
)



// @Summary 登录
// @Produce  json
// @Param userName path int true "userName"
// @Param passWord path int true "password"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/login [Post]
func GetAuth(ctx *gin.Context){
		userName := ctx.PostForm("userName")
		passWord := ctx.PostForm("passWord")

		info,err := models.Login(userName,passWord)
		if err != nil {
			println(err)
		}else{
		println(info)
		}

}


// @Summary 获取用户列表
// @Produce  json
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/user [Post]
func AddUser(ctx *gin.Context){
	userName := ctx.PostForm("userName")
	passWord := ctx.PostForm("passWord")
	isCreate := models.AddUser(userName,passWord)
	println(isCreate)
}
