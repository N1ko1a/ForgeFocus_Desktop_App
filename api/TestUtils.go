package main

import "fmt"

// Helper function to compare two slices of maps
func compareMapsSlice(a, b []map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !compareMaps(a[i], b[i]) {
			return false
		}
	}
	return true
}

// Helper function to compare two maps
func compareMaps(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if bVal, ok := b[k]; !ok || bVal != v {
			return false
		}
	}
	return true
}

func compareEvents(a, b *EventsSchema) bool {
	if a.Title != b.Title {
		return false
	} else if a.Date != b.Date {
		return false
	} else if a.FromDate != b.FromDate {
		return false
	} else if a.ToDate != b.ToDate {
		return false
	}

	return true
}

func compereEventsReturns(a, b []*EventsSchema) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !compareEvents(a[i], b[i]) { // Dereference pointers when calling compareButtons
			return false
		}
	}
	return true
}

func compareButtons(a, b *ButtonsSchema) bool {
	if a.Name != b.Name {
		return false
	}
	return true
}

func compereButtonReturns(a, b []*ButtonsSchema) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !compareButtons(a[i], b[i]) { // Dereference pointers when calling compareButtons
			return false
		}
	}
	return true
}
func compareTodos(a, b *TodoSchema) bool {
	if a.Content != b.Content {
		return false
	} else if a.Workspace != b.Workspace {
		return false
	} else if a.Completed != b.Completed {
		return false
	}

	return true
}

func compereTodosReturns(a, b []*TodoSchema) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !compareTodos(a[i], b[i]) { // Dereference pointers when calling compareButtons
			return false
		}
	}
	return true
}

// I need this so that i can dereference values to print them in the error
func getTodosAsString(todos []*TodoSchema) string {
	var todosStr string
	for _, todo := range todos {
		todosStr += fmt.Sprintf("{Content: %s, Workspace: %s, Completed: %t}\n", todo.Content, todo.Workspace, todo.Completed)
	}
	return todosStr
}

func getEventsAsString(events []*EventsSchema) string {
	var eventsStr string
	for _, event := range events {
		eventsStr += fmt.Sprintf("{Date: %v, Title: %s, FromDate: %v, ToDate: %v}\n", event.Date, event.Title, event.FromDate, event.ToDate)
	}
	return eventsStr
}
