package models

type GUSResource struct {
	GUSResourceId     int    `json:"GUSResourceId,omitempty" db:"GUSResourceId"`
	GUSResourceNamePl string `json:"GUSResourceNamePl"  db:"GUSResourceNamePl"`
	GUSResourceNameEn string `json:"GUSResourceNameEn,omitempty"  db:"GUSResourceNameEn"`
	GUSId             string `json:"GUSId" db:"GUSId"`
}

type GUSResourceId struct {
	GUSResourceId     int    `json:"GUSResourceId,omitempty" db:"GUSResourceId"`
	GUSResourceNamePl string `json:"GUSResourceNamePl"  db:"GUSResourceNamePl"`
	GUSId             string `json:"GUSId" db:"GUSId"`
}
