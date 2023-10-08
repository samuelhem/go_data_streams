package datastreams

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type Worker struct {
	applicationPool *ApplicationPool
	workerId        uuid.UUID
	messageQueue    chan *Message
	quitChannel     chan bool
}

func NewWorker(applicationPool *ApplicationPool, messageQueue chan *Message) *Worker {
	return &Worker{
		applicationPool: applicationPool,
		workerId:        uuid.New(),
		messageQueue:    messageQueue,
		quitChannel:     make(chan bool),
	}
}

func (w *Worker) Start() {
	slog.Info("Starting", "worker", w.workerId)
	go func() {
		for {
			select {
			case message := <-w.messageQueue:
				go w.dispatch(message)
			case <-w.quitChannel:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.quitChannel <- true
	}()
}

func (w *Worker) dispatch(message *Message) {

	ctx, close := context.WithTimeout(context.Background(), 5*time.Second)
	defer close()

	apps := w.applicationPool.GetAllApplicationsForStreamEvent(message.Stream, message.Event.Name)
	for _, app := range apps {
		slog.Info(fmt.Sprintf("Dispatching message %+v", message))
		_, err := w.applicationPool.connectionPool[app.Name].Exchange(ctx, message)
		if err != nil {
			slog.Error(fmt.Sprintf("Error sending message to application %s", app.Name), "error", err)
		}
	}
}
