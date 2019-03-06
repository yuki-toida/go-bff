package main

import (
	"context"
	"github.com/go-kit/kit/log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-bff/bff/adapter/repositories/repository_email"
	"go-bff/bff/adapter/repositories/repository_profile"
	"go-bff/bff/adapter/repositories/repository_user"
	"go-bff/bff/external/config"
	"go-bff/bff/external/microservices"
	"go-bff/bff/external/mysql"
	"go-bff/bff/external/web"
	"go-bff/bff/registry"
	"net/http"
	"os"
	"strconv"
)

func main() {
	c := config.Load()
	db := mysql.Connect(c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Name)
	defer db.Close()

	db.LogMode(true)
	db.DropTableIfExists(&repository_user.User{}, &repository_profile.Profile{}, &repository_email.Email{})
	db.AutoMigrate(&repository_user.User{}, &repository_profile.Profile{}, &repository_email.Email{})

	for i := 1; i < 3; i++ {
		p := repository_profile.Profile{Name: strconv.Itoa(i)}
		db.Create(&p)
		u := repository_user.User{ProfileID: p.ID}
		db.Create(&u)
		for j := 1; j < 4; j++ {
			e := strconv.Itoa(j) + "@hacobu.jp"
			db.Create(&repository_email.Email{Email: e, UserID: u.ID})
		}
	}

	emailconn := microservices.Connect(c.Email.Host, c.Email.Port)
	defer emailconn.Close()

	repositories := registry.NewRepositories(db)
	usecases := registry.NewUseCases(repositories)
	cxt := context.Background()
	microservices := registry.NewMicroServices(emailconn)
	handler := web.Handle(usecases, microservices, cxt)

	port := ":8080"
	logger := log.NewLogfmtLogger(os.Stderr)
	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(port, handler))

}
