package worker

type Job struct {
	ID int
}

type Result struct {
	JobID  int
	Output string
}
