package main

import (
	"fmt"

	"github.com/yanzay/tbot"
)

const TELEGRAM_TOKEN = "YUOR_TOKEN_BOT"
const HTTP_GAME = "https://openalmanac.altervista.org/ifgamebook/index.html" //YOUR GAME BOOK OR IF ON YOUR SITE

type application struct {
	client *tbot.Client
}

func main() {
	token := TELEGRAM_TOKEN
	bot := tbot.New(token)
	app := &application{}
	app.client = bot.Client()
	bot.HandleMessage("/start", app.gameHandler)
	bot.HandleCallback(app.callbackHandler)
	bot.Start()
}

func (a *application) gameHandler(m *tbot.Message) {
	_, err := a.client.SendGame(m.Chat.ID, "ultrasecret") //Insert Short Name used/configured in Bot Father
	if err != nil {
		fmt.Println(err)
	}
}

func (a *application) callbackHandler(cq *tbot.CallbackQuery) {
	a.client.AnswerCallbackQuery(cq.ID, tbot.OptURL(HTTP_GAME))
}
