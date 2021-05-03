package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// PLAN
	// 2 URL only
		// load_wp.com/call_wp/blog
		// load_wp.com/call_wp/tips
	// load data from .env file for WP URL & construct different URL parameter
	// call the APIs sequentially
	// connect to Redis and save the data in Redis
		// Redis Keys will also be in env variable

	// Future enhancement
		// Call the Wordpress APIs parallely [tags/categories/post-list]
		// run this service periodically on its own, without invoking.

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
}

func main() {
	addr := flag.String("addr", ":4001", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)

	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}