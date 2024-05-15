package models

type Car struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CarModelId int    `json:"car_model_id"`
}

type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CarModel struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	BrandID int    `json:"brand_id"`
}
