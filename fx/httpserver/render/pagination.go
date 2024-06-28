package render

import (
	"context"
	"fmt"

	"fx.prodigy9.co/data"
)

type Pagination struct {
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	TotalItems   int `json:"total_items"`
	LimitPerPage int `json:"limit_per_page"`
}

func (p *Pagination) Execute(ctx context.Context, table string, condition string, args []interface{}) (err error) {
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s %s`, table, condition)
	if err = data.Get(ctx, &p.TotalItems, query, args...); err != nil {
		return
	}

	p.TotalPages = (p.TotalItems + p.LimitPerPage - 1) / p.LimitPerPage
	return
}
