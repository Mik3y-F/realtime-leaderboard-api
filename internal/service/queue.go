package service

type ScoreUpdateEvent struct {
	PlayerID  string
	NewScore  int
	TimeStamp int64 // Optional: To record when the event took place
}

type RabbitMQService interface {
	// PublishScoreUpdate publishes a score update event to a Message Queue exchange.
	PublishScoreUpdate(event ScoreUpdateEvent) error

	// ConsumeScoreUpdates starts consuming score update messages from a queue.
	// The provided callback is called for each consumed message.
	ConsumeScoreUpdates(callback func(event ScoreUpdateEvent) error) error
}
