package service_test

import (
	"essemfly/go_base_app/internal/persistence/database"
	"essemfly/go_base_app/internal/service"
	"testing"

	"gotest.tools/v3/assert"
)

func TestMyService(t *testing.T) {
	mockDb := &database.MockDatabase{}

	myService := service.NewMyService(mockDb)

	result := myService.SomeMethod()
	assert.Equal(t, "hello query", result)
}
