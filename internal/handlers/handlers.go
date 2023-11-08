package handlers

import (
	"job-portal/internal/auth"
	"job-portal/internal/middleware"
	"job-portal/internal/repository"
	"job-portal/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func API(a *auth.Auth, c repository.Repository) *gin.Engine {

	r := gin.New()

	m, err := middleware.NewMid(a)
	s := service.NewServiceStore(c)
	h := handler{
		a: a,
		s: s,
	}

	if err != nil {
		log.Panic().Msg("middlewares not set up")
	}

	r.Use(m.Log(), gin.Recovery())

	r.POST("/api/register", h.UserRegister)
	r.POST("/api/login", h.UserLogin)
	r.POST("/api/companies", m.Authenticate(h.CreateCompany))
	r.GET("/api/companies", m.Authenticate(h.ViewCompany))
	r.GET("/api/companies/:id", m.Authenticate(h.GetCompanyById))
	r.POST("/api/companies/:id/jobs", m.Authenticate(h.AddJob))
	r.GET("/api/jobs", m.Authenticate(h.ViewJobs))
	r.GET("/api/jobs/:id", m.Authenticate(h.ViewJobById))
	r.GET("/api/companies/:id/jobs", m.Authenticate(h.ViewJobByCompany))
	r.POST("/api/job/applications/:id", m.Authenticate(h.ApplyForJob))

	// Return the prepared Gin engine
	return r
}
