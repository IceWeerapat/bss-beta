package data

import (
	"context"
	"fmt"
	"time"

	"fx.prodigy9.co/data"
	"fx.prodigy9.co/httpserver/httperrors"
	"fx.prodigy9.co/httpserver/render"
)

type ListDatasForm struct {
}

type ListDatasByCriteriaForm struct {
	StaffID   int    `in:"query=staff_id" qb:"col='a.staff_id';op='='"`
	StartDate string `in:"query=start_date" qb:"col='d.created_at';op='>='"`
	EndDate   string `in:"query=end_date" qb:"col='d.created_at';op='<='"`
	OrderBy   string `validate:"oneof=created_at updated_at" in:"query=order_by" qb:"order"`
	Sort      string `validate:"oneof=desc asc" in:"query=sort" qb:"sort_direction"`
	Page      int    `in:"query=page;default=1" qb:"page"`
	Limit     int    `in:"query=limit;default=0" qb:"limit"`
}

type DataWithStaffResponse struct {
	ID                  int       `json:"id" db:"id"`
	LicenseNumber       string    `json:"license_number" db:"license_number"`
	DateRegistration    time.Time `json:"date_registration" db:"date_registration"`
	Province            string    `json:"province" db:"province"`
	Type                string    `json:"type" db:"type"`
	Characteristic      string    `json:"characteristic" db:"characteristic"`
	Brand               string    `json:"brand" db:"brand"`
	Model               string    `json:"model" db:"model"`
	Year                int       `json:"year" db:"year"`
	Color               string    `json:"color" db:"color"`
	VehicleNumber       string    `json:"vehicle_number" db:"vehicle_number"`
	LocationVehicle     string    `json:"location_vehicle" db:"location_vehicle"`
	EngineBrand         string    `json:"engine_brand" db:"engine_brand"`
	EngineNumber        string    `json:"engine_number" db:"engine_number"`
	LocationEngine      string    `json:"location_engine" db:"location_engine"`
	Fuel                string    `json:"fuel" db:"fuel"`
	GasTankNumber       string    `json:"gas_tank_number" db:"gas_tank_number"`
	CylinderCount       int       `json:"cylinder_count" db:"cylinder_count"`
	CC                  int       `json:"cc" db:"cc"`
	Horsepower          int       `json:"horsepower" db:"horsepower"`
	AxleAndWheelCount   int       `json:"axle_and_wheel_count" db:"axle_and_wheel_count"`
	VehicleWeight       float64   `json:"vehicle_weight" db:"vehicle_weight"`
	LoadWeight          float64   `json:"load_weight" db:"load_weight"`
	TotalWeight         float64   `json:"total_weight" db:"total_weight"`
	SeatCount           int       `json:"seat_count" db:"seat_count"`
	LatestOwnerNumber   string    `json:"latest_owner_number" db:"latest_owner_number"`
	OwnershipDate       time.Time `json:"ownership_date" db:"ownership_date"`
	Owner               string    `json:"owner" db:"owner"`
	IDCardNumber        string    `json:"id_card_number" db:"id_card_number"`
	BirthDate           time.Time `json:"birth_date" db:"birth_date"`
	Nationality         string    `json:"nationality" db:"nationality"`
	Address             string    `json:"address" db:"address"`
	Phone               string    `json:"phone" db:"phone"`
	LeaseContractNumber string    `json:"lease_contract_number" db:"lease_contract_number"`
	ContractDate        time.Time `json:"contract_date" db:"contract_date"`
	LastTaxPaidDate     time.Time `json:"last_tax_paid_date" db:"last_tax_paid_date"`
	TaxDueDate          time.Time `json:"tax_due_date" db:"tax_due_date"`
	ReceiptNumber       string    `json:"receipt_number" db:"receipt_number"`
	TaxAmount           float64   `json:"tax_amount" db:"tax_amount"`
	AdditionalAmount    float64   `json:"additional_amount" db:"additional_amount"`
	Recorder            string    `json:"recorder" db:"recorder"`
	Registrar           string    `json:"registrar" db:"registrar"`
	LastUpdateDate      time.Time `json:"last_update_date" db:"last_update_date"`
	PictureFront        string    `json:"picture_front" db:"picture_front"`
	PictureSide         string    `json:"picture_side" db:"picture_side"`
	PictureBack         string    `json:"picture_back" db:"picture_back"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	CreatedBy           string    `json:"created_by" db:"created_by"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy           string    `json:"updated_by" db:"updated_by"`
	StaffID             int       `json:"staff_id" db:"staff_id"`
}

func (l *ListDatasForm) ListData(ctx context.Context, datas *[]*Data) (err error) {
	query := `
		SELECT * 
		FROM datas
		ORDER BY ID ASC
	`
	return data.Select(ctx, datas, query)
}

func (l *ListDatasByCriteriaForm) ListDataByCriteria(ctx context.Context, datasRes *[]*DataWithStaffResponse, page *render.Pagination) (err error) {

	selectField := `SELECT d.*, a.staff_id `
	joinedQuery := `datas d LEFT JOIN audit_log a ON d.id = a.data_id  `
	qb := &data.QueryBuilder{Data: *l}
	whereDefault := " AND a.action='create' "
	whereClause := qb.Where(" AND ")

	if err = page.Execute(ctx, joinedQuery, whereClause.Sql+whereDefault, qb.Args); err != nil {
		return httperrors.NewError(httperrors.ErrInternal.Code(), err.Error())
	}

	query := fmt.Sprintf("%s FROM %s %s %s", selectField, joinedQuery, whereDefault, whereClause.Page().Sql)
	if err = data.Select(ctx, datasRes, query, qb.Args...); err != nil {
		return httperrors.NewError(httperrors.ErrInternal.Code(), err.Error())
	}

	return

}
