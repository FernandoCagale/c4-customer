package notify

type Notify interface {
	GetNotify() ([]string, error)
}
