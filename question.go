package go_braintest

import (
	"github.com/google/uuid"
	"time"
)

type Question struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	Point     uint64    `json:"point"`
	Options   []*Option `json:"options"`
	TimeLimit int64     `json:"time_limit"` // in seconds
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

// AddOption menambahkan opsi baru ke dalam pertanyaan.
func (q *Question) AddOption(option *Option) {
	q.Options = append(q.Options, option)
}

// RemoveOption menghapus opsi dari pertanyaan berdasarkan ID opsi.
func (q *Question) RemoveOption(optionID uuid.UUID) {
	var updatedOptions []*Option
	for _, option := range q.Options {
		if option.ID != optionID {
			updatedOptions = append(updatedOptions, option)
		}
	}
	q.Options = updatedOptions
}

// UpdateTimeLimit memperbarui batas waktu pertanyaan.
func (q *Question) UpdateTimeLimit(newTimeLimit int64) {
	q.TimeLimit = newTimeLimit
	q.UpdatedAt = time.Now().Unix()
}

// GetCorrectOptions mengembalikan daftar opsi yang benar dari pertanyaan.
func (q *Question) GetCorrectOptions() []*Option {
	var correctOptions []*Option
	for _, option := range q.Options {
		if option.IsCorrect {
			correctOptions = append(correctOptions, option)
		}
	}
	return correctOptions
}

// IsCorrectAnswer memeriksa apakah jawaban yang diberikan benar berdasarkan ID opsi.
func (q *Question) IsCorrectAnswer(answerID uuid.UUID) bool {
	for _, option := range q.Options {
		if option.ID == answerID && option.IsCorrect {
			return true
		}
	}
	return false
}
