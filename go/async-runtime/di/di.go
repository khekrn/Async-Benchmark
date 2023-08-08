package di

import (
	"errors"

	"github.com/khekrn/async/runtime"
)

type ApplicationContext struct {
	instanceMap map[string]*runtime.WorkflowTask
}

func (a *ApplicationContext) RegisterInstance(name string, instance *runtime.WorkflowTask) {
	a.instanceMap[name] = instance
}

func (a *ApplicationContext) GetInstance(name string) (*runtime.WorkflowTask, error) {
	if instance, ok := a.instanceMap[name]; ok {
		return instance, nil
	} else {
		return nil, errors.New("instance not found for the given name " + name)
	}
}
