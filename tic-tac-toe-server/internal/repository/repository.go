package repository

import "context"

type Entity interface {
	GetId() uint64
}

type PostgresRepository[T Entity] interface {
	FindById(ctx context.Context, id uint64) (*T, error)
	FindAll(ctx context.Context) ([]*T, error)
	DeleteById(ctx context.Context, id uint64) error
	UpdateById(ctx context.Context, entity *T) error
}
