package main

import (
	"awesomeProject1/pkg/models/mysql"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// получаем доступ из любого места
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")
	dsn := flag.String("dsn", "root:QWEqwe136479@/snippetbox?parseTime=true", "Название MySQL источника данных")
	flag.Parse()

	//кастомные ошибки
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	// как
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}
	//  пока не созд тут структуру(указатель)  что выше не было доступно app.
	//  потому что мы созд указатель и таким обрзо сделали видимсоть?

	//  обяъснить Олегу Созд указатель а где общая структура???
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запуск на %v", *addr)

	err = srv.ListenAndServe()
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
