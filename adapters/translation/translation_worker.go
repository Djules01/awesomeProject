package translationadapter

import (
	"context"
	"fmt"
	"time"
)

type TranslationJob struct {
	TodoID int
	Titre  string
}

type TranslationService struct {
	Queue chan TranslationJob
}

func NewTranslationService(bufferSize int) *TranslationService {
	return &TranslationService{
		Queue: make(chan TranslationJob, bufferSize),
	}
}

func (s *TranslationService) AddJob(job TranslationJob) bool {
	select {
	case s.Queue <- job:
		fmt.Println("Job ajoute dans la queue :", job.TodoID)
		return true
	default:
		fmt.Println("Queue pleine, job refuse :", job.TodoID)
		return false
	}
}

func (s *TranslationService) StartWorker() {
	go func() {
		for job := range s.Queue {
			s.processJob(job)
		}
	}()
}

func (s *TranslationService) processJob(job TranslationJob) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("Debut traduction fictive todo :", job.TodoID)

	select {
	case <-time.After(15 * time.Second):
		fmt.Println("Traduction terminee todo :", job.TodoID)
	case <-ctx.Done():
		fmt.Println("Timeout traduction todo :", job.TodoID)
	}
}
