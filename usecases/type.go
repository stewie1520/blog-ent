package usecases

import "context"

type Command[T any] interface {
	Validate() error
	Execute(ctx context.Context) (T, error)
}

type Query[T any] interface {
	Validate() error
	Execute(ctx context.Context) (T, error)
}
