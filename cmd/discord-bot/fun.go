package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/utils/json/option"

	infinite "github.com/infinite-industries/infinite-go"
)

var infinite_client = infinite.New()

func init() {
	// seed the random number generator
	rand.Seed(time.Now().Unix())
}

func (h *handler) cmdFun(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {

	// add logging here - if there is no response whatsover, what happened?

	guild_id := data.Event.GuildID.String()
	//	if guild_id == "" {
	//		guild_id = "dm"
	//	}

	if counter, err := bot_processed_requests_total.GetMetricWithLabelValues(guild_id); err != nil {
		log.Printf("Error accessing metric counter")
	} else {
		counter.Inc()
	}

	// log request
	// TODO: use h.s.Guild(guild_id) to get server name.
	log.Printf("'%s' interacted in channel '%s', server id '%s'", data.Event.Sender().Username, data.Event.Channel.Name, data.Event.GuildID)
	// fetch events
	// TODO: is it possible for this to timeout w/out error & response?  See 10/25/2023 @ 21:49.
	events, err := infinite_client.Events.CurrentVerified()
	if err != nil {
		log.Printf("Error accessing the Infinite API: %s", err)
		return &api.InteractionResponseData{
			Content: option.NewNullableString("Sorry, I couldn't access the Infinite Industries API"),
		}
	}
	// pick the Title / URL for a random one
	event := events[rand.Intn(len(events))]
	// return a markdown response: [Title](URL)
	output := fmt.Sprintf("[%s](https://infinite.industries/events/%s)", event.Title, event.ID)
	return &api.InteractionResponseData{
		Content: option.NewNullableString(output),
	}
}
