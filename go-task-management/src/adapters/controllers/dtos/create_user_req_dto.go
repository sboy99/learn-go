package dtos

type CreateUserReqDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
