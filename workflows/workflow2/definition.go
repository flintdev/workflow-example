package workflow2

import (
	"encoding/json"
	"fmt"
	workflowFramework "github.com/flintdev/workflow-engine/engine"
)

func ParseDefinition() workflowFramework.Workflow {
	var w workflowFramework.Workflow
	definition := `{
	"name": "workflow2",
	"startAt": "step1",
	"trigger": {
		"model": "approval",
		"eventType": "ADDED"
	},
	"steps": {
		"step1": {
			"type": "automation",
			"nextSteps": [{
					"name": "step2",
					"condition": {
						"key": "$.workflow2.step1.field1.field2",
						"value": "test1",
						"operator": "="
					}
				},
				{
					"name": "step3",
					"condition": {
						"key": "step1.result",
						"value": "Failure",
						"operator": "="
					}
				}
			]
		},
		"step2": {
			"type": "automation",
			"nextSteps": [{
				"name": "step4"
			}]
		},
		"step3": {
			"type": "automation",
			"nextSteps": [{
				"name": "step4"
			}]
		},
		"step4": {
			"type": "automation",
			"nextSteps": [{
				"name": "end"
			}]
		},
		"end": {
			"nextSteps": []
		}
	}
}`
	err := json.Unmarshal([]byte(definition), &w)

	if err != nil {
		fmt.Println(err.Error())
	}
	return w

}