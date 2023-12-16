package controllers

import (
	"backend/models"
	"backend/models/types"
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	beego.Controller
}

func (u *AuthController) RegisterUser() {
	var user types.UserType
	requestBody := u.Ctx.Input.RequestBody
	// var requestData map[string]interface{}

	if err := json.Unmarshal(requestBody, &user); err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	jwtToken, err := models.CreateUser(user)
	if err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	u.Ctx.Output.Context.SetCookie("auth", jwtToken)
	u.Ctx.Output.JSON(map[string]string{"messaage": "User Created", "token": jwtToken}, false, false)
	// u.Ctx.Output.SetStatus(201)
}

func (u *AuthController) LoginUser() {
	var user types.UserType
	requestBody := u.Ctx.Input.RequestBody
	// var requestData map[string]interface{}

	if err := json.Unmarshal(requestBody, &user); err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	jwtToken, err := models.LoginUser(user)
	if err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	u.Ctx.Output.Context.SetCookie("auth", jwtToken)
	u.Ctx.Output.JSON(map[string]string{"messaage": "User Created", "token": jwtToken}, false, false)
}
