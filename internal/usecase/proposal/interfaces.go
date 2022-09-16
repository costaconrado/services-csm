// Package usecase implements application business logic. Each logic group in own file.
package proposal

import (
	"context"

	"github.com/costaconrado/services-csm/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=proposal_test

type (
	Proposal interface {
		Get(context.Context, uint) (entity.Proposal, error)
		List(context.Context, int) ([]entity.Proposal, error)
		Create(context.Context, entity.Proposal) (entity.Proposal, error)
		Update(context.Context, entity.Proposal) (entity.Proposal, error)
		ChangeStage(context.Context, entity.Proposal, entity.DealStage) (entity.Proposal, error)
	}

	ProposalRepo interface {
		GetProposal(context.Context, uint) (entity.Proposal, error)
		ListProposals(context.Context, int) ([]entity.Proposal, error)
		CreateProposal(context.Context, entity.Proposal) (entity.Proposal, error)
		UpdateProposal(context.Context, entity.Proposal) (entity.Proposal, error)
	}
)
