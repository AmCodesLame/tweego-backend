package models

import (
	"backend/models/types"
	"backend/prisma/db"
	"encoding/json"
	"fmt"
)

// "errors"
// "strconv"
// "time"
// "backend/models/types"
// "backend/prisma/db"
// "encoding/json"
// "fmt"

// "golang.org/x/crypto/bcrypt"

func CreateTweet(t types.TweetType) error {
	client, ctx, err := db.Connect()
	if err != nil {
		return err
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	var id int
	id = t.UserID
	if t.UserID == 0 {
		id, err = GetUserIdByUname(t.User.Username)
		if err != nil {
			return err
		}
	}

	_, err = client.Tweet.CreateOne(
		db.Tweet.Tweet.Set(t.Tweet),
		db.Tweet.User.Link(db.User.ID.Equals(id)),
		db.Tweet.Image.Set(t.Image),
	).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func TweetsByUser(u types.TweetType, username string) (data []byte, err error) {
	client, ctx, err := db.Connect()
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	userId := u.UserID
	if u.UserID == 0 {
		if username == "" || u.User.Username != username {
			return nil, fmt.Errorf("Please provide the correct username")
		}
		userId, err = GetUserIdByUname(username)
		if userId == 0 {
			return nil, fmt.Errorf("No Username found")
		}
	}

	tweets, err := client.Tweet.FindMany(db.Tweet.UserID.Equals(userId)).Exec(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return nil, err
	}
	var ResponseArray []json.RawMessage
	for _, val := range tweets {
		Response, err := json.Marshal(val)
		if err != nil {
			fmt.Printf("Error in parsing JSON: %v", err.Error())
			return nil, err
		}
		ResponseArray = append(ResponseArray, json.RawMessage(Response))
	}
	response, err := json.Marshal(ResponseArray)
	return response, nil
}

func DeleteTweet(tweet types.TweetType) error {
	client, ctx, err := db.Connect()
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	_, err = client.Tweet.FindUnique(
		db.Tweet.ID.Equals(tweet.ID),
	).Delete().Exec(ctx)
	if err != nil {
		return err
	}

	// if err := bcrypt.CompareHashAndPassword(userDb.Password, []byte(user.Password)); err != nil {
	// 	return err
	// }
	return nil
}

func GetTweetById(id int) ([]byte, error) {
	// id, err := strconv.Atoi(TweetId)

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return nil, err
	// }
	client, ctx, err := db.Connect()
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	tweet, err := client.Tweet.FindUnique(
		db.Tweet.ID.Equals(id),
	).With(db.Tweet.User.Fetch()).Exec(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return nil, fmt.Errorf("Tweet Not Found!")
	}

	result, err := json.Marshal(tweet)
	if err != nil {
		fmt.Printf("Error in parsing JSON: %v", err.Error())
		return nil, err
	}
	return result, nil
}
