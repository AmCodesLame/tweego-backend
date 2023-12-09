package controllers

import (
	"backend/models"
	"fmt"

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
	Response, err := models.GetUserById(id)
	if err != nil {
		u.Ctx.WriteString(fmt.Sprintf("Error: %v", err.Error()))
		return
	}
	u.Ctx.Output.Header("Content-Type", "application/json")
	u.Ctx.Output.Body(Response)
}
