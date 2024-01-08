package v1

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sqlc_test/constants"
	"github.com/sqlc_test/models"
	"github.com/sqlc_test/utils"
	"go.uber.org/zap"
)

type AuthorsControllers struct {
	model *models.Queries
}

func NewAuthorsController(db *sql.DB, logger *zap.Logger) (*AuthorsControllers, error) {
	authorModel := models.New(db)
	return &AuthorsControllers{
		model: authorModel,
	}, nil
}

func (ctrl *AuthorsControllers) GetAthors(c *fiber.Ctx) error {
	UserID := c.Params(constants.ParamUid)
	// string to int32

	UserIDInt, err := strconv.Atoi(UserID)
	if err != nil {
		return nil
	}

	author, err := ctrl.model.GetAuthor(c.Context(), int32(UserIDInt))
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetAuthor)
	}

	return utils.JSONSuccess(c, 200, author)
}
