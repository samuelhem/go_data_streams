package datastreams

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

func NewApplication(name string, uri string) (uuid.UUID, *Application) {
	id := uuid.New()
	return id, &Application{
		Id:       &UUID{Value: id.String()},
		Name:     name,
		Hostname: uri,
		Events:   make(map[string]EventState),
	}
}

type ApplicationPool struct {
	registeredApps map[uuid.UUID]*Application
}

func NewApplicationPool(initalSize int32) *ApplicationPool {
	return &ApplicationPool{
		registeredApps: make(map[uuid.UUID]*Application),
	}
}

func (pool *ApplicationPool) RegisterApplication(app *Application) error {
	if err := app.validate(); err != nil {
		return err
	}
        //reinit events
        app.Events = make(map[string]EventState)
	pool.registeredApps[parseUUID(app.Id)] = app
	return nil
}

func (pool *ApplicationPool) GetApplication(applicationId uuid.UUID) *Application {
	return pool.registeredApps[applicationId]
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

func (a *Application) AddEvent(event string) {
	a.Events[event] = EventState_SUBSCRIBED
	slog.Info(fmt.Sprintf("Application: %s -- Subscribing to the event %s", a.Name, event))
}

func (a *Application) RemoveEvent(event string) {
	delete(a.Events, event)
}
