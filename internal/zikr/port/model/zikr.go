package model

type Zikr struct {
	Arabic    string `json:"arabic"`
	Uzbek     string `json:"uzbek"`
	Pronounce string `json:"pronounce"`
}

type PatchCount struct {
	Guid     string `json:"guid"`
	UserGuid string `json:"userGuid"`
	Count    int    `json:"count"`
}

type GetZikr struct {
	Guid       string `json:"guid"`
	Arabic     string `json:"arabic"`
	Uzbek      string `json:"uzbek"`
	Pronounce  string `json:"pronounce"`
	Count      int    `json:"count"`
	IsFavorite bool   `json:"isfavorite"`
}

type Zikrs struct {
	Zikrs []GetZikr `json:"zikrs"`
}
