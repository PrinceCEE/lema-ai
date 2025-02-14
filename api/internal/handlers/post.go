package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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
	GetPost(postId uint) (*models.Post, error)
	GetPosts(userId uint) ([]*models.Post, error)
	DeletePost(postId uint) error
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
	UserID uint   `json:"userId" validate:"required,gt=0"`
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

	postId, err := strconv.Atoi(chi.URLParam(r, "post_id"))
	if err != nil {
		resp.Message = err.Error()
		response.SendErrorResponse(w, resp, http.StatusBadRequest)
		return
	}

	post, err := h.postService.GetPost(uint(postId))
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

	userId, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		resp.Message = "Invalid user ID"
		response.SendErrorResponse(w, resp, http.StatusBadRequest)
		return
	}

	posts, err := h.postService.GetPosts(uint(userId))
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

	postId, err := strconv.Atoi(chi.URLParam(r, "post_id"))
	if err != nil {
		resp.Message = err.Error()
		response.SendErrorResponse(w, resp, http.StatusBadRequest)
		return
	}

	err = h.postService.DeletePost(uint(postId))
	if err != nil {
		code := apperror.GetErrorStatusCode(err)
		resp.Message = err.Error()
		response.SendErrorResponse(w, resp, code)
		return
	}

	resp.Message = "Post deleted successfully"
	response.SendResponse(w, resp, nil)
}
