package console

type ConsoleImplementation struct {
}

type ConsoleInterface interface {
}

func New() ConsoleInterface {

	return &ConsoleImplementation{}

}
