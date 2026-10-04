package core_service_blacklist

import (
	"context"
	"time"
)

type Blacklist interface {
	Add(ctx context.Context, token string, ttl time.Duration) error
	Contains(ctx context.Context, token string) (bool, error)
}