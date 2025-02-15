package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/princecee/lema-ai/config"
	"github.com/princecee/lema-ai/internal/db/models"
	apperror "github.com/princecee/lema-ai/pkg/error"
	"github.com/princecee/lema-ai/pkg/json"
	"github.com/princecee/lema-ai/pkg/response"
	"github.com/princecee/lema-ai/pkg/validator"
	"github.com/rs/zerolog"
)

type PostService interface {
	CreatePost(p *models.Post) error
	GetPost(postId string) (*models.Post, error)
	GetPosts(userId string) ([]*models.Post, error)
	DeletePost(postId string) error
}

type PostHandler struct {
	postService PostService
	config      *config.Config
	logger      zerolog.Logger
}

func NewPostHandler(postService PostService, cfg *config.Config, l zerolog.Logger) *PostHandler {
	return &PostHandler{postService, cfg, l}
}

type createPostData struct {
	Title  string `json:"title" validate:"required,min=5"`
	Body   string `json:"body" validate:"required"`
	UserID string `json:"userId" validate:"required,uuid"`
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	resp := response.Response[any]{}
	data := new(createPostData)

	err := json.ReadJSON(r.Body, data)
	defer r.Body.Close()
	if err != nil {
		resp.Message = err.Error()
		response.SendErrorResponse(w, resp, http.StatusBadRequest)
		return
	}

	validationErrors := validator.ValidateData(data)
	if validationErrors != nil {
		resp.Message = apperror.ErrBadRequest.Error()
		resp.Data = validationErrors
		response.SendErrorResponse(w, resp, http.StatusBadRequest)
		return
	}

	post := &models.Post{
		ID:     uuid.NewString(),
		Title:  data.Title,
		Body:   data.Body,
		UserID: data.UserID,
	}

	err = h.postService.CreatePost(post)
	if err != nil {
		code := apperror.GetErrorStatusCode(err)
		resp.Message = err.Error()
		response.SendErrorResponse(w, resp, code)
		return
	}

	resp.Message = "Post created successfully"
	resp.Data = post
	response.SendResponse(w, resp, nil)
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	resp := response.Response[any]{}

	postId := chi.URLParam(r, "post_id")
	if !validator.IsValidUUID(postId) {
		resp.Message = "Invalid post ID"
		response.SendErrorResponse(w, resp, http.StatusBadRequest)
		return
	}

	post, err := h.postService.GetPost(postId)
	if err != nil {
		code := apperror.GetErrorStatusCode(err)
		resp.Message = err.Error()
		response.SendErrorResponse(w, resp, code)
		return
	}

	resp.Message = "Post fetched successfully"
	resp.Data = post
	response.SendResponse(w, resp, nil)
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	resp := response.Response[any]{}

	userId := r.URL.Query().Get("user_id")
	if !validator.IsValidUUID(userId) {
		resp.Message = "Invalid user ID"
		response.SendErrorResponse(w, resp, http.StatusBadRequest)
		return
	}

	posts, err := h.postService.GetPosts(userId)
	if err != nil {
		code := apperror.GetErrorStatusCode(err)
		resp.Message = err.Error()
		response.SendErrorResponse(w, resp, code)
		return
	}

	resp.Message = "Posts fetched successfully"
	resp.Data = posts
	response.SendResponse(w, resp, nil)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	resp := response.Response[any]{}

	postId := chi.URLParam(r, "post_id")
	if !validator.IsValidUUID(postId) {
		resp.Message = "Invalid post ID"
		response.SendErrorResponse(w, resp, http.StatusBadRequest)
		return
	}

	err := h.postService.DeletePost(postId)
	if err != nil {
		code := apperror.GetErrorStatusCode(err)
		resp.Message = err.Error()
		response.SendErrorResponse(w, resp, code)
		return
	}

	resp.Message = "Post deleted successfully"
	response.SendResponse(w, resp, nil)
}
