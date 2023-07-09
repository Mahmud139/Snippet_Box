package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"html/template"
	"fmt"
	"runtime/debug"

	"github.com/mahmud139/Snippet_Box/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippet  *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", "localhost:8080", "HTTP network address")
	dsn := flag.String("dsn", "web:mahmud@/snippetbox?parseTime=true", "MySql data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache("M:/Projects/Snippet_box/ui/html/")
	if err != nil {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		errorLog.Output(2, trace)
		errorLog.Panic(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippet:  &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %v\n", *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
