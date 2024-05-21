package model

type Count struct {
	UserGuid string `json:"user_guid"`
	ZikrGuid string `json:"zikr_guid"`
	Count    int64  `json:"count"`
}
