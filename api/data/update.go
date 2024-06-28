package data

import (
	"context"
	"net/http"
	"time"

	ad "bss-beta-backend/audit_log"

	d "fx.prodigy9.co/data"
)

type UpdateDatasForm struct {
	ID      int `in:"path=id"`
	Payload struct {
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
		UpdatedBy           string    `json:"updated_by" db:"updated_by"`
		StaffID             int       `json:"staff_id"`
		EffectField         string    `json:"effect_field"`
	} `in:"body=json"`
}

func (u *UpdateDatasForm) UpdateData(ctx context.Context, resp *http.ResponseWriter, data *Data) error {
	return d.Run(ctx, func(s d.Scope) error {
		updateQuery := `
        UPDATE datas
            SET license_number=$1, date_registration=$2, province=$3, type=$4, characteristic=$5, brand=$6,
                model=$7, year=$8, color=$9, vehicle_number=$10, location_vehicle=$11, engine_brand=$12, 
                engine_number=$13, location_engine=$14, fuel=$15, gas_tank_number=$16, cylinder_count=$17, cc=$18, 
                horsepower=$19, axle_and_wheel_count=$20, vehicle_weight=$21, load_weight=$22, total_weight=$23, 
                seat_count=$24, latest_owner_number=$25, ownership_date=$26, owner=$27, id_card_number=$28, 
                birth_date=$29, nationality=$30, address=$31, phone=$32, lease_contract_number=$33, 
                contract_date=$34, last_tax_paid_date=$35, tax_due_date=$36, receipt_number=$37, tax_amount=$38,
                additional_amount=$39, recorder=$40, registrar=$41, last_update_date=$42, picture_front=$43,
                picture_side=$44, picture_back=$45, updated_by=$46
            WHERE id=$47
            RETURNING *
        `
		if err := s.Get(data, updateQuery, u.Payload.LicenseNumber, u.Payload.DateRegistration, u.Payload.Province,
			u.Payload.Type, u.Payload.Characteristic, u.Payload.Brand, u.Payload.Model, u.Payload.Year,
			u.Payload.Color, u.Payload.VehicleNumber, u.Payload.LocationVehicle, u.Payload.EngineBrand,
			u.Payload.EngineNumber, u.Payload.LocationEngine, u.Payload.Fuel, u.Payload.GasTankNumber,
			u.Payload.CylinderCount, u.Payload.CC, u.Payload.Horsepower, u.Payload.AxleAndWheelCount,
			u.Payload.VehicleWeight, u.Payload.LoadWeight, u.Payload.TotalWeight, u.Payload.SeatCount,
			u.Payload.LatestOwnerNumber, u.Payload.OwnershipDate, u.Payload.Owner, u.Payload.IDCardNumber,
			u.Payload.BirthDate, u.Payload.Nationality, u.Payload.Address, u.Payload.Phone,
			u.Payload.LeaseContractNumber, u.Payload.ContractDate, u.Payload.LastTaxPaidDate,
			u.Payload.TaxDueDate, u.Payload.ReceiptNumber, u.Payload.TaxAmount, u.Payload.AdditionalAmount,
			u.Payload.Recorder, u.Payload.Registrar, u.Payload.LastUpdateDate, u.Payload.PictureFront,
			u.Payload.PictureSide, u.Payload.PictureBack, u.Payload.UpdatedBy, u.ID); err != nil {
			return err
		}

		createAuditLog := &ad.CreateAuditLogForm{
			StaffID:     u.Payload.StaffID,
			DataID:      data.ID,
			Action:      ad.ActionUpdate,
			EffectField: u.Payload.EffectField,
		}

		auditLog := &ad.AuditLog{}
		if err := createAuditLog.CreateAuditLog(s.Context(), auditLog); err != nil {
			return err
		}

		return nil
	})
}

// package data

// import (
// 	"context"
// 	"net/http"
// 	"strings"
// 	"time"

// 	ad "bss-beta-backend/audit_log"

// 	d "fx.prodigy9.co/data"
// )

