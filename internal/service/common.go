package service

import (
	"example/komposervice/internal/config"
	"example/komposervice/internal/schema"
	"example/komposervice/internal/tasks"
	"example/komposervice/pkg/lib/worker"
)

func Ping() {
	// mailers.SendHTML("iduchungho@gmail.com")
	// log.Println("PONG")
}

func HealthCheck() schema.HealthCheckResponse {
	return schema.HealthCheckResponse{
		Status: "ok",
	}
}

func WorkerCheck() error {
	return worker.Exec(config.CriticalQueue, worker.NewTask(
		tasks.WorkerHealthCheck,
		1,
	))
}
