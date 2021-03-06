package views

import (
	"github.com/abdullin/omni/core/env"
	"github.com/abdullin/omni/core/spec"
	"github.com/abdullin/omni/lang"
	"github.com/abdullin/seq"
)

var useCases = []env.UseCaseFactory{
	given_no_tasks_get_inbox_returns_empty,
	given_inbox_task_get_inbox_returns_it,
	given_inbox_task_deleted_inbox_returns_nothing,
}

func given_no_tasks_get_inbox_returns_empty() *env.UseCase {

	return &env.UseCase{
		Name: "Given no tasks, GET /inbox returns nothing",
		When: spec.GetJSON("/views/inbox", nil),
		ThenResponse: spec.ReturnJSON(seq.Map{
			"tasks.length": 0,
		}),
	}
}

func given_inbox_task_get_inbox_returns_it() *env.UseCase {

	e1 := lang.NewTaskAdded(event(), task(), "Write a use case")
	e2 := lang.NewTaskMovedToInbox(event(), e1.TaskId)

	return &env.UseCase{
		Name:  "Given inbox task, GET /inbox returns it",
		Given: spec.GivenEvents(e1, e2),
		When:  spec.GetJSON("/views/inbox", nil),
		ThenResponse: spec.ReturnJSON(seq.Map{
			"tasks": seq.Map{
				"length": 1,
				"[0]": seq.Map{
					"name":   e1.Name,
					"taskId": e1.TaskId,
				},
			},
		}),
	}
}

func given_inbox_task_deleted_inbox_returns_nothing() *env.UseCase {

	e1 := lang.NewTaskAdded(event(), task(), "Write a use case")
	e2 := lang.NewTaskMovedToInbox(event(), e1.TaskId)
	e3 := lang.NewTaskRemoved(event(), e1.TaskId)

	return &env.UseCase{
		Name:  "Given inbox task that is deleted, GET /inbox returns nothing",
		Given: spec.GivenEvents(e1, e2, e3),
		When:  spec.GetJSON("/views/inbox", nil),
		ThenResponse: spec.ReturnJSON(seq.Map{
			"tasks": seq.Map{
				"length": 0,
			},
		}),
	}
}

func event() lang.EventId {
	return lang.NewEventId()
}
func task() lang.TaskId {
	return lang.NewTaskId()
}
