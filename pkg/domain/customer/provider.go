package customer

import "github.com/google/wire"

var Set = wire.NewSet(NewUseCase, NewInMemoryRepository, NewGormRepository, NewProducerKafka,
	wire.Bind(new(Repository), new(*InMemoryRepository)), // in InMemoryRepository
	//wire.Bind(new(Repository), new(*GormRepository)), // in PostgresRepository
	wire.Bind(new(Producer), new(*ProducerKafka)), // in Kafka
	wire.Bind(new(UseCase), new(*CustomerUseCase)))
