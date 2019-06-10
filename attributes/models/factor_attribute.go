package models

import "time"

type FactorAttribute struct {
	FactorAttributeId int     `json:"attributeId,omitempty" db:"AttributeId"`
	FactorId          int     `json:"attributeEnumId" db:"AttributeEnumId"`
	SourceId          int     `json:"sourceId" db:"SourceId"`
	Value             float32 `json:"value" db:"Value"`
	Uncertainty       float32 `json:"uncertainty" db:"Uncertainty"`
	ValueUnitId1      int     `json:"valueUnitId1" db:"ValueUnitId1"`
	ValueUnitId2      int     `json:"valueUnitId2" db:"ValueUnitId2"`
	ValueUnitId3      int     `json:"valueUnitId3" db:"ValueUnitId3"`
}

type FactorAttributeWithName struct {
	FactorAttributeId   int     `json:"attributeId" db:"AttributeId"`
	AttributeEnumNamePl string  `json:"attributeEnumNamePl" db:"AttributeEnumNamePl"`
	Value               float32 `json:"value" db:"Value"`
	Uncertainty         float32 `json:"uncertainty" db:"Uncertainty"`
	ValueUnitId1        int     `json:"valueUnitId1" db:"ValueUnitId1"`
	ValueUnitId2        int     `json:"valueUnitId2" db:"ValueUnitId2"`
	ValueUnitId3        int     `json:"valueUnitId3" db:"ValueUnitId3"`
}

type FactorAttributeModel struct {
	FactorAttributeId int       `json:"attributeId" db:"AttributeId"`
	FactorId          int       `json:"attributeEnumId" db:"AttributeEnumId"`
	ResourceId        int       `json:"resourceId,omitempty" db:"ResourceId"`
	SourceId          int       `json:"sourceId" db:"SourceId"`
	SourceDate        time.Time `json:"sourceDate" db:"SourceDate"`
	SourceDescription string    `json:"sourceDescription" db:"SourceDescription"`
	FactorName        string    `json:"attributeEnumNamePl" db:"AttributeEnumNamePl"`
	Value             float32   `json:"value" db:"Value"`
	ValueUnitId1      *int      `json:"valueUnitId1" db:"ValueUnitId1"`
	ValueUnitId2      *int      `json:"valueUnitId2" db:"ValueUnitId2"`
	ValueUnitId3      *int      `json:"valueUnitId3" db:"ValueUnitId3"`
	UnitId1ShortName  *string   `json:"unit1shortname" db:"unit1shortname"`
	UnitId2ShortName  *string   `json:"Unit2ShortName" db:"unit2shortname"`
	UnitId3ShortName  *string   `json:"Unit3ShortName" db:"unit3shortname"`
	Uncertainty       float32   `json:"uncertainty" db:"Uncertainty"`
}