// type UpdateDatasForm struct {
// 	ID      int `in:"path=id"`
// 	Payload struct {
// 		LicenseNumber       string    `json:"license_number" db:"license_number"`
// 		DateRegistration    time.Time `json:"date_registration" db:"date_registration"`
// 		Province            string    `json:"province" db:"province"`
// 		Type                string    `json:"type" db:"type"`
// 		Characteristic      string    `json:"characteristic" db:"characteristic"`
// 		Brand               string    `json:"brand" db:"brand"`
// 		Model               string    `json:"model" db:"model"`
// 		Year                int       `json:"year" db:"year"`
// 		Color               string    `json:"color" db:"color"`
// 		VehicleNumber       string    `json:"vehicle_number" db:"vehicle_number"`
// 		LocationVehicle     string    `json:"location_vehicle" db:"location_vehicle"`
// 		EngineBrand         string    `json:"engine_brand" db:"engine_brand"`
// 		EngineNumber        string    `json:"engine_number" db:"engine_number"`
// 		LocationEngine      string    `json:"location_engine" db:"location_engine"`
// 		Fuel                string    `json:"fuel" db:"fuel"`
// 		GasTankNumber       string    `json:"gas_tank_number" db:"gas_tank_number"`
// 		CylinderCount       int       `json:"cylinder_count" db:"cylinder_count"`
// 		CC                  int       `json:"cc" db:"cc"`
// 		Horsepower          int       `json:"horsepower" db:"horsepower"`
// 		AxleAndWheelCount   int       `json:"axle_and_wheel_count" db:"axle_and_wheel_count"`
// 		VehicleWeight       float64   `json:"vehicle_weight" db:"vehicle_weight"`
// 		LoadWeight          float64   `json:"load_weight" db:"load_weight"`
// 		TotalWeight         float64   `json:"total_weight" db:"total_weight"`
// 		SeatCount           int       `json:"seat_count" db:"seat_count"`
// 		LatestOwnerNumber   string    `json:"latest_owner_number" db:"latest_owner_number"`
// 		OwnershipDate       time.Time `json:"ownership_date" db:"ownership_date"`
// 		Owner               string    `json:"owner" db:"owner"`
// 		IDCardNumber        string    `json:"id_card_number" db:"id_card_number"`
// 		BirthDate           time.Time `json:"birth_date" db:"birth_date"`
// 		Nationality         string    `json:"nationality" db:"nationality"`
// 		Address             string    `json:"address" db:"address"`
// 		Phone               string    `json:"phone" db:"phone"`
// 		LeaseContractNumber string    `json:"lease_contract_number" db:"lease_contract_number"`
// 		ContractDate        time.Time `json:"contract_date" db:"contract_date"`
// 		LastTaxPaidDate     time.Time `json:"last_tax_paid_date" db:"last_tax_paid_date"`
// 		TaxDueDate          time.Time `json:"tax_due_date" db:"tax_due_date"`
// 		ReceiptNumber       string    `json:"receipt_number" db:"receipt_number"`
// 		TaxAmount           float64   `json:"tax_amount" db:"tax_amount"`
// 		AdditionalAmount    float64   `json:"additional_amount" db:"additional_amount"`
// 		Recorder            string    `json:"recorder" db:"recorder"`
// 		Registrar           string    `json:"registrar" db:"registrar"`
// 		LastUpdateDate      time.Time `json:"last_update_date" db:"last_update_date"`
// 		PictureFront        string    `json:"picture_front" db:"picture_front"`
// 		PictureSide         string    `json:"picture_side" db:"picture_side"`
// 		PictureBack         string    `json:"picture_back" db:"picture_back"`
// 		UpdatedBy           string    `json:"updated_by" db:"updated_by"`
// 		StaffID             int       `json:"staff_id"`
// 	} `in:"body=json"`
// }

// func (u *UpdateDatasForm) UpdateData(ctx context.Context, resp *http.ResponseWriter, data *Data) error {
// 	return d.Run(ctx, func(s d.Scope) error {
// 		// Fetch original data
// 		originalData := &Data{}
// 		query := `SELECT * FROM datas WHERE id=$1`
// 		if err := s.Get(originalData, query, u.ID); err != nil {
// 			return err
// 		}

