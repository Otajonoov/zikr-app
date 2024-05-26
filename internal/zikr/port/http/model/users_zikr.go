package model

type Count struct {
	UserGuid string `json:"user_guid"`
	ZikrGuid string `json:"zikr_guid"`
	Count    int64  `json:"count"`
}

type IsFavorite struct {
	UserGuId string `json:"user_guid"`
	ZikrGuid string `json:"zikr_guid"`
	IsFav    bool   `json:"is_favorite"`
}

type Reyting struct {
	Limit    int64  `json:"limit" binding:"required" default:"10"`
	Page     int64  `json:"page" binding:"required" default:"1"`
	ZikrGuid string `json:"zikr_guid"`
}

type ReytingInfo struct {
	UserGuid  string `json:"user_guid"`
	Username  string `json:"username"`
	ZikrCount int64  `json:"zikr_count"`
}

type ReytingResponse struct {
	Reytings *[]Reyting `json:"reytings"`
}
