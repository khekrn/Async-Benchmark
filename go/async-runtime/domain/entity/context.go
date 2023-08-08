package entity

type WorkflowContext struct {
	Variables map[string]any
}

func (w *WorkflowContext) SetVariable(key string, value any) {
	w.Variables[key] = value
}

func (w *WorkflowContext) HasVariable(key string) bool {
	result := false
	if _, ok := w.Variables[key]; ok {
		result = ok
	}

	return result
}

func (w *WorkflowContext) DeleteVariable(key string) {
	delete(w.Variables, key)
}
