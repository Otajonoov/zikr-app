package model

import "time"

type Zikr struct {
	Arabic    string
	Uzbek     string
	Pronounce string
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

type Count int

type GetZikr struct {
	Guid       string
	UserGuid   string
	Arabic     string
	Uzbek      string
	Pronounce  string
	Count      int
	IsFavorite bool
}

type Zikrs struct {
	Zikrs []GetZikr
}
