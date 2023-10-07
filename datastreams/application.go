package datastreams

import (
	"errors"

	"github.com/google/uuid"
)

func NewApplication(name string, uri string) *Application {
    return &Application{
        Name:     name,
        Hostname: uri,
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

func (pool *ApplicationPool) RegisterApplication(app *Application) (error) {

    if err := app.validate(); err != nil {
        return err
    }
    uuid := uuid.New()
    app.Id = &UUID{Value: uuid.String()}
    pool.registeredApps[uuid] = app
    return nil
}

func (a *Application) validate() error {
    if a.Name == "" {
        return errors.New("Application name is required")
    }
    return nil 
}
