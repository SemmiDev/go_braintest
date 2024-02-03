package go_braintest

import (
	"github.com/google/uuid"
	"time"
)

type Exam struct {
	ID        uuid.UUID `json:"id"`
	CreatorID uuid.UUID `json:"creator_id"`

	Title       string `json:"title"`
	Description string `json:"description"`

	Questions         []*Question `json:"questions"`
	TotalPoint        uint64      `json:"total_point"`
	MinPointForPass   uint64      `json:"min_point_for_pass"`
	RetakeAttempts    uint8       `json:"retake_attempts"`
	IsRandomQuestions bool        `json:"is_random_questions"`

	Participants []*Participant `json:"participants"`

	StartFrom         int64 `json:"start_from"`
	EndAt             int64 `json:"end_at"`
	Duration          int64 `json:"duration"` // in minutes
	IsTimePerQuestion bool  `json:"is_time_per_question"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CalculateTotalPoint menghitung total poin ujian berdasarkan poin setiap pertanyaan.
func (exam *Exam) CalculateTotalPoint() {
	totalPoint := uint64(0)

	for _, question := range exam.Questions {
		totalPoint += question.Point
	}

	exam.TotalPoint = totalPoint
}

// IsPassingScore memeriksa apakah skor yang diberikan melebihi batas kelulusan ujian.
func (exam *Exam) IsPassingScore(score uint64) bool {
	return score >= exam.MinPointForPass
}

// IsOngoing memeriksa apakah ujian masih berlangsung pada saat pemanggilan fungsi.
func (exam *Exam) IsOngoing() bool {
	currentTime := time.Now().Unix()
	return currentTime >= exam.StartFrom && currentTime <= exam.EndAt
}

// RemainingTime mengembalikan waktu tersisa dalam detik hingga berakhirnya ujian.
func (exam *Exam) RemainingTime() int64 {
	currentTime := time.Now().Unix()
	if currentTime >= exam.StartFrom && currentTime <= exam.EndAt {
		return exam.EndAt - currentTime
	}
	return 0
}

// IsExpired memeriksa apakah ujian telah berakhir pada saat pemanggilan fungsi.
func (exam *Exam) IsExpired() bool {
	currentTime := time.Now().Unix()
	return currentTime > exam.EndAt
}

// AddQuestion menambahkan pertanyaan ke dalam daftar pertanyaan ujian.
func (exam *Exam) AddQuestion(question *Question) {
	exam.Questions = append(exam.Questions, question)
	exam.CalculateTotalPoint()
}

// RemoveQuestion menghapus pertanyaan dari daftar pertanyaan ujian berdasarkan ID pertanyaan.
func (exam *Exam) RemoveQuestion(questionID uuid.UUID) {
	for i, question := range exam.Questions {
		if question.ID == questionID {
			exam.Questions = append(exam.Questions[:i], exam.Questions[i+1:]...)
			exam.CalculateTotalPoint()
			break
		}
	}
}

// UpdateQuestion mengganti pertanyaan lama dengan pertanyaan baru berdasarkan ID pertanyaan.
func (exam *Exam) UpdateQuestion(questionID uuid.UUID, newQuestion *Question) {
	for i, question := range exam.Questions {
		if question.ID == questionID {
			exam.Questions[i] = newQuestion
			exam.CalculateTotalPoint()
			break
		}
	}
}

// AddParticipant menambahkan peserta ke dalam daftar peserta ujian.
func (exam *Exam) AddParticipant(participant *Participant) {
	exam.Participants = append(exam.Participants, participant)
}

// RemoveParticipant menghapus peserta dari daftar peserta ujian berdasarkan ID peserta.
func (exam *Exam) RemoveParticipant(participantID uuid.UUID) {
	for i, participant := range exam.Participants {
		if participant.ID == participantID {
			exam.Participants = append(exam.Participants[:i], exam.Participants[i+1:]...)
			break
		}
	}
}

// UpdateParticipant mengganti data peserta lama dengan data peserta baru berdasarkan ID peserta.
func (exam *Exam) UpdateParticipant(participantID uuid.UUID, newParticipant *Participant) {
	for i, participant := range exam.Participants {
		if participant.ID == participantID {
			exam.Participants[i] = newParticipant
			break
		}
	}
}
