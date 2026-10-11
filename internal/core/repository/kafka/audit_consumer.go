package core_kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
	core_logger "github.com/vadim1100/todo_app/internal/core/logger"
	"go.uber.org/zap"
)

type AuditInserter interface {
	Insert(
		ctx context.Context,
		event string,
		userID int,
		taskID *int,
		payload []byte,
	) error
}

type AuditConsumer struct {
	client *kgo.Client
	audit  AuditInserter
	logger *core_logger.Logger
}

func NewAuditConsumer(
	brokers []string,
	topic string,
	groupID string,
	audit AuditInserter,
	logger *core_logger.Logger,
) *AuditConsumer {
	opts := []kgo.Opt{
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup(groupID),
		kgo.ConsumeTopics(topic),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
		kgo.DisableAutoCommit(),
		kgo.Balancers(kgo.CooperativeStickyBalancer()),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		panic(fmt.Sprintf("failed to create kafka client: %v", err))
	}

	return &AuditConsumer{
		client: client,
		audit:  audit,
		logger: logger,
	}
}

func (c *AuditConsumer) Run(ctx context.Context) {
	defer c.client.Close()

	for {
		if ctx.Err() != nil {
			c.logger.Warn("cancelled context")
			return
		}

		fetches := c.client.PollFetches(ctx)

		if errs := fetches.Errors(); len(errs) > 0 {
			for _, e := range errs {
				c.logger.Error("kafka fetch failed", zap.Error(e.Err))
			}
			continue
		}

		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()

			var event TaskEvent

			if err := json.Unmarshal(record.Value, &event); err != nil {
				c.logger.Error("bad event", zap.Error(err))
				_ = c.client.CommitRecords(ctx, record)
				continue
			}

			var taskID *int

			if event.TaskID != 0 {
				taskID = &event.TaskID
			}

			if err := c.audit.Insert(ctx, event.Event, event.UserID, taskID, record.Value); err != nil {
				c.logger.Error("audit insert failed", zap.Error(err))
				continue
			}

			if err := c.client.CommitRecords(ctx, record); err != nil {
				c.logger.Error("commit failed", zap.Error(err))
			}
		}
	}
}
