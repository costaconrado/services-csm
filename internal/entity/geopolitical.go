// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

type Country struct {
	NativeName  string `json:"nativeName"       example:"Brasil"`
	EnglishName string `json:"enlishName"  example:"Brazil"`
	Acronym     string `json:"acronym"     example:"BR"`
	RegionId    string `json:"regionId"  example:"1"`
	SubRegionId string `json:"subRegionId"  example:"1"`
}

type Region struct {
	Name    string `json:"Name"  example:"Latin America"`
	Acronym string `json:"acronym"     example:"LATAM"`
}

type SubRegion struct {
	Name     string `json:"Name"  example:"Latin America"`
	Acronym  string `json:"acronym"     example:"LATAM"`
	RegionId string `json:"regionId"  example:"1"`
}
