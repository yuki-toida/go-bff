package main

import (
	"context"
	"github.com/go-kit/kit/log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-bff/bff/adapter/gateways/db/db_email"
	"go-bff/bff/adapter/gateways/db/db_profile"
	"go-bff/bff/adapter/gateways/db/db_user"
	"go-bff/bff/external/config"
	"go-bff/bff/external/grpc"
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
	db.DropTableIfExists(&db_user.User{}, &db_profile.Profile{}, &db_email.Email{})
	db.AutoMigrate(&db_user.User{}, &db_profile.Profile{}, &db_email.Email{})

	for i := 1; i < 3; i++ {
		p := db_profile.Profile{Name: strconv.Itoa(i)}
		db.Create(&p)
		u := db_user.User{ProfileID: p.ID}
		db.Create(&u)
		for j := 1; j < 4; j++ {
			e := strconv.Itoa(j) + "@hacobu.jp"
			db.Create(&db_email.Email{Email: e, UserID: u.ID})
		}
	}

	notifyConn := grpc.Connect(c.NotifyService.Host, c.NotifyService.Port)
	defer notifyConn.Close()

	reg := registry.New(db, context.Background(), notifyConn)
	handler := web.Handle(reg)

	port := ":8080"
	logger := log.NewLogfmtLogger(os.Stderr)

	if err := logger.Log("msg", "HTTP", "addr", port); err != nil {
		panic(err)
	}
	if err := logger.Log("err", http.ListenAndServe(port, handler)); err != nil {
		panic(err)
	}

}
