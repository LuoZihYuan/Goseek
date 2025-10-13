package models

type ProductsSearchResponse struct {
	Products   []*Product `json:"products"`
	TotalFound int        `json:"total_found"`
	SearchTime string     `json:"search_time"`
}
