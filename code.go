package queue

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func (q *Queue) handle(wr http.ResponseWriter, req *http.Request) error {
	channel := strings.TrimPrefix(req.URL.Path, "/queue/")
	email := strings.TrimSpace(req.PostFormValue("email"))
	msg := Message{}
	msg.Payload = []byte(email)
	msg.Timestamp = time.Now()
	msg.IPAddress = req.RemoteAddr
	msg.Referer = req.Header.Get("Referer")
	b, _ := json.Marshal(msg)
	go func() {
		q.sendToChannel(channel, b)
	}()
	wr.WriteHeader(http.StatusCreated)
	wr.Header().Add("Content-Type", "application/json")
	wr.Write(okResponse)
	return nil
}
