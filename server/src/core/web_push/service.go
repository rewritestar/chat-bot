package webpush

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"chat-bot/src/common-service/push"
	core_values "chat-bot/src/core/values"
	"chat-bot/src/core/web_push/domain"

	"github.com/SherClockHolmes/webpush-go"
)

type webPushService struct {
	svc push.PushService
}

func NewWebPush(svc push.PushService) WebPushService {
	return &webPushService{
		svc,
	}
}

func (s *webPushService) Push(reqData domain.RequestPush) {
	s.push(reqData)
}

func (s *webPushService) push(reqData domain.RequestPush) {
	reqMsg := domain.RequestMessage{
		Title: reqData.Title,
		Body:  reqData.Body,
		Url:   fmt.Sprintf("%s/room/%d", os.Getenv(core_values.EnvClientUrl), reqData.RoomID),
	}
	msgByte, err := json.Marshal(reqMsg)
	if err != nil {
		log.Println(err.Error())
		return
	}

	subList, err := s.svc.FindAllByUserID(reqData.UserID)
	if err != nil {
		return
	}

	var resp *http.Response
	for _, sub := range *subList {
		subInfo := &webpush.Subscription{
			Endpoint: sub.Endpoint,
			Keys: webpush.Keys{
				Auth:   sub.Auth,
				P256dh: sub.P256dh,
			},
		}

		resp, err = webpush.SendNotification(msgByte, subInfo, &webpush.Options{
			Subscriber:      os.Getenv(core_values.EnvAdminEmail),
			VAPIDPublicKey:  os.Getenv(core_values.EnvVapidPublicKey),
			VAPIDPrivateKey: os.Getenv(core_values.EnvVapidPrivateKey),
			TTL:             30,
		})
		if err != nil {
			log.Println(err.Error())
			continue
		}
	}

	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
}
