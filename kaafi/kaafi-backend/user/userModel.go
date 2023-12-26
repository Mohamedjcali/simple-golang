package user
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type User struct {
	Id          primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty"`
	Phonenumber string	 `json:"phonenumber,omitempty" bson:"phonenumber,omitempty" `
	Name        string    `json:"name,omitempty" bson:"name,omitempty"`
	Age         int       `json:"age,omitempty" bson:"age,omitempty"`
	Gender      string    `json:"gender,omitempty" bson:"gender,omitempty"`
	Images      []string  `json:"images,omitempty" bson:"images,omitempty"`
	Location    string    `json:"location,omitempty" bson:"location,omitempty"`
	ProfilePic  string    `json:"profilepic,omitempty" bson:"profilepic,omitempty"`
	Bio         string    `json:"bio,omitempty" bson:"bio,omitempty"`
	Work        bool      `json:"work,omitempty" bson:"work,omitempty"`
	Shiinaay    int       `json:"shiinaay,omitempty" bson:"shiinaay,omitempty"`

}