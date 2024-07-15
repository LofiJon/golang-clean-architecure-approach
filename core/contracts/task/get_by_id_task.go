package task

type GetByIdTask interface {
	Execute(id string) error
}
