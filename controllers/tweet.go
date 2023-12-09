package controllers

import (
	"backend/models"
	// "encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type TweetController struct {
	beego.Controller
}

func (u *TweetController) Get() {
	Response := models.ReturnMsg()
	u.Data["json"] = Response
	u.ServeJSON()
}
func (u *TweetController) Post() {
	Response := models.ReturnMsg()
	u.Data["json"] = Response
	u.ServeJSON()
}

func (u *TweetController) GetById() {
	id := u.GetString(":id")
	Response := models.ReturnMsgParam(id)
	u.Data["json"] = Response
	u.ServeJSON()
}
