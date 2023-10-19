package user

import (
	"context"

	"github.com/stewie1520/blog_ent/ent"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/stewie1520/blog_ent/config"
	"github.com/stewie1520/blog_ent/core"
	"github.com/stewie1520/blog_ent/tools/securities"
	"github.com/stewie1520/blog_ent/usecases"
	"go.uber.org/zap"
)

var _ usecases.Command[*TokensResponse] = (*RegisterCommand)(nil)

type RegisterCommand struct {
	app           core.App
	userClient    *ent.UserClient
	accountClient *ent.AccountClient

	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Bio      string `json:"bio"`
}

func NewRegisterCommand(app core.App) *RegisterCommand {
	return &RegisterCommand{
		app:           app,
		userClient:    app.Dao().User,
		accountClient: app.Dao().Account,
	}
}

func (cmd *RegisterCommand) Validate() error {
	return validation.ValidateStruct(cmd,
		validation.Field(&cmd.FullName, validation.Required, validation.Length(1, 255)),
		validation.Field(&cmd.Password, validation.Required, validation.Length(8, 255)),
		validation.Field(&cmd.Email, validation.Required, is.Email, validation.Length(1, 255)),
	)
}

func (cmd *RegisterCommand) Execute(ctx context.Context) (*TokensResponse, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	if password, err := securities.HashPassword(cmd.Password); err != nil {
		return nil, err
	} else {
		cmd.Password = password
	}

	tx, err := cmd.app.Dao().BeginTx(ctx, nil)
	if err != nil {
		return rollback(tx, cmd.app, err)
	}

	accountClientTx := tx.Account
	userClientTx := tx.User

	dbAccountBuilder := accountClientTx.Create()
	dbAccountBuilder.SetID(uuid.New())
	dbAccountBuilder.SetEmail(cmd.Email)
	dbAccountBuilder.SetPassword(cmd.Password)
	dbAccount, err := dbAccountBuilder.Save(ctx)

	if err != nil {
		return rollback(tx, cmd.app, err)
	}

	dbUserBuilder := userClientTx.Create()
	dbUserBuilder.SetID(uuid.New())
	dbUserBuilder.SetFullName(cmd.FullName)
	dbUserBuilder.SetAccountID(dbAccount.ID)
	dbUserBuilder.SetBio(cmd.Bio)
	dbUser, err := dbUserBuilder.Save(ctx)

	if err != nil {
		return rollback(tx, cmd.app, err)
	}

	err = tx.Commit()
	if err != nil {
		return rollback(tx, cmd.app, err)
	}

	return createTokens(cmd.app.Config(), dbUser.ID.String(), dbAccount.ID.String())
}

func createTokens(config *config.Config, userId string, accountId string) (*TokensResponse, error) {
	accessToken, err := securities.NewPaseto(map[string]string{
		"userId":    userId,
		"accountId": accountId,
		"type":      "access",
	},
		config.Token.PrivateKey,
		config.Token.AccessTokenTTL,
	)

	if err != nil {
		return nil, err
	}

	refreshToken, err := securities.NewPaseto(map[string]string{
		"userId":    userId,
		"accountId": accountId,
		"type":      "refresh",
	},
		config.Token.PrivateKey,
		config.Token.RefreshTokenTTL,
	)

	if err != nil {
		return nil, err
	}

	return &TokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func rollback(tx *ent.Tx, app core.App, err error) (*TokensResponse, error) {
	if rerr := tx.Rollback(); rerr != nil {
		err = rerr
		app.Log().Info("Rollback error:", zap.Error(rerr))
	}

	return nil, err
}
