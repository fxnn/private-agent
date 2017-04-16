package json

import (
	"github.com/fxnn/deadbox/model"
	"time"
)

type WorkerRequest struct {
	id      string
	timeout time.Time
}

func (r *WorkerRequest) Id() model.WorkerRequestId {
	return model.WorkerRequestId(r.id)
}
func (r *WorkerRequest) Timeout() time.Time {
	return r.timeout
}

func AsWorkerRequest(request model.WorkerRequest) WorkerRequest {
	return WorkerRequest{id: string(request.Id()), timeout: request.Timeout()}
}
func AsWorkerRequests(requests []model.WorkerRequest) []WorkerRequest {
	result := make([]WorkerRequest, len(requests))
	for i, request := range requests {
		result[i] = AsWorkerRequest(request)
	}
	return result
}