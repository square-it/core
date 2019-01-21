package property

import "github.com/project-flogo/core/data"

func init() {
	SetDefaultManager(NewManager(make(Properties)))
}

var defaultManager *Manager

func SetDefaultManager(manager *Manager) {
	defaultManager = manager
}

func DefaultManager() *Manager {
	return defaultManager
}

func NewManager(properties Properties) *Manager {

	manager := &Manager{properties: properties}
	return manager
}

type Manager struct {
	properties Properties
}

func (m *Manager) GetProperty(name string) (interface{}, bool) {
	var value interface{}

	property, exists := m.properties[name]
	if exists {
		value = property.Value
	}
	return value, exists
}

func (m *Manager) Finalize(processors ...PostProcessor) error {

	for _, processor := range processors {
		processor(m.properties)
	}

	return nil
}

type PostProcessor func(properties Properties) error

type Property struct {
	Value    interface{}
	DataType data.Type
}

type Properties map[string]*Property
