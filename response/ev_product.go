package response

type EvProduct struct {
	EvBaseModel
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CategoryID  int64      `json:"category_id"`
	Type        string     `json:"type"`
	Category    EvCategory `json:"category"`
}

type EvCategory struct {
	EvBaseModel
	Name string `json:"name"`
	Slug string `json:"slug"`
}
