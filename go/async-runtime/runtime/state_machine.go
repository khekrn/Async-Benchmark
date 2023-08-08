package runtime

import (
	"github.com/khekrn/async/domain/entity"
	"github.com/khekrn/async/domain/model"
)

type WorkflowTask interface {
	Process(worklowContext *entity.WorkflowContext)
}

type StateMachine struct {
	TaskList []string
}

func (st *StateMachine) Run(workflowRequest model.WorkflowRequest) {

}

func NewStateMachine() *StateMachine {
	return &StateMachine{}
}
