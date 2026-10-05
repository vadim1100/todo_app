package core_redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Blacklist struct {
	client *redis.Client
}

func NewBlacklist(client *redis.Client) *Blacklist {
	return &Blacklist{
		client: client,
	}
}

func (b *Blacklist) Add(ctx context.Context, token string, ttl time.Duration) error {
	key := NewBlacklistKey(token)
	if err := b.client.Set(ctx, key, "1", ttl).Err(); err != nil {
		return fmt.Errorf("add to blacklist: %w", err)
	}
	return nil
}

func (b *Blacklist) Contains(ctx context.Context, token string) (bool, error) {
	key := NewBlacklistKey(token)
	n, err := b.client.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("check blacklist: %w", err)
	}
	return n > 0, nil
}