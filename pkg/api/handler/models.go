package handler

type followRequest struct {
	Followee int `json:"value"`
}

type userCreateRequest struct {
	Nickname string `json:"nickname" validate:"required"`
}
