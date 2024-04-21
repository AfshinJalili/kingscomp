package repository

import (
	"context"
	"github.com/redis/rueidis"
	"kingscomp/internal/entity"
)

var _ CommonBehaviour[entity.Entity] = &RedisCommonBehaviour[entity.Entity]{}

type CommonBehaviour[T entity.Entity] interface {
	Get(ctx context.Context, id entity.ID) (T, error)
	Save(ctx context.Context, ent entity.Entity) error
}

type RedisCommonBehaviour[T entity.Entity] struct {
	client rueidis.Client
}

func NewRedisCommonBehaviour[T entity.Entity]() *RedisCommonBehaviour[T] {

}

func (r RedisCommonBehaviour[T]) Get(ctx context.Context, id entity.ID) (T, error) {

}

func (r RedisCommonBehaviour[T]) Save(ctx context.Context, ent entity.Entity) error {
	//TODO implement me
	panic("implement me")
}
