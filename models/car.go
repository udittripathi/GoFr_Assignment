package models

type Car struct {
	ID          string    `json:"id"`
	Make        string    `json:"make"`
	Model       string    `json:"model"`
	EntryTime   string `json:"entry_time"`
	RepairStatus string   `json:"repair_status"`
}
