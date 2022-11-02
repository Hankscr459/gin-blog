package dto

type CreateOrgInput struct {
	Name   string   `json:"name"`
	Access []Access `json:"access"`
}

type Access struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
}
