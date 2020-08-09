package Types

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
 * Claim
 */
type Claim struct {
	Email string `json:"email"`
	Id primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}
