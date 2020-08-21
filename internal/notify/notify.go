package notify

type Notify interface {
	GetNotify(headers map[string]string) ([]string, error)
}
