package main

import (
	"os"

	"github.com/dedit-hs/go_final-project_bmt-be/configs"
	"github.com/dedit-hs/go_final-project_bmt-be/middlewares"
	"github.com/dedit-hs/go_final-project_bmt-be/routes"
)

// HELLO

func init() {
	configs.LoadEnv()
	configs.InitDB()
}

func main() {
	e := routes.New()
	middlewares.GlobalMiddleware(e)
	e.Logger.Fatal(e.Start(envPortOr("3003")))
}

func envPortOr(port string) string {
	envPort := os.Getenv("FPAPP_PORT")
	if envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
