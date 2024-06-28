package auditlog

import (
	"context"

	"fx.prodigy9.co/data"
)

type CreateAuditLogForm struct {
	StaffID     int    `json:"staff_id" db:"staff_id"`
	DataID      int    `json:"data_id" db:"data_id"`
	Action      string `json:"action" db:"action"`
	EffectField string `json:"effect_field" db:"effect_field"`
}

func (c *CreateAuditLogForm) CreateAuditLog(ctx context.Context, auditLog *AuditLog) error {
	query := `
		INSERT INTO audit_log (staff_id, data_id, action, effect_field)
		VALUES ($1, $2, $3, $4)
		RETURNING *
	`
	return data.Get(ctx, auditLog, query, c.StaffID, c.DataID, c.Action, c.EffectField)
}
