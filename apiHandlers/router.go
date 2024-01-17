package apiHandlers

import (
	"TeamManagementSystem/api"

	"encoding/gob"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Router(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	store := session.New()
	//encryptcookie.New(encryptcookie.Config{
	//	Key: "secret-thirty-2-character-string",
	//})
	gob.Register(map[string]interface{}{})

	api := app.Group("/TeamManagementSystem/api")
	check := app.Group("/")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	DefaultMappings(check, store)
	RouteMappings(api,store)
}

func RouteMappings(cg fiber.Router, store *session.Store) {
cg.Post("/CreateTeam", api.CreateTeamApi)
cg.Put("/UpdateTeam", api.UpdateTeamApi)
cg.Delete("/DeleteTeam", api.DeleteTeamApi)
cg.Get("/FindTeam", api.FindTeamApi)
cg.Get("/FindallTeam", api.FindallTeamApi)
cg.Post("/CreateMember", api.CreateMemberApi)
cg.Put("/UpdateMember", api.UpdateMemberApi)
cg.Delete("/DeleteMember", api.DeleteMemberApi)
cg.Get("/FindMember", api.FindMemberApi)
cg.Get("/FindallMember", api.FindallMemberApi)
cg.Post("/CreateRoles", api.CreateRolesApi)
cg.Put("/UpdateRoles", api.UpdateRolesApi)
cg.Delete("/DeleteRoles", api.DeleteRolesApi)
cg.Get("/FindRoles", api.FindRolesApi)
cg.Get("/FindallRoles", api.FindallRolesApi)


}

func DefaultMappings(cg fiber.Router, store *session.Store) {
	cg.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "TeamManagementSystem-APP359 service is up and running", "version": "1.0"})
	})
}