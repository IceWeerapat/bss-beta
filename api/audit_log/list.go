package auditlog

import (
	"context"

	"fx.prodigy9.co/data"
)

type ListAuditLogForm struct {
}

func (l *ListAuditLogForm) ListAuditLogs(ctx context.Context, auditLogs *[]*AuditLogResult) (err error) {
	query := `
		SELECT DISTINCT al.id, 
                CONCAT(s.first_name, ' ', s.last_name) AS staff, 
                d.license_number, 
                al.action, 
                al.effect_field, 
                al.created_at 
		FROM audit_log al
		LEFT JOIN staffs s ON al.staff_id = s.id
		LEFT JOIN datas d ON al.data_id = d.id
		ORDER BY al.id ASC
	`
	return data.Select(ctx, auditLogs, query)
}
