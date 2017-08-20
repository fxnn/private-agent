package it

import (
	"github.com/boltdb/bolt"
	"github.com/fxnn/deadbox/config"
	"github.com/fxnn/deadbox/daemon"
	"github.com/fxnn/deadbox/drop"
	"github.com/fxnn/deadbox/model"
	"github.com/fxnn/deadbox/worker"
	"net/url"
	"os"
	"testing"
	"time"
)

const workerDbFileName = "worker.boltdb"
const workerName = "itWorker"
const dropDbFileName = "drop.boltdb"
const dropName = "itDrop"
const port = "54123"

func assertWorkerTimeoutInFuture(actualWorker model.Worker, t *testing.T) {
	if actualWorker.Timeout.Before(time.Now()) {
		t.Fatalf("expected worker timeout to be in the future, but was %s", actualWorker.Timeout)
	}
}
func assertWorkerName(actualWorker model.Worker, workerName string, t *testing.T) {
	if string(actualWorker.Id) != workerName {
		t.Fatalf("expected worker to be %s, but was %v", workerName, actualWorker)
	}
}
func assertNumberOfWorkers(actualWorkers []model.Worker, expectedNumber int, t *testing.T) {
	if len(actualWorkers) != expectedNumber {
		t.Fatalf("expected %d workers, but got %v", expectedNumber, actualWorkers)
	}
}
func assertNumberOfRequests(actualRequests []model.WorkerRequest, expectedNumber int, t *testing.T) {
	if len(actualRequests) != expectedNumber {
		t.Fatalf("expected %d requests, but got %v", expectedNumber, actualRequests)
	}
}

func runDropDaemon(t *testing.T) drop.DaemonizedDrop {

	cfg := config.Drop{Name: dropName, ListenAddress: ":" + port}
	db, err := bolt.Open(dropDbFileName, 0664, bolt.DefaultOptions)
	if err != nil {
		t.Fatalf("could not open Drop's BoltDB: %s", err)
	}

	dropDaemon := drop.New(cfg, db)
	dropDaemon.OnStop(func() error {
		db.Close()
		os.Remove(dropDbFileName)
		return nil
	})
	dropDaemon.Start()

	return dropDaemon
}

func runWorkerDaemon(t *testing.T) daemon.Daemon {

	cfg := config.Worker{Name: workerName, DropUrl: parseUrlOrPanic("http://localhost:" + port)}
	db, err := bolt.Open(workerDbFileName, 0664, bolt.DefaultOptions)
	if err != nil {
		t.Fatalf("could not open Worker's BoltDB: %s", err)
	}

	workerDaemon := worker.New(cfg, db)
	workerDaemon.OnStop(func() error {
		db.Close()
		os.Remove(workerDbFileName)
		return nil
	})
	workerDaemon.Start()

	return workerDaemon
}

func stopDaemon(d daemon.Daemon, t *testing.T) {
	t.Log(d.Stop())
}

func parseUrlOrPanic(s string) *url.URL {
	result, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	return result
}