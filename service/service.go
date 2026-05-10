package service

import (
	"context"
	"go-project/repository"
)

type TestService interface {
	GetMessage() string
	SaveDBTest(ctx context.Context, body string) (repository.DBTestRecord, error)
}

type testService struct {
	repo repository.TestRepository
}

func NewTestService(repo repository.TestRepository) TestService {
	return &testService{repo: repo}
}

func (s *testService) GetMessage() string {
	return s.repo.GetHello()
}

func (s *testService) SaveDBTest(ctx context.Context, body string) (repository.DBTestRecord, error) {
	record, err := s.repo.CreateDBTest(ctx, body)
	if err != nil {
		return repository.DBTestRecord{}, err
	}

	return s.repo.GetDBTestByID(ctx, record.ID)
}
