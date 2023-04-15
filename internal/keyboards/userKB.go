package keyboards

import tele "gopkg.in/telebot.v3"

var (
	userMenu       = &tele.ReplyMarkup{ResizeKeyboard: true, RemoveKeyboard: true}
	userInlineMenu = &tele.ReplyMarkup{}
	BtnHelp        = userMenu.Text("ℹ Help")
	BtnShowSnippet = userMenu.Text("👁 Show")
	btnPrev        = userInlineMenu.Data("⬅", "prev", "1")
	btnNext        = userInlineMenu.Data("➡", "next", "2")
)

func ShowSnippets() *tele.ReplyMarkup {
	userInlineMenu.Inline(userInlineMenu.Row(btnPrev, btnNext))
	return userInlineMenu
}

func StartKB() *tele.ReplyMarkup {
	userMenu.Reply(
		userMenu.Row(BtnShowSnippet),
		userMenu.Row(BtnHelp))
	return userMenu
}
