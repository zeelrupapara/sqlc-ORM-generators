package routes

import (
	"database/sql"
	"fmt"
	"sync"

	"go.uber.org/zap"

	"github.com/sqlc_test/config"
	"github.com/sqlc_test/constants"
	controller "github.com/sqlc_test/controllers/api/v1"

	// "github.com/sqlc_test/constants"

	"github.com/gofiber/fiber/v2"
	"github.com/sqlc_test/middlewares"
)

var mu sync.Mutex

// Setup func
func Setup(app *fiber.App, db *sql.DB, logger *zap.Logger, config config.AppConfig) error {
	mu.Lock()

	app.Use(middlewares.LogHandler(logger))

	app.Static("/assets/", "./assets")

	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Render("./assets/index.html", fiber.Map{})
	})

	router := app.Group("/api")
	v1 := router.Group("/v1")

	middlewares := middlewares.NewMiddleware(config, logger)

	err := setupAuthorController(v1, db, logger, middlewares)
	if err != nil {
		return err
	}

	// err := setupUserController(v1, goqu, logger, middlewares)
	// if err != nil {
	// 	return err
	// }

	mu.Unlock()
	return nil
}

// func setupUserController(v1 fiber.Router, goqu *goqu.Database, logger *zap.Logger, middlewares middlewares.Middleware, events *events.Events) error {
// 	userController, err := controller.NewUserController(goqu, logger, events)
// 	if err != nil {
// 		return err
// 	}

// 	userRouter := v1.Group("/users")
// 	userRouter.Post("/", userController.CreateUser)
// 	userRouter.Get(fmt.Sprintf("/:%s", constants.ParamUid), middlewares.Authenticated, userController.GetUser)
// 	return nil
// }

func setupAuthorController(v1 fiber.Router, db *sql.DB, logger *zap.Logger, middlewares middlewares.Middleware) error {
	authorController, err := controller.NewAuthorsController(db, logger)
	if err != nil {
		return err
	}

	authorRouter := v1.Group("/authors")
	authorRouter.Get(fmt.Sprintf("/:%s", constants.ParamUid), authorController.GetAthors)
	return nil
}
