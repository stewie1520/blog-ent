package core

import (
	"github.com/stewie1520/blog_ent/config"
	"github.com/stewie1520/blog_ent/ent"
	hook "github.com/stewie1520/blog_ent/hooks"
	"go.uber.org/zap"
)

type App interface {
	Dao() *ent.Client
	Log() *zap.Logger

	Bootstrap() error
	Config() *config.Config
	IsDebug() bool

	// OnAfterAccountCreated hook is triggered after an account is created in identity service (SuperTokens for e.g)
	// This is useful when you want to create an user in your database after an account is created in identity service
	OnAfterAccountCreated() *hook.Hook[*AccountCreatedEvent]

	// OnUnauthorizedAccess Thrown when a protected backend API is accessed without a session.
	// The default behaviour of this is to clear session tokens (if any) and send a 401 to the frontend.
	OnUnauthorizedAccess() *hook.Hook[*UnauthorizedAccessEvent]
}
