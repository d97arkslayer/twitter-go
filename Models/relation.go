package Models

/**
 * Relation
 * Model to relations between users
 */
type Relation struct {
	UserId string `bson:"userId" json:"userId"`
	UserRelationId string `bson:"userRelationId" json:"userRelationId"`
}
