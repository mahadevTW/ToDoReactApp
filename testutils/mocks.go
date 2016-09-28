package testutils

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockToDoRepository struct {
	mock.Mock
}

func (m *MockToDoRepository) Delete(db *sql.DB, id int) error {
	args := m.Called(db, id)
	return args.Error(0)
}
