package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	token := "7546240919:AAFGiv8wxv_kXf7owVm0CLA-wZxnab84vBo"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch {
		case strings.HasPrefix(update.Message.Text, "/start"):
			// Главное меню
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я расскажу тебе о парсинге данных. Что хочешь узнать?")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Расскажи о парсинге"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

		case update.Message.Text == "Расскажи о парсинге":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "О чём хочешь узнать?\n1. Методы парсинга\n2. Популярные библиотеки\n3. Проблемы при парсинге")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Методы парсинга"),
					tgbotapi.NewKeyboardButton("Популярные библиотеки"),
					tgbotapi.NewKeyboardButton("Проблемы при парсинге"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("В начало"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

		case update.Message.Text == "Методы парсинга":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Методы парсинга данных:\n1. Использование регулярных выражений\n2. Парсинг с помощью HTML парсеров\n3. API запросы для получения данных")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Регулярные выражения"),
					tgbotapi.NewKeyboardButton("HTML парсеры"),
					tgbotapi.NewKeyboardButton("API запросы"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("В начало"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

		case update.Message.Text == "Регулярные выражения":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Регулярные выражения — это мощный инструмент для поиска и извлечения данных из строк текста. Они используются, например, для извлечения данных из HTML-кода или логов.")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("В начало"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

		case update.Message.Text == "HTML парсеры":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "HTML парсеры, такие как BeautifulSoup или GoQuery, позволяют извлекать данные из HTML-страниц, анализируя DOM-дерево.")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("В начало"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

		case update.Message.Text == "API запросы":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Парсинг через API запросы позволяет получить данные в структурированном формате, например, JSON. Это позволяет обходиться без анализа HTML.")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("В начало"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

		case update.Message.Text == "Популярные библиотеки":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Популярные библиотеки для парсинга данных:\n1. BeautifulSoup (Python)\n2. GoQuery (Go)\n3. Cheerio (Node.js)")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("BeautifulSoup"),
					tgbotapi.NewKeyboardButton("GoQuery"),
					tgbotapi.NewKeyboardButton("Cheerio"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("В начало"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

		case update.Message.Text == "BeautifulSoup":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "BeautifulSoup — это библиотека Python для парсинга HTML и XML документов. Она позволяет легко извлекать информацию из веб-страниц.")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("В начало"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

		case update.Message.Text == "GoQuery":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "GoQuery — это библиотека Go для парсинга HTML. Она использует jQuery-подобный синтаксис для работы с DOM.")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("В начало"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

		case update.Message.Text == "Cheerio":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Cheerio — это библиотека Node.js, которая предоставляет jQuery-подобный интерфейс для работы с HTML.")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("В начало"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

		case update.Message.Text == "Проблемы при парсинге":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Проблемы при парсинге:\n1. Защита от парсинга (CAPTCHA, защита от ботов)\n2. Изменения в структуре HTML\n3. Большие объемы данных")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("CAPTCHA и защита от ботов"),
					tgbotapi.NewKeyboardButton("Изменения в HTML"),
					tgbotapi.NewKeyboardButton("Большие объемы данных"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("В начало"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

		case update.Message.Text == "CAPTCHA и защита от ботов":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Парсинг с использованием CAPTCHA требует дополнительных методов обхода, например, с использованием сервисов для решения CAPTCHA.")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("В начало"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

		case update.Message.Text == "Изменения в HTML":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Изменения в структуре HTML могут привести к поломке парсера. Лучше использовать более устойчивые методы парсинга, например, API.")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("В начало"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

		case update.Message.Text == "Большие объемы данных":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "При работе с большими объемами данных важно учитывать скорость работы парсера и возможные ограничения сервера.")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("В начало"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

		case update.Message.Text == "В начало":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в главное меню. Что хотите узнать?")
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Расскажи о парсинге"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)
		}
	}
}
