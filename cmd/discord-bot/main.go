package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {}

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalln("No $BOT_TOKEN given.")
	}

	h := newHandler(state.New("Bot " + token))
	h.s.AddInteractionHandler(h)
	h.s.AddIntents(gateway.IntentGuilds)
	h.s.AddHandler(func(*gateway.ReadyEvent) {
		me, _ := h.s.Me()
		log.Printf("Connected to the gateway as '%s'", me.Tag())

		// request a list of 100 guilds
		guilds, _ := h.s.Guilds()
		//		var guild_names []string
		log.Print("The bot is on the following servers: ")
		for _, g := range guilds {
			//			guild_names = append(guild_names, g.Name)
			log.Printf("\tserver=%s, id=%s", g.Name, g.ID.String())
		}
	})

	if err := overwriteCommands(h.s); err != nil {
		log.Fatalln("cannot update commands:", err)
	}

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Printf("Prometheus metrics exposition at :2112/metrics")
		if err := http.ListenAndServe(":2112", nil); err != nil {
			log.Fatal(err)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := h.s.Connect(ctx); err != nil {
		log.Fatalln("cannot connect:", err)
	}
}
