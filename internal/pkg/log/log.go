package log

type Logger struct {
	buf    []byte     // for accumulating text to write
}

func (l Logger) Log() {

}

var someVar int

func init() {
	someVar = 1
}

func GetSomeVar() int {
	return someVar
}

func New() Logger {
	buf := make([]byte, 3)
	return Logger{buf: buf}
}