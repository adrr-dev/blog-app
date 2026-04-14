// Package middleware contains middleware
package middleware

import (
	"context"
	"net/http"
	"strconv"
)

type ctxKey string

const userKey ctxKey = "user_id"

type MiddleWare struct{}

func NewMiddleWare() *MiddleWare {
	newMiddleWare := &MiddleWare{}
	return newMiddleWare
}

func (m MiddleWare) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("user_id")
		if err != nil {
			http.Redirect(w, r, "/loginpage", http.StatusSeeOther)
			return
		}

		strID := cookie.Value
		u64, _ := strconv.ParseUint(strID, 10, 0)

		userID := uint(u64)

		ctx := context.WithValue(r.Context(), userKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m MiddleWare) GetID(ctx context.Context) (uint, bool) {
	id, ok := ctx.Value(userKey).(uint)
	return id, ok
}
