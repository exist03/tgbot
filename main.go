package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"tgbot/internal/handlers"
	"tgbot/pkg/models/mysql"
	"time"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found")
	}
	pref := tele.Settings{
		Token: os.Getenv("BOT_TOKEN"),
		Poller: &tele.LongPoller{
			Timeout: 5 * time.Second,
		},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		panic("can`t create bot")
	}
	db, err := openDB("quest:quest@/snippets")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("db is open")
	defer db.Close()

	m := mysql.SnippetModel{DB: db}
	handlers.InitHandlers(b, &m)
	b.Start()
}
func openDB(dst string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dst)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
