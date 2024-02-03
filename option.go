package go_braintest

import "github.com/google/uuid"

type Option struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	IsCorrect bool      `json:"is_correct"`
}

func (o *Option) ToggleCorrectStatus() {
	o.IsCorrect = !o.IsCorrect
}

// UpdateText memperbarui teks opsi.
func (o *Option) UpdateText(newText string) {
	o.Text = newText
}
