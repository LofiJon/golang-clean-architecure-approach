package task

type DeleteTask interface {
	Execute(id string) error
}
