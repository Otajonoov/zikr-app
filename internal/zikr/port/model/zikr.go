package model

import "time"

type Zikr struct {
	Arabic    string `json:"arabic"`
	Uzbek     string `json:"uzbek"`
	Pronounce string `json:"pronounce"`
}

type ZikrSave struct {
	Guid       string
	UserGuid   string
	Arabic     string
	Uzbek      string
	Pronounce  string
	Count      int
	IsFavorite bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PatchCount struct {
	Count int
}

type GetZikr struct {
	Guid       string `json:"guid"`
	UserEmail  string `json:"user_email"`
	Arabic     string `json:"arabic"`
	Uzbek      string `json:"uzbek"`
	Pronounce  string `json:"pronounce"`
	Count      int    `json:"count"`
	IsFavorite bool   `json:"is_favorite"`
}

type Zikrs struct {
	Zikrs []GetZikr
}
