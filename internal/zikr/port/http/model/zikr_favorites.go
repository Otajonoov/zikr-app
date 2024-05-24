package model

type IsFavorite struct {
	UserGuId string `json:"user_guid"`
	ZikrGuid string `json:"zikr_guid"`
	IsFav    bool   `json:"is_favorite"`
}
