package handlers

import (
	"dev-clash/internal/dto"
	"dev-clash/pkg/logger"
	"dev-clash/pkg/server_utils/configure_headers"
	custom_errors "dev-clash/pkg/server_utils/errors"
	"dev-clash/pkg/server_utils/response_message"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)
	var newUser dto.CreateUser

	json.NewDecoder(r.Body).Decode(&newUser)

	user, err := h.User.CreateUser(&newUser)

	if err != nil {
		custom_errors.ErrorResponse(w, err, logger.GetLoger())
	} else {
		safetyUser := dto.SafetyUser{Username: user.Username, Email: user.Email}
		logger.Info(fmt.Sprintf("User created succesfully: %+v", safetyUser))
		response_message.WrapperResponseJSON(w, 201, safetyUser)
	}
}
