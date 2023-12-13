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

func (u *TweetController) PostTweet() {
	var tweet types.TweetType
	requestBody := u.Ctx.Input.RequestBody
	// var requestData map[string]interface{}

	if err := json.Unmarshal(requestBody, &tweet); err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	// fmt.Println(tweet.ID == 0)
	if err := models.CreateTweet(tweet); err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	u.Ctx.Output.JSON(map[string]string{"messaage": "Tweet Posted"}, false, false)
	// u.Ctx.Output.SetStatus(201)
}

func (u *TweetController) GetTweetsByUser() {
	var tweetUser types.TweetType
	requestBody := u.Ctx.Input.RequestBody
	uname := u.GetString("username")
	// var requestData map[string]interface{}

	if err := json.Unmarshal(requestBody, &tweetUser); err != nil {
		fmt.Println("Error in parsing the JSON")
		// u.Ctx.Output.SetStatus(400)
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	Response, err := models.TweetsByUser(tweetUser, uname)
	if err != nil {
		fmt.Println("Error in Fetching Tweets")
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	u.Ctx.Output.Header("Content-Type", "application/json")
	u.Ctx.Output.Body(Response)
}

func (u *TweetController) DeleteTweet() {
	var tweet types.TweetType
	requestBody := u.Ctx.Input.RequestBody
	// var requestData map[string]interface{}

	if err := json.Unmarshal(requestBody, &tweet); err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	if err := models.DeleteTweet(tweet); err != nil {
		u.Ctx.WriteString(ErrGen(err))
		return
	}
	u.Ctx.Output.JSON(map[string]string{"messaage": "Tweet Deletetd"}, false, false)
}

func (u *TweetController) GetTweetById() {
	id, _ := u.GetInt(":id")
	fmt.Println(id)
	tweet, err := models.GetTweetById(id)
	if err != nil {
		u.Ctx.Output.SetStatus(400)
		u.Ctx.Output.Body([]byte(ErrGen(err)))
		return
	}
	u.Ctx.Output.Header("Content-Type", "application/json")
	u.Ctx.Output.Body(tweet)
}
