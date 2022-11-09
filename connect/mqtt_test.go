package connect

import (
	"github.com/aide-cloud/universal/alog"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"testing"
	"time"
)

func TestCreateMqttClient(t *testing.T) {
	cli := NewMQTTClient(NewDefaultMqttClientConfig(), alog.NewLogger())
	cli.AppendTopic(ClientPrefix, NewTopicConfig(0, func(client mqtt.Client, msg mqtt.Message) {
		t.Log(ClientPrefix, "收到消息", string(msg.Payload()))
	}))

	cli.AppendTopic("aide-family-1", NewTopicConfig(0, func(client mqtt.Client, msg mqtt.Message) {
		t.Log("aide-family-1", "收到消息", string(msg.Payload()))
	}))

	cli.Subscribe()
	time.Sleep(time.Second * 1)

	count := 0
	for {
		count++
		cli.Publish(PublishMessage{Topic: ClientPrefix, Qos: 0, Payload: "123"})
		time.Sleep(time.Second * 1)
		if count > 10 {
			cli.Publish(PublishMessage{Topic: "aide-family-1", Qos: 0, Payload: "123"})
			cli.RemoveTopic(ClientPrefix)
		}

		if count > 20 {
			cli.ClearTopicSet()
			t.Log("清空主题")
			break
		}
	}
}
