package records

type Voices struct {
	UserId     string `db:"user_id"`
	PetitionId string `db:"petition_id"`
}
