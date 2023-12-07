package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	perf "github.com/gb-duke/ie2_db/src/decorators"
	"github.com/gb-duke/ie2_db/src/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/streadway/amqp"
)

func main() {

	r := mux.NewRouter()

	conn, ch := initRMQ()
	if conn == nil || ch == nil {
		panic("RMQ Connection failed...")
	}

	defer conn.Close()
	defer ch.Close()

	// TODO: make the api version load dynamically so we don't hardcode it
	userStore := handlers.UsersStore{}
	userStore.Init()
	r.HandleFunc("/api/v1/users", perf.RestPerf(userStore.GetAll, ch)).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", perf.RestPerf(userStore.Get, ch)).Methods("GET")
	r.HandleFunc("/api/v1/users", perf.RestPerf(userStore.Create, ch)).Methods("POST")
	r.HandleFunc("/api/v1/users/{id}", perf.RestPerf(userStore.Update, ch)).Methods("PUT")
	r.HandleFunc("/api/v1/users/{id}", perf.RestPerf(userStore.Delete, ch)).Methods("DELETE")

	duStore := handlers.DataUploadStore{}
	duStore.Init()
	r.HandleFunc("/api/v1/datauploads", perf.RestPerf(duStore.GetAll, ch)).Methods("GET")
	r.HandleFunc("/api/v1/datauploads/{id}", perf.RestPerf(duStore.Get, ch)).Methods("GET")
	r.HandleFunc("/api/v1/datauploads", perf.RestPerf(duStore.Create, ch)).Methods("POST")
	r.HandleFunc("/api/v1/datauploads/{id}", perf.RestPerf(duStore.Update, ch)).Methods("PUT")
	r.HandleFunc("/api/v1/datauploads/{id}", perf.RestPerf(duStore.Delete, ch)).Methods("DELETE")

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*", "http://localhost:3000"},
		AllowedHeaders: []string{
			"Accept",
			"Accept-Encoding",
			"Accept-Language",
			"Cache-Control",
			"Connection",
			"DNT",
			"Host",
			"Origin",
			"Pragma",
			"Referer",
			"User-Agent",
		},
		AllowedMethods: []string{
			"DELETE",
			"GET",
			"OPTIONS",
			"POST",
			"PUT",
		},
	})

	// TODO: set port in env
	log.Println("API Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsOpts.Handler(r)))
}

func initRMQ() (*amqp.Connection, *amqp.Channel) {
	queue := os.Getenv("RMQ_QNAME")
	uname := os.Getenv("RMQ_UNAME")
	pwd := os.Getenv("RMQ_PWD")
	domain := os.Getenv("RMQ_URL")

	if queue == "" {
		panic("RMQ Queue Name is empty")
	}

	if uname == "" {
		panic("RMQ Username is empty")
	}

	if pwd == "" {
		panic("RMQ Pwd is empty")
	}

	if domain == "" {
		panic("RMQ Domain is empty")
	}

	rmq := fmt.Sprintf("amqp://%s:%s@%s/", uname, pwd, domain)
	conn, err := amqp.Dial(rmq)
	if err != nil {

		// wait a moment and try again...
		time.Sleep(10 * time.Second)

		conn, err = amqp.Dial(rmq)

		if err != nil {
			panic(err)
		}
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	q, err := ch.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(q)

	return conn, ch
}
