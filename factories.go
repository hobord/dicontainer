package gdic

// difactory is a function that builds the object instance
// you can pass options to the factory function
type difactory func(opts ...interface{}) (interface{}, error)

// AddFactory adds factory to the container
// if factory is exist then return error
func AddFactory[T any](name InstanceName, factory func(opts ...interface{}) (T, error)) error {
	// get type of the interface
	itype := GetType[T]()

	// check if factory is nil
	if factory == nil {
		return ErrNilFactory
	}

	// lock container for manipulating the factories map
	store.Lck.Lock()
	defer store.Lck.Unlock()

	// if type is not exist in the container then create the map for it
	if _, ok := store.factories[itype]; !ok {
		store.factories[itype] = make(map[InstanceName]difactory)
	}

	// if factory is exist then return error
	if _, ok := store.factories[itype][name]; ok {
		return ErrFactoryExist
	}

	// add factory to the container
	store.factories[itype][name] = func(opts ...interface{}) (interface{}, error) {
		return factory(opts...)
	}

	return nil
}

// ReplaceFactory replaces factory in the container
// if factory is not exist then add it
func ReplaceFactory[T any](name InstanceName, factory func(opts ...interface{}) (T, error)) error {
	// get type of the interface
	itype := GetType[T]()

	// check if factory is nil
	if factory == nil {
		return ErrNilFactory
	}

	// lock container for manipulating the factories map
	store.Lck.Lock()
	defer store.Lck.Unlock()

	// if type is not exist in the container then create the map for it
	if _, ok := store.factories[itype]; !ok {
		store.factories[itype] = make(map[InstanceName]difactory)
	}

	// add factory to the container
	store.factories[itype][name] = func(opts ...interface{}) (interface{}, error) {
		return factory(opts...)
	}

	return nil
}

// DeleteFactory deletes factory from the container
func DeleteFactory[T any](name InstanceName) {
	// get type of the interface
	itype := GetType[T]()

	// lock container for manipulating the factories map
	store.Lck.Lock()
	defer store.Lck.Unlock()

	// if type is not exist in the container then create the map for it
	if _, ok := store.factories[itype]; !ok {
		return
	}

	// delete factory from the container
	delete(store.factories[itype], name)
}

// IsFactoryExist checks if factory is exist in the container
func IsFactoryExist[T any](name InstanceName) bool {
	// get type of the interface
	itype := GetType[T]()

	store.Lck.RLock()
	defer store.Lck.RUnlock()

	// check if type is exist in the container
	if _, ok := store.factories[itype]; !ok {
		return false
	}

	// check if factory is exist
	_, ok := store.factories[itype][name]

	return ok
}

// ConvertFactoryOptions converts options to the factory options
func ConvertFactoryOptions[T any](opts []interface{}) []T {
	fopts := make([]T, len(opts))

	for i, opt := range opts {
		fopts[i] = opt.(T)
	}

	return fopts
}
