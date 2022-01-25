package mqproxy

type Produce interface {
	produce(dest string, data interface{}) error
}

type Require interface {
	require(dest, name, group string) (interface{}, error)
}

type Consume interface {
	consume(dest, name, group string, businessFunc func(data interface{}) error) error
}

type MqServer interface {
	Produce
}

type MqClient interface {
	Require
	Consume
}
