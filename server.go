package main

import (
  "net/http"
  "encoding/json"
  "fmt"
  "github.com/franela/goreq"
)

var BOT_AUTH = "" // Your Bot Auth Key.

// Struct prototype for response JSON.
type test_struct struct {
  Message struct {
    Text string
    Chat struct {
      Id int
    }
  }
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(w, "Welcome to a Golang Server!")
    fmt.Println("Touchdown! @ /")
  })

  http.HandleFunc("/BotPush",func(w http.ResponseWriter, request *http.Request) {
    fmt.Println("Touchdown! @ /BotPush")
    decoder := json.NewDecoder(request.Body)

  	var t test_struct
  	err := decoder.Decode(&t)

  	if err != nil {
  	   panic(err)
  	}

  	fmt.Println("Telegram User ID:", t.Message.Chat.Id)
    fmt.Println("Text:", t.Message.Text)


    type Item struct {
      Chat_id int `chat_id:""`
      Text string `text:""`
    }

    item := Item {
      Chat_id: t.Message.Chat.Id,
      Text: "You said: " + t.Message.Text,
    }

    res, err := goreq.Request{
      Method: "POST",
      Uri: "https://api.telegram.org/bot" + BOT_AUTH + "/sendMessage",
      QueryString: item,
      ShowDebug:   false,
    }.Do()

    fmt.Println("Response Sent!", res)
    fmt.Fprintf(w, "Thanks!")
  })
  panic(http.ListenAndServe(":8001", nil))
}
