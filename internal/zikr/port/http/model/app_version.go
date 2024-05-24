package model

type AppVersion struct {
	Anndroid    string `json:"android_version"`
	Ios         string `json:"ios_version"`
	ForceUpdate bool   `json:"force_update"`
}
