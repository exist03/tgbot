package handlers

import (
	tele "gopkg.in/telebot.v3"
	"strings"
	keyboards2 "tgbot/internal/keyboards"
	"tgbot/pkg/models/mysql"
)

func InitHandlers(b *tele.Bot, m *mysql.SnippetModel) {
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello!", keyboards2.StartKB())
	})

	b.Handle("/admin", func(c tele.Context) error {
		return c.Send("Hello!", keyboards2.StartKB())
	})

	b.Handle(&keyboards2.BtnHelp, func(c tele.Context) error {
		return c.Send("If you have any problems contact @qazwsxedc327")
	})

	b.Handle("/laba", func(c tele.Context) error {
		return c.Send("Choose", keyboards2.LabaKB())
	})

	b.Handle(&keyboards2.BtnShowSnippet, func(c tele.Context) error {
		return c.Send("smth", keyboards2.ShowSnippets())
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		snippetSlice := strings.Split(c.Text(), " ")
		if len(snippetSlice) != 2 {
			return c.Send("Incorrect message")
		}
		err := m.Insert(c.Sender().ID, snippetSlice[0], snippetSlice[1])
		if err != nil {
			return err
		}
		return c.Send(snippetSlice[0])
	})

	b.Handle(tele.OnCallback, func(c tele.Context) error {
		return c.Send(c.Query())
	})
}
