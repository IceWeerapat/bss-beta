package auditlog

import (
	"time"

	"fx.prodigy9.co/app"
)

var App = app.Build().
	Controllers(AuditLogCtr{})

const (
	ActionCreate = "create"
	ActionUpdate = "update"
	ActionDelete = "delete"
)

type AuditLog struct {
	ID          int       `json:"id" db:"id"`
	StaffID     int       `json:"staff_id" db:"staff_id"`
	DataID      int       `json:"data_id" db:"data_id"`
	Action      string    `json:"action" db:"action"`
	EffectField string    `json:"effect_field" db:"effect_field"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type AuditLogResult struct {
	ID            int       `json:"id" db:"id"`
	Staff         string    `json:"staff" db:"staff"`
	LicenseNumber string    `json:"license_number" db:"license_number"`
	Action        string    `json:"action" db:"action"`
	EffectField   string    `json:"effect_field" db:"effect_field"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
