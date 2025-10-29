package handlers

import (
	"dev-clash/internal/dto"
	"dev-clash/pkg/logger"
	"dev-clash/pkg/server_utils/configure_headers"
	custom_errors "dev-clash/pkg/server_utils/errors"
	"dev-clash/pkg/server_utils/response_message"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)
	var newUser dto.CreateUser

	json.NewDecoder(r.Body).Decode(&newUser)

	// CHECKING CORRECT JSON
	if newUser.Username == "" && newUser.Email == "" && newUser.Password == "" {
		wError := custom_errors.New(errors.New("empty body"), 400)
		wError.AddResponseData("Request body is empty")
		custom_errors.ErrorResponse(w, wError, logger.GetLoger())
		return
	}

	user, err := h.User.CreateUser(&newUser)

	if err != nil {
		custom_errors.ErrorResponse(w, err, logger.GetLoger())
	} else {
		safetyUser := dto.SafetyUser{Username: user.Username, Email: user.Email}
		logger.Info(fmt.Sprintf("User created succesfully: %+v", safetyUser))
		response_message.WrapperResponseJSON(w, 201, safetyUser)
	}
}

func (h *Handlers) GetUserByID(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	//GETTING USER ID FROM REQUEST
	userID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		wError := custom_errors.New(err, 400)
		wError.AddLogData(fmt.Sprintf("Invalid user ID: %v. ID must be decimal", mux.Vars(r)["id"]))
		wError.AddResponseData(fmt.Sprintf("Invalid user ID: %v. ID must be decimal", mux.Vars(r)["id"]))
		custom_errors.ErrorResponse(w, wError, logger.GetLoger())
		return
	}
	
	findedUser, err := h.User.FindUserByID(userID)
		if err != nil {
		custom_errors.ErrorResponse(w, err, logger.GetLoger())
	} else {
		
		safetyUser := dto.SafetyUser{
			Username: findedUser.Username, 
			Email: findedUser.Email, 
			Description: dto.NullStringToValid(findedUser.Description),
			Status: dto.NullStringToValid(findedUser.Status),
			ModeratorTimes: findedUser.ModeratorTimes,
			ParticipantTimes: findedUser.ParticipantTimes,
			PrizeTimes: findedUser.PrizeTimes,
			Skills: findedUser.Skills,
		}
		logger.Info(fmt.Sprintf("User finded succesfully: %+v", safetyUser))
		response_message.WrapperResponseJSON(w, 200, safetyUser)
	}
}