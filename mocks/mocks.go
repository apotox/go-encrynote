package mocks

import "github.com/stretchr/testify/mock"

type QueueDeleteNoteMock struct {
	mock.Mock
}

func (r *QueueDeleteNoteMock) Publish(id string) error {

	args := r.Called(id)
	return args.Error(0)
}

func (r *QueueDeleteNoteMock) Init() error {

	args := r.Called()
	return args.Error(0)
}
