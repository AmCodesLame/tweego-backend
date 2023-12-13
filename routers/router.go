// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"backend/controllers"
	"backend/middleware"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.InsertFilter("/user/*", beego.BeforeRouter, middleware.Authenticate)
	beego.InsertFilter("/tweet/*", beego.BeforeRouter, middleware.Authenticate)
	authNS := beego.NewNamespace("/auth",
		beego.NSRouter("/register", &controllers.NewuserController{}, "post:RegisterUserPlease"),
	)

	userNS := beego.NewNamespace("/user",
		//get user list
		beego.NSRouter("/", &controllers.UserController{}, "get:Get"),
		//create user
		beego.NSRouter("/useraction", &controllers.UserController{}, "post:CreateUser"),
		//update user
		beego.NSRouter("/useraction", &controllers.UserController{}, "put:UpdateUser"),
		//delete user
		beego.NSRouter("/useraction", &controllers.UserController{}, "delete:DelUser"),
		//get user by username
		beego.NSRouter("/getuser/:uname:string", &controllers.UserController{}, "get:GetUserByUsername"),
		//get user by id
		beego.NSRouter("/getuser/:id:int", &controllers.UserController{}, "get:GetUserById"),
		//
	)

	tweetNS := beego.NewNamespace("/tweet",
		//get teet list
		// beego.NSRouter("/", &controllers.TweetController{}, "get:Get"),

		//post tweet
		beego.NSRouter("/tweetaction", &controllers.TweetController{}, "post:PostTweet"),
		//update tweet
		// beego.NSRouter("/tweetaction", &controllers.TweetController{}, "put:UpdateTweet"),
		//delete tweet
		beego.NSRouter("/tweetaction", &controllers.TweetController{}, "delete:DeleteTweet"),

		//get tweet by username
		beego.NSRouter("/gettweet/:username:string", &controllers.TweetController{}, "post:GetTweetsByUser"),

		//get tweet by id
		beego.NSRouter("/gettweet/:id:int", &controllers.TweetController{}, "get:GetTweetById"),
	)

	beego.AddNamespace(userNS, tweetNS, authNS)

	// ns := beego.NewNamespace("/v1",
	// 	beego.NSNamespace("/object",
	// 		beego.NSInclude(
	// 			&controllers.ObjectController{},
	// 		),
	// 	),
	// 	beego.NSNamespace("/user",
	// 		beego.NSInclude(
	// 			&controllers.UserController{},
	// 		),
	// 	),
	// )
	// beego.AddNamespace(ns)
}
