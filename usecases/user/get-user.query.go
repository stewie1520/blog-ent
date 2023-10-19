package user

import (
	"context"

	"github.com/stewie1520/blog_ent/ent"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/stewie1520/blog_ent/core"
	"github.com/stewie1520/blog_ent/usecases"
)

var _ usecases.Query[*ent.User] = (*GetUserQuery)(nil)

func NewGetUserQuery(app core.App) *GetUserQuery {
	return &GetUserQuery{
		app:        app,
		userClient: app.Dao().User,
	}
}

type GetUserQuery struct {
	app        core.App
	userClient *ent.UserClient

	ID string `json:"id"`
}

// Validate implements Query.
func (q *GetUserQuery) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(&q.ID, validation.Required, is.UUIDv4),
	)
}

// Execute implements Query.
func (q *GetUserQuery) Execute(ctx context.Context) (*ent.User, error) {
	if err := q.Validate(); err != nil {
		return nil, err
	}

	user, err := q.userClient.Get(ctx, uuid.MustParse(q.ID))
	return user, err
}
