package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}


	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/healthz", Healthz)


	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()


}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	for k, vs := range r.Header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}

	w.Header().Add("VERSION", os.Getenv("VERSION"))
	log.Print("ip==", r.RemoteAddr)

	io.WriteString(w, "ok")

}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
