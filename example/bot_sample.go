package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/eezhal92/faq"
)

func main() {
	q := faq.Question{
		Text:   "Hi",
		Answer: "Sup, What are you looking for?",
		Choices: []faq.Question{
			faq.Question{
				Text:   "Are you open?",
				Answer: "Yes",
			},
			faq.Question{
				Text:   "What do you have?",
				Answer: "Here's the items",
				Choices: []faq.Question{
					faq.Question{Text: "Fried Rice", Answer: "You will like it!"},
					faq.Question{Text: "Curry", Answer: "Thank you!!"},
				},
			},
		},
	}

	sessionCounter := 1
	driverContainer := make(map[string]*faq.Driver)

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		nextSession := strconv.Itoa(sessionCounter)
		sessionCounter += 1

		driver := faq.NewDriver(q, "back", "repeat", "reset")
		driverContainer[nextSession] = driver
		reply := driver.Boot()

		data := struct {
			Message faq.Reply `json:"message"`
			Session string    `json:"session"`
		}{
			Message: reply,
			Session: nextSession,
		}
		result, _ := json.Marshal(data)

		w.Write(result)
	})
	http.HandleFunc("/ask", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		session := r.URL.Query().Get("session")
		cmd := r.URL.Query().Get("cmd")

		driver, ok := driverContainer[session]

		if !ok {
			payload := struct {
				Message string `json:"message"`
			}{Message: "Session not found"}
			result, _ := json.Marshal(payload)
			w.Write(result)
			return
		}

		reply := driver.Ask(cmd)

		result, _ := json.Marshal(reply)
		w.Write(result)
	})

	address := "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Println(err.Error())
	}
}
