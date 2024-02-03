package go_braintest

import (
	"github.com/google/uuid"
	"time"
)

type QuestionBuilder struct {
	question Question
}

type QuestionOption func(*QuestionBuilder)

func NewQuestionBuilder(text string, point uint64, options []*Option) *QuestionBuilder {
	return &QuestionBuilder{
		question: Question{
			ID:        uuid.New(),
			Text:      text,
			Point:     point,
			Options:   options,
			TimeLimit: 0, // Default time limit (in seconds)
			CreatedAt: time.Now().Unix(),
		},
	}
}

func WithQuestionTimeLimit(timeLimit int64) QuestionOption {
	return func(qb *QuestionBuilder) {
		qb.question.TimeLimit = timeLimit
	}
}

func WithQuestionUpdatedAt(updatedAt int64) QuestionOption {
	return func(qb *QuestionBuilder) {
		qb.question.UpdatedAt = updatedAt
	}
}

func (qb *QuestionBuilder) Build(options ...QuestionOption) Question {
	for _, option := range options {
		option(qb)
	}
	return qb.question
}
