package repo

import (
	"context"
	"fmt"

	"github.com/costaconrado/services-csm/internal/entity"
	"github.com/costaconrado/services-csm/pkg/postgres"
)

const _defaultEntityCap = 64

type ProposalRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *ProposalRepo {
	pg.DB.AutoMigrate(&Proposal{})
	pg.DB.AutoMigrate(&Customer{})
	pg.DB.AutoMigrate(&CustomerContact{})
	pg.DB.AutoMigrate(&Offering{})

	return &ProposalRepo{pg}
}

func (r *ProposalRepo) GetProposal(ctx context.Context, ID uint) (entity.Proposal, error) {
	model := &Proposal{}
	err := r.DB.Preload("CustomerContact.Customer").Preload("Offerings.Addons").First(&model, ID).Error
	var proposal entity.Proposal

	if err != nil {
		return proposal, fmt.Errorf("ProposalRepo - GetProposal: %w", err)
	}

	proposal, err = model.Unmarshal()
	if err != nil {
		return proposal, fmt.Errorf("ProposalRepo - GetProposal: %w", err)
	}

	return proposal, nil
}

func (r *ProposalRepo) ListProposals(ctx context.Context, page int) ([]entity.Proposal, error) {
	var proposalModels []Proposal
	var proposalEntities []entity.Proposal
	err := r.DB.
		Limit(10).
		Offset(page * 10).
		Preload("CustomerContact.Customer").
		Find(&proposalModels).Error

	if err != nil {
		return nil, fmt.Errorf("ProposalRepo - GetProposal: %w", err)
	}

	proposalEntities, err = UnmarshalProposalList(proposalModels)
	if err != nil {
		return nil, fmt.Errorf("ProposalRepo - GetProposal: %w", err)
	}

	return proposalEntities, nil
}

func (r *ProposalRepo) CreateProposal(ctx context.Context, proposal entity.Proposal) (entity.Proposal, error) {
	model := &Proposal{}
	model.Marshal(proposal)

	err := r.DB.Create(model).Error
	if err != nil {
		return proposal, fmt.Errorf("ProposalRepo - CreateProposal: %w", err)
	}

	proposal, err = model.Unmarshal()
	if err != nil {
		return proposal, fmt.Errorf("ProposalRepo - CreateProposal: %w", err)
	}
	return proposal, nil
}

func (r *ProposalRepo) UpdateProposal(ctx context.Context, proposal entity.Proposal) (entity.Proposal, error) {
	model := &Proposal{}
	model.Marshal(proposal)

	err := r.DB.Save(proposal).Error
	if err != nil {
		return proposal, fmt.Errorf("ProposalRepo - UpdateProposal: %w", err)
	}

	proposal, err = model.Unmarshal()
	if err != nil {
		return proposal, fmt.Errorf("ProposalRepo - UpdateProposal: %w", err)
	}
	return proposal, nil
}
