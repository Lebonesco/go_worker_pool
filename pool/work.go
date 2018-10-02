package pool

type Work struct {
	ID	int
	Job string
	Func func(a string, b int) (interface{})
}