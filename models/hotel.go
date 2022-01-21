package models


type Hotel struct {
	Name string `json:"name"`
	TotalRooms int `json:"total_rooms"`
	OccupiedRooms int `json:"occupied_rooms"`
	CostPerDay int `json:"cost_per_day"`
}

type HotelIn struct {
	Name string `json:"name"`
	TotalRooms int `json:"total_rooms"`
	CostPerDay int `json:"cost_per_day"`
}