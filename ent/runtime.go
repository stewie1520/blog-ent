// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/stewie1520/blog_ent/ent/account"
	"github.com/stewie1520/blog_ent/ent/schema"
	"github.com/stewie1520/blog_ent/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescEmail is the schema descriptor for email field.
	accountDescEmail := accountFields[1].Descriptor()
	// account.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	account.EmailValidator = accountDescEmail.Validators[0].(func(string) error)
	// accountDescPassword is the schema descriptor for password field.
	accountDescPassword := accountFields[2].Descriptor()
	// account.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	account.PasswordValidator = accountDescPassword.Validators[0].(func(string) error)
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountFields[3].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountFields[4].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() time.Time)
	// account.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	account.UpdateDefaultUpdatedAt = accountDescUpdatedAt.UpdateDefault.(func() time.Time)
	// accountDescID is the schema descriptor for id field.
	accountDescID := accountFields[0].Descriptor()
	// account.DefaultID holds the default value on creation for the id field.
	account.DefaultID = accountDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescFullName is the schema descriptor for full_name field.
	userDescFullName := userFields[1].Descriptor()
	// user.FullNameValidator is a validator for the "full_name" field. It is called by the builders before save.
	user.FullNameValidator = userDescFullName.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