// 		// Update data
// 		updateQuery := `
//         UPDATE datas
//             SET license_number=$1, date_registration=$2, province=$3, type=$4, characteristic=$5, brand=$6,
//                 model=$7, year=$8, color=$9, vehicle_number=$10, location_vehicle=$11, engine_brand=$12,
//                 engine_number=$13, location_engine=$14, fuel=$15, gas_tank_number=$16, cylinder_count=$17, cc=$18,
//                 horsepower=$19, axle_and_wheel_count=$20, vehicle_weight=$21, load_weight=$22, total_weight=$23,
//                 seat_count=$24, latest_owner_number=$25, ownership_date=$26, owner=$27, id_card_number=$28,
//                 birth_date=$29, nationality=$30, address=$31, phone=$32, lease_contract_number=$33,
//                 contract_date=$34, last_tax_paid_date=$35, tax_due_date=$36, receipt_number=$37, tax_amount=$38,
//                 additional_amount=$39, recorder=$40, registrar=$41, last_update_date=$42, picture_front=$43,
//                 picture_side=$44, picture_back=$45, updated_by=$46
//             WHERE id=$47
//             RETURNING *
//         `
// 		if err := s.Get(data, updateQuery, u.Payload.LicenseNumber, u.Payload.DateRegistration, u.Payload.Province,
// 			u.Payload.Type, u.Payload.Characteristic, u.Payload.Brand, u.Payload.Model, u.Payload.Year,
// 			u.Payload.Color, u.Payload.VehicleNumber, u.Payload.LocationVehicle, u.Payload.EngineBrand,
// 			u.Payload.EngineNumber, u.Payload.LocationEngine, u.Payload.Fuel, u.Payload.GasTankNumber,
// 			u.Payload.CylinderCount, u.Payload.CC, u.Payload.Horsepower, u.Payload.AxleAndWheelCount,
// 			u.Payload.VehicleWeight, u.Payload.LoadWeight, u.Payload.TotalWeight, u.Payload.SeatCount,
// 			u.Payload.LatestOwnerNumber, u.Payload.OwnershipDate, u.Payload.Owner, u.Payload.IDCardNumber,
// 			u.Payload.BirthDate, u.Payload.Nationality, u.Payload.Address, u.Payload.Phone,
// 			u.Payload.LeaseContractNumber, u.Payload.ContractDate, u.Payload.LastTaxPaidDate,
// 			u.Payload.TaxDueDate, u.Payload.ReceiptNumber, u.Payload.TaxAmount, u.Payload.AdditionalAmount,
// 			u.Payload.Recorder, u.Payload.Registrar, u.Payload.LastUpdateDate, u.Payload.PictureFront,
// 			u.Payload.PictureSide, u.Payload.PictureBack, u.Payload.UpdatedBy, u.ID); err != nil {
// 			return err
// 		}

