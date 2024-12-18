package worker

func DoWork(counter *int) {
	c := *counter
	c++
	*counter = c
}
