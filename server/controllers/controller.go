package controller

import (
	"Chat-Application/database"
	"Chat-Application/models"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "secret"

func Hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to the chat application")
}

func Register(c *fiber.Ctx) error {
	fmt.Println("Recieved a registration request")

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "failed to parse the body request",
		})
	}
	fmt.Println("user name: ", data["name"], "email:", data["email"])
	// check if the email already exists

	var existingUser models.User

	if err := database.DB.Where("email= ?", data["email"]).First(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "the email already exists",
		})
	}

	// passwords
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error hashing the password",
		})
	}

	fmt.Printf("creating the user....")

	user := &models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: hashedPassword,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error creating the database",
		})
	}

	fmt.Println("user registered successfully")
	return c.JSON(fiber.Map{
		"message": "User registered Successfully",
	})
}

func Login(c *fiber.Ctx) error {
	fmt.Println("the request for login is succeded")

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error parsing the body",
		})
	}
	var user models.User
	database.DB.Where("email=?", data["email"]).first(&user)
	if user.ID == 0 {
		fmt.Println("User not found")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
	if err != nil {
		fmt.Println("User not found")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// creating a JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(user.ID)),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error generating token:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	// creating a cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), // expire in 24 hour
		HTTPOnly: true,
		Secure:   true,
	}
	c.Cookie(&cookie)

	fmt.Println("authentication successfull")

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Message": "Login Successfull",
	})

}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), //
		HTTPOnly: true,
		Secure:   true,
	}
	c.Cookie(&cookie)

	fmt.Println("authentication successfull")

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Message": "Logout Successfull",
	})

}
