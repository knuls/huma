package creator

import (
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Creator struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	FirstName string             `json:"firstName" bson:"firstName" validate:"required"`
	LastName  string             `json:"lastName" bson:"lastName" validate:"required"`
	Password  string             `json:"password" bson:"password" validate:"required"`
	Verified  bool               `json:"verified" bson:"verified"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt" validate:"required"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt" validate:"required"`
}

func NewCreator() *Creator {
	return &Creator{}
}

func (c *Creator) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
