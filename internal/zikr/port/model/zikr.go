package model

type Zikr struct {
	Arabic    string `json:"arabic"`
	Uzbek     string `json:"uzbek"`
	Pronounce string `json:"pronounce"`
}

type GetZikr struct {
	Id        string `json:"id"`
	Arabic    string `json:"arabic"`
	Uzbek     string `json:"uzbek"`
	Pronounce string `json:"pronounce"`
}

type Id struct {
	Id string `json:"id"`
}

type Zikrs struct {
	Zikrs []GetZikr `json:"zikrs"`
}
