package dtos

type CreateTaskDto struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}
