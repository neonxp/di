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
		serviceA, err := di.GetByType[ServiceA]() // <- Get dependency from container by type
		if err != nil {
			return nil, err
		}

		return &ServiceB{
			ServiceA: serviceA[0],
		}, nil
	})

	// Do work...
	service, err := di.Get[ServiceB]("serviceB") // <- Get instantinated service B
	if err != nil {
		panic(err)
	}
	service.DoStuff() // Output: Hello, world!
}

func ExampleGet_interface() {
	di.Register("worker1", func() (*Worker1, error) {
		return &Worker1{}, nil
	})
	di.Register("worker2", func() (*Worker2, error) {
		return &Worker2{}, nil
	})
	workers, err := di.GetByInterface[Worker]()
	if err != nil {
		panic(err)
	}
	for _, w := range workers {
		w.Do()
	}
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

type Worker interface {
	Do()
}

type Worker1 struct{}

func (w *Worker1) Do() {
	fmt.Println("Worker 1 says hello")
}

type Worker2 struct{}

func (w *Worker2) Do() {
	fmt.Println("Worker 2 says hello")
}
