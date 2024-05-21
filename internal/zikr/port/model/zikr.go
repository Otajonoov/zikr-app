package model

type Zikr struct {
	Arabic    string `json:"arabic"`
	Uzbek     string `json:"uzbek"`
	Pronounce string `json:"pronounce"`
}

type Response struct {
	Result string `json:"result"`
}

type PatchCount struct {
	Count int
}

type GetZikr struct {
	Guid       string `json:"guid"`
	Arabic     string `json:"arabic"`
	Uzbek      string `json:"uzbek"`
	Pronounce  string `json:"pronounce"`
	Count      int    `json:"count"`
	IsFavorite bool   `json:"is_favorite"`
}

type Zikrs struct {
	Zikrs []GetZikr `json:"zikrs"`
}
