package models

type FactorId struct {
	AttributeEnumId        int    `json:"attributeEnumId" db:"AttributeEnumId"`
	AttributeEnumShortName string `json:"attributeEnumShortName" db:"AttributeEnumShortName"`
}

type Factor struct {
	AttributeEnumId            int    `json:"attributeEnumId,omitempty" db:"AttributeEnumId"`
	AttributeEnumShortName     string `json:"attributeEnumShortName" db:"AttributeEnumShortName"`
	AttributeEnumNamePl        string `json:"attributeEnumNamePl" db:"AttributeEnumNamePl"`
	AttributeEnumNameEn        string `json:"attributeEnumNameEn" db:"AttributeEnumNameEn"`
	AttributeEnumDescriptionPl string `json:"attributeEnumDescriptionPl" db:"AttributeEnumDescriptionPl"`
	AttributeEnumDescriptionEn string `json:"attributeEnumDescriptionEn" db:"AttributeEnumDescriptionEn"`
}

type ResourceFactorId struct {
	ResourceId  int `json:"resourceId" db:"ResourceId"`
	AttributeId int `json:"attributeId" db:"AttributeId"`
}
