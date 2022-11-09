package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/cli"
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
)

type hello struct {
	app.Compo
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.Div().Class("image-title").Body(
			app.Img().
				Alt("butterfly").
				Src("/web/logo2.png").
				Width(400).Height(300),
		),
		app.Div().Body(
			app.H1().Class("hello-title").Text("Hello World!"),
		))
}

func main() {

	hF := &hello{}
	app.Route(helloPath, hF)

	app.RunWhenOnBrowser()

	ctx, cancel := cli.ContextWithSignals(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()
	defer exit()

	hB := &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		// Resources:   app.CustomProvider(".", helloPath),
		Icon: app.Icon{
			Default: "/web/192.png",
			Large:   "/web/logo2.png",
		},
		// Resources: app.LocalDir("/web"),
		// Resources: app.LocalDir(""),H1
		Styles: []string{
			"/web/hello-main.css",
		},
		Title: "hello exampler",
	}

	if useGin {
		r := gin.Default()

		foo := func(c *gin.Context) {
			fmt.Printf("requestd path : %s\n", c.Request.URL)
			hB.ServeHTTP(c.Writer, c.Request)
		}

		r.GET(helloPath, foo)
		r.GET("/web/hello-main.css", foo)
		r.GET("/favicon.ico", foo)
		r.GET("/web/logo2.png", foo)
		r.GET("/web/logo.png", foo)
		r.GET("/web/192.png", foo)
		r.GET("/app.css", foo)
		r.GET("/wasm_exec.js", foo)
		r.GET("/web/app.wasm", foo)
		r.GET("/app.js", foo)
		r.GET("/manifest.webmanifest", foo)
		r.GET("/app-worker.js", foo)

		srv := &http.Server{
			Addr:    ":8000",
			Handler: r,
		}
		fmt.Println("*** started ***")
		go func() {
			<-ctx.Done()
			// fmt.Println("someone invoked cancel")
			srv.Shutdown(context.Background())
			// fmt.Println("shutdown issued")
		}()

		go func() {
			if err := srv.ListenAndServe(); err != nil {
				if err != http.ErrServerClosed {
					log.Fatalf("failed serving with gin: %s", err)
				}
			}
		}()
		<-ctx.Done()
		time.Sleep(100 * time.Millisecond)

		fmt.Println("*** ended ***")

	} else {
		http.Handle(helloPath, hB)

		fmt.Println("started")
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Fatal(err)
		}
	}
}

var (
	helloPath = "/"
	// helloPath = "/helo"
	// useGin = false
	useGin = true
)

func exit() {
	err := recover()
	if err != nil {
		app.Logf("command failed: %s", errors.Newf("%v", err))
		os.Exit(-1)
	}
}
