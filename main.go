package main

import (
	workflowFramework "github.com/flintdev/workflow-engine/engine"
	"workflow-example/workflows"
	workflow1 "workflow-example/workflows/workflow1"
)

func main() {
	app := workflowFramework.CreateApp()
	app.RegisterWorkflow(workflow1.Definition, workflow1.Steps, workflow1.Trigger)
	app.RegisterConfig(workflows.ParseConfig)
	app.Start()
}
