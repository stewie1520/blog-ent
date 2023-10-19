package user

import (
	"context"

	"github.com/stewie1520/blog_ent/ent"
	"github.com/stewie1520/blog_ent/ent/account"
	"github.com/stewie1520/blog_ent/ent/user"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/stewie1520/blog_ent/core"
	"github.com/stewie1520/blog_ent/tools/securities"
	"github.com/stewie1520/blog_ent/usecases"
)

var _ usecases.Command[*TokensResponse] = (*LoginCommand)(nil)

type LoginCommand struct {
	app           core.App
	userClient    *ent.UserClient
	accountClient *ent.AccountClient

	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginCommand(app core.App) *LoginCommand {
	return &LoginCommand{
		app:           app,
		userClient:    app.Dao().User,
		accountClient: app.Dao().Account,
	}
}

func (cmd *LoginCommand) Validate() error {
	return validation.ValidateStruct(cmd,
		validation.Field(&cmd.Email, validation.Required, is.Email, validation.Length(1, 255)),
		validation.Field(&cmd.Password, validation.Required, validation.Length(8, 255)),
	)
}

func (cmd *LoginCommand) Execute(ctx context.Context) (*TokensResponse, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	dbAccount, err := cmd.accountClient.Query().Where(account.Email(cmd.Email)).First(ctx)

	if ent.IsNotFound(err) {
		return nil, usecases.ErrInvalidCredentials
	}

	if err != nil {
		return nil, err
	}

	if ok := securities.CompareHashAndPassword(dbAccount.Password, cmd.Password); !ok {
		cmd.app.Log().Info("Invalid credentials")
		return nil, usecases.ErrInvalidCredentials
	}

	dbUser, err := cmd.userClient.Query().Where(user.HasAccountWith(account.ID(dbAccount.ID))).First(ctx)

	if err != nil {
		return nil, err
	}

	return createTokens(cmd.app.Config(), dbUser.ID.String(), dbAccount.ID.String())
}
