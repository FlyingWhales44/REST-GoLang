package models

type Resource struct {
	ResourceNamePl        string `json:"resourceNamePl"`
	ResourceNameEn        string `json:"resourceNameEn"`
	ResourceDescriptionPl string `json:"resourceDescriptionPl"`
	ResourceDescriptionEn string `json:"resourceDescriptionEn"`
	AttributeIds          []int  `json:"attributeIds"`
}

type AttributeRes struct {
	ResourceId  int `json:"resourceId" db:"ResourceId"`
	AttributeId int `json:"attributeId" db:"AttributeId"`
}

type ResourceEdit struct {
	ResourceId            int    `json:"resourceId" db:"ResourceId"`
	ResourceNamePl        string `json:"resourceNamePl" db:"ResourceNamePl"`
	ResourceNameEn        string `json:"resourceNameEn" db:"ResourceNameEn"`
	ResourceDescriptionPl string `json:"resourceDescriptionPl" db:"ResourceDescriptionPl"`
	ResourceDescriptionEn string `json:"resourceDescriptionEn" db:"ResourceDescriptionEn"`
}

type ResourceModel struct {
	ResourceId            int                    `json:"resourceId" db:"ResourceId"`
	ResourceNamePl        string                 `json:"resourceNamePl" db:"ResourceNamePl"`
	ResourceNameEn        string                 `json:"resourceNameEn" db:"ResourceNameEn"`
	ResourceDescriptionPl string                 `json:"resourceDescriptionPl" db:"ResourceDescriptionPl"`
	ResourceDescriptionEn string                 `json:"resourceDescriptionEn" db:"ResourceDescriptionEn"`
	Attributes            []FactorAttributeModel `json:"attributes" db:"Attributes"`
}
