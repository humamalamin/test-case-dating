package auth

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/humamalamin/test-case-dating/api/domains/entities/auth"
	"github.com/humamalamin/test-case-dating/api/domains/interfaces"
	"github.com/humamalamin/test-case-dating/api/services"
	"github.com/humamalamin/test-case-dating/pkg/manager"
	"github.com/rs/zerolog/log"

	jsonHelper "github.com/humamalamin/test-case-dating/helpers/json"
)

var validate = validator.New()
var (
	code string
)

type Auth struct {
	Service interfaces.AuthService
}

// Login implements interfaces.AuthHandler.
func (h *Auth) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			code = "[Handler] Login - 1"
			log.Error().Err(err).Msg(code)
			jsonHelper.ErrorResponse(w, "400", http.StatusBadRequest, nil, "Failed read request body")
			return
		}

		if validateErr := validate.Struct(&req); validateErr != nil {
			code = "[Handler] Login - 2"
			log.Error().Err(validateErr).Msg(code)
			jsonHelper.ErrorResponse(w, "442", http.StatusBadRequest, nil, validateErr.Error())
			return
		}

		reqEntity := auth.Auth{
			Email:    req.Email,
			Password: req.Password,
		}

		result, err := h.Service.Login(r.Context(), &reqEntity)
		if err != nil {
			code = "[Handler] Login - 3"
			log.Error().Err(err).Msg(code)
			jsonHelper.ErrorResponse(w, "500", http.StatusInternalServerError, nil, err.Error())
			return
		}

		jsonHelper.SuccessResponse(w, "200", http.StatusOK, result, "success", nil)
	})
}

// Register implements interfaces.AuthHandler.
func (h *Auth) Register() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			code = "[Handler] Register - 1"
			log.Error().Err(err).Msg(code)
			jsonHelper.ErrorResponse(w, "400", http.StatusBadRequest, nil, "Failed read request body")
			return
		}

		if validateErr := validate.Struct(&req); validateErr != nil {
			code = "[Handler] Register - 2"
			log.Error().Err(validateErr).Msg(code)
			jsonHelper.ErrorResponse(w, "442", http.StatusBadRequest, nil, validateErr.Error())
			return
		}

		gender := 1
		if strings.ToLower(req.Gender) == "wanita" {
			gender = 0
		}

		reqEntity := &auth.Auth{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Gender:    gender,
			Email:     req.Email,
			Password:  req.Password,
		}

		err := h.Service.Register(r.Context(), reqEntity)
		if err != nil {
			code = "[Handler] Register - 3"
			log.Error().Err(err).Msg(code)
			jsonHelper.ErrorResponse(w, "500", http.StatusInternalServerError, nil, err.Error())
			return
		}

		jsonHelper.SuccessResponse(w, "201", http.StatusCreated, nil, "register success", nil)
	})
}

func NewAuthHandler(mgr manager.Manager) interfaces.AuthHandler {
	handler := new(Auth)
	handler.Service = services.NewAuthService(mgr)

	return handler
}
