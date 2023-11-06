package handlers

import (
	"encoding/json"
	"job-portal/internal/middleware"
	"job-portal/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

func (h *handler) CreateCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var newCom models.NewCompany
	err := json.NewDecoder(c.Request.Body).Decode(&newCom)
	if err != nil {
		log.Error().Str("Trace Id", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	validate := validator.New()
	err = validate.Struct(newCom)

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": "Bad Request"})
		return
	}
	com, err := h.s.CreateCompany(ctx, newCom)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, com)

}
func (h *handler) ViewCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	data, err := h.s.ViewCompany(ctx)

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *handler) GetCompanyById(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	id := c.Param("id")
	cId, err := strconv.Atoi(id)
	if err != nil {
		// Handle invalid ID
		log.Error().Err(err).Str("Trace Id", traceId)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	// Call the service layer to get company information
	company, err := h.s.GetCompanyInfoByID(ctx, cId)
	if err != nil {
		// Handle errors, e.g., company not found
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the company data as JSON response
	c.JSON(http.StatusOK, company)
}
