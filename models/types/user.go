package types

import (
	"backend/prisma/db"
	"time"
)

type UserType struct {
	ID int `json:"id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Email       string `json:"email"`
	Password    string `json:"password"`
	Username    string `json:"username"`
	Displayname string `json:"displayname"`
	PFP         string `json:"pfp"`
	Bio         string `json:"bio"`

	Tweets []TweetType `json:"tweets"`
}

type UpdateUserType struct {
	Password    string `json:"password"`
	NewPassword string `json:"newpassword"`
	Username    string `json:"username"`
	Displayname string `json:"displayname"`
	PFP         string `json:"pfp"`
	Bio         string `json:"bio"`
}

type TweetType struct {
	ID int `json:"id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Tweet string `json:"tweet"`
	Image string `json:"image"`

	UserID int      `json:"user_id"`
	User   UserType `json:"user"`
}

func (u *UserType) puts(data *db.UserModel) {

}
