package go_braintest

import "github.com/google/uuid"

type Participant struct {
	ID           uuid.UUID   `json:"id"`
	Email        string      `json:"email"`
	HighestPoint uint64      `json:"highest_point"`
	AveragePoint float64     `json:"average_point"`
	TakeExams    []*TakeExam `json:"take_exams"`
}

// AddTakeExam menambahkan hasil ujian ke daftar ujian peserta.
func (p *Participant) AddTakeExam(exam *TakeExam) {
	p.TakeExams = append(p.TakeExams, exam)
}

// GetLatestTakeExam mengembalikan ujian terakhir yang diambil oleh peserta.
func (p *Participant) GetLatestTakeExam() *TakeExam {
	if len(p.TakeExams) > 0 {
		return p.TakeExams[len(p.TakeExams)-1]
	}
	return nil
}

// GetTotalExams mengembalikan jumlah total ujian yang diambil oleh peserta.
func (p *Participant) GetTotalExams() int {
	return len(p.TakeExams)
}

// UpdateHighestPoint memperbarui nilai tertinggi peserta jika nilai baru lebih tinggi.
func (p *Participant) UpdateHighestPoint(newPoint uint64) {
	if newPoint > p.HighestPoint {
		p.HighestPoint = newPoint
	}
}

// CalculateAverageScore menghitung rata-rata nilai ujian yang diambil oleh peserta.
func (p *Participant) UpdateAverageScore() {
	var total uint64
	for _, exam := range p.TakeExams {
		total += exam.FinalPoint
	}
	p.AveragePoint = float64(total) / float64(len(p.TakeExams))
}
