package cmd

import (
	"CSIE_ADMS_Back-End/internal/dao"
	"CSIE_ADMS_Back-End/internal/model/entity"
	"CSIE_ADMS_Back-End/utility"
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"
)

func loginFunc(r *ghttp.Request) (string, interface{}) {
	email := r.Get("email").String()
	password := r.Get("password").String()

	if email == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("email or password is empty"))
		r.ExitAll()
	}

	// Get user by email
	user := entity.Users{}
	err := dao.Users.Ctx(context.TODO()).Scan(&user, "email", email)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail("email not found"))
		r.ExitAll()
	}

	// Check password
	checkPWD, err := utility.ComparePWD(password, user.Password)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(err.Error()))
		r.ExitAll()
	}
	if !checkPWD {
		r.Response.WriteJson(gtoken.Fail("password error"))
		r.ExitAll()
	}
	return strconv.Itoa(user.Id), ""
}

func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = -1
		r.Response.WriteJsonExit(respData)
	}

	// Get user by id
	userId := respData.GetString("userKey")
	user := entity.Users{}
	err := dao.Users.Ctx(context.TODO()).Scan(&user, "id", userId)
	if err != nil {
		respData.Code = -1
		respData.Msg = err.Error()
		r.Response.WriteJsonExit(respData)
		return
	}

	respData.Code = 0
	respData.Data = g.Map{
		"type":     "Bearer",
		"token":    respData.GetString("token"),
		"username": user.Username,
		"role":     user.Role,
	}
	r.Response.WriteJson(respData)
	return
}

func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	userID := 0
	err := gconv.Struct(respData.GetString("userKey"), &userID)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code": -1,
			"msg":  "Please login",
		})
	}
	r.SetCtxVar("id", userID)
	r.Middleware.Next()
}
