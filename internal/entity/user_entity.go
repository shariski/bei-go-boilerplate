package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Password string             `bson:"password"`
	Name     string             `bson:"name"`
}

func (u *User) CollName() string {
	return "users"
}
