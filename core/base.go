package core

import (
	"github.com/stewie1520/blog_ent/config"
	"github.com/stewie1520/blog_ent/db"
	"github.com/stewie1520/blog_ent/ent"
	hook "github.com/stewie1520/blog_ent/hooks"
	"github.com/stewie1520/blog_ent/log"
	"go.uber.org/zap"
)

var _ App = (*BaseApp)(nil)

type BaseAppConfig struct {
	*config.Config
	IsDebug bool
}

type BaseApp struct {
	config BaseAppConfig
	dao    *ent.Client
	logger *zap.Logger

	onAfterAccountCreated *hook.Hook[*AccountCreatedEvent]
	onUnauthorizedAccess  *hook.Hook[*UnauthorizedAccessEvent]
}

func NewBaseApp(config BaseAppConfig) *BaseApp {
	app := &BaseApp{
		config:                config,
		onAfterAccountCreated: &hook.Hook[*AccountCreatedEvent]{},
		onUnauthorizedAccess:  &hook.Hook[*UnauthorizedAccessEvent]{},
	}

	return app
}

func (app *BaseApp) Bootstrap() error {
	if logger, err := log.New(); err != nil {
		return err
	} else {
		app.logger = logger
	}

	if err := app.initDatabase(); err != nil {
		return err
	}

	return nil
}

func (app *BaseApp) IsDebug() bool {
	return app.config.IsDebug
}

func (app *BaseApp) Dao() *ent.Client {
	return app.dao
}

func (app *BaseApp) Log() *zap.Logger {
	return app.logger
}

func (app *BaseApp) Config() *config.Config {
	return app.config.Config
}

func (app *BaseApp) OnAfterAccountCreated() *hook.Hook[*AccountCreatedEvent] {
	return app.onAfterAccountCreated
}

func (app *BaseApp) OnUnauthorizedAccess() *hook.Hook[*UnauthorizedAccessEvent] {
	return app.onUnauthorizedAccess
}

func (app *BaseApp) initDatabase() error {

	entClient, err := db.NewPostgresDBX(app.config.DatabaseURL)

	if err != nil {
		return err
	}

	app.dao = entClient

	return nil
}
