package Types

/**
 * Tweet
 * Struct to create a tweet
 */
type Tweet struct {
	Message string `bson:"message" json:"message"`
}