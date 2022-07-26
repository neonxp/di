package di

import (
	"fmt"
	"sync"
)

var ( // singletones
	services sync.Map
	cache    sync.Map
)

func init() {
	services = sync.Map{}
	cache = sync.Map{}
}

// Register service in di
func Register[T any](id string, constructor func() (*T, error)) {
	services.Store(id, constructor)
}

// Get services by type
func GetByType[T any]() ([]*T, error) {
	var err error
	result := []*T{}
	services.Range(func(id, constructor any) bool {
		if constructor, ok := constructor.(func() (*T, error)); ok {
			if instance, ok := cache.Load(id); ok {
				if instance, ok := instance.(*T); ok {
					result = append(result, instance)
				}
				return true
			}
			instance, instErr := constructor()
			if instErr != nil {
				err = instErr
				return false
			}
			cache.Store(id, instance)
			result = append(result, instance)
		}
		return true
	})
	return result, err
}

// Get services by interface
func GetByInterface[Interface any]() ([]Interface, error) {
	var err error
	result := []Interface{}
	services.Range(func(id, constructor any) bool {
		if constructor, ok := constructor.(func() (Interface, error)); ok {
			if instance, ok := cache.Load(id); ok {
				if instance, ok := instance.(Interface); ok {
					result = append(result, instance)
				}
				return true
			}
			instance, instErr := constructor()
			if instErr != nil {
				err = instErr
				return false
			}
			cache.Store(id, instance)
			result = append(result, instance)
		}
		return true
	})
	return result, err
}

// Get service by id and type
func Get[T any](id string) (*T, error) {
	if instance, ok := cache.Load(id); ok {
		if instance, ok := instance.(*T); ok {
			return instance, nil
		}
		return nil, fmt.Errorf("invalid type for service %s (%t)", id, instance)
	}
	if constructor, ok := services.Load(id); ok {
		if constructor, ok := constructor.(func() (*T, error)); ok {
			instance, err := constructor()
			if err != nil {
				return nil, err
			}
			cache.Store(id, instance)
			return instance, nil
		}
		return nil, fmt.Errorf("invalid constructor")
	}
	return nil, fmt.Errorf("unknown service %s", id)
}
