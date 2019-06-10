package models

type Quantity struct {
	QuantityId     int    `json:"QuantityId,omitempty" db:"QuantityId"`
	QuantityNamePl string `json:"QuantityNamePl" db:"QuantityNamePl"`
	QuantityNameEn string `json:"QuantityNameEn" db:"QuantityNameEn"`
	BaseUnitID     int    `json:"BaseUnitID" db:"BaseUnitID"`
}

type GetQuantity struct {
	QuantityId     int    `json:"QuantityId,omitempty" db:"QuantityId"`
	QuantityNamePl string `json:"QuantityNamePl" db:"QuantityNamePl"`
	QuantityNameEn string `json:"QuantityNameEn" db:"QuantityNameEn"`
	UnitShortName  string `json:"unitShortName" db:"UnitShortName"`
	BaseUnitID     int    `json:"BaseUnitID" db:"BaseUnitID"`
}

type QuantityId struct {
	QuantityId     string `json:"quantityId" db:"QuantityId"`
	QuantityNamePl string `json:"quantityNamePl" db:"QuantityNamePl"`
	UnitShortName  string `json:"unitShortName" db:"UnitShortName"`
}
