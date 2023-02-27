package organization

import (
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organization struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" validate:"required,alphanum"`
	CreatorID primitive.ObjectID `json:"creatorId" bson:"userId" validate:"required,oid"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt" validate:"required"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt" validate:"required"`
}

func NewOrganization() *Organization {
	return &Organization{}
}

func (c *Organization) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
