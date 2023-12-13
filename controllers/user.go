package controllers

import (
	"backend/models"
	"backend/models/types"
	"encoding/json"
	"fmt"

	// "encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	Response, err := models.GetUserList()
	if err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	u.Ctx.Output.Header("Content-Type", "application/json")
	u.Ctx.Output.Body(Response)
}
func (u *UserController) Post() {
	id := u.GetString(":uname")

	Response := models.ReturnMsgParam(id)
	u.Data["json"] = Response
	u.ServeJSON()
}
func (u *UserController) Post1() {
	id := u.GetString(":id")

	Response := models.ReturnMsgParam(id)
	u.Data["json"] = Response
	u.ServeJSON()
}

func (u *UserController) GetUserByUsername() {
	id := u.GetString(":uname")
	Response, err := models.GetUserByUsername(id)
	if err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	u.Ctx.Output.Header("Content-Type", "application/json")
	u.Ctx.Output.Body(Response)
}
func (u *UserController) GetUserById() {
	id, err := u.GetInt(":id")
	if err != nil {
		u.Ctx.Output.SetStatus(400)
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	Response, err := models.GetUserById(id)
	if err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	u.Ctx.Output.Header("Content-Type", "application/json")
	u.Ctx.Output.Body(Response)
}

func (u *UserController) CreateUser() {
	var user types.UserType
	requestBody := u.Ctx.Input.RequestBody
	// var requestData map[string]interface{}

	if err := json.Unmarshal(requestBody, &user); err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	if _, err := models.CreateUser(user); err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	u.Ctx.Output.JSON(map[string]string{"messaage": "User Created"}, false, false)
	// u.Ctx.Output.SetStatus(201)
}

func (u *UserController) UpdateUser() {
	var user types.UpdateUserType

	requestBody := u.Ctx.Input.RequestBody

	if err := json.Unmarshal(requestBody, &user); err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	if user.Password == "" || user.Username == "" {
		u.Ctx.Output.JSON(map[string]string{"messaage": "Please provide the Username and Password"}, false, false)
		return
	}
	if err := models.UpdateUser(user); err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	u.Ctx.Output.JSON(map[string]string{"messaage": "User Updated"}, false, false)

}

func (u *UserController) DelUser() {
	var user types.UserType

	requestBody := u.Ctx.Input.RequestBody

	if err := json.Unmarshal(requestBody, &user); err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	if err := models.DeleteUser(user); err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	u.Ctx.Output.JSON(map[string]string{"messaage": "User Deletetd"}, false, false)

}

func ErrGen(e error) string {
	return fmt.Sprintf("Error: %s", e.Error())
}
