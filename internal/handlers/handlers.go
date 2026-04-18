// Package handlers contains the handling
package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/adrr-dev/blog-app/internal/domain"
)

type Service interface {
	NewUser(username, password string) error
	FetchUser(username, password string) (*domain.User, error)
	FetchUserByID(id uint) (*domain.User, error)

	FetchPosts(userID uint) ([]domain.Post, error)
	NewPost(userID uint, content string) error
	RandomPosts() ([]domain.Post, error)
	DeletePost(postID, userID uint) error
}

type MiddleWare interface {
	GetID(ctx context.Context) (uint, bool)
}

type Components interface {
	Login() templ.Component
	CreateAccount(notice string) templ.Component
	Dashboard(user *domain.User) templ.Component
	Home(user *domain.User) templ.Component
	Posts(posts []domain.Post) templ.Component
	Feed(posts []domain.Post) templ.Component
}

type Handling struct {
	service    Service
	middleware MiddleWare
	components Components
}

func NewHandling(service Service, middleware MiddleWare, components Components) *Handling {
	newHandling := &Handling{service: service, middleware: middleware, components: components}
	return newHandling
}

func (h Handling) Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := h.service.FetchUser(username, password)
	if err != nil {
		http.Redirect(w, r, "/createaccount", http.StatusSeeOther)
	}

	userID := user.ID
	strID := fmt.Sprintf("%d", userID)

	cookie := &http.Cookie{
		Name:     "user_id",
		Value:    strID,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (h Handling) LoginPage(w http.ResponseWriter, r *http.Request) {
	component := h.components.Login()
	err := component.Render(context.Background(), w)
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

func (h Handling) CreateAccount(w http.ResponseWriter, r *http.Request) {
	notice := r.FormValue("notice")
	if notice != "" {
		notice = "account already exists"
	}
	component := h.components.CreateAccount(notice)
	err := component.Render(context.Background(), w)
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

func (h Handling) NewAccount(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	_, err := h.service.FetchUser(username, password)
	if err == nil {
		http.Redirect(w, r, "/createaccount?notice=already_exists", http.StatusSeeOther)
		return
	}

	err = h.service.NewUser(username, password)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, "/loginpage", http.StatusSeeOther)
}

func (h Handling) DashboardPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, ok := h.middleware.GetID(ctx)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.service.FetchUserByID(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	component := h.components.Dashboard(user)
	err = component.Render(context.Background(), w)
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

func (h Handling) HomePage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, ok := h.middleware.GetID(ctx)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.service.FetchUserByID(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	component := h.components.Home(user)
	err = component.Render(context.Background(), w)
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

func (h Handling) NewPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, ok := h.middleware.GetID(ctx)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	content := r.FormValue("content")

	err := h.service.NewPost(id, content)
	if err != nil {
		http.Error(w, "would not create post", http.StatusInternalServerError)
	}

	posts, err := h.service.FetchPosts(id)
	if err != nil {
		http.Error(w, "could not fetch posts", http.StatusInternalServerError)
		return
	}
	component := h.components.Posts(posts)
	err = component.Render(context.Background(), w)
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

func (h Handling) FeedPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, ok := h.middleware.GetID(ctx)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	posts, err := h.service.FetchPosts(id)
	if err != nil {
		http.Error(w, "could not fetch posts", http.StatusInternalServerError)
		return
	}
	component := h.components.Feed(posts)
	err = component.Render(context.Background(), w)
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

func (h Handling) DeletePost(w http.ResponseWriter, r *http.Request) {
	strID := r.FormValue("id")
	u64, _ := strconv.ParseUint(strID, 10, 0)
	postID := uint(u64)

	ctx := r.Context()
	userID, _ := h.middleware.GetID(ctx)

	err := h.service.DeletePost(postID, userID)
	if err != nil {
		http.NotFound(w, r)
		return
	}
}
