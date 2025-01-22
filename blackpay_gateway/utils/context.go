package utils

import (
	"context"
)

type contextKey string

const userIDKey contextKey = "userID"

// AddToContext adds a key-value pair to the context.
func AddToContext(ctx context.Context, key, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

// GetFromContext retrieves a value from the context by key.
func GetFromContext(ctx context.Context, key interface{}) interface{} {
	return ctx.Value(key)
}
