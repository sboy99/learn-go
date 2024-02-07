package dtos

// ------------------------------REGISTER------------------------------- //

type RegisterReqDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type RegisterResDto struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// ------------------------------LOGIN------------------------------- //

type LoginReqDto struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type LoginResDto struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
