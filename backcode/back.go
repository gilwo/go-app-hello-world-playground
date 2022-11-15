package backcode

import (
	"context"
	"fmt"
	"goappex"
	"log"
	"net/http"
	"os"
	"syscall"
	"time"

	_ "goappex/frontcode"

	"github.com/gin-gonic/gin"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/cli"
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
)

func init() {
	goappex.Mainback = mainback
}

var (
	hB *app.Handler = &app.Handler{
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
		Title: "hello exampler 2",
	}

	webrtcDCB *app.Handler = &app.Handler{
		Name:        "webrtc data channels example",
		Description: "webrtc data channels front side for the example within pion : https://github.com/pion/webrtc/blob/master/examples/data-channels ",
		Styles: []string{
			"/web/webrtc.css",
			"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css",
		},
		Title: "webrtc data-channels wasm example",
		// Resources: ,
		LoadingLabel: "bluuuuuu{progress}%",
	}
)

func mainback() {

	if goappex.Mainfront == nil {
		panic("cant find front code logic")
	}
	goappex.Mainfront()

	ctx, cancel := cli.ContextWithSignals(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()
	defer exit()

	var hRef http.Handler
	if useGin {
		hRef = initGin()
	} else {
		http.Handle(goappex.HelloPath, hB)
		http.Handle(goappex.WebrtcDataChannelsPath, webrtcDCB)
	}
	srv := &http.Server{
		Addr:    ":8000",
		Handler: hRef,
	}
	fmt.Printf("*** started on <%v> ***", srv.Addr)
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
}

var (
	// useGin = false
	useGin = true
)

func initGin() *gin.Engine {

	r := gin.Default()

	foo := func(c *gin.Context) {
		fmt.Printf("requestd path : %s\n", c.Request.URL)
		hB.ServeHTTP(c.Writer, c.Request)
	}
	bar := func(c *gin.Context) {
		fmt.Printf("requestd path : %s\n", c.Request.URL)
		webrtcDCB.ServeHTTP(c.Writer, c.Request)
	}

	r.GET(goappex.WebrtcDataChannelsPath, bar)
	r.GET(goappex.HelloPath, foo)
	r.GET("/web/hello-main.css", foo)
	r.GET("/web/webrtc.css", foo)
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

	return r
}

func exit() {
	err := recover()
	if err != nil {
		app.Logf("command failed: %s", errors.Newf("%v", err))
		os.Exit(-1)
	}
}
