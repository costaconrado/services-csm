// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "time"

type Proposal struct {
	CustomerContact CustomerContact `json:"customerContact"`
	Customer        Customer        `json:"customer"`
	CreationDate    time.Time       `json:"creationDate"  example:"2022-09-15T21:54:42.123Z"`
	Offerings       []Offering      `json:"offerings"`
	Stage           DealStage       `json:"stage"  example:"0"`
}

type CustomerContact struct {
	Name     string   `json:"name"  example:"John Doe"`
	Email    string   `json:"email"  example:"john.doe@example.com"`
	Phone    string   `json:"phone"  example:"+55 (31) 9999 9999"`
	Customer Customer `json:"customer"`
}

type Customer struct {
	Name     string `json:"name"  example:"Example Co."`
	Segment  string `json:"segment"  example:"Enterprise"`
	Vertical string `json:"vertical"  example:"Banking and Securities"`
}

type Offering struct {
	Name      string     `json:"name"  example:"Kubernetes Adminstration Training"`
	SKU       string     `json:"sku"  example:"KUB101"`
	UnitValue float32    `json:"unitValue"  example:"4100.50"`
	Quantity  uint       `json:"quantity"  example:"12"`
	Addons    []Offering `json:"addons"`
}

type DealStage uint8

const (
	STAGE_ENGAGE DealStage = iota
	STAGE_CLOSED_LOST
	STAGE_CLOSED_WON
)
