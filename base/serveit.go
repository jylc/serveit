package base

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Application 应用
type Application struct {
	components []IComponent
	profile    string
}

func NewApplication(opts ...Option) *Application {
	app := &Application{}
	app.components = make([]IComponent, 0)

	for _, opt := range opts {
		opt(app)
	}
	return nil
}

func (a *Application) Startup() error {

	sg := make(chan os.Signal, 1)
	signal.Notify(sg, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	select {
	case s := <-sg:
		log.Printf("[INFO] receive shutdown signal (%v)", s)

	}
	return nil
}

func (a *Application) Shutdown() {

}

// Register 注册组件
func (a *Application) Register(components ...IComponent) {
	if len(components) == 0 {
		log.Printf("[INFO] no components need to be registered")
		return
	}
	for _, component := range components {
		if a.Find(component.Name()) {
			log.Printf("[INFO] components (%v) has existed", component.Name())
		} else {
			a.components = append(a.components, component)
			component.Init()
			component.Run()
		}
	}
}

// Find 通过组件的名字查找其是否存在
func (a *Application) Find(name string) bool {
	for _, component := range a.components {
		if component.Name() == name {
			return true
		}
	}
	return false
}
