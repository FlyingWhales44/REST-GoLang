package models

import "time"

type EnergyResource struct {
	EnergyResourceAttributeId   int     `json:"energyResourceAttributeId,omitempty"  db:"EnergyResourceAttributeId"`
	EnergyResourceAttributeName string  `json:"energyResourceAttributeName" db:"EnergyResourceAttributeName"`
	GUSResourceId               int     `json:"GUSResourceId" db:"GUSResourceId"`
	SourceId                    int     `json:"sourceId" db:"SourceId"`
	CO2Value                    float32 `json:"CO2Value" db:"CO2Value"`
	NCVValue                    float32 `json:"NCVValue" db:"NCVValue"`
	CO2UnitId                   int     `json:"CO2UnitId" db:"CO2UnitId"`
}

type EnergyResourceAttribute struct {
	EnergyResourceAttributeId int     `json:"energyResourceAttributeId"  db:"EnergyResourceAttributeId"`
	EnergyResourceId          int     `json:"energyResourceId" db:"EnergyResourceId"`
	SourceId                  int     `json:"sourceId" db:"SourceId"`
	CO2Value                  float32 `json:"CO2Value" db:"CO2Value"`
	NCVValue                  float32 `json:"NCVValue" db:"NCVValue"`
	CO2UnitId                 int     `json:"CO2UnitId" db:"CO2UnitId"`
}

type EnergyResourceAttributeEdit struct {
	EnergyResourceAttributeId int     `json:"energyResourceAttributeId"  db:"EnergyResourceAttributeId"`
	EnergyResourceId          int     `json:"energyResourceId" db:"EnergyResourceId"`
	SourceId                  int     `json:"sourceId" db:"SourceId"`
	CO2Value                  float32 `json:"CO2Value" db:"CO2Value"`
	NCVValue                  float32 `json:"NCVValue" db:"NCVValue"`
	CO2UnitId                 int     `json:"CO2UnitId" db:"CO2UnitId"`
	EnergyResourceName        string  `json:"energyResourceName" db:"EnergyResourceName"`
	GUSResourceId             int     `json:"GUSResourceId" db:"GUSResourceId"`
}

type EnergyResourceAttributeForm struct {
	EnergyResourceAttributeId int       `json:"energyResourceAttributeId"  db:"EnergyResourceAttributeId"`
	EnergyResourceId          int       `json:"energyResourceId" db:"EnergyResourceId"`
	EnergyResourceName        string    `json:"energyResourceName" db:"EnergyResourceName"`
	GUSResourceId             int       `json:"GUSResourceId" db:"GUSResourceId"`
	GUSId                     int       `json:"GUSId" db:"GUSId"`
	SourceId                  int       `json:"sourceId" db:"SourceId"`
	SourceDescription         string    `json:"sourceDescription" db:"SourceDescription"`
	SourceDate                time.Time `json:"sourceDate" db:"SourceDate"`
	CO2Value                  float32   `json:"CO2Value" db:"CO2Value"`
	NCVValue                  float32   `json:"NCVValue" db:"NCVValue"`
	CO2UnitId                 int       `json:"CO2UnitId" db:"CO2UnitId"`
	Ratio                     float32   `json:"ratio" db:"Ratio"`
	UnitShortName             string    `json:"unitShortName" db:"UnitShortName"`
}

type EnergyResourceEdit struct {
	EnergyResourceId   int    `json:"energyResourceId"  db:"EnergyResourceId"`
	EnergyResourceName string `json:"energyResourceName" db:"EnergyResourceName"`
	GUSResourceId      int    `json:"GUSResourceId" db:"GUSResourceId"`
}
