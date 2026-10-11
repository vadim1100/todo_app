package core_kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client *kgo.Client
}

func NewProducer(brokers []string, topic string) *Producer {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.DefaultProduceTopic(topic),
		kgo.RequiredAcks(kgo.AllISRAcks()),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to create producer: %v", err))
	}

	return &Producer{client: client}
}

func (p *Producer) Produce(ctx context.Context, key string, event any) error {
	b, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshal event: %w", err)
	}

	record := &kgo.Record{
		Key:   []byte(key),
		Value: b,
	}

	if err := p.client.ProduceSync(ctx, record).FirstErr(); err != nil {
		return fmt.Errorf("produce: %w", err)
	}

	return nil
}

func (p *Producer) Close() {
	p.client.Close()
}
