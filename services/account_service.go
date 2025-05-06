package services

import (
	"database/sql"
	"strconv"

	"github.com/Elagoht/go-pass/db"
	"github.com/Elagoht/go-pass/models"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type AccountService struct {
	db *sql.DB
}

func NewAccountService() *AccountService {
	return &AccountService{
		db: db.DB,
	}
}

/* Handles the creation of a new account */
func (service *AccountService) CreateAccount(account *models.Account) (*models.Account, error) {
	if err := validate.Struct(account); err != nil {
		return nil, err
	}

	result, err := service.db.Exec(
		"INSERT INTO accounts (platform, url, identity, passphrase, notes) VALUES (?, ?, ?, ?, ?)",
		account.Platform,
		account.URL,
		account.Identity,
		account.Passphrase,
		account.Notes,
	)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	account.Id = int(id)
	return account, nil
}

/* Handles the retrieval of all accounts */
func (service *AccountService) GetAllAccounts() ([]models.Account, error) {
	rows, err := service.db.Query("SELECT id, platform, url, identity, passphrase, notes FROM accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []models.Account{}
	for rows.Next() {
		var account models.Account
		if err := rows.Scan(
			&account.Id,
			&account.Platform,
			&account.URL,
			&account.Identity,
			&account.Passphrase,
			&account.Notes,
		); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

/* Handles the retrieval of a single account by ID */
func (service *AccountService) GetAccountByID(id string) (*models.Account, error) {
	accountID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	var account models.Account
	err = service.db.QueryRow(
		"SELECT id, platform, url, identity, passphrase, notes FROM accounts WHERE id = ?",
		accountID,
	).Scan(
		&account.Id,
		&account.Platform,
		&account.URL,
		&account.Identity,
		&account.Passphrase,
		&account.Notes,
	)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &account, nil
}

/* Handles the update of an account by ID */
func (service *AccountService) UpdateAccount(id string, account *models.Account) (*models.Account, error) {
	accountID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	if err := validate.Struct(account); err != nil {
		return nil, err
	}

	result, err := service.db.Exec(
		"UPDATE accounts SET platform = ?, url = ?, identity = ?, passphrase = ?, notes = ? WHERE id = ?",
		account.Platform, account.URL, account.Identity, account.Passphrase, account.Notes, accountID,
	)

	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	account.Id = accountID
	return account, nil
}

/* Handles the deletion of an account by ID */
func (service *AccountService) DeleteAccount(id string) error {
	accountID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	result, err := service.db.Exec("DELETE FROM accounts WHERE id = ?", accountID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
