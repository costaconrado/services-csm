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
		h.GET("/list", r.listProposals)
		h.POST("/", r.addProposal)
		h.GET("/:ID", r.getProposal)
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

// @Summary     Proposal
// @Description Get a list of proposals
// @ID          list-proposal
// @Tags  	    proposal
// @Produce     json
// @Success     200 {object} []entity.Proposal
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /proposal/list [get]
func (r *proposalRoutes) listProposals(c *gin.Context) {
	page, err := strconv.ParseInt(c.DefaultQuery("page", "0"), 10, 0)
	if err != nil {
		r.log.Error(err, "http - v1 - listProposal")
		errorResponse(c, http.StatusBadRequest, "invalid page")

		return
	}

	var proposals []entity.Proposal
	proposals, err = r.proposal.List(c, int(page))

	if err != nil {
		r.log.Error(err, "http - v1 - listProposal")
		errorResponse(c, http.StatusInternalServerError, "internal server error")

		return
	}

	c.JSON(http.StatusOK, proposals)
}

func (r *proposalRoutes) addProposal(c *gin.Context) {
	var proposal entity.Proposal
	err := c.BindJSON(&proposal)

	if err != nil {
		r.log.Error(err, "http - v1 - getProposal")
		errorResponse(c, http.StatusBadRequest, "invalid proposal")

		return
	}

	proposal, err = r.proposal.Create(c, proposal)

	if err != nil {
		r.log.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusInternalServerError, "internal server error")

		return
	}

	c.JSON(http.StatusOK, proposal)
}
