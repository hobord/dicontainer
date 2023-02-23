package main

import (
	"github.com/hobord/gdic"
	"github.com/hobord/gdic/example/service/service2"
)

func main() {
	// Resolve the service2.Service2 interface.
	// the gdic will call the factory function registered for the service2.Service2 interface.
	// and service2.Service2 factory function will call other factory functions to resolve the dependencies.

	// with 'gdic.Defalult' You will get the default instance of the interface,
	// but you can also get a named instance of the interface by passing the name of the instance.
	// but then you will need to register the factory function with the same name.
	srv2, err := gdic.Resolve[service2.Service2](gdic.Defalult)
	if err != nil {
		panic(err)
	}

	srv2.DoSomething()
}
