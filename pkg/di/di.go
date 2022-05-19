package di

import (
	"go.uber.org/dig"
	"sync"
)

var (
	instanceDI *dig.Container
	onceDI     sync.Once
	Providers  []interface{}
)



func GetDIInstance() *dig.Container {
	onceDI.Do(func() {
		c := dig.New()
		for _, v := range Providers {
			c.Provide(v)
		}
		instanceDI = c
	})
	return instanceDI
}

