package di_test

import (
	"fmt"

	"go.neonxp.dev/di"
)

func ExampleGet() {
	di.Register("serviceA", func() (*ServiceA, error) { // <- Register service A
		return &ServiceA{}, nil
	})
	di.Register("serviceB", func() (*ServiceB, error) { // <- Register service B, that depends from service A
		serviceA, err := di.Get[ServiceA]() // <- Get dependency from container by type
		if err != nil {
			return nil, err
		}

		return &ServiceB{
			ServiceA: serviceA[0],
		}, nil
	})

	// Do work...
	service, err := di.GetById[ServiceB]("serviceB") // <- Get instantinated service B
	if err != nil {
		panic(err)
	}
	service.DoStuff() // Output: Hello, world!
}

type ServiceA struct{}

func (d *ServiceA) DoStuff() {
	fmt.Println("Hello, world!")
}

type ServiceB struct {
	ServiceA *ServiceA
}

func (d *ServiceB) DoStuff() {
	d.ServiceA.DoStuff()
}
