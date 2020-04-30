package pointersanderrors

import "fmt"

//Stringer a interface to declare String()
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
