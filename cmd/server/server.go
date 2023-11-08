package server

import (
	"example/komposervice/api"
	"example/komposervice/internal/config"
	"example/komposervice/internal/service"
	"example/komposervice/internal/tasks"
	"example/komposervice/pkg/lib/job"
	"example/komposervice/pkg/lib/worker"
	"fmt"
	"time"
)

// Server
// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func Server() error {
	server := api.New()
	return server.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
}

func AsyncWorker(concurrency int) error {
	w := worker.NewServer(concurrency, worker.Queue{
		config.CriticalQueue: 6, // processed 60% of the time
		config.DefaultQueue:  3, // processed 30% of the time
		config.LowQueue:      1, // processed 10% of the time
	})
	w.HandleFunctions(tasks.Path())
	return w.Run()
}

func JobLaunch() {
	j := job.New()
	j.Scheduler(service.Ping, 5*time.Second)
	if err := j.Launch(); err != nil {
		panic(err)
	}
}
