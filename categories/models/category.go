package models

type Category struct {
	CategoryNamePl        string `json:"categoryNamePl"`
	CategoryNameEn        string `json:"categoryNameEn"`
	CategoryDescriptionPl string `json:"categoryDescriptionPl"`
	CategoryDescriptionEn string `json:"categoryDescriptionEn"`
	CategoryParentsIds    []int  `json:"categoryParentsIds"`
}

type EditCategory struct {
	CategoryId            int    `json:"categoryId" db:"CategoryId"`
	CategoryNamePl        string `json:"categoryNamePl" db:"CategoryNamePl"`
	CategoryNameEn        string `json:"categoryNameEn" db:"CategoryNameEn"`
	CategoryDescriptionPl string `json:"categoryDescriptionPl" db:"CategoryDescriptionPl"`
	CategoryDescriptionEn string `json:"categoryDescriptionEn" db:"CategoryDescriptionEn"`
}

type CategoryForm struct {
	CategoryId            int            `json:"categoryId" db:"CategoryId"`
	CategoryNamePl        string         `json:"categoryNamePl" db:"CategoryNamePl"`
	CategoryNameEn        string         `json:"categoryNameEn" db:"CategoryNameEn"`
	CategoryDescriptionPl string         `json:"categoryDescriptionPl" db:"CategoryDescriptionPl"`
	CategoryDescriptionEn string         `json:"categoryDescriptionEn" db:"CategoryDescriptionEn"`
	CategoryParents       []EditCategory `json:"categoryParents" db:"CategoryParents"`
}

type Hierarchy struct {
	CategoryId int `json:"categoryId" db:"CategoryId"`
	ParentId   int `json:"parentId" db:"ParentId"`
}
