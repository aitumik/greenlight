package main

import (
	"context"
	"greenlight/internal/data"
	"net/http"
)

type contextKey string

const userContext = contextKey("user")

func (app *application) contextSetUser(r *http.Request, u *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContext, u)
	return r.WithContext(ctx)
}

func (app *application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContext).(*data.User)
	if !ok {
		panic("missing value in request context")
	}
	return user
}
