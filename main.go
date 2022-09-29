package main

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// LINE Botクライアント生成する
	// BOTにチャネルシークレットとチャネルトークンを引数にわたす
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
	}

	//天気情報を取得
	result, err := weather.GetWeather()

	if err != nil {
		log.Fatal(err)
	}

	//テキストメッセージを生成
	message := linebot.NewTextMessage(result)
	//テキストメッセージをユーザーに配信
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
