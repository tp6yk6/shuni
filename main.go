// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	//"strconv"
	"math/rand" //亂數
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New("7f429e5e0a1c48de9ffacb668eda4b56", "fyHN9FqU0g2pOfX3d1jG0ZHxdU75biicfisoJTGBpuW/WwWghm1DwipvW9hG4fEo5GJGSGXKAsccP1wsqek+LEf08R0U5qIW1q7sK8DuYzME6erVbzOXmNOTLF+X2FvXJh6EXbVLpEt/RzXcLVP1qwdB04t89/1O/w1cDnyilFU=")
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				reply := message.Text
				choose :=[] string {"幹不要罵髒話啦", "衝三毀", "幹拎" } 
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				if message.Text == "幹" {
					reply = choose[r.Intn(len(choose))]
				}
				if _, err = bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage(reply)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

