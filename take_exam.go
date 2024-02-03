package go_braintest

import (
	"github.com/google/uuid"
	"time"
)

type TakeExam struct {
	ID         uuid.UUID `json:"id"`
	FinalPoint uint64    `json:"final_point"`
	StartAt    int64     `json:"start_at"`
	FinishAt   int64     `json:"finish_at"`
	Duration   int64     `json:"duration"`
	Note       string    `json:"note"`
}

// IsInProgress memeriksa apakah ujian masih berlangsung pada saat pemanggilan fungsi.
func (te *TakeExam) IsInProgress() bool {
	currentTime := time.Now().Unix()
	return currentTime >= te.StartAt && currentTime <= te.FinishAt
}

// IsExpired memeriksa apakah ujian sudah berakhir pada saat pemanggilan fungsi.
func (te *TakeExam) IsExpired() bool {
	currentTime := time.Now().Unix()
	return currentTime > te.FinishAt
}

// RemainingTime mengembalikan waktu tersisa dalam detik hingga berakhirnya ujian.
func (te *TakeExam) RemainingTime() int64 {
	currentTime := time.Now().Unix()
	if currentTime >= te.StartAt && currentTime <= te.FinishAt {
		return te.FinishAt - currentTime
	}
	return 0
}

// UpdateFinalPoint memperbarui nilai akhir peserta ujian.
func (te *TakeExam) UpdateFinalPoint(newFinalPoint uint64) {
	te.FinalPoint = newFinalPoint
}

// ExtendDuration memperpanjang durasi ujian berdasarkan waktu tambahan yang diberikan dalam detik.
func (te *TakeExam) ExtendDuration(additionalTime int64) {
	te.FinishAt += additionalTime
	te.Duration += additionalTime
}

// AddNote menambahkan catatan untuk ujian.
func (te *TakeExam) AddNote(note string) {
	te.Note += "\n" + note
}
