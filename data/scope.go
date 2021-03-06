package data

// Scope is a set of attributes that are accessible
type Scope interface {
	// GetValue gets the specified value
	GetValue(name string) (value interface{}, exists bool)

	// SetValue sets the specified  value
	SetValue(name string, value interface{}) error
}

// SimpleScope is a basic implementation of a scope
type SimpleScope struct {
	parentScope Scope
	values      map[string]interface{}
}

// NewSimpleScope creates a new SimpleScope
func NewSimpleScope(values map[string]interface{}, parentScope Scope) Scope {

	scope := &SimpleScope{
		parentScope: parentScope,
		values:      make(map[string]interface{}),
	}

	for name, value := range values {
		scope.values[name] = value
	}

	return scope
}

func (s *SimpleScope) GetValue(name string) (value interface{}, exists bool) {
	value, found := s.values[name]

	if found {
		return value, true
	}

	if s.parentScope != nil {
		return s.parentScope.GetValue(name)
	}

	return nil, false
}

func (s *SimpleScope) SetValue(name string, value interface{}) error {

	s.values[name] = value
	return nil
	//attr, found := s.values[name]
	//
	//if found {
	//	attr.SetValue(value)
	//	return nil
	//}

	//return errors.New("attribute not in scope")
}
