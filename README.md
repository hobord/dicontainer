# Dependecy injection container

This is a simple dependency injection container for GO using generics.

You should register Your factories in the container and then the container use them to resolve Your dependencies.

```go
func init() {
	// Register the factory function for the dal.Reader interface.
	// The factory function will be called once when the dal.Reader interface is resolved.
	err := gdic.AddFactory(gdic.Defalult, func() (Reader, error) {
        	// if reader needs some dependencies, you can resolve them here.
		dependency, err := gdic.Resolve[dep.Other](gdic.Defalult)
		if err != nil {
		    return nil, err
		}

		return NewReader(dependency), nil
	})

	if err != nil {
		panic(err)
	}
}
```

Then You can use the container to get instance in your code:

```go

func main() {
    // Resolve the dal.Reader interface.
    reader, err := gdic.Resolve[dal.Reader](gdic.Defalult)
    if err != nil {
        panic(err)
    }

    // Use the dal.Reader instance.
    reader.Read()
}
```

You have option to store multiple instances with same interface using instance diferent name when You register Your factory. 
The default InstanceName is `gdic.Defalult`.
But You can use any string as instance name. (define as const)

```go

const (
    CustomInstanceName = "customInstanceName"
)

func init() {
    // Register the factory function for the dal.Reader interface.
    // The factory function will be called once when the dal.Reader interface is resolved.
    err := gdic.AddFactory(CustomInstanceName, func() (Reader, error) {
        // if reader needs some dependencies, you can resolve them here.
        dependency, err := gdic.Resolve[dep.Other](gdic.Defalult)
        if err != nil {
        	return nil, err
        }

        return NewReader(dependency), nil
    })

    if err != nil {
        panic(err)
    }
}
```

Then You can use the container to get instance in your code:
```go

func main() {
    // Resolve the dal.Reader interface with custom instance name.
    reader, err := gdic.Resolve[dal.Reader](CustomInstanceName)
    if err != nil {
        panic(err)
    }

    // Use the dal.Reader instance.
    reader.Read()
}

```

Take a look at the example folder for more details.
