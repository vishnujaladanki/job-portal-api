package handlers

import (
	"encoding/json"
	"job-portal/internal/middleware"
	"job-portal/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

func (h *handler) ForgotPassword(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Define a new struct for login data
	var login struct {
		Email string `json:"email" validate:"required,email"`
	}

	// Attempt to decode JSON from the request body into the login variable
	err := json.NewDecoder(c.Request.Body).Decode(&login)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	// Create a new validator and validate the login variable
	validate := validator.New()
	err = validate.Struct(login)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	// Attempt to authenticate the user with the email
	_, err = h.s.CheckEmail(login.Email)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "given email is not registered with job portal"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "otp sucessfully sent to registred email"})

}

func (h *handler) ResetPassword(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var login models.Reset
	// Attempt to decode JSON from the request body into the login variable
	err := json.NewDecoder(c.Request.Body).Decode(&login)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	// Create a new validator and validate the login variable
	validate := validator.New()
	err = validate.Struct(login)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	if login.NewPassword != login.ConfirmPassword {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "newpassword and confirmpassword are not equal"})
		return
	}
	// Attempt to authenticate the user with the email
	_, err = h.s.UpdatePassword(login)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error while resetting the newpassword"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "sucessfully reset the password"})

}
