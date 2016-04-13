package models

import (
	//"github.com/novikk/redline/models/mongo"

	"sharit-backend/models/mongo"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

// User is a user :D
type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	IDuser    string        `bson:"iduser,omitempty"`
	Email     string        `bson:"email,omitempty"`
	Pass      string        `bson:"pass,omitempty"`
	Name      string        `bson:"name,omitempty"`
	Surname   string        `bson:"surname,omitempty"`
	Stars     string        `bson:"stars,omitempty"`
	ItemsUser Items         `bson:"itemsUser,omitempty"`
	X         string        `bson:"x,omitempty"`
	Y         string        `bson:"y,omitempty"`
}

//Users is a list of User
type Users []User

// Create creates a user with its information in the database
func (u *User) Create() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C(" ")
	err := c.Insert(u)
	return err
}

// FindUserByID returns a user found by steamid
func FindUserByID(id string) (User, error) {
	var u User

	db := mongo.Conn()
	defer db.Close()

	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Find(bson.M{"iduser": id}).One(&u)

	return u, err
}

// UpdateEmail updates user email
func (u *User) UpdateEmail(email string) error {
	db := mongo.Conn()
	defer db.Close()

	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$set": bson.M{"email": email}})

	return err
}

// GetAllUsers returns all users
func GetAllUsers() (Users, error) {
	db := mongo.Conn()
	defer db.Close()
	var p Users
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Find(bson.M{}).All(&p)
	return p, err
}
