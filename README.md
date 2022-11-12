# thecatsapi SDK
___


- Use the API Key that was emailed to you when you signed up to authenticate every request you send to the API.
- Without it you only have access to a tiny amount of data in a random order, and cannot  vote.

If you don’t have an API Key, just head over to 'https://thecatapi.com/#pricing' and get one for free



---

### `List` method GET:
#### Search & Itterate through all public images.

```go
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
}
```

---
### `getImage` method GET
#### Gets the image matching the image_id parameter passed.
```go
const APIKey = "" // your api key
const BaseUrl = "https://api.thecatapi.com/v1"
const UserID = "User-123" // your name to vote

func main() {
    catAPIClient := catsapi.NewClient(BaseUrl, APIKey, UserID)

    ctx := context.Background()

    cat, err := catAPIClient.GetImage(ctx, "byQhFO7iV")
        if err != nil {
        log.Fatal(err)
    }

    fmt.Println(cat.Info())
}
```

---
### `CreateVote` method POST
#### Vote an Image Up or Down. Contains the “image_id” you want to Vote on, and “value”:true to Up Vote, or “value”:false to Down Vote.
```go
const APIKey = "" // your api key
const BaseUrl = "https://api.thecatapi.com/v1"
const UserID = "User-123" // your name to vote

func main() {
    catAPIClient := catsapi.NewClient(BaseUrl, APIKey, UserID)
    
    ctx := context.Background()
    
    if err := catAPIClient.CreateVote(ctx, list[1].ID, true); err != nil {
        log.Fatal(err)
    }

    for _, vote := range myVotes {
        fmt.Println(vote.Info())
    }
}
```


---
### `DeleteVote` method DELETE
#### Delete a Vote from your Account
```go
const APIKey = "" // your api key
const BaseUrl = "https://api.thecatapi.com/v1"
const UserID = "User-123" // your name to vote

func main() {
    catAPIClient := catsapi.NewClient(BaseUrl, APIKey, UserID)
    
    ctx := context.Background()
    
    delVote, err := catAPIClient.DeleteVote(ctx, "932156")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(delVote.Message)
}
```