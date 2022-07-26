# Go Dependency Inject Container

Main repo: [https://gitrepo.ru/neonxp/di](https://gitrepo.ru/neonxp/di). Github is only mirror.

Simple dependecy inject container with generics.

Use for your own risk!

## Usage

### Register dependency

```go
di.Register("service id", func () (*Service, error) { /* construct service */ })
```

### Get dependency

Get dependencies by type:

```go
services, err := di.GetByType[Service]()
```

Get dependencies by type and id:
```go
service, err := di.Get[Service]("service id")
```

Get dependencies by interface:
```go
services, err := di.GetByInterface[Worker]() // Worker is interface for many workers
```

### Go doc

```go
package di // import "go.neonxp.dev/di"

func Get[T any](id string) (*T, error)
func GetByInterface[Interface any]() ([]Interface, error)
func GetByType[T any]() ([]*T, error)
func Register[T any](id string, constructor func() (*T, error))
```

### Example

```go
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

// Do work ...
service, err := di.Get[ServiceB]("serviceB") // <- Get instantinated service B
if err != nil {
    panic(err)
}

service.DoStuff() // Output: Hello, world!


// Services ...
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

```

## License

GPLv3
