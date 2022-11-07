package connect

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"math/rand"
	"sync"
	"time"
)

const broker = "broker.emqx.io" // MQTT Broker 连接地址
const port = 1883
const username = "emqx"
const password = "******"
const ClientPrefix = "aide-family-"

type (
	// MqttClientConfig mqtt client
	MqttClientConfig struct {
		Broker    string
		Port      int
		Username  string
		Password  string
		ClientId  string
		Keepalive time.Duration
	}

	// MqttClientConfigOption mqtt client config option
	MqttClientConfigOption func(*MqttClientConfig)

	MQTTClient struct {
		mqtt.Client
		Logger   *log.Logger
		lock     sync.RWMutex
		topicSet map[string]TopicConfig
	}

	TopicConfig struct {
		Qos     byte
		Handler mqtt.MessageHandler
	}

	PublishMessage struct {
		Topic    string // 主题
		Payload  string // 消息内容
		Qos      byte   // 服务质量
		Retained bool   // 是否保留消息
	}

	SubscribeMessage struct {
		Topic    string // 主题
		Qos      byte   // 服务质量
		Callback func(client mqtt.Client, msg mqtt.Message)
	}
)

// NewMqttClientConfig new mqtt client config
func NewMqttClientConfig(opts ...MqttClientConfigOption) *MqttClientConfig {
	config := &MqttClientConfig{}
	for _, opt := range opts {
		opt(config)
	}
	return config
}

// NewDefaultMqttClientConfig new default mqtt client config
func NewDefaultMqttClientConfig() *MqttClientConfig {
	return NewMqttClientConfig(
		WithBroker(broker),
		WithPort(port),
		WithUsername(username),
		WithPassword(password),
		WithKeepalive(30*time.Second),
		WithClientId(fmt.Sprintf("%s%d", ClientPrefix, rand.Intn(1000))),
	)
}

// newMqttClient create mqtt client
func newMqttClient(cfg ...*MqttClientConfig) mqtt.Client {
	var conf = NewDefaultMqttClientConfig()
	if len(cfg) > 0 {
		conf = cfg[0]
	}
	connectAddress := fmt.Sprintf("tcp://%s:%d", conf.Broker, conf.Port)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(connectAddress)
	opts.SetUsername(conf.Username)
	opts.SetPassword(conf.Password)
	opts.SetClientID(conf.ClientId)
	opts.SetKeepAlive(conf.Keepalive)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	// 如果连接失败，则终止程序
	if token.WaitTimeout(3*time.Second) && token.Error() != nil {
		log.Fatal(token.Error())
	}
	return client
}

// NewMQTTClient create mqtt client
func NewMQTTClient(cfg *MqttClientConfig, logger *log.Logger) *MQTTClient {
	cli := &MQTTClient{}
	cli.Client = newMqttClient(cfg)
	cli.Logger = logger
	cli.topicSet = make(map[string]TopicConfig)
	return cli
}

func NewTopicConfig(qos byte, Handler mqtt.MessageHandler) TopicConfig {
	return TopicConfig{
		Qos:     qos,
		Handler: Handler,
	}
}

// Publish  message
func (m *MQTTClient) Publish(msg PublishMessage) {
	if token := m.Client.Publish(msg.Topic, msg.Qos, msg.Retained, msg.Payload); token.Wait() && token.Error() != nil {
		m.Logger.Printf("publish failed, topic: %s, payload: %s\n", msg.Topic, msg.Payload)
	}
}

// Subscribe subscribe topic
func (m *MQTTClient) Subscribe() {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for topic, config := range m.topicSet {
		m.Logger.Printf("subscribe topic: %s\n", topic)
		m.Client.Subscribe(topic, config.Qos, config.Handler)
	}
}

// AppendTopic append topic
func (m *MQTTClient) AppendTopic(topic string, config TopicConfig) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.topicSet == nil {
		m.topicSet = make(map[string]TopicConfig)
	}
	m.topicSet[topic] = config
	m.Client.Subscribe(topic, config.Qos, config.Handler)
}

// RemoveTopic remove topic
func (m *MQTTClient) RemoveTopic(topic string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.Client.Unsubscribe(topic)
	delete(m.topicSet, topic)
}

// ClearTopicSet clear topic set
func (m *MQTTClient) ClearTopicSet() {
	m.lock.Lock()
	defer m.lock.Unlock()
	for topic := range m.topicSet {
		m.Client.Unsubscribe(topic)
	}
	m.topicSet = make(map[string]TopicConfig)
}

// Unsubscribe 关闭主题监听
func (m *MQTTClient) Unsubscribe(topic string) {
	if _, ok := m.topicSet[topic]; ok {
		m.Client.Unsubscribe(topic)
	}
}

// WithBroker set broker
func WithBroker(broker string) MqttClientConfigOption {
	return func(config *MqttClientConfig) {
		config.Broker = broker
	}
}

// WithPort set port
func WithPort(port int) MqttClientConfigOption {
	return func(config *MqttClientConfig) {
		config.Port = port
	}
}

// WithUsername set username
func WithUsername(username string) MqttClientConfigOption {
	return func(config *MqttClientConfig) {
		config.Username = username
	}
}

// WithPassword set password
func WithPassword(password string) MqttClientConfigOption {
	return func(config *MqttClientConfig) {
		config.Password = password
	}
}

// WithClientId set keepalive
func WithClientId(clientId string) MqttClientConfigOption {
	return func(config *MqttClientConfig) {
		config.ClientId = clientId
	}
}

// WithKeepalive set keepalive
func WithKeepalive(keepalive time.Duration) MqttClientConfigOption {
	return func(config *MqttClientConfig) {
		config.Keepalive = keepalive
	}
}
