package go_braintest

import (
	"github.com/google/uuid"
	"time"
)

type TakeExamBuilder struct {
	takeExam TakeExam
}

type TakeExamOption func(*TakeExamBuilder)

func NewTakeExamBuilder(finalPoint uint64) *TakeExamBuilder {
	return &TakeExamBuilder{
		takeExam: TakeExam{
			ID:         uuid.New(),
			FinalPoint: finalPoint,
			StartAt:    time.Now().Unix(),
		},
	}
}

func WithTakeExamFinishTime(finishAt time.Time) TakeExamOption {
	return func(teb *TakeExamBuilder) {
		teb.takeExam.FinishAt = finishAt.Unix()
	}
}

func WithTakeExamDuration(duration int64) TakeExamOption {
	return func(teb *TakeExamBuilder) {
		teb.takeExam.Duration = duration
	}
}

func WithTakeExamNote(note string) TakeExamOption {
	return func(teb *TakeExamBuilder) {
		teb.takeExam.Note = note
	}
}

func (teb *TakeExamBuilder) Build(options ...TakeExamOption) TakeExam {
	for _, option := range options {
		option(teb)
	}
	return teb.takeExam
}
