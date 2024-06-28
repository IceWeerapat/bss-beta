package data

import (
	"context"
	"strconv"
	"time"

	"fx.prodigy9.co/data"
)

type SummaryDataForm struct {
	StaffID string `in:"path=staff_id"`
}

func (s *SummaryDataForm) Summary(ctx context.Context, summary *SummaryResponse) (err error) {

	today := time.Now().Truncate(24 * time.Hour)
	weekStart := today.AddDate(0, 0, -int(today.Weekday()))
	monthStart := today.AddDate(0, 0, -int(today.Day())+1)

	queryBase := `
		SELECT  COUNT(DISTINCT d.id)
		FROM datas d
		LEFT JOIN audit_log a ON d.id = a.data_id
		WHERE a.action = 'create' AND a.staff_id = $1 AND d.created_at >= $2 AND d.created_at < $3
	`
	if s.StaffID == "" {
		s.StaffID = "0"
	}
	staffID, err := strconv.Atoi(s.StaffID)
	if err != nil {
		return err
	}

	if staffID != 0 {
		err = data.Get(ctx, &summary.Today, queryBase, staffID, today, today.AddDate(0, 0, 1))
		if err != nil {
			return err
		}

		err = data.Get(ctx, &summary.ThisWeek, queryBase, staffID, weekStart, weekStart.AddDate(0, 0, 7))
		if err != nil {
			return err
		}

		err = data.Get(ctx, &summary.ThisMonth, queryBase, staffID, monthStart, monthStart.AddDate(0, 1, 0))
		if err != nil {
			return err
		}

		err = data.Get(ctx, &summary.Total, `
			SELECT COUNT(DISTINCT d.id)
			FROM datas d
			LEFT JOIN audit_log a ON d.id = a.data_id
			WHERE a.action = 'create' AND a.staff_id = $1
		`, s.StaffID)
	} else {
		queryBase := `
			SELECT COUNT(DISTINCT d.id)
			FROM datas d
			LEFT JOIN audit_log a ON d.id = a.data_id
			WHERE a.action = 'create' AND d.created_at >= $1 AND d.created_at < $2
		`
		err = data.Get(ctx, &summary.Today, queryBase, today, today.AddDate(0, 0, 1))
		if err != nil {
			return err
		}
		err = data.Get(ctx, &summary.ThisWeek, queryBase, weekStart, weekStart.AddDate(0, 0, 7))
		if err != nil {
			return err
		}
		err = data.Get(ctx, &summary.ThisMonth, queryBase, monthStart, monthStart.AddDate(0, 1, 0))
		if err != nil {
			return err
		}
		err = data.Get(ctx, &summary.Total, `
			SELECT COUNT(*)
			FROM datas
		`)
	}

	return err
}
