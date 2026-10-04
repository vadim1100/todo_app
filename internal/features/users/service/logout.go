package users_service

import (
	"context"
	"fmt"
	"time"
)

func (s *UsersService) Logout(ctx context.Context, token string) error {
	claims, err := s.jwtManager.Parse(token)
	if err != nil {
		return fmt.Errorf("parse token for logout: %w", err)
	}

	ttl := time.Until(claims.ExpiresAt.Time)
	if ttl <= 0 {
		return nil
	}

	if err := s.blacklist.Add(ctx, token, ttl); err != nil {
		return fmt.Errorf("add token to blacklist: %w", err)
	}
	
	return nil
}