package dto

type Error struct {
	Message string `json:"message"`
}

type PageParamsInput struct {
	Filter     string   `json:"filter"`
	SearchType []string `json:"search_type"`
	Limit      string   `json:"limit"`
	Page       string   `json:"page"`
}
