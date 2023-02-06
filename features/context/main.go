package myctx

// This package tells you how to use the context.Context API.

import (
	"context"
	"fmt"
)

type (
	contextKeyClientClaims struct{}
	contextKeyUser         struct{}
)

type ClientClaims struct {
	Sub   string
	Name  string
	Email string
}

type User struct {
	ID    string
	Name  string
	Email string
}

func Main() {
	ctx := context.Background()

	// Arrange data
	claims := ClientClaims{
		Sub:   "123",
		Name:  "John Doe",
		Email: "test@example.com",
	}
	user := User{
		ID:    "123",
		Name:  "John Doe",
		Email: "test@example.com",
	}

	// Set data
	ctx = setClientClaims(ctx, claims)
	ctx = setUser(ctx, user)

	// Get data
	claims, ok := getClientClaims(ctx)
	if !ok {
		panic("no client claims")
	}
	user, ok = getUser(ctx)
	if !ok {
		panic("no user")
	}

	// Show data
	fmt.Println("claims:", claims)
	fmt.Println("user:", user)
}

func setClientClaims(ctx context.Context, claims ClientClaims) context.Context {
	return context.WithValue(ctx, contextKeyClientClaims{}, claims)
}

func getClientClaims(ctx context.Context) (ClientClaims, bool) {
	claims, ok := ctx.Value(contextKeyClientClaims{}).(ClientClaims)
	return claims, ok
}

func setUser(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, contextKeyUser{}, user)
}

func getUser(ctx context.Context) (User, bool) {
	user, ok := ctx.Value(contextKeyUser{}).(User)
	return user, ok
}
