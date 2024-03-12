package metric

import "github.com/bongerka/metricaler/internal/repository"

type Service struct {
	repo repository.MetricRepository
}

func NewService(repo repository.MetricRepository) *Service {
	return &Service{
		repo: repo,
	}
}
