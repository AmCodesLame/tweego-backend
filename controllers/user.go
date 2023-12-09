package controllers

import (
	"backend/models"
	// "encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	Response := models.ReturnMsg()
	u.Data["json"] = Response
	u.ServeJSON()
}
func (u *UserController) Post() {
	Response := models.ReturnMsg()
	u.Data["json"] = Response
	u.ServeJSON()
}

func (u *UserController) GetById() {
	id := u.GetString(":id")
	Response := models.ReturnMsgParam(id)
	u.Data["json"] = Response
	u.ServeJSON()
}
