package records

type Likes struct {
	UserId     string `db:"user_id"`
	PetitionId string `db:"petition_id"`
}
