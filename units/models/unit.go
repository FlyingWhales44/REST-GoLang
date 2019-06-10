package models

type Unit struct {
	UnitId        int     `json:"unitId,omitempty" db:"UnitId"`
	UnitNamePl    string  `json:"unitNamePl" db:"UnitNamePl"`
	UnitNameEn    string  `json:"unitNameEn" db:"UnitNameEn"`
	QuantityId    int     `json:"quantityId" db:"QuantityId"`
	Ratio         float32 `json:"ratio" db:"Ratio"`
	UnitShortName string  `json:"unitShortName" db:"UnitShortName"`
}

type GetUnit struct {
	UnitId         int     `json:"unitId,omitempty" db:"UnitId"`
	QuantityNamePl string  `json:"QuantityNamePl" db:"QuantityNamePl"`
	UnitNamePl     string  `json:"unitNamePl" db:"UnitNamePl"`
	UnitNameEn     string  `json:"unitNameEn" db:"UnitNameEn"`
	QuantityId     int     `json:"quantityId" db:"QuantityId"`
	Ratio          float32 `json:"ratio" db:"Ratio"`
	UnitShortName  string  `json:"unitShortName" db:"UnitShortName"`
}

type UnitId struct {
	UnitId        int    `json:"unitId,omitempty" db:"UnitId"`
	UnitShortName string `json:"unitShortName" db:"UnitShortName"`
}
