package data

import (
	"context"
	"net/http"
	"time"

	ad "bss-beta-backend/audit_log"

	d "fx.prodigy9.co/data"
)

type CreateDatasForm struct {
	Payload struct {
		LicenseNumber       string    `json:"license_number"`
		DateRegistration    time.Time `json:"date_registration"`
		Province            string    `json:"province"`
		Type                string    `json:"type"`
		Characteristic      string    `json:"characteristic"`
		Brand               string    `json:"brand"`
		Model               string    `json:"model"`
		Year                int       `json:"year"`
		Color               string    `json:"color"`
		VehicleNumber       string    `json:"vehicle_number"`
		LocationVehicle     string    `json:"location_vehicle"`
		EngineBrand         string    `json:"engine_brand"`
		EngineNumber        string    `json:"engine_number"`
		LocationEngine      string    `json:"location_engine"`
		Fuel                string    `json:"fuel"`
		GasTankNumber       string    `json:"gas_tank_number"`
		CylinderCount       int       `json:"cylinder_count"`
		CC                  int       `json:"cc"`
		Horsepower          int       `json:"horsepower"`
		AxleAndWheelCount   int       `json:"axle_and_wheel_count"`
		VehicleWeight       float64   `json:"vehicle_weight"`
		LoadWeight          float64   `json:"load_weight"`
		TotalWeight         float64   `json:"total_weight"`
		SeatCount           int       `json:"seat_count"`
		LatestOwnerNumber   string    `json:"latest_owner_number"`
		OwnershipDate       time.Time `json:"ownership_date"`
		Owner               string    `json:"owner"`
		IDCardNumber        string    `json:"id_card_number"`
		BirthDate           time.Time `json:"birth_date"`
		Nationality         string    `json:"nationality"`
		Address             string    `json:"address"`
		Phone               string    `json:"phone"`
		LeaseContractNumber string    `json:"lease_contract_number"`
		ContractDate        time.Time `json:"contract_date"`
		LastTaxPaidDate     time.Time `json:"last_tax_paid_date"`
		TaxDueDate          time.Time `json:"tax_due_date"`
		ReceiptNumber       string    `json:"receipt_number"`
		TaxAmount           float64   `json:"tax_amount"`
		AdditionalAmount    float64   `json:"additional_amount"`
		Recorder            string    `json:"recorder"`
		Registrar           string    `json:"registrar"`
		LastUpdateDate      time.Time `json:"last_update_date"`
		PictureFront        string    `json:"picture_front"`
		PictureSide         string    `json:"picture_side"`
		PictureBack         string    `json:"picture_back"`
		CreatedBy           string    `json:"created_by"`
		UpdatedBy           string    `json:"updated_by"`
		StaffID             int       `json:"staff_id"`
	} `in:"body=json"`
}

func (c *CreateDatasForm) CreateData(ctx context.Context, resp *http.ResponseWriter, data *Data) (err error) {
	return d.Run(ctx, func(s d.Scope) error {
		query := `
			INSERT INTO datas (
				license_number, date_registration, province, type, characteristic, brand, model, year,
				color, vehicle_number, location_vehicle, engine_brand, engine_number, location_engine, fuel,
				gas_tank_number, cylinder_count, cc, horsepower, axle_and_wheel_count, vehicle_weight, load_weight,
				total_weight, seat_count, latest_owner_number, ownership_date, owner, id_card_number, birth_date,
				nationality, address, phone, lease_contract_number, contract_date, last_tax_paid_date, tax_due_date,
				receipt_number, tax_amount, additional_amount, recorder, registrar, last_update_date, picture_front,
				picture_side, picture_back, created_by, updated_by
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, 
				$21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, 
				$40, $41, $42, $43, $44, $45, $46, $47)
			RETURNING *
		`
		if err = s.Get(data, query, c.Payload.LicenseNumber, c.Payload.DateRegistration, c.Payload.Province,
			c.Payload.Type, c.Payload.Characteristic, c.Payload.Brand, c.Payload.Model, c.Payload.Year, c.Payload.Color,
			c.Payload.VehicleNumber, c.Payload.LocationVehicle, c.Payload.EngineBrand, c.Payload.EngineNumber, c.Payload.LocationEngine,
			c.Payload.Fuel, c.Payload.GasTankNumber, c.Payload.CylinderCount, c.Payload.CC, c.Payload.Horsepower,
			c.Payload.AxleAndWheelCount, c.Payload.VehicleWeight, c.Payload.LoadWeight, c.Payload.TotalWeight, c.Payload.SeatCount,
			c.Payload.LatestOwnerNumber, c.Payload.OwnershipDate, c.Payload.Owner, c.Payload.IDCardNumber, c.Payload.BirthDate,
			c.Payload.Nationality, c.Payload.Address, c.Payload.Phone, c.Payload.LeaseContractNumber, c.Payload.ContractDate,
			c.Payload.LastTaxPaidDate, c.Payload.TaxDueDate, c.Payload.ReceiptNumber, c.Payload.TaxAmount, c.Payload.AdditionalAmount,
			c.Payload.Recorder, c.Payload.Registrar, c.Payload.LastUpdateDate, c.Payload.PictureFront, c.Payload.PictureSide,
			c.Payload.PictureBack, c.Payload.CreatedBy, c.Payload.UpdatedBy,
		); err != nil {
			return err
		}

		createAuditLog := &ad.CreateAuditLogForm{
			StaffID:     c.Payload.StaffID,
			DataID:      data.ID,
			Action:      ad.ActionCreate,
			EffectField: "all",
		}
		auditLog := &ad.AuditLog{}
		if err := createAuditLog.CreateAuditLog(s.Context(), auditLog); err != nil {
			return err
		}

		return nil
	})
}
