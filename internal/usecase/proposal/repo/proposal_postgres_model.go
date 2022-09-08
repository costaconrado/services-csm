package repo

import (
	"time"

	"gorm.io/gorm"
)

type Proposal struct {
	gorm.Model
	CreationDate      time.Time  `json:"creationDate"  example:"2022-09-15T21:54:42.123Z"`
	Offerings         []Offering `json:"offerings"`
	Stage             DealStage  `json:"stage"  example:"0"`
	CustomerID        uint       `json:"CustomerID"`
	CustomerContactID uint       `json:"CustomerContactID"`
}

type CustomerContact struct {
	gorm.Model
	Name       string     `json:"name"  example:"John Doe"`
	Email      string     `json:"email"  example:"john.doe@example.com"`
	Phone      string     `json:"phone"  example:"+55 (31) 9999 9999"`
	Proposals  []Proposal `json:"proposal"`
	CustomerID uint       `json:"CustomerID"`
}

type Customer struct {
	gorm.Model
	Name             string            `json:"name"  example:"Example Co."`
	Segment          string            `json:"segment"  example:"Enterprise"`
	Vertical         string            `json:"vertical"  example:"Banking and Securities"`
	Proposals        []Proposal        `json:"proposal"`
	CustomerContacts []CustomerContact `json:"-"`
}

type Offering struct {
	gorm.Model
	Name             string     `json:"name"  example:"Kubernetes Adminstration Training"`
	SKU              string     `json:"sku"  example:"KUB101"`
	UnitValue        float32    `json:"unitValue"  example:"4100.50"`
	Quantity         uint       `json:"quantity"  example:"12"`
	Addons           []Offering `json:"addons" gorm:"foreignKey:ParentOfferingID"`
	ProposalID       uint       `json:"proposalID"`
	ParentOfferingID uint       `json:"parentOfferingID"`
}

type DealStage uint8

const (
	STAGE_ENGAGE DealStage = iota
	STAGE_CLOSED_LOST
	STAGE_CLOSED_WON
)
