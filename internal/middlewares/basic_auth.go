package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func (mw *Middlewares) BasicAuthMW() func(*fiber.Ctx) error {
	user1name := os.Getenv("USER_1_NAME")
	user1pass := os.Getenv("USER_1_PASS")
	user2name := os.Getenv("USER_2_NAME")
	user2pass := os.Getenv("USER_2_PASS")
	basicAuth := basicauth.Config{
		Users: map[string]string{
			user1name: user1pass,
			user2name: user2pass,
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == user1name && pass == user1pass {
				return true
			}
			if user == user2name && pass == user2pass {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			c.Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "You do not have accesss to this page",
			})
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}

	err := basicauth.New(basicAuth)

	return err
}
