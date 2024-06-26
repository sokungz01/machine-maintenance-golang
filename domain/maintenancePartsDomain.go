package domain

import (
	"time"
)

type MaintenanceParts struct {
	MaintenancePartID int       `json:"maintenancePart" db:"maintenancePartID"`
	ServiceID         int       `json:"serviceID" db:"serviceID"`
	ItemID            int       `json:"itemID" db:"itemID"`
	Item              Item      `json:"item"`
	Qty               int       `json:"qty" db:"qty"`
	CreatedDate       time.Time `json:"createdDate" db:"createdDate"`
}

type MaintenancePartsRepository interface {
	CreateMaintenanceParts(serviceID int, itemID int, qty int, createdDate time.Time) error
	GetMaintenacnePartsByServiceID(serviceID int) (*[]MaintenanceParts, error)
}
