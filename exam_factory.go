package go_braintest

import (
	"github.com/google/uuid"
	"time"
)

type ExamBuilder struct {
	exam Exam
}

type ExamOption func(*ExamBuilder)

func NewExamBuilder(title, description string) *ExamBuilder {
	return &ExamBuilder{
		exam: Exam{
			ID:          uuid.New(),
			Title:       title,
			Description: description,
		},
	}
}

func WithIsTimePerQuestion(isTimePerQuestion bool) ExamOption {
	return func(eb *ExamBuilder) {
		eb.exam.IsTimePerQuestion = isTimePerQuestion
	}
}

func WithQuestion(question *Question) ExamOption {
	return func(eb *ExamBuilder) {
		eb.exam.Questions = append(eb.exam.Questions, question)
		eb.exam.CalculateTotalPoint()
	}
}

func WithQuestions(questions []*Question) ExamOption {
	return func(eb *ExamBuilder) {
		eb.exam.Questions = questions
		eb.exam.CalculateTotalPoint()
	}
}

func WithRandomQuestions(isRandomQuestions bool) ExamOption {
	return func(eb *ExamBuilder) {
		eb.exam.IsRandomQuestions = isRandomQuestions
	}
}

func WithTotalPoint(totalPoint uint64) ExamOption {
	return func(eb *ExamBuilder) {
		eb.exam.TotalPoint = totalPoint
	}
}

func WithMinPointForPass(minPointForPass uint64) ExamOption {
	return func(eb *ExamBuilder) {
		eb.exam.MinPointForPass = minPointForPass
	}
}

func WithStartAndEndTime(startFrom, endAt time.Time) ExamOption {
	return func(eb *ExamBuilder) {
		eb.exam.StartFrom = startFrom.Unix()
		eb.exam.EndAt = endAt.Unix()

		eb.exam.Duration = endAt.Unix() - startFrom.Unix()
	}
}

func WithDuration(duration int64) ExamOption {
	return func(eb *ExamBuilder) {
		eb.exam.Duration = duration
	}
}

func WithRetakeAttempts(retakeAttempts uint8) ExamOption {
	return func(eb *ExamBuilder) {
		eb.exam.RetakeAttempts = retakeAttempts
	}
}

func (eb *ExamBuilder) Build(options ...ExamOption) Exam {
	for _, option := range options {
		option(eb)
	}

	eb.exam.CreatedAt = time.Now()
	eb.exam.UpdatedAt = time.Now()

	return eb.exam
}
