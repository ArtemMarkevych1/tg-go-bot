import "github.com/spf13/cobra"
import (
    "fmt"
    "github.com/spf13/cobra"
    "gopkg.in/telebot.v3"
)

token := os.Getenv("TELE_TOKEN")

func main() {
    bot, err := telebot.NewBot(telebot.Settings{
        Token:  token,
        Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
    })

    if err != nil {
        fmt.Println(err)
        return
    }

    bot.Handle(telebot.OnText, onText)

    bot.Start()
}

func onText(ctx telebot.Context) error {
    msg := ctx.Message().Text
    fmt.Println(msg)

    return ctx.Send("Привет! Я получил ваше сообщение: " + msg)
}


bot, err := telebot.NewBot(telebot.Settings{
    Token:  token,
    Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
})

if err != nil {
    fmt.Println(err)
    return
}

func onText(ctx telebot.Context) error {
    msg := ctx.Message().Text
    fmt.Println(msg)

    return ctx.Send("Привет! Я получил ваше сообщение: " + msg)
}