package main

import (
	"context"
	"fmt"
	"github.com/lukinairina90/thecatsapi/catsapi"
	"log"
)

const APIKey = "" // your api key
const BaseUrl = "https://api.thecatapi.com/v1"
const UserID = "User-123" // your name to vote

func main() {
	catAPIClient := catsapi.NewClient(BaseUrl, APIKey, UserID)

	ctx := context.Background()

	list, err := catAPIClient.List(ctx, &catsapi.ListParams{
		Limit:     100,
		Page:      3,
		DescOrder: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range list {
		fmt.Println(val.Info())
	}

	cat, err := catAPIClient.GetImage(ctx, "byQhFO7iV")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cat.Info())

	if err := catAPIClient.CreateVote(ctx, list[1].ID, true); err != nil {
		log.Fatal(err)
	}

	myVotes, err := catAPIClient.GetVotes(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, vote := range myVotes {
		fmt.Println(vote.Info())
	}

	delVote, err := catAPIClient.DeleteVote(ctx, "932156")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(delVote.Message)
}
