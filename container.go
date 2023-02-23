package gdic

import (
	"sync"
)

// difactory is a function that builds the object instance
type difactory func() (interface{}, error)

// InstanceName is a name of instance for store in the container
type InstanceName string

const (
	Defalult InstanceName = "default"
	// you can define more instance name in Your project
)

// store is a module level singleton container
var store container

// container is store instances and factories
type container struct {
	Lck sync.RWMutex

	// instances is storage for the created instances
	instances map[string]map[InstanceName]interface{}
	// factories is storage for the factories for creating instances into the container
	factories map[string]map[InstanceName]difactory
}

// create module level singleton container
func init() {
	store = container{
		instances: make(map[string]map[InstanceName]interface{}),
		factories: make(map[string]map[InstanceName]difactory),
	}
}
