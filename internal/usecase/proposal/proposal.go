package proposal

import (
	"context"
	"fmt"

	"github.com/costaconrado/services-csm/internal/entity"
)

type ProposalUseCase struct {
	repo ProposalRepo
}

// New -.
func New(r ProposalRepo) *ProposalUseCase {
	return &ProposalUseCase{
		repo: r,
	}
}

func (uc *ProposalUseCase) Get(ctx context.Context, ID uint) (entity.Proposal, error) {
	proposal, err := uc.repo.GetProposal(ctx, ID)
	if err != nil {
		return entity.Proposal{}, fmt.Errorf("ProposalUseCase - GetProposal: %w", err)
	}

	return proposal, nil
}

func (uc *ProposalUseCase) Create(ctx context.Context, proposal entity.Proposal) (entity.Proposal, error) {
	proposalSync, err := uc.repo.CreateProposal(ctx, proposal)
	if err != nil {
		return proposal, fmt.Errorf("ProposalUseCase - CreateProposal: %w", err)
	}

	return proposalSync, nil
}

func (uc *ProposalUseCase) Update(ctx context.Context, proposal entity.Proposal) (entity.Proposal, error) {
	proposalSync, err := uc.repo.UpdateProposal(ctx, proposal)
	if err != nil {
		return proposal, fmt.Errorf("ProposalUseCase - UpdateProposal: %w", err)
	}

	return proposalSync, nil
}

func (uc *ProposalUseCase) ChangeStage(ctx context.Context, proposal entity.Proposal, stage entity.DealStage) (entity.Proposal, error) {
	proposalSync, err := uc.repo.ChangeStageProposal(ctx, proposal, stage)
	if err != nil {
		return proposal, fmt.Errorf("ProposalUseCase - ChangeStageProposal: %w", err)
	}

	return proposalSync, nil
}
