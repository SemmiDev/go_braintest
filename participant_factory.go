package go_braintest

import "github.com/google/uuid"

type ParticipantBuilder struct {
	participant Participant
}

type ParticipantOption func(*ParticipantBuilder)

func NewParticipantBuilder(email string) *ParticipantBuilder {
	return &ParticipantBuilder{
		participant: Participant{
			ID:           uuid.New(),
			Email:        email,
			HighestPoint: 0, // Default highest point
		},
	}
}

func WithParticipantHighestPoint(highestPoint uint64) ParticipantOption {
	return func(pb *ParticipantBuilder) {
		pb.participant.HighestPoint = highestPoint
	}
}

func WithParticipantTakeExams(takeExams []*TakeExam) ParticipantOption {
	return func(pb *ParticipantBuilder) {
		pb.participant.TakeExams = takeExams
	}
}

func WithParticipantAveragePoint(averagePoint float64) ParticipantOption {
	return func(pb *ParticipantBuilder) {
		pb.participant.AveragePoint = averagePoint
	}
}

func (pb *ParticipantBuilder) Build(options ...ParticipantOption) Participant {
	for _, option := range options {
		option(pb)
	}
	return pb.participant
}
