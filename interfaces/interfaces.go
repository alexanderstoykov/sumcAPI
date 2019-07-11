package interfaces

type ScheduleGenerator interface {
	GenerateSchedule(stop int) map[int][]int
}

type ScheduleProvider interface {
	CallAPI(stop int) []byte
}
