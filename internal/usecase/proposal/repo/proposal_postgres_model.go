package repo

import (
	"encoding/json"
	"time"

	"github.com/costaconrado/services-csm/internal/entity"
	"gorm.io/gorm"
)

type CustomerContact struct {
	gorm.Model
	Name       string
	Email      string
	Phone      string
	CustomerID *uint
	Customer   *Customer
}

type Proposal struct {
	gorm.Model
	Offerings         []Offering
	Stage             entity.DealStage
	CustomerContactID *uint
	CustomerContact   *CustomerContact
	CreatedAt         time.Time
}

type Customer struct {
	gorm.Model
	Name     string
	Segment  string
	Vertical string
}

type Offering struct {
	gorm.Model
	Name             string
	SKU              string
	UnitValue        float32
	Quantity         uint
	Addons           []Offering `gorm:"foreignKey:ParentOfferingID"`
	ProposalID       uint
	ParentOfferingID *uint
}

func (model *CustomerContact) Unmarshal() (entity.CustomerContact, error) {
	var ent entity.CustomerContact
	modelJson, marshalError := json.Marshal(model)
	if marshalError != nil {
		return ent, marshalError
	}

	unmarshalError := json.Unmarshal(modelJson, &ent)
	if unmarshalError != nil {
		return ent, unmarshalError
	}
	return ent, nil
}

func (model *Proposal) Unmarshal() (entity.Proposal, error) {
	var ent entity.Proposal
	modelJson, marshalError := json.Marshal(model)
	if marshalError != nil {
		return ent, marshalError
	}

	unmarshalError := json.Unmarshal(modelJson, &ent)
	if unmarshalError != nil {
		return ent, unmarshalError
	}
	return ent, nil
}

func (model *Customer) Unmarshal() (entity.Customer, error) {
	var ent entity.Customer
	modelJson, marshalError := json.Marshal(model)
	if marshalError != nil {
		return ent, marshalError
	}

	unmarshalError := json.Unmarshal(modelJson, &ent)
	if unmarshalError != nil {
		return ent, unmarshalError
	}
	return ent, nil
}

func (model *Offering) Unmarshal() (entity.Offering, error) {
	var ent entity.Offering
	modelJson, marshalError := json.Marshal(model)
	if marshalError != nil {
		return ent, marshalError
	}

	unmarshalError := json.Unmarshal(modelJson, &ent)
	if unmarshalError != nil {
		return ent, unmarshalError
	}
	return ent, nil
}

func (model *CustomerContact) Marshal(ent entity.CustomerContact) error {
	entJson, err := json.Marshal(ent)
	if err != nil {
		return err
	}
	unErr := json.Unmarshal(entJson, &model)
	if unErr != nil {
		return unErr
	}
	return nil
}

func (model *Proposal) Marshal(ent entity.Proposal) error {
	entJson, err := json.Marshal(ent)
	if err != nil {
		return err
	}
	unErr := json.Unmarshal(entJson, &model)
	if unErr != nil {
		return unErr
	}
	return nil
}

func (model *Customer) Marshal(ent entity.Customer) error {
	entJson, err := json.Marshal(ent)
	if err != nil {
		return err
	}
	unErr := json.Unmarshal(entJson, &model)
	if unErr != nil {
		return unErr
	}
	return nil
}

func (model *Offering) Marshal(ent entity.Offering) error {
	entJson, err := json.Marshal(ent)
	if err != nil {
		return err
	}
	unErr := json.Unmarshal(entJson, &model)
	if unErr != nil {
		return unErr
	}
	return nil
}
