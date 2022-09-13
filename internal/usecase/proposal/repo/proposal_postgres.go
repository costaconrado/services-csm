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
	return &ProposalRepo{pg}
}

func (r *ProposalRepo) GetProposal(ctx context.Context, ID uint) (entity.Proposal, error) {
	model := &Proposal{}
	err := r.DB.First(&model, ID).Error
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
