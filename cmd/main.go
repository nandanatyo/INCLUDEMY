package main

import (
	"includemy/internal/handler/rest"
	"includemy/internal/repository"
	"includemy/internal/service"
	"includemy/pkg/bcrypt"
	"includemy/pkg/config"
	"includemy/pkg/database/mysql"
	"includemy/pkg/jwt"
	"includemy/pkg/middleware"
	"includemy/pkg/supabase"
	"log"
)

func main() {

	config.LoadEnvirontment()
	config.LoadMidtransConfig()
	db := mysql.ConnectDatabase()
	err := mysql.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(db)
	bcrypt := bcrypt.Init()
	jwt := jwt.Init()
	supabase := supabase.Init()
	svc := service.NewService(repo, bcrypt, jwt, supabase)
	middleware := middleware.Init(jwt, svc)
	r := rest.NewRest(svc, middleware)
	r.MountEndpoints()
	r.Run()
}
