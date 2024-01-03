package v1

type Likes struct {
	UserId     string `db:"user_id"`
	PetitionId string `db:"petition_id"`
}
