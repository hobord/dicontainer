package main

import (
	"sync"
	"testing"

	"github.com/hobord/gdic"
	"github.com/hobord/gdic/example/service/service2"
)

func Test_race(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			srv2, err := gdic.Resolve[service2.Service2](gdic.Default)
			if err != nil {
				panic(err)
			}

			srv2.DoSomething()
		}()
	}

	wg.Wait()
}
