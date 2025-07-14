package main

// Soal no 2
import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
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

var rdbglobal = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {
	// setRedisData("adss", "{
	// 		“realname”:”Aberto Doni Sianturi”,
	// 		“email”:”adss@gmail.com”
	// 		“password”:”f7c3bc1d808e0 . . . 441”
	// 		}");
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

var ctx = context.Background()

func setRedisData(key, value string) string {
	err := rdbglobal.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdbglobal.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	// fmt.Println("key", val)
	return val
}

// dari dokumentasi
func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
