package model

type Zikr struct {
	Arabic    string `json:"arabic"`
	Uzbek     string `json:"uzbek"`
	Pronounce string `json:"pronounce"`
}

type GetZikr struct {
	Guid      string `json:"guid"`
	Arabic    string `json:"arabic"`
	Uzbek     string `json:"uzbek"`
	Pronounce string `json:"pronounce"`
}

type Zikrs struct {
	Zikrs []GetZikr `json:"zikrs"`
}

type Favorites struct {
	UserId int `json:"user_id"`
	ZikrId int `json:"zikr_id"`
}

type Id struct {
	Id int `json:"id"`
}
