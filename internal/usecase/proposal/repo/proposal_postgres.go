package repo

import (
	"github.com/costaconrado/services-csm/pkg/postgres"
)

const _defaultEntityCap = 64

type ProposalRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *ProposalRepo {
	return &ProposalRepo{pg}
}

// func (r *ProposalRepo) GetHistory(ctx context.Context) ([]entity.Proposal, error) {

// }
