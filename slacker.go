package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"

	"golang.org/x/net/websocket"
)

type Slacker struct {
	Token string
	WS    *websocket.Conn
	ID    string
}

type slackLoginResponse struct {
	Ok    bool         `json:"ok"`
	Error string       `json:"error"`
	Url   string       `json:"url"`
	Self  responseSelf `json:"self"`
}

type responseSelf struct {
	Id string `json:"id"`
}

func NewSlacker(token string) Slacker {
	return Slacker{Token: token}
}

func (s *Slacker) Connect() error {
	resp, err := http.Get("https://slack.com/api/rtm.start?token=" + s.Token)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("API request failure: " + resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	var slackResponse slackLoginResponse
	err = json.Unmarshal(body, &slackResponse)
	if err != nil {
		return err
	}

	if !slackResponse.Ok {
		return errors.New("Error: " + slackResponse.Error)
	}

	wsurl := slackResponse.Url
	id := slackResponse.Self.Id

	ws, err := websocket.Dial(wsurl, "", "https://api.slack.com/")
	if err != nil {
		log.Fatal(err)
	}

	s.ID = id
	s.WS = ws

	return nil
}

func (s *Slacker) ListChannels() (string, error) {
	resp, err := http.Get("https://slack.com/api/channels.list?token=" + s.Token + "&pretty=1")
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New("API request failure: " + resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(body), nil
}

type SlackMessage struct {
	Id      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func (s *Slacker) GetMessage() (SlackMessage, error) {
	var sm SlackMessage
	err := websocket.JSON.Receive(s.WS, &sm)
	return sm, err
}

var counter uint64 = 0

func (s *Slacker) SendMessage(msg, channel string) error {
	sm := SlackMessage{
		Id:      atomic.AddUint64(&counter, 1),
		Type:    "message",
		Channel: channel,
		Text:    msg,
	}
	return websocket.JSON.Send(s.WS, sm)
}
