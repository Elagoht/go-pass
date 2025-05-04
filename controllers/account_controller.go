package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	errorHandler "github.com/Elagoht/go-pass/error"
	"github.com/Elagoht/go-pass/models"
	"github.com/Elagoht/go-pass/services"
	"github.com/Elagoht/go-pass/utils"
	"github.com/gorilla/mux"
)

// HTTP Controller for the Account model
type AccountController struct {
	service *services.AccountService
}

/* Instantiates a new AccountController */
func NewAccountController() *AccountController {
	return &AccountController{
		service: services.NewAccountService(),
	}
}

/* Handles the creation of a new account */
func (controller *AccountController) CreateAccount(writer http.ResponseWriter, request *http.Request) {
	var account models.Account
	if err := json.NewDecoder(request.Body).Decode(&account); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, "Invalid data")
		return
	}

	result, err := controller.service.CreateAccount(&account)
	if err != nil {
		errorHandler.HandleValidationError(writer, err)
		return
	}

	utils.RespondWithSuccess(writer, http.StatusCreated, result)
}

/* Handles the retrieval of all accounts */
func (controller *AccountController) GetAccounts(writer http.ResponseWriter, request *http.Request) {
	accounts, err := controller.service.GetAllAccounts()
	if err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, "Accounts retrieval failed")
		return
	}

	utils.RespondWithSuccess(writer, http.StatusOK, accounts)
}

/* Handles the retrieval of a single account by ID */
func (controller *AccountController) GetAccount(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	account, err := controller.service.GetAccountByID(vars["id"])
	if err == sql.ErrNoRows {
		utils.RespondWithError(writer, http.StatusNotFound, "Account not found")
		return
	} else if err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, "Account retrieval failed")
		return
	}

	utils.RespondWithSuccess(writer, http.StatusOK, account)
}

/* Handles the update of an account by ID */
func (controller *AccountController) UpdateAccount(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	var account models.Account
	if err := json.NewDecoder(request.Body).Decode(&account); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, "Invalid data")
		return
	}

	result, err := controller.service.UpdateAccount(vars["id"], &account)
	if err == sql.ErrNoRows {
		utils.RespondWithError(writer, http.StatusNotFound, "Account not found")
		return
	} else if err != nil {
		errorHandler.HandleValidationError(writer, err)
		return
	}

	utils.RespondWithSuccess(writer, http.StatusOK, result)
}

/* Handles the deletion of an account by ID */
func (controller *AccountController) DeleteAccount(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	err := controller.service.DeleteAccount(vars["id"])
	if err == sql.ErrNoRows {
		utils.RespondWithError(writer, http.StatusNotFound, "Account not found")
		return
	} else if err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, "Account deletion failed")
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
