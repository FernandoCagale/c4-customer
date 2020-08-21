package customer

type Producer interface {
	Producer(topic string, payload interface{}) error
}
