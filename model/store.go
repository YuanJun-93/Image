package model

type Image struct {
	ImageID   int64  `json:"image_id" db:"image_id"`
	ImageCost int64  `json:"image_cost" db:"image_cost"`
	UserID    int64  `json:"user_id" db:"user_id"`
	ImageName string `json:"image_name" db:"image_name"`
	ImageUrl  string `json:"image_url" db:"image_url"`
}
