package types

import (
	"backend/prisma/db"
	"time"
)

type UserType struct {
	id int

	createdAt time.Time
	updatedAt time.Time

	email       string
	username    string
	displayname string
	pfp         string
	tweets      []TweetType
}

type TweetType struct {
	id int

	createdAt time.Time
	updatedAt time.Time

	tweet string
	image string

	userId int
	user   UserType
}

func (u *UserType) puts(data *db.UserModel) {

}
