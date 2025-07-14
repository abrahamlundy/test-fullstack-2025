package main

// Soal no 2
import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	RealName string
	Email    string
	Password string // hashed password
}

var users = map[string]User{
	"aberto": {
		RealName: "Aberto Doni Sianturi",
		Email:    "aberto@gmail.com",
		Password: hash("123456"),
	},
}

func main() {

	// Struktur User
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	app.Get("/", func(c *fiber.Ctx) error {
		html := `
        <html>
        <head><title>Login</title></head>
        <body>
            <h2>Form Login</h2>
            <form action="/login" method="POST">
                <label>Username:</label><br>
                <input type="text" name="username"><br>
                <label>Password:</label><br>
                <input type="password" name="password"><br><br>
                <input type="submit" value="Login">
            </form>
        </body>
        </html>`
		return c.Type("html").SendString(html)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		user, ada := users[username]
		if !ada {
			fmt.Println("Login gagal: user tidak ada.")
			return c.Status(401).SendString("User tidak ada.")
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			fmt.Println("Login gagal: password salah", username)
			return c.Status(401).SendString("Password salah.")
		}

		fmt.Println("Login berhasil untuk:", username)
		return c.SendString("Login berhasil. Selamat datang, " + user.RealName + "!")
	})

	app.Listen(":3000")
}

func hash(pw string) string {
	h, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Gagal hash password:", err)
		return ""
	}
	return string(h)
}
