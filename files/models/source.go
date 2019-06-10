package models

import "time"

type SourceForm struct {
	SourceDate        string `json:"sourceDate"`
	DOI               string `json:"doi"`
	SourceDescription string `json:"sourceDescription"`
	BibText           string `json:"bibText"`
	FilePathURL       string `json:"filePathURL"`
}

type SourceInfo struct {
	SourceId          int       `json:"sourceId" db:"SourceId"`
	SourceDate        time.Time `json:"sourceDate" db:"SourceDate"`
	DOI               string    `json:"doi" db:"DOI"`
	SourceDescription string    `json:"sourceDescription" db:"SourceDescription"`
}

type Source struct {
	SourceId          int    `json:"sourceId" db:"SourceId"`
	SourceDate        string `json:"sourceDate" db:"SourceDate"`
	DOI               string `json:"doi" db:"DOI"`
	SourceDescription string `json:"sourceDescription" db:"SourceDescription"`
	BibText           string `json:"bibText" db:"BibText"`
	FilePathURL       string `json:"filePathURL" db:"FilePathURL"`
}
