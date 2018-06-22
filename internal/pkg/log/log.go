package log

type Logger struct {
	buf    []byte     // for accumulating text to write
}

var someVar int

func init() {
	someVar = 1
}

func GetSomeVar() int {
	return someVar
}
