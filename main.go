package main

import (
	"log"
	"worker_pool/work"
	"worker_pool/worker"
)

const WORKER_COUNT = 50_000
const JOB_COUNT = 500_000

func main() {
	log.Println("starting application...")
	collector := worker.StartDispatcher(WORKER_COUNT) // start up worker pool

	for i, job := range work.CreateJobs(JOB_COUNT) {
		collector.Work <-worker.Work{Job: job, ID: i}
	}
}	