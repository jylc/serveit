package base

import (
	"log"
	"sync/atomic"
)

const (
	ComponentPrepare = iota
	ComponentRunning
	ComponentStopped
)

// Component 组件
type Component struct {
	app    *Application
	name   string
	status uint32
}

type IComponent interface {
	Init()
	Name() string
	Run()
}

func (c *Component) Init() {
	c.status = ComponentPrepare
}

func (c *Component) Name() string {
	return c.name
}

func (c *Component) Run() {
	atomic.StoreUint32(&c.status, ComponentRunning)
	log.Printf("[INFO] component (%v) runnning", c.name)
}

func (c *Component) Stop() {
	atomic.StoreUint32(&c.status, ComponentStopped)
	log.Printf("[INFO] component (%v) stopped", c.name)

}
