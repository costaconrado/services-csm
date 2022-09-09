// Package usecase implements application business logic. Each logic group in own file.
package proposal

import (
	"context"

	"github.com/costaconrado/services-csm/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=proposal_test

type (
	Proposal interface {
		Create(context.Context, entity.Proposal) (entity.Proposal, error)
		Lock(context.Context, entity.Proposal) error
		History(context.Context) ([]entity.Proposal, error)
	}

	ProposalRepo interface {
		GetProposal(context.Context, uint) (entity.Proposal, error)
		CreateProposal(context.Context, entity.Proposal) (entity.Proposal, error)
		UpdateProposal(context.Context, entity.Proposal) (entity.Proposal, error)
		ChangeStageProposal(context.Context, entity.Proposal, entity.DealStage) (entity.Proposal, error)
	}

	ProposalWebAPI interface {
		Translate(entity.Proposal) (entity.Proposal, error)
	}
)
