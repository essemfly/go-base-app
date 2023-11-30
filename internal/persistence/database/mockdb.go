package database

type MockDatabase struct{}

func (m *MockDatabase) QuerySomething() string {
	return "hello query"
}
