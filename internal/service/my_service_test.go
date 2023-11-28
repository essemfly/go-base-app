package service

import (
	"essemfly/go_base_app/internal/db"
	"testing"

	"gotest.tools/v3/assert"
)

func TestMyService(t *testing.T) {
	mockDb := &db.MockDatabase{}

	myService := NewMyService(mockDb)

	result := myService.SomeMethod()
	assert.Equal(t, "hello query", result)
}
