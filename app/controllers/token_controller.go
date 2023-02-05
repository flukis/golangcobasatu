package controllers

import "github.com/gofiber/fiber/v2"

func RenewToken(c *fiber.Ctx) error {
	/*
	 * 1. Get time now
	 * 2. Get claims from jwt
	 * 3. Set expires time
	 * 4. Check if now > access token time
	 * 5. Create new refresh token
	 * 6. Checking received data from JSON body
	 */
}
