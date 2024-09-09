package middleware

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
	"sygap_new_knowledge_management/backend/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CustomRecoverMiddleware(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			panicMessage := fmt.Sprintf("%v", r)
			log.Printf("Recovered from panic: %v", panicMessage)

			// trace detail error go runtime
			stackTrace := debug.Stack()
			log.Printf("Panic problem: %v", string(stackTrace))
			c.Status(500).JSON(utils.ResponseData{
				StatusCode: 500,
				Message:    "Error panic",
				Error:      panicMessage,
				Data:       string(stackTrace),
			})
		}

		duration := 240 * time.Second
		ctx, cancel := context.WithTimeout(c.Context(), duration)
		defer cancel()
		c.SetUserContext(ctx)
	}()
	return c.Next()
}
