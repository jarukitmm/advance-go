package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"advance-go/internal/config"
	"advance-go/internal/ping"
	"advance-go/internal/project"
	"advance-go/internal/score"
)

type Route struct {
	Name     string
	Path     string
	Method   string
	Endpoint gin.HandlerFunc
}

func Init(conf *config.Config) http.Handler {
	pj := project.NewHandler(conf)
	sc := score.NewHandler(conf)
	apiv1 := []Route{
		{
			Name:     "common ping",
			Method:   http.MethodGet,
			Path:     "/ping",
			Endpoint: ping.Endpoint,
		},
		{
			Name:     "show project code",
			Method:   http.MethodGet,
			Path:     "/project",
			Endpoint: pj.GetProjectCode,
		},
		{
			Name:     "score",
			Method:   http.MethodPost,
			Path:     "/score",
			Endpoint: sc.GetScore,
		},
	}
	ro := gin.New()

	v1 := ro.Group("/v1")
	for _, r := range apiv1 {
		v1.Handle(r.Method, r.Path, r.Endpoint)
	}
	return ro
}
