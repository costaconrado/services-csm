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
		return entity.Proposal{}, fmt.Errorf("ProposalUseCase - Get: %w", err)
	}

	return proposal, nil
}

func (uc *ProposalUseCase) List(ctx context.Context, page int) ([]entity.Proposal, error) {
	proposal, err := uc.repo.ListProposals(ctx, page)
	if err != nil {
		return nil, fmt.Errorf("ProposalUseCase - List: %w", err)
	}

	return proposal, nil
}

func (uc *ProposalUseCase) Create(ctx context.Context, proposal entity.Proposal) (entity.Proposal, error) {
	proposalSync, err := uc.repo.CreateProposal(ctx, proposal)
	if err != nil {
		return proposal, fmt.Errorf("ProposalUseCase - Create: %w", err)
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
	proposalSync, err := uc.Get(ctx, proposal.ID)
	proposalSync.Stage = stage
	proposalSync, err = uc.Update(ctx, proposal)

	if err != nil {
		return proposal, fmt.Errorf("ProposalUseCase - ChangeStage: %w", err)
	}

	return proposalSync, nil
}
