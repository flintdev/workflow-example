package workflow2

import (
	workflowFramework "github.com/flintdev/workflow-engine/engine"
	"github.com/flintdev/workflow-engine/handler"
	"workflow-example/workflows/workflow2/steps/step1"
	"workflow-example/workflows/workflow2/steps/step2"
	"workflow-example/workflows/workflow2/steps/step3"
	"workflow-example/workflows/workflow2/steps/step4"
)

func Definition() workflowFramework.Workflow {
	w := ParseDefinition()
	return w
}

func Steps() map[string]func(handler handler.Handler) {
	StepFuncMap := map[string]func(handler handler.Handler){
		"step1": step1.Execute,
		"step2": step2.Execute,
		"step3": step3.Execute,
		"step4": step4.Execute,
	}
	return StepFuncMap
}

func Trigger(event workflowFramework.Event) bool {
	return TriggerCondition(event)
}

type FilePath string
