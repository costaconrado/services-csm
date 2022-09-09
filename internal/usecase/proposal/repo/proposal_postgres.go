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

func (r *ProposalRepo) ChangeStageProposal(ctx context.Context, proposal entity.Proposal, stage entity.DealStage) (entity.Proposal, error) {
	if proposal.ID == 0 {
		return proposal, fmt.Errorf("ProposalRepo - ChangeStageProposal: Proposal's ID is nil")
	}

	model := &Proposal{}
	model.Marshal(proposal)
	err := r.DB.Model(&model).Update("stage", stage).Error
	if err != nil {
		return proposal, fmt.Errorf("ProposalRepo - ChangeStageProposal: %w", err)
	}

	proposal, err = model.Unmarshal()
	if err != nil {
		return proposal, fmt.Errorf("ProposalRepo - ChangeStageProposal: %w", err)
	}
	return proposal, nil
}
