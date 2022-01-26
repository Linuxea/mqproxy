package mqproxy

type Produce interface {
	produce(dest string, data interface{}) error
}

type Require interface {
	require(dest, name, group string) (interface{}, error)
}

type businessFunc func(data interface{}) error
type errHandler func(err error)

type Consume interface {
	consume(dest, name, group string, bf businessFunc, eh errHandler)
}

type MqServer interface {
	Produce
}

type MqClient interface {
	Require
	Consume
}
