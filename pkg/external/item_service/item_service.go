package item_service

import "fmt"

// Service Supposedly imported gRPC interface
type Service struct {
	FictionalDatabase map[string]*Item
}

// get retrieves an item using its name as id
func (s *Service) Get(name string) (*Item, error) {
	item, exists := s.FictionalDatabase[name]
	if !exists {
		return nil, fmt.Errorf("[%s] not found", name)
	}
	return item, nil
}
