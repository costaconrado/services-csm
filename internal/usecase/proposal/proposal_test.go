package proposal_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/costaconrado/services-csm/internal/entity"
	usecase "github.com/costaconrado/services-csm/internal/usecase/proposal"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var errNotFound = errors.New("proposal not found")
var errInternalServErr = errors.New("internal server error")
var emptyProposal = entity.Proposal{}
var validPersistentProposal = entity.Proposal{
	ID:        10,
	CreatedAt: time.Now(),
	Stage:     entity.STAGE_ENGAGE,
	CustomerContact: entity.CustomerContact{
		ID:   1,
		Name: "Joseph",
		Customer: entity.Customer{
			ID:       1,
			Name:     "Ultra Banking Inc.",
			Vertical: "Enterprise",
			Segment:  "Banking and Securing",
		},
	},
	Offerings: []entity.Offering{
		{
			ID:        1,
			SKU:       "KUB101",
			UnitValue: 1000,
			Quantity:  5,
			Name:      "Kubernetes Administration Training",
		},
	},
}

var validPersistentWonProposal = entity.Proposal{
	ID:        10,
	CreatedAt: time.Now(),
	Stage:     entity.STAGE_CLOSED_WON,
	CustomerContact: entity.CustomerContact{
		ID:   1,
		Name: "Joseph",
		Customer: entity.Customer{
			ID:       1,
			Name:     "Ultra Banking Inc.",
			Vertical: "Enterprise",
			Segment:  "Banking and Securing",
		},
	},
	Offerings: []entity.Offering{
		{
			ID:        1,
			SKU:       "KUB101",
			UnitValue: 1000,
			Quantity:  5,
			Name:      "Kubernetes Administration Training",
		},
	},
}

var validNonPersistentProposal = entity.Proposal{
	ID:        0,
	CreatedAt: time.Now(),
	Stage:     entity.STAGE_ENGAGE,
	CustomerContact: entity.CustomerContact{
		ID:   0,
		Name: "Joseph",
		Customer: entity.Customer{
			ID:       0,
			Name:     "Ultra Banking Inc.",
			Vertical: "Enterprise",
			Segment:  "Banking and Securing",
		},
	},
	Offerings: []entity.Offering{
		{
			ID:        0,
			SKU:       "KUB101",
			UnitValue: 1000,
			Quantity:  5,
			Name:      "Kubernetes Administration Training",
		},
	},
}

type test struct {
	name  string
	mock  func()
	input interface{}
	res   interface{}
	err   error
}

func proposal(t *testing.T) (*usecase.ProposalUseCase, *MockProposalRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockProposalRepo(mockCtl)

	proposal := usecase.New(repo)

	return proposal, repo
}

func TestGet(t *testing.T) {
	t.Parallel()
	proposal, repo := proposal(t)

	tests := []test{
		{
			name: "Wrong ID",
			mock: func() {
				repo.EXPECT().GetProposal(context.Background(), 0).Return(emptyProposal, errNotFound)
			},
			res:   emptyProposal,
			input: 0,
			err:   errNotFound,
		},
		{
			name: "Valid ID",
			mock: func() {
				repo.EXPECT().GetProposal(context.Background(), 10).Return(validPersistentProposal, nil)
			},
			res:   validPersistentProposal,
			input: 10,
			err:   nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			input, _ := tc.res.(uint)
			res, err := proposal.Get(context.Background(), input)

			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestCreate(t *testing.T) {
	t.Parallel()
	proposal, repo := proposal(t)

	tests := []test{
		{
			name: "Valid ID",
			mock: func() {
				repo.EXPECT().CreateProposal(context.Background(), validNonPersistentProposal).Return(validPersistentProposal, nil)
			},
			input: validNonPersistentProposal,
			res:   validPersistentProposal,
			err:   nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			input, _ := tc.input.(entity.Proposal)
			res, err := proposal.Create(context.Background(), input)

			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()
	proposal, repo := proposal(t)

	type Input struct {
		proposal entity.Proposal
	}

	tests := []test{
		{
			name: "Valid Change",
			mock: func() {
				repo.EXPECT().UpdateProposal(context.Background(), validPersistentWonProposal).Return(validPersistentWonProposal, nil)
			},
			input: Input{proposal: validPersistentWonProposal},
			res:   validPersistentWonProposal,
			err:   nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			input, _ := tc.input.(Input)
			res, err := proposal.Update(context.Background(), input.proposal)

			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestChangeStage(t *testing.T) {
	t.Parallel()
	proposal, repo := proposal(t)

	type Input struct {
		proposal entity.Proposal
		stage    entity.DealStage
	}

	tests := []test{
		{
			name: "Engage to Won",
			mock: func() {
				repo.EXPECT().GetProposal(context.Background(), 10).Return(validPersistentProposal, nil)
				repo.EXPECT().UpdateProposal(context.Background(), validPersistentWonProposal).Return(validPersistentWonProposal, nil)
			},
			input: Input{proposal: validPersistentProposal, stage: entity.STAGE_CLOSED_WON},
			res:   validPersistentWonProposal,
			err:   nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			input, _ := tc.input.(Input)
			res, err := proposal.ChangeStage(context.Background(), input.proposal, input.stage)

			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestList(t *testing.T) {
	t.Parallel()

	proposal, repo := proposal(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				repo.EXPECT().ListProposals(context.Background(), 0).Return(nil, nil)
			},
			res: []entity.Translation(nil),
			err: nil,
		},
		{
			name: "result with error",
			mock: func() {
				repo.EXPECT().ListProposals(context.Background(), 0).Return(nil, errInternalServErr)
			},
			res: []entity.Translation(nil),
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := proposal.List(context.Background(), 0)

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
