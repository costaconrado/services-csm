// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "time"

type Proposal struct {
	ID              uint
	CustomerContact CustomerContact
	CreatedAt       time.Time `example:"2022-09-15T21:54:42.123Z"`
	Offerings       []Offering
	Stage           DealStage `example:"0"`
}

type CustomerContact struct {
	ID       uint
	Name     string `example:"John Doe"`
	Email    string `example:"john.doe@example.com"`
	Phone    string `example:"+55 (31) 9999 9999"`
	Customer Customer
}

type Customer struct {
	ID       uint
	Name     string `example:"Example Co."`
	Segment  string `example:"Enterprise"`
	Vertical string `example:"Banking and Securities"`
}

type Offering struct {
	ID        uint
	Name      string  `example:"Kubernetes Adminstration Training"`
	SKU       string  `example:"KUB101"`
	UnitValue float32 `example:"4100.50"`
	Quantity  uint    `example:"12"`
	Addons    []Offering
}

type DealStage uint8

const (
	STAGE_ENGAGE DealStage = iota
	STAGE_CLOSED_LOST
	STAGE_CLOSED_WON
)