// 		// Compare original data with updated data
// 		var changedFields []string
// 		if originalData.LicenseNumber != u.Payload.LicenseNumber {
// 			changedFields = append(changedFields, "license_number")
// 		}
// 		if originalData.DateRegistration != u.Payload.DateRegistration {
// 			changedFields = append(changedFields, "date_registration")
// 		}
// 		if originalData.Province != u.Payload.Province {
// 			changedFields = append(changedFields, "province")
// 		}
// 		if originalData.Type != u.Payload.Type {
// 			changedFields = append(changedFields, "type")
// 		}
// 		if originalData.Characteristic != u.Payload.Characteristic {
// 			changedFields = append(changedFields, "characteristic")
// 		}
// 		if originalData.Brand != u.Payload.Brand {
// 			changedFields = append(changedFields, "brand")
// 		}
// 		if originalData.Model != u.Payload.Model {
// 			changedFields = append(changedFields, "model")
// 		}
// 		if originalData.Year != u.Payload.Year {
// 			changedFields = append(changedFields, "year")
// 		}
// 		if originalData.Color != u.Payload.Color {
// 			changedFields = append(changedFields, "color")
// 		}
// 		if originalData.VehicleNumber != u.Payload.VehicleNumber {
// 			changedFields = append(changedFields, "vehicle_number")
// 		}
// 		if originalData.LocationVehicle != u.Payload.LocationVehicle {
// 			changedFields = append(changedFields, "location_vehicle")
// 		}
// 		if originalData.EngineBrand != u.Payload.EngineBrand {
// 			changedFields = append(changedFields, "engine_brand")
// 		}
// 		if originalData.EngineNumber != u.Payload.EngineNumber {
// 			changedFields = append(changedFields, "engine_number")
// 		}
// 		if originalData.LocationEngine != u.Payload.LocationEngine {
// 			changedFields = append(changedFields, "location_engine")
// 		}
// 		if originalData.Fuel != u.Payload.Fuel {
// 			changedFields = append(changedFields, "fuel")
// 		}
// 		if originalData.GasTankNumber != u.Payload.GasTankNumber {
// 			changedFields = append(changedFields, "gas_tank_number")
// 		}
// 		if originalData.CylinderCount != u.Payload.CylinderCount {
// 			changedFields = append(changedFields, "cylinder_count")
// 		}
// 		if originalData.CC != u.Payload.CC {
// 			changedFields = append(changedFields, "cc")
// 		}
// 		if originalData.Horsepower != u.Payload.Horsepower {
// 			changedFields = append(changedFields, "horsepower")
// 		}
// 		if originalData.AxleAndWheelCount != u.Payload.AxleAndWheelCount {
// 			changedFields = append(changedFields, "axle_and_wheel_count")
// 		}
// 		if originalData.VehicleWeight != u.Payload.VehicleWeight {
// 			changedFields = append(changedFields, "vehicle_weight")
// 		}
// 		if originalData.LoadWeight != u.Payload.LoadWeight {
// 			changedFields = append(changedFields, "load_weight")
// 		}
// 		if originalData.TotalWeight != u.Payload.TotalWeight {
// 			changedFields = append(changedFields, "total_weight")
// 		}
// 		if originalData.SeatCount != u.Payload.SeatCount {
// 			changedFields = append(changedFields, "seat_count")
// 		}
// 		if originalData.LatestOwnerNumber != u.Payload.LatestOwnerNumber {
// 			changedFields = append(changedFields, "latest_owner_number")
// 		}
// 		if originalData.OwnershipDate != u.Payload.OwnershipDate {
// 			changedFields = append(changedFields, "ownership_date")
// 		}
// 		if originalData.Owner != u.Payload.Owner {
// 			changedFields = append(changedFields, "owner")
// 		}
// 		if originalData.IDCardNumber != u.Payload.IDCardNumber {
// 			changedFields = append(changedFields, "id_card_number")
// 		}
// 		if originalData.BirthDate != u.Payload.BirthDate {
// 			changedFields = append(changedFields, "birth_date")
// 		}
// 		if originalData.Nationality != u.Payload.Nationality {
// 			changedFields = append(changedFields, "nationality")
// 		}
// 		if originalData.Address != u.Payload.Address {
// 			changedFields = append(changedFields, "address")
// 		}
// 		if originalData.Phone != u.Payload.Phone {
// 			changedFields = append(changedFields, "phone")
// 		}
// 		if originalData.LeaseContractNumber != u.Payload.LeaseContractNumber {
// 			changedFields = append(changedFields, "lease_contract_number")
// 		}
// 		if originalData.ContractDate != u.Payload.ContractDate {
// 			changedFields = append(changedFields, "contract_date")
// 		}
// 		if originalData.LastTaxPaidDate != u.Payload.LastTaxPaidDate {
// 			changedFields = append(changedFields, "last_tax_paid_date")
// 		}
// 		if originalData.TaxDueDate != u.Payload.TaxDueDate {
// 			changedFields = append(changedFields, "tax_due_date")
// 		}
// 		if originalData.ReceiptNumber != u.Payload.ReceiptNumber {
// 			changedFields = append(changedFields, "receipt_number")
// 		}
// 		if originalData.TaxAmount != u.Payload.TaxAmount {
// 			changedFields = append(changedFields, "tax_amount")
// 		}
// 		if originalData.AdditionalAmount != u.Payload.AdditionalAmount {
// 			changedFields = append(changedFields, "additional_amount")
// 		}
// 		if originalData.Recorder != u.Payload.Recorder {
// 			changedFields = append(changedFields, "recorder")
// 		}
// 		if originalData.Registrar != u.Payload.Registrar {
// 			changedFields = append(changedFields, "registrar")
// 		}
// 		if originalData.LastUpdateDate != u.Payload.LastUpdateDate {
// 			changedFields = append(changedFields, "last_update_date")
// 		}
// 		if originalData.PictureFront != u.Payload.PictureFront {
// 			changedFields = append(changedFields, "picture_front")
// 		}
// 		if originalData.PictureSide != u.Payload.PictureSide {
// 			changedFields = append(changedFields, "picture_side")
// 		}
// 		if originalData.PictureBack != u.Payload.PictureBack {
// 			changedFields = append(changedFields, "picture_back")
// 		}
// 		if originalData.UpdatedBy != u.Payload.UpdatedBy {
// 			changedFields = append(changedFields, "updated_by")
// 		}

// 		// Create audit log
// 		createAuditLog := &ad.CreateAuditLogForm{
// 			StaffID:     u.Payload.StaffID,
// 			DataID:      data.ID,
// 			Action:      ad.ActionUpdate,
// 			EffectField: strings.Join(changedFields, "|"),
// 		}
// 		auditLog := &ad.AuditLog{}
// 		if err := createAuditLog.CreateAuditLog(s.Context(), auditLog); err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// }
