package models

import (
	// "errors"
	// "strconv"
	// "time"
	"backend/models/types"
	"backend/prisma/db"
	"encoding/json"
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

////////////////////////////

// ///////////////////////

func LoginUser(user types.UserType) (string, error) {
	client, ctx, err := db.Connect()
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	userDb, err := client.User.FindUnique(
		db.User.Username.Equals(user.Username),
	).Exec(ctx)
	if err != nil {
		return "", fmt.Errorf("User Not Found, {%v}", err.Error())
	}
	if err := bcrypt.CompareHashAndPassword(userDb.Password, []byte(user.Password)); err != nil {
		return "", fmt.Errorf("Password Incorrect! {%v}", err.Error())
	}
	claims := jwt.MapClaims{
		"username":  user.Username,
		"id":        userDb.ID,
		"IssuedAt":  time.Now().UTC().Unix(),
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
		"Issuer":    "tweego",
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("HMACKEY")))
	if err != nil {
		fmt.Println("Error In JWT Signing", err)
	}
	return tokenString, nil
}

func GetUserByUsername(uname string) (data []byte, err error) {
	client, ctx, err := db.Connect()
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	user, err := client.User.FindUnique(
		db.User.Username.Equals(uname),
	).With(db.User.Tweets.Fetch()).Exec(ctx)

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return nil, err
	}

	result, err := json.Marshal(user)

	if err != nil {
		fmt.Printf("Error in parsing JSON: %v", err.Error())
		return nil, err
	}
	fmt.Printf("user: %s\n", result)
	return result, nil
}

func GetUserById(userid int) (data []byte, err error) {
	username, err := GetUsernameByUserid(userid)
	if err != nil {
		return nil, err
	}
	return GetUserByUsername(username)
}

func CreateUser(user types.UserType) (string, error) {
	client, ctx, err := db.Connect()
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	encPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	createdPost, err := client.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Username.Set(user.Username),
		db.User.Password.Set(encPass),
		db.User.Displayname.Set(user.Displayname),
		db.User.Bio.Set(user.Bio),
		db.User.Pfp.Set(user.PFP),
	).Exec(ctx)
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"username":  user.Username,
		"id":        &createdPost.ID,
		"IssuedAt":  time.Now().UTC().Unix(),
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
		"Issuer":    "tweego",
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("HMACKEY")))
	if err != nil {
		fmt.Println("Error In JWT Signing", err)
	}

	result, _ := json.MarshalIndent(createdPost, "", "  ")
	fmt.Printf("created user: %s\n", result)
	return tokenString, nil
}

func UpdateUser(user types.UpdateUserType) error {

	client, ctx, err := db.Connect()
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	userDb, err := client.User.FindUnique(
		db.User.Username.Equals(user.Username),
	).Exec(ctx)
	if err != nil {
		return fmt.Errorf("User Not Found, {%v}", err.Error())
	}

	if err := bcrypt.CompareHashAndPassword(userDb.Password, []byte(user.Password)); err != nil {
		return fmt.Errorf("Password Incorrect! {%v}", err.Error())
	}
	if user.NewPassword != "" {
		encPass, err := bcrypt.GenerateFromPassword([]byte(user.NewPassword), 10)
		if err != nil {
			return err
		}
		_, err = client.User.FindUnique(
			db.User.Username.Equals(user.Username),
		).Update(
			db.User.Password.Set(encPass),
			db.User.Displayname.Set(user.Displayname),
			db.User.Bio.Set(user.Bio),
			db.User.Pfp.Set(user.PFP),
		).Exec(ctx)
		return nil
	}

	_, err = client.User.FindUnique(
		db.User.Username.Equals(user.Username),
	).Update(
		db.User.Displayname.Set(user.Displayname),
		db.User.Bio.Set(user.Bio),
		db.User.Pfp.Set(user.PFP),
	).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(user types.UserType) error {
	client, ctx, err := db.Connect()
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	userDb, err := client.User.FindUnique(
		db.User.Username.Equals(user.Username),
	).Exec(ctx)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword(userDb.Password, []byte(user.Password)); err != nil {
		return err
	}

	_, err = client.User.FindUnique(
		db.User.Username.Equals(user.Username),
	).Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func GetUserList() (data []byte, err error) {
	client, ctx, err := db.Connect()
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	user, err := client.User.FindMany().Exec(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return nil, err
	}
	var ResponseArray []json.RawMessage
	for _, val := range user {
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

func GetUserIdByUname(uname string) (int, error) {
	if uname == "" {
		return 0, fmt.Errorf("Username not provided")
	}
	client, ctx, err := db.Connect()
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	user, err := client.User.FindUnique(
		db.User.Username.Equals(uname),
	).Exec(ctx)

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return 0, err
	}
	return user.ID, nil
}

func GetUsernameByUserid(userid int) (string, error) {
	if userid == 0 {
		return "", fmt.Errorf("UserId not provided")
	}
	client, ctx, err := db.Connect()
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	user, err := client.User.FindUnique(
		db.User.ID.Equals(userid),
	).Exec(ctx)

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return "", fmt.Errorf("Username not found by UserID")
	}
	return user.Username, nil
}

///////////////////////////

// var (
// 	UserList map[string]*User
// )

// func init() {
// 	UserList = make(map[string]*User)
// 	u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
// 	UserList["user_11111"] = &u
// }

// type User struct {
// 	Id       string
// 	Username string
// 	Password string
// 	Profile  Profile
// }

// type Profile struct {
// 	Gender  string
// 	Age     int
// 	Address string
// 	Email   string
// }

// func AddUser(u User) string {
// 	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
// 	UserList[u.Id] = &u
// 	return u.Id
// }

// func GetUser(uid string) (u *User, err error) {
// 	if u, ok := UserList[uid]; ok {
// 		return u, nil
// 	}
// 	return nil, errors.New("User not exists")
// }

// func GetAllUsers() map[string]*User {
// 	return UserList
// }

// func UpdateUser(uid string, uu *User) (a *User, err error) {
// 	if u, ok := UserList[uid]; ok {
// 		if uu.Username != "" {
// 			u.Username = uu.Username
// 		}
// 		if uu.Password != "" {
// 			u.Password = uu.Password
// 		}
// 		if uu.Profile.Age != 0 {
// 			u.Profile.Age = uu.Profile.Age
// 		}
// 		if uu.Profile.Address != "" {
// 			u.Profile.Address = uu.Profile.Address
// 		}
// 		if uu.Profile.Gender != "" {
// 			u.Profile.Gender = uu.Profile.Gender
// 		}
// 		if uu.Profile.Email != "" {
// 			u.Profile.Email = uu.Profile.Email
// 		}
// 		return u, nil
// 	}
// 	return nil, errors.New("User Not Exist")
// }

// func Login(username, password string) bool {
// 	for _, u := range UserList {
// 		if u.Username == username && u.Password == password {
// 			return true
// 		}
// 	}
// 	return false
// }

// func DeleteUser(uid string) {
// 	delete(UserList, uid)
// }
