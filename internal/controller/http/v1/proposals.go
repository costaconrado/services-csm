package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/costaconrado/services-csm/internal/entity"
	usecase "github.com/costaconrado/services-csm/internal/usecase/proposal"
	"github.com/costaconrado/services-csm/pkg/logger"
)

type proposalRoutes struct {
	proposal usecase.Proposal
	log      logger.Interface
}

func newProposalRoutes(handler *gin.RouterGroup, t usecase.Proposal, l logger.Interface) {
	r := &proposalRoutes{t, l}

	h := handler.Group("/proposal")
	{
		h.GET("/get/:ID", r.getProposal)
	}
}

func (r *proposalRoutes) getProposal(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("ID"), 10, 0)
	if err != nil {
		r.log.Error(err, "http - v1 - getProposal")
		errorResponse(c, http.StatusBadRequest, "invalid ID")

		return
	}

	var proposal entity.Proposal
	proposal, err = r.proposal.Get(c, uint(ID))

	if err != nil {
		r.log.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusNotFound, "proposal not found")

		return
	}

	c.JSON(http.StatusOK, proposal)
}
