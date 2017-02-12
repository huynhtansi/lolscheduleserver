package main

import (
	"fmt"
	"github.com/huynhtansi/lolscheduleserver/app/config"
	"log"
	"flag"
	"github.com/huynhtansi/lolscheduleserver/app/core/database"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"lolscheduleserver/app/api"
	"github.com/huynhtansi/lolscheduleserver/app/core/entity"
)

var flagProfile string


func init() {
	flag.StringVar(&flagProfile, "profile", "", "-profile")
}

func bootstrap()  {

	// Init Environment
	if err := config.InitConf(flagProfile); err != nil {
		//panic(err)
		log.Fatal(err)
	}
	// Connect to database
	if err := database.Connect(config.Env.Database); err != nil {
		panic(err)
	}

	//Create tables base on Entities (in public schema)
	database.SQL.CreateTable(entity.Team{})
}

func main()  {

	fmt.Print("Stating Lol Schedule server...")
	bootstrap()


	e := echo.New()

	if config.Env.Debug {
		//e.Debug(true)
		//echopprof.Wrapper(e)
	}

	////////////////// Root middleware ///////////////
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())


	e.GET("/favicon.ico", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	//api := api.API{}

	// Run the server
	address := fmt.Sprintf("%s:%d", config.Env.Host, config.Env.Port)
	if config.Env.Debug {
		e.Logger.Fatal(e.Start(address))
	}
	e.Logger.Fatal(e.StartTLS(address, "../config/lottery_server.crt", "../config/lottery_server.key"))
}
