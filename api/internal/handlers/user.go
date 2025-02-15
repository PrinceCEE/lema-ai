package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/princecee/lema-ai/config"
	"github.com/princecee/lema-ai/internal/db/models"
	apperror "github.com/princecee/lema-ai/pkg/error"
	"github.com/princecee/lema-ai/pkg/pagination"
	"github.com/princecee/lema-ai/pkg/response"
	"github.com/princecee/lema-ai/pkg/validator"
	"github.com/rs/zerolog"
)

type UserService interface {
	GetUsers(page, limt int) (*pagination.GetUsersResult, error)
	GetUserCount() (int64, error)
	GetUser(id string) (*models.User, error)
}

type UserHandler struct {
	userService UserService
	config      *config.Config
	logger      zerolog.Logger
}

func NewUserHandler(userService UserService, cfg *config.Config, l zerolog.Logger) *UserHandler {
	return &UserHandler{userService, cfg, l}
}

type GetUsersQuery struct {
	Page  int `validate:"required"`
	Limit int `validate:"required"`
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	resp := response.Response[any]{}

	query := GetUsersQuery{}
	page, limit, err := pagination.FormatPaginationQuery(r.URL.Query().Get("page"), r.URL.Query().Get("limit"))
	if err != nil {
		resp.Message = err.Error()
		response.SendErrorResponse(w, resp, http.StatusBadRequest)
		return
	}

	query.Page = page
	query.Limit = limit

	validationErrors := validator.ValidateData(query)
	if validationErrors != nil {
		resp.Message = apperror.ErrBadRequest.Error()
		resp.Data = validationErrors
		response.SendErrorResponse(w, resp, http.StatusBadRequest)
		return
	}

	getUsersResp, err := h.userService.GetUsers(query.Page, query.Limit)
	if err != nil {
		code := apperror.GetErrorStatusCode(err)
		resp.Message = err.Error()
		response.SendErrorResponse(w, resp, code)
		return
	}

	resp.Message = "Users fetched successfully"
	resp.Data = getUsersResp
	response.SendResponse(w, resp, nil)
}

func (h *UserHandler) GetUsersCount(w http.ResponseWriter, r *http.Request) {
	resp := response.Response[any]{}

	count, err := h.userService.GetUserCount()
	if err != nil {
		code := apperror.GetErrorStatusCode(err)
		resp.Message = err.Error()
		response.SendErrorResponse(w, resp, code)
		return
	}

	resp.Message = "Users count fetched successfully"
	resp.Data = map[string]int64{"count": count}
	response.SendResponse(w, resp, nil)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	resp := response.Response[any]{}

	userId := chi.URLParam(r, "user_id")
	if !validator.IsValidUUID(userId) {
		resp.Message = "Invalid user ID"
		response.SendErrorResponse(w, resp, http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUser(userId)
	if err != nil {
		code := apperror.GetErrorStatusCode(err)
		resp.Message = err.Error()
		response.SendErrorResponse(w, resp, code)
		return
	}

	resp.Message = "User fetched successfully"
	resp.Data = user
	response.SendResponse(w, resp, nil)
}
