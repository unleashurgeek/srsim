package handler

import "sort"

// Priority EventHandler that will execute listeners in order of priority (ascending)
type PriorityEventHandler[E event] struct {
	listeners priorityListeners[E]
}

// Emit an event to all subscribed listeners, in order of priority (ascending)
func (handler *PriorityEventHandler[E]) Emit(event E) {
	for _, listener := range handler.listeners {
		listener.listener(event)
	}
}

// Subscribe a listener to this event handler with the given priority. Listeners are executed
// in ascending order.
func (handler *PriorityEventHandler[E]) Subscribe(listener Listener[E], priority int) {
	pl := priorityListener[E]{listener: listener, priority: priority}
	handler.listeners = append(handler.listeners, pl)
	sort.Sort(handler.listeners)
}

type priorityListener[E event] struct {
	listener Listener[E]
	priority int
}

type priorityListeners[E event] []priorityListener[E]

func (a priorityListeners[E]) Len() int           { return len(a) }
func (a priorityListeners[E]) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a priorityListeners[E]) Less(i, j int) bool { return a[i].priority < a[j].priority }
