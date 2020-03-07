package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	// BroadcastURL は BOT にてブロードキャスト配信をおこなうための API URL
	BroadcastURL = "https://api.line.me/v2/bot/message/broadcast"
)

var (
	// AccessToken は LINE Developers にて提供される Messaging API のアクセストークン
	AccessToken = os.Getenv("LINE_ACCESS_TOKEN")
)

// Message is struct which have message to post
type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// PostMessages は送信するメッセージ群を保持する構造体
type PostMessages struct {
	Messages []Message `json:"messages"`
}

// PostBroadcast send messages to LINE Channel
func PostBroadcast(postMessages PostMessages) {
	jsonByte, err := json.Marshal(postMessages)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	req, err := http.NewRequest(
		"POST",
		BroadcastURL,
		bytes.NewBuffer(jsonByte),
	)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+AccessToken)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		fmt.Println("Succeed to send.")
	} else {
		fmt.Println(response)
	}
	return
}
