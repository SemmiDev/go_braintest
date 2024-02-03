package go_braintest

import "github.com/google/uuid"

type OptionBuilder struct {
	option Option
}

type OptionOption func(*OptionBuilder)

func NewOptionBuilder(text string, isCorrect bool) *OptionBuilder {
	return &OptionBuilder{
		option: Option{
			ID:        uuid.New(),
			Text:      text,
			IsCorrect: isCorrect,
		},
	}
}

func (ob *OptionBuilder) Build() Option {
	return ob.option
}
