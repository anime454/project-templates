package domain

type VehicleStatus string

type Vehicle struct {
	VehicleID      string `json:"vehicle_id"`
	LicensePlate   string `json:"license_plate"`
	RegisterStatus bool   `json:"register_status"`
	RegisteredAt   string `json:"registered_at"`
}
