package keyboards

import tele "gopkg.in/telebot.v3"

var (
	labaMenu         = &tele.ReplyMarkup{ResizeKeyboard: true, RemoveKeyboard: true}
	BtnShowLaba      = userMenu.Text("Show table")
	BtnNewRecordLaba = userMenu.Text("Create")
)

func LabaKB() *tele.ReplyMarkup {
	labaMenu.Reply(
		labaMenu.Row(BtnShowLaba, BtnNewRecordLaba),
	)
	return labaMenu
}
