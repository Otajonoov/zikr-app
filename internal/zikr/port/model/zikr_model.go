package model

type Zikr struct {
	Arabic    string `json:"arabic"`
	Uzbek     string `json:"uzbek"`
	Pronounce string `json:"pronounce"`
}

type Zikrs struct {
	Zikrs []Zikr `json:"zikrs"`
}
