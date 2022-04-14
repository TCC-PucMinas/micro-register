package service

import (
	"os"

	"github.com/TCC-PucMinas/micro-register/helpers"
	"github.com/nats-io/nats.go"
)

type Nats struct {
	Url      string `json:"url"`
	Encoding string `json:"encoding"`
	Message  []byte
	Nats     *nats.EncodedConn
}

func (n *Nats) Connect() error {
	n.Url = os.Getenv("NATS_SERVER")
	//n.Url = "localhost:4222"
	n.Encoding = nats.JSON_ENCODER
	nc, err := nats.Connect(n.Url)

	if err != nil {
		return err
	}
	c, er := nats.NewEncodedConn(nc, n.Encoding)

	if er != nil {
		return err
	}

	n.Nats = c
	return nil
}

func attemptsPublishMail(n *Nats, subject string, mailAlert *EmailCommunicate, retry int) error {
	err := n.Nats.Publish(subject, mailAlert)

	if err != nil && helpers.EqualRetry(err, retry) {
		return attemptsPublishMail(n, subject, mailAlert, retry+1)
	}

	return err
}

func (n *Nats) PublishAlertEmail(subject string, mailAlert *EmailCommunicate) error {
	return attemptsPublishMail(n, subject, mailAlert, 1)
}
