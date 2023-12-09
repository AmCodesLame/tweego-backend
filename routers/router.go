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

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	userNS := beego.NewNamespace("/user",
		//get user list
		beego.NSRouter("/", &controllers.UserController{}, "get:Get"),
		//post user
		beego.NSRouter("/", &controllers.UserController{}, "post:Post"),
		//get user by id
		beego.NSRouter("/:id", &controllers.UserController{}, "get:GetById"),
	)

	tweetNS := beego.NewNamespace("/tweet",
		//get user list
		beego.NSRouter("/", &controllers.TweetController{}, "get:Get"),
		//post user
		beego.NSRouter("/", &controllers.TweetController{}, "post:Post"),
		//get user by id
		beego.NSRouter("/:id", &controllers.TweetController{}, "get:GetById"),
	)
	beego.AddNamespace(userNS, tweetNS)

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
