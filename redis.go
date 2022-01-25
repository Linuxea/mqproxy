package mqproxy

import (
	"github.com/go-redis/redis"
)

type RedisMq struct {
	redisCli *redis.Client
}

func (r *RedisMq) produce(dest string, data interface{}) error {
	r.redisCli.Publish(dest, data)
	return nil
}

func (r *RedisMq) require(dest, name, group string) (interface{}, error) {
	stream := dest
	subscribe := r.redisCli.Subscribe(stream)
	return (<-subscribe.Channel()).Payload, nil
}

func (r *RedisMq) consume(dest, name, group string, businessFunc func(data interface{}) error) error {
	b, err := r.require(dest, name, group)
	if err != nil {
		return err
	}

	return businessFunc(b)
}
