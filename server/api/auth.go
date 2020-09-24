package api

import "github.com/gin-gonic/gin"



// @Summary 登录
// @Produce  json
// @Param userName path int true "userName"
// @Param passWord path int true "password"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/login/{id} [Post]
func GetAuth(ctx *gin.Context){
		userName := ctx.PostForm("userName")
		println(userName)

}


// @Summary 获取用户列表
// @Produce  json
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/user [Get]
func GetUser(ctx *gin.Context){
	msg := ctx.PostForm("msg")
	println(msg)


}
