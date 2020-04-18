package main

import (
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	gitlab "github.com/xanzy/go-gitlab"
)

func main() {
	http.HandleFunc("/webhook", handler)
	http.ListenAndServe(":8888", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		logrus.Error("Bad request. ", err)
		return
	}
	logrus.Info("Request Received: ", string(body))
	event, err := gitlab.ParseWebhook(gitlab.WebhookEventType(r), body)
	if err != nil {
		http.Error(w, "Can't parse webhook", http.StatusInternalServerError)
		logrus.Error("Can't parse webhook. ", err)
		return
	}
	logrus.Info("Success.")
	logrus.Info(event)
	w.WriteHeader(204)
}
