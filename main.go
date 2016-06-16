package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-martini/martini"
	"github.com/kardianos/service"
)

var (
	logger service.Logger

	config *service.Config = &service.Config{
		Name:        "sayhi",
		DisplayName: "hi, boy",
	}
)

var (
	ErrNotFound []byte = []byte("404 Not Found")
)

const (
	PORT   string = ":1234"
	layout string = "2006-01-02 15:04:05"
)

type App struct {
	Pid int
}

func (a *App) Start(s service.Service) (err error) {
	go a.run(s)
	return nil
}

func (a *App) run(s service.Service) {
	a.Pid = os.Getpid()
	m := martini.Classic()

	m.Use(static)
	m.Use(parseForm)

	m.Get("/", func() string {
		now := time.Now()
		return "Hello, boy! Today is " + now.Format(layout)
	})

	m.Get("/login", func() []byte {
		tpl, err := AssetHtml("login")
		if err != nil {
			return ErrNotFound
		}
		return tpl
	})

	m.Post("/login/check", func(r *http.Request) []byte {
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if strings.EqualFold("kyf", username) &&
			strings.EqualFold("keyongfeng", password) {
			return encodeResponse(true, "success")
		} else {
			return encodeResponse(false, "username or password is invalid")
		}

	})

	err := http.ListenAndServe(PORT, m)
	if err != nil {
		logger.Error(err)
	}
}

func (a *App) Stop(s service.Service) (err error) {
	process, err := os.FindProcess(a.Pid)
	if err != nil {
		return err
	}

	return process.Kill()
}

func main() {
	app := &App{}
	s, err := service.New(app, config)
	if err != nil {
		log.Printf("init service err:%v", err)
		os.Exit(1)
	}

	action := ""

	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	if len(action) > 0 {
		err = service.Control(s, action)
		if err != nil {
			log.Printf("init service err:%v", err)
			os.Exit(1)
		}
		return
	}

	logger, err = s.Logger(nil)
	if err != nil {
		return
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
