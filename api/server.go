package api

import (
	"example/komposervice/api/middleware"
	"example/komposervice/api/router"
	"example/komposervice/internal/config"
	"example/komposervice/internal/service"
	"example/komposervice/pkg/lib/job"
	"example/komposervice/pkg/lib/mailers"
	"example/komposervice/pkg/lib/worker"
	"example/komposervice/pkg/sentry"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	sentry.Init()
	mailers.Config(config.Email, config.EmailAppPassword)
	worker.SetBroker(config.RedisHost, config.RedisPort, config.RedisPassword)
}

type Server struct {
	engine *gin.Engine
}

func New() *Server {
	if config.StageStatus != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	return &Server{
		engine: gin.Default(),
	}
}

func (s *Server) middleware(mdws ...func(*gin.Engine)) {
	for _, m := range mdws {
		m(s.engine)
	}
}

func (s *Server) backgroundTask(tasks ...func()) {
	for _, t := range tasks {
		go t()
	}
}

func (s *Server) router(routers ...func(*gin.Engine)) {
	for _, r := range routers {
		r(s.engine)
	}
}

func (s *Server) Scheduler() {
	j := job.New()
	j.Scheduler(service.Ping, 5*time.Second)

	if err := j.Launch(); err != nil {
		panic(err)
	}
}

func (s *Server) Run(addr string) error {
	s.backgroundTask(
		s.Scheduler,
	)
	s.middleware(
		middleware.GinMiddleware,
		middleware.Sentry,
	)
	s.router(
		router.Common,
		router.Auth,
		router.Docs,
	)
	return s.engine.Run(addr)
}
