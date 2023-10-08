package datastreams

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	util "github.com/samuelhem/go_data_streams/util"
)

func NewApplication(name string, uri string) (uuid.UUID, *Application) {
	id := uuid.New()
	return id, &Application{
		Id:       &UUID{Value: id.String()},
		Name:     name,
		Hostname: uri,
                Streams: make(map[string]*DataStream),
	}
}

type ApplicationPool struct {
	registeredApps map[uuid.UUID]*Application
        connectionPool map[string]DataStreamServiceClient
}


func NewApplicationPool(initalSize int32) *ApplicationPool {
	return &ApplicationPool{
		registeredApps: make(map[uuid.UUID]*Application),
                connectionPool: make(map[string]DataStreamServiceClient),
	}
}

func (pool *ApplicationPool) RegisterApplication(app *Application) error {
	if err := app.validate(); err != nil {
		return err
	}


        slog.Info(fmt.Sprintf("Registering application on: %s", app.Hostname))
        //init reverse grpc connection 
        pool.connectionPool[app.Name] = NewDataStreamServiceClient(util.InitializeGrpcConnection(app.Hostname))

        //reinit events
        app.Streams = make(map[string]*DataStream)
	pool.registeredApps[parseUUID(app.Id)] = app
	return nil
}

func (pool *ApplicationPool) GetApplication(applicationId uuid.UUID) *Application {
	return pool.registeredApps[applicationId]
}

func (pool *ApplicationPool) GetAllApplicationsForStreamEvent(stream string, event string) []*Application {
    var res []*Application
    for _, app := range pool.registeredApps {
        if app.Streams[stream] != nil && app.Streams[stream].Events[event] != nil && app.Streams[stream].Events[event].State == EventState_SUBSCRIBED {
            res = append(res, app)
        }
    }
    return res
}

func parseUUID(id *UUID) uuid.UUID {
	res, _ := uuid.Parse(id.Value)
        return res
}

func (pool *ApplicationPool) IsRegistered(applicationId uuid.UUID) bool {
	return pool.registeredApps[applicationId] != nil
}

func (a *Application) validate() error {
	if a.Name == "" {
		return errors.New("Application name is required")
	}
	return nil
}

func (a *Application) AddStream(name string) { 
        a.Streams[name] = &DataStream{
            Name: name,
            Events: make(map[string]*Event),
        }
}

func (a *Application) AddEvent(stream string, event *Event) {
        if a.Streams[stream] == nil {
            slog.Info(fmt.Sprintf("Stream %s not found creating new one", stream))
            a.AddStream(stream)
        }


        if a.Streams[stream].Events[event.Name] != nil {
            slog.Info(fmt.Sprintf("Event %s already exists in stream %s", event.Name, stream))
            return
        }

        event.State = EventState_SUBSCRIBED

        a.Streams[stream].Events[event.Name] = event
        slog.Info(fmt.Sprintf("Application %s -- Added event %s to stream %s", a.Name, event.Name, stream))
}

func (a *Application) RemoveEvent(stream string, event *Event) {

}



