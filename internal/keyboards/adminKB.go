package keyboards

import tele "gopkg.in/telebot.v3"

var (
	adminMenu    = &tele.ReplyMarkup{ResizeKeyboard: true, RemoveKeyboard: true}
	btnShowUsers = userMenu.Text("Users list")
)
