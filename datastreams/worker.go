package datastreams

import (
	"fmt"
	"log/slog"

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
	slog.Info("Starting", "worker",  w.workerId)
	go func() {
		for {
			select {
			case message := <-w.messageQueue:
				dispatch(message)
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

func (w *Worker) processMessage(message *Message) {

}

type ForwardMessagingDispatcher struct{}

func NewForwardMessagingDispatcher() *ForwardMessagingDispatcher {
	return &ForwardMessagingDispatcher{}
}

func dispatch(message *Message) {
    slog.Info(fmt.Sprintf("Dispatching message %+v", message))
}
