package data

import (
	"fmt"
	"reflect"
	"strings"
)

type QueryBuilder struct {
	Data interface{}
	Sql  string
	Args []interface{}
}

func (qb *QueryBuilder) Where(separator ...string) *QueryBuilder {
	var sep string
	var conditions []string
	var values []interface{}

	if len(separator) > 0 {
		sep = separator[0]
	}

	valueOfFilter := reflect.ValueOf(qb.Data)
	if valueOfFilter.Kind() != reflect.Struct {
		return nil
	}

	typeOfFilter := reflect.TypeOf(qb.Data)

	for i := 0; i < typeOfFilter.NumField(); i++ {
		field := typeOfFilter.Field(i)
		fieldValue := valueOfFilter.Field(i)

		if fieldValue.IsZero() {
			continue
		}

		fieldName := field.Tag.Get("qb")
		if fieldName == "" {
			continue
		}

		qbTags := strings.Split(fieldName, ";")
		var cols []string
		op := "="
		for _, tag := range qbTags {
			if strings.HasPrefix(tag, "col=") {
				col := strings.TrimPrefix(tag, "col=")
				col = strings.Trim(col, "'")
				cols = strings.Split(col, ",")
			} else if strings.HasPrefix(tag, "op=") {
				op = strings.TrimPrefix(tag, "op=")
				op = strings.Trim(op, "'")
			}
		}
		if len(cols) == 0 {
			continue
		}

		var subConditions []string
		for _, col := range cols {
			values = append(values, valueOfFilter.Field(i).Interface())
			subConditions = append(subConditions, fmt.Sprintf("%s %s $%d", col, op, len(values)))
		}
		conditions = append(conditions, fmt.Sprintf("(%s)", strings.Join(subConditions, " OR ")))
	}

	if len(conditions) > 0 {
		qb.Sql = "WHERE " + strings.Join(conditions, fmt.Sprintf(" %s ", sep))
	}
	qb.Args = values

	return qb
}

func (qb *QueryBuilder) Order() *QueryBuilder {
	valueOfFilter := reflect.ValueOf(qb.Data)
	if valueOfFilter.Kind() != reflect.Struct {
		return nil
	}

	typeOfFilter := reflect.TypeOf(qb.Data)
	var order, sortDirection string
	for i := 0; i < typeOfFilter.NumField(); i++ {
		field := typeOfFilter.Field(i)
		fieldName := field.Tag.Get("qb")
		fieldValue := valueOfFilter.Field(i)
		if fieldName == "" {
			continue
		}

		qbTags := strings.Split(fieldName, ";")
		for _, tag := range qbTags {
			if tag == "order" {
				order = fieldValue.String()
			} else if tag == "sort_direction" {
				sortDirection = fieldValue.String()
			}
		}

	}
	if order == "" {
		return qb
	}
	qb.Sql = fmt.Sprintf("%s ORDER BY %s %s", qb.Sql, order, sortDirection)

	return qb
}

func (qb *QueryBuilder) Page() *QueryBuilder {
	valueOfFilter := reflect.ValueOf(qb.Data)
	if valueOfFilter.Kind() != reflect.Struct {
		return nil
	}

	typeOfFilter := reflect.TypeOf(qb.Data)
	var page, limit int
	for i := 0; i < typeOfFilter.NumField(); i++ {
		field := typeOfFilter.Field(i)
		fieldName := field.Tag.Get("qb")
		fieldValue := valueOfFilter.Field(i)
		if fieldName == "" {
			continue
		}

		qbTags := strings.Split(fieldName, ";")
		for _, tag := range qbTags {
			if tag == "page" {
				page = int(fieldValue.Int())
			} else if tag == "limit" {
				limit = int(fieldValue.Int())
			}
		}

	}

	if limit > 0 {
		offset := (page - 1) * limit
		qb.Sql = fmt.Sprintf(`%s LIMIT %d OFFSET %d`, qb.Sql, limit, offset)
	}

	return qb
}

func (qb *QueryBuilder) GroupBy(fields []string) *QueryBuilder {
	if len(fields) == 0 {
		return qb
	}
	qb.Sql = fmt.Sprintf("%s GROUP BY %s", qb.Sql, strings.Join(fields, ", "))
	return qb
}
