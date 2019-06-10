package models

type BasicUnit struct {
	UnitId        int     `json:"unitId,omitempty" db:"UnitId"`
	UnitNamePl    string  `json:"unitNamePl" db:"UnitNamePl"`
	UnitNameEn    string  `json:"unitNameEn" db:"UnitNameEn"`
	UnitShortName string  `json:"unitShortName" db:"UnitShortName"`
	Ratio         float32 `json:"ratio,omitempty" db:"Ratio"`
}
