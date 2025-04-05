package oswrap

import "sync"

const MAX_ROUTINE_CONTEXT_MAP_SIZE = 10000

var routinesContexts map[int]RoutineContext = make(map[int]RoutineContext, MAX_ROUTINE_CONTEXT_MAP_SIZE)
var mutex sync.Mutex

type RoutineContext map[string]any

func createRoutineContext(go_routine_id int) (context RoutineContext) {
	context = RoutineContext{"go_routine_id": go_routine_id}
	routinesContexts[go_routine_id] = context
	return context
}

func RoutineContextWithValue(context RoutineContext, key string, value any) RoutineContext {
	context[key] = value
	return context
}

func DeleteRoutineContext() {
	mutex.Lock()
	defer mutex.Unlock()
	routine_id := GetGoRoutineId()
	delete(routinesContexts, routine_id)
}

func GetRoutineContext() map[string]any {
	mutex.Lock()
	defer mutex.Unlock()
	routine_id := GetGoRoutineId()
	routine_context, ok := routinesContexts[routine_id]
	if !ok {
		routine_context = createRoutineContext(routine_id)
	}
	return routine_context
}
