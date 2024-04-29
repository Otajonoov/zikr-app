package model

type Zikr struct {
	Id         int
	UserId     int    `json:"user_id"`
	Arabic     string `json:"arabic"`
	Uzbek      string `json:"uzbek"`
	Pronounce  string `json:"pronounce"`
	IsFavorite bool   `json:"is_favorite"`
}

type GetZikr struct {
	Id        string `json:"id"`
	Arabic    string `json:"arabic"`
	Uzbek     string `json:"uzbek"`
	Pronounce string `json:"pronounce"`
}

type Favorites struct {
	UserId int `json:"user_id"`
	ZikrId int `json:"zikr_id"`
}

type Id struct {
	Id int `json:"id"`
}

type Zikrs struct {
	Zikrs []GetZikr `json:"zikrs"`
}
